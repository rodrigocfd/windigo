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

// IUnknown COM interface, base to all COM interfaces.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown interface {
	// ⚠️ You must defer IUnknown.Release() on the returned COM object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
	QueryInterface(riid co.IID) IUnknown

	// Creates a clone of the COM object.
	//
	// ⚠️ You must defer IUnknown.Release() on the returned COM object.
	//
	// Example:
	//
	//		var myObj IUnknown // initialized somewhere
	//
	//		otherObj := myObj.AddRef()
	//		defer otherObj.Release()
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
	AddRef() IUnknown

	// Releases the COM pointer and sets the internal pointer to nil.
	//
	// Never fails, can be called any number of times.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
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
	ret, _, _ := syscall.Syscall((*me.ppv).QueryInterface, 3,
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
	syscall.Syscall((*me.ppv).AddRef, 1,
		uintptr(unsafe.Pointer(me.ppv)), 0, 0)
	return NewIUnknown(me.ppv) // simply copy the pointer into a new object
}

func (me *_IUnknown) Release() uint32 {
	ret := uintptr(0)
	if me.Ptr() != nil {
		ret, _, _ = syscall.Syscall((*me.ppv).Release, 1,
			uintptr(unsafe.Pointer(me.ppv)), 0, 0)
		if ret == 0 { // COM pointer was released
			me.ppv = nil
		}
	}
	return uint32(ret)
}

func (me *_IUnknown) Ptr() **comvt.IUnknown {
	return me.ppv
}
