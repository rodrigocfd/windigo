//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IShellFolder] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	item, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp")
//
//	folder, _ := shell.BindToHandler[shell.IShellFolder](
//		item, rel, nil, co.BHID_SFObject)
//
// [IShellFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellfolder
type IShellFolder struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellFolder) IID() co.IID {
	return co.IID_IShellFolder
}

// [BindToObject] method. Not implemented as a method of [IShellFolder] because
// Go doesn't support generic methods.
//
// [BindToObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtoobject
func BindToObject[T any, P ole.ComCtor[T]](
	iShellFolder *IShellFolder,
	releaser *ole.Releaser,
	pidl *ITEMIDLIST,
	bindCtx *ole.IBindCtx,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	riidGuid := win.GuidFrom(pObj.IID())

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*iShellFolder.Ppvt())).BindToObject,
		uintptr(*pidl),
		uintptr(unsafe.Pointer(pBindCtx)),
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

// [BindToStorage] method. Not implemented as a method of [IShellFolder] because
// Go doesn't support generic methods.
//
// [BindToStorage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtostorage
func BindToStorage[T any, P ole.ComCtor[T]](
	iShellFolder *IShellFolder,
	releaser *ole.Releaser,
	pidl *ITEMIDLIST,
	bindCtx *ole.IBindCtx,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	riidGuid := win.GuidFrom(pObj.IID())

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*iShellFolder.Ppvt())).BindToStorage,
		uintptr(*pidl),
		uintptr(unsafe.Pointer(pBindCtx)),
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

type _IShellFolderVt struct {
	ole.IUnknownVt
	ParseDisplayName uintptr
	EnumObjects      uintptr
	BindToObject     uintptr
	BindToStorage    uintptr
	CompareIDs       uintptr
	CreateViewObject uintptr
	GetAttributesOf  uintptr
	GetUIObjectOf    uintptr
	GetDisplayNameOf uintptr
	SetNameOf        uintptr
}
