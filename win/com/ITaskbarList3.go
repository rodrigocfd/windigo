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

// ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
type ITaskbarList3 struct {
	ITaskbarList2
}

type iTaskbarList3Vtbl struct {
	iTaskbarList2Vtbl
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

func (me *ITaskbarList3) coCreateInstance() {
	if me.lpVtbl == 0 { // if not created yet
		me.IUnknown.coCreateInstance(
			&co.Guid_ITaskbarList, &co.Guid_ITaskbarList3)
	}
}

func (me *ITaskbarList3) SetProgressValue(
	hwnd win.HWND, ullCompleted, ullTotal uint64) {

	me.coCreateInstance()
	lpVtbl := (*iTaskbarList3Vtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall6(lpVtbl.SetProgressValue, 4,
		uintptr(unsafe.Pointer(me)), uintptr(hwnd),
		uintptr(ullCompleted), uintptr(ullTotal),
		0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("ITaskbarList3.SetProgressValue failed."))
	}
}

func (me *ITaskbarList3) SetProgressState(hwnd win.HWND, tbpFlags co.TBPF) {
	me.coCreateInstance()
	lpVtbl := (*iTaskbarList3Vtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.SetProgressState, 3,
		uintptr(unsafe.Pointer(me)), uintptr(hwnd), uintptr(tbpFlags))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("ITaskbarList3.SetProgressState failed."))
	}
}
