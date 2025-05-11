//go:build windows

package oleaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IDispatch] COM interface.
//
// [IDispatch]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDispatch) IID() co.IID {
	return co.IID_IDispatch
}

// [GetIDsOfNames] method.
//
// [GetIDsOfNames]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-getidsofnames
func (me *IDispatch) GetIDsOfNames(
	lcid win.LCID,
	member string,
	parameters ...string,
) ([]MEMBERID, error) {
	nParams := uint(1 + len(parameters)) // member + parameters
	nullGuid := win.GuidFrom(co.IID_NULL)
	memberIds := make([]MEMBERID, nParams) // will be returned

	allStrs16 := wstr.NewArray()
	allStrs16.Append(member)
	allStrs16.Append(parameters...)

	strPtrs := make([]*uint16, 0, nParams)
	for i := uint(0); i < nParams; i++ {
		strPtrs = append(strPtrs, allStrs16.PtrOf(i))
	}

	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).GetIDsOfNames,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&nullGuid)),
		uintptr(unsafe.Pointer(&strPtrs[0])),
		uintptr(nParams),
		uintptr(lcid),
		uintptr(unsafe.Pointer(&memberIds[0])))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return memberIds, nil
	} else {
		return nil, hr
	}
}

// [GetTypeInfo] method.
//
// # Example
//
//	var iDisp oleaut.IDispatch // initialized somewhere
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	nfo, _ := iDisp.GetTypeInfo(rel, win.LCID_USER_DEFAULT)
//
// [GetTypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (me *IDispatch) GetTypeInfo(releaser *ole.Releaser, lcid win.LCID) (*ITypeInfo, error) {
	var ppvtQueried **ole.IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).GetTypeInfo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0, uintptr(lcid),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := ole.ComObj[ITypeInfo](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetTypeInfoCount] method.
//
// If the object provides type information, this number is 1; otherwise the
// number is 0.
//
// [GetTypeInfoCount]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfocount
func (me *IDispatch) GetTypeInfoCount() (uint, error) {
	var pctInfo uint32
	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).GetTypeInfoCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&pctInfo)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return uint(pctInfo), nil
	} else {
		return 0, hr
	}
}

// [Invoke] method.
//
// This is a low-level method, prefer using IDispatch.InvokeGet(),
// IDispatch.InvokeMethod() or IDispatch.InvokePut().
//
// If the remote call raises an exception, the returned error will be an
// instance of *ole.[EXCEPINFO].
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) Invoke(
	releaser *ole.Releaser,
	dispIdMember MEMBERID,
	lcid win.LCID,
	flags co.DISPATCH,
	dispParams *DISPPARAMS,
) (*VARIANT, error) {
	var remoteErr _EXCEPINFO // in case of remote error, will be converted to *EXCEPINFO
	defer remoteErr.Free()

	remoteResult := NewVariantEmpty(releaser) // result returned from the remote call
	nullGuid := win.GuidFrom(co.IID_NULL)

	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).Invoke,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(dispIdMember),
		uintptr(unsafe.Pointer(&nullGuid)),
		uintptr(lcid), uintptr(flags),
		uintptr(unsafe.Pointer(dispParams)),
		uintptr(unsafe.Pointer(remoteResult)),
		uintptr(unsafe.Pointer(&remoteErr)),
		0) // puArgErr is not retrieved

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return remoteResult, nil
	} else if hr == co.HRESULT_DISP_E_EXCEPTION {
		return nil, remoteErr.Serialize()
	} else {
		return nil, hr
	}
}

// Calls [Invoke] with co.DISPATCH_PROPERTYGET.
//
// If the remote call raises an exception, the returned error will be an
// instance of *ole.[EXCEPINFO].
//
// # Example
//
//	ole.CoInitializeEx(co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer ole.CoUninitialize()
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	clsId, _ := ole.CLSIDFromProgID("Excel.Application")
//	iExcel, _ := ole.CoCreateInstance[oleaut.IDispatch](
//		rel, clsId, co.CLSCTX_LOCAL_SERVER)
//
//	vBooks, _ := iExcel.InvokeGet(rel, "Workbooks")
//	iBooks, _ := vBooks.IDispatch(rel)
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokeGet(
	releaser *ole.Releaser,
	propertyName string,
	params ...interface{},
) (*VARIANT, error) {
	return me.rawInvoke(releaser, co.DISPATCH_PROPERTYGET, propertyName, params...)
}

// Calls [Invoke] with co.DISPATCH_METHOD.
//
// If the remote call raises an exception, the returned error will be an
// instance of *ole.[EXCEPINFO].
//
// Parameters must be one of the valid VARIANT types.
//
// # Example
//
//	ole.CoInitializeEx(co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer ole.CoUninitialize()
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	clsId, _ := ole.CLSIDFromProgID("Excel.Application")
//	iExcel, _ := ole.CoCreateInstance[oleaut.IDispatch](
//		rel, clsId, co.CLSCTX_LOCAL_SERVER)
//
//	vBooks, _ := iExcel.InvokeGet(rel, "Workbooks")
//	iBooks, _ := vBooks.IDispatch(rel)
//	vFile, _ := iBooks.InvokeMethod(rel, "Open", "C:\\Temp\\file.xlsx")
//	iFile, _ := vFile.IDispatch(rel)
//	iFile.InvokeMethod(rel, "SaveAs", "C:\\Temp\\copy.xlsx")
//	iFile.InvokeMethod(rel, "Close")
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokeMethod(
	releaser *ole.Releaser,
	methodName string,
	params ...interface{},
) (*VARIANT, error) {
	return me.rawInvoke(releaser, co.DISPATCH_METHOD, methodName, params...)
}

// Calls [Invoke] with co.DISPATCH_PROPERTYPUT.
//
// If the remote call raises an exception, the returned error will be an
// instance of *oleaut.[EXCEPINFO].
//
// Parameter must be one of the valid VARIANT types.
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokePut(
	releaser *ole.Releaser,
	propertyName string,
	value interface{},
) (*VARIANT, error) {
	return me.rawInvoke(releaser, co.DISPATCH_PROPERTYPUT, propertyName, value)
}

func (me *IDispatch) rawInvoke(
	releaser *ole.Releaser,
	method co.DISPATCH,
	methodName string,
	params ...interface{},
) (*VARIANT, error) {
	memberIds, err := me.GetIDsOfNames(win.LCID_USER_DEFAULT, methodName) // will return 1 element
	if err != nil {
		return nil, err
	}

	localRel := ole.NewReleaser()
	defer localRel.Release()

	arrVars := make([]VARIANT, 0, len(params))
	for i := len(params) - 1; i >= 0; i-- { // in reverse order
		arrVars = append(arrVars, *NewVariant(localRel, params[i])) // copy bytes, and trust they won't be changed
	}

	var dp DISPPARAMS
	if len(params) > 0 {
		dp.SetArgs(arrVars)
	}
	if method == co.DISPATCH_PROPERTYPUT {
		dp.SetNamedArgs(co.DISPID_PROPERTYPUT)
	}

	v, err := me.Invoke(releaser, memberIds[0], win.LCID_USER_DEFAULT, method, &dp)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type _IDispatchVt struct {
	ole.IUnknownVt
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}
