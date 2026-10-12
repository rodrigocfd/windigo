//go:build windows

package windxgi

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIAdapter2] COM interface.
//
// [IDXGIAdapter2]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgiadapter2
type IDXGIAdapter2 struct{ IDXGIAdapter1 }

type _IDXGIAdapter2Vt struct {
	_IDXGIAdapter1Vt
	GetDesc2 uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIAdapter2) IID() *co.IID {
	return &codxgi.IID_IDXGIAdapter2
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIAdapter2) AddRef(releaser *win.OleReleaser) *IDXGIAdapter2 {
	return utl.OleNewFromAddRef[*IDXGIAdapter2](me, releaser)
}

// [GetDesc2] method.
//
// [GetDesc2]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiadapter2-getdesc2
func (me *IDXGIAdapter2) GetDesc2() (DXGI_ADAPTER_DESC2, error) {
	return utl.OleCallReturnStruct[DXGI_ADAPTER_DESC2](me,
		utl.Vt[_IDXGIAdapter2Vt](me.Ppvt()).GetDesc2)
}
