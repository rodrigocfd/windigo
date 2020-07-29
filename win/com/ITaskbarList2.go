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

// ITaskbarList2 > ITaskbarList > IUnknown.
type ITaskbarList2 struct {
	ITaskbarList
}

type iTaskbarList2Vtbl struct {
	iTaskbarListVtbl
	MarkFullscreenWindow uintptr
}

func (me *ITaskbarList2) coCreateInstance() {
	if me.lpVtbl == 0 { // if not created yet
		me.IUnknown.coCreateInstance(
			&co.Guid_ITaskbarList, &co.Guid_ITaskbarList2)
	}
}

func (me *ITaskbarList2) MarkFullscreenWindow(
	hwnd win.HWND, fFullScreen bool) {

	me.coCreateInstance()
	lpVtbl := (*iTaskbarList2Vtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.MarkFullscreenWindow, 3,
		uintptr(unsafe.Pointer(me)), uintptr(hwnd),
		uintptr(boolToUintptr(fFullScreen)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("ITaskbarList2.MarkFullscreenWindow failed."))
	}
}
