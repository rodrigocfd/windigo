/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// IGraphBuilder > IFilterGraph > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
	IGraphBuilder struct{ IFilterGraph }

	IGraphBuilderVtbl struct {
		IFilterGraphVtbl
		Connect                 uintptr
		Render                  uintptr
		RenderFile              uintptr
		AddSourceFilter         uintptr
		SetLogFile              uintptr
		Abort                   uintptr
		ShouldOperationContinue uintptr
	}
)

// Typically uses CLSCTX_INPROC_SERVER.
//
// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIGraphBuilder(dwClsContext co.CLSCTX) *IGraphBuilder {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0xe436ebb3, 0x524f, 0x11ce, 0x9f53, 0x0020af0ba770), // CLSID_FilterGraph
		nil,
		dwClsContext,
		win.NewGuid(0x56a868a9, 0x0ad4, 0x11ce, 0xb03a, 0x0020af0ba770), // IID_IGraphBuilder
	)
	if err != nil {
		panic(err)
	}
	return &IGraphBuilder{
		IFilterGraph{
			IUnknown: *iUnk,
		},
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-connect
func (me *IGraphBuilder) Connect(pinOut, pinIn *IPin) {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Connect, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pinOut.Ppv)),
		uintptr(unsafe.Pointer(pinIn.Ppv)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.Connect"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-render
func (me *IGraphBuilder) Render(pinOut *IPin) {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Render, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pinOut.Ppv)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.Render"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-renderfile
func (me *IGraphBuilder) RenderFile(file string) {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).RenderFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(file))), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.RenderFile"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-setlogfile
func (me *IGraphBuilder) SetLogFile(hFile win.HFILE) {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).SetLogFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hFile), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.SetLogFile"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *IGraphBuilder) Abort() {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Abort, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.Abort"))
	}
}
