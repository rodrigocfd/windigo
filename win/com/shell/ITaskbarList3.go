package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

type _ITaskbarList3Vtbl struct {
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

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct {
	ITaskbarList2 // Base ITaskbarList2 > ITaskbarList > IUnknown.
}

// Calls IUnknown.CoCreateInstance() to return ITaskbarList3.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateITaskbarList3(dwClsContext co.CLSCTX) ITaskbarList3 {
	iUnk := win.CoCreateInstance(
		CLSID.TaskbarList, nil, dwClsContext, IID.ITaskbarList3)
	return ITaskbarList3{
		ITaskbarList2{
			ITaskbarList{IUnknown: iUnk},
		},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).RegisterTab, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndMDI))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(
	hwnd win.HWND, hIcon win.HICON, pszDescription string) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetOverlayIcon, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszDescription))), 0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hwnd win.HWND, tbpFlags co.TBPF) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetProgressState, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(tbpFlags))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(
	hwnd win.HWND, ullCompleted, ullTotal uint64) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetProgressValue, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(ullCompleted), uintptr(ullTotal), 0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabActive, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndMDI))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabOrder, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hwnd win.HWND, prcClip *win.RECT) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetThumbnailClip, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(unsafe.Pointer(prcClip)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hwnd win.HWND, pszTip string) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetThumbnailTooltip, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszTip))))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(
	hwnd win.HWND, cButtons uint32, pButton *THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarAddButtons, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(cButtons), uintptr(unsafe.Pointer(pButton)), 0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(
	hwnd win.HWND, himl win.HIMAGELIST) {

	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarSetImageList, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(himl))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(
	hwnd win.HWND, cButtons uint32, pButton *THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarUpdateButtons, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(cButtons), uintptr(unsafe.Pointer(pButton)), 0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).UnregisterTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
