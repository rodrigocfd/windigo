package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

// ITypeInfo virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfoVtbl struct {
	win.IUnknownVtbl
	GetTypeAttr          uintptr
	GetTypeComp          uintptr
	GetFuncDesc          uintptr
	GetVarDesc           uintptr
	GetNames             uintptr
	GetRefTypeOfImplType uintptr
	GetImplTypeFlags     uintptr
	GetIDsOfNames        uintptr
	Invoke               uintptr
	GetDocumentation     uintptr
	GetDllEntry          uintptr
	GetRefTypeInfo       uintptr
	AddressOfMember      uintptr
	CreateInstance       uintptr
	GetMops              uintptr
	GetContainingTypeLib uintptr
	ReleaseTypeAttr      uintptr
	ReleaseFuncDesc      uintptr
	ReleaseVarDesc       uintptr
}

//------------------------------------------------------------------------------

// ITypeInfo COM interface.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer ITypeInfo.ReleaseFuncDesc() on the returned object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getfuncdesc
func (me *ITypeInfo) GetFuncDesc(index int) *FUNCDESC {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*ITypeInfoVtbl)(unsafe.Pointer(*me.Ppv)).GetFuncDesc, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(index), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*FUNCDESC)(unsafe.Pointer(pv))
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *ITypeInfo) ReleaseFuncDesc(funcDesc *FUNCDESC) {
	ret, _, _ := syscall.Syscall(
		(*ITypeInfoVtbl)(unsafe.Pointer(*me.Ppv)).ReleaseFuncDesc, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(funcDesc)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
