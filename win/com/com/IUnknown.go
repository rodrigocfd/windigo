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
type IUnknown struct{ ppv **comvt.IUnknown }

// Returns the underlying pointer to the COM virtual table.
func (me *IUnknown) Ptr() **comvt.IUnknown {
	return me.ppv
}

// ⚠️ You must defer IUnknown.Release() on the new object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef() IUnknown {
	syscall.Syscall((*me.ppv).AddRef, 1,
		uintptr(unsafe.Pointer(me.ppv)), 0, 0)
	return IUnknown{ppv: me.ppv} // simply copy the pointer into a new object
}

// ⚠️ The returned pointer must be used to construct a COM object; you must
// defer its Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(riid co.IID) IUnknown {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall((*me.ppv).QueryInterface, 3,
		uintptr(unsafe.Pointer(me.ppv)),
		uintptr(unsafe.Pointer(win.GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IUnknown{ppv: ppvQueried}
	} else {
		panic(hr)
	}
}

// Releases the COM pointer. Never fails, can be called any number of times.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() uint32 {
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
