package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ipropertybag
type IPropertyBag interface {
	com.IUnknown

	// The errorLog can be nil.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ipropertybag-read
	Read(propName string, errorLog IErrorLog) VARIANT

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ipropertybag-write
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
	ret, _, _ := syscall.Syscall6(
		(*automvt.IPropertyBag)(unsafe.Pointer(*me.Ptr())).Read, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(propName))),
		uintptr(unsafe.Pointer(&buf)),
		uintptr(unsafe.Pointer(errorLog.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return buf
	} else {
		panic(hr)
	}
}

func (me *_IPropertyBag) Write(propName string, value *VARIANT) {
	ret, _, _ := syscall.Syscall(
		(*automvt.IPropertyBag)(unsafe.Pointer(*me.Ptr())).Write, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(propName))),
		uintptr(unsafe.Pointer(value)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
