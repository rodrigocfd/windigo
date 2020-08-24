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
	_ITaskbarList2 struct{ _ITaskbarList }

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
	//
	// ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList2 struct{ _ITaskbarList2 }

	_ITaskbarList2Vtbl struct {
		_ITaskbarListVtbl
		MarkFullscreenWindow uintptr
	}
)

func (me *_ITaskbarList2) CoCreateInstance(dwClsContext co.CLSCTX) {
	me._IUnknown.coCreateInstance(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList2)
}

func (me *_ITaskbarList2) MarkFullscreenWindow(
	hwnd win.HWND, fFullScreen bool) {

	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList2Vtbl)(unsafe.Pointer(me.pVtb())).MarkFullscreenWindow, 3,
		uintptr(unsafe.Pointer(me.uintptr)), uintptr(hwnd),
		uintptr(boolToUintptr(fFullScreen)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("ITaskbarList2.MarkFullscreenWindow failed."))
	}
}
