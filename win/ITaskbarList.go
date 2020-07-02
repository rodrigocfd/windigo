/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
)

type iTaskbarList struct {
	IUnknown
}

type iTaskbarListVtbl struct {
	iUnknownVtbl
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

//------------------------------------------------------------------------------

type iTaskbarList2 struct {
	iTaskbarList
}

type iTaskbarList2Vtbl struct {
	iTaskbarListVtbl
	MarkFullscreenWindow uintptr
}

//------------------------------------------------------------------------------

type ITaskbarList3 struct {
	iTaskbarList2
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

func CoCreateITaskbarList3() *ITaskbarList3 {
	iUnk := coCreateInstance(&co.Guid_ITaskbarList, &co.Guid_ITaskbarList3)
	return &ITaskbarList3{
		iTaskbarList2: iTaskbarList2{
			iTaskbarList: iTaskbarList{
				IUnknown: *iUnk,
			},
		},
	}
}

func (v *ITaskbarList3) SetProgressValue(hwnd HWND,
	ullCompleted, ullTotal uint64) {

	lpVtbl := (*iTaskbarList3Vtbl)(unsafe.Pointer(v.lpVtbl))
	ret, _, _ := syscall.Syscall6(lpVtbl.SetProgressValue, 4,
		uintptr(unsafe.Pointer(v)), uintptr(hwnd),
		uintptr(ullCompleted), uintptr(ullTotal),
		0, 0)
	if co.ERROR(ret) != co.ERROR_S_OK {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("ITaskbarList3.SetProgressValue failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (v *ITaskbarList3) SetProgressState(hwnd HWND, tbpFlags co.TBPF) {
	lpVtbl := (*iTaskbarList3Vtbl)(unsafe.Pointer(v.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.SetProgressState, 3,
		uintptr(unsafe.Pointer(v)), uintptr(hwnd), uintptr(tbpFlags))
	if co.ERROR(ret) != co.ERROR_S_OK {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("ITaskbarList3.SetProgressState failed: %d %s",
			lerr, lerr.Error()))
	}
}
