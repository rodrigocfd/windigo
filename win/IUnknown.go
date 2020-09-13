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
//
// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
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
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
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
