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
	_ITaskbarList3Impl struct{ _ITaskbarList2Impl }

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
	//
	// ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList3 struct{ _ITaskbarList3Impl }

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

func (me *_ITaskbarList3Impl) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList3)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(
	hwnd win.HWND, ullCompleted, ullTotal uint64) {

	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.SetProgressValue, 4, me.uintptr,
		uintptr(hwnd), uintptr(ullCompleted), uintptr(ullTotal), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetProgressValue").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hwnd win.HWND, tbpFlags co.TBPF) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetProgressState, 3, me.uintptr,
		uintptr(hwnd), uintptr(tbpFlags))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetProgressState").Error())
	}
}
