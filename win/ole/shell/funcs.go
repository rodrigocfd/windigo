//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [SHCreateItemFromIDList] function.
//
// Return type is typically [IShellItem] of [IShellItem2].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var item *shell.IShellItem
//	shell.SHCreateItemFromParsingName(
//		rel, "C:\\Temp\\foo.txt", &item)
//
//	idl, _ := shell.SHGetIDListFromObject(rel, item)
//
//	var sameItem *shell.IShellItem2
//	shell.SHCreateItemFromIDList(rel, idl, &sameItem)
//
// [SHCreateItemFromIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromidlist
func SHCreateItemFromIDList(releaser *ole.Releaser, pidl *ITEMIDLIST, ppOut interface{}) error {
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(_SHCreateItemFromIDList.Addr(),
		uintptr(*pidl),
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

var _SHCreateItemFromIDList = dll.Shell32.NewProc("SHCreateItemFromIDList")

// [SHCreateItemFromParsingName] function.
//
// Return type is typically [IShellItem] of [IShellItem2].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var item *shell.IShellItem
//	shell.SHCreateItemFromParsingName(
//		rel,
//		"C:\\Temp\\foo.txt",
//		&item,
//	)
//
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func SHCreateItemFromParsingName(
	releaser *ole.Releaser,
	folderOrFilePath string,
	ppOut interface{},
) error {
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	folderOrFilePath16 := wstr.NewBufWith[wstr.Stack20](folderOrFilePath, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(_SHCreateItemFromParsingName.Addr(),
		uintptr(folderOrFilePath16.UnsafePtr()),
		0,
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

var _SHCreateItemFromParsingName = dll.Shell32.NewProc("SHCreateItemFromParsingName")

// [SHCreateItemFromRelativeName] function.
//
// [SHCreateItemFromRelativeName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromrelativename
func SHCreateItemFromRelativeName(
	releaser *ole.Releaser,
	parent *IShellItem,
	name string,
	bindCtx *ole.IBindCtx,
	ppOut interface{},
) error {
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pOut.IID())

	name16 := wstr.NewBufWith[wstr.Stack20](name, wstr.ALLOW_EMPTY)

	var pBindCtx **ole.IUnknownVt
	if bindCtx != nil {
		pBindCtx = bindCtx.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(_SHCreateItemFromRelativeName.Addr(),
		uintptr(unsafe.Pointer(parent.Ppvt())),
		uintptr(name16.UnsafePtr()),
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

var _SHCreateItemFromRelativeName = dll.Shell32.NewProc("SHCreateItemFromRelativeName")

// [SHCreateShellItemArray] function.
//
// [SHCreateShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateshellitemarray
func SHCreateShellItemArray(
	releaser *ole.Releaser,
	pidlParent *ITEMIDLIST,
	parent *IShellFolder,
	pidlChildren []*ITEMIDLIST,
) (*IShellItemArray, error) {
	var ppvtQueried **ole.IUnknownVt

	var pidlParentObj ITEMIDLIST
	if pidlParent != nil {
		pidlParentObj = *pidlParent
	}

	var pParent **ole.IUnknownVt
	if parent != nil {
		pParent = parent.Ppvt()
	}

	var pidlChildrenObjsPtr *ITEMIDLIST
	if pidlChildren != nil {
		pidlChildrenObjs := make([]ITEMIDLIST, 0, len(pidlChildren))
		for _, pidl := range pidlChildren {
			pidlChildrenObjs = append(pidlChildrenObjs, *pidl)
		}
		pidlChildrenObjsPtr = &pidlChildrenObjs[0]
	}

	ret, _, _ := syscall.SyscallN(_SHCreateShellItemArray.Addr(),
		uintptr(pidlParentObj),
		uintptr(unsafe.Pointer(pParent)),
		uintptr(len(pidlChildren)),
		uintptr(unsafe.Pointer(pidlChildrenObjsPtr)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj *IShellItemArray
		utl.ComCreateObj(&pObj, unsafe.Pointer(ppvtQueried))
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHCreateShellItemArray = dll.Shell32.NewProc("SHCreateShellItemArray")

// [SHCreateShellItemArrayFromIDLists] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var item1, item2 *shell.IShellItem
//	shell.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item1)
//	shell.SHCreateItemFromParsingName(rel, "C:\\Temp\\bar.txt", &item2)
//
//	pidl1, _ := shell.SHGetIDListFromObject(rel, &item1.IUnknown)
//	pidl2, _ := shell.SHGetIDListFromObject(rel, &item2.IUnknown)
//
//	arr, _ := shell.SHCreateShellItemArrayFromIDLists(rel, pidl1, pidl2)
//
// [SHCreateShellItemArrayFromIDLists]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateshellitemarrayfromidlists?redirectedfrom=MSDN
func SHCreateShellItemArrayFromIDLists(
	releaser *ole.Releaser,
	pidls ...*ITEMIDLIST,
) (*IShellItemArray, error) {
	var ppvtQueried **ole.IUnknownVt

	pidlObjs := make([]ITEMIDLIST, 0, len(pidls))
	for _, pidl := range pidls {
		pidlObjs = append(pidlObjs, *pidl)
	}

	ret, _, _ := syscall.SyscallN(_SHCreateShellItemArrayFromIDLists.Addr(),
		uintptr(uint32(len(pidls))),
		uintptr(unsafe.Pointer(&pidlObjs[0])),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj *IShellItemArray
		utl.ComCreateObj(&pObj, unsafe.Pointer(ppvtQueried))
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHCreateShellItemArrayFromIDLists = dll.Shell32.NewProc("SHCreateShellItemArrayFromIDLists")

// [SHGetDesktopFolder] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	folder, _ := shell.SHGetDesktopFolder(rel)
//
// [SHGetDesktopFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shgetdesktopfolder
func SHGetDesktopFolder(releaser *ole.Releaser) (*IShellFolder, error) {
	var ppvtQueried **ole.IUnknownVt
	ret, _, _ := syscall.SyscallN(_SHGetDesktopFolder.Addr(),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj *IShellFolder
		utl.ComCreateObj(&pObj, unsafe.Pointer(ppvtQueried))
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHGetDesktopFolder = dll.Shell32.NewProc("SHGetDesktopFolder")

// [SHGetKnownFolderItem] function.
//
// Return type is typically [IShellItem] of [IShellItem2].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var desktop *shell.IShellItem
//	shell.SHGetKnownFolderItem(
//		rel,
//		co.FOLDERID_Desktop,
//		co.KF_FLAG_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	path, _ := desktop.GetDisplayName(co.SIGDN_FILESYSPATH)
//	println(path)
//
// [SHGetKnownFolderItem]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shgetknownfolderitem
func SHGetKnownFolderItem(
	releaser *ole.Releaser,
	kfid co.FOLDERID,
	flags co.KF_FLAG,
	hToken win.HANDLE, // HACCESSTOKEN
	ppOut interface{},
) error {
	pOut := utl.ComValidateObj(ppOut).(ole.ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **ole.IUnknownVt
	guidKfid := win.GuidFrom(kfid)
	guidIid := win.GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(_SHGetKnownFolderItem.Addr(),
		uintptr(unsafe.Pointer(&guidKfid)),
		uintptr(flags),
		uintptr(hToken),
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

var _SHGetKnownFolderItem = dll.Shell32.NewProc("SHGetKnownFolderItem")

// [SHGetIDListFromObject] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var item *shell.IShellItem
//	shell.SHCreateItemFromParsingName(
//		rel, "C:\\Temp\\foo.txt", &item)
//
//	idl, _ := shell.SHGetIDListFromObject(rel, &item.IUnknown)
//
// [SHGetIDListFromObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetidlistfromobject
func SHGetIDListFromObject(releaser *ole.Releaser, obj *ole.IUnknown) (*ITEMIDLIST, error) {
	var idl ITEMIDLIST
	ret, _, _ := syscall.SyscallN(_SHGetIDListFromObject.Addr(),
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

var _SHGetIDListFromObject = dll.Shell32.NewProc("SHGetIDListFromObject")
