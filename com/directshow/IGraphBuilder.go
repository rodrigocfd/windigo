/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
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
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
	//
	// IGraphBuilder > IFilterGraph > IUnknown.
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

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *IGraphBuilder) CoCreateInstance(dwClsContext co.CLSCTX) *IGraphBuilder {
	ppv, err := win.CoCreateInstance(
		win.NewGuid(0xe436ebb3, 0x524f, 0x11ce, 0x9f53_0020af0ba770), // CLSID_FilterGraph
		nil,
		dwClsContext,
		win.NewGuid(0x56a868a9, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770)) // IID_IGraphBuilder

	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "CoCreateInstance/IGraphBuilder"))
	}
	me.Ppv = ppv
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-renderfile
func (me *IGraphBuilder) RenderFile(file string) *IGraphBuilder {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).AddFilter, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(file))), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.RenderFile").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *IGraphBuilder) Abort() *IGraphBuilder {
	ret, _, _ := syscall.Syscall(
		(*IGraphBuilderVtbl)(unsafe.Pointer(*me.Ppv)).Abort, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.Abort").Error())
	}
	return me
}
