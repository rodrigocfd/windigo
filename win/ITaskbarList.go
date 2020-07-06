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

func (me *ITaskbarList3) coCreateInstance() {
	if me.lpVtbl == 0 { // if not created yet
		me.IUnknown.coCreateInstance(
			&co.Guid_ITaskbarList, &co.Guid_ITaskbarList3)
	}
}

func (me *ITaskbarList3) SetProgressValue(hwnd HWND,
	ullCompleted, ullTotal uint64) {

	me.coCreateInstance()
	lpVtbl := (*iTaskbarList3Vtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall6(lpVtbl.SetProgressValue, 4,
		uintptr(unsafe.Pointer(me)), uintptr(hwnd),
		uintptr(ullCompleted), uintptr(ullTotal),
		0, 0)
	if co.ERROR(ret) != co.ERROR_S_OK {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("ITaskbarList3.SetProgressValue failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (me *ITaskbarList3) SetProgressState(hwnd HWND, tbpFlags co.TBPF) {
	me.coCreateInstance()
	lpVtbl := (*iTaskbarList3Vtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.SetProgressState, 3,
		uintptr(unsafe.Pointer(me)), uintptr(hwnd), uintptr(tbpFlags))
	if co.ERROR(ret) != co.ERROR_S_OK {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("ITaskbarList3.SetProgressState failed: %d %s",
			lerr, lerr.Error()))
	}
}
