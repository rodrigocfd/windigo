/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
	//
	// ITaskbarList > IUnknown.
	ITaskbarList struct{ _ITaskbarListImpl }

	_ITaskbarListImpl struct{ _IUnknownImpl }

	_ITaskbarListVtbl struct {
		_IUnknownVtbl
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *ITaskbarList) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		dwClsContext,
		_Win.NewGuid(0x56fdf342, 0xfd6d, 0x11d0, 0x958a_006097c9a090)) // IID_ITaskbarList
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *_ITaskbarListImpl) HrInit() {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	syscall.Syscall(vTbl.HrInit, 1, uintptr(me.ptr), 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *_ITaskbarListImpl) AddTab(hwnd HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.AddTab, 2, uintptr(me.ptr), uintptr(hwnd), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList.AddTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *_ITaskbarListImpl) DeleteTab(hwnd HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.DeleteTab, 2, uintptr(me.ptr), uintptr(hwnd), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList.DeleteTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *_ITaskbarListImpl) ActivateTab(hwnd HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.ActivateTab, 1, uintptr(me.ptr), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList.ActivateTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *_ITaskbarListImpl) SetActiveAlt(hwnd HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetActiveAlt, 1, uintptr(me.ptr), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList.SetActiveAlt").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
	//
	// ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList2 struct{ _ITaskbarList2Impl }

	_ITaskbarList2Impl struct{ _ITaskbarListImpl }

	_ITaskbarList2Vtbl struct {
		_ITaskbarListVtbl
		MarkFullscreenWindow uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *ITaskbarList2) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		dwClsContext,
		_Win.NewGuid(0x602d4995, 0xb13a, 0x429b, 0xa66e_1935e44f4317)) // IID_ITaskbarList2
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *_ITaskbarList2Impl) MarkFullscreenWindow(
	hwnd HWND, fFullScreen bool) {

	vTbl := (*_ITaskbarList2Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.MarkFullscreenWindow, 3, uintptr(me.ptr),
		uintptr(hwnd), _Win.BoolToUintptr(fFullScreen))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList2.MarkFullscreenWindow").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
	//
	// ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList3 struct{ _ITaskbarList3Impl }

	_ITaskbarList3Impl struct{ _ITaskbarList2Impl }

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

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ns-shobjidl_core-thumbbutton
	THUMBBUTTON struct {
		DwMask  co.THB
		IId     uint32
		IBitmap uint32
		HIcon   HICON
		SzTip   [260]uint16
		DwFlags co.THBF
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *ITaskbarList3) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		dwClsContext,
		_Win.NewGuid(0xea1afb91, 0x9e28, 0x4b86, 0x90e9_9e9f8a5eefaf)) // IID_ITaskbarList3
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *_ITaskbarList3Impl) SetProgressValue(
	hwnd HWND, ullCompleted, ullTotal uint64) {

	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.SetProgressValue, 4, uintptr(me.ptr),
		uintptr(hwnd), uintptr(ullCompleted), uintptr(ullTotal), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetProgressValue").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *_ITaskbarList3Impl) SetProgressState(hwnd HWND, tbpFlags co.TBPF) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetProgressState, 3, uintptr(me.ptr),
		uintptr(hwnd), uintptr(tbpFlags))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetProgressState").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *_ITaskbarList3Impl) RegisterTab(hwndTab, hwndMDI HWND) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.RegisterTab, 3, uintptr(me.ptr),
		uintptr(hwndTab), uintptr(hwndMDI))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.RegisterTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *_ITaskbarList3Impl) UnregisterTab(hwndTab HWND) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.UnregisterTab, 2, uintptr(me.ptr),
		uintptr(hwndTab), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.UnregisterTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *_ITaskbarList3Impl) SetTabOrder(hwndTab, hwndInsertBefore HWND) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetTabOrder, 3, uintptr(me.ptr),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetTabOrder").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *_ITaskbarList3Impl) SetTabActive(hwndTab, hwndMDI HWND) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.SetTabActive, 4, uintptr(me.ptr),
		uintptr(hwndTab), uintptr(hwndMDI), 0, 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetTabActive").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *_ITaskbarList3Impl) ThumbBarAddButtons(
	hwnd HWND, cButtons uint32, pButton *THUMBBUTTON) {

	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.ThumbBarAddButtons, 4, uintptr(me.ptr),
		uintptr(hwnd), uintptr(cButtons), uintptr(unsafe.Pointer(pButton)), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.ThumbBarAddButtons").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *_ITaskbarList3Impl) ThumbBarUpdateButtons(
	hwnd HWND, cButtons uint32, pButton *THUMBBUTTON) {

	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.ThumbBarUpdateButtons, 4, uintptr(me.ptr),
		uintptr(hwnd), uintptr(cButtons), uintptr(unsafe.Pointer(pButton)), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.ThumbBarUpdateButtons").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *_ITaskbarList3Impl) ThumbBarSetImageList(hwnd HWND, himl HIMAGELIST) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.ThumbBarSetImageList, 3, uintptr(me.ptr),
		uintptr(hwnd), uintptr(himl))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.ThumbBarSetImageList").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *_ITaskbarList3Impl) SetOverlayIcon(
	hwnd HWND, hIcon HICON, pszDescription string) {

	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.SetOverlayIcon, 4, uintptr(me.ptr),
		uintptr(hwnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pszDescription))), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetOverlayIcon").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *_ITaskbarList3Impl) SetThumbnailTooltip(hwnd HWND, pszTip string) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetThumbnailTooltip, 3, uintptr(me.ptr),
		uintptr(hwnd), uintptr(unsafe.Pointer(Str.ToUint16Ptr(pszTip))))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetThumbnailTooltip").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *_ITaskbarList3Impl) SetThumbnailClip(hwnd HWND, prcClip *RECT) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetThumbnailClip, 3, uintptr(me.ptr),
		uintptr(hwnd), uintptr(unsafe.Pointer(prcClip)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetThumbnailClip").Error())
	}
}
