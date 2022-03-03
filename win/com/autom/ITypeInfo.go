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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypeinfo
type ITypeInfo interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-addressofmember
	AddressOfMember(memberId MEMBERID, invokeKind automco.INVOKEKIND) uintptr

	// ⚠️ You must defer IUnknown.Release() on the returned COM object. If
	// iUnkOuter is not null, you must defer IUnknown.Release() on it too.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-createinstance
	CreateInstance(iUnkOuter *com.IUnknown, riid co.IID) com.IUnknown

	// Example:
	//
	//  var info autom.ITypeInfo // initialized somewhere
	//  var funDesc *autom.FUNCDESC
	//
	//  docum := info.GetDocumentation(funDesc.Memid)
	//  fmt.Printf("Function name: %s\n", docum.Name)
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getdocumentation
	GetDocumentation(memId MEMBERID) TypeDoc

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
	GetFuncDesc(index int) *FUNCDESC

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getidsofnames
	GetIDsOfNames(names []string) []MEMBERID

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
	GetTypeAttr() *TYPEATTR

	// ⚠️ You must defer ITypeInfo.ReleaseVarDesc() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-getvardesc
	GetVarDesc(index int) *VARDESC

	// Retrieves a resumed information of all functions of this COM interface,
	// by calling ITypeInfo.GetTypeAttr(), ITypeInfo.GetFuncDesc() and
	// ITypeInfo.GetDocumentation().
	ListFunctions() []FuncDescResume

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
	ReleaseFuncDesc(funcDesc *FUNCDESC)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasetypeattr
	ReleaseTypeAttr(typeAttr *TYPEATTR)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasevardesc
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
	ret, _, _ := syscall.Syscall6(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).AddressOfMember, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(memberId), uintptr(invokeKind),
		uintptr(unsafe.Pointer(&pv)),
		0, 0)

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

	ret, _, _ := syscall.Syscall6(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).CreateInstance, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pppvOuter)),
		uintptr(unsafe.Pointer(win.GuidFromIid(riid))),
		uintptr(unsafe.Pointer(&ppvQueried)),
		0, 0)

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
	ret, _, _ := syscall.Syscall6(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetDocumentation, 6,
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
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetFuncDesc, 3,
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

	ret, _, _ := syscall.Syscall6(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetTypeAttr, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pNames[0])),
		uintptr(len(names)),
		uintptr(unsafe.Pointer(&memberIds[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return memberIds
	} else {
		panic(hr)
	}
}

func (me *_ITypeInfo) GetTypeAttr() *TYPEATTR {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetTypeAttr, 3,
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
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).GetVarDesc, 3,
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
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).ReleaseFuncDesc, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(funcDesc)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

func (me *_ITypeInfo) ReleaseTypeAttr(typeAttr *TYPEATTR) {
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).ReleaseTypeAttr, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(typeAttr)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

func (me *_ITypeInfo) ReleaseVarDesc(varDesc *VARDESC) {
	ret, _, _ := syscall.Syscall(
		(*automvt.ITypeInfo)(unsafe.Pointer(*me.Ptr())).ReleaseVarDesc, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(varDesc)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}
