//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IDataObject] COM interface.
//
// [IDataObject]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-idataobject
type IDataObject struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDataObject) IID() co.IID {
	return co.IID_IDataObject
}

// [GetCanonicalFormatEtc] method.
//
// [GetCanonicalFormatEtc]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-getcanonicalformatetc
func (me *IDataObject) GetCanonicalFormatEtc(etcIn *FORMATETC) (FORMATETC, error) {
	var etcOut FORMATETC
	ret, _, _ := syscall.SyscallN(
		(*vt.IDataObject)(unsafe.Pointer(*me.Ppvt())).GetCanonicalFormatEtc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(etcIn)), uintptr(unsafe.Pointer(&etcOut)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return etcOut, nil
	} else {
		return FORMATETC{}, hr
	}
}

// [GetData] method.
//
// ⚠️ You must defer ReleaseStgMedium() on the returned object.
//
// [GetData]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-getdata
func (me *IDataObject) GetData(etc *FORMATETC) (STGMEDIUM, error) {
	var stg STGMEDIUM
	ret, _, _ := syscall.SyscallN(
		(*vt.IDataObject)(unsafe.Pointer(*me.Ppvt())).GetData,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(etc)), uintptr(unsafe.Pointer(&stg)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stg, nil
	} else {
		return STGMEDIUM{}, hr
	}
}

// [QueryGetData] method.
//
// [QueryGetData]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-querygetdata
func (me *IDataObject) QueryGetData(etc *FORMATETC) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.IDataObject)(unsafe.Pointer(*me.Ppvt())).QueryGetData,
		uintptr(unsafe.Pointer(me.Ppvt())), uintptr(unsafe.Pointer(etc)))
	return wutil.ErrorAsHResult(ret)
}
