//go:build windows

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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
type IGraphBuilder interface {
	IFilterGraph

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
	Abort()

	// ⚠️ You must defer IBaseFilter.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-addsourcefilter
	AddSourceFilter(fileName, filterName string) IBaseFilter

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-connect
	Connect(pinOut, pinIn IPin)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-render
	Render(pinOut IPin)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-renderfile
	RenderFile(file string) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-setlogfile
	SetLogFile(hFile win.HFILE)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-shouldoperationcontinue
	ShouldOperationContinue() bool
}

type _IGraphBuilder struct{ IFilterGraph }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IGraphBuilder.Release().
//
// Example:
//
//		gb := dshow.NewIGraphBuilder(
//			com.CoCreateInstance(
//				dshowco.CLSID_FilterGraph, nil,
//				comco.CLSCTX_INPROC_SERVER,
//				dshowco.IID_IGraphBuilder),
//		)
//		defer gb.Release()
func NewIGraphBuilder(base com.IUnknown) IGraphBuilder {
	return &_IGraphBuilder{IFilterGraph: NewIFilterGraph(base)}
}

func (me *_IGraphBuilder) Abort() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).Abort, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IGraphBuilder) AddSourceFilter(
	fileName, filterName string) IBaseFilter {

	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).AddSourceFilter, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(filterName))),
		uintptr(unsafe.Pointer(&ppvQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIBaseFilter(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IGraphBuilder) Connect(pinOut, pinIn IPin) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).Connect, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pinOut.Ptr())),
		uintptr(unsafe.Pointer(pinIn.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IGraphBuilder) Render(pinOut IPin) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).Render, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pinOut.Ptr())), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IGraphBuilder) RenderFile(file string) error {
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

func (me *_IGraphBuilder) SetLogFile(hFile win.HFILE) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IGraphBuilder)(unsafe.Pointer(*me.Ptr())).SetLogFile, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hFile), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IGraphBuilder) ShouldOperationContinue() bool {
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
