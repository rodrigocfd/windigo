package dshow

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
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

// Calls CoCreateInstance(), typically with CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/medfound/using-the-directshow-evr-filter
func NewEnhancedVideoRenderer(dwClsContext co.CLSCTX) IBaseFilter {
	iUnk := win.CoCreateInstance(
		dshowco.CLSID_EnhancedVideoRenderer, nil, dwClsContext,
		dshowco.IID_IBaseFilter)
	return IBaseFilter{
		IMediaFilter{
			IPersist{IUnknown: iUnk},
		},
	}
}

// Calls CoCreateInstance(), typically with CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func NewVideoMixingRenderer9(dwClsContext co.CLSCTX) IBaseFilter {
	iUnk := win.CoCreateInstance(
		dshowco.CLSID_VideoMixingRenderer9, nil, dwClsContext,
		dshowco.IID_IBaseFilter)
	return IBaseFilter{
		IMediaFilter{
			IPersist{IUnknown: iUnk},
		},
	}
}

// Calls IUnknown.QueryInterface() to return IMFGetService.
//
// ⚠️ You must defer Release().
func (me *IBaseFilter) QueryIMFGetService() IMFGetService {
	iUnk := me.QueryInterface(dshowco.IID_IMFGetService)
	return IMFGetService{IUnknown: iUnk}
}
