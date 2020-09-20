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
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
	//
	// IFilterGraph > IUnknown.
	IFilterGraph struct{ win.IUnknown }

	IFilterGraphVtbl struct {
		win.IUnknownVtbl
		AddFilter            uintptr
		RemoveFilter         uintptr
		EnumFilters          uintptr
		FindFilterByName     uintptr
		ConnectDirect        uintptr
		Reconnect            uintptr
		Disconnect           uintptr
		SetDefaultSyncSource uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-addfilter
func (me *IFilterGraph) AddFilter(
	pFilter IBaseFilter, name string) *IFilterGraph {

	ret, _, _ := syscall.Syscall(
		(*IFilterGraphVtbl)(unsafe.Pointer(*me.Ppv)).AddFilter, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pFilter.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(name))))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFilterGraph.AddFilter").Error())
	}
	return me
}
