package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
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

// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIFilterGraph(dwClsContext co.CLSCTX) IFilterGraph {
	clsidFilterGraph := win.NewGuid(0xe436ebb3, 0x524f, 0x11ce, 0x9f53, 0x0020af0ba770)
	iidIFilterGraph := win.NewGuid(0x56a8689f, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770)

	iUnk, err := win.CoCreateInstance(
		clsidFilterGraph, nil, dwClsContext, iidIFilterGraph)
	if err != nil {
		panic(err)
	}
	return IFilterGraph{IUnknown: iUnk}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-addfilter
func (me *IFilterGraph) AddFilter(pFilter *IBaseFilter, name string) error {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).AddFilter, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pFilter.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(name))))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return lerr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-disconnect
func (me *IFilterGraph) Disconnect(ppin *IPin) {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).Disconnect, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(ppin.Ppv)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
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

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
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

	if lerr := err.ERROR(ret); lerr == err.VFW_E_NOT_FOUND {
		return IBaseFilter{}, false
	} else if lerr != err.S_OK {
		panic(lerr)
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
func (me *IFilterGraph) Reconnect(ppin *IPin) {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).Reconnect, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(ppin.Ppv)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-removefilter
func (me *IFilterGraph) RemoveFilter(pFilter *IBaseFilter) {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).RemoveFilter, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pFilter.Ppv)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-setdefaultsyncsource
func (me *IFilterGraph) SetDefaultSyncSource() {
	ret, _, _ := syscall.Syscall(
		(*_IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).SetDefaultSyncSource, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
