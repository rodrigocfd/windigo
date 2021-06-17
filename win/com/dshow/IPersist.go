package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IPersistVtbl struct {
	win.IUnknownVtbl
	GetClassID uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct {
	win.IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
func (me *IPersist) GetClassID() *win.GUID {
	clsid := win.GUID{}
	ret, _, _ := syscall.Syscall(
		(*_IPersistVtbl)(unsafe.Pointer(*me.Ppv)).GetClassID, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&clsid)), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
	return &clsid
}
