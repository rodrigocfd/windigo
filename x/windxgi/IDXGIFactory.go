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

// [IDXGIFactory] COM interface.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	factory, _ := win.CreateDXGIFactory(rel)
//
// [IDXGIFactory]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgifactory
type IDXGIFactory struct{ IDXGIObject }

type _IDXGIFactoryVt struct {
	_IDXGIObjectVt
	EnumAdapters          uintptr
	MakeWindowAssociation uintptr
	GetWindowAssociation  uintptr
	CreateSwapChain       uintptr
	CreateSoftwareAdapter uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIFactory) IID() *co.IID {
	return &codxgi.IID_IDXGIFactory
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIFactory) AddRef(releaser *win.OleReleaser) *IDXGIFactory {
	return utl.OleNewFromAddRef[*IDXGIFactory](me, releaser)
}

// [CreateSoftwareAdapter] method.
//
// [CreateSoftwareAdapter]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-createsoftwareadapter
func (me *IDXGIFactory) CreateSoftwareAdapter(
	releaser *win.OleReleaser,
	hModule win.HINSTANCE,
) (*IDXGIAdapter, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactoryVt](me.Ppvt()).CreateSoftwareAdapter,
		me.Ppvt(),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGIAdapter](ret, ppvtQueried, releaser)
}

// [CreateSwapChain] method.
//
// [CreateSwapChain]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-createswapchain
func (me *IDXGIFactory) CreateSwapChain(
	releaser *win.OleReleaser,
	device *win.IUnknown,
	pDesc *DXGI_SWAP_CHAIN_DESC,
) (*IDXGISwapChain, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactoryVt](me.Ppvt()).CreateSoftwareAdapter,
		me.Ppvt(),
		device.Ppvt(),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGISwapChain](ret, ppvtQueried, releaser)
}

// [EnumAdapters] method.
//
// [EnumAdapters]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-enumadapters
func (me *IDXGIFactory) EnumAdapters(releaser *win.OleReleaser) ([]*IDXGIAdapter, error) {
	var index uint32
	var ppvtQueried uintptr
	var adapters []*IDXGIAdapter

	for {
		ret, _, _ := syscall.SyscallN(
			utl.Vt[_IDXGIFactoryVt](me.Ppvt()).EnumAdapters,
			me.Ppvt(),
			uintptr(index),
			uintptr(unsafe.Pointer(&ppvtQueried)))

		if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
			pObj := utl.OleNew[*IDXGIAdapter](ppvtQueried, releaser)
			adapters = append(adapters, pObj)
			index++
		} else if hr == codxgi.HRESULT_DXGI_ERROR_NOT_FOUND {
			return adapters, nil // no further adapters
		} else { // actual error
			return nil, hr
		}
	}
}

// [GetWindowAssociation] method.
//
// [GetWindowAssociation]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-getwindowassociation
func (me *IDXGIFactory) GetWindowAssociation() (win.HWND, error) {
	return utl.OleCallReturnStruct[win.HWND](me,
		utl.Vt[_IDXGIFactoryVt](me.Ppvt()).GetWindowAssociation)
}

// [MakeWindowAssociation] method.
//
// [MakeWindowAssociation]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-makewindowassociation
func (me *IDXGIFactory) MakeWindowAssociation(hWnd win.HWND, flags codxgi.DXGI_MWA) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactoryVt](me.Ppvt()).MakeWindowAssociation,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(flags))
	return utl.HresultToError(ret)
}
