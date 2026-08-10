//go:build windows

package winwic

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cowic"
)

// [IWICComponentInfo] COM interface.
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

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IWICComponentInfo) AddRef(releaser *win.OleReleaser) *IWICComponentInfo {
	return utl.OleNewFromAddRef[*IWICComponentInfo](me, releaser)
}

// [GetAuthor] method.
//
// [GetAuthor]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwiccomponentinfo-getauthor
func (me *IWICComponentInfo) GetAuthor() (string, error) {
	return oleCallAllocBufRetStr(me, utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetAuthor)
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
	return oleCallAllocBufRetStr(me, utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetFriendlyName)
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
	return oleCallAllocBufRetStr(me, utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetSpecVersion)
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
	return oleCallAllocBufRetStr(me, utl.Vt[_IWICComponentInfoVt](me.Ppvt()).GetVersion)
}
