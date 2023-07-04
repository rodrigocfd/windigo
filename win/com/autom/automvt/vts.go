//go:build windows

package automvt

import (
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
)

// [IDispatch] virtual table.
//
// [IDispatch]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch struct {
	comvt.IUnknown
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

// [IErrorLog] virtual table.
//
// [IErrorLog]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ierrorlog
type IErrorLog struct {
	comvt.IUnknown
	AddError uintptr
}

// [IPropertyBag] virtual table.
//
// [IPropertyBag]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-ipropertybag
type IPropertyBag struct {
	comvt.IUnknown
	Read  uintptr
	Write uintptr
}

// [ITypeInfo] virtual table.
//
// [ITypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct {
	comvt.IUnknown
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
