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

// [IDXGIAdapter] COM interface.
//
// [IDXGIAdapter]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgiadapter
type IDXGIAdapter struct{ IDXGIObject }

type _IDXGIAdapterVt struct {
	_IDXGIObjectVt
	EnumOutputs           uintptr
	GetDesc               uintptr
	CheckInterfaceSupport uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIAdapter) IID() *co.IID {
	return &codxgi.IID_IDXGIAdapter
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIAdapter) AddRef(releaser *win.OleReleaser) *IDXGIAdapter {
	return utl.OleNewFromAddRef[*IDXGIAdapter](me, releaser)
}

// [CheckInterfaceSupport] method.
//
// [CheckInterfaceSupport]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiadapter-checkinterfacesupport
func (me *IDXGIAdapter) CheckInterfaceSupport(pInterfaceName *co.GUID) (int, error) {
	var umdVersion int64
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIAdapterVt](me.Ppvt()).CheckInterfaceSupport,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pInterfaceName)),
		uintptr(unsafe.Pointer(&umdVersion)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(umdVersion), nil
	} else {
		return 0, hr
	}
}

// [EnumOutputs] method.
//
// [EnumOutputs]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiadapter-enumoutputs
func (me *IDXGIAdapter) EnumOutputs(releaser *win.OleReleaser) ([]*IDXGIOutput, error) {
	var index uint32
	var ppvtQueried uintptr
	var adapters []*IDXGIOutput

	for {
		ret, _, _ := syscall.SyscallN(
			utl.Vt[_IDXGIAdapterVt](me.Ppvt()).EnumOutputs,
			me.Ppvt(),
			uintptr(index),
			uintptr(unsafe.Pointer(&ppvtQueried)))

		if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
			pObj := utl.OleNew[*IDXGIOutput](ppvtQueried, releaser)
			adapters = append(adapters, pObj)
		} else if hr == codxgi.HRESULT_DXGI_ERROR_NOT_FOUND {
			return adapters, nil // no further adapters
		} else { // actual error
			return nil, hr
		}
	}
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiadapter-getdesc
func (me *IDXGIAdapter) GetDesc() (DXGI_ADAPTER_DESC, error) {
	return utl.OleCallReturnStruct[DXGI_ADAPTER_DESC](me,
		utl.Vt[_IDXGIAdapterVt](me.Ppvt()).GetDesc)
}
