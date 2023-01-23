//go:build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist interface {
	IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
	GetClassID() *win.GUID
}

type _IPersist struct{ IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPersist.Release().
func NewIPersist(base IUnknown) IPersist {
	return &_IPersist{IUnknown: base}
}

func (me *_IPersist) GetClassID() *win.GUID {
	clsid := &win.GUID{}
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPersist)(unsafe.Pointer(*me.Ptr())).GetClassID,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(clsid)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return clsid
	} else {
		panic(hr)
	}
}
