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

// [IShellFolder] COM interface.
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
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *winsh.IShellFolder
//	_ = item.BindToHandler(rel, nil, &cosh.BHID_SFObject, &folder)
//
// [IShellFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellfolder
type IShellFolder struct{ win.IUnknown }

type _IShellFolderVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellFolder) IID() *co.IID {
	return &cosh.IID_IShellFolder
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IShellFolder) AddRef(releaser *win.OleReleaser) *IShellFolder {
	return utl.OleNewFromAddRef[*IShellFolder](me, releaser)
}

// [BindToObject] method.
//
// [BindToObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtoobject
func (me *IShellFolder) BindToObject(
	releaser *win.OleReleaser,
	pidl *ITEMIDLIST,
	bindCtx *win.IBindCtx,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellFolderVt](me.Ppvt()).BindToObject,
		me.Ppvt(),
		uintptr(*pidl),
		utl.OlePpvtOrNil(bindCtx),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [BindToStorage] method.
//
// [BindToStorage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-bindtostorage
func (me *IShellFolder) BindToStorage(
	releaser *win.OleReleaser,
	pidl *ITEMIDLIST,
	bindCtx *win.IBindCtx,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellFolderVt](me.Ppvt()).BindToStorage,
		me.Ppvt(),
		uintptr(*pidl),
		utl.OlePpvtOrNil(bindCtx),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [CompareIDs] method.
//
// [CompareIDs]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-compareids
func (me *IShellFolder) CompareIDs(
	sortingRule uint16,
	sortingFlags cosh.SHCIDS,
	pidl1, pidl2 *ITEMIDLIST,
) (int, error) {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellFolderVt](me.Ppvt()).CompareIDs,
		me.Ppvt(),
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
	releaser *win.OleReleaser,
	hwndOwner win.HWND,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellFolderVt](me.Ppvt()).CreateViewObject,
		me.Ppvt(),
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [EnumObjects] method.
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
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *winsh.IShellFolder
//	_ = item.BindToHandler(rel, nil, &cosh.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(
//		rel,
//		win.HWND(0),
//		cosh.SHCONTF_FOLDERS|cosh.SHCONTF_NONFOLDERS|cosh.SHCONTF_INCLUDEHIDDEN,
//	)
//
// [EnumObjects]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-enumobjects
func (me *IShellFolder) EnumObjects(
	releaser *win.OleReleaser,
	hWnd win.HWND,
	flags cosh.SHCONTF,
) (*IEnumIDList, error) {
	return utl.OleNewFromCallWithoutParms[*IEnumIDList](me, releaser,
		utl.Vt[_IShellFolderVt](me.Ppvt()).EnumObjects)
}

// [ParseDisplayName] method.
//
// [ParseDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-parsedisplayname
func (me *IShellFolder) ParseDisplayName(
	releaser *win.OleReleaser,
	hWnd win.HWND,
	bindCtx *win.IBindCtx,
	displayName string,
	attributes cosh.SFGAO,
) (*ITEMIDLIST, cosh.SFGAO, error) {
	var wDisplayName wstr.BufEncoder
	var chEaten uint32
	var idl ITEMIDLIST

	var pSfgao *cosh.SFGAO
	if attributes != cosh.SFGAO(0) {
		pSfgao = &attributes
	}

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellFolderVt](me.Ppvt()).ParseDisplayName,
		me.Ppvt(),
		uintptr(hWnd),
		utl.OlePpvtOrNil(bindCtx),
		uintptr(wDisplayName.AllowEmpty(displayName)),
		uintptr(unsafe.Pointer(&chEaten)),
		uintptr(unsafe.Pointer(&idl)),
		uintptr(unsafe.Pointer(pSfgao)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, cosh.SFGAO(0), hr
	}

	pIdl := &idl
	releaser.Add(pIdl)
	return pIdl, *pSfgao, nil
}

// [SetNameOf] method.
//
// [SetNameOf]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-setnameof
func (me *IShellFolder) SetNameOf(
	releaser *win.OleReleaser,
	hWnd win.HWND,
	pidl *ITEMIDLIST,
	name string,
	flags cosh.SHGDN,
) (*ITEMIDLIST, error) {
	var idlChild ITEMIDLIST
	var wName wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellFolderVt](me.Ppvt()).SetNameOf,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(*pidl)),
		uintptr(wName.AllowEmpty(name)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&idlChild)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}

	pIdlChild := &idlChild
	releaser.Add(pIdlChild)
	return pIdlChild, nil
}
