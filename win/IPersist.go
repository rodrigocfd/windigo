package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/errco"
)

// IPersist virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersistVtbl struct {
	IUnknownVtbl
	GetClassID uintptr
}

//------------------------------------------------------------------------------

// IPersist COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct {
	IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ipersist-getclassid
func (me *IPersist) GetClassID() *GUID {
	clsid := GUID{}
	ret, _, _ := syscall.Syscall(
		(*IPersistVtbl)(unsafe.Pointer(*me.Ppv)).GetClassID, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&clsid)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return &clsid
	} else {
		panic(hr)
	}
}
