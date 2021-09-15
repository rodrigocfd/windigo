package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/errco"
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

// Calls CoCreateInstance(). Usually context is CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer ITaskbarList3.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func NewITaskbarList3(context co.CLSCTX) ITaskbarList3 {
	iUnk := win.CoCreateInstance(
		shellco.CLSID_TaskbarList, nil, context,
		shellco.IID_ITaskbarList3)
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

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(
	hWnd win.HWND, hIcon win.HICON, description string) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetOverlayIcon, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(description))), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hWnd win.HWND, flags shellco.TBPF) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetProgressState, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(
	hWnd win.HWND, completed, total uint64) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetProgressValue, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(completed), uintptr(total), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabActive, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabOrder, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetThumbnailClip, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(unsafe.Pointer(rcClip)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hwnd win.HWND, tip string) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).SetThumbnailTooltip, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), uintptr(unsafe.Pointer(win.Str.ToNativePtr(tip))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarAddButtons, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(
	hWnd win.HWND, hImgl win.HIMAGELIST) {

	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarSetImageList, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(hImgl))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).ThumbBarUpdateButtons, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList3Vtbl)(unsafe.Pointer(*me.Ppv)).UnregisterTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
