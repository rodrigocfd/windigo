package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
type IGraphBuilder struct{ IFilterGraph }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IGraphBuilder.Release().
//
// Example:
//
//  gb := dshow.NewIGraphBuilder(
//      win.CoCreateInstance(
//          dshowco.CLSID_FilterGraph, nil,
//          co.CLSCTX_INPROC_SERVER,
//          dshowco.IID_IGraphBuilder),
//  )
//  defer gb.Release()
func NewIGraphBuilder(base win.IUnknown) IGraphBuilder {
	return IGraphBuilder{IFilterGraph: NewIFilterGraph(base)}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *IGraphBuilder) Abort() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).Abort, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// ⚠️ You must defer IBaseFilter.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-addsourcefilter
func (me *IGraphBuilder) AddSourceFilter(
	fileName, filterName string) IBaseFilter {

	var ppvQueried win.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).AddSourceFilter, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(filterName))),
		uintptr(unsafe.Pointer(&ppvQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIBaseFilter(ppvQueried)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-connect
func (me *IGraphBuilder) Connect(pinOut, pinIn *IPin) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).Connect, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pinOut.Ptr())),
		uintptr(unsafe.Pointer(pinIn.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-render
func (me *IGraphBuilder) Render(pinOut *IPin) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).Render, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pinOut.Ptr())), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-renderfile
func (me *IGraphBuilder) RenderFile(file string) error {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).RenderFile, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(file))), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-setlogfile
func (me *IGraphBuilder) SetLogFile(hFile win.HFILE) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).SetLogFile, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hFile), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-shouldoperationcontinue
func (me *IGraphBuilder) ShouldOperationContinue() bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).ShouldOperationContinue, 1,
		uintptr(unsafe.Pointer(me.Ptr())), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
