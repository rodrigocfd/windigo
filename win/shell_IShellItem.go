//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IShellItem] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Usually created with [SHCreateItemFromParsingName].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem) IID() co.IID {
	return co.IID_IShellItem
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
//	var desktop *win.IShellItem
//	_ = win.SHGetKnownFolderItem(
//		rel,
//		co.FOLDERID_Desktop,
//		co.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	var enumItems *win.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, co.BHID_EnumItems, &enumItems)
//
// [BindToHandler]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-bindtohandler
func (me *IShellItem) BindToHandler(
	releaser *OleReleaser,
	bindCtx *IBindCtx,
	bhid co.BHID,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidBhid := GuidFrom(bhid)
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).BindToHandler,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(ppvtOrNil(bindCtx)),
		uintptr(unsafe.Pointer(&guidBhid)),
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

// [Compare] method.
//
// [Compare]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-compare
func (me *IShellItem) Compare(si *IShellItem, hint co.SICHINT) (bool, error) {
	var piOrder uint32
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).Compare,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())),
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
func (me *IShellItem) GetAttributes(mask co.SFGAO) (attrs co.SFGAO, exactMatch bool, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetAttributes,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&mask)),
		uintptr(unsafe.Pointer(&attrs)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		exactMatch, hr = true, nil
	} else if hr == co.HRESULT_S_FALSE {
		exactMatch, hr = false, nil
	} else {
		attrs, exactMatch = co.SFGAO(0), false
	}
	return
}

// [GetDisplayName] method.
//
// Example:
//
//	var shi win.IShellItem // initialized somewhere
//
//	fullPath, _ := shi.GetDisplayName(co.SIGDN_FILESYSPATH)
//
// [GetDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName co.SIGDN) (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetDisplayName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(sigdnName),
		uintptr(unsafe.Pointer(&pv)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(pv).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetParent] method.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
func (me *IShellItem) GetParent(releaser *OleReleaser) (*IShellItem, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetParent,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItem{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IShellItemVt struct {
	_IUnknownVt
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}
