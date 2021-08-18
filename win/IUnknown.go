package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// IUnknown virtual table, base to all COM virtual tables.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

//------------------------------------------------------------------------------

// IUnknown COM interface, base to all COM interfaces.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	Ppv **IUnknownVtbl // Pointer to pointer to the COM virtual table.
}

// Returns a pointer to a pointer to the IUnknown virtual table, which can be
// cast into the specific virtual table type.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer IUnknown.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(
	rclsid co.CLSID, pUnkOuter *IUnknown,
	dwClsContext co.CLSCTX, riid co.IID) IUnknown {

	var ppv **IUnknownVtbl

	var ppOuterVtbl ***IUnknownVtbl = nil
	if pUnkOuter != nil {
		ppOuterVtbl = &pUnkOuter.Ppv
	}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(NewGuidFromClsid(rclsid))),
		uintptr(unsafe.Pointer(ppOuterVtbl)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(NewGuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppv)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IUnknown{ppv}
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer IUnknown.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef() IUnknown {
	syscall.Syscall((*me.Ppv).AddRef, 1,
		uintptr(unsafe.Pointer(me.Ppv)), 0, 0)
	return IUnknown{me.Ppv} // simply copy the pointer into a new object
}

// Returns a pointer to a pointer to the IUnknown virtual table, which can be
// cast into the specific virtual table type.
//
// ⚠️ You must defer IUnknown.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(riid co.IID) IUnknown {
	var ppvQueried **IUnknownVtbl
	ret, _, _ := syscall.Syscall((*me.Ppv).QueryInterface, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(NewGuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IUnknown{ppvQueried}
	} else {
		panic(hr)
	}
}

// Releases the COM pointer. Never fails, can be called any number of times.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() uint32 {
	var ret uintptr
	if me.Ppv != nil {
		ret, _, _ = syscall.Syscall((*me.Ppv).Release, 1,
			uintptr(unsafe.Pointer(me.Ppv)), 0, 0)
		if ret == 0 { // COM pointer was released
			me.Ppv = nil
		}
	}
	return uint32(ret)
}
