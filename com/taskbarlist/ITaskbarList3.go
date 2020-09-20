/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package taskbarlist

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
	//
	// ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList3 struct{ ITaskbarList2 }

	ITaskbarList3Vtbl struct {
		ITaskbarList2Vtbl
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

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *ITaskbarList3) CoCreateInstance(dwClsContext co.CLSCTX) *ITaskbarList3 {
	ppv, err := win.CoCreateInstance(
		win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		nil,
		dwClsContext,
		win.NewGuid(0xea1afb91, 0x9e28, 0x4b86, 0x90e9_9e9f8a5eefaf)) // IID_ITaskbarList3

	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "CoCreateInstance/ITaskbarList2"))
	}
	me.Ppv = ppv
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(
	hwnd win.HWND, ullCompleted, ullTotal uint64) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall6(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetProgressValue, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(ullCompleted), uintptr(ullTotal), 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetProgressValue").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(
	hwnd win.HWND, tbpFlags TBPF) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetProgressState, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(tbpFlags))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetProgressState").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) *ITaskbarList3 {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).RegisterTab, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndMDI))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.RegisterTab").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) *ITaskbarList3 {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).UnregisterTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.UnregisterTab").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(
	hwndTab, hwndInsertBefore win.HWND) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabOrder, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetTabOrder").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(
	hwndTab, hwndMDI win.HWND) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall6(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabActive, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndMDI), 0, 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetTabActive").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(
	hwnd win.HWND, cButtons uint32, pButton *THUMBBUTTON) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall6(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarAddButtons, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(cButtons), uintptr(unsafe.Pointer(pButton)), 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.ThumbBarAddButtons").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(
	hwnd win.HWND, cButtons uint32, pButton *THUMBBUTTON) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall6(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarUpdateButtons, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(cButtons), uintptr(unsafe.Pointer(pButton)), 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.ThumbBarUpdateButtons").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(
	hwnd win.HWND, himl win.HIMAGELIST) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarSetImageList, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(himl))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.ThumbBarSetImageList").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(
	hwnd win.HWND, hIcon win.HICON, pszDescription string) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall6(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetOverlayIcon, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszDescription))), 0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetOverlayIcon").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(
	hwnd win.HWND, pszTip string) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetThumbnailTooltip, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszTip))))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetThumbnailTooltip").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(
	hwnd win.HWND, prcClip *win.RECT) *ITaskbarList3 {

	ret, _, _ := syscall.Syscall(
		(*ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetThumbnailClip, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(unsafe.Pointer(prcClip)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList3.SetThumbnailClip").Error())
	}
	return me
}
