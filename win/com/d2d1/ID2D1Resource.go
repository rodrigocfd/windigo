package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1resource
type ID2D1Resource struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1Resource.Release().
func NewID2D1Resource(base win.IUnknown) ID2D1Resource {
	return ID2D1Resource{IUnknown: base}
}

// ⚠️ You must defer ID2D1Factory.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1resource-getfactory
func (me *ID2D1Resource) GetFactory() ID2D1Factory {
	var ppvQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1Resource)(unsafe.Pointer(*me.Ptr())).GetFactory, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)),
		0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1Factory(ppvQueried)
	} else {
		panic(hr)
	}
}
