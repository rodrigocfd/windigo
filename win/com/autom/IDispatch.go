package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

// IDispatch virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatchVtbl struct {
	win.IUnknownVtbl
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
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer ITypeInfo.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (me *IDispatch) GetTypeInfo(lcid win.LCID) ITypeInfo {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(
		(*IDispatchVtbl)(unsafe.Pointer(*me.Ppv)).GetTypeInfo, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, uintptr(lcid),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return ITypeInfo{
			win.IUnknown{Ppv: ppQueried},
		}
	} else {
		panic(hr)
	}
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
