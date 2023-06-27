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

// [IErrorLog] COM interface.
//
// [IErrorLog]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ierrorlog
type IErrorLog interface {
	com.IUnknown

	// [AddError] COM method.
	//
	// [AddError]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ierrorlog-adderror
	AddError(propName string, exceps []EXCEPINFO)
}

type _IErrorLog struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IErrorLog.Release().
func NewIErrorLog(base com.IUnknown) IErrorLog {
	return &_IErrorLog{IUnknown: base}
}

func (me *_IErrorLog) AddError(propName string, exceps []EXCEPINFO) {
	ret, _, _ := syscall.SyscallN(
		(*automvt.IPropertyBag)(unsafe.Pointer(*me.Ptr())).Write,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(propName))),
		uintptr(unsafe.Pointer(&exceps[0])))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
