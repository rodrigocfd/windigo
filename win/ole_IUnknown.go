package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Typically uses CLSCTX_INPROC_SERVER. Panics if the COM object cannot be
// created.
//
// ⚠️ The returned IUnknown can be used to construct another COM object; either
// way you must defer Release().
//
// Example for an ordinary COM object:
//
//  comObject := shell.NewITaskbarList(
//      win.CoCreateInstance(
//          shellco.CLSID_TaskbarList, nil,
//          co.CLSCTX_INPROC_SERVER,
//          shellco.IID_ITaskbarList),
//  )
//  defer comObject.Release()
//
// Example for COM Automation:
//
//  clsId, _ := win.CLSIDFromProgID("Excel.Application")
//  root := win.CoCreateInstance(
//      clsId, nil, co.CLSCTX_SERVER, co.IID_IUNKNOWN)
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
	dwClsContext co.CLSCTX, riid co.IID) IUnknown {

	var ppvQueried **IUnknownVtbl

	var ppvOuter ***IUnknownVtbl
	if pUnkOuter != nil { // was the outer pointer requested?
		pUnkOuter.Release()
		ppvOuter = &pUnkOuter.ppv
	}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(GuidFromClsid(rclsid))),
		uintptr(unsafe.Pointer(ppvOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IUnknown{ppv: ppvQueried}
	} else {
		panic(hr)
	}
}

//------------------------------------------------------------------------------

// IUnknown virtual table, base to all COM virtual tables.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

// IUnknown COM interface, base to all COM interfaces.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct{ ppv **IUnknownVtbl }

// Returns the underlying pointer to the COM virtual table.
func (me *IUnknown) Ptr() **IUnknownVtbl {
	return me.ppv
}

// ⚠️ You must defer IUnknown.Release() on the new object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef() IUnknown {
	syscall.Syscall((*me.ppv).AddRef, 1,
		uintptr(unsafe.Pointer(me.ppv)), 0, 0)
	return IUnknown{ppv: me.ppv} // simply copy the pointer into a new object
}

// ⚠️ The returned pointer must be used to construct a COM object; you must
// defer its Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(riid co.IID) IUnknown {
	var ppvQueried **IUnknownVtbl
	ret, _, _ := syscall.Syscall((*me.ppv).QueryInterface, 3,
		uintptr(unsafe.Pointer(me.ppv)),
		uintptr(unsafe.Pointer(GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IUnknown{ppv: ppvQueried}
	} else {
		panic(hr)
	}
}

// Releases the COM pointer. Never fails, can be called any number of times.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() uint32 {
	ret := uintptr(0)
	if me.Ptr() != nil {
		ret, _, _ = syscall.Syscall((*me.ppv).Release, 1,
			uintptr(unsafe.Pointer(me.ppv)), 0, 0)
		if ret == 0 { // COM pointer was released
			me.ppv = nil
		}
	}
	return uint32(ret)
}
