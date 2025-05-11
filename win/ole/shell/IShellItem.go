//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IShellItem] COM interface.
//
// Usually created with [SHCreateItemFromParsingName].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	ish, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\foo.txt")
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
type IShellItem struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem) IID() co.IID {
	return co.IID_IShellItem
}

// [BindToHandler] method. Not implemented as a method of [IShellItem] because
// Go doesn't support generic methods.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	desktop, _ := shell.SHGetKnownFolderItem[shell.IShellItem](
//		rel, co.FOLDERID_Desktop, co.KF_FLAG_DEFAULT, win.HANDLE(0))
//
//	enumItems, _ := shell.BindToHandler[shell.IEnumShellItems](
//		desktop, rel, nil, co.BHID_EnumItems)
//
// [BindToHandler]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-bindtohandler
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
func BindToHandler[T any, P ole.ComCtor[T]](
	iShellItem *IShellItem,
	releaser *ole.Releaser,
	bindCtx *ole.IBindCtx,
	bhid co.BHID,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	bhidGuid := win.GuidFrom(bhid)
	riidGuid := win.GuidFrom(pObj.IID())

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*iShellItem.Ppvt())).BindToHandler,
		uintptr(unsafe.Pointer(iShellItem.Ppvt())),
		uintptr(unsafe.Pointer(pBindCtx)),
		uintptr(unsafe.Pointer(&bhidGuid)),
		uintptr(unsafe.Pointer(&riidGuid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj.Set(ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
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
// # Example
//
//	var shi shell.IShellItem // initialized somewhere
//
//	fullPath, _ := shi.GetDisplayName(co.SIGDN_FILESYSPATH)
//
// [GetDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName co.SIGDN) (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetDisplayName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(sigdnName), uintptr(unsafe.Pointer(&pv)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer ole.HTASKMEM(pv).CoTaskMemFree()
		name := wstr.WstrPtrToStr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetParent] method.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
func (me *IShellItem) GetParent(releaser *ole.Releaser) (*IShellItem, error) {
	var ppvtQueried **ole.IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetParent,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := ole.ComObj[IShellItem](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IShellItemVt struct {
	ole.IUnknownVt
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}
