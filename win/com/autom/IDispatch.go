package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

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
// Example:
//
//  var iDisp autom.IDispatch // initialized somewhere
//
//  tyInfo := iDisp.GetTypeInfo(win.LCID_SYSTEM_DEFAULT)
//  defer tyInfo.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (me *IDispatch) GetTypeInfo(lcid win.LCID) ITypeInfo {
	var ppQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall6(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfo, 4,
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

// If the object provides type information, this number is 1; otherwise the
// number is 0.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
func (me *IDispatch) GetTypeInfoCount() int {
	var pctInfo uint32
	ret, _, _ := syscall.Syscall(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pctInfo)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(pctInfo)
	} else {
		panic(hr)
	}
}
