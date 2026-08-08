//go:build windows

package windxgi

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIAdapter1] COM interface.
//
// Implements [OleResource].
//
// [IDXGIAdapter1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgiadapter1
type IDXGIAdapter1 struct{ IDXGIAdapter }

type _IDXGIAdapter1Vt struct {
	_IDXGIAdapterVt
	GetDesc1 uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIAdapter1) IID() *co.IID {
	return &codxgi.IID_IDXGIAdapter1
}

// [GetDesc1] method.
//
// [GetDesc1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiadapter1-getdesc1
func (me *IDXGIAdapter1) GetDesc1() (DXGI_ADAPTER_DESC1, error) {
	return utl.OleCallReturnStruct[DXGI_ADAPTER_DESC1](me,
		utl.Vt[_IDXGIAdapter1Vt](me.Ppvt()).GetDesc1)
}
