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
	"wingows/win/proc"
)

type (
	baseIUnknown struct{ uintptr }

	// IUnknown is the base to all COM interfaces.
	IUnknown struct{ baseIUnknown }

	vtbIUnknown struct {
		QueryInterface uintptr
		AddRef         uintptr
		Release        uintptr
	}
)

// Creates any COM interface, returning the base IUnknown.
// To retrieve the other interface itself, cast the inner lpVtbl.
func (me *baseIUnknown) coCreateInstance(
	clsid *co.CLSID, dwClsContext co.CLSCTX, iid *co.IID) {

	if me.uintptr != 0 {
		panic("Trying to CoCreateInstance() an IUnknown already created.")
	}

	clsidFlip := cloneFlipLastUint64Clsid(clsid)
	iidFlip := cloneFlipLastUint64Iid(iid)
	retIUnk := &baseIUnknown{}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(&clsidFlip)), 0,
		uintptr(dwClsContext), uintptr(unsafe.Pointer(&iidFlip)),
		uintptr(unsafe.Pointer(&retIUnk)), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(lerr.Format("CoCreateInstance failed."))
	}
	me.uintptr = retIUnk.uintptr
}

// Queries any COM interface, returning the base IUnknown.
// To retrieve the other interface itself, cast the inner lpVtbl.
func (me *baseIUnknown) queryInterface(iid *co.IID) IUnknown {
	if me.uintptr == 0 {
		panic("Calling queryInterface on empty IUnknown.")
	}

	lpVtbl := (*vtbIUnknown)(unsafe.Pointer(me.uintptr))
	iidFlip := cloneFlipLastUint64Iid(iid)
	retIUnk := &baseIUnknown{}

	ret, _, _ := syscall.Syscall(lpVtbl.AddRef, 3,
		uintptr(unsafe.Pointer(me)), uintptr(unsafe.Pointer(&iidFlip)),
		uintptr(unsafe.Pointer(&retIUnk.uintptr)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("IUnknown.QueryInterface failed."))
	}
	return IUnknown{*retIUnk}
}

func (me *baseIUnknown) AddRef() uint32 {
	if me.uintptr == 0 {
		panic("Calling AddRef on empty IUnknown.")
	}

	pVtbTy := (*vtbIUnknown)(unsafe.Pointer(me.uintptr))
	ret, _, _ := syscall.Syscall(pVtbTy.AddRef, 1,
		uintptr(unsafe.Pointer(me)), 0, 0)
	return uint32(ret)
}

func (me *baseIUnknown) Release() uint32 {
	if me.uintptr == 0 {
		panic("Calling Release on empty IUnknown.")
	}

	pVtbTy := (*vtbIUnknown)(unsafe.Pointer(me.uintptr))
	ret, _, _ := syscall.Syscall(pVtbTy.Release, 1,
		uintptr(unsafe.Pointer(&me.uintptr)), 0, 0)
	return uint32(ret)
}

// Returns a new GUID with the last uint64 member bytes flipped.
// This allows us to have easy literal declaration for GUID constants.
// With a literal declaration, the last uint64 will have its bits flipped.
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
