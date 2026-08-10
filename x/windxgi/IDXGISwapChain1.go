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

// [IDXGISwapChain1] COM interface.
//
// [IDXGISwapChain1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgiswapchain1
type IDXGISwapChain1 struct{ IDXGISwapChain }

type _IDXGISwapChain1Vt struct {
	_IDXGISwapChainVt
	GetDesc1                 uintptr
	GetFullscreenDesc        uintptr
	GetHwnd                  uintptr
	GetCoreWindow            uintptr
	Present1                 uintptr
	IsTemporaryMonoSupported uintptr
	GetRestrictToOutput      uintptr
	SetBackgroundColor       uintptr
	GetBackgroundColor       uintptr
	SetRotation              uintptr
	GetRotation              uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGISwapChain1) IID() *co.IID {
	return &codxgi.IID_IDXGISwapChain1
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGISwapChain1) AddRef(releaser *win.OleReleaser) *IDXGISwapChain1 {
	return utl.OleNewFromAddRef[*IDXGISwapChain1](me, releaser)
}

// [GetBackgroundColor] method.
//
// [GetBackgroundColor]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-getbackgroundcolor
func (me *IDXGISwapChain1) GetBackgroundColor() (DXGI_RGBA, error) {
	return utl.OleCallReturnStruct[DXGI_RGBA](me,
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetBackgroundColor)
}

// [GetCoreWindow] method.
//
// [GetCoreWindow]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-getcorewindow
func (me *IDXGISwapChain1) GetCoreWindow(
	releaser *win.OleReleaser,
	piid *co.IID,
	ppOut interface{},
) error {
	_ = utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetCoreWindow,
		me.Ppvt(),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [GetDesc1] method.
//
// [GetDesc1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-getdesc1
func (me *IDXGISwapChain1) GetDesc1() (DXGI_SWAP_CHAIN_DESC1, error) {
	return utl.OleCallReturnStruct[DXGI_SWAP_CHAIN_DESC1](me,
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetDesc1)
}

// [GetFullscreenDesc] method.
//
// [GetFullscreenDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-getfullscreendesc
func (me *IDXGISwapChain1) GetFullscreenDesc() (DXGI_SWAP_CHAIN_FULLSCREEN_DESC, error) {
	return utl.OleCallReturnStruct[DXGI_SWAP_CHAIN_FULLSCREEN_DESC](me,
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetFullscreenDesc)
}

// [GetHwnd] method.
//
// [GetHwnd]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-gethwnd
func (me *IDXGISwapChain1) GetHwnd() (win.HWND, error) {
	return utl.OleCallReturnStruct[win.HWND](me,
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetHwnd)
}

// [GetRestrictToOutput] method.
//
// [GetRestrictToOutput]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-getrestricttooutput
func (me *IDXGISwapChain1) GetRestrictToOutput(releaser *win.OleReleaser) (*IDXGIOutput, error) {
	return utl.OleNewFromCallWithoutParms[*IDXGIOutput](me, releaser,
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetRestrictToOutput)
}

// [GetRotation] method.
//
// [GetRotation]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-getrotation
func (me *IDXGISwapChain1) GetRotation() (codxgi.DXGI_MODE_ROTATION, error) {
	return utl.OleCallReturnStruct[codxgi.DXGI_MODE_ROTATION](me,
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).GetRotation)
}

// [IsTemporaryMonoSupported] method.
//
// [IsTemporaryMonoSupported]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-istemporarymonosupported
func (me *IDXGISwapChain1) IsTemporaryMonoSupported() bool {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).IsTemporaryMonoSupported,
		me.Ppvt())
	return ret != 0
}

// [Present1] method.
//
// [Present1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-present1
func (me *IDXGISwapChain1) Present1(
	syncInterval int,
	presentFlags codxgi.DXGI_PRESENT,
	pPresentParameters *DXGI_PRESENT_PARAMETERS,
) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).Present1,
		me.Ppvt(),
		uintptr(uint32(syncInterval)),
		uintptr(presentFlags),
		uintptr(unsafe.Pointer(pPresentParameters)))
	return utl.HresultToError(ret)
}

// [SetBackgroundColor] method.
//
// [SetBackgroundColor]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-setbackgroundcolor
func (me *IDXGISwapChain1) SetBackgroundColor(pColor *DXGI_RGBA) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).SetBackgroundColor,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pColor)))
	return utl.HresultToError(ret)
}

// [SetRotation] method.
//
// [SetRotation]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgiswapchain1-setrotation
func (me *IDXGISwapChain1) SetRotation(rotation codxgi.DXGI_MODE_ROTATION) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISwapChain1Vt](me.Ppvt()).SetRotation,
		me.Ppvt(),
		uintptr(rotation))
	return utl.HresultToError(ret)
}
