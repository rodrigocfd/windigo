/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api"
	"wingows/co"
)

var (
	globalBaseSubclassId  = uint32(0)  // incremented at each subclass installed
	globalSubclassProcPtr = uintptr(0) // necessary for RemoveWindowSubclass
)

// Base to all native child control types, like Button and Edit.
// Allows control subclassing.
type controlNativeBase struct {
	controlIdGuard
	hwnd       api.HWND
	msgs       windowDepotMsg
	subclassId uint32
}

func makeNativeControlBase(ctrlId co.ID) controlNativeBase {
	return controlNativeBase{
		controlIdGuard: makeCtrlIdGuard(ctrlId),
	}
}

// Returns the underlying HWND handle of this native control.
func (me *controlNativeBase) Hwnd() api.HWND {
	return me.hwnd
}

// Exposes all the control subclass methods that can be handled.
// The subclass will be installed if at least 1 message was added.
func (me *controlNativeBase) OnSubclassMsg() *windowDepotMsg {
	if me.Hwnd() != 0 {
		panic("Cannot add subclass message after the control was created.")
	}
	return &me.msgs
}

func (me *controlNativeBase) create(exStyle co.WS_EX, className, title string,
	style co.WS, x, y int32, width, height uint32, parent Window) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}

	me.hwnd = api.CreateWindowEx(exStyle, className, title, style,
		x, y, width, height, parent.Hwnd(), api.HMENU(me.controlIdGuard.CtrlId()),
		parent.Hwnd().GetInstance(), nil)

	if len(me.msgs.mapMsgs) > 0 || // at last 1 subclass message was added?
		len(me.msgs.mapCmds) > 0 {

		if globalSubclassProcPtr == 0 {
			globalSubclassProcPtr = syscall.NewCallback(subclassProc)
		}
		globalBaseSubclassId++
		me.subclassId = globalBaseSubclassId

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hwnd.SetWindowSubclass(globalSubclassProcPtr,
			me.subclassId, unsafe.Pointer(me))
	}
}

func subclassProc(hwnd api.HWND, msg co.WM,
	wParam api.WPARAM, lParam api.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	pMe := (*controlNativeBase)(unsafe.Pointer(dwRefData))

	// Save *nativeControlBase from being collected by GC; stored won't be used.
	hwnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	if pMe != nil && pMe.hwnd != 0 {
		userRet, wasHandled := pMe.msgs.processMessage(msg,
			wmBase{WParam: wParam, LParam: lParam}) // try to process the message with an user handler

		if msg == co.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			pMe.hwnd.RemoveWindowSubclass(globalSubclassProcPtr, pMe.subclassId)
		}
		if wasHandled && msg != co.WM_LBUTTONUP {
			// For some reason, if we don't call DefSubclassProc with WM_LBUTTONUP,
			// all parent window messages are routed to this proc, and it becomes
			// unresponsive. So we return user result only if not WM_LBUTTONUP.
			return userRet
		}
	} else if msg == co.WM_NCDESTROY { // https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		hwnd.RemoveWindowSubclass(globalSubclassProcPtr, pMe.subclassId)
	}

	return hwnd.DefSubclassProc(msg, wParam, lParam) // message was not processed
}
