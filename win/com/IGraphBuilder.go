/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"unsafe"
	"wingows/co"
)

type (
	_IGraphBuilder struct{ _IFilterGraph }

	// IGraphBuilder > IFilterGraph > IUnknown.
	IGraphBuilder struct{ _IGraphBuilder }

	_IGraphBuilderVtbl struct {
		_IFilterGraphVtbl
		Connect                 uintptr
		Render                  uintptr
		RenderFile              uintptr
		AddSourceFilter         uintptr
		SetLogFile              uintptr
		Abort                   uintptr
		ShouldOperationContinue uintptr
	}
)

func (me *_IGraphBuilder) CoCreateInstance(dwClsContext co.CLSCTX) {
	me._IUnknown.coCreateInstance(
		&co.CLSID_FilterGraph, dwClsContext, &co.IID_IGraphBuilder)
}

func (me *IGraphBuilder) Abort() {
	ret, _, _ := syscall.Syscall(
		(*_IGraphBuilderVtbl)(unsafe.Pointer(me.pVtb())).Abort, 1,
		uintptr(unsafe.Pointer(me.uintptr)), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("IGraphBuilder.Abort failed."))
	}
}
