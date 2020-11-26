/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

var (
	_globalBaseSubclassId  = uint32(0)  // incremented at each subclass installed
	_globalSubclassProcPtr = uintptr(0) // necessary for RemoveWindowSubclass
)

//------------------------------------------------------------------------------

// Base to all native child controls.
type _NativeControlBase struct {
	hwnd         win.HWND
	ctrlId       int
	parent       Parent
	eventsWmSubc *_EventsWm // for subclassing
	subclassId   uint32
}

// Constructor.
func _NewNativeControlBase(parent Parent, ctrlId ...int) *_NativeControlBase {
	ourCtrlId := 0

	switch numCtrlId := len(ctrlId); {
	case numCtrlId > 1:
		panic("You cannot pass more than 1 control ID.")
	case numCtrlId == 1:
		ourCtrlId = ctrlId[0]
	default:
		ourCtrlId = _global.NewAutoCtrlId()
	}

	return &_NativeControlBase{
		ctrlId:       ourCtrlId,
		parent:       parent,
		eventsWmSubc: _NewEventsWm(),
	}
}

// Returns the underlying HWND handle of this window.
func (me *_NativeControlBase) Hwnd() win.HWND {
	return me.hwnd
}

// Returns the control ID.
func (me *_NativeControlBase) CtrlId() int {
	return me.ctrlId
}

// Exposes all the window messages that can be handled with subclassing.
//
// Cannot be called after the control was created.
func (me *_NativeControlBase) OnSubclass() *_EventsWm {
	if me.hwnd != 0 {
		panic("Cannot subclass after the control was created.")
	}
	return me.eventsWmSubc
}

// Calls CreateWindowEx(), installs subclass if needed.
func (me *_NativeControlBase) create(
	className, title string, pos Pos, size Size, style co.WS, exStyle co.WS_EX) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Attempting to create %s twice.", className))
	}

	me.hwnd = win.CreateWindowEx(exStyle, className, title, style,
		int32(pos.X), int32(pos.Y), int32(size.Cx), int32(size.Cy),
		me.parent.Hwnd(), win.HMENU(me.ctrlId), me.parent.Hwnd().GetInstance(),
		nil)

	me.installSubclassIfNeeded()
}

// Calls GetDlgItem() to retrieve the control HWND, installs subclass if needed.
func (me *_NativeControlBase) createAssignDlg() {
	if me.hwnd != 0 {
		panic(fmt.Sprintf("Control was already created in dialog."))
	}
	me.hwnd = me.parent.Hwnd().GetDlgItem(int32(me.ctrlId))
	me.installSubclassIfNeeded()
}

func (me *_NativeControlBase) installSubclassIfNeeded() {
	if me.eventsWmSubc.hasMessages() {
		if _globalSubclassProcPtr == 0 {
			_globalSubclassProcPtr = syscall.NewCallback(_globalSubclassProc)
		}
		_globalBaseSubclassId++
		me.subclassId = _globalBaseSubclassId

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hwnd.SetWindowSubclass(_globalSubclassProcPtr,
			me.subclassId, unsafe.Pointer(me))
	}
}

// WNDPROC for subclassing all child controls.
func _globalSubclassProc(
	hwnd win.HWND, msg co.WM, wParam win.WPARAM, lParam win.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	pMe := (*_NativeControlBase)(unsafe.Pointer(dwRefData))

	// If the retrieved *_ControlNativeBase stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	if pMe != nil && pMe.hwnd != 0 {
		retVal, useRetVal, wasHandled := pMe.eventsWmSubc.processMessage( // try to process the message with an user handler
			msg, Wm{WParam: wParam, LParam: lParam})

		if msg == co.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			pMe.hwnd.RemoveWindowSubclass(_globalSubclassProcPtr, pMe.subclassId)
		}
		if wasHandled && msg != co.WM_LBUTTONUP {
			// For some reason, if we don't call DefSubclassProc with WM_LBUTTONUP,
			// all parent window messages are routed to this proc, and it becomes
			// unresponsive. So we return user result only if not WM_LBUTTONUP.
			if useRetVal {
				return retVal
			}
			return 0
		}
	} else if msg == co.WM_NCDESTROY { // https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		hwnd.RemoveWindowSubclass(_globalSubclassProcPtr, pMe.subclassId)
	}

	return hwnd.DefSubclassProc(msg, wParam, lParam) // message was not processed
}
