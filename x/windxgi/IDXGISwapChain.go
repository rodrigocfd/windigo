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

// [IDXGISwapChain] COM interface.
//
// [IDXGISwapChain]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgiswapchain
type IDXGISwapChain struct{ IDXGIDeviceSubObject }

type _IDXGISwapChainVt struct {
	_IDXGIDeviceSubObjectVt
	Present             uintptr
	GetBuffer           uintptr
	SetFullscreenState  uintptr
	GetFullscreenState  uintptr
	GetDesc             uintptr
	ResizeBuffers       uintptr
	ResizeTarget        uintptr
	GetContainingOutput uintptr
	GetFrameStatistics  uintptr
	GetLastPresentCount uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGISwapChain) IID() *co.IID {
	return &codxgi.IID_IDXGISwapChain
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGISwapChain) AddRef(releaser *win.OleReleaser) *IDXGISwapChain {
	return utl.OleNewFromAddRef[*IDXGISwapChain](me, releaser)
}

// [GetBuffer] method.
//
// [GetBuffer]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getbuffer
func (me *IDXGISwapChain) GetBuffer(
	releaser *win.OleReleaser,
	bufferIndex int,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).GetBuffer,
		me.Ppvt(),
		uintptr(uint32(bufferIndex)),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [GetContainingOutput] method.
//
// [GetContainingOutput]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getcontainingoutput
func (me *IDXGISwapChain) GetContainingOutput(releaser *win.OleReleaser) (*IDXGIOutput, error) {
	return utl.OleNewFromCallWithoutParms[*IDXGIOutput](me, releaser,
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).GetContainingOutput)
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getdesc
func (me *IDXGISwapChain) GetDesc() (DXGI_SWAP_CHAIN_DESC, error) {
	return utl.OleCallReturnStruct[DXGI_SWAP_CHAIN_DESC](me,
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).GetDesc)
}

// [GetFrameStatistics] method.
//
// [GetFrameStatistics]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getframestatistics
func (me *IDXGISwapChain) GetFrameStatistics() (DXGI_FRAME_STATISTICS, error) {
	return utl.OleCallReturnStruct[DXGI_FRAME_STATISTICS](me,
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).GetFrameStatistics)
}

// [GetFullscreenState] method.
//
// [GetFullscreenState]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getfullscreenstate
func (me *IDXGISwapChain) GetFullscreenState(
	releaser *win.OleReleaser,
) (isFullScreen bool, output *IDXGIOutput, hr error) {
	var bFullScreen win.BOOL
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).GetFullscreenState,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&bFullScreen)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		if bFullScreen.Ok() {
			pObj := utl.OleNew[*IDXGIOutput](ppvtQueried, releaser)
			return true, pObj, nil
		} else {
			return false, nil, nil
		}
	} else {
		return false, nil, hr
	}
}

// [GetLastPresentCount] method.
//
// [GetLastPresentCount]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getlastpresentcount
func (me *IDXGISwapChain) GetLastPresentCount() (int, error) {
	var c uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).GetLastPresentCount,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&c)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(c), nil
	} else {
		return 0, hr
	}
}

// [Present] method.
//
// [Present]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-present
func (me *IDXGISwapChain) Present(syncInterval int, flags codxgi.DXGI_PRESENT) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).Present,
		me.Ppvt(),
		uintptr(uint32(syncInterval)),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [ResizeBuffers] method.
//
// [ResizeBuffers]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-resizebuffers
func (me *IDXGISwapChain) ResizeBuffers(
	bufferCount int,
	szBackBuffer win.SIZE,
	newFormat codxgi.DXGI_FORMAT,
	flags codxgi.DXGI_SWAP_CHAIN_FLAG,
) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).ResizeBuffers,
		me.Ppvt(),
		uintptr(uint32(bufferCount)),
		uintptr(uint32(szBackBuffer.Cx)),
		uintptr(uint32(szBackBuffer.Cy)),
		uintptr(newFormat),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [ResizeTarget] method.
//
// [ResizeTarget]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-resizetarget
func (me *IDXGISwapChain) ResizeTarget(pNewTargetParams *DXGI_MODE_DESC) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).ResizeTarget,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pNewTargetParams)))
	return utl.HresultToError(ret)
}

// [SetFullscreenState] method.
//
// [SetFullscreenState]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-setfullscreenstate
func (me *IDXGISwapChain) SetFullscreenState(fullScreen bool, target *IDXGIOutput) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChainVt](me.Ppvt()).SetFullscreenState,
		me.Ppvt(),
		utl.BoolToUintptr(fullScreen),
		utl.OlePpvtOrNil(target))
	return utl.HresultToError(ret)
}
