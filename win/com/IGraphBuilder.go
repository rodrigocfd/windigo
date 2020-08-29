/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"wingows/co"
	"wingows/win"
)

type (
	_IGraphBuilder struct{ _IFilterGraphImpl }

	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
	//
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
	me.coCreateInstancePtr(
		&co.CLSID_FilterGraph, dwClsContext, &co.IID_IGraphBuilder)
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-igraphbuilder-abort
func (me *IGraphBuilder) Abort() {
	vTbl := (*_IGraphBuilderVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.Abort, 1, me.uintptr, 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IGraphBuilder.Abort").Error())
	}
}
