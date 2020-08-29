/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"wingows/co"
	"wingows/win"
)

type (
	_ITaskbarList2Impl struct{ _ITaskbarListImpl }

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
	//
	// ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList2 struct{ _ITaskbarList2Impl }

	_ITaskbarList2Vtbl struct {
		_ITaskbarListVtbl
		MarkFullscreenWindow uintptr
	}
)

func (me *_ITaskbarList2Impl) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList2)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *_ITaskbarList2Impl) MarkFullscreenWindow(
	hwnd win.HWND, fFullScreen bool) {

	vTbl := (*_ITaskbarList2Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.MarkFullscreenWindow, 3, me.uintptr,
		uintptr(hwnd), boolToUintptr(fFullScreen))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList2.MarkFullscreenWindow").Error())
	}
}
