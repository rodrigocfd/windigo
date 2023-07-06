//go:build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IUnknown] COM interface, base to all COM interfaces.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown interface {
	// [QueryInterface] COM method.
	//
	// ⚠️ You must defer IUnknown.Release() on the returned COM object.
	//
	// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
	QueryInterface(riid co.IID) IUnknown

	// [AddRef] COM method.
	//
	// Creates a clone of the COM object.
	//
	// ⚠️ You must defer IUnknown.Release() on the returned COM object.
	//
	// # Example
	//
	//	var myObj IUnknown // initialized somewhere
	//
	//	otherObj := myObj.AddRef()
	//	defer otherObj.Release()
	//
	// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
	AddRef() IUnknown

	// [Release] COM method.
	//
	// If the underlying COM pointer is not nil, calls [Release] and sets it to
	// nil.
	//
	// Never fails, can safely be called any number of times.
	//
	// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
	Release() uint32

	// Returns the underlying pointer to pointer to the COM virtual table.
	//
	// If you want to check whether the object contains a valid, initialized
	// pointer, prefer using the com.IsObj() function, which is safer.
	//
	// Don't use this pointer to create a new COM object, this can cause a
	// resource leak.
	//
	// This method is used internally by the library, don't use unless you know
	// what you're doing.
	Ptr() **comvt.IUnknown
}

type _IUnknown struct{ ppv **comvt.IUnknown }

// Constructs an IUnknown object from a pointer to a pointer to its virtual
// table.
//
// This function is the building block of the COM interface object chain, and it
// should be used only if you're creating an object from a raw virtual table
// pointer.
//
// ⚠️ You must defer IUnknown.Release().
func NewIUnknown(ppv **comvt.IUnknown) IUnknown {
	return &_IUnknown{ppv: ppv}
}

func (me *_IUnknown) QueryInterface(riid co.IID) IUnknown {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN((*me.ppv).QueryInterface,
		uintptr(unsafe.Pointer(me.ppv)),
		uintptr(unsafe.Pointer(win.GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIUnknown(ppvQueried)
	} else {
		panic(hr)
	}
}

func (me *_IUnknown) AddRef() IUnknown {
	syscall.SyscallN((*me.ppv).AddRef,
		uintptr(unsafe.Pointer(me.ppv)))
	return NewIUnknown(me.ppv) // simply copy the pointer into a new object
}

func (me *_IUnknown) Release() uint32 {
	var refCount uintptr
	if me.Ptr() != nil {
		refCount, _, _ = syscall.SyscallN((*me.ppv).Release,
			uintptr(unsafe.Pointer(me.ppv)))
		if refCount == 0 { // COM pointer was released
			me.ppv = nil
		}
	}
	return uint32(refCount)
}

func (me *_IUnknown) Ptr() **comvt.IUnknown {
	return me.ppv
}
