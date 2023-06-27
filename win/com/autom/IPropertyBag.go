//go:build windows

package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IPropertyBag] COM interface.
//
// [IPropertyBag]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ipropertybag
type IPropertyBag interface {
	com.IUnknown

	// [Read] COM method.
	//
	// The errorLog can be nil.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the returned object.
	//
	// [Read]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ipropertybag-read
	Read(propName string, errorLog IErrorLog) VARIANT

	// [Write] COM method.
	//
	// [Write]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ipropertybag-write
	Write(propName string, value *VARIANT)
}

type _IPropertyBag struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPropertyBag.Release().
func NewIPropertyBag(base com.IUnknown) IPropertyBag {
	return &_IPropertyBag{IUnknown: base}
}

func (me *_IPropertyBag) Read(propName string, errorLog IErrorLog) VARIANT {
	buf := NewVariantEmpty()
	ret, _, _ := syscall.SyscallN(
		(*automvt.IPropertyBag)(unsafe.Pointer(*me.Ptr())).Read,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(propName))),
		uintptr(unsafe.Pointer(&buf)),
		uintptr(unsafe.Pointer(errorLog.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return buf
	} else {
		panic(hr)
	}
}

func (me *_IPropertyBag) Write(propName string, value *VARIANT) {
	ret, _, _ := syscall.SyscallN(
		(*automvt.IPropertyBag)(unsafe.Pointer(*me.Ptr())).Write,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(propName))),
		uintptr(unsafe.Pointer(value)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
