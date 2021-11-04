package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
type IBaseFilter struct{ IMediaFilter }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IBaseFilter.Release().
//
// Example for an Enhanced Video Renderer:
//
//  evh := dshow.NewIBaseFilter(
//      win.CoCreateInstance(
//          dshowco.CLSID_EnhancedVideoRenderer, nil,
//          co.CLSCTX_INPROC_SERVER,
//          dshowco.IID_IBaseFilter),
//  )
//  defer evh.Release()
//
// Example for a Video Media Renderer 9:
//
//  vmr9 := dshow.NewIBaseFilter(
//      win.CoCreateInstance(
//          dshowco.CLSID_VideoMixingRenderer9, nil,
//          co.CLSCTX_INPROC_SERVER,
//          dshowco.IID_IBaseFilter),
//  )
//  defer vmr9.Release()
func NewIBaseFilter(base win.IUnknown) IBaseFilter {
	return IBaseFilter{IMediaFilter: NewIMediaFilter(base)}
}

// ⚠️ You must defer IEnumPins.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-enumpins
func (me *IBaseFilter) EnumPins() IEnumPins {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).EnumPins, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumPins(ppQueried)
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer IPin.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-findpin
func (me *IBaseFilter) FindPin(id string) (IPin, bool) {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).FindPin, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(id))),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(ppQueried), true
	} else if hr == errco.VFW_E_NOT_FOUND {
		return IPin{}, false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-joinfiltergraph
func (me *IBaseFilter) JoinFilterGraph(
	graph *IFilterGraph, name string) error {

	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).JoinFilterGraph, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(graph.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(name))))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// ⚠️ You must defer IFilterGraph.Release() on PGraph field.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-queryfilterinfo
func (me *IBaseFilter) QueryFilterInfo(info *FILTER_INFO) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).QueryFilterInfo, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(info)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// Returns false if the method is not supported.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-queryvendorinfo
func (me *IBaseFilter) QueryVendorInfo() (string, bool) {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).QueryVendorInfo, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		win.CoTaskMemFree(pv)
		return name, true
	} else if hr == errco.E_NOTIMPL {
		return "", false
	} else {
		panic(hr)
	}
}
