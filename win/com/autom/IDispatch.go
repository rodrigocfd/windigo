//go:build windows

package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-getidsofnames
	GetIDsOfNames(lcid win.LCID,
		member string, parameters ...string) ([]MEMBERID, error)

	// ⚠️ You must defer ITypeInfo.Release() on the returned object.
	//
	// Example:
	//
	//		var iDisp autom.IDispatch // initialized somewhere
	//
	//		tyInfo := iDisp.GetTypeInfo(win.LCID_SYSTEM_DEFAULT)
	//		defer tyInfo.Release()
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
	GetTypeInfo(lcid win.LCID) ITypeInfo

	// If the object provides type information, this number is 1; otherwise the
	// number is 0.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
	GetTypeInfoCount() int

	// This is a low-level method, prefer using IDispatch.InvokePut(),
	// IDispatch.InvokeMethod() or IDispatch.InvokeGet().
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the VarResult member of the
	// returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
	Invoke(dispIdMember MEMBERID, lcid win.LCID,
		flags automco.DISPATCH, dispParams *DISPPARAMS) (_InvokeRet, error)

	// This helper method calls IDispatch.GetIDsOfNames(), builds the DISPPARAMS
	// array and calls IDispatch.Invoke() with automco.DISPATCH_PROPERTYGET.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the VarResult member of the
	// returned object.
	InvokeGet(methodName string, params ...VARIANT) (_InvokeRet, error)

	// This helper method calls IDispatch.GetIDsOfNames(), builds the DISPPARAMS
	// array and calls IDispatch.Invoke() with automco.DISPATCH_METHOD.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the VarResult member of the
	// returned object.
	InvokeMethod(methodName string, params ...VARIANT) (_InvokeRet, error)

	// This helper method calls IDispatch.GetIDsOfNames(), builds the DISPPARAMS
	// array and calls IDispatch.Invoke() with automco.DISPATCH_PROPERTYPUT.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the VarResult member of the
	// returned object.
	InvokePut(methodName string, params ...VARIANT) (_InvokeRet, error)

	// This helper method calls IDispatch.GetTypeInfo() with
	// win.LCID_SYSTEM_DEFAULT, then calls ITypeInfo.ListFunctions().
	ListFunctions() []FuncDescResume
}

type _IDispatch struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IDispatch.Release().
func NewIDispatch(base com.IUnknown) IDispatch {
	return &_IDispatch{IUnknown: base}
}

// Constructs an automation IDispatch object by calling CLSIDFromProgID().
//
// Panics if progId is invalid.
//
// ⚠️ You must defer IDispatch.Release().
//
// Example:
//
//		excelApp := autom.NewIDispatchFromProgId("Excel.Application")
//		defer excelApp.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func NewIDispatchFromProgId(progId string) IDispatch {
	clsId, err := com.CLSIDFromProgID("Excel.Application")
	if err != nil {
		panic(err)
	}

	return NewIDispatch(
		com.CoCreateInstance(clsId, nil,
			comco.CLSCTX_LOCAL_SERVER, automco.IID_IDispatch),
	)
}

func (me *_IDispatch) GetIDsOfNames(
	lcid win.LCID, member string, parameters ...string) ([]MEMBERID, error) {

	numStrs := 1 + len(parameters)
	memberIds := make([]MEMBERID, numStrs)

	oleStrs := make([]*uint16, 0, numStrs)
	oleStrs = append(oleStrs, win.Str.ToNativePtr(member))
	for _, parameter := range parameters {
		oleStrs = append(oleStrs, win.Str.ToNativePtr(parameter))
	}

	ret, _, _ := syscall.Syscall6(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetIDsOfNames, 6,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.GuidFromIid(comco.IID_NULL))),
		uintptr(unsafe.Pointer(&oleStrs[0])), uintptr(numStrs),
		uintptr(lcid), uintptr(unsafe.Pointer(&memberIds[0])))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return memberIds, nil
	} else if hr == errco.DISP_E_UNKNOWNNAME || hr == errco.DISP_E_UNKNOWNLCID {
		return nil, hr
	} else {
		panic(hr)
	}
}

func (me *_IDispatch) GetTypeInfo(lcid win.LCID) ITypeInfo {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfo, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, uintptr(lcid),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewITypeInfo(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IDispatch) GetTypeInfoCount() int {
	var pctInfo uint32
	ret, _, _ := syscall.Syscall(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pctInfo)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(pctInfo)
	} else {
		panic(hr)
	}
}

type _InvokeRet struct {
	VarResult VARIANT
	ExcepInfo EXCEPINFO
	ArgErr    uint32
}

func (me *_IDispatch) Invoke(
	dispIdMember MEMBERID, lcid win.LCID,
	flags automco.DISPATCH, dispParams *DISPPARAMS) (_InvokeRet, error) {

	var invokeRet _InvokeRet

	ret, _, _ := syscall.Syscall9(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).Invoke, 9,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(dispIdMember),
		uintptr(unsafe.Pointer(win.GuidFromIid(comco.IID_NULL))),
		uintptr(lcid), uintptr(flags),
		uintptr(unsafe.Pointer(dispParams)),
		uintptr(unsafe.Pointer(&invokeRet.VarResult)),
		uintptr(unsafe.Pointer(&invokeRet.ExcepInfo)),
		uintptr(unsafe.Pointer(&invokeRet.ArgErr)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return invokeRet, nil
	} else {
		panic(hr)
	}
}

func (me *_IDispatch) invokeCall(
	methodName string,
	flags automco.DISPATCH, params ...VARIANT) (_InvokeRet, error) {

	// https://docs.microsoft.com/en-us/previous-versions/office/troubleshoot/office-developer/automate-excel-from-c

	memIds, err := me.GetIDsOfNames(win.LCID_USER_DEFAULT, methodName)
	if err != nil {
		return _InvokeRet{}, err
	}

	var dp DISPPARAMS
	if len(params) > 0 {
		dp.SetArgs(params...)
	}
	if (flags & automco.DISPATCH_PROPERTYPUT) != 0 {
		dp.SetNamedArgs(automco.DISPID_PROPERTYPUT)
	}

	return me.Invoke(memIds[0], win.LCID_SYSTEM_DEFAULT, flags, &dp)
}

func (me *_IDispatch) InvokeGet(
	methodName string, params ...VARIANT) (_InvokeRet, error) {

	return me.invokeCall(methodName, automco.DISPATCH_PROPERTYGET, params...)
}

func (me *_IDispatch) InvokeMethod(
	methodName string, params ...VARIANT) (_InvokeRet, error) {

	return me.invokeCall(methodName, automco.DISPATCH_METHOD, params...)
}

func (me *_IDispatch) InvokePut(
	methodName string, params ...VARIANT) (_InvokeRet, error) {

	return me.invokeCall(methodName, automco.DISPATCH_PROPERTYPUT, params...)
}

func (me *_IDispatch) ListFunctions() []FuncDescResume {
	info := me.GetTypeInfo(win.LCID_SYSTEM_DEFAULT)
	defer info.Release()

	return info.ListFunctions()
}
