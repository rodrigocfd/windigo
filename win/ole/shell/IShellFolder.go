//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IShellFolder] COM interface.
//
// Implements [ole.ComObj] and [ole.ComResource].
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
//	item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
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
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToObject,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(*pidl),
		uintptr(unsafe.Pointer(pBindCtx)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(ole.ComObj)
		releaser.Add(pOut)
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
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToStorage,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(*pidl),
		uintptr(unsafe.Pointer(pBindCtx)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(ole.ComObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

// [CompareIDs] method.
//
// [CompareIDs]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-compareids
func (me *IShellFolder) CompareIDs(
	sortingRule uint16,
	sortingFlags co.SHCIDS,
	pidl1, pidl2 *ITEMIDLIST,
) (int, error) {
	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).CompareIDs,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(sortingRule)|uint32(sortingFlags)),
		uintptr(*pidl1),
		uintptr(*pidl2))

	if hr := co.HRESULT(ret); hr.Succeeded() {
		return int(hr.Code()), nil
	} else {
		return 0, hr
	}
}

// [CreateViewObject] method.
//
// Return type is typically [IShellView].
//
// [CreateViewObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-createviewobject
func (me *IShellFolder) CreateViewObject(
	releaser *ole.Releaser,
	hwndOwner win.HWND,
	ppOut interface{},
) error {
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).CreateViewObject,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(ole.ComObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

// [ParseDisplayName] method.
//
// [ParseDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-parsedisplayname
func (me *IShellFolder) ParseDisplayName(
	releaser *ole.Releaser,
	hWnd win.HWND,
	bindCtx *ole.IBindCtx,
	displayName string,
	attributes co.SFGAO,
) (*ITEMIDLIST, co.SFGAO, error) {
	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	displayName16 := wstr.NewBufWith[wstr.Stack20](displayName, wstr.ALLOW_EMPTY)
	var chEaten uint32
	var idl ITEMIDLIST

	var pSfgao *co.SFGAO
	if attributes != co.SFGAO(0) {
		pSfgao = &attributes
	}

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).ParseDisplayName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pBindCtx)),
		uintptr(displayName16.UnsafePtr()),
		uintptr(unsafe.Pointer(&chEaten)),
		uintptr(unsafe.Pointer(&idl)),
		uintptr(unsafe.Pointer(pSfgao)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pIdl := &idl
		releaser.Add(pIdl)
		return pIdl, *pSfgao, nil
	} else {
		return nil, co.SFGAO(0), hr
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
