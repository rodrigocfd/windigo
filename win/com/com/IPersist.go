package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// IPersist COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct{ IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPersist.Release().
func NewIPersist(base IUnknown) IPersist {
	return IPersist{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
func (me *IPersist) GetClassID() *win.GUID {
	clsid := &win.GUID{}
	ret, _, _ := syscall.Syscall(
		(*comvt.IPersist)(unsafe.Pointer(*me.Ptr())).GetClassID, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(clsid)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return clsid
	} else {
		panic(hr)
	}
}
