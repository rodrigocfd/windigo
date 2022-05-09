package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 interface {
	ITaskbarList2

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
	RegisterTab(hwndTab, hwndMDI win.HWND)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
	SetOverlayIcon(hWnd win.HWND, hIcon win.HICON, description string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
	SetProgressState(hWnd win.HWND, flags shellco.TBPF)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
	SetProgressValue(hWnd win.HWND, completed, total uint64)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
	SetTabActive(hwndTab, hwndMDI win.HWND)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
	SetTabOrder(hwndTab, hwndInsertBefore win.HWND)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
	SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
	SetThumbnailTooltip(hwnd win.HWND, tip string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
	ThumbBarAddButtons(hWnd win.HWND, buttons []THUMBBUTTON)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
	ThumbBarSetImageList(hWnd win.HWND, hImgl win.HIMAGELIST)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
	ThumbBarUpdateButtons(hWnd win.HWND, buttons []THUMBBUTTON)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
	UnregisterTab(hwndTab win.HWND)
}

type _ITaskbarList3 struct{ ITaskbarList2 }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList3.Release().
//
// Example:
//
//		taskbl := shell.NewITaskbarList3(
//			com.CoCreateInstance(
//				shellco.CLSID_TaskbarList, nil,
//				comco.CLSCTX_INPROC_SERVER,
//				shellco.IID_ITaskbarList3),
//		)
//		defer taskbl.Release()
func NewITaskbarList3(base com.IUnknown) ITaskbarList3 {
	return &_ITaskbarList3{ITaskbarList2: NewITaskbarList2(base)}
}

func (me *_ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).RegisterTab, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetOverlayIcon(
	hWnd win.HWND, hIcon win.HICON, description string) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetOverlayIcon, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(description))), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetProgressState(hWnd win.HWND, flags shellco.TBPF) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetProgressState, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetProgressValue(
	hWnd win.HWND, completed, total uint64) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetProgressValue, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(completed), uintptr(total), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetTabActive, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetTabOrder, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetThumbnailClip, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(unsafe.Pointer(rcClip)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetThumbnailTooltip(hwnd win.HWND, tip string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetThumbnailTooltip, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwnd), uintptr(unsafe.Pointer(win.Str.ToNativePtr(tip))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) ThumbBarAddButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarAddButtons, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) ThumbBarSetImageList(
	hWnd win.HWND, hImgl win.HIMAGELIST) {

	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarSetImageList, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(hImgl))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) ThumbBarUpdateButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarUpdateButtons, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) UnregisterTab(hwndTab win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).UnregisterTab, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
