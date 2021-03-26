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
	clsidEnhancedVideoRenderer := win.NewGuid(0xfa10746c, 0x9b63, 0x4b6c, 0xbc49, 0xfc300ea5f256)
	iidIBaseFilter := win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770)

	iUnk, lerr := win.CoCreateInstance(
		clsidEnhancedVideoRenderer, nil, dwClsContext, iidIBaseFilter)
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
	clsidVideoMixingRenderer9 := win.NewGuid(0x51b4abf3, 0x748f, 0x4e3b, 0xa276, 0xc828330e926a)
	iidIBaseFilter := win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770)

	iUnk, lerr := win.CoCreateInstance(
		clsidVideoMixingRenderer9, nil, dwClsContext, iidIBaseFilter)
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
	iidIMFGetService := win.NewGuid(0xfa993888, 0x4383, 0x415a, 0xa930, 0xdd472a8cf6f7)

	iUnk, lerr := me.QueryInterface(iidIMFGetService)
	if lerr != nil {
		return IMFGetService{}, lerr
	}
	return IMFGetService{IUnknown: iUnk}, nil
}
