package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ibindctx
type IBindCtx struct{ IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IBindCtx.Release().
func NewIBindCtx(base IUnknown) IBindCtx {
	return IBindCtx{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-releaseboundobjects
func (me *IBindCtx) ReleaseBoundObjects() {
	ret, _, _ := syscall.Syscall(
		(*comvt.IBindCtx)(unsafe.Pointer(*me.Ptr())).ReleaseBoundObjects, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-revokeobjectparam
func (me *IBindCtx) RevokeObjectParam(key string) {
	ret, _, _ := syscall.Syscall(
		(*comvt.IBindCtx)(unsafe.Pointer(*me.Ptr())).RevokeObjectParam, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(key))),
		0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
