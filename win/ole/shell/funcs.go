//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
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
//	item, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\foo.txt")
//
//	idl, _ := shell.SHGetIDListFromObject(rel, item)
//
//	sameItem, _ := shell.SHCreateItemFromIDList[shell.IShellItem2](
//		rel, idl)
//
// [SHCreateItemFromIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromidlist
func SHCreateItemFromIDList[T any, P ole.ComCtor[T]](
	releaser *ole.Releaser,
	pidl *ITEMIDLIST,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	guidIid := win.GuidFrom(pObj.IID())

	ret, _, _ := syscall.SyscallN(_SHCreateItemFromIDList.Addr(),
		uintptr(*pidl), uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj.Set(ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
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
//	ish, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\foo.txt")
//
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func SHCreateItemFromParsingName[T any, P ole.ComCtor[T]](
	releaser *ole.Releaser,
	folderOrFilePath string,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	riidGuid := win.GuidFrom(pObj.IID())

	folderOrFilePath16 := wstr.NewBufWith[wstr.Stack20](folderOrFilePath, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(_SHCreateItemFromParsingName.Addr(),
		uintptr(folderOrFilePath16.UnsafePtr()),
		0, uintptr(unsafe.Pointer(&riidGuid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj.Set(ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _SHCreateItemFromParsingName = dll.Shell32.NewProc("SHCreateItemFromParsingName")

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
		pObj := ole.ComObj[IShellFolder](ppvtQueried)
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
//	desktop, _ := shell.SHGetKnownFolderItem[shell.IShellItem](
//		rel, co.FOLDERID_Desktop, co.KF_FLAG_DEFAULT, win.HANDLE(0))
//
//	path, _ := desktop.GetDisplayName(co.SIGDN_FILESYSPATH)
//	println(path)
//
// [SHGetKnownFolderItem]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shgetknownfolderitem
func SHGetKnownFolderItem[T any, P ole.ComCtor[T]](
	releaser *ole.Releaser,
	kfid co.FOLDERID,
	flags co.KF_FLAG,
	hToken win.HANDLE, // HACCESSTOKEN
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **ole.IUnknownVt
	kfidGuid := win.GuidFrom(kfid)
	riidGuid := win.GuidFrom(pObj.IID())

	ret, _, _ := syscall.SyscallN(_SHGetKnownFolderItem.Addr(),
		uintptr(unsafe.Pointer(&kfidGuid)),
		uintptr(flags), uintptr(hToken),
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

var _SHGetKnownFolderItem = dll.Shell32.NewProc("SHGetKnownFolderItem")

// [SHGetIDListFromObject] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	item, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\foo.txt")
//
//	idl, _ := shell.SHGetIDListFromObject(rel, item)
//
// [SHGetIDListFromObject]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shgetidlistfromobject
func SHGetIDListFromObject(releaser *ole.Releaser, obj ole.ComPtr) (*ITEMIDLIST, error) {
	var idl ITEMIDLIST
	ret, _, _ := syscall.SyscallN(_SHGetIDListFromObject.Addr(),
		uintptr(unsafe.Pointer(obj.Ppvt())), uintptr(unsafe.Pointer(&idl)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pIdl := &idl
		releaser.Add(pIdl)
		return pIdl, nil
	} else {
		return nil, hr
	}
}

var _SHGetIDListFromObject = dll.Shell32.NewProc("SHGetIDListFromObject")
