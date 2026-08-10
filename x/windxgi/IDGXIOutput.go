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

// [IDXGIOutput] COM interface.
//
// [IDXGIOutput]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgioutput
type IDXGIOutput struct{ IDXGIObject }

type _IDXGIOutputVt struct {
	_IDXGIObjectVt
	GetDesc                     uintptr
	GetDisplayModeList          uintptr
	FindClosestMatchingMode     uintptr
	WaitForVBlank               uintptr
	TakeOwnership               uintptr
	ReleaseOwnership            uintptr
	GetGammaControlCapabilities uintptr
	SetGammaControl             uintptr
	GetGammaControl             uintptr
	SetDisplaySurface           uintptr
	GetDisplaySurfaceData       uintptr
	GetFrameStatistics          uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIOutput) IID() *co.IID {
	return &codxgi.IID_IDXGIOutput
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIOutput) AddRef(releaser *win.OleReleaser) *IDXGIOutput {
	return utl.OleNewFromAddRef[*IDXGIOutput](me, releaser)
}

// [FindClosestMatchingMode] method.
//
// [FindClosestMatchingMode]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-findclosestmatchingmode
func (me *IDXGIOutput) FindClosestMatchingMode(
	pModeToMatch *DXGI_MODE_DESC,
	concernedDevice *win.IUnknown,
) (DXGI_MODE_DESC, error) {
	var closestMatch DXGI_MODE_DESC
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).FindClosestMatchingMode,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pModeToMatch)),
		uintptr(unsafe.Pointer(&closestMatch)),
		utl.OlePpvtOrNil(concernedDevice))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return DXGI_MODE_DESC{}, hr
	}
	return closestMatch, nil
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getdesc
func (me *IDXGIOutput) GetDesc() (DXGI_OUTPUT_DESC, error) {
	return utl.OleCallReturnStruct[DXGI_OUTPUT_DESC](me,
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetDesc)
}

// [GetDisplayModeList] method.
//
// [GetDisplayModeList]: https://learn.microsoft.com/en-us/windows/win32/direct3ddxgi/dxgi-enum-modes
func (me *IDXGIOutput) GetDisplayModeList(
	enumFormat codxgi.DXGI_FORMAT,
	flags codxgi.DXGI_ENUM_MODES,
) ([]DXGI_MODE_DESC, error) {
	var numModes uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetDisplayModeList,
		me.Ppvt(),
		uintptr(enumFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(&numModes)),
		0)
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}

	desc := make([]DXGI_MODE_DESC, numModes)
	ret, _, _ = syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetDisplayModeList,
		me.Ppvt(),
		uintptr(enumFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(&numModes)),
		uintptr(unsafe.Pointer(&desc[0])))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}
	return desc, nil
}

// [GetDisplaySurfaceData] method.
//
// [GetDisplaySurfaceData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getdisplaysurfacedata
func (me *IDXGIOutput) GetDisplaySurfaceData(dest *IDXGISurface) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetDisplaySurfaceData,
		me.Ppvt(),
		dest.Ppvt())
	return utl.HresultToError(ret)
}

// [GetFrameStatistics] method.
//
// [GetFrameStatistics]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_frame_statistics
func (me *IDXGIOutput) GetFrameStatistics() (DXGI_FRAME_STATISTICS, error) {
	return utl.OleCallReturnStruct[DXGI_FRAME_STATISTICS](me,
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetFrameStatistics)
}

// [GetGammaControl] method.
//
// [GetGammaControl]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getgammacontrol
func (me *IDXGIOutput) GetGammaControl() (DXGI_GAMMA_CONTROL, error) {
	return utl.OleCallReturnStruct[DXGI_GAMMA_CONTROL](me,
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetGammaControl)
}

// [GetGammaControlCapabilities] method.
//
// [GetGammaControlCapabilities]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getgammacontrolcapabilities
func (me *IDXGIOutput) GetGammaControlCapabilities() (DXGI_GAMMA_CONTROL_CAPABILITIES, error) {
	return utl.OleCallReturnStruct[DXGI_GAMMA_CONTROL_CAPABILITIES](me,
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).GetGammaControlCapabilities)
}

// [ReleaseOwnership] method.
//
// [ReleaseOwnership]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-releaseownership
func (me *IDXGIOutput) ReleaseOwnership() {
	utl.OleCallWithoutParms(me, utl.Vt[_IDXGIOutputVt](me.Ppvt()).ReleaseOwnership)
}

// [SetDisplaySurface] method.
//
// [SetDisplaySurface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-setdisplaysurface
func (me *IDXGIOutput) SetDisplaySurface(scanoutSurface *IDXGISurface) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).SetDisplaySurface,
		me.Ppvt(),
		scanoutSurface.Ppvt())
	return utl.HresultToError(ret)
}

// [SetGammaControl] method.
//
// [SetGammaControl]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-setgammacontrol
func (me *IDXGIOutput) SetGammaControl(pArray *DXGI_GAMMA_CONTROL) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).SetGammaControl,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pArray)))
	return utl.HresultToError(ret)
}

// [TakeOwnership] method.
//
// [TakeOwnership]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-takeownership
func (me *IDXGIOutput) TakeOwnership(device *win.IUnknown, exclusive bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputVt](me.Ppvt()).TakeOwnership,
		me.Ppvt(),
		device.Ppvt(),
		utl.BoolToUintptr(exclusive))
	return utl.HresultToError(ret)
}

// [WaitForVBlank] method.
//
// [WaitForVBlank]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-waitforvblank
func (me *IDXGIOutput) WaitForVBlank() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IDXGIOutputVt](me.Ppvt()).WaitForVBlank)
}
