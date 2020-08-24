/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

type (
	_ITaskbarList3 struct{ _ITaskbarList2 }

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
	//
	// ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList3 struct{ _ITaskbarList3 }

	_ITaskbarList3Vtbl struct {
		_ITaskbarList2Vtbl
		SetProgressValue      uintptr
		SetProgressState      uintptr
		RegisterTab           uintptr
		UnregisterTab         uintptr
		SetTabOrder           uintptr
		SetTabActive          uintptr
		ThumbBarAddButtons    uintptr
		ThumbBarUpdateButtons uintptr
		ThumbBarSetImageList  uintptr
		SetOverlayIcon        uintptr
		SetThumbnailTooltip   uintptr
		SetThumbnailClip      uintptr
	}
)

func (me *_ITaskbarList3) CoCreateInstance(dwClsContext co.CLSCTX) {
	me._IUnknown.coCreateInstance(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList3)
}

func (me *ITaskbarList3) SetProgressValue(
	hwnd win.HWND, ullCompleted, ullTotal uint64) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(me.pVtb())).SetProgressValue, 4,
		uintptr(unsafe.Pointer(me.uintptr)), uintptr(hwnd),
		uintptr(ullCompleted), uintptr(ullTotal),
		0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(fmt.Sprintf("ITaskbarList3.SetProgressValue failed. %s",
			lerr.Error()))
	}
}

func (me *ITaskbarList3) SetProgressState(hwnd win.HWND, tbpFlags co.TBPF) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(me.pVtb())).SetProgressState, 3,
		uintptr(unsafe.Pointer(me.uintptr)), uintptr(hwnd), uintptr(tbpFlags))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(fmt.Sprintf("ITaskbarList3.SetProgressState failed. %s",
			lerr.Error()))
	}
}
