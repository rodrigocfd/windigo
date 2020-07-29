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
	baseIGraphBuilder struct{ baseIFilterGraph }

	// IGraphBuilder > IFilterGraph > IUnknown.
	IGraphBuilder struct{ baseIGraphBuilder }

	vtbIGraphBuilder struct {
		vtbIFilterGraph
		Connect                 uintptr
		Render                  uintptr
		RenderFile              uintptr
		AddSourceFilter         uintptr
		SetLogFile              uintptr
		Abort                   uintptr
		ShouldOperationContinue uintptr
	}
)

func (me *baseIGraphBuilder) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.baseIUnknown.coCreateInstance(
		&co.CLSID_FilterGraph, dwClsContext, &co.IID_IGraphBuilder)
}

func (me *IGraphBuilder) Abort() {
	ret, _, _ := syscall.Syscall(
		(*vtbIGraphBuilder)(unsafe.Pointer(me.pVtb())).Abort, 1,
		uintptr(unsafe.Pointer(me.uintptr)), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("IGraphBuilder.Abort failed."))
	}
}
