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

// [IDXGIFactory2] COM interface.
//
// [IDXGIFactory2]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgifactory2
type IDXGIFactory2 struct{ IDXGIFactory1 }

type _IDXGIFactory2Vt struct {
	_IDXGIFactory1Vt
	IsWindowedStereoEnabled       uintptr
	CreateSwapChainForHwnd        uintptr
	CreateSwapChainForCoreWindow  uintptr
	GetSharedResourceAdapterLuid  uintptr
	RegisterStereoStatusWindow    uintptr
	RegisterStereoStatusEvent     uintptr
	UnregisterStereoStatus        uintptr
	RegisterOcclusionStatusWindow uintptr
	RegisterOcclusionStatusEvent  uintptr
	UnregisterOcclusionStatus     uintptr
	CreateSwapChainForComposition uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIFactory2) IID() *co.IID {
	return &codxgi.IID_IDXGIFactory2
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIFactory2) AddRef(releaser *win.OleReleaser) *IDXGIFactory2 {
	return utl.OleNewFromAddRef[*IDXGIFactory2](me, releaser)
}

// [CreateSwapChainForComposition] method.
//
// [CreateSwapChainForComposition]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-createswapchainforcomposition
func (me *IDXGIFactory2) CreateSwapChainForComposition(
	releaser *win.OleReleaser,
	device *win.IUnknown,
	pDesc *DXGI_SWAP_CHAIN_DESC1,
	restrictToOutput *IDXGIOutput,
) (*IDXGISwapChain1, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).CreateSwapChainForComposition,
		me.Ppvt(),
		utl.OlePpvtOrNil(device),
		uintptr(unsafe.Pointer(pDesc)),
		utl.OlePpvtOrNil(restrictToOutput),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGISwapChain1](ret, ppvtQueried, releaser)
}

// [CreateSwapChainForHwnd] method.
//
// [CreateSwapChainForHwnd]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-createswapchainforhwnd
func (me *IDXGIFactory2) CreateSwapChainForHwnd(
	releaser *win.OleReleaser,
	device *win.IUnknown,
	hWnd win.HWND,
	pDesc *DXGI_SWAP_CHAIN_DESC1,
	pFullscreenDesc *DXGI_SWAP_CHAIN_FULLSCREEN_DESC,
	restrictToOutput *IDXGIOutput,
) (*IDXGISwapChain1, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).CreateSwapChainForHwnd,
		me.Ppvt(),
		utl.OlePpvtOrNil(device),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(unsafe.Pointer(pFullscreenDesc)),
		utl.OlePpvtOrNil(restrictToOutput),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGISwapChain1](ret, ppvtQueried, releaser)
}

// [GetSharedResourceAdapterLuid] method.
//
// [GetSharedResourceAdapterLuid]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-getsharedresourceadapterluid
func (me *IDXGIFactory2) GetSharedResourceAdapterLuid() (co.LUID, error) {
	return utl.OleCallReturnStruct[co.LUID](me,
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).GetSharedResourceAdapterLuid)
}

// [IsWindowedStereoEnabled] method.
//
// [IsWindowedStereoEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-iswindowedstereoenabled
func (me *IDXGIFactory2) IsWindowedStereoEnabled() bool {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).IsWindowedStereoEnabled,
		me.Ppvt())
	return ret != 0
}

// [RegisterOcclusionStatusEvent] method.
//
// Paired with [IDXGIFactory2.UnregisterStereoStatus].
//
// [RegisterOcclusionStatusEvent]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-registerocclusionstatusevent
func (me *IDXGIFactory2) RegisterOcclusionStatusEvent(hEvent win.HANDLE) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).RegisterOcclusionStatusEvent,
		me.Ppvt(),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(&cookie)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return cookie, nil
}

// [RegisterOcclusionStatusWindow] method.
//
// Paired with [IDXGIFactory2.UnregisterOcclusionStatus].
//
// [RegisterOcclusionStatusWindow]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-registerocclusionstatuswindow
func (me *IDXGIFactory2) RegisterOcclusionStatusWindow(hWnd win.HWND, msg co.WM) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).RegisterOcclusionStatusWindow,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(unsafe.Pointer(&cookie)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return cookie, nil
}

// [RegisterStereoStatusEvent] method.
//
// Paired with [IDXGIFactory2.UnregisterStereoStatus].
//
// [RegisterStereoStatusEvent]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-registerstereostatusevent
func (me *IDXGIFactory2) RegisterStereoStatusEvent(hEvent win.HANDLE) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).RegisterStereoStatusEvent,
		me.Ppvt(),
		uintptr(hEvent),
		uintptr(unsafe.Pointer(&cookie)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return cookie, nil
}

// [RegisterStereoStatusWindow] method.
//
// Paired with [IDXGIFactory2.UnregisterStereoStatus].
//
// [RegisterStereoStatusWindow]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-registerstereostatuswindow
func (me *IDXGIFactory2) RegisterStereoStatusWindow(hWnd win.HWND, msg co.WM) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).RegisterStereoStatusWindow,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(unsafe.Pointer(&cookie)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return cookie, nil
}

// [UnregisterOcclusionStatus] method.
//
// Paired with [IDXGIFactory2.RegisterOcclusionStatusEvent] or
// [IDXGIFactory2.RegisterStereoStatusWindow].
//
// [UnregisterOcclusionStatus]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-unregisterocclusionstatus
func (me *IDXGIFactory2) UnregisterOcclusionStatus(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).UnregisterOcclusionStatus,
		me.Ppvt(),
		uintptr(cookie))
	return utl.HresultToError(ret)
}

// [UnregisterStereoStatus] method.
//
// Paired with [IDXGIFactory2.RegisterStereoStatusEvent] or
// [IDXGIFactory2.RegisterStereoStatusWindow].
//
// [UnregisterStereoStatus]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgifactory2-unregisterstereostatus
func (me *IDXGIFactory2) UnregisterStereoStatus(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIFactory2Vt](me.Ppvt()).UnregisterStereoStatus,
		me.Ppvt(),
		uintptr(cookie))
	return utl.HresultToError(ret)
}
