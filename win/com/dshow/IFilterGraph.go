package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IFilterGraphVtbl struct {
	win.IUnknownVtbl
	AddFilter            uintptr
	RemoveFilter         uintptr
	EnumFilters          uintptr
	FindFilterByName     uintptr
	ConnectDirect        uintptr
	Reconnect            uintptr
	Disconnect           uintptr
	SetDefaultSyncSource uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
type IFilterGraph struct {
	win.IUnknown // Base IUnknown.
}

// Calls IUnknown.CoCreateInstance() to return IFilterGraph.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIFilterGraph(dwClsContext co.CLSCTX) IFilterGraph {
	iUnk := win.CoCreateInstance(
		dshowco.CLSID_FilterGraph, nil, dwClsContext,
		dshowco.IID_IFilterGraph)
	return IFilterGraph{IUnknown: iUnk}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-addfilter
func (me *IFilterGraph) AddFilter(filter *IBaseFilter, name string) error {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).AddFilter, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(filter.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(name))))

	if err := errco.ERROR(ret); err != errco.S_OK {
		return err
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-connectdirect
func (me *IFilterGraph) ConnectDirect(pinOut, pinIn *IPin, pmt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.Syscall6(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).AddFilter, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pinOut.Ppv)),
		uintptr(unsafe.Pointer(pinIn.Ppv)),
		uintptr(unsafe.Pointer(pmt)), 0, 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-disconnect
func (me *IFilterGraph) Disconnect(pin *IPin) {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).Disconnect, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pin.Ppv)), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-enumfilters
func (me *IFilterGraph) EnumFilters() IEnumFilters {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).EnumFilters, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
	return IEnumFilters{
		win.IUnknown{Ppv: ppvQueried},
	}
}

// If the filter was not found, returns false.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-findfilterbyname
func (me *IFilterGraph) FindFilterByName(pName string) (IBaseFilter, bool) {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).FindFilterByName, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if err := errco.ERROR(ret); err == errco.VFW_E_NOT_FOUND {
		return IBaseFilter{}, false
	} else if err != errco.S_OK {
		panic(err)
	}

	return IBaseFilter{
		IMediaFilter{
			IPersist{
				win.IUnknown{Ppv: ppvQueried},
			},
		},
	}, true
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-reconnect
func (me *IFilterGraph) Reconnect(pin *IPin) {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).Reconnect, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pin.Ppv)), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-removefilter
func (me *IFilterGraph) RemoveFilter(filter *IBaseFilter) {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).RemoveFilter, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(filter.Ppv)), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-setdefaultsyncsource
func (me *IFilterGraph) SetDefaultSyncSource() {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).SetDefaultSyncSource, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}
