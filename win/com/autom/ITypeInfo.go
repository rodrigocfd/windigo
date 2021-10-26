package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// ITypeInfo COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct{ win.IUnknown }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer ITypeInfo.Release().
func NewITypeInfo(ptr win.IUnknownPtr) ITypeInfo {
	return ITypeInfo{
		IUnknown: win.NewIUnknown(ptr),
	}
}

// ⚠️ You must defer ITypeInfo.ReleaseFuncDesc() on the returned object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getfuncdesc
func (me *ITypeInfo) GetFuncDesc(index int) *FUNCDESC {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).GetFuncDesc, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*FUNCDESC)(unsafe.Pointer(pv))
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer ITypeInfo.ReleaseVarDesc() on the returned object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getvardesc
func (me *ITypeInfo) GetVarDesc(index int) *VARDESC {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).GetVarDesc, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*VARDESC)(unsafe.Pointer(pv))
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *ITypeInfo) ReleaseFuncDesc(funcDesc *FUNCDESC) {
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).ReleaseFuncDesc, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(funcDesc)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasevardesc
func (me *ITypeInfo) ReleaseVarDesc(varDesc *VARDESC) {
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).ReleaseVarDesc, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(varDesc)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
