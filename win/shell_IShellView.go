//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IShellView] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IShellView]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellview
type IShellView struct{ IOleWindow }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellView) IID() co.IID {
	return co.IID_IShellView
}

// [DestroyViewWindow] method.
//
// [DestroyViewWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-destroyviewwindow
func (me *IShellView) DestroyViewWindow() error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).DestroyViewWindow,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [EnableModeless] method.
//
// [EnableModeless]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-enablemodeless
func (me *IShellView) EnableModeless(enable bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).EnableModeless,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(enable))
	return utl.ErrorAsHResult(ret)
}

// [Refresh] method.
//
// [Refresh]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-refresh
func (me *IShellView) Refresh() error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).Refresh,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SaveViewState] method.
//
// [SaveViewState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-saveviewstate
func (me *IShellView) SaveViewState() error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).SaveViewState,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [TranslateAccelerator] method.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-translateaccelerator
func (me *IShellView) TranslateAccelerator(msg *MSG) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).TranslateAccelerator,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(msg)))
	return utl.ErrorAsHResult(ret)
}

// [UIActivate] method.
//
// [UIActivate]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-uiactivate
func (me *IShellView) UIActivate(state co.SVUIA) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).UIActivate,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(state))
	return utl.ErrorAsHResult(ret)
}

type _IShellViewVt struct {
	_IOleWindowVt
	TranslateAccelerator  uintptr
	EnableModeless        uintptr
	UIActivate            uintptr
	Refresh               uintptr
	CreateViewWindow      uintptr
	DestroyViewWindow     uintptr
	GetCurrentInfo        uintptr
	AddPropertySheetPages uintptr
	SaveViewState         uintptr
	SelectItem            uintptr
	GetItemObject         uintptr
}
