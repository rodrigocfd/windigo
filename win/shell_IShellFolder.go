//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IShellFolder] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// # Example
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *win.IShellItem
//	win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
// [IShellFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellfolder
type IShellFolder struct{ IUnknown }

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
	releaser *OleReleaser,
	pidl *ITEMIDLIST,
	bindCtx *IBindCtx,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToObject,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(*pidl),
		uintptr(ppvtOrNil(bindCtx)),
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

// [BindToStorage] method.
//
// [BindToStorage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtostorage
func (me *IShellFolder) BindToStorage(
	releaser *OleReleaser,
	pidl *ITEMIDLIST,
	bindCtx *IBindCtx,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToStorage,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(*pidl),
		uintptr(ppvtOrNil(bindCtx)),
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
	releaser *OleReleaser,
	hwndOwner HWND,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).CreateViewObject,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndOwner),
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

// [ParseDisplayName] method.
//
// [ParseDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-parsedisplayname
func (me *IShellFolder) ParseDisplayName(
	releaser *OleReleaser,
	hWnd HWND,
	bindCtx *IBindCtx,
	displayName string,
	attributes co.SFGAO,
) (*ITEMIDLIST, co.SFGAO, error) {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pDisplayName := wbuf.PtrAllowEmpty(displayName)

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
		uintptr(ppvtOrNil(bindCtx)),
		uintptr(pDisplayName),
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

// [SetNameOf] method.
//
// [SetNameOf]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-setnameof
func (me *IShellFolder) SetNameOf(
	releaser *OleReleaser,
	hWnd HWND,
	pidl *ITEMIDLIST,
	name string,
	flags co.SHGDN,
) (*ITEMIDLIST, error) {
	var idlChild ITEMIDLIST

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pName := wbuf.PtrAllowEmpty(name)

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).SetNameOf,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(*pidl)),
		uintptr(pName),
		uintptr(flags),
		uintptr(unsafe.Pointer(&idlChild)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pIdlChild := &idlChild
		releaser.Add(pIdlChild)
		return pIdlChild, nil
	} else {
		return nil, hr
	}
}

type _IShellFolderVt struct {
	_IUnknownVt
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
