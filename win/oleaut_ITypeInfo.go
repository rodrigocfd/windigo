//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [ITypeInfo] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo struct{ IUnknown }

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
		uintptr(memberId),
		uintptr(invokeKind),
		uintptr(unsafe.Pointer(&addr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return addr, nil
	} else {
		return 0, hr
	}
}

// [CreateInstance] method.
//
// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-createinstance
func (me *ITypeInfo) CreateInstance(
	releaser *OleReleaser,
	unkOuter *IUnknown,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).CreateInstance,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(ppvtOrNil(unkOuter)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.OleCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(OleObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

// [GetContainingTypeLib] method.
//
// Returns the type library and its index.
//
// [GetContainingTypeLib]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getcontainingtypelib
func (me *ITypeInfo) GetContainingTypeLib(releaser *OleReleaser) (*ITypeLib, uint, error) {
	var ppvtQueried **_IUnknownVt
	var index uint32

	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetContainingTypeLib,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&index)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &ITypeLib{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, uint(index), nil
	} else {
		return nil, 0, hr
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
		uintptr(memberId),
		uintptr(invokeKind),
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

// [GetFuncDesc] method.
//
// The [OleReleaser] is responsible for freeing the resources by calling
// [ReleaseFuncDesc].
//
// Example:
//
//	var nfo *win.ITypeInfo // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	funcDesc, _ := nfo.GetFuncDesc(rel, 0)
//	println(funcDesc.Memid)
//
// [GetFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getfuncdesc
// [ReleaseFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *ITypeInfo) GetFuncDesc(releaser *OleReleaser, index uint) (*FuncDescData, error) {
	var pFuncDesc *FUNCDESC
	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetFuncDesc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&pFuncDesc)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pData := &FuncDescData{pFuncDesc, me}
		releaser.Add(pData)
		return pData, nil
	} else {
		return nil, hr
	}
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
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetIDsOfNames,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&strPtrs[0])),
		uintptr(uint32(len(names))),
		uintptr(unsafe.Pointer(&memIds[0])))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return memIds, nil
	} else {
		return nil, hr
	}
}

// [GetImplTypeFlags] method.
//
// [GetImplTypeFlags]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getimpltypeflags
func (me *ITypeInfo) GetImplTypeFlags(index uint) (co.IMPLTYPEFLAG, error) {
	var flags co.IMPLTYPEFLAG
	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetImplTypeFlags,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&flags)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return flags, nil
	} else {
		return co.IMPLTYPEFLAG(0), hr
	}
}

// [GetMops] method.
//
// [GetMops]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getmops
func (me *ITypeInfo) GetMops(memberId MEMBERID) (string, error) {
	var mops BSTR
	defer mops.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).GetMops,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(memberId),
		uintptr(unsafe.Pointer(&mops)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return mops.String(), nil
	} else {
		return "", hr
	}
}

// [ReleaseFuncDesc] method.
//
// [ReleaseFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *ITypeInfo) _ReleaseFuncDesc(pFuncDesc *FUNCDESC) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITypeInfoVt)(unsafe.Pointer(*me.Ppvt())).ReleaseFuncDesc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pFuncDesc)))
	return utl.ErrorAsHResult(ret)
}

type _ITypeInfoVt struct {
	_IUnknownVt
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
