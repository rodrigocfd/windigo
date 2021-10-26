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

type _TypeDoc struct {
	Name        string
	DocString   string
	HelpContext uint32
	HelpFile    string
}

// Example:
//
//  var info autom.ITypeInfo // initialized somewhere
//  var funDesc *autom.FUNCDESC
//
//  docum := info.GetDocumentation(funDesc.Memid)
//  fmt.Printf("Method name: %s\n", docum.Name)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdocumentation
func (me *ITypeInfo) GetDocumentation(memId MEMBERID) _TypeDoc {
	var name, docString, helpContext, helpFile uintptr
	ret, _, _ := syscall.Syscall6(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).GetDocumentation, 6,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(memId),
		uintptr(unsafe.Pointer(&name)), uintptr(unsafe.Pointer(&docString)),
		uintptr(unsafe.Pointer(&helpContext)), uintptr(unsafe.Pointer(&helpFile)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		var ret _TypeDoc
		if name != 0 {
			bstr := BSTR(name)
			defer bstr.SysFreeString()
			ret.Name = bstr.String()
		}
		if docString != 0 {
			bstr := BSTR(docString)
			defer bstr.SysFreeString()
			ret.DocString = bstr.String()
		}
		if helpContext != 0 {
			ret.HelpContext = *(*uint32)(unsafe.Pointer(helpContext))
		}
		if helpFile != 0 {
			bstr := BSTR(helpFile)
			defer bstr.SysFreeString()
			ret.HelpFile = bstr.String()
		}
		return ret
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer ITypeInfo.ReleaseFuncDesc() on the returned object.
//
// Example:
//
//  var info autom.ITypeInfo // initialized somewhere
//  var attr *autom.TYPEATTR
//
//  for i := 0; i < int(attr.CFuncs); i++ {
//      funDesc := info.GetFuncDesc(i)
//      defer info.ReleaseFuncDesc(funDesc)
//
//      fmt.Printf("Member ID: %d, invoke kind: %d\n",
//          funDesc.Memid, funDesc.Invkind)
//  }
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

// ⚠️ You must defer ITypeInfo.ReleaseTypeAttr() on the returned object.
//
// Example:
//
//  var info autom.ITypeInfo // initialized somewhere
//
//  attr := tyInfo.GetTypeAttr()
//  defer info.ReleaseTypeAttr(attr)
//
//  fmt.Printf("Num funcs: %d, GUID: %s\n",
//      attr.CFuncs, attr.Guid.String())
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-gettypeattr
func (me *ITypeInfo) GetTypeAttr() *TYPEATTR {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).GetTypeAttr, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*TYPEATTR)(unsafe.Pointer(pv))
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

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasetypeattr
func (me *ITypeInfo) ReleaseTypeAttr(typeAttr *TYPEATTR) {
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).ReleaseTypeAttr, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(typeAttr)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasevardesc
func (me *ITypeInfo) ReleaseVarDesc(varDesc *VARDESC) {
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfoVtbl)(unsafe.Pointer(*me.Ptr())).ReleaseVarDesc, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(varDesc)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}
