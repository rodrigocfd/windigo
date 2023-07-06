//go:build windows

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

// [ITaskbarList3] COM interface.
//
// [ITaskbarList3]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 interface {
	ITaskbarList2

	// [RegisterTab] COM method.
	//
	// [RegisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
	RegisterTab(hwndTab, hwndMDI win.HWND)

	// [SetOverlayIcon] COM method.
	//
	// [SetOverlayIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
	SetOverlayIcon(hWnd win.HWND, hIcon win.HICON, description string)

	// [SetProgressState] COM method.
	//
	// [SetProgressState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
	SetProgressState(hWnd win.HWND, flags shellco.TBPF)

	// [SetProgressValue] COM method.
	//
	// [SetProgressValue]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
	SetProgressValue(hWnd win.HWND, completed, total uint64)

	// [SetTabActive] COM method.
	//
	// [SetTabActive]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
	SetTabActive(hwndTab, hwndMDI win.HWND)

	// [SetTabOrder] COM method.
	//
	// [SetTabOrder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
	SetTabOrder(hwndTab, hwndInsertBefore win.HWND)

	// [SetThumbnailClip] COM method.
	//
	// [SetThumbnailClip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
	SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT)

	// [SetThumbnailTooltip] COM method.
	//
	// [SetThumbnailTooltip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
	SetThumbnailTooltip(hwnd win.HWND, tip string)

	// [ThumbBarAddButtons] COM method.
	//
	// [ThumbBarAddButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
	ThumbBarAddButtons(hWnd win.HWND, buttons []THUMBBUTTON)

	// [ThumbBarSetImageList] COM method.
	//
	// [ThumbBarSetImageList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
	ThumbBarSetImageList(hWnd win.HWND, hImgl win.HIMAGELIST)

	// [ThumbBarUpdateButtons] COM method.
	//
	// [ThumbBarUpdateButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
	ThumbBarUpdateButtons(hWnd win.HWND, buttons []THUMBBUTTON)

	// [UnregisterTab] COM method.
	//
	// [UnregisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
	UnregisterTab(hwndTab win.HWND)
}

type _ITaskbarList3 struct{ ITaskbarList2 }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList3.Release().
//
// # Example
//
//	taskbl := shell.NewITaskbarList3(
//		com.CoCreateInstance(
//			shellco.CLSID_TaskbarList, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			shellco.IID_ITaskbarList3),
//	)
//	defer taskbl.Release()
func NewITaskbarList3(base com.IUnknown) ITaskbarList3 {
	return &_ITaskbarList3{ITaskbarList2: NewITaskbarList2(base)}
}

func (me *_ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).RegisterTab,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetOverlayIcon(
	hWnd win.HWND, hIcon win.HICON, description string) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetOverlayIcon,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(description))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetProgressState(hWnd win.HWND, flags shellco.TBPF) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetProgressState,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetProgressValue(
	hWnd win.HWND, completed, total uint64) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetProgressValue,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(completed), uintptr(total))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetTabActive,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetTabOrder,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetThumbnailClip,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(unsafe.Pointer(rcClip)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) SetThumbnailTooltip(hwnd win.HWND, tip string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetThumbnailTooltip,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwnd), uintptr(unsafe.Pointer(win.Str.ToNativePtr(tip))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) ThumbBarAddButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarAddButtons,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) ThumbBarSetImageList(
	hWnd win.HWND, hImgl win.HIMAGELIST) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarSetImageList,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(hImgl))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) ThumbBarUpdateButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarUpdateButtons,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList3) UnregisterTab(hwndTab win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).UnregisterTab,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
