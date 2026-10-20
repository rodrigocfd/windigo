//go:build windows

package windxgi

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIDisplayControl] COM interface.
//
// [IDXGIDisplayControl]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgidisplaycontrol
type IDXGIDisplayControl struct{ win.IUnknown }

type _IDXGIDisplayControlVt struct {
	utl.IUnknownVt
	IsStereoEnabled  uintptr
	SetStereoEnabled uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIDisplayControl) IID() *co.IID {
	return &codxgi.IID_IDXGIDisplayControl
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIDisplayControl) AddRef(releaser *win.OleReleaser) *IDXGIDisplayControl {
	return utl.OleNewFromAddRef[*IDXGIDisplayControl](me, releaser)
}

// [IsStereoEnabled] method.
//
// [IsStereoEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgidisplaycontrol-isstereoenabled
func (me *IDXGIDisplayControl) IsStereoEnabled() bool {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDisplayControlVt](me.Ppvt()).IsStereoEnabled,
		me.Ppvt())
	return ret != 0
}

// [SetStereoEnabled] method.
//
// [SetStereoEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgidisplaycontrol-setstereoenabled
func (me *IDXGIDisplayControl) SetStereoEnabled(enable bool) {
	_, _, _ = syscall.SyscallN(
		utl.Vt[_IDXGIDisplayControlVt](me.Ppvt()).SetStereoEnabled,
		me.Ppvt(),
		utl.BoolToUintptr(enable))
}
