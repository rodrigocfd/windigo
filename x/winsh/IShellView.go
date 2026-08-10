//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IShellView] COM interface.
//
// [IShellView]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellview
type IShellView struct{ IOleWindow }

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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellView) IID() *co.IID {
	return &cosh.IID_IShellView
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IShellView) AddRef(releaser *win.OleReleaser) *IShellView {
	return utl.OleNewFromAddRef[*IShellView](me, releaser)
}

// [DestroyViewWindow] method.
//
// [DestroyViewWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-destroyviewwindow
func (me *IShellView) DestroyViewWindow() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IShellViewVt](me.Ppvt()).DestroyViewWindow)
}

// [EnableModeless] method.
//
// [EnableModeless]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-enablemodeless
func (me *IShellView) EnableModeless(enable bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellViewVt](me.Ppvt()).EnableModeless,
		me.Ppvt(),
		utl.BoolToUintptr(enable))
	return utl.HresultToError(ret)
}

// [Refresh] method.
//
// [Refresh]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-refresh
func (me *IShellView) Refresh() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IShellViewVt](me.Ppvt()).Refresh)
}

// [SaveViewState] method.
//
// [SaveViewState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-saveviewstate
func (me *IShellView) SaveViewState() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IShellViewVt](me.Ppvt()).SaveViewState)
}

// [TranslateAccelerator] method.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-translateaccelerator
func (me *IShellView) TranslateAccelerator(pMsg *win.MSG) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellViewVt](me.Ppvt()).TranslateAccelerator,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pMsg)))
	return utl.HresultToError(ret)
}

// [UIActivate] method.
//
// [UIActivate]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-uiactivate
func (me *IShellView) UIActivate(state cosh.SVUIA) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellViewVt](me.Ppvt()).UIActivate,
		me.Ppvt(),
		uintptr(state))
	return utl.HresultToError(ret)
}
