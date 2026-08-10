//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IShellItem] COM interface.
//
// Usually created with [SHCreateItemFromParsingName].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem struct{ win.IUnknown }

type _IShellItemVt struct {
	utl.IUnknownVt
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem) IID() *co.IID {
	return &cosh.IID_IShellItem
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IShellItem) AddRef(releaser *win.OleReleaser) *IShellItem {
	return utl.OleNewFromAddRef[*IShellItem](me, releaser)
}

// [BindToHandler] method.
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
//	var desktop *winsh.IShellItem
//	_ = winsh.SHGetKnownFolderItem(
//		rel,
//		&cosh.FOLDERID_Desktop,
//		cosh.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	var enumItems *winsh.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, &cosh.BHID_EnumItems, &enumItems)
//
// [BindToHandler]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-bindtohandler
func (me *IShellItem) BindToHandler(
	releaser *win.OleReleaser,
	bindCtx *win.IBindCtx,
	pBhid *co.BHID,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItemVt](me.Ppvt()).BindToHandler,
		me.Ppvt(),
		utl.OlePpvtOrNil(bindCtx),
		uintptr(unsafe.Pointer(pBhid)),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [Compare] method.
//
// [Compare]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-compare
func (me *IShellItem) Compare(si *IShellItem, hint cosh.SICHINT) (bool, error) {
	var piOrder int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItemVt](me.Ppvt()).Compare,
		me.Ppvt(),
		si.Ppvt(),
		uintptr(hint),
		uintptr(unsafe.Pointer(&piOrder)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return true, nil
	} else if hr == co.HRESULT_S_FALSE {
		return false, nil
	} else {
		return false, hr
	}
}

// [GetAttributes] method.
//
// [GetAttributes]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getattributes
func (me *IShellItem) GetAttributes(mask cosh.SFGAO) (attrs cosh.SFGAO, exactMatch bool, hr error) {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItemVt](me.Ppvt()).GetAttributes,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&mask)),
		uintptr(unsafe.Pointer(&attrs)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		exactMatch, hr = true, nil
	} else if hr == co.HRESULT_S_FALSE {
		exactMatch, hr = false, nil
	} else {
		attrs, exactMatch = cosh.SFGAO(0), false
	}
	return
}

// [GetDisplayName] method.
//
// Example:
//
//	var shi *winsh.IShellItem // initialized somewhere
//
//	fullPath, _ := shi.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//
// [GetDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName cosh.SIGDN) (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItemVt](me.Ppvt()).GetDisplayName,
		me.Ppvt(),
		uintptr(sigdnName),
		uintptr(unsafe.Pointer(&pv)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	defer win.HTASKMEM(pv).CoTaskMemFree()
	name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
	return name, nil
}

// [GetParent] method.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
func (me *IShellItem) GetParent(releaser *win.OleReleaser) (*IShellItem, error) {
	return utl.OleNewFromCallWithoutParms[*IShellItem](me, releaser,
		utl.Vt[_IShellItemVt](me.Ppvt()).GetParent)
}
