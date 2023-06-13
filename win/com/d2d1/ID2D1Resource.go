//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [ID2D1Resource] COM interface.
//
// [ID2D1Resource]: https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1resource
type ID2D1Resource interface {
	com.IUnknown

	// [GetFactory] COM method.
	//
	// ⚠️ You must defer ID2D1Factory.Release() on the returned object.
	//
	// [GetFactory]: https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1resource-getfactory
	GetFactory() ID2D1Factory
}

type _ID2D1Resource struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1Resource.Release().
func NewID2D1Resource(base com.IUnknown) ID2D1Resource {
	return &_ID2D1Resource{IUnknown: base}
}

func (me *_ID2D1Resource) GetFactory() ID2D1Factory {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1Resource)(unsafe.Pointer(*me.Ptr())).GetFactory,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1Factory(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}
