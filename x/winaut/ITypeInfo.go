//go:build windows

package winaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/coaut"
)

// [ITypeInfo] COM interface.
//
// [ITypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct{ win.IUnknown }

type _ITypeInfoVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITypeInfo) IID() *co.IID {
	return &coaut.IID_ITypeInfo
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *ITypeInfo) AddRef(releaser *win.OleReleaser) *ITypeInfo {
	return utl.OleNewFromAddRef[*ITypeInfo](me, releaser)
}

// [AddressOfMember] method.
//
// [AddressOfMember]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-addressofmember
func (me *ITypeInfo) AddressOfMember(
	memberId MEMBERID,
	invokeKind coaut.INVOKEKIND,
) (uintptr, error) {
	var addr uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).AddressOfMember,
		me.Ppvt(),
		uintptr(memberId),
		uintptr(invokeKind),
		uintptr(unsafe.Pointer(&addr)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return addr, nil
}

// [CreateInstance] method.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-createinstance
func (me *ITypeInfo) CreateInstance(
	releaser *win.OleReleaser,
	unkOuter *win.IUnknown,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).CreateInstance,
		me.Ppvt(),
		utl.OlePpvtOrNil(unkOuter),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [GetContainingTypeLib] method.
//
// Returns the type library and its index.
//
// [GetContainingTypeLib]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getcontainingtypelib
func (me *ITypeInfo) GetContainingTypeLib(releaser *win.OleReleaser) (*ITypeLib, int, error) {
	var ppvtQueried uintptr
	var index uint32

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetContainingTypeLib,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&index)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, 0, hr
	}
	pObj := utl.OleNew[*ITypeLib](ppvtQueried, releaser)
	return pObj, int(index), nil
}

// [GetDllEntry] method.
//
// [GetDllEntry]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdllentry
func (me *ITypeInfo) GetDllEntry(
	memberId MEMBERID,
	invokeKind coaut.INVOKEKIND,
) (TypeInfoDllEntry, error) {
	var dllName, name BSTR
	defer dllName.SysFreeString()
	defer name.SysFreeString()
	var ordinal16 uint16

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetDllEntry,
		me.Ppvt(),
		uintptr(memberId),
		uintptr(invokeKind),
		uintptr(unsafe.Pointer(&dllName)),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(&ordinal16)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return TypeInfoDllEntry{}, hr
	}
	return TypeInfoDllEntry{
		DllName: dllName.String(),
		Name:    name.String(),
		Ordinal: int(ordinal16),
	}, nil
}

// Returned by [ITypeInfo.GetDllEntry].
type TypeInfoDllEntry struct {
	DllName string
	Name    string
	Ordinal int
}

// [GetDocumentation] method.
//
// [GetDocumentation]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdocumentation
func (me *ITypeInfo) GetDocumentation(memberId MEMBERID) (TypeInfoDoc, error) {
	var name, docStr, helpFile BSTR
	defer name.SysFreeString()
	defer docStr.SysFreeString()
	defer helpFile.SysFreeString()
	var helpCtx uint32

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetDocumentation,
		me.Ppvt(),
		uintptr(memberId),
		uintptr(unsafe.Pointer(&name)),
		uintptr(unsafe.Pointer(&docStr)),
		uintptr(unsafe.Pointer(&helpCtx)),
		uintptr(unsafe.Pointer(&helpFile)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return TypeInfoDoc{}, hr
	}
	return TypeInfoDoc{
		Name:        name.String(),
		DocString:   docStr.String(),
		HelpContext: int(helpCtx),
		HelpFile:    helpFile.String(),
	}, nil
}

// Returned by [ITypeInfo.GetDocumentation].
type TypeInfoDoc struct {
	Name        string
	DocString   string
	HelpContext int
	HelpFile    string
}

// [GetFuncDesc] method.
//
// The [win.OleReleaser] is responsible for freeing the resources by calling
// [ReleaseFuncDesc].
//
// Example:
//
//	var nfo *winaut.ITypeInfo // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	funcDesc, _ := nfo.GetFuncDesc(rel, 0)
//	println(funcDesc.Memid)
//
// [GetFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getfuncdesc
// [ReleaseFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *ITypeInfo) GetFuncDesc(releaser *win.OleReleaser, index int) (*FuncDescData, error) {
	var pFuncDesc *FUNCDESC
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetFuncDesc,
		me.Ppvt(),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&pFuncDesc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}
	pData := &FuncDescData{pFuncDesc, me}
	releaser.Add(pData)
	return pData, nil
}

// [GetIDsOfNames] method.
//
// [GetIDsOfNames]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getidsofnames
func (me *ITypeInfo) GetIDsOfNames(names ...string) ([]MEMBERID, error) {
	strPtrs := make([]*uint16, 0, len(names))
	for _, name := range names {
		strPtrs = append(strPtrs, wstr.EncodeToPtr(name))
	}

	memIds := make([]MEMBERID, len(names)) // to be returned

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetIDsOfNames,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&strPtrs[0])),
		uintptr(uint32(len(names))),
		uintptr(unsafe.Pointer(&memIds[0])))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}
	return memIds, nil
}

// [GetImplTypeFlags] method.
//
// [GetImplTypeFlags]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getimpltypeflags
func (me *ITypeInfo) GetImplTypeFlags(index int) (coaut.IMPLTYPEFLAG, error) {
	var flags coaut.IMPLTYPEFLAG
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetImplTypeFlags,
		me.Ppvt(),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&flags)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return coaut.IMPLTYPEFLAG(0), hr
	}
	return flags, nil
}

// [GetMops] method.
//
// [GetMops]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getmops
func (me *ITypeInfo) GetMops(memberId MEMBERID) (string, error) {
	var mops BSTR
	defer mops.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).GetMops,
		me.Ppvt(),
		uintptr(memberId),
		uintptr(unsafe.Pointer(&mops)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}
	return mops.String(), nil
}

// [ReleaseFuncDesc] method.
//
// [ReleaseFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *ITypeInfo) _ReleaseFuncDesc(pFuncDesc *FUNCDESC) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeInfoVt](me.Ppvt()).ReleaseFuncDesc,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pFuncDesc)))
	return utl.HresultToError(ret)
}
