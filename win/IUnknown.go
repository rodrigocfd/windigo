/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win/proc"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	IUnknown struct{ _IUnknownImpl }

	_IUnknownImpl struct{ uintptr }

	_IUnknownVtbl struct {
		QueryInterface uintptr
		AddRef         uintptr
		Release        uintptr
	}
)

// Returns a pointer to the virtual table.
func (me *_IUnknownImpl) pVtbl() unsafe.Pointer {
	// https://www.codeproject.com/Articles/633/Introduction-to-COM-What-It-Is-and-How-to-Use-It
	pptr := (*struct{ uintptr })(unsafe.Pointer(me.uintptr))
	return unsafe.Pointer(pptr.uintptr)
}

// Creates any COM interface, returning the base IUnknown.
func (me *_IUnknownImpl) coCreateInstancePtr(
	clsid *GUID, dwClsContext co.CLSCTX, iid *GUID) {

	if me.uintptr != 0 {
		panic("IUnknown already created, CoCreateInstance not called.")
	}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(clsid)), 0,
		uintptr(dwClsContext), uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&me.uintptr)), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "CoCreateInstance").Error())
	}
}

// Queries any COM interface, returning the base IUnknown.
//
// To retrieve the queried interface, cast the virtual table pointer.
func (me *_IUnknownImpl) queryInterface(iid *GUID) IUnknown {
	if me.uintptr == 0 {
		panic("Calling queryInterface on empty IUnknown.")
	}

	retIUnk := IUnknown{}

	vTbl := (*_IUnknownVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.QueryInterface, 3, me.uintptr,
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&retIUnk.uintptr)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(NewWinError(lerr, "IUnknown.QueryInterface").Error())
	}
	return retIUnk
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *_IUnknownImpl) AddRef() uint32 {
	if me.uintptr == 0 {
		panic("Calling AddRef on empty IUnknown.")
	}

	vTbl := (*_IUnknownVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.AddRef, 1, me.uintptr, 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *_IUnknownImpl) Release() uint32 {
	if me.uintptr == 0 {
		panic("Calling Release on empty IUnknown.")
	}

	vTbl := (*_IUnknownVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.Release, 1, me.uintptr, 0, 0)
	return uint32(ret)
}

// Returns a new GUID with the last uint64 member bytes flipped. This allows us
// to have easy literal declaration for GUID constants. With a literal
// declaration, the last uint64 will have its bits flipped.
//
// This function is called to make the conversion when needed internally.
// func cloneFlipLastUint64(guid *co.GUID) co.GUID {
// 	buf64 := [8]byte{}
// 	binary.BigEndian.PutUint64(buf64[:], guid.Data4)
// 	guidCopy := *guid
// 	guidCopy.Data4 = binary.LittleEndian.Uint64(buf64[:])
// 	return guidCopy
// }

// func cloneFlipLastUint64Clsid(clsid *co.CLSID) co.CLSID {
// 	return co.CLSID(cloneFlipLastUint64((*co.GUID)(clsid))) // specialization for CLSID
// }

// func cloneFlipLastUint64Iid(iid *co.IID) co.IID {
// 	return co.IID(cloneFlipLastUint64((*co.GUID)(iid))) // specialization for IID
// }
