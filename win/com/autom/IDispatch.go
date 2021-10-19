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
type IDispatch struct{ win.IUnknown }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IDispatch.Release().
func NewIDispatch(ptr win.IUnknownPtr) IDispatch {
	return IDispatch{
		IUnknown: win.NewIUnknown(ptr),
	}
}

// ⚠️ You must defer ITypeInfo.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (me *IDispatch) GetTypeInfo(lcid win.LCID) ITypeInfo {
	var ppQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall6(
		(*IDispatchVtbl)(unsafe.Pointer(*me.Ptr())).GetTypeInfo, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, uintptr(lcid),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return ITypeInfo{
			win.NewIUnknown(ppQueried),
		}
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
func (me *IDispatch) GetTypeInfoCount() int {
	var pctInfo uint32
	ret, _, _ := syscall.Syscall(
		(*IDispatchVtbl)(unsafe.Pointer(*me.Ptr())).GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pctInfo)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(pctInfo)
	} else {
		panic(hr)
	}
}
