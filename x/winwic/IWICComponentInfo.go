//go:build windows

package winwic

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICComponentInfo] COM interface.
//
// Implements [OleResource].
//
// [IWICComponentInfo]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nn-wincodec-iwiccomponentinfo
type IWICComponentInfo struct{ win.IUnknown }

type _IWICComponentInfoVt struct {
	utl.IUnknownVt
	GetComponentType uintptr
	GetCLSID         uintptr
	GetSigningStatus uintptr
	GetAuthor        uintptr
	GetVendorGUID    uintptr
	GetVersion       uintptr
	GetSpecVersion   uintptr
	GetFriendlyName  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IWICComponentInfo) IID() *co.IID {
	return &cowic.IID_IWICComponentInfo
}

// [GetAuthor] method.
//
// [GetAuthor]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getauthor
func (me *IWICComponentInfo) GetAuthor() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetAuthor,
		me.Ppvt(),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetAuthor,
		me.Ppvt(),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

// [GetCLSID] method.
//
// [GetCLSID]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getclsid
func (me *IWICComponentInfo) GetCLSID() (co.CLSID, error) {
	return utl.OleCallReturnStruct[co.CLSID](me,
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetCLSID)
}

// [GetFriendlyName] method.
//
// [GetFriendlyName]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getfriendlyname
func (me *IWICComponentInfo) GetFriendlyName() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetFriendlyName,
		me.Ppvt(),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetFriendlyName,
		me.Ppvt(),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

// [GetSigningStatus] method.
//
// [GetSigningStatus]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getsigningstatus
func (me *IWICComponentInfo) GetSigningStatus() (cowic.WIC_COMPONENTSIGN, error) {
	return utl.OleCallReturnStruct[cowic.WIC_COMPONENTSIGN](me,
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetSigningStatus)
}

// [GetSpecVersion] method.
//
// [GetSpecVersion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getspecversion
func (me *IWICComponentInfo) GetSpecVersion() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetSpecVersion,
		me.Ppvt(),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetSpecVersion,
		me.Ppvt(),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}

// [GetVendorGUID] method.
//
// [GetVendorGUID]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getvendorguid
func (me *IWICComponentInfo) GetVendorGUID() (co.GUID, error) {
	return utl.OleCallReturnStruct[co.GUID](me,
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetVendorGUID)
}

// [GetVersion] method.
//
// [GetVersion]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getversion
func (me *IWICComponentInfo) GetVersion() (string, error) {
	var szRequired uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetVersion,
		me.Ppvt(),
		0, 0,
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szRequired)
	ret, _, _ = syscall.SyscallN(
		utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetVersion,
		me.Ppvt(),
		uintptr(szRequired),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&szRequired)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.DecodeSlice(buf), nil
	} else {
		return "", hr
	}
}
