package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
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
func CoCreateIGraphBuilder(dwClsContext co.CLSCTX) (IGraphBuilder, error) {
	clsidFilterGrapth := win.NewGuid(0xe436ebb3, 0x524f, 0x11ce, 0x9f53, 0x0020af0ba770)
	iidIGraphBuilder := win.NewGuid(0x56a868a9, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770)

	iUnk, lerr := win.CoCreateInstance(
		clsidFilterGrapth, nil, dwClsContext, iidIGraphBuilder)
	if lerr != nil {
		return IGraphBuilder{}, lerr
	}
	return IGraphBuilder{
		IFilterGraph{IUnknown: iUnk},
	}, nil
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
func (me *IGraphBuilder) QueryIBasicAudio() (IBasicAudio, error) {
	iidIBasicAudio := win.NewGuid(0x56a868b3, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770)

	iUnk, lerr := me.QueryInterface(iidIBasicAudio)
	if lerr != nil {
		return IBasicAudio{}, lerr
	}
	return IBasicAudio{
		IDispatch{IUnknown: iUnk},
	}, nil
}

// ⚠️ You must defer Release().
//
// Calls IUnknown.QueryInterface() to return IMediaControl.
func (me *IGraphBuilder) QueryIMediaControl() (IMediaControl, error) {
	iidIMediaControl := win.NewGuid(0x56a868b1, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770)

	iUnk, lerr := me.QueryInterface(iidIMediaControl)
	if lerr != nil {
		return IMediaControl{}, lerr
	}
	return IMediaControl{
		IDispatch{IUnknown: iUnk},
	}, nil
}

// ⚠️ You must defer Release().
//
// Calls IUnknown.QueryInterface() to return IMediaSeeking.
func (me *IGraphBuilder) QueryIMediaSeeking() (IMediaSeeking, error) {
	iidIMediaSeeking := win.NewGuid(0x36b73880, 0xc2c8, 0x11cf, 0x8b46, 0x00805f6cef60)

	iUnk, lerr := me.QueryInterface(iidIMediaSeeking)
	if lerr != nil {
		return IMediaSeeking{}, lerr
	}
	return IMediaSeeking{IUnknown: iUnk}, nil
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
