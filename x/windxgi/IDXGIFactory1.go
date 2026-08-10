//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIFactory1] COM interface.
//
// [IDXGIFactory1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgifactory1
type IDXGIFactory1 struct{ IDXGIFactory }

type _IDXGIFactory1Vt struct {
	_IDXGIFactoryVt
	EnumAdapters1 uintptr
	IsCurrent     uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIFactory1) IID() *co.IID {
	return &codxgi.IID_IDXGIFactory1
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIFactory1) AddRef(releaser *win.OleReleaser) *IDXGIFactory1 {
	return utl.OleNewFromAddRef[*IDXGIFactory1](me, releaser)
}

// [EnumAdapters1] method.
//
// [EnumAdapters1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-enumadapters
func (me *IDXGIFactory1) EnumAdapters1(releaser *win.OleReleaser) ([]*IDXGIAdapter1, error) {
	var index uint32
	var ppvtQueried uintptr
	var adapters []*IDXGIAdapter1

	for {
		ret, _, _ := syscall.SyscallN(
			utl.Vt[_IDXGIFactory1Vt](me.Ppvt()).EnumAdapters1,
			me.Ppvt(),
			uintptr(index),
			uintptr(unsafe.Pointer(&ppvtQueried)))

		if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
			pObj := utl.OleNew[*IDXGIAdapter1](ppvtQueried, releaser)
			adapters = append(adapters, pObj)
		} else if hr == codxgi.HRESULT_DXGI_ERROR_NOT_FOUND {
			return adapters, nil // no further adapters
		} else { // actual error
			return nil, hr
		}
	}
}

// [IsCurrent] method.
//
// [IsCurrent]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory1-iscurrent
func (me *IDXGIFactory1) IsCurrent() bool {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory1Vt](me.Ppvt()).IsCurrent,
		me.Ppvt())
	return ret != 0
}
