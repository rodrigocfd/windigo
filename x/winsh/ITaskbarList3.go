//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [ITaskbarList3] COM interface.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var taskbl *winsh.ITaskbarList3
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList3]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct{ ITaskbarList2 }

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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList3) IID() *co.IID {
	return &cosh.IID_ITaskbarList3
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *ITaskbarList3) AddRef(releaser *win.OleReleaser) *ITaskbarList3 {
	return utl.OleNewFromAddRef[*ITaskbarList3](me, releaser)
}

// [RegisterTab] method.
//
// [RegisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).RegisterTab,
		me.Ppvt(),
		uintptr(hwndTab),
		uintptr(hwndMDI))
	return utl.HresultToError(ret)
}

// [SetOverlayIcon] method.
//
// [SetOverlayIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(hWnd win.HWND, hIcon win.HICON, description string) error {
	var wDescription wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetOverlayIcon,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(hIcon),
		uintptr(wDescription.AllowEmpty(description)))
	return utl.HresultToError(ret)
}

// [SetProgressState] method.
//
// [SetProgressState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hWnd win.HWND, flags cosh.TBPF) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetProgressState,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [SetProgressValue] method.
//
// Panics if completed or total is negative.
//
// [SetProgressValue]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(hWnd win.HWND, completed, total int) error {
	utl.PanicNeg(completed, total)
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetProgressValue,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(completed),
		uintptr(total))
	return utl.HresultToError(ret)
}

// [SetTabActive] method.
//
// [SetTabActive]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetTabActive,
		me.Ppvt(),
		uintptr(hwndTab),
		uintptr(hwndMDI))
	return utl.HresultToError(ret)
}

// [SetTabOrder] method.
//
// [SetTabOrder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetTabOrder,
		me.Ppvt(),
		uintptr(hwndTab),
		uintptr(hwndInsertBefore))
	return utl.HresultToError(ret)
}

// [SetThumbnailClip] method.
//
// [SetThumbnailClip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hWnd win.HWND, pRcClip *win.RECT) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetThumbnailClip,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pRcClip)))
	return utl.HresultToError(ret)
}

// [SetThumbnailTooltip] method.
//
// [SetThumbnailTooltip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hWnd win.HWND, tip string) error {
	var wTip wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).SetThumbnailTooltip,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(wTip.EmptyIsNil(tip)))
	return utl.HresultToError(ret)
}

// [ThumbBarAddButtons] method.
//
// [ThumbBarAddButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(hWnd win.HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).ThumbBarAddButtons,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(uint32(len(buttons))),
		uintptr(unsafe.Pointer(&buttons[0])))
	return utl.HresultToError(ret)
}

// [ThumbBarSetImageList] method.
//
// [ThumbBarSetImageList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(hWnd win.HWND, hImgl win.HIMAGELIST) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).ThumbBarSetImageList,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(hImgl))
	return utl.HresultToError(ret)
}

// [ThumbBarUpdateButtons] method.
//
// [ThumbBarUpdateButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(hWnd win.HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).ThumbBarUpdateButtons,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(uint32(len(buttons))),
		uintptr(unsafe.Pointer(&buttons[0])))
	return utl.HresultToError(ret)
}

// [UnregisterTab] method.
//
// [UnregisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList3Vt](me.Ppvt()).UnregisterTab,
		me.Ppvt(),
		uintptr(hwndTab))
	return utl.HresultToError(ret)
}
