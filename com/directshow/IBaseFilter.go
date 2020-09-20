/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package directshow

import (
	"windigo/co"
	"windigo/win"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
	//
	// IBaseFilter > IMediaFilter > IPersist > IUnknown.
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

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
//
// https://docs.microsoft.com/en-us/windows/win32/medfound/using-the-directshow-evr-filter
func (me *IBaseFilter) CoCreateEnhancedVideoRenderer(dwClsContext co.CLSCTX) *IBaseFilter {
	ppv, err := win.CoCreateInstance(
		win.NewGuid(0xfa10746c, 0x9b63, 0x4b6c, 0xbc49_fc300ea5f256), // CLSID_EnhancedVideoRenderer
		nil,
		dwClsContext,
		win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IBaseFilter

	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "CoCreateInstance/EnhancedVideoRenderer"))
	}
	me.Ppv = (**win.IUnknownVtbl)(ppv)
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
//
// https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func (me *IBaseFilter) CoCreateVideoMixingRenderer9(dwClsContext co.CLSCTX) *IBaseFilter {
	ppv, err := win.CoCreateInstance(
		win.NewGuid(0x51b4abf3, 0x748f, 0x4e3b, 0xa276_c828330e926a), // CLSID_VideoMixingRenderer9
		nil,
		dwClsContext,
		win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IBaseFilter

	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "CoCreateInstance/VideoMixingRenderer9"))
	}
	me.Ppv = (**win.IUnknownVtbl)(ppv)
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
//
// https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
func (me *IBaseFilter) QueryIMFGetService() IMFGetService {
	ppv, err := me.QueryInterface(
		win.NewGuid(0xfa993888, 0x4383, 0x415a, 0xa930_dd472a8cf6f7)) // IID_IMFGetService
	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "IBaseFilter.QueryIMFGetService"))
	}
	return IMFGetService{
		win.IUnknown{
			Ppv: ppv,
		},
	}
}
