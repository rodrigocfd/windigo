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
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	IUnknown struct {
		Ppv **IUnknownVtbl // Pointer to pointer to the COM virtual table.
	}

	IUnknownVtbl struct {
		QueryInterface uintptr
		AddRef         uintptr
		Release        uintptr
	}
)

// Queries any COM interface, returning the base IUnknown pointer.
//
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(riid *GUID) (**IUnknownVtbl, co.ERROR) {
	var ppvObject **IUnknownVtbl = nil
	ret, _, _ := syscall.Syscall((*me.Ppv).QueryInterface, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvObject)))
	return ppvObject, co.ERROR(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef() uint32 {
	ret, _, _ := syscall.Syscall((*me.Ppv).AddRef, 1,
		uintptr(unsafe.Pointer(me.Ppv)), 0, 0)
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() uint32 {
	ret, _, _ := syscall.Syscall((*me.Ppv).Release, 1,
		uintptr(unsafe.Pointer(me.Ppv)), 0, 0)
	return uint32(ret)
}
