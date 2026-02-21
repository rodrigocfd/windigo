//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IDXGIAdapter] COM interface.
//
// Implements [OleObj] and [OleResource].
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
	return &co.IID_IDXGIAdapter
}

// [CheckInterfaceSupport] method.
//
// [CheckInterfaceSupport]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiadapter-checkinterfacesupport
func (me *IDXGIAdapter) CheckInterfaceSupport(pInterfaceName *co.GUID) (int, error) {
	var umdVersion int64
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIAdapterVt)(unsafe.Pointer(*me.Ppvt())).CheckInterfaceSupport,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
func (me *IDXGIAdapter) EnumOutputs(releaser *OleReleaser) ([]*IDXGIOutput, error) {
	var index uint32
	var ppvtQueried **_IUnknownVt
	var adapters []*IDXGIOutput

	for {
		ret, _, _ := syscall.SyscallN(
			(*_IDXGIAdapterVt)(unsafe.Pointer(*me.Ppvt())).EnumOutputs,
			uintptr(unsafe.Pointer(me.Ppvt())),
			uintptr(index),
			uintptr(unsafe.Pointer(&ppvtQueried)))

		if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
			var pObj *IDXGIOutput
			com_buildObj(&pObj, ppvtQueried, releaser)
			adapters = append(adapters, pObj)
		} else if hr == co.HRESULT_DXGI_ERROR_NOT_FOUND {
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
	var ad DXGI_ADAPTER_DESC
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIAdapterVt)(unsafe.Pointer(*me.Ppvt())).GetDesc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ad)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ad, nil
	} else {
		return DXGI_ADAPTER_DESC{}, hr
	}
}

// [IDXGIDeviceSubObject] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IDXGIDeviceSubObject]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgidevicesubobject
type IDXGIDeviceSubObject struct{ IDXGIObject }

type _IDXGIDeviceSubObjectVt struct {
	_IDXGIObjectVt
	GetDevice uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIDeviceSubObject) IID() *co.IID {
	return &co.IID_IDXGIDeviceSubObject
}

// [GetDevice] method.
//
// [GetDevice]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevicesubobject-getdevice
func (me *IDXGIDeviceSubObject) GetDevice(releaser *OleReleaser, ppOut interface{}) error {
	piid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt

	ret, _, _ := syscall.SyscallN(
		(*_IDXGIDeviceSubObjectVt)(unsafe.Pointer(*me.Ppvt())).GetDevice,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
}

// [IDXGIFactory] COM interface.
//
// Implements [OleObj] and [OleResource].
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
	return &co.IID_IDXGIFactory
}

// [CreateSoftwareAdapter] method.
//
// [CreateSoftwareAdapter]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-createsoftwareadapter
func (me *IDXGIFactory) CreateSoftwareAdapter(releaser *OleReleaser, hModule HINSTANCE) (*IDXGIAdapter, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateSoftwareAdapter,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IDXGIAdapter](ret, ppvtQueried, releaser)
}

// [CreateSwapChain] method.
//
// [CreateSwapChain]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-createswapchain
func (me *IDXGIFactory) CreateSwapChain(
	releaser *OleReleaser,
	device OleObj,
	pDesc *DXGI_SWAP_CHAIN_DESC,
) (*IDXGISwapChain, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIFactoryVt)(unsafe.Pointer(*me.Ppvt())).CreateSoftwareAdapter,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(device.Ppvt())),
		uintptr(unsafe.Pointer(pDesc)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IDXGISwapChain](ret, ppvtQueried, releaser)
}

// [EnumAdapters] method.
//
// [EnumAdapters]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-enumadapters
func (me *IDXGIFactory) EnumAdapters(releaser *OleReleaser) ([]*IDXGIAdapter, error) {
	var index uint32
	var ppvtQueried **_IUnknownVt
	var adapters []*IDXGIAdapter

	for {
		ret, _, _ := syscall.SyscallN(
			(*_IDXGIFactoryVt)(unsafe.Pointer(*me.Ppvt())).EnumAdapters,
			uintptr(unsafe.Pointer(me.Ppvt())),
			uintptr(index),
			uintptr(unsafe.Pointer(&ppvtQueried)))

		if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
			var pObj *IDXGIAdapter
			com_buildObj(&pObj, ppvtQueried, releaser)
			adapters = append(adapters, pObj)
		} else if hr == co.HRESULT_DXGI_ERROR_NOT_FOUND {
			return adapters, nil // no further adapters
		} else { // actual error
			return nil, hr
		}
	}
}

// [GetWindowAssociation] method.
//
// [GetWindowAssociation]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-getwindowassociation
func (me *IDXGIFactory) GetWindowAssociation() (HWND, error) {
	var hWnd HWND
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIFactoryVt)(unsafe.Pointer(*me.Ppvt())).GetWindowAssociation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hWnd)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hWnd, nil
	} else {
		return HWND(0), hr
	}
}

// [MakeWindowAssociation] method.
//
// [MakeWindowAssociation]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgifactory-makewindowassociation
func (me *IDXGIFactory) MakeWindowAssociation(hWnd HWND, flags co.DXGI_MWA) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIFactoryVt)(unsafe.Pointer(*me.Ppvt())).MakeWindowAssociation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [IDXGIObject] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IDXGIObject]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgiobject
type IDXGIObject struct{ IUnknown }

type _IDXGIObjectVt struct {
	_IUnknownVt
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIObject) IID() *co.IID {
	return &co.IID_IDXGIObject
}

// [GetParent] method.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-getparent
func (me *IDXGIObject) GetParent(releaser *OleReleaser, ppOut interface{}) error {
	piid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt

	ret, _, _ := syscall.SyscallN(
		(*_IDXGIObjectVt)(unsafe.Pointer(*me.Ppvt())).GetParent,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
}

// [GetPrivateData] method.
//
// [GetPrivateData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-getprivatedata
func (me *IDXGIObject) GetPrivateData(pName *co.GUID, szData int, pData unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIObjectVt)(unsafe.Pointer(*me.Ppvt())).GetPrivateData,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pName)),
		uintptr(uint32(szData)),
		uintptr(pData))
	return utl.HresultToError(ret)
}

// [GetPrivateData] method, specialized to return an [IUnknown]-derived object.
//
// [GetPrivateData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-getprivatedata
func (me *IDXGIObject) GetPrivateDataInterface(releaser *OleReleaser, pName *co.GUID, ppOut interface{}) error {
	com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt

	if err := me.GetPrivateData(pName, int(unsafe.Sizeof(uintptr(0))), unsafe.Pointer(&ppvtQueried)); err != nil {
		return err
	}
	return com_buildObj(ppOut, ppvtQueried, releaser)
}

// [SetPrivateData] method.
//
// [SetPrivateData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-setprivatedata
func (me *IDXGIObject) SetPrivateData(pName *co.GUID, szData int, pData unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIObjectVt)(unsafe.Pointer(*me.Ppvt())).SetPrivateData,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pName)),
		uintptr(uint32(szData)),
		uintptr(pData))
	return utl.HresultToError(ret)
}

// [SetPrivateDataInterface] method.
//
// [SetPrivateDataInterface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiobject-setprivatedatainterface
func (me *IDXGIObject) SetPrivateDataInterface(pName *co.GUID, obj OleObj) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIObjectVt)(unsafe.Pointer(*me.Ppvt())).SetPrivateDataInterface,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(obj.Ppvt())))
	return utl.HresultToError(ret)
}

// [IDXGIOutput] COM interface.
//
// Implements [OleObj] and [OleResource].
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
	return &co.IID_IDXGIOutput
}

// [FindClosestMatchingMode] method.
//
// [FindClosestMatchingMode]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-findclosestmatchingmode
func (me *IDXGIOutput) FindClosestMatchingMode(
	pModeToMatch *DXGI_MODE_DESC,
	concernedDevice OleObj,
) (DXGI_MODE_DESC, error) {
	var closestMatch DXGI_MODE_DESC
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).FindClosestMatchingMode,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pModeToMatch)),
		uintptr(unsafe.Pointer(&closestMatch)),
		uintptr(com_ppvtOrNil(concernedDevice)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return closestMatch, nil
	} else {
		return DXGI_MODE_DESC{}, hr
	}
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getdesc
func (me *IDXGIOutput) GetDesc() (DXGI_OUTPUT_DESC, error) {
	var desc DXGI_OUTPUT_DESC
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetDesc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&desc)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return desc, nil
	} else {
		return DXGI_OUTPUT_DESC{}, hr
	}
}

// [GetDisplayModeList] method.
//
// [GetDisplayModeList]: https://learn.microsoft.com/en-us/windows/win32/direct3ddxgi/dxgi-enum-modes
func (me *IDXGIOutput) GetDisplayModeList(
	enumFormat co.DXGI_FORMAT,
	flags co.DXGI_ENUM_MODES,
) ([]DXGI_MODE_DESC, error) {
	var numModes uint32
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetDisplayModeList,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(enumFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(&numModes)),
		0)
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}

	desc := make([]DXGI_MODE_DESC, numModes)
	ret, _, _ = syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetDisplayModeList,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(enumFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(&numModes)),
		uintptr(unsafe.Pointer(unsafe.SliceData(desc))))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return desc, nil
	} else {
		return nil, hr
	}
}

// [GetDisplaySurfaceData] method.
//
// [GetDisplaySurfaceData]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getdisplaysurfacedata
func (me *IDXGIOutput) GetDisplaySurfaceData(dest *IDXGISurface) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetDisplaySurfaceData,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(dest.Ppvt())))
	return utl.HresultToError(ret)
}

// [GetFrameStatistics] method.
//
// [GetFrameStatistics]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_frame_statistics
func (me *IDXGIOutput) GetFrameStatistics() (DXGI_FRAME_STATISTICS, error) {
	var stats DXGI_FRAME_STATISTICS
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetFrameStatistics,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&stats)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stats, nil
	} else {
		return DXGI_FRAME_STATISTICS{}, hr
	}
}

// [GetGammaControl] method.
//
// [GetGammaControl]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getgammacontrol
func (me *IDXGIOutput) GetGammaControl() (DXGI_GAMMA_CONTROL, error) {
	var g DXGI_GAMMA_CONTROL
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetGammaControl,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&g)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return g, nil
	} else {
		return DXGI_GAMMA_CONTROL{}, hr
	}
}

// [GetGammaControlCapabilities] method.
//
// [GetGammaControlCapabilities]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-getgammacontrolcapabilities
func (me *IDXGIOutput) GetGammaControlCapabilities() (DXGI_GAMMA_CONTROL_CAPABILITIES, error) {
	var gammaCaps DXGI_GAMMA_CONTROL_CAPABILITIES
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).GetGammaControlCapabilities,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&gammaCaps)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return gammaCaps, nil
	} else {
		return DXGI_GAMMA_CONTROL_CAPABILITIES{}, hr
	}
}

// [ReleaseOwnership] method.
//
// [ReleaseOwnership]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-releaseownership
func (me *IDXGIOutput) ReleaseOwnership() {
	_, _, _ = syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).ReleaseOwnership,
		uintptr(unsafe.Pointer(me.Ppvt())))
}

// [SetDisplaySurface] method.
//
// [SetDisplaySurface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-setdisplaysurface
func (me *IDXGIOutput) SetDisplaySurface(scanoutSurface *IDXGISurface) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).SetDisplaySurface,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(scanoutSurface.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetGammaControl] method.
//
// [SetGammaControl]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-setgammacontrol
func (me *IDXGIOutput) SetGammaControl(pArray *DXGI_GAMMA_CONTROL) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).SetGammaControl,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pArray)))
	return utl.HresultToError(ret)
}

// [TakeOwnership] method.
//
// [TakeOwnership]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-takeownership
func (me *IDXGIOutput) TakeOwnership(device OleObj, exclusive bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).TakeOwnership,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(device.Ppvt())),
		utl.BoolToUintptr(exclusive))
	return utl.HresultToError(ret)
}

// [WaitForVBlank] method.
//
// [WaitForVBlank]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgioutput-waitforvblank
func (me *IDXGIOutput) WaitForVBlank() error {
	return com_callErr(me,
		(*_IDXGIOutputVt)(unsafe.Pointer(*me.Ppvt())).WaitForVBlank)
}

// [IDXGISurface] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IDXGISurface]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgisurface
type IDXGISurface struct{ IDXGIDeviceSubObject }

type _IDXGISurfaceVt struct {
	_IDXGIDeviceSubObjectVt
	GetDesc uintptr
	Map     uintptr
	Unmap   uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGISurface) IID() *co.IID {
	return &co.IID_IDXGISurface
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface-getdesc
func (me *IDXGISurface) GetDesc() (DXGI_SURFACE_DESC, error) {
	var desc DXGI_SURFACE_DESC
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISurfaceVt)(unsafe.Pointer(*me.Ppvt())).GetDesc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&desc)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return desc, nil
	} else {
		return DXGI_SURFACE_DESC{}, hr
	}
}

// [Map] method.
//
// [Map]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface-map
func (me *IDXGISurface) Map(flags co.DXGI_MAP) (DXGI_MAPPED_RECT, error) {
	var lockedRect DXGI_MAPPED_RECT
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISurfaceVt)(unsafe.Pointer(*me.Ppvt())).Map,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&lockedRect)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return lockedRect, nil
	} else {
		return DXGI_MAPPED_RECT{}, hr
	}
}

// [Unmap] method.
//
// [Unmap]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface-unmap
func (me *IDXGISurface) Unmap() error {
	return com_callErr(me,
		(*_IDXGISurfaceVt)(unsafe.Pointer(*me.Ppvt())).Unmap)
}

// [IDXGISwapChain] COM interface.
//
// Implements [OleObj] and [OleResource].
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
	return &co.IID_IDXGISwapChain
}

// [GetBuffer] method.
//
// [GetBuffer]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getbuffer
func (me *IDXGISwapChain) GetBuffer(
	releaser *OleReleaser,
	bufferIndex int,
	ppOut interface{},
) error {
	piid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt

	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).GetBuffer,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(bufferIndex)),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
}

// [GetContainingOutput] method.
//
// [GetContainingOutput]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getcontainingoutput
func (me *IDXGISwapChain) GetContainingOutput(releaser *OleReleaser) (*IDXGIOutput, error) {
	return com_callObj[*IDXGIOutput](me, releaser,
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).GetContainingOutput)
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getdesc
func (me *IDXGISwapChain) GetDesc() (DXGI_SWAP_CHAIN_DESC, error) {
	var desc DXGI_SWAP_CHAIN_DESC
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).GetDesc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&desc)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return desc, nil
	} else {
		return DXGI_SWAP_CHAIN_DESC{}, hr
	}
}

// [GetFrameStatistics] method.
//
// [GetFrameStatistics]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getframestatistics
func (me *IDXGISwapChain) GetFrameStatistics() (DXGI_FRAME_STATISTICS, error) {
	var stats DXGI_FRAME_STATISTICS
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).GetFrameStatistics,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&stats)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stats, nil
	} else {
		return DXGI_FRAME_STATISTICS{}, hr
	}
}

// [GetFullscreenState] method.
//
// [GetFullscreenState]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-getfullscreenstate
func (me *IDXGISwapChain) GetFullscreenState(
	releaser *OleReleaser,
) (isFullScreen bool, output *IDXGIOutput, hr error) {
	var bFullScreen BOOL
	var ppvtQueried **_IUnknownVt

	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).GetFullscreenState,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bFullScreen)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		if bFullScreen.Ok() {
			var pObj *IDXGIOutput
			com_buildObj(&pObj, ppvtQueried, releaser)
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
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).GetLastPresentCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
func (me *IDXGISwapChain) Present(syncInterval int, flags co.DXGI_PRESENT) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).Present,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(syncInterval)),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [ResizeBuffers] method.
//
// [ResizeBuffers]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-resizebuffers
func (me *IDXGISwapChain) ResizeBuffers(
	bufferCount int,
	szBackBuffer SIZE,
	newFormat co.DXGI_FORMAT,
	flags co.DXGI_SWAP_CHAIN_FLAG,
) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).ResizeBuffers,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).ResizeTarget,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pNewTargetParams)))
	return utl.HresultToError(ret)
}

// [SetFullscreenState] method.
//
// [SetFullscreenState]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiswapchain-setfullscreenstate
func (me *IDXGISwapChain) SetFullscreenState(fullScreen bool, target *IDXGIOutput) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDXGISwapChainVt)(unsafe.Pointer(*me.Ppvt())).SetFullscreenState,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(fullScreen),
		uintptr(com_ppvtOrNil(target)))
	return utl.HresultToError(ret)
}
