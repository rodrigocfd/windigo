/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type (
	// IBaseFilter > IMediaFilter > IPersist > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
	IBaseFilter struct{ IMediaFilter }

	IBaseFilterVtbl struct {
		IMediaFilterVtbl
		EnumPins        uintptr
		FindPin         uintptr
		QueryFilterInfo uintptr
		JoinFilterGraph uintptr
		QueryVendorInfo uintptr
	}
)

// Typically uses CLSCTX_INPROC_SERVER.
//
// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/medfound/using-the-directshow-evr-filter
func CoCreateEnhancedVideoRenderer(dwClsContext co.CLSCTX) *IBaseFilter {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0xfa10746c, 0x9b63, 0x4b6c, 0xbc49, 0xfc300ea5f256), // CLSID_EnhancedVideoRenderer
		nil,
		dwClsContext,
		win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770), // IID_IBaseFilter
	)
	if err != nil {
		panic(err)
	}
	return &IBaseFilter{
		IMediaFilter{
			IPersist{
				IUnknown: *iUnk,
			},
		},
	}
}

// Typically uses CLSCTX_INPROC_SERVER.
//
// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func CoCreateVideoMixingRenderer9(dwClsContext co.CLSCTX) *IBaseFilter {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0x51b4abf3, 0x748f, 0x4e3b, 0xa276, 0xc828330e926a), // CLSID_VideoMixingRenderer9
		nil,
		dwClsContext,
		win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770), // IID_IBaseFilter
	)
	if err != nil {
		panic(err)
	}
	return &IBaseFilter{
		IMediaFilter{
			IPersist{
				IUnknown: *iUnk,
			},
		},
	}
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
func (me *IBaseFilter) QueryIMFGetService() *IMFGetService {
	iUnk, err := me.QueryInterface(
		win.NewGuid(0xfa993888, 0x4383, 0x415a, 0xa930, 0xdd472a8cf6f7), // IID_IMFGetService
	)
	if err != nil {
		panic(err)
	}
	return &IMFGetService{
		IUnknown: *iUnk,
	}
}
