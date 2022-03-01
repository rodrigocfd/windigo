package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/errco"
)

// IPropertyBag COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ipropertybag
type IPropertyBag struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPropertyBag.Release().
func NewIPropertyBag(base com.IUnknown) IPropertyBag {
	return IPropertyBag{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ipropertybag-write
func (me *IPropertyBag) Write(propName string, value *VARIANT) {
	ret, _, _ := syscall.Syscall(
		(*automvt.IPropertyBag)(unsafe.Pointer(*me.Ptr())).Write, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(propName))),
		uintptr(unsafe.Pointer(value)),
	)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
