//go:build windows

package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [ITypeInfo] COM interface.
//
// [ITypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo interface {
	com.IUnknown

	// [AddressOfMember] COM method.
	//
	// [AddressOfMember]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-addressofmember
	AddressOfMember(memberId MEMBERID, invokeKind automco.INVOKEKIND) uintptr

	// [CreateInstance] COM method.
	//
	// ⚠️ You must defer IUnknown.Release() on the returned COM object. If
	// iUnkOuter is not null, you must defer IUnknown.Release() on it too.
	//
	// [CreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-createinstance
	CreateInstance(iUnkOuter *com.IUnknown, riid co.IID) com.IUnknown

	// [GetDocumentation] COM method.
	//
	// # Example
	//
	//	var info autom.ITypeInfo // initialized somewhere
	//	var funDesc *autom.FUNCDESC
	//
	//	docum := info.GetDocumentation(funDesc.Memid)
	//	fmt.Printf("Function name: %s\n", docum.Name)
	//
	// [GetDocumentation]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdocumentation
	GetDocumentation(memId MEMBERID) TypeDoc

	// [GetFuncDesc] COM method.
	//
	// ⚠️ You must defer ITypeInfo.ReleaseFuncDesc() on the returned object.
	//
	// # Example
	//
	//	var info autom.ITypeInfo // initialized somewhere
	//	var attr *autom.TYPEATTR
	//
	//	for i := 0; i < int(attr.CFuncs); i++ {
	//		funDesc := info.GetFuncDesc(i)
	//		defer info.ReleaseFuncDesc(funDesc)
	//
	//		fmt.Printf("Member ID: %d, invoke kind: %d\n",
	//			funDesc.Memid, funDesc.Invkind)
	//	}
	//
	// [GetFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getfuncdesc
	GetFuncDesc(index int) *FUNCDESC

	// [GetIDsOfNames] COM method.
	//
	// [GetIDsOfNames]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getidsofnames
	GetIDsOfNames(names []string) []MEMBERID

	// [GetTypeAttr] COM method.
	//
	// ⚠️ You must defer ITypeInfo.ReleaseTypeAttr() on the returned object.
	//
	// # Example
	//
	//	var info autom.ITypeInfo // initialized somewhere
	//
	//	attr := tyInfo.GetTypeAttr()
	//	defer info.ReleaseTypeAttr(attr)
	//
	//	fmt.Printf("Num funcs: %d, GUID: %s\n",
	//		attr.CFuncs, attr.Guid.String())
	//
	// [GetTypeAttr]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-gettypeattr
	GetTypeAttr() *TYPEATTR

	// [GetVarDesc] COM method.
	//
	// ⚠️ You must defer ITypeInfo.ReleaseVarDesc() on the returned object.
	//
	// [GetVarDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getvardesc
	GetVarDesc(index int) *VARDESC

	// This helper method retrieves a resumed information of all functions of
	// this COM interface, by calling ITypeInfo.GetTypeAttr(),
	// ITypeInfo.GetFuncDesc() and ITypeInfo.GetDocumentation().
	ListFunctions() []FuncDescResume

	// [ReleaseFuncDesc] COM method.
	//
	// [ReleaseFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
	ReleaseFuncDesc(funcDesc *FUNCDESC)

	// [ReleaseTypeAttr] COM method.
	//
	// [ReleaseTypeAttr]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasetypeattr
	ReleaseTypeAttr(typeAttr *TYPEATTR)

	// [ReleaseVarDesc] COM method.
	//
	// [ReleaseVarDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasevardesc
	ReleaseVarDesc(varDesc *VARDESC)
}

type _ITypeInfo struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITypeInfo.Release().
func NewITypeInfo(base com.IUnknown) ITypeInfo {
	return &_ITypeInfo{IUnknown: base}
}

func (me *_ITypeInfo) AddressOfMember(
	memberId MEMBERID, invokeKind automco.INVOKEKIND) uintptr {

	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).AddressOfMember,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(memberId), uintptr(invokeKind),
		uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pv
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) CreateInstance(
	iUnkOuter *com.IUnknown, riid co.IID) com.IUnknown {

	var ppvQueried **comvt.IUnknown

	var pppvOuter ***comvt.IUnknown
	if iUnkOuter != nil { // was the outer pointer requested?
		(*iUnkOuter).Release() // release if existing
		var ppvOuterBuf **comvt.IUnknown
		pppvOuter = &ppvOuterBuf // we'll request the outer pointer
	}

	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).CreateInstance,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pppvOuter)),
		uintptr(unsafe.Pointer(win.GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		if iUnkOuter != nil {
			*iUnkOuter = com.NewIUnknown(*pppvOuter)
		}
		return com.NewIUnknown(ppvQueried)
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) GetDocumentation(memberId MEMBERID) TypeDoc {
	var name, docString, helpContext, helpFile uintptr
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetDocumentation,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(memberId),
		uintptr(unsafe.Pointer(&name)), uintptr(unsafe.Pointer(&docString)),
		uintptr(unsafe.Pointer(&helpContext)), uintptr(unsafe.Pointer(&helpFile)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		var ret TypeDoc
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
			// Documented as *uint32, but apparently returned as the value itself.
			ret.HelpContext = uint32(helpContext)
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

func (me *_ITypeInfo) GetFuncDesc(index int) *FUNCDESC {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetFuncDesc,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*FUNCDESC)(unsafe.Pointer(pv))
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) GetIDsOfNames(names []string) []MEMBERID {
	pNames := make([]*uint16, 0, len(names))
	for _, name := range names {
		pNames = append(pNames, win.Str.ToNativePtr(name))
	}

	memberIds := make([]MEMBERID, len(names))

	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetTypeAttr,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pNames[0])),
		uintptr(uint32(len(names))),
		uintptr(unsafe.Pointer(&memberIds[0])))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return memberIds
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) GetTypeAttr() *TYPEATTR {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetTypeAttr,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*TYPEATTR)(unsafe.Pointer(pv))
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) GetVarDesc(index int) *VARDESC {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetVarDesc,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return (*VARDESC)(unsafe.Pointer(pv))
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) ListFunctions() []FuncDescResume {
	attr := me.GetTypeAttr()
	defer me.ReleaseTypeAttr(attr)

	resumes := make([]FuncDescResume, 0, attr.CFuncs)

	for i := 0; i < int(attr.CFuncs); i++ {
		funDesc := me.GetFuncDesc(i)
		defer me.ReleaseFuncDesc(funDesc) // will pile up at the end of the function, but it's fine

		docum := me.GetDocumentation(funDesc.Memid)

		resumes = append(resumes, FuncDescResume{
			MemberId:     funDesc.Memid,
			Name:         docum.Name,
			FuncKind:     funDesc.Funckind,
			InvokeKind:   funDesc.Invkind,
			NumParams:    int(funDesc.CParams),
			NumOptParams: int(funDesc.CParamsOpt),
			Flags:        funDesc.WFuncFlags,
		})
	}

	return resumes
}

func (me *_ITypeInfo) ReleaseFuncDesc(funcDesc *FUNCDESC) {
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).ReleaseFuncDesc,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(funcDesc)))

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

func (me *_ITypeInfo) ReleaseTypeAttr(typeAttr *TYPEATTR) {
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).ReleaseTypeAttr,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(typeAttr)))

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

func (me *_ITypeInfo) ReleaseVarDesc(varDesc *VARDESC) {
	ret, _, _ := syscall.SyscallN(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).ReleaseVarDesc,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(varDesc)))

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}
