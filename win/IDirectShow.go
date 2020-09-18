/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
	//
	// IBaseFilter > IMediaFilter > IPersist > IUnknown.
	IBaseFilter struct{ _IBaseFilterImpl }

	_IBaseFilterImpl struct{ _IMediaFilterImpl }

	_IBaseFilterVtbl struct {
		_IMediaFilterVtbl
		EnumPins        uintptr
		FindPin         uintptr
		QueryFilterInfo uintptr
		JoinFilterGraph uintptr
		QueryVendorInfo uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
//
// https://docs.microsoft.com/en-us/windows/win32/medfound/using-the-directshow-evr-filter
func (me *IBaseFilter) CoCreateEnhancedVideoRenderer(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0xfa10746c, 0x9b63, 0x4b6c, 0xbc49_fc300ea5f256), // CLSID_EnhancedVideoRenderer
		dwClsContext,
		_Win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IBaseFilter
}

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
//
// https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func (me *IBaseFilter) CoCreateVideoMixingRenderer9(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x51b4abf3, 0x748f, 0x4e3b, 0xa276_c828330e926a), // CLSID_VideoMixingRenderer9
		dwClsContext,
		_Win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IBaseFilter
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
//
// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
func (me *_IBaseFilterImpl) QueryIMFGetService() IMFGetService {
	return IMFGetService{
		_IMFGetServiceImpl{
			_IUnknownImpl{
				ptr: me.QueryInterface(
					_Win.NewGuid(0xfa993888, 0x4383, 0x415a, 0xa930_dd472a8cf6f7)), // IID_IMFGetService
			},
		},
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
	//
	// IDispatch > IUnknown.
	IDispatch struct{ _IDispatchImpl }

	_IDispatchImpl struct{ _IUnknownImpl }

	_IDispatchVtbl struct {
		_IUnknownVtbl
		GetTypeInfoCount uintptr
		GetTypeInfo      uintptr
		GetIDsOfNames    uintptr
		Invoke           uintptr
	}
)

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
	//
	// IFilterGraph > IUnknown.
	IFilterGraph struct{ _IFilterGraphImpl }

	_IFilterGraphImpl struct{ _IUnknownImpl }

	_IFilterGraphVtbl struct {
		_IUnknownVtbl
		AddFilter            uintptr
		RemoveFilter         uintptr
		EnumFilters          uintptr
		FindFilterByName     uintptr
		ConnectDirect        uintptr
		Reconnect            uintptr
		Disconnect           uintptr
		SetDefaultSyncSource uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-addfilter
func (me *_IFilterGraphImpl) AddFilter(pFilter IBaseFilter, name string) {
	vTbl := (*_IFilterGraphVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.AddFilter, 3, uintptr(me.ptr),
		uintptr(pFilter.ptr),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(name))))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IFilterGraph.AddFilter").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
	//
	// IGraphBuilder > IFilterGraph > IUnknown.
	IGraphBuilder struct{ _IGraphBuilderImpl }

	_IGraphBuilderImpl struct{ _IFilterGraphImpl }

	_IGraphBuilderVtbl struct {
		_IFilterGraphVtbl
		Connect                 uintptr
		Render                  uintptr
		RenderFile              uintptr
		AddSourceFilter         uintptr
		SetLogFile              uintptr
		Abort                   uintptr
		ShouldOperationContinue uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *IGraphBuilder) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0xe436ebb3, 0x524f, 0x11ce, 0x9f53_0020af0ba770), // CLSID_FilterGraph
		dwClsContext,
		_Win.NewGuid(0x56a868a9, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IGraphBuilder
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *_IGraphBuilderImpl) Abort() {
	vTbl := (*_IGraphBuilderVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.Abort, 1, uintptr(me.ptr), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IGraphBuilder.Abort").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
	//
	// IMediaFilter > IPersist > IUnknown.
	IMediaFilter struct{ _IMediaFilterImpl }

	_IMediaFilterImpl struct{ _IPersistImpl }

	_IMediaFilterVtbl struct {
		_IPersistVtbl
		Stop          uintptr
		Pause         uintptr
		Run           uintptr
		GetState      uintptr
		SetSyncSource uintptr
		GetSyncSource uintptr
	}
)

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
	//
	// IMFGetService > IUnknown.
	IMFGetService struct{ _IMFGetServiceImpl }

	_IMFGetServiceImpl struct{ _IUnknownImpl }

	_IMFGetServiceVtbl struct {
		_IUnknownVtbl
		GetService uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
func (me *_IMFGetServiceImpl) GetService(
	guidService, riid *GUID) unsafe.Pointer {

	vTbl := (*_IMFGetServiceVtbl)(me.pVtbl())
	var ppvObject unsafe.Pointer = nil
	ret, _, _ := syscall.Syscall6(vTbl.GetService, 4, uintptr(me.ptr),
		uintptr(unsafe.Pointer(guidService)), uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvObject)), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IMFGetService.GetService").Error())
	}
	return ppvObject
}

// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
//
// https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
func (me *_IMFGetServiceImpl) GetIMFVideoDisplayControl() IMFVideoDisplayControl {
	return IMFVideoDisplayControl{
		_IMFVideoDisplayControlImpl{
			_IUnknownImpl{
				ptr: me.GetService(
					_Win.NewGuid(0x1092a86c, 0xab1a, 0x459a, 0xa336_831fbc4d11ff),  // MR_VIDEO_RENDER_SERVICE
					_Win.NewGuid(0xa490b1e4, 0xab84, 0x4d31, 0xa1b2_181e03b1077a)), // IID_IMFVideoDisplayControl
			},
		},
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
	//
	// IMFVideoDisplayControl > IUnknown.
	IMFVideoDisplayControl struct{ _IMFVideoDisplayControlImpl }

	_IMFVideoDisplayControlImpl struct{ _IUnknownImpl }

	_IMFVideoDisplayControlVtbl struct {
		_IUnknownVtbl
		GetNativeVideoSize uintptr
		GetIdealVideoSize  uintptr
		SetVideoPosition   uintptr
		GetVideoPosition   uintptr
		SetAspectRatioMode uintptr
		GetAspectRatioMode uintptr
		SetVideoWindow     uintptr
		GetVideoWindow     uintptr
		RepaintVideo       uintptr
		GetCurrentImage    uintptr
		SetBorderColor     uintptr
		GetBorderColor     uintptr
		SetRenderingPrefs  uintptr
		GetRenderingPrefs  uintptr
		SetFullscreen      uintptr
		GetFullscreen      uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setaspectratiomode
func (me *_IMFVideoDisplayControlImpl) SetAspectRatioMode(mode co.MFVideoARMode) {
	vTbl := (*_IMFVideoDisplayControlVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetAspectRatioMode, 2, uintptr(me.ptr),
		uintptr(mode), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IMFVideoDisplayControl.SetAspectRatioMode").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/evr/nf-evr-imfvideodisplaycontrol-setvideowindow
func (me *_IMFVideoDisplayControlImpl) SetVideoWindow(hwndVideo HWND) {
	vTbl := (*_IMFVideoDisplayControlVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetVideoWindow, 2, uintptr(me.ptr),
		uintptr(hwndVideo), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IMFVideoDisplayControl.SetVideoWindow").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
	//
	// IPersist > IUnknown.
	IPersist struct{ _IPersistImpl }

	_IPersistImpl struct{ _IUnknownImpl }

	_IPersistVtbl struct {
		_IUnknownVtbl
		GetClassID uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
func (me *_IPersistImpl) GetClassID() *GUID {
	vTbl := (*_IPersistVtbl)(me.pVtbl())
	clsid := GUID{}
	ret, _, _ := syscall.Syscall(vTbl.GetClassID, 2, uintptr(me.ptr),
		uintptr(unsafe.Pointer(&clsid)), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IPersist.GetClassID").Error())
	}
	return &clsid
}
