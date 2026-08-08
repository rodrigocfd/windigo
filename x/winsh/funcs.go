//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [CommandLineToArgv] function.
//
// Typically used with [GetCommandLine].
//
// [CommandLineToArgv]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-commandlinetoargvw
func CommandLineToArgv(cmdLine string) ([]string, error) {
	var wCmdLine wstr.BufEncoder
	var pNumArgs int32

	ret, _, err := syscall.SyscallN(
		dll.Shell.Load(&_shell_CommandLineToArgvW, "CommandLineToArgvW"),
		uintptr(wCmdLine.EmptyIsNil(cmdLine)),
		uintptr(unsafe.Pointer(&pNumArgs)))
	if ret == 0 {
		return nil, co.ERROR(err)
	}

	lpPtrs := unsafe.Slice((**uint16)(unsafe.Pointer(ret)), pNumArgs) // []*uint16
	strs := make([]string, 0, pNumArgs)

	for _, lpPtr := range lpPtrs {
		strs = append(strs, wstr.DecodePtr(lpPtr))
	}
	return strs, nil
}

var _shell_CommandLineToArgvW *syscall.Proc

// [SHCreateItemFromIDList] function.
//
// Return type is typically [IShellItem] of [IShellItem2].
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
//	_ = winsh.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
//	pidl, _ := winsh.SHGetIDListFromObject(rel, &item.IUnknown)
//
//	var sameItem *winsh.IShellItem2
//	_ = winsh.SHCreateItemFromIDList(rel, pidl, &sameItem)
//
// [SHCreateItemFromIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromidlist
func SHCreateItemFromIDList(releaser *win.OleReleaser, pidl *ITEMIDLIST, ppOut interface{}) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHCreateItemFromIDList, "SHCreateItemFromIDList"),
		uintptr(*pidl),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

var _shell_SHCreateItemFromIDList *syscall.Proc

// [SHCreateItemFromParsingName] function.
//
// Return type is typically [IShellItem] of [IShellItem2].
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
//	_ = winsh.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func SHCreateItemFromParsingName(
	releaser *win.OleReleaser,
	folderOrFilePath string,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr
	var wFolderOrFilePath wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHCreateItemFromParsingName, "SHCreateItemFromParsingName"),
		uintptr(wFolderOrFilePath.EmptyIsNil(folderOrFilePath)),
		0,
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

var _shell_SHCreateItemFromParsingName *syscall.Proc

// [SHCreateItemFromRelativeName] function.
//
// [SHCreateItemFromRelativeName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromrelativename
func SHCreateItemFromRelativeName(
	releaser *win.OleReleaser,
	parent *IShellItem,
	name string,
	bindCtx *win.IBindCtx,
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr
	var wName wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHCreateItemFromRelativeName, "SHCreateItemFromRelativeName"),
		uintptr(unsafe.Pointer(parent.Ppvt())),
		uintptr(wName.AllowEmpty(name)),
		utl.OlePpvtOrNil(bindCtx),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

var _shell_SHCreateItemFromRelativeName *syscall.Proc

// [SHCreateMemStream] function.
//
// Creates an [IStream] projection over a slice, which must remain valid in
// memory throughout IStream's lifetime.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	data := []byte{0x10, 0x11, 0x12}
//	defer runtime.KeepAlive(data)
//
//	stream, _ := winsh.SHCreateMemStream(rel, data)
//
// [SHCreateMemStream]: https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func SHCreateMemStream(releaser *win.OleReleaser, src []byte) (*win.IStream, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Shlwapi.Load(&_shlwapi_SHCreateMemStream, "SHCreateMemStream"),
		uintptr(unsafe.Pointer(&src[0])),
		uintptr(uint32(len(src))))
	if ret == 0 {
		return nil, co.HRESULT_E_OUTOFMEMORY
	}

	pObj := utl.OleNew[*win.IStream](ret, releaser)
	return pObj, nil
}

var _shlwapi_SHCreateMemStream *syscall.Proc

// [SHCreateShellItemArray] function.
//
// [SHCreateShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateshellitemarray
func SHCreateShellItemArray(
	releaser *win.OleReleaser,
	pidlParent *ITEMIDLIST,
	parent *IShellFolder,
	pidlChildren []*ITEMIDLIST,
) (*IShellItemArray, error) {
	var ppvtQueried uintptr

	var pidlParentObj ITEMIDLIST
	if pidlParent != nil {
		pidlParentObj = *pidlParent
	}

	var pidlChildrenObjsPtr *ITEMIDLIST
	if pidlChildren != nil {
		pidlChildrenObjs := make([]ITEMIDLIST, 0, len(pidlChildren))
		for _, pidl := range pidlChildren {
			pidlChildrenObjs = append(pidlChildrenObjs, *pidl)
		}
		pidlChildrenObjsPtr = &pidlChildrenObjs[0]
	}

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHCreateShellItemArray, "SHCreateShellItemArray"),
		uintptr(pidlParentObj),
		utl.OlePpvtOrNil(parent),
		uintptr(uint32(len(pidlChildren))),
		uintptr(unsafe.Pointer(pidlChildrenObjsPtr)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IShellItemArray](ret, ppvtQueried, releaser)
}

var _shell_SHCreateShellItemArray *syscall.Proc

// [SHCreateShellItemArrayFromIDLists] function.
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
//	var item1, item2 *win.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item1)
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\bar.txt", &item2)
//
//	pidl1, _ := winsh.SHGetIDListFromObject(rel, &item1.IUnknown)
//	pidl2, _ := winsh.SHGetIDListFromObject(rel, &item2.IUnknown)
//
//	arr, _ := winsh.SHCreateShellItemArrayFromIDLists(rel, pidl1, pidl2)
//
// [SHCreateShellItemArrayFromIDLists]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateshellitemarrayfromidlists?redirectedfrom=MSDN
func SHCreateShellItemArrayFromIDLists(
	releaser *win.OleReleaser,
	pidls ...*ITEMIDLIST,
) (*IShellItemArray, error) {
	var ppvtQueried uintptr

	pidlObjs := make([]ITEMIDLIST, 0, len(pidls))
	for _, pidl := range pidls {
		pidlObjs = append(pidlObjs, *pidl)
	}

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHCreateShellItemArrayFromIDLists, "SHCreateShellItemArrayFromIDLists"),
		uintptr(uint32(len(pidls))),
		uintptr(unsafe.Pointer(&pidlObjs[0])),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IShellItemArray](ret, ppvtQueried, releaser)
}

var _shell_SHCreateShellItemArrayFromIDLists *syscall.Proc

// [Shell_NotifyIcon] function.
//
// [Shell_NotifyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func Shell_NotifyIcon(message cosh.NIM, pData *NOTIFYICONDATA) error {
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_Shell_NotifyIconW, "Shell_NotifyIconW"),
		uintptr(message),
		uintptr(unsafe.Pointer(pData)))
	if ret == 0 {
		return co.ERROR_UNIDENTIFIED_ERROR
	}
	return nil
}

var _shell_Shell_NotifyIconW *syscall.Proc

// [Shell_NotifyIconGetRect] function.
//
// [Shell_NotifyIconGetRect]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyicongetrect
func Shell_NotifyIconGetRect(pIdent *NOTIFYICONIDENTIFIER) (win.RECT, error) {
	var rc win.RECT
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_Shell_NotifyIconGetRect, "Shell_NotifyIconGetRect"),
		uintptr(unsafe.Pointer(pIdent)),
		uintptr(unsafe.Pointer(&rc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return win.RECT{}, hr
	}
	return rc, nil
}

var _shell_Shell_NotifyIconGetRect *syscall.Proc

// [SHGetDesktopFolder] function.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	folder, _ := winsh.SHGetDesktopFolder(rel)
//
// [SHGetDesktopFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shgetdesktopfolder
func SHGetDesktopFolder(releaser *win.OleReleaser) (*IShellFolder, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHGetDesktopFolder, "SHGetDesktopFolder"),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IShellFolder](ret, ppvtQueried, releaser)
}

var _shell_SHGetDesktopFolder *syscall.Proc

// [SHGetFileInfo] function.
//
// ⚠️ You must defer [HICON.DestroyIcon] on the HIcon member of the returned
// struct.
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(path string, fileAttrs co.FILE_ATTRIBUTE, flags cosh.SHGFI) (SHFILEINFO, error) {
	var wPath wstr.BufEncoder
	var sfi SHFILEINFO

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&utl.Shell_SHGetFileInfoW, "SHGetFileInfoW"), // note: syscall.Proc from utl
		uintptr(wPath.AllowEmpty(path)),
		uintptr(fileAttrs),
		uintptr(unsafe.Pointer(&sfi)),
		unsafe.Sizeof(sfi),
		uintptr(flags))

	if ret == 0 {
		return SHFILEINFO{}, co.ERROR_UNIDENTIFIED_ERROR
	}
	return sfi, nil
}

// [SHGetKnownFolderItem] function.
//
// Return type is typically [IShellItem] of [IShellItem2].
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
//	_ = winsh.SHGetKnownFolderItem(
//		rel,
//		&co.FOLDERID_Desktop,
//		co.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	path, _ := desktop.GetDisplayName(co.SIGDN_FILESYSPATH)
//	println(path)
//
// [SHGetKnownFolderItem]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shgetknownfolderitem
func SHGetKnownFolderItem(
	releaser *win.OleReleaser,
	pKfid *cosh.FOLDERID,
	flags cosh.KF,
	hToken win.HANDLE, // HACCESSTOKEN
	ppOut interface{},
) error {
	piid := utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHGetKnownFolderItem, "SHGetKnownFolderItem"),
		uintptr(unsafe.Pointer(pKfid)),
		uintptr(flags),
		uintptr(hToken),
		uintptr(unsafe.Pointer(piid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

var _shell_SHGetKnownFolderItem *syscall.Proc

// [SHGetIDListFromObject] function.
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
//	_ = winsh.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
//	idl, _ := winsh.SHGetIDListFromObject(rel, &item.IUnknown)
//
// [SHGetIDListFromObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetidlistfromobject
func SHGetIDListFromObject(releaser *win.OleReleaser, obj *win.IUnknown) (*ITEMIDLIST, error) {
	var idl ITEMIDLIST
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHGetIDListFromObject, "SHGetIDListFromObject"),
		uintptr(unsafe.Pointer(obj.Ppvt())),
		uintptr(unsafe.Pointer(&idl)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pIdl := &idl
		releaser.Add(pIdl)
		return pIdl, nil
	} else {
		return nil, hr
	}
}

var _shell_SHGetIDListFromObject *syscall.Proc

// [SHGetPropertyStoreFromIDList] function.
//
// [SHGetPropertyStoreFromIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetpropertystorefromidlist
func SHGetPropertyStoreFromIDList(
	releaser *win.OleReleaser,
	pidl *ITEMIDLIST,
	flags cosh.GPS,
) (*IPropertyStore, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHGetPropertyStoreFromIDList, "SHGetPropertyStoreFromIDList"),
		uintptr(*pidl),
		uintptr(flags),
		uintptr(unsafe.Pointer(&cosh.IID_IPropertyStore)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IPropertyStore](ret, ppvtQueried, releaser)
}

var _shell_SHGetPropertyStoreFromIDList *syscall.Proc

// [SHGetPropertyStoreFromParsingName] function.
//
// [SHGetPropertyStoreFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetpropertystorefromparsingname
func SHGetPropertyStoreFromParsingName(
	releaser *win.OleReleaser,
	folderOrFilePath string,
	bindCtx *win.IBindCtx,
	flags cosh.GPS,
) (*IPropertyStore, error) {
	var ppvtQueried uintptr
	var wFolderOrFilePath wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_SHGetPropertyStoreFromParsingName, "SHGetPropertyStoreFromParsingName"),
		uintptr(wFolderOrFilePath.AllowEmpty(folderOrFilePath)),
		utl.OlePpvtOrNil(bindCtx),
		uintptr(flags),
		uintptr(unsafe.Pointer(&cosh.IID_IPropertyStore)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IPropertyStore](ret, ppvtQueried, releaser)
}

var _shell_SHGetPropertyStoreFromParsingName *syscall.Proc
