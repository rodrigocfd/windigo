package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IGraphBuilderVtbl struct {
	_IFilterGraphVtbl
	Connect                 uintptr
	Render                  uintptr
	RenderFile              uintptr
	AddSourceFilter         uintptr
	SetLogFile              uintptr
	Abort                   uintptr
	ShouldOperationContinue uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
type IGraphBuilder struct {
	IFilterGraph // Base IFilterGraph > IUnknown.
}

// Calls IUnknown.CoCreateInstance() to return IGraphBuilder.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIGraphBuilder(dwClsContext co.CLSCTX) IGraphBuilder {
	iUnk := win.CoCreateInstance(
		dshowco.CLSID_FilterGraph, nil, dwClsContext,
		dshowco.IID_IGraphBuilder)
	return IGraphBuilder{
		IFilterGraph{IUnknown: iUnk},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *IGraphBuilder) Abort() {
	ret, _, _ := syscall.Syscall(
		(*_IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Abort, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-connect
func (me *IGraphBuilder) Connect(pinOut, pinIn *IPin) {
	ret, _, _ := syscall.Syscall(
		(*_IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Connect, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pinOut.Ppv)),
		uintptr(unsafe.Pointer(pinIn.Ppv)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// ⚠️ You must defer Release().
//
// Calls IUnknown.QueryInterface() to return IBasicAudio.
func (me *IGraphBuilder) QueryIBasicAudio() IBasicAudio {
	iUnk := me.QueryInterface(dshowco.IID_IBasicAudio)
	return IBasicAudio{
		win.IDispatch{IUnknown: iUnk},
	}
}

// ⚠️ You must defer Release().
//
// Calls IUnknown.QueryInterface() to return IMediaControl.
func (me *IGraphBuilder) QueryIMediaControl() IMediaControl {
	iUnk := me.QueryInterface(dshowco.IID_IMediaControl)
	return IMediaControl{
		win.IDispatch{IUnknown: iUnk},
	}
}

// ⚠️ You must defer Release().
//
// Calls IUnknown.QueryInterface() to return IMediaSeeking.
func (me *IGraphBuilder) QueryIMediaSeeking() IMediaSeeking {
	iUnk := me.QueryInterface(dshowco.IID_IMediaSeeking)
	return IMediaSeeking{IUnknown: iUnk}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-render
func (me *IGraphBuilder) Render(pinOut *IPin) {
	ret, _, _ := syscall.Syscall(
		(*_IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Render, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pinOut.Ppv)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-renderfile
func (me *IGraphBuilder) RenderFile(file string) error {
	ret, _, _ := syscall.Syscall(
		(*_IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).RenderFile, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(file))), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return lerr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-setlogfile
func (me *IGraphBuilder) SetLogFile(hFile win.HFILE) {
	ret, _, _ := syscall.Syscall(
		(*_IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).SetLogFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hFile), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
