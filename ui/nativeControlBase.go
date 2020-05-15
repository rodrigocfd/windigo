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

var baseSubclassId = uint32(0) // incremented at each subclass installed

// Base to all child control types.
type nativeControlBase struct {
	ctrlIdGuard
	hwnd        api.HWND
	subclassMsg windowMsg
	subclassId         uint32
	subclassProc uintptr
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

func (me *nativeControlBase) create(exStyle c.WS_EX, className, title string,
	style c.WS, x, y int32, width, height uint32, parent Window) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}

	me.hwnd = api.CreateWindowEx(exStyle, className, title, style,
		x, y, width, height, parent.Hwnd(), api.HMENU(me.ctrlIdGuard.CtrlId()),
		parent.Hwnd().GetInstance(), nil)

	if len(me.subclassMsg.msgs) > 0 || // at last 1 subclass message was added?
		len(me.subclassMsg.cmds) > 0 ||
		len(me.subclassMsg.nfys) > 0 {

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		baseSubclassId++
		me.hwnd.SetWindowSubclass(syscall.NewCallback(subclassProc),
			baseSubclassId, unsafe.Pointer(me))
	}
}

func subclassProc(hwnd api.HWND, msg c.WM, wParam api.WPARAM, lParam api.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	base := (*windowBase)(unsafe.Pointer(dwRefData))

	// Save *nativeControlBase from being collected by GC; stored won't be used.
	hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))

	if base != nil {
		if base.hwnd != 0 {

		} else {

		}
	} else if msg == c.WM_NCDESTROY {
		hwnd.RemoveWindowSubclass(subclassProc uintptr, uIdSubclass uint32)
	}

	return hwnd.DefSubclassProc(msg, wParam, lParam)
}
