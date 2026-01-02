//go:build windows

package win

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IDispatch] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IDispatch]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
type IDispatch struct{ IUnknown }

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
	lcid LCID,
	member string,
	parameters ...string,
) ([]MEMBERID, error) {
	nParams := 1 + len(parameters) // member + parameters
	nullGuid := GuidFrom(co.IID_NULL)
	memberIds := make([]MEMBERID, nParams) // to be returned

	strPtrs := make([]*uint16, 0, nParams)
	strPtrs = append(strPtrs, wstr.EncodeToPtr(member))
	for _, param := range parameters {
		strPtrs = append(strPtrs, wstr.EncodeToPtr(param))
	}

	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).GetIDsOfNames,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&nullGuid)),
		uintptr(unsafe.Pointer(&strPtrs[0])),
		uintptr(uint32(nParams)),
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
// Example:
//
//	var iDisp oleaut.IDispatch // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	nfo, _ := iDisp.GetTypeInfo(rel, win.LCID_USER_DEFAULT)
//
// [GetTypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-gettypeinfo
func (me *IDispatch) GetTypeInfo(releaser *OleReleaser, lcid LCID) (*ITypeInfo, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).GetTypeInfo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0,
		uintptr(lcid),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &ITypeInfo{IUnknown{ppvtQueried}}
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
func (me *IDispatch) GetTypeInfoCount() (int, error) {
	var pctInfo uint32
	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).GetTypeInfoCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&pctInfo)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(pctInfo), nil
	} else {
		return 0, hr
	}
}

// [Invoke] method.
//
// This is a low-level method, prefer using [IDispatch.InvokeGet],
// [IDispatch.InvokeMethod] or [IDispatch.InvokePut].
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
func (me *IDispatch) Invoke(
	releaser *OleReleaser,
	dispIdMember MEMBERID,
	lcid LCID,
	flags co.DISPATCH,
	dispParams *DISPPARAMS,
) (*VARIANT, error) {
	var remoteErr _EXCEPINFO // in case of remote error, will be converted to *EXCEPINFO
	defer remoteErr.Free()

	remoteResult := NewVariantEmpty(releaser) // result returned from the remote call
	nullGuid := GuidFrom(co.IID_NULL)

	ret, _, _ := syscall.SyscallN(
		(*_IDispatchVt)(unsafe.Pointer(*me.Ppvt())).Invoke,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(dispIdMember),
		uintptr(unsafe.Pointer(&nullGuid)),
		uintptr(lcid),
		uintptr(flags),
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

// Calls [Invoke] with [co.DISPATCH_PROPERTYGET].
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// Parameters must be one of the valid [VARIANT] types.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//	var excel *win.IDispatch
//	_ = win.CoCreateInstance(
//		rel,
//		clsId,
//		nil,
//		co.CLSCTX_LOCAL_SERVER,
//		&excel,
//	)
//
//	varBooks, _ := excel.InvokeGet(rel, "Workbooks")
//	dispBooks, _ := varBooks.IDispatch(rel)
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokeGet(
	releaser *OleReleaser,
	propertyName string,
	params ...interface{},
) (*VARIANT, error) {
	return me.rawInvoke(releaser, co.DISPATCH_PROPERTYGET, propertyName, params...)
}

// Calls [Invoke] with [co.DISPATCH_PROPERTYGET], and tries to convert the
// [VARIANT] result to an [IDispatch] object.
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// Parameters must be one of the valid [VARIANT] types.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//	var excel *win.IDispatch
//	_ = win.CoCreateInstance(
//		rel,
//		clsId,
//		nil,
//		co.CLSCTX_LOCAL_SERVER,
//		&excel,
//	)
//
//	books, _ := excel.InvokeGetIDispatch(rel, "Workbooks")
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokeGetIDispatch(
	releaser *OleReleaser,
	propertyName string,
	params ...interface{},
) (*IDispatch, error) {
	variant, err := me.InvokeGet(releaser, propertyName, params...)
	if err != nil {
		return nil, err
	}
	if idisp, ok := variant.IDispatch(releaser); ok {
		return idisp, nil
	} else {
		return nil, fmt.Errorf("InvokeGet \"%s\" didn't return an IDispatch object", propertyName)
	}
}

// Calls [Invoke] with [co.DISPATCH_METHOD].
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// Parameters must be one of the valid [VARIANT] types.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//	var excel *win.IDispatch
//	_ = win.CoCreateInstance(
//		rel,
//		clsId,
//		nil,
//		co.CLSCTX_LOCAL_SERVER,
//		&excel,
//	)
//
//	varBooks, _ := excel.InvokeGet(rel, "Workbooks")
//	dispBooks, _ := varBooks.IDispatch(rel)
//
//	varFile, _ := dispBooks.InvokeMethod(rel, "Open", "C:\\Temp\\file.xlsx")
//	dispFile, _ := varFile.IDispatch(rel)
//
//	_, _ = dispFile.InvokeMethod(rel, "SaveAs", "C:\\Temp\\copy.xlsx")
//	_, _ = dispFile.InvokeMethod(rel, "Close")
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokeMethod(
	releaser *OleReleaser,
	methodName string,
	params ...interface{},
) (*VARIANT, error) {
	return me.rawInvoke(releaser, co.DISPATCH_METHOD, methodName, params...)
}

// Calls [Invoke] with [co.DISPATCH_METHOD], and tries to convert the
// [VARIANT] result to an [IDispatch] object.
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// Parameters must be one of the valid [VARIANT] types.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//	var excel *win.IDispatch
//	_ = win.CoCreateInstance(
//		rel,
//		clsId,
//		nil,
//		co.CLSCTX_LOCAL_SERVER,
//		&excel,
//	)
//
//	books, _ := excel.InvokeGetIDispatch(rel, "Workbooks")
//	file, _ := books.InvokeMethodIDispatch(rel, "Open", "C:\\Temp\\file.xlsx")
//	_, _ = file.InvokeMethod(rel, "SaveAs", "C:\\Temp\\copy.xlsx")
//	_, _ = file.InvokeMethod(rel, "Close")
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokeMethodIDispatch(
	releaser *OleReleaser,
	methodName string,
	params ...interface{},
) (*IDispatch, error) {
	variant, err := me.InvokeMethod(releaser, methodName, params...)
	if err != nil {
		return nil, err
	}
	if idisp, ok := variant.IDispatch(releaser); ok {
		return idisp, nil
	} else {
		return nil, fmt.Errorf("InvokeMethod \"%s\" didn't return an IDispatch object", methodName)
	}
}

// Calls [Invoke] with [co.DISPATCH_PROPERTYPUT].
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// Parameter must be one of the valid [VARIANT] types.
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokePut(
	releaser *OleReleaser,
	propertyName string,
	value interface{},
) (*VARIANT, error) {
	return me.rawInvoke(releaser, co.DISPATCH_PROPERTYPUT, propertyName, value)
}

// Calls [Invoke] with [co.DISPATCH_PROPERTYPUT], and tries to convert the
// [VARIANT] result to an [IDispatch] object.
//
// If the remote call raises an exception, the returned error will be an
// instance of *[EXCEPINFO].
//
// Parameter must be one of the valid [VARIANT] types.
//
// [Invoke]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
func (me *IDispatch) InvokePutIDispatch(
	releaser *OleReleaser,
	propertyName string,
	value interface{},
) (*IDispatch, error) {
	variant, err := me.InvokePut(releaser, propertyName, value)
	if err != nil {
		return nil, err
	}
	if idisp, ok := variant.IDispatch(releaser); ok {
		return idisp, nil
	} else {
		return nil, fmt.Errorf("InvokePut \"%s\" didn't return an IDispatch object", propertyName)
	}
}

func (me *IDispatch) rawInvoke(
	releaser *OleReleaser,
	method co.DISPATCH,
	methodName string,
	params ...interface{},
) (*VARIANT, error) {
	memberIds, err := me.GetIDsOfNames(LCID_USER_DEFAULT, methodName) // will return 1 element
	if err != nil {
		return nil, err
	}

	localRel := NewOleReleaser()
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

	v, err := me.Invoke(releaser, memberIds[0], LCID_USER_DEFAULT, method, &dp)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type _IDispatchVt struct {
	_IUnknownVt
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IPicture] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nn-ocidl-ipicture
type IPicture struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPicture) IID() co.IID {
	return co.IID_ITaskbarList
}

// [get_Attributes] method.
//
// [get_Attributes]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_attributes
func (me *IPicture) Attributes() (co.PICATTR, error) {
	var attr co.PICATTR
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Attributes,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&attr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return attr, nil
	} else {
		return co.PICATTR(0), hr
	}
}

// [get_CurDC] method.
//
// [get_CurDC]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_curdc
func (me *IPicture) CurDC() (HDC, error) {
	var hdc HDC
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_CurDC,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hdc)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hdc, nil
	} else {
		return HDC(0), hr
	}
}

// [get_Handle] method.
//
// [get_Handle]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_handle
func (me *IPicture) Handle() (HBITMAP, error) {
	var hBmp HBITMAP
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Handle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hBmp)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hBmp, nil
	} else {
		return HBITMAP(0), hr
	}
}

// [get_Height] method.
//
// If you need both width and height, call [IPicture.Size], which returns both.
//
// Note that this method returns the height in HIMETRIC units. To convert it to
// pixels, use [win.HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [get_Height]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_height
func (me *IPicture) Height() (int, error) {
	var cy int32
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Height,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cy)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cy), nil
	} else {
		return 0, hr
	}
}

// [get_hPal] method.
//
// [get_hPal]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_hpal
func (me *IPicture) HPal() (HPALETTE, error) {
	var hPal HPALETTE
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_hPal,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hPal)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hPal, nil
	} else {
		return HPALETTE(0), hr
	}
}

// [get_KeepOriginalFormat] method.
//
// [get_KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_keeporiginalformat
func (me *IPicture) KeepOriginalFormat() (bool, error) {
	var keep BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_KeepOriginalFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&keep)))
	return utl.HresultToBoolError(int32(keep), ret)
}

// [PictureChanged] method.
//
// [PictureChanged]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-picturechanged
func (me *IPicture) PictureChanged() error {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).PictureChanged,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// [Render] method.
//
// Example:
//
//	var wnd *ui.Main // initialized somewhere
//	var pic *win.IPicture
//
//	wnd.On().WmPaint(func() {
//		var ps win.PAINTSTRUCT
//		hdc, _ := wnd.Hwnd().BeginPaint(&ps)
//		defer wnd.Hwnd().EndPaint(&ps)
//
//		sz, _ := pic.Size()
//		_, _ = pic.Render(hdc,
//			win.POINT{},
//			win.SIZE{Cx: ps.RcPaint.Right, Cy: ps.RcPaint.Bottom},
//			win.POINT{X: 0, Y: sz.Cy},
//			win.SIZE{Cx: sz.Cx, Cy: -sz.Cy},
//		)
//	})
//
// [Render]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-render
func (me *IPicture) Render(
	hdc HDC,
	destOffset POINT,
	destSz SIZE,
	srcOffset POINT,
	srcSz SIZE,
) (metafileBounds RECT, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Render,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hdc),
		uintptr(destOffset.X),
		uintptr(destOffset.Y),
		uintptr(destSz.Cx),
		uintptr(destSz.Cy),
		uintptr(srcOffset.X),
		uintptr(srcOffset.Y),
		uintptr(srcSz.Cx),
		uintptr(srcSz.Cy),
		uintptr(unsafe.Pointer(&metafileBounds)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		metafileBounds = RECT{}
	}
	return
}

// [SaveAsFile] method.
//
// [SaveAsFile]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-saveasfile
func (me *IPicture) SaveAsFile(stream *IStream, saveCopy bool) (numBytesWritten int, hr error) {
	var written32 int32
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).SaveAsFile,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		utl.BoolToUintptr(saveCopy),
		uintptr(unsafe.Pointer(&written32)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(written32), nil
	} else {
		return 0, hr
	}
}

// [SelectPicture] method.
//
// [SelectPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-selectpicture
func (me *IPicture) SelectPicture(hdc HDC) (HDC, HBITMAP, error) {
	var hdcOut HDC
	var hBmp HBITMAP

	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).SelectPicture,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&hdcOut)),
		uintptr(unsafe.Pointer(&hBmp)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hdcOut, hBmp, nil
	} else {
		return HDC(0), HBITMAP(0), hr
	}
}

// [put_KeepOriginalFormat] method.
//
// [put_KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-put_keeporiginalformat
func (me *IPicture) SetKeepOriginalFormat(keep bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Put_KeepOriginalFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(keep))
	return utl.HresultToError(ret)
}

// [set_hPal] method.
//
// [set_hPal]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-set_hpal
func (me *IPicture) SetHPal(hPal HPALETTE) error {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Set_hPal,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hPal))
	return utl.HresultToError(ret)
}

// Calls [IPicture.Width] and [IPicture.Height] at once.
//
// If you need both width and height, call [IPicture.Size], which returns both.
//
// Note that this method returns the size in HIMETRIC units. To convert it to
// pixels, use [HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [IPicture.Width]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_width
// [IPicture.Height]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_height
func (me *IPicture) Size() (SIZE, error) {
	width, err := me.Width()
	if err != nil {
		return SIZE{}, err
	}

	height, err := me.Height()
	if err != nil {
		return SIZE{}, err
	}

	return SIZE{Cx: int32(width), Cy: int32(height)}, nil
}

// Calls [IPicture.Width] and [IPicture.Height], then convers from HIMETRIC
// units to pixels with [HDC.HiMetricToPixel].
//
// If hdc is zero, the method will retrieve the HDC for the whole screen with
// [HWND.GetDC].
//
// Example:
//
//	hdcScreen, _ := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
//	sz, _ := pic.SizePixels(hdcScreen)
func (me *IPicture) SizePixels(hdc HDC) (SIZE, error) {
	myHdc := hdc
	if myHdc == 0 {
		myHdc, err := HWND(0).GetDC() // DC of the entire screen
		if err != nil {
			return SIZE{}, err
		}
		defer HWND(0).ReleaseDC(myHdc)
	}

	himetricX, err := me.Width()
	if err != nil {
		return SIZE{}, err
	}
	himetricY, err := me.Height()
	if err != nil {
		return SIZE{}, err
	}

	pixelX, pixelY := myHdc.HiMetricToPixel(himetricX, himetricY)
	return SIZE{Cx: int32(pixelX), Cy: int32(pixelY)}, nil
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_type
func (me *IPicture) Type() (co.PICTYPE, error) {
	var picty co.PICTYPE
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Type,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&picty)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return picty, nil
	} else {
		return co.PICTYPE(0), hr
	}
}

// [get_Width] method.
//
// Note that this method returns the width in HIMETRIC units. To convert it to
// pixels, use [HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [get_Width]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_width
func (me *IPicture) Width() (int, error) {
	var cx int32
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Width,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cx)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cx), nil
	} else {
		return 0, hr
	}
}

type _IPictureVt struct {
	_IUnknownVt
	Get_Handle             uintptr
	Get_hPal               uintptr
	Get_Type               uintptr
	Get_Width              uintptr
	Get_Height             uintptr
	Render                 uintptr
	Set_hPal               uintptr
	Get_CurDC              uintptr
	SelectPicture          uintptr
	Get_KeepOriginalFormat uintptr
	Put_KeepOriginalFormat uintptr
	PictureChanged         uintptr
	SaveAsFile             uintptr
	Get_Attributes         uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IPropertyStore] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IPropertyStore]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nn-propsys-ipropertystore
type IPropertyStore struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPropertyStore) IID() co.IID {
	return co.IID_IPropertyStore
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-commit
func (me *IPropertyStore) Commit() error {
	ret, _, _ := syscall.SyscallN(
		(*_IPropertyStoreVt)(unsafe.Pointer(*me.Ppvt())).Commit,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// Returns all [co.PKEY] values by calling [IPropertyStore.GetCount] and
// [IPropertyStore.GetAt].
func (me *IPropertyStore) Enum() ([]co.PKEY, error) {
	count, hr := me.GetCount()
	if hr != nil {
		return nil, hr
	}

	pkeys := make([]co.PKEY, 0, count)
	for i := 0; i < count; i++ {
		pkey, hr := me.GetAt(i)
		if hr != nil {
			return nil, hr
		}
		pkeys = append(pkeys, pkey)
	}
	return pkeys, nil
}

// [GetAt] method.
//
// [GetAt]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-getat
func (me *IPropertyStore) GetAt(index int) (co.PKEY, error) {
	var guidPkey GUID
	ret, _, _ := syscall.SyscallN(
		(*_IPropertyStoreVt)(unsafe.Pointer(*me.Ppvt())).GetAt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&guidPkey)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.PKEY(guidPkey.String()), nil
	} else {
		return co.PKEY(""), hr
	}
}

// [GetCount] method.
//
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-getcount
func (me *IPropertyStore) GetCount() (int, error) {
	var cProps uint32
	ret, _, _ := syscall.SyscallN(
		(*_IPropertyStoreVt)(unsafe.Pointer(*me.Ppvt())).GetCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cProps)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cProps), nil
	} else {
		return 0, hr
	}
}

type _IPropertyStoreVt struct {
	_IUnknownVt
	GetCount uintptr
	GetAt    uintptr
	GetValue uintptr
	SetValue uintptr
	Commit   uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

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
func (me *ITypeInfo) GetContainingTypeLib(releaser *OleReleaser) (*ITypeLib, int, error) {
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
		return pObj, int(index), nil
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
			Ordinal: int(ordinal16),
		}, nil
	} else {
		return ITypeInfoDllEntry{}, hr
	}
}

// Returned by [ITypeInfo.GetDllEntry].
type ITypeInfoDllEntry struct {
	DllName string
	Name    string
	Ordinal int
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
			HelpContext: int(helpCtx),
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
	HelpContext int
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
func (me *ITypeInfo) GetFuncDesc(releaser *OleReleaser, index int) (*FuncDescData, error) {
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
func (me *ITypeInfo) GetImplTypeFlags(index int) (co.IMPLTYPEFLAG, error) {
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
	return utl.HresultToError(ret)
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [ITypeLib] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITypeLib]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypelib
type ITypeLib struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITypeLib) IID() co.IID {
	return co.IID_ITypeLib
}

type _ITypeLibVt struct {
	GetTypeInfoCount  uintptr
	GetTypeInfo       uintptr
	GetTypeInfoType   uintptr
	GetTypeInfoOfGuid uintptr
	GetLibAttr        uintptr
	GetTypeComp       uintptr
	GetDocumentation  uintptr
	IsName            uintptr
	FindName          uintptr
	ReleaseTLibAttr   uintptr
}
