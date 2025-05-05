//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win/co"
)

// [ITypeInfo] COM interface.
//
// [ITypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct{ IUnknown }

// Returns the unique COM interface identifier.
func (*ITypeInfo) IID() co.IID {
	return co.IID_ITypeInfo
}

// [AddressOfMember] method.
//
// [AddressOfMember]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-addressofmember
func (me *ITypeInfo) AddressOfMember(memberId MEMBERID, invokeKind co.INVOKEKIND) (uintptr, error) {
	var addr uintptr
	ret, _, _ := syscall.SyscallN(
		(*vt.ITypeInfo)(unsafe.Pointer(*me.Ppvt())).AddressOfMember,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId), uintptr(invokeKind),
		uintptr(unsafe.Pointer(&addr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return addr, nil
	} else {
		return 0, hr
	}
}

// // [CreateInstance] method.
// //
// // [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-createinstance
// func (me *ITypeInfo) CreateInstance(releaser *Releaser, riid co.IID) (*IUnknown, error) {
// 	var ppvtQueried **vt.IUnknown
// 	guidIid := win.GuidFrom(riid)
//
// 	ret, _, _ := syscall.SyscallN(
// 		(*vt.ITypeInfo)(unsafe.Pointer(*me.Ppvt())).CreateInstance,
// 		uintptr(unsafe.Pointer(me.Ppvt())), 0, // don't query pUnkOuter
// 		uintptr(unsafe.Pointer(&guidIid)),
// 		uintptr(unsafe.Pointer(&ppvtQueried)))
//
// 	if hr := errco.ERROR(ret); hr == errco.S_OK {
// 		pUnk := NewIUnknown(ppvtQueried)
// 		releaser.Add(pUnk)
// 		return pUnk, nil
// 	} else {
// 		return nil, hr
// 	}
// }

// [GetDllEntry] method.
//
// [GetDllEntry]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdllentry
func (me *ITypeInfo) GetDllEntry(
	memberId MEMBERID,
	invokeKind co.INVOKEKIND,
) (_ITypeInfoDll, error) {
	var dllName, name BSTR
	defer dllName.SysFreeString()
	defer name.SysFreeString()
	var ordinal16 uint16

	ret, _, _ := syscall.SyscallN(
		(*vt.ITypeInfo)(unsafe.Pointer(*me.Ppvt())).GetDllEntry,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId), uintptr(invokeKind),
		uintptr(unsafe.Pointer(&dllName)),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(&ordinal16)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return _ITypeInfoDll{
			DllName: dllName.String(),
			Name:    name.String(),
			Ordinal: uint(ordinal16),
		}, nil
	} else {
		return _ITypeInfoDll{}, hr
	}
}

type _ITypeInfoDll struct {
	DllName string
	Name    string
	Ordinal uint
}

// [GetDocumentation] method.
//
// [GetDocumentation]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdocumentation
func (me *ITypeInfo) GetDocumentation(memberId MEMBERID) (_ITypeInfoDoc, error) {
	var name, docStr, helpFile BSTR
	defer name.SysFreeString()
	defer docStr.SysFreeString()
	defer helpFile.SysFreeString()
	var helpCtx uint32

	ret, _, _ := syscall.SyscallN(
		(*vt.ITypeInfo)(unsafe.Pointer(*me.Ppvt())).GetDocumentation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(&docStr)),
		uintptr(unsafe.Pointer(&helpCtx)),
		uintptr(unsafe.Pointer(&helpFile)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return _ITypeInfoDoc{
			Name:        name.String(),
			DocString:   docStr.String(),
			HelpContext: uint(helpCtx),
			HelpFile:    helpFile.String(),
		}, nil
	} else {
		return _ITypeInfoDoc{}, hr
	}
}

type _ITypeInfoDoc struct {
	Name        string
	DocString   string
	HelpContext uint
	HelpFile    string
}
