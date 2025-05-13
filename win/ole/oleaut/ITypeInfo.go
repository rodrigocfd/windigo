//go:build windows

package oleaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [ITypeInfo] COM interface.
//
// [ITypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITypeInfo) IID() co.IID {
	return co.IID_ITypeInfo
}

// [AddressOfMember] method.
//
// [AddressOfMember]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-addressofmember
func (me *ITypeInfo) AddressOfMember(
	memberId MEMBERID,
	invokeKind co.INVOKEKIND,
) (uintptr, error) {
	var addr uintptr
	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).AddressOfMember,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId), uintptr(invokeKind),
		uintptr(unsafe.Pointer(&addr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return addr, nil
	} else {
		return 0, hr
	}
}

// [CreateInstance] method. Not implemented as a method of [ITypeInfo] because
// Go doesn't support generic methods.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-createinstance
func CreateInstance[T any, P ole.ComCtor[T]](
	iTypeInfo *ITypeInfo,
	releaser *ole.Releaser,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	riidGuid := win.GuidFrom(pObj.IID())

	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*iTypeInfo.Ppvt())).CreateInstance,
		uintptr(unsafe.Pointer(iTypeInfo.Ppvt())),
		uintptr(unsafe.Pointer(&riidGuid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj.Set(ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetDllEntry] method.
//
// [GetDllEntry]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdllentry
func (me *ITypeInfo) GetDllEntry(
	memberId MEMBERID,
	invokeKind co.INVOKEKIND,
) (ITypeInfoDllEntry, error) {
	var dllName, name BSTR
	defer dllName.SysFreeString()
	defer name.SysFreeString()
	var ordinal16 uint16

	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetDllEntry,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId), uintptr(invokeKind),
		uintptr(unsafe.Pointer(&dllName)),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(&ordinal16)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ITypeInfoDllEntry{
			DllName: dllName.String(),
			Name:    name.String(),
			Ordinal: uint(ordinal16),
		}, nil
	} else {
		return ITypeInfoDllEntry{}, hr
	}
}

// Returned by [ITypeInfo.GetDllEntry].
type ITypeInfoDllEntry struct {
	DllName string
	Name    string
	Ordinal uint
}

// [GetDocumentation] method.
//
// [GetDocumentation]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdocumentation
func (me *ITypeInfo) GetDocumentation(memberId MEMBERID) (ITypeInfoDoc, error) {
	var name, docStr, helpFile BSTR
	defer name.SysFreeString()
	defer docStr.SysFreeString()
	defer helpFile.SysFreeString()
	var helpCtx uint32

	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetDocumentation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(&docStr)),
		uintptr(unsafe.Pointer(&helpCtx)),
		uintptr(unsafe.Pointer(&helpFile)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ITypeInfoDoc{
			Name:        name.String(),
			DocString:   docStr.String(),
			HelpContext: uint(helpCtx),
			HelpFile:    helpFile.String(),
		}, nil
	} else {
		return ITypeInfoDoc{}, hr
	}
}

// Returned by [ITypeInfo.GetDocumentation].
type ITypeInfoDoc struct {
	Name        string
	DocString   string
	HelpContext uint
	HelpFile    string
}

type _ITypeInfoVt struct {
	ole.IUnknownVt
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
