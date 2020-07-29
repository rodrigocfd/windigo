/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

type (
	baseITaskbarList2 struct{ baseITaskbarList }

	// ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList2 struct{ baseITaskbarList2 }

	vtbITaskbarList2 struct {
		tvbITaskbarList
		MarkFullscreenWindow uintptr
	}
)

func (me *baseITaskbarList2) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.baseIUnknown.coCreateInstance(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList2)
}

func (me *baseITaskbarList2) MarkFullscreenWindow(
	hwnd win.HWND, fFullScreen bool) {

	ret, _, _ := syscall.Syscall(
		(*vtbITaskbarList2)(unsafe.Pointer(me.pVtb())).MarkFullscreenWindow, 3,
		uintptr(unsafe.Pointer(me.uintptr)), uintptr(hwnd),
		uintptr(boolToUintptr(fFullScreen)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("ITaskbarList2.MarkFullscreenWindow failed."))
	}
}
