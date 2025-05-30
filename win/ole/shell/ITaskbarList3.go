//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [ITaskbarList3] COM interface.
//
// Implements [ole.ComObj] and [ole.ComResource].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var taskbl *shell.ITaskbarList3
//	ole.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList3]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct{ ITaskbarList2 }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList3) IID() co.IID {
	return co.IID_ITaskbarList3
}

// [RegisterTab] method.
//
// [RegisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).RegisterTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(hwndMDI))
	return utl.ErrorAsHResult(ret)
}

// [SetOverlayIcon] method.
//
// [SetOverlayIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(hWnd win.HWND, hIcon win.HICON, description string) error {
	description16 := wstr.NewBufWith[wstr.Stack20](description, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetOverlayIcon,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(hIcon),
		uintptr(description16.UnsafePtr()))
	return utl.ErrorAsHResult(ret)
}

// [SetProgressState] method.
//
// [SetProgressState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hWnd win.HWND, flags co.TBPF) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetProgressState,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(flags))
	return utl.ErrorAsHResult(ret)
}

// [SetProgressValue] method.
//
// [SetProgressValue]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(hWnd win.HWND, completed, total uint) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetProgressValue,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(completed),
		uintptr(total))
	return utl.ErrorAsHResult(ret)
}

// [SetTabActive] method.
//
// [SetTabActive]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetTabActive,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(hwndMDI))
	return utl.ErrorAsHResult(ret)
}

// [SetTabOrder] method.
//
// [SetTabOrder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetTabOrder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(hwndInsertBefore))
	return utl.ErrorAsHResult(ret)
}

// [SetThumbnailClip] method.
//
// [SetThumbnailClip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetThumbnailClip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rcClip)))
	return utl.ErrorAsHResult(ret)
}

// [SetThumbnailTooltip] method.
//
// [SetThumbnailTooltip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hWnd win.HWND, tip string) error {
	tip16 := wstr.NewBufWith[wstr.Stack20](tip, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetThumbnailTooltip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(tip16.UnsafePtr()))
	return utl.ErrorAsHResult(ret)
}

// [ThumbBarAddButtons] method.
//
// [ThumbBarAddButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(hWnd win.HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).ThumbBarAddButtons,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(uint32(len(buttons))),
		uintptr(unsafe.Pointer(&buttons[0])))
	return utl.ErrorAsHResult(ret)
}

// [ThumbBarSetImageList] method.
//
// [ThumbBarSetImageList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(hWnd win.HWND, hImgl win.HIMAGELIST) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).ThumbBarSetImageList,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(hImgl))
	return utl.ErrorAsHResult(ret)
}

// [ThumbBarUpdateButtons] method.
//
// [ThumbBarUpdateButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(hWnd win.HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).ThumbBarUpdateButtons,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(uint32(len(buttons))),
		uintptr(unsafe.Pointer(&buttons[0])))
	return utl.ErrorAsHResult(ret)
}

// [UnregisterTab] method.
//
// [UnregisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).UnregisterTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab))
	return utl.ErrorAsHResult(ret)
}

type _ITaskbarList3Vt struct {
	_ITaskbarList2Vt
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
