/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"encoding/binary"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
	"wingows/win/proc"
)

type (
	_IUnknownPtr struct{ uintptr } // IUnknown pointer itself, which has a pointer to virtual table
	_IUnknown    struct{ uintptr } // container, which has a pointer to actual IUnknown

	// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	IUnknown struct{ _IUnknown }

	_IUnknownVtbl struct {
		QueryInterface uintptr
		AddRef         uintptr
		Release        uintptr
	}
)

// Returns a pointer to IUnknown virtual table.
func (me *_IUnknown) pVtb() *_IUnknownVtbl {
	// https://www.codeproject.com/Articles/633/Introduction-to-COM-What-It-Is-and-How-to-Use-It
	ptrIUnk := (*_IUnknownPtr)(unsafe.Pointer(me.uintptr))
	ptrVtb := (*_IUnknownVtbl)(unsafe.Pointer(ptrIUnk.uintptr))
	return ptrVtb
}

// Creates any COM interface, returning the base IUnknown.
// To retrieve the other interface itself, cast the inner lpVtbl.
func (me *_IUnknown) coCreateInstance(
	clsid *co.CLSID, dwClsContext co.CLSCTX, iid *co.IID) {

	if me.uintptr != 0 {
		panic("IUnknown already created, CoCreateInstance not called.")
	}

	clsidFlip := cloneFlipLastUint64Clsid(clsid)
	iidFlip := cloneFlipLastUint64Iid(iid)

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(&clsidFlip)), 0,
		uintptr(dwClsContext), uintptr(unsafe.Pointer(&iidFlip)),
		uintptr(unsafe.Pointer(&me.uintptr)), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "CoCreateInstance").Error())
	}
}

// Queries any COM interface, returning the base IUnknown.
//
// To retrieve the other interface itself, cast the inner lpVtbl.
func (me *_IUnknown) queryInterface(iid *co.IID) IUnknown {
	if me.uintptr == 0 {
		panic("Calling queryInterface on empty IUnknown.")
	}

	iidFlip := cloneFlipLastUint64Iid(iid)
	retIUnk := IUnknown{}

	ret, _, _ := syscall.Syscall(me.pVtb().QueryInterface, 3,
		uintptr(unsafe.Pointer(me.uintptr)),
		uintptr(unsafe.Pointer(&iidFlip)),
		uintptr(unsafe.Pointer(&retIUnk.uintptr)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IUnknown.QueryInterface").Error())
	}
	return retIUnk
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *_IUnknown) AddRef() uint32 {
	if me.uintptr == 0 {
		panic("Calling AddRef on empty IUnknown.")
	}

	ret, _, _ := syscall.Syscall(me.pVtb().AddRef, 1,
		uintptr(unsafe.Pointer(me.uintptr)), 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *_IUnknown) Release() uint32 {
	if me.uintptr == 0 {
		panic("Calling Release on empty IUnknown.")
	}

	ret, _, _ := syscall.Syscall(me.pVtb().Release, 1,
		uintptr(unsafe.Pointer(me.uintptr)), 0, 0)
	return uint32(ret)
}

// Returns a new GUID with the last uint64 member bytes flipped. This allows us
// to have easy literal declaration for GUID constants. With a literal
// declaration, the last uint64 will have its bits flipped.
//
// This function is called to make the conversion when needed internally.
func cloneFlipLastUint64(guid *co.GUID) co.GUID {
	buf64 := [8]byte{}
	binary.BigEndian.PutUint64(buf64[:], guid.Data4)
	guidCopy := *guid
	guidCopy.Data4 = binary.LittleEndian.Uint64(buf64[:])
	return guidCopy
}

func cloneFlipLastUint64Clsid(clsid *co.CLSID) co.CLSID {
	return co.CLSID(cloneFlipLastUint64((*co.GUID)(clsid))) // specialization for CLSID
}

func cloneFlipLastUint64Iid(iid *co.IID) co.IID {
	return co.IID(cloneFlipLastUint64((*co.GUID)(iid))) // specialization for IID
}
