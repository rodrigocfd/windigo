//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [ITaskbarList3] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	taskbl, _ := ole.CoCreateInstance[shell.ITaskbarList3](
//		rel,
//		co.CLSID_TaskbarList,
//		co.CLSCTX_INPROC_SERVER,
//	)
//
// [ITaskbarList3]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct{ ITaskbarList2 }

// Returns the unique COM interface identifier.
func (*ITaskbarList3) IID() co.IID {
	return co.IID_ITaskbarList3
}

// [RegisterTab] method.
//
// [RegisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).RegisterTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab), uintptr(hwndMDI))
	return wutil.ErrorAsHResult(ret)
}

// [SetOverlayIcon] method.
//
// [SetOverlayIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(hWnd win.HWND, hIcon win.HICON, description string) error {
	description16 := wstr.NewBufWith[wstr.Stack20](description, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetOverlayIcon,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(hIcon), uintptr(description16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

// [SetProgressState] method.
//
// [SetProgressState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hWnd win.HWND, flags co.TBPF) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetProgressState,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(flags))
	return wutil.ErrorAsHResult(ret)
}

// [SetProgressValue] method.
//
// [SetProgressValue]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(hWnd win.HWND, completed, total uint) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetProgressValue,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(completed), uintptr(total))
	return wutil.ErrorAsHResult(ret)
}

// [SetTabActive] method.
//
// [SetTabActive]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetTabActive,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab), uintptr(hwndMDI))
	return wutil.ErrorAsHResult(ret)
}

// [SetTabOrder] method.
//
// [SetTabOrder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetTabOrder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab), uintptr(hwndInsertBefore))
	return wutil.ErrorAsHResult(ret)
}

// [SetThumbnailClip] method.
//
// [SetThumbnailClip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetThumbnailClip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(unsafe.Pointer(rcClip)))
	return wutil.ErrorAsHResult(ret)
}

// [SetThumbnailTooltip] method.
//
// [SetThumbnailTooltip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hwnd win.HWND, tip string) error {
	tip16 := wstr.NewBufWith[wstr.Stack20](tip, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).SetThumbnailTooltip,
		uintptr(unsafe.Pointer(me.Ppvt())), uintptr(hwnd), uintptr(tip16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

// [ThumbBarAddButtons] method.
//
// [ThumbBarAddButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(hWnd win.HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).ThumbBarAddButtons,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])))
	return wutil.ErrorAsHResult(ret)
}

// [ThumbBarSetImageList] method.
//
// [ThumbBarSetImageList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(hWnd win.HWND, hImgl win.HIMAGELIST) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).ThumbBarSetImageList,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(hImgl))
	return wutil.ErrorAsHResult(ret)
}

// [ThumbBarUpdateButtons] method.
//
// [ThumbBarUpdateButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(hWnd win.HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).ThumbBarUpdateButtons,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])))
	return wutil.ErrorAsHResult(ret)
}

// [UnregisterTab] method.
//
// [UnregisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList3)(unsafe.Pointer(*me.Ppvt())).UnregisterTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab))
	return wutil.ErrorAsHResult(ret)
}
