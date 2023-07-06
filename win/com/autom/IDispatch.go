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

// [IDispatch] COM interface.
//
// [IDispatch]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch interface {
	com.IUnknown

	// [GetIDsOfNames] COM method.
	//
	// [GetIDsOfNames]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-getidsofnames
	GetIDsOfNames(lcid win.LCID,
		member string, parameters ...string) ([]MEMBERID, error)

	// [GetTypeInfo] COM method.
	//
	// ⚠️ You must defer ITypeInfo.Release() on the returned object.
	//
	// # Example
	//
	//	var iDisp autom.IDispatch // initialized somewhere
	//
	//	tyInfo := iDisp.GetTypeInfo(win.LCID_SYSTEM_DEFAULT)
	//	defer tyInfo.Release()
	//
	// [GetTypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
	GetTypeInfo(lcid win.LCID) ITypeInfo

	// [GetTypeInfoCount] COM method.
	//
	// If the object provides type information, this number is 1; otherwise the
	// number is 0.
	//
	// [GetTypeInfoCount]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
	GetTypeInfoCount() int

	// [Invoke] COM method.
	//
	// This is a low-level method, prefer using IDispatch.InvokePut(),
	// IDispatch.InvokeMethod() or IDispatch.InvokeGet().
	//
	// If the Invoke() call itself fails, an ordinary HRESULT is returned in the
	// form of an errco.ERROR.
	//
	// If the remote call fails, an *autom.ExceptionInfo is returned.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the returned VARIANT.
	//
	// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
	Invoke(dispIdMember MEMBERID, lcid win.LCID,
		flags automco.DISPATCH, dispParams *DISPPARAMS) (VARIANT, error)

	// This helper method calls IDispatch.GetIDsOfNames(), builds the DISPPARAMS
	// array and calls IDispatch.Invoke() with automco.DISPATCH_PROPERTYGET.
	//
	// If the Invoke() call itself fails, an ordinary HRESULT is returned in the
	// form of an errco.ERROR.
	//
	// If the remote call fails, an *autom.ExceptionInfo is returned.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the returned VARIANT.
	//
	// # Example
	//
	//	xlApp, err := autom.NewIDispatchFromProgId("Excel.Application")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer xlApp.Release()
	//
	//	variRet, err := xlApp.InvokeGet("Workbooks")
	//	if err != nil {
	//		switch realErr := err.(type) {
	//		case *autom.ExceptionInfo:
	//			println("Invoke error", realErr.Code, realErr.Description)
	//		default:
	//			println("Ordinary error", realErr.Error())
	//		}
	//	}
	//	defer variRet.VariantClear()
	InvokeGet(methodName string, params ...VARIANT) (VARIANT, error)

	// This helper method calls IDispatch.GetIDsOfNames(), builds the DISPPARAMS
	// array and calls IDispatch.Invoke() with automco.DISPATCH_METHOD.
	//
	// If the Invoke() call itself fails, an ordinary HRESULT is returned in the
	// form of an errco.ERROR.
	//
	// If the remote call fails, an *autom.ExceptionInfo is returned.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the returned VARIANT.
	//
	// # Example
	//
	//	xlApp, err := autom.NewIDispatchFromProgId("Excel.Application")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer xlApp.Release()
	//
	//	ret, err := xlApp.InvokeMethod("Quit")
	//	if err != nil {
	//		switch realErr := err.(type) {
	//		case *autom.ExceptionInfo:
	//			println("Invoke error", realErr.Code, realErr.Description)
	//		default:
	//			println("Ordinary error", realErr.Error())
	//		}
	//	}
	//	defer variRet.VariantClear()
	InvokeMethod(methodName string, params ...VARIANT) (VARIANT, error)

	// This helper method calls IDispatch.GetIDsOfNames(), builds the DISPPARAMS
	// array and calls IDispatch.Invoke() with automco.DISPATCH_PROPERTYPUT.
	//
	// If the Invoke() call itself fails, an ordinary HRESULT is returned in the
	// form of an errco.ERROR.
	//
	// If the remote call fails, an *autom.ExceptionInfo is returned.
	//
	// ⚠️ You must defer VARIANT.VariantClear() on the returned VARIANT.
	//
	// # Example
	//
	//	xlApp, err := autom.NewIDispatchFromProgId("Excel.Application")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	trueVal := autom.NewVariantInt32(1)
	//	defer trueVal.VariantClear()
	//
	//	ret, err := xlApp.InvokePut("Visible", trueVal)
	//	if err != nil {
	//		switch realErr := err.(type) {
	//		case *autom.ExceptionInfo:
	//			println("Invoke error", realErr.Code, realErr.Description)
	//		default:
	//			println("Ordinary error", realErr.Error())
	//		}
	//	}
	//	defer variRet.VariantClear()
	InvokePut(methodName string, params ...VARIANT) (VARIANT, error)

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

// Helper function which constructs an automation IDispatch object by calling
// [CLSIDFromProgID] and [CoCreateInstance].
//
// Note that the target application must be installed, otherwise the call will
// fail.
//
// ⚠️ You must defer IDispatch.Release().
//
// # Example
//
//	excelApp, _ := autom.NewIDispatchFromProgId("Excel.Application")
//	defer excelApp.Release()
//
// [CLSIDFromProgID]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func NewIDispatchFromProgId(progId string) (IDispatch, error) {
	clsId, err := com.CLSIDFromProgID(progId)
	if err != nil {
		return nil, err
	}

	return NewIDispatch(
		com.CoCreateInstance(clsId, nil,
			comco.CLSCTX_LOCAL_SERVER, automco.IID_IDispatch),
	), nil
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

	ret, _, _ := syscall.SyscallN(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetIDsOfNames,
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
	ret, _, _ := syscall.SyscallN(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfo,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, uintptr(lcid),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewITypeInfo(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IDispatch) GetTypeInfoCount() int {
	var pctInfo uint32
	ret, _, _ := syscall.SyscallN(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).GetTypeInfoCount,
		uintptr(unsafe.Pointer(me.Ptr())), uintptr(unsafe.Pointer(&pctInfo)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(pctInfo)
	} else {
		panic(hr)
	}
}

func (me *_IDispatch) Invoke(
	dispIdMember MEMBERID, lcid win.LCID,
	flags automco.DISPATCH, dispParams *DISPPARAMS) (VARIANT, error) {

	var retExcep EXCEPINFO
	var retVari VARIANT

	ret, _, _ := syscall.SyscallN(
		(*automvt.IDispatch)(unsafe.Pointer(*me.Ptr())).Invoke,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(dispIdMember),
		uintptr(unsafe.Pointer(win.GuidFromIid(comco.IID_NULL))),
		uintptr(lcid), uintptr(flags),
		uintptr(unsafe.Pointer(dispParams)),
		uintptr(unsafe.Pointer(&retVari)),
		uintptr(unsafe.Pointer(&retExcep)),
		0) // puArgErr is not retrieved

	if hr := errco.ERROR(ret); hr == errco.S_OK { // Invoke() call succeeded
		if retExcep.BstrSource == 0 && retExcep.BstrDescription == 0 && retExcep.BstrHelpFile == 0 { // all good
			return retVari, nil
		} else { // Invoke() call succeeded, remote call returned error
			defer retVari.VariantClear() // if any; not returned to user
			e0, e1, e2 := retExcep.ReleaseStrings()
			userExcep := &ExceptionInfo{
				Code:        int32(retExcep.WCode),
				Source:      e0,
				Description: e1,
				HelpFile:    e2,
			}
			if userExcep.Code == 0 {
				userExcep.Code = retExcep.Scode
			}
			return VARIANT{}, userExcep
		}
	} else { // Invoke() call failed
		return VARIANT{}, errco.ERROR(ret)
	}
}

func (me *_IDispatch) invokeCall(
	methodName string,
	flags automco.DISPATCH, params ...VARIANT) (VARIANT, error) {

	// https://learn.microsoft.com/en-us/previous-versions/office/troubleshoot/office-developer/automate-excel-from-c

	memIds, err := me.GetIDsOfNames(win.LCID_USER_DEFAULT, methodName)
	if err != nil {
		return VARIANT{}, err
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
	methodName string, params ...VARIANT) (VARIANT, error) {

	return me.invokeCall(methodName, automco.DISPATCH_PROPERTYGET, params...)
}

func (me *_IDispatch) InvokeMethod(
	methodName string, params ...VARIANT) (VARIANT, error) {

	return me.invokeCall(methodName, automco.DISPATCH_METHOD, params...)
}

func (me *_IDispatch) InvokePut(
	methodName string, params ...VARIANT) (VARIANT, error) {

	return me.invokeCall(methodName, automco.DISPATCH_PROPERTYPUT, params...)
}

func (me *_IDispatch) ListFunctions() []FuncDescResume {
	info := me.GetTypeInfo(win.LCID_SYSTEM_DEFAULT)
	defer info.Release()

	return info.ListFunctions()
}
