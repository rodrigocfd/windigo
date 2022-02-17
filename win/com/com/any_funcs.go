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

// Used to retrieve class IDs to create COM Automation objects.
//
// If the progId is invalid, error returns errco.CO_E_CLASSSTRING.
//
// Example:
//
//  clsId, _ := com.CLSIDFromProgID("Excel.Application")
//
//  mainObj := com.CoCreateInstance(
//      clsId, nil, comco.CLSCTX_SERVER, comco.IID_IUnknown)
//  defer mainObj.Release()
//
//  excel := mainObj.QueryInterface(automco.IID_IDispatch)
//  defer excel.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	var guid win.GUID
	ret, _, _ := syscall.Syscall(proc.CLSIDFromProgID.Addr(), 2,
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(progId))),
		uintptr(unsafe.Pointer(&guid)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

// Typically uses CLSCTX_INPROC_SERVER. Panics if the COM object cannot be
// created.
//
// ⚠️ The returned IUnknown can be used to construct another COM object; either
// way you must defer Release().
//
// Example for an ordinary COM object:
//
//  comObject := shell.NewITaskbarList(
//      com.CoCreateInstance(
//          shellco.CLSID_TaskbarList, nil,
//          comco.CLSCTX_INPROC_SERVER,
//          shellco.IID_ITaskbarList),
//  )
//  defer comObject.Release()
//
// Example for COM Automation:
//
//  clsId, _ := com.CLSIDFromProgID("Excel.Application")
//  root := com.CoCreateInstance(
//      clsId, nil, comco.CLSCTX_SERVER, comco.IID_IUnknown)
//  defer root.Release()
//
//  excel := autom.NewIDispatch(
//      root.QueryInterface(automco.IID_IDispatch))
//  defer excel.Release()
//
//  for _, f := range excel.ListFunctions() {
//      println(f.Name, f.NumParams, f.FuncKind, f.Flags)
//  }
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(
	rclsid co.CLSID, pUnkOuter *IUnknown,
	dwClsContext comco.CLSCTX, riid co.IID) IUnknown {

	var ppvQueried **comvt.IUnknown

	var ppvOuter ***comvt.IUnknown
	if pUnkOuter != nil { // was the outer pointer requested?
		pUnkOuter.Release()
		ppvOuter = &pUnkOuter.ppv
	}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(win.GuidFromClsid(rclsid))),
		uintptr(unsafe.Pointer(ppvOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(win.GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IUnknown{ppv: ppvQueried}
	} else {
		panic(hr)
	}
}

// Loads the COM module. This needs to be done only once in your application.
// Typically uses COINIT_APARTMENTTHREADED.
//
// ⚠️ You must defer CoUninitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit comco.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(coInit), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
