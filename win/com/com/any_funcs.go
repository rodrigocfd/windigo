//go:build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [CLSIDFromProgID] function.
//
// Used to retrieve class IDs to create COM Automation objects. If the progId is
// invalid, error returns errco.CO_E_CLASSSTRING.
//
// # Example
//
//	clsId, _ := com.CLSIDFromProgID("Excel.Application")
//
//	mainObj := com.CoCreateInstance(
//		clsId, nil, comco.CLSCTX_SERVER, comco.IID_IUnknown)
//	defer mainObj.Release()
//
//	excel := mainObj.QueryInterface(automco.IID_IDispatch)
//	defer excel.Release()
//
// [CLSIDFromProgID]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	var guid win.GUID
	ret, _, _ := syscall.SyscallN(proc.CLSIDFromProgID.Addr(),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(progId))),
		uintptr(unsafe.Pointer(&guid)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

// [CoCreateGuid] function.
//
// This function creates a globally unique 128-bit integer.
//
// [CoCreateGuid]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateguid
func CoCreateGuid() win.GUID {
	var guid win.GUID
	ret, _, _ := syscall.SyscallN(proc.CoCreateGuid.Addr(),
		uintptr(unsafe.Pointer(&guid)))
	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return guid
	} else {
		panic(hr)
	}
}

// [CoCreateInstance] function.
//
// Creates a COM object from its CLSID + IID. The iUnkOuter is usually nil.
//
// Panics if the COM object cannot be created.
//
// ⚠️ You must defer IUnknown.Release() on the returned COM object. If iUnkOuter
// is not null, you must defer IUnknown.Release() on it too.
//
// # Example
//
//	comObject := shell.NewITaskbarList(
//		com.CoCreateInstance(
//			shellco.CLSID_TaskbarList, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			shellco.IID_ITaskbarList),
//	)
//	defer comObject.Release()
//
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(
	rclsid co.CLSID,
	iUnkOuter *IUnknown,
	dwClsContext comco.CLSCTX,
	riid co.IID) IUnknown {

	var ppvQueried **comvt.IUnknown

	var pppvOuter ***comvt.IUnknown
	if iUnkOuter != nil { // was the outer pointer requested?
		(*iUnkOuter).Release() // release if existing
		var ppvOuterBuf **comvt.IUnknown
		pppvOuter = &ppvOuterBuf // we'll request the outer pointer
	}

	ret, _, _ := syscall.SyscallN(proc.CoCreateInstance.Addr(),
		uintptr(unsafe.Pointer(win.GuidFromClsid(rclsid))),
		uintptr(unsafe.Pointer(pppvOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(win.GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		if iUnkOuter != nil {
			*iUnkOuter = NewIUnknown(*pppvOuter)
		}
		return NewIUnknown(ppvQueried)
	} else {
		panic(hr)
	}
}

// [CoInitializeEx] function.
//
// Loads the COM module. This needs to be done only once in your application.
// Typically uses COINIT_APARTMENTTHREADED.
//
// ⚠️ You must defer CoUninitialize().
//
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit comco.COINIT) {
	ret, _, _ := syscall.SyscallN(proc.CoInitializeEx.Addr(),
		0, uintptr(coInit))
	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// [CoUninitialize] function.
//
// [CoUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.SyscallN(proc.CoUninitialize.Addr())
}

// This helper function returns true if the COM object is not nil, and contains
// an initialized internal pointer.
func IsObj(obj IUnknown) bool {
	return obj != nil && obj.Ptr() != nil
}

// [OleInitialize] function.
//
// ⚠️ You must defer OleUninitialize().
//
// [OleInitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-oleinitialize
func OleInitialize() {
	ret, _, _ := syscall.SyscallN(proc.OleInitialize.Addr(),
		0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// [OleUninitialize] function.
//
// [OleUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-oleuninitialize
func OleUninitialize() {
	syscall.SyscallN(proc.OleUninitialize.Addr())
}
