/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
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
func (me *_IBaseFilterImpl) CoCreateEnhancedVideoRenderer(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0xfa10746c, 0x9b63, 0x4b6c, 0xbc49_fc300ea5f256), // CLSID_EnhancedVideoRenderer
		dwClsContext,
		_Win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IBaseFilter
}

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
//
// https://docs.microsoft.com/en-us/windows/win32/directshow/video-mixing-renderer-filter-9
func (me *_IBaseFilterImpl) CoCreateVideoMixingRenderer9(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x51b4abf3, 0x748f, 0x4e3b, 0xa276_c828330e926a), // CLSID_VideoMixingRenderer9
		dwClsContext,
		_Win.NewGuid(0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IBaseFilter
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

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
	//
	// IGraphBuilder > IFilterGraph > IUnknown.
	IGraphBuilder struct{ _IGraphBuilder }

	_IGraphBuilder struct{ _IFilterGraphImpl }

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
func (me *_IGraphBuilder) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0xe436ebb3, 0x524f, 0x11ce, 0x9f53_0020af0ba770), // CLSID_FilterGraph
		dwClsContext,
		_Win.NewGuid(0x56a868a9, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IGraphBuilder
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *IGraphBuilder) Abort() {
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

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
	//
	// ITaskbarList > IUnknown.
	ITaskbarList struct{ _ITaskbarListImpl }

	_ITaskbarListImpl struct{ _IUnknownImpl }

	_ITaskbarListVtbl struct {
		_IUnknownVtbl
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *_ITaskbarListImpl) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		dwClsContext,
		_Win.NewGuid(0x56fdf342, 0xfd6d, 0x11d0, 0x958a_006097c9a090)) // IID_ITaskbarList
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *_ITaskbarListImpl) ActivateTab(hwnd HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.ActivateTab, 1, uintptr(me.ptr), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList.ActivateTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *_ITaskbarListImpl) SetActiveAlt(hwnd HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetActiveAlt, 1, uintptr(me.ptr), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList.SetActiveAlt").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
	//
	// ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList2 struct{ _ITaskbarList2Impl }

	_ITaskbarList2Impl struct{ _ITaskbarListImpl }

	_ITaskbarList2Vtbl struct {
		_ITaskbarListVtbl
		MarkFullscreenWindow uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *_ITaskbarList2Impl) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		dwClsContext,
		_Win.NewGuid(0x602d4995, 0xb13a, 0x429b, 0xa66e_1935e44f4317)) // IID_ITaskbarList2
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *_ITaskbarList2Impl) MarkFullscreenWindow(
	hwnd HWND, fFullScreen bool) {

	vTbl := (*_ITaskbarList2Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.MarkFullscreenWindow, 3, uintptr(me.ptr),
		uintptr(hwnd), _Win.BoolToUintptr(fFullScreen))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList2.MarkFullscreenWindow").Error())
	}
}

//------------------------------------------------------------------------------

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
	//
	// ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList3 struct{ _ITaskbarList3Impl }

	_ITaskbarList3Impl struct{ _ITaskbarList2Impl }

	_ITaskbarList3Vtbl struct {
		_ITaskbarList2Vtbl
		SetProgressValue      uintptr
		SetProgressState      uintptr
		RegisterTab           uintptr
		UnregisterTab         uintptr
		SetTabOrder           uintptr
		SetTabActive          uintptr
		ThumbBarAddButtons    uintptr
		ThumbBarUpdateButtons uintptr
		ThumbBarSetImageList  uintptr
		SetOverlayIcon        uintptr
		SetThumbnailTooltip   uintptr
		SetThumbnailClip      uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *_ITaskbarList3Impl) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		_Win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090), // CLSID_TaskbarList
		dwClsContext,
		_Win.NewGuid(0xea1afb91, 0x9e28, 0x4b86, 0x90e9_9e9f8a5eefaf)) // IID_ITaskbarList3
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(
	hwnd HWND, ullCompleted, ullTotal uint64) {

	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall6(vTbl.SetProgressValue, 4, uintptr(me.ptr),
		uintptr(hwnd), uintptr(ullCompleted), uintptr(ullTotal), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetProgressValue").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hwnd HWND, tbpFlags co.TBPF) {
	vTbl := (*_ITaskbarList3Vtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetProgressState, 3, uintptr(me.ptr),
		uintptr(hwnd), uintptr(tbpFlags))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "ITaskbarList3.SetProgressState").Error())
	}
}
