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

// IGraphBuilder > IFilterGraph > IUnknown.
type IGraphBuilder struct {
	IFilterGraph
}

type iGraphBuilderVtbl struct {
	iFilterGraphVtbl
	Connect                 uintptr
	Render                  uintptr
	RenderFile              uintptr
	AddSourceFilter         uintptr
	SetLogFile              uintptr
	Abort                   uintptr
	ShouldOperationContinue uintptr
}

func (me *IGraphBuilder) coCreateInstance() {
	if me.lpVtbl == 0 { // if not created yet
		me.IUnknown.coCreateInstance(
			&co.CLSID_FilterGraph, &co.IID_IGraphBuilder)
	}
}

func (me *IGraphBuilder) Abort() {
	me.coCreateInstance()
	lpVtbl := (*iGraphBuilderVtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.Abort, 1,
		uintptr(unsafe.Pointer(me)), 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("IGraphBuilder.Abort failed."))
	}
}
