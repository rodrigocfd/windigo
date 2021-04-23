package dshow

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _IBaseFilterVtbl struct {
	_IMediaFilterVtbl
	EnumPins        uintptr
	FindPin         uintptr
	QueryFilterInfo uintptr
	JoinFilterGraph uintptr
	QueryVendorInfo uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
type IBaseFilter struct {
	IMediaFilter // Base IMediaFilter > IPersist > IUnknown.
}

// Calls IUnknown.CoCreateInstance() to return IBaseFilter.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/medfound/using-the-directshow-evr-filter
func CoCreateEnhancedVideoRenderer(dwClsContext co.CLSCTX) (IBaseFilter, error) {
	iUnk, lerr := win.CoCreateInstance(
		CLSID.EnhancedVideoRenderer, nil, dwClsContext, IID.IBaseFilter)
	if lerr != nil {
		return IBaseFilter{}, lerr
	}
	return IBaseFilter{
		IMediaFilter{
			IPersist{IUnknown: iUnk},
		},
	}, nil
}

// Calls IUnknown.CoCreateInstance() to return IBaseFilter.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func CoCreateVideoMixingRenderer9(dwClsContext co.CLSCTX) (IBaseFilter, error) {
	iUnk, lerr := win.CoCreateInstance(
		CLSID.VideoMixingRenderer9, nil, dwClsContext, IID.IBaseFilter)
	if lerr != nil {
		return IBaseFilter{}, lerr
	}
	return IBaseFilter{
		IMediaFilter{
			IPersist{IUnknown: iUnk},
		},
	}, nil
}

// Calls IUnknown.QueryInterface() to return IMFGetService.
//
// ⚠️ You must defer Release().
func (me *IBaseFilter) QueryIMFGetService() (IMFGetService, error) {
	iUnk, lerr := me.QueryInterface(IID.IMFGetService)
	if lerr != nil {
		return IMFGetService{}, lerr
	}
	return IMFGetService{IUnknown: iUnk}, nil
}
