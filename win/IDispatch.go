package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/errco"
)

// IDispatch virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatchVtbl struct {
	IUnknownVtbl
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

//------------------------------------------------------------------------------

// IDispatch COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch struct {
	IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
func (me *IDispatch) GetTypeInfoCount() int {
	var pctinfo uint32
	ret, _, _ := syscall.Syscall(
		(*IDispatchVtbl)(unsafe.Pointer(*me.Ppv)).GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pctinfo)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(pctinfo)
	} else {
		panic(hr)
	}
}
