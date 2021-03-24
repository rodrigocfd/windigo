package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IDispatchVtbl struct {
	win.IUnknownVtbl
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch struct {
	win.IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
func (me *IDispatch) GetTypeInfoCount() int {
	pctinfo := uint32(0)
	ret, _, _ := syscall.Syscall(
		(*_IDispatchVtbl)(unsafe.Pointer(*me.Ppv)).GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pctinfo)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return int(pctinfo)
}
