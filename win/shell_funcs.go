//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [CommandLineToArgv] function.
//
// Typically used with [GetCommandLine].
//
// [CommandLineToArgv]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-commandlinetoargvw
func CommandLineToArgv(cmdLine string) ([]string, error) {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pCmdLine := wbuf.PtrEmptyIsNil(cmdLine)

	var pNumArgs int32

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_CommandLineToArgvW, "CommandLineToArgvW"),
		uintptr(pCmdLine),
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

var _CommandLineToArgvW *syscall.Proc

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
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
//	pidl, _ := win.SHGetIDListFromObject(rel, &item.IUnknown)
//
//	var sameItem *win.IShellItem2
//	_ = win.SHCreateItemFromIDList(rel, pidl, &sameItem)
//
// [SHCreateItemFromIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromidlist
func SHCreateItemFromIDList(releaser *OleReleaser, pidl *ITEMIDLIST, ppOut interface{}) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHCreateItemFromIDList, "SHCreateItemFromIDList"),
		uintptr(*pidl),
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

var _SHCreateItemFromIDList *syscall.Proc

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
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func SHCreateItemFromParsingName(
	releaser *OleReleaser,
	folderOrFilePath string,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pFolderOrFilePath := wbuf.PtrEmptyIsNil(folderOrFilePath)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHCreateItemFromParsingName, "SHCreateItemFromParsingName"),
		uintptr(pFolderOrFilePath),
		0,
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

var _SHCreateItemFromParsingName *syscall.Proc

// [SHCreateItemFromRelativeName] function.
//
// [SHCreateItemFromRelativeName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromrelativename
func SHCreateItemFromRelativeName(
	releaser *OleReleaser,
	parent *IShellItem,
	name string,
	bindCtx *IBindCtx,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(pOut.IID())

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pName := wbuf.PtrAllowEmpty(name)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHCreateItemFromRelativeName, "SHCreateItemFromRelativeName"),
		uintptr(unsafe.Pointer(parent.Ppvt())),
		uintptr(pName),
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

var _SHCreateItemFromRelativeName *syscall.Proc

// [SHCreateShellItemArray] function.
//
// [SHCreateShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateshellitemarray
func SHCreateShellItemArray(
	releaser *OleReleaser,
	pidlParent *ITEMIDLIST,
	parent *IShellFolder,
	pidlChildren []*ITEMIDLIST,
) (*IShellItemArray, error) {
	var ppvtQueried **_IUnknownVt

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
		dll.Load(dll.SHELL32, &_SHCreateShellItemArray, "SHCreateShellItemArray"),
		uintptr(pidlParentObj),
		uintptr(ppvtOrNil(parent)),
		uintptr(uint32(len(pidlChildren))),
		uintptr(unsafe.Pointer(pidlChildrenObjsPtr)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItemArray{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHCreateShellItemArray *syscall.Proc

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
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item1)
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\bar.txt", &item2)
//
//	pidl1, _ := win.SHGetIDListFromObject(rel, &item1.IUnknown)
//	pidl2, _ := win.SHGetIDListFromObject(rel, &item2.IUnknown)
//
//	arr, _ := win.SHCreateShellItemArrayFromIDLists(rel, pidl1, pidl2)
//
// [SHCreateShellItemArrayFromIDLists]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateshellitemarrayfromidlists?redirectedfrom=MSDN
func SHCreateShellItemArrayFromIDLists(
	releaser *OleReleaser,
	pidls ...*ITEMIDLIST,
) (*IShellItemArray, error) {
	var ppvtQueried **_IUnknownVt

	pidlObjs := make([]ITEMIDLIST, 0, len(pidls))
	for _, pidl := range pidls {
		pidlObjs = append(pidlObjs, *pidl)
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHCreateShellItemArrayFromIDLists, "SHCreateShellItemArrayFromIDLists"),
		uintptr(uint32(len(pidls))),
		uintptr(unsafe.Pointer(&pidlObjs[0])),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItemArray{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHCreateShellItemArrayFromIDLists *syscall.Proc

// [SHGetDesktopFolder] function.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	folder, _ := win.SHGetDesktopFolder(rel)
//
// [SHGetDesktopFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shgetdesktopfolder
func SHGetDesktopFolder(releaser *OleReleaser) (*IShellFolder, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHGetDesktopFolder, "SHGetDesktopFolder"),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellFolder{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHGetDesktopFolder *syscall.Proc

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
//	_ = win.SHGetKnownFolderItem(
//		rel,
//		co.FOLDERID_Desktop,
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
	releaser *OleReleaser,
	kfid co.FOLDERID,
	flags co.KF,
	hToken HANDLE, // HACCESSTOKEN
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidKfid := GuidFrom(kfid)
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHGetKnownFolderItem, "SHGetKnownFolderItem"),
		uintptr(unsafe.Pointer(&guidKfid)),
		uintptr(flags),
		uintptr(hToken),
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

var _SHGetKnownFolderItem *syscall.Proc

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
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
//	idl, _ := win.SHGetIDListFromObject(rel, &item.IUnknown)
//
// [SHGetIDListFromObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetidlistfromobject
func SHGetIDListFromObject(releaser *OleReleaser, obj *IUnknown) (*ITEMIDLIST, error) {
	var idl ITEMIDLIST
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHGetIDListFromObject, "SHGetIDListFromObject"),
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

var _SHGetIDListFromObject *syscall.Proc

// [SHGetPropertyStoreFromIDList] function.
//
// [SHGetPropertyStoreFromIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetpropertystorefromidlist
func SHGetPropertyStoreFromIDList(
	releaser *OleReleaser,
	pidl *ITEMIDLIST,
	flags co.GPS,
) (*IPropertyStore, error) {
	var ppvtQueried **_IUnknownVt
	guid := GuidFrom(co.IID_IPropertyStore)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHGetPropertyStoreFromIDList, "SHGetPropertyStoreFromIDList"),
		uintptr(*pidl),
		uintptr(flags),
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IPropertyStore{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHGetPropertyStoreFromIDList *syscall.Proc

// [SHGetPropertyStoreFromParsingName] function.
//
// [SHGetPropertyStoreFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetpropertystorefromparsingname
func SHGetPropertyStoreFromParsingName(
	releaser *OleReleaser,
	folderOrFilePath string,
	bindCtx *IBindCtx,
	flags co.GPS,
) (*IPropertyStore, error) {
	var ppvtQueried **_IUnknownVt
	guid := GuidFrom(co.IID_IPropertyStore)

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pFolderOrFilePath := wbuf.PtrAllowEmpty(folderOrFilePath)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHGetPropertyStoreFromParsingName, "SHGetPropertyStoreFromParsingName"),
		uintptr(pFolderOrFilePath),
		uintptr(ppvtOrNil(bindCtx)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IPropertyStore{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHGetPropertyStoreFromParsingName *syscall.Proc

// [Shell_NotifyIcon] function.
//
// [Shell_NotifyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func Shell_NotifyIcon(message co.NIM, data *NOTIFYICONDATA) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_Shell_NotifyIconW, "Shell_NotifyIconW"),
		uintptr(message),
		uintptr(unsafe.Pointer(data)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

var _Shell_NotifyIconW *syscall.Proc

// [Shell_NotifyIconGetRect] function.
//
// [Shell_NotifyIconGetRect]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyicongetrect
func Shell_NotifyIconGetRect(identifier *NOTIFYICONIDENTIFIER) (RECT, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_Shell_NotifyIconGetRect, "Shell_NotifyIconGetRect"),
		uintptr(unsafe.Pointer(identifier)),
		uintptr(unsafe.Pointer(&rc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return RECT{}, hr
	}
	return rc, nil
}

var _Shell_NotifyIconGetRect *syscall.Proc

// [SHGetFileInfo] function.
//
// ⚠️ You must defer [HICON.DestroyIcon] on the HIcon member of the returned
// struct.
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(path string, fileAttrs co.FILE_ATTRIBUTE, flags co.SHGFI) (SHFILEINFO, error) {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pPath := wbuf.PtrAllowEmpty(path)

	var sfi SHFILEINFO

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHELL32, &_SHGetFileInfoW, "SHGetFileInfoW"),
		uintptr(pPath),
		uintptr(fileAttrs),
		uintptr(unsafe.Pointer(&sfi)),
		unsafe.Sizeof(sfi),
		uintptr(flags))

	if (flags&co.SHGFI_EXETYPE) == 0 || (flags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			return SHFILEINFO{}, co.ERROR_INVALID_PARAMETER
		}
	}
	if (flags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			return SHFILEINFO{}, co.ERROR_INVALID_PARAMETER
		}
	}

	return sfi, nil
}

var _SHGetFileInfoW *syscall.Proc
