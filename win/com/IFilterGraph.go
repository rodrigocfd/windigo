/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// IFilterGraph > IUnknown.
type IFilterGraph struct {
	IUnknown
}

type iFilterGraphVtbl struct {
	iUnknownVtbl
	AddFilter            uintptr
	RemoveFilter         uintptr
	EnumFilters          uintptr
	FindFilterByName     uintptr
	ConnectDirect        uintptr
	Reconnect            uintptr
	Disconnect           uintptr
	SetDefaultSyncSource uintptr
}

// func (me *IFilterGraph) coCreateInstance() {
// 	if me.lpVtbl == 0 { // if not created yet
// 		me.IUnknown.coCreateInstance(
// 			&co.Guid_IFilterGraph, &co.Guid_IGraphBuilder)
// 	}
// }

// func (me *IFilterGraph) SetDefaultSyncSource() {
// 	me.coCreateInstance()
// 	lpVtbl := (*iFilterGraphVtbl)(unsafe.Pointer(me.lpVtbl))
// 	ret, _, _ := syscall.Syscall(lpVtbl.SetDefaultSyncSource, 1,
// 		uintptr(unsafe.Pointer(me)), 0, 0)

// 	lerr := co.ERROR(ret)
// 	if lerr != co.ERROR_S_OK {
// 		me.Release() // free resource
// 		panic(lerr.Format("IFilterGraph.SetDefaultSyncSource failed."))
// 	}
// }
