//go:build windows

package shell

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
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
//	var item *shell.IShellItem
//	shell.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *shell.IShellFolder
//	shell.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
// [IShellFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellfolder
type IShellFolder struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellFolder) IID() co.IID {
	return co.IID_IShellFolder
}

// [BindToObject] method.
//
// [BindToObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtoobject
func (me *IShellFolder) BindToObject(
	releaser *ole.Releaser,
	pidl *ITEMIDLIST,
	bindCtx *ole.IBindCtx,
	ppOut interface{},
) error {
	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(utl.ComRetrieveIid(ppOut))

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToObject,
		uintptr(*pidl),
		uintptr(unsafe.Pointer(pBindCtx)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried))
		releaser.Add(reflect.ValueOf(ppOut).Elem().Interface().(ole.ComResource))
		return nil
	} else {
		return hr
	}
}

// [BindToStorage] method.
//
// [BindToStorage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtostorage
func (me *IShellFolder) BindToStorage(
	releaser *ole.Releaser,
	pidl *ITEMIDLIST,
	bindCtx *ole.IBindCtx,
	ppOut interface{},
) error {
	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(utl.ComRetrieveIid(ppOut))

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToStorage,
		uintptr(*pidl),
		uintptr(unsafe.Pointer(pBindCtx)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried))
		releaser.Add(reflect.ValueOf(ppOut).Elem().Interface().(ole.ComResource))
		return nil
	} else {
		return hr
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
