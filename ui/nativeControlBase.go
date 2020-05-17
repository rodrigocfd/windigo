/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

var baseSubclassId = uint32(0)   // incremented at each subclass installed
var subclassProcPtr = uintptr(0) // necessary for RemoveWindowSubclass

// Base to all child control types.
type nativeControlBase struct {
	ctrlIdGuard
	hwnd        api.HWND
	subclassMsg windowMsg
	subclassId  uint32
}

func makeNativeControlBase(ctrlId c.ID) nativeControlBase {
	return nativeControlBase{
		ctrlIdGuard: makeCtrlIdGuard(ctrlId),
	}
}

// Returns the underlying HWND handle of this native control.
func (me *nativeControlBase) Hwnd() api.HWND {
	return me.hwnd
}

// Exposes all the control subclass methods that can be handled.
// The subclass will be installed if at least 1 message was added.
func (me *nativeControlBase) OnSubclassMsg() *windowMsg {
	return &me.subclassMsg
}

func (me *nativeControlBase) create(exStyle c.WS_EX, className, title string,
	style c.WS, x, y int32, width, height uint32, parent Window) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.hwnd = api.CreateWindowEx(exStyle, className, title, style,
		x, y, width, height, parent.Hwnd(), api.HMENU(me.ctrlIdGuard.CtrlId()),
		parent.Hwnd().GetInstance(), nil)

	if len(me.subclassMsg.msgs) > 0 || // at last 1 subclass message was added?
		len(me.subclassMsg.cmds) > 0 ||
		len(me.subclassMsg.nfys) > 0 {

		if subclassProcPtr == 0 {
			subclassProcPtr = syscall.NewCallback(subclassProc)
		}
		baseSubclassId++
		me.subclassId = baseSubclassId

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hwnd.SetWindowSubclass(subclassProcPtr,
			me.subclassId, unsafe.Pointer(me))
	}
}

func subclassProc(hwnd api.HWND, msg c.WM, wParam api.WPARAM, lParam api.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	pMe := (*nativeControlBase)(unsafe.Pointer(dwRefData))

	// Save *nativeControlBase from being collected by GC; stored won't be used.
	hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	if pMe != nil && pMe.hwnd != 0 {
		userResult, wasProcessed := pMe.subclassMsg.processMessage(msg,
			wmBase{WParam: wParam, LParam: lParam}) // try to process the message with an user handler

		if msg == c.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			pMe.hwnd.RemoveWindowSubclass(subclassProcPtr, pMe.subclassId)
		}
		if wasProcessed && msg != c.WM_LBUTTONUP {
			// For some reason, if we don't call DefSubclassProc with WM_LBUTTONUP,
			// all parent window messages are routed to this proc, and it becomes
			// unresponsive. So user return is not used.
			return userResult
		}
	} else if msg == c.WM_NCDESTROY { // https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		hwnd.RemoveWindowSubclass(subclassProcPtr, pMe.subclassId)
	}

	return hwnd.DefSubclassProc(msg, wParam, lParam) // message was not processed
}
