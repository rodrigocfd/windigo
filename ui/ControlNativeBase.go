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

// Base to all native child control types, like Button and Edit.
//
// Allows control subclassing.
type _ControlNativeBase struct {
	hwnd       win.HWND
	msgs       _DepotMsg
	subclassId uint32
}

// Returns the underlying HWND handle of this native control.
func (me *_ControlNativeBase) Hwnd() win.HWND {
	return me.hwnd
}

// Retrieves the command ID for this control.
func (me *_ControlNativeBase) Id() int {
	return int(me.hwnd.GetDlgCtrlID())
}

// Exposes all the control subclass methods that can be handled.
//
// The subclass will be installed in create() if at least 1 message was added.
func (me *_ControlNativeBase) OnSubclassMsg() *_DepotMsg {
	if me.hwnd != 0 {
		panic("Cannot add subclass message after the control was created.")
	}
	return &me.msgs
}

func (me *_ControlNativeBase) create(
	exStyle co.WS_EX, className, title string, style co.WS,
	x, y int, width, height uint, parent Window, ctrlId int) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}

	me.hwnd = win.CreateWindowEx(exStyle, className, title, style,
		int32(x), int32(y), int32(width), int32(height),
		parent.Hwnd(), win.HMENU(ctrlId), parent.Hwnd().GetInstance(), nil)

	if me.msgs.hasMessages() {
		if _globalSubclassProcPtr == 0 {
			_globalSubclassProcPtr = syscall.NewCallback(subclassProc)
		}
		_globalBaseSubclassId++
		me.subclassId = _globalBaseSubclassId

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hwnd.SetWindowSubclass(_globalSubclassProcPtr,
			me.subclassId, unsafe.Pointer(me))
	}
}

func subclassProc(hwnd win.HWND, msg co.WM,
	wParam win.WPARAM, lParam win.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	pMe := (*_ControlNativeBase)(unsafe.Pointer(dwRefData))

	// If the retrieved *nativeControlBase stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	if pMe != nil && pMe.hwnd != 0 {
		userRet, wasHandled := pMe.msgs.processMessage(msg, // try to process the message with an user handler
			Wm{WParam: wParam, LParam: lParam})

		if msg == co.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			pMe.hwnd.RemoveWindowSubclass(_globalSubclassProcPtr, pMe.subclassId)
		}
		if wasHandled && msg != co.WM_LBUTTONUP {
			// For some reason, if we don't call DefSubclassProc with WM_LBUTTONUP,
			// all parent window messages are routed to this proc, and it becomes
			// unresponsive. So we return user result only if not WM_LBUTTONUP.
			return userRet
		}
	} else if msg == co.WM_NCDESTROY { // https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		hwnd.RemoveWindowSubclass(_globalSubclassProcPtr, pMe.subclassId)
	}

	return hwnd.DefSubclassProc(msg, wParam, lParam) // message was not processed
}
