package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
type IBaseFilter interface {
	IMediaFilter

	// ⚠️ You must defer IEnumPins.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-enumpins
	EnumPins() IEnumPins

	// ⚠️ You must defer IPin.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-findpin
	FindPin(id string) (IPin, bool)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-joinfiltergraph
	JoinFilterGraph(graph IFilterGraph, name string) error

	// ⚠️ You must defer IFilterGraph.Release() on PGraph field of the info
	// object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-queryfilterinfo
	QueryFilterInfo(info *FILTER_INFO)

	// Returns false if the method is not supported.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ibasefilter-queryvendorinfo
	QueryVendorInfo() (string, bool)
}

type _IBaseFilter struct{ IMediaFilter }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IBaseFilter.Release().
//
// Example for an Enhanced Video Renderer:
//
//		evh := dshow.NewIBaseFilter(
//			com.CoCreateInstance(
//				dshowco.CLSID_EnhancedVideoRenderer, nil,
//				comco.CLSCTX_INPROC_SERVER,
//				dshowco.IID_IBaseFilter),
//		)
//		defer evh.Release()
//
// Example for a Video Media Renderer 9:
//
//		vmr9 := dshow.NewIBaseFilter(
//			com.CoCreateInstance(
//				dshowco.CLSID_VideoMixingRenderer9, nil,
//				comco.CLSCTX_INPROC_SERVER,
//				dshowco.IID_IBaseFilter),
//		)
//		defer vmr9.Release()
func NewIBaseFilter(base com.IUnknown) IBaseFilter {
	return &_IBaseFilter{IMediaFilter: NewIMediaFilter(base)}
}

func (me *_IBaseFilter) EnumPins() IEnumPins {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).EnumPins, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumPins(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IBaseFilter) FindPin(id string) (IPin, bool) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).FindPin, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(id))),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(com.NewIUnknown(ppQueried)), true
	} else if hr == errco.VFW_E_NOT_FOUND {
		return nil, false
	} else {
		panic(hr)
	}
}

func (me *_IBaseFilter) JoinFilterGraph(graph IFilterGraph, name string) error {
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

func (me *_IBaseFilter) QueryFilterInfo(info *FILTER_INFO) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).QueryFilterInfo, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(info)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IBaseFilter) QueryVendorInfo() (string, bool) {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IBaseFilter)(unsafe.Pointer(*me.Ptr())).QueryVendorInfo, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name, true
	} else if hr == errco.E_NOTIMPL {
		return "", false
	} else {
		panic(hr)
	}
}
