//go:build windows

package win

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IEnumIDList] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	_ = item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(rel, win.HWND(0),
//		co.SHCONTF_FOLDERS|co.SHCONTF_NONFOLDERS|co.SHCONTF_INCLUDEHIDDEN)
//
// [IEnumIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumidlist
type IEnumIDList struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumIDList) IID() co.IID {
	return co.IID_IEnumIDList
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-clone
func (me *IEnumIDList) Clone(releaser *OleReleaser) (*IEnumIDList, error) {
	return com_callBuildObj[*IEnumIDList](me, releaser,
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Clone)
}

// Returns all [ITEMIDLIST] values by calling [IEnumIDList.Next].
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
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	_ = item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(
//		rel,
//		win.HWND(0),
//		co.SHCONTF_FOLDERS|co.SHCONTF_NONFOLDERS|co.SHCONTF_INCLUDEHIDDEN,
//	)
//
//	pidls, _ := pidlList.Enum(rel)
//	for _, pidl := range pidls {
//		var child *win.IShellItem
//		_ = win.SHCreateItemFromIDList(rel, pidl, &child)
//		name, _ := child.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(name)
//	}
func (me *IEnumIDList) Enum(releaser *OleReleaser) ([]*ITEMIDLIST, error) {
	items := make([]*ITEMIDLIST, 0)
	var item *ITEMIDLIST
	var hr error

	for {
		item, hr = me.Next(releaser)
		if hr != nil { // actual error
			return nil, hr
		} else if item == nil { // no more items to fetch
			return items, nil
		} else { // item fetched
			items = append(items, item)
		}
	}
}

// [Next] method.
//
// If there are no more items, nil is returned.
//
// This is a low-level method, prefer using [IEnumIDList.Enum].
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-next
func (me *IEnumIDList) Next(releaser *OleReleaser) (*ITEMIDLIST, error) {
	var idlFetched ITEMIDLIST
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
		1,
		uintptr(unsafe.Pointer(&idlFetched)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pIdlFetched := &idlFetched
		releaser.Add(pIdlFetched)
		return pIdlFetched, nil
	} else if hr == co.HRESULT_S_FALSE {
		return nil, nil
	} else {
		return nil, hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-reset
func (me *IEnumIDList) Reset() error {
	return com_callNoParm(me,
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Reset)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-skip
func (me *IEnumIDList) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}

type _IEnumIDListVt struct {
	_IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IEnumShellItems] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var enumItems *win.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, co.BHID_EnumItems, &enumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(path)
//	}
//
// [IEnumShellItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumshellitems
type IEnumShellItems struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumShellItems) IID() co.IID {
	return co.IID_IEnumShellItems
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-clone
func (me *IEnumShellItems) Clone(releaser *OleReleaser) (*IEnumShellItems, error) {
	return com_callBuildObj[*IEnumShellItems](me, releaser,
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Clone)
}

// Returns all [IShellItem] values by calling [IEnumShellItems.Next].
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
//	var enumItems *win.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, co.BHID_EnumItems, &enumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(path)
//	}
func (me *IEnumShellItems) Enum(releaser *OleReleaser) ([]*IShellItem, error) {
	items := make([]*IShellItem, 0)
	var item *IShellItem
	var hr error

	for {
		item, hr = me.Next(releaser)
		if hr != nil { // actual error
			return nil, hr
		} else if item == nil { // no more items to fetch
			return items, nil
		} else { // item fetched
			items = append(items, item)
		}
	}
}

// [Next] method.
//
// If there are no more items, nil is returned.
//
// This is a low-level method, prefer using [IEnumShellItems.Enum].
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-next
func (me *IEnumShellItems) Next(releaser *OleReleaser) (*IShellItem, error) {
	var ppvtQueried **_IUnknownVt
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
		1,
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj *IShellItem
		com_buildObj(&pObj, ppvtQueried, releaser)
		return pObj, nil
	} else if hr == co.HRESULT_S_FALSE {
		return nil, nil
	} else {
		return nil, hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-reset
func (me *IEnumShellItems) Reset() error {
	return com_callNoParm(me,
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Reset)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-skip
func (me *IEnumShellItems) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}

type _IEnumShellItemsVt struct {
	_IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IFileDialog] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IFileDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
type IFileDialog struct{ IModalWindow }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileDialog) IID() co.IID {
	return co.IID_IFileDialog
}

// [AddPlace] method.
//
// [AddPlace]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-addplace
func (me *IFileDialog) AddPlace(si *IShellItem, fdap co.FDAP) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).AddPlace,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())),
		uintptr(fdap))
	return utl.HresultToError(ret)
}

// [Advise] method.
//
// Paired with [IFileDialog.Unadvise].
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-advise
func (me *IFileDialog) Advise(events *IFileDialogEvents) (cookie uint32, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).Advise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(events.Ppvt())),
		uintptr(unsafe.Pointer(&cookie)))
	if hr = co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return cookie, nil
}

// [ClearClientData] method.
//
// [ClearClientData]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-clearclientdata
func (me *IFileDialog) ClearClientData() error {
	return com_callNoParm(me,
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).ClearClientData)
}

// [Close] method.
//
// [Close]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-close
func (me *IFileDialog) Close(hr co.ERROR) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).Close,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hr))
	return utl.HresultToError(ret)
}

// [GetCurrentSelection] method.
//
// [GetCurrentSelection]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getcurrentselection
func (me *IFileDialog) GetCurrentSelection(releaser *OleReleaser) (*IShellItem, error) {
	return com_callBuildObj[*IShellItem](me, releaser,
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetCurrentSelection)
}

// [GetFileName] method.
//
// [GetFileName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfilename
func (me *IFileDialog) GetFileName() (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetFileName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&pv)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(pv).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetFileTypeIndex] method.
//
// [GetFileTypeIndex]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfiletypeindex
func (me *IFileDialog) GetFileTypeIndex() (int, error) {
	var idx uint32
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetFileTypeIndex,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&idx)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(idx), nil
	} else {
		return 0, hr
	}
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfolder
func (me *IFileDialog) GetFolder(releaser *OleReleaser) (*IShellItem, error) {
	return com_callBuildObj[*IShellItem](me, releaser,
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetFolder)
}

// [GetOptions] method.
//
// [GetOptions]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getoptions
func (me *IFileDialog) GetOptions() (co.FOS, error) {
	var fos co.FOS
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&fos)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return fos, nil
	} else {
		return co.FOS(0), hr
	}
}

// [GetResult] method.
//
// Returns the selected item after user confirmation, for single-selection
// dialogs – those without [co.FOS_ALLOWMULTISELECT] option.
//
// For multi-selection dialogs, use [IFileOpenDialog.GetResults].
//
// [GetResult]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getresult
func (me *IFileDialog) GetResult(releaser *OleReleaser) (*IShellItem, error) {
	return com_callBuildObj[*IShellItem](me, releaser,
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetResult)
}

// [SetClientGuid] method.
//
// [SetClientGuid]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setclientguid
func (me *IFileDialog) SetClientGuid(guid *GUID) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetClientGuid,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(guid)))
	return utl.HresultToError(ret)
}

// [SetDefaultExtension] method.
//
// [SetDefaultExtension]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setdefaultextension
func (me *IFileDialog) SetDefaultExtension(defaultExt string) error {
	var wDefaultExt wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetDefaultExtension,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wDefaultExt.EmptyIsNil(defaultExt)))
	return utl.HresultToError(ret)
}

// [SetDefaultFolder] method.
//
// [SetDefaultFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setdefaultfolder
func (me *IFileDialog) SetDefaultFolder(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetDefaultFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetFileName] method.
//
// [SetFileName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilename
func (me *IFileDialog) SetFileName(name string) error {
	var wName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wName.EmptyIsNil(name)))
	return utl.HresultToError(ret)
}

// [SetFileNameLabel] method.
//
// [SetFileNameLabel]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilenamelabel
func (me *IFileDialog) SetFileNameLabel(label string) error {
	var wLabel wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileNameLabel,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wLabel.EmptyIsNil(label)))
	return utl.HresultToError(ret)
}

// [SetFileTypeIndex] method.
//
// The index is one-based.
//
// [SetFileTypeIndex]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypeindex
func (me *IFileDialog) SetFileTypeIndex(index int) error {
	if index < 1 {
		panic(fmt.Sprintf("Index is one-based: %d.", index))
	}

	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileTypeIndex,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)))
	return utl.HresultToError(ret)
}

// [SetFileTypes] method.
//
// Example:
//
//	var fd win.IFileDialog // initialized somewhere
//
//	_ = fd.SetFileTypes([]win.COMDLG_FILTERSPEC{
//		{Name: "MP3 audio files", Spec: "*.mp3"},
//		{Name: "All files", Spec: "*.*"},
//	})
//
// [SetFileTypes]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypes
func (me *IFileDialog) SetFileTypes(filterSpec []COMDLG_FILTERSPEC) error {
	nativeFilters := make([]_COMDLG_FILTERSPEC, 0, len(filterSpec))
	for _, fs := range filterSpec {
		nativeFilters = append(nativeFilters, _COMDLG_FILTERSPEC{
			PszName: (*uint16)(wstr.EncodeToPtr(fs.Name)),
			PszSpec: (*uint16)(wstr.EncodeToPtr(fs.Spec)),
		})
	}

	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileTypes,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(len(filterSpec))),
		uintptr(unsafe.Pointer(unsafe.SliceData(nativeFilters))))
	return utl.HresultToError(ret)
}

// [SetFilter] method.
//
// [SetFilter]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilter
func (me *IFileDialog) SetFilter(filter *IShellItemFilter) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFilter,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(filter.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetFolder] method.
//
// [SetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfolder
func (me *IFileDialog) SetFolder(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetOkButtonLabel] method.
//
// [SetOkButtonLabel]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setokbuttonlabel
func (me *IFileDialog) SetOkButtonLabel(text string) error {
	var wText wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetOkButtonLabel,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wText.EmptyIsNil(text)))
	return utl.HresultToError(ret)
}

// [SetOptions] method.
//
// Example:
//
//	var fd win.IFileDialog // initialized somewhere
//
//	curOpts, _ := fd.GetOptions()
//	_ = fd.SetOptions(curOpts |
//		co.FOS_FORCEFILESYSTEM |
//		co.FOS_FILEMUSTEXIST |
//		co.FOS_ALLOWMULTISELECT,
//	)
//
// [SetOptions]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setoptions
func (me *IFileDialog) SetOptions(fos co.FOS) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(fos))
	return utl.HresultToError(ret)
}

// [SetTitle] method.
//
// [SetTitle]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-settitle
func (me *IFileDialog) SetTitle(title string) error {
	var wTitle wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetTitle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wTitle.EmptyIsNil(title)))
	return utl.HresultToError(ret)
}

// [Unadvise] method.
//
// Paired with [IFileDialog.Advise].
//
// [Unadvise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-unadvise
func (me *IFileDialog) Unadvise(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).Unadvise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(cookie))
	return utl.HresultToError(ret)
}

type _IFileDialogVt struct {
	_IModalWindowVt
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IFileOpenDialog] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var fod *win.IFileOpenDialog
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOpenDialog,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&fod,
//	)
//
//	defOpts, _ := fod.GetOptions()
//	_ = fod.SetOptions(defOpts |
//		co.FOS_FORCEFILESYSTEM |
//		co.FOS_FILEMUSTEXIST,
//	)
//
//	_ = fod.SetFileTypes([]win.COMDLG_FILTERSPEC{
//		{Name: "Text files", Spec: "*.txt"},
//		{Name: "All files", Spec: "*.*"},
//	})
//	_ = fod.SetFileTypeIndex(1)
//
//	if ok, _ := fod.Show(hWnd); ok {
//		item, _ := fod.GetResult(rel)
//		fileName, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(fileName)
//	}
//
// [IFileOpenDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog struct{ IFileDialog }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOpenDialog) IID() co.IID {
	return co.IID_IFileOpenDialog
}

// [GetResults] method.
//
// Returns the selected items after user confirmation, for multi-selection
// dialogs – those with [co.FOS_ALLOWMULTISELECT] option.
//
// For single-selection dialogs, use [IFileDialog.GetResult].
//
// [GetResults]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults(releaser *OleReleaser) (*IShellItemArray, error) {
	return com_callBuildObj[*IShellItemArray](me, releaser,
		(*_IFileOpenDialogVt)(unsafe.Pointer(*me.Ppvt())).GetResults)
}

// [GetSelectedItems] method.
//
// [GetSelectedItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems(releaser *OleReleaser) (*IShellItemArray, error) {
	return com_callBuildObj[*IShellItemArray](me, releaser,
		(*_IFileOpenDialogVt)(unsafe.Pointer(*me.Ppvt())).GetSelectedItems)
}

type _IFileOpenDialogVt struct {
	_IFileDialogVt
	GetResults       uintptr
	GetSelectedItems uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IFileOperation] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var op *win.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL,
//		&op,
//	)
//
// [IFileOperation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileoperation
type IFileOperation struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOperation) IID() co.IID {
	return co.IID_IFileOperation
}

// [Advise] method.
//
// Paired with [IFileOperation.Unadvise].
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-advise
func (me *IFileOperation) Advise(fops *IFileOperationProgressSink) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).Advise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(fops.Ppvt())),
		uintptr(unsafe.Pointer(&cookie)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return cookie, nil
	} else {
		return 0, hr
	}
}

// [ApplyPropertiesToItem] method.
//
// [ApplyPropertiesToItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-applypropertiestoitem
func (me *IFileOperation) ApplyPropertiesToItem(item *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).ApplyPropertiesToItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())))
	return utl.HresultToError(ret)
}

// [CopyItem] method.
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
//	var op *win.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL,
//		&op,
//	)
//
//	var file, dest *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &file)
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\mydir", &dest)
//
//	_ = op.CopyItem(file, dest, "new name.txt", nil)
//	_ = op.PerformOperations()
//
// [CopyItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-copyitem
func (me *IFileOperation) CopyItem(
	item, destFolder *IShellItem,
	copyName string,
	fops *IFileOperationProgressSink,
) error {
	var wCopyName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).CopyItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(wCopyName.EmptyIsNil(copyName)),
		uintptr(com_ppvtOrNil(fops)))
	return utl.HresultToError(ret)
}

// [DeleteItem] method.
//
// [DeleteItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-deleteitem
func (me *IFileOperation) DeleteItem(item *IShellItem, fops *IFileOperationProgressSink) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).DeleteItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(com_ppvtOrNil(fops)))
	return utl.HresultToError(ret)
}

// [GetAnyOperationsAborted] method.
//
// [GetAnyOperationsAborted]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-getanyoperationsaborted
func (me *IFileOperation) GetAnyOperationsAborted() (bool, error) {
	var bVal BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).GetAnyOperationsAborted,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bVal)))
	return utl.HresultToBoolError(int32(bVal), ret)
}

// [MoveItem] method.
//
// [MoveItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-moveitem
func (me *IFileOperation) MoveItem(
	item, destFolder *IShellItem,
	newName string,
	fops *IFileOperationProgressSink,
) error {
	var wNewName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).MoveItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(wNewName.AllowEmpty(newName)),
		uintptr(com_ppvtOrNil(fops)))
	return utl.HresultToError(ret)
}

// [NewItem] method.
//
// [NewItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-newitem
func (me *IFileOperation) NewItem(
	destFolder *IShellItem,
	fileAtt co.FILE_ATTRIBUTE,
	name, templateName string,
	fops *IFileOperationProgressSink,
) error {
	var wName, wTemplateName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).NewItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(fileAtt),
		uintptr(wName.AllowEmpty(name)),
		uintptr(wTemplateName.EmptyIsNil(templateName)),
		uintptr(com_ppvtOrNil(fops)))
	return utl.HresultToError(ret)
}

// [PerformOperations] method.
//
// [PerformOperations]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-performoperations
func (me *IFileOperation) PerformOperations() error {
	return com_callNoParm(me,
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).PerformOperations)
}

// [RenameItem] method.
//
// [RenameItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-renameitem
func (me *IFileOperation) RenameItem(
	item *IShellItem,
	newName string,
	fops *IFileOperationProgressSink,
) error {
	var wNewName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).RenameItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(wNewName.EmptyIsNil(newName)),
		uintptr(com_ppvtOrNil(fops)))
	return utl.HresultToError(ret)
}

// [SetOperationFlags] method.
//
// [SetOperationFlags]:
func (me *IFileOperation) SetOperationFlags(flags co.FOF) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).SetOperationFlags,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [SetOwnerWindow] method.
//
// [SetOwnerWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-setownerwindow
func (me *IFileOperation) SetOwnerWindow(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).SetOwnerWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [SetProgressMessage] method.
//
// [SetProgressMessage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-setprogressmessage
func (me *IFileOperation) SetProgressMessage(message string) error {
	var wMessage wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).SetProgressMessage,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wMessage.EmptyIsNil(message)))
	return utl.HresultToError(ret)
}

// [Unadvise] method.
//
// Paired with [IFileOperation.Advise].
//
// [Unadvise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-unadvise
func (me *IFileOperation) Unadvise(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).Unadvise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(cookie))
	return utl.HresultToError(ret)
}

type _IFileOperationVt struct {
	_IUnknownVt
	Advise                  uintptr
	Unadvise                uintptr
	SetOperationFlags       uintptr
	SetProgressMessage      uintptr
	SetProgressDialog       uintptr
	SetProperties           uintptr
	SetOwnerWindow          uintptr
	ApplyPropertiesToItem   uintptr
	ApplyPropertiesToItems  uintptr
	RenameItem              uintptr
	RenameItems             uintptr
	MoveItem                uintptr
	MoveItems               uintptr
	CopyItem                uintptr
	CopyItems               uintptr
	DeleteItem              uintptr
	DeleteItems             uintptr
	NewItem                 uintptr
	PerformOperations       uintptr
	GetAnyOperationsAborted uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IFileSaveDialog] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var fsd *win.IFileSaveDialog
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileSaveDialog,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&fsd,
//	)
//
//	_ = fsd.SetFileTypes([]win.COMDLG_FILTERSPEC{
//		{Name: "Text files", Spec: "*.txt"},
//		{Name: "All files", Spec: "*.*"},
//	})
//	_ = fsd.SetFileTypeIndex(1)
//
//	_ = fsd.SetFileName("default-file-name.txt")
//
//	if ok, _ := fsd.Show(hWnd); ok {
//		item, _ := fsd.GetResult(rel)
//		txtPath, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(txtPath)
//	}
//
// [IFileSaveDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog struct{ IFileDialog }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileSaveDialog) IID() co.IID {
	return co.IID_IFileSaveDialog
}

// [ApplyProperties] method.
//
// [ApplyProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-applyproperties
func (me *IFileSaveDialog) ApplyProperties(
	item *IShellItem,
	store *IPropertyStore,
	hwnd HWND,
	sink *IFileOperationProgressSink,
) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).ApplyProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(unsafe.Pointer(store.Ppvt())),
		uintptr(hwnd),
		uintptr(com_ppvtOrNil(sink)))
	return utl.HresultToError(ret)
}

// [GetProperties] method.
//
// [GetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-getproperties
func (me *IFileSaveDialog) GetProperties(releaser *OleReleaser) (*IPropertyStore, error) {
	return com_callBuildObj[*IPropertyStore](me, releaser,
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).GetProperties)
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setproperties
func (me *IFileSaveDialog) SetProperties(store *IPropertyStore) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).SetProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(store.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetSaveAsItem] method.
//
// [SetSaveAsItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(item *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).SetSaveAsItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())))
	return utl.HresultToError(ret)
}

type _IFileSaveDialogVt struct {
	_IFileDialogVt
	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IModalWindow] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IModalWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IModalWindow) IID() co.IID {
	return co.IID_IModalWindow
}

// [Show] method.
//
// Returns false if user cancelled.
//
// [Show]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner HWND) (bool, error) {
	ret, _, _ := syscall.SyscallN(
		(*_IModalWindowVt)(unsafe.Pointer(*me.Ppvt())).Show,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndOwner))

	if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
		return true, nil
	} else if wErr == co.ERROR_CANCELLED {
		return false, nil
	} else {
		return false, wErr.ToHresult()
	}
}

type _IModalWindowVt struct {
	_IUnknownVt
	Show uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IOleWindow] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IOleWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-iolewindow
type IOleWindow struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IOleWindow) IID() co.IID {
	return co.IID_IOleWindow
}

// [ContextSensitiveHelp] method.
//
// [ContextSensitiveHelp]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-iolewindow-contextsensitivehelp
func (me *IOleWindow) ContextSensitiveHelp() (bool, error) {
	var bVal BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IOleWindowVt)(unsafe.Pointer(*me.Ppvt())).ContextSensitiveHelp,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bVal)))
	return utl.HresultToBoolError(int32(bVal), ret)
}

// [GetWindow] method.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-iolewindow-getwindow
func (me *IOleWindow) GetWindow() (HWND, error) {
	var hWnd HWND
	ret, _, _ := syscall.SyscallN(
		(*_IOleWindowVt)(unsafe.Pointer(*me.Ppvt())).GetWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hWnd)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hWnd, nil
	} else {
		return HWND(0), hr
	}
}

type _IOleWindowVt struct {
	_IUnknownVt
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IShellFolder] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	_ = item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
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
	iid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(iid)

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToObject,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(*pidl),
		uintptr(com_ppvtOrNil(bindCtx)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
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
	iid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(iid)

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).BindToStorage,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(*pidl),
		uintptr(com_ppvtOrNil(bindCtx)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
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
	iid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt
	guidIid := GuidFrom(iid)

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).CreateViewObject,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
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
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	_ = item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(
//		rel,
//		win.HWND(0),
//		co.SHCONTF_FOLDERS|co.SHCONTF_NONFOLDERS|co.SHCONTF_INCLUDEHIDDEN,
//	)
//
// [EnumObjects]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-enumobjects
func (me *IShellFolder) EnumObjects(
	releaser *OleReleaser,
	hWnd HWND,
	flags co.SHCONTF,
) (*IEnumIDList, error) {
	return com_callBuildObj[*IEnumIDList](me, releaser,
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).EnumObjects)
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
	var wDisplayName wstr.BufEncoder
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
		uintptr(com_ppvtOrNil(bindCtx)),
		uintptr(wDisplayName.AllowEmpty(displayName)),
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
	var wName wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		(*_IShellFolderVt)(unsafe.Pointer(*me.Ppvt())).SetNameOf,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(*pidl)),
		uintptr(wName.AllowEmpty(name)),
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IShellItem] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Usually created with [SHCreateItemFromParsingName].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem) IID() co.IID {
	return co.IID_IShellItem
}

// [BindToHandler] method.
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
//	var enumItems *win.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, co.BHID_EnumItems, &enumItems)
//
// [BindToHandler]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-bindtohandler
func (me *IShellItem) BindToHandler(
	releaser *OleReleaser,
	bindCtx *IBindCtx,
	bhid co.BHID,
	ppOut interface{},
) error {
	iid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt
	guidBhid := GuidFrom(bhid)
	guidIid := GuidFrom(iid)

	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).BindToHandler,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(com_ppvtOrNil(bindCtx)),
		uintptr(unsafe.Pointer(&guidBhid)),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
}

// [Compare] method.
//
// [Compare]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-compare
func (me *IShellItem) Compare(si *IShellItem, hint co.SICHINT) (bool, error) {
	var piOrder int32
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
// Example:
//
//	var shi win.IShellItem // initialized somewhere
//
//	fullPath, _ := shi.GetDisplayName(co.SIGDN_FILESYSPATH)
//
// [GetDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName co.SIGDN) (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetDisplayName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(sigdnName),
		uintptr(unsafe.Pointer(&pv)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(pv).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetParent] method.
//
// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
func (me *IShellItem) GetParent(releaser *OleReleaser) (*IShellItem, error) {
	return com_callBuildObj[*IShellItem](me, releaser,
		(*_IShellItemVt)(unsafe.Pointer(*me.Ppvt())).GetParent)
}

type _IShellItemVt struct {
	_IUnknownVt
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IShellItem2] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Usually created with [SHCreateItemFromParsingName].
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
//	var item *win.IShellItem2
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
// It can also be queried from an [IShellItem] object:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	var item2 *win.IShellItem2
//	_ = item.QueryInterface(rel, &item2)
//
// [IShellItem2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem2
type IShellItem2 struct{ IShellItem }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem2) IID() co.IID {
	return co.IID_IShellItem2
}

// [GetBool] method.
//
// [GetBool]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getbool
func (me *IShellItem2) GetBool(pkey co.PKEY) (bool, error) {
	guidPkey := PropertykeyFrom(pkey)
	var bVal BOOL

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetBool,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&bVal)))
	return utl.HresultToBoolError(int32(bVal), ret)
}

// [GetCLSID] method.
//
// [GetCLSID]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getclsid
func (me *IShellItem2) GetCLSID(pkey co.PKEY) (co.CLSID, error) {
	guidPkey := PropertykeyFrom(pkey)
	var guidClsid GUID

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetCLSID,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&guidClsid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.CLSID(guidClsid.String()), nil
	} else {
		return co.CLSID(""), hr
	}
}

// [GetFileTime] method.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getfiletime
func (me *IShellItem2) GetFileTime(pkey co.PKEY) (time.Time, error) {
	guidPkey := PropertykeyFrom(pkey)
	var ft FILETIME

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetFileTime,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&ft)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ft.ToTime(), nil
	} else {
		return time.Time{}, hr
	}
}

// [GetInt32] method.
//
// [GetInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getint32
func (me *IShellItem2) GetInt32(pkey co.PKEY) (int32, error) {
	guidPkey := PropertykeyFrom(pkey)
	var i int32

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetInt32,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&i)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return i, nil
	} else {
		return 0, hr
	}
}

// [GetPropertyStore] method.
//
// [GetPropertyStore]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getpropertystore
func (me *IShellItem2) GetPropertyStore(releaser *OleReleaser, flags co.GPS) (*IPropertyStore, error) {
	var ppvtQueried **_IUnknownVt
	guid := GuidFrom(co.IID_IPropertyStore)

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetPropertyStore,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(flags),
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IPropertyStore](ret, ppvtQueried, releaser)
}

// [GetString] method.
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
//	var item *win.IShellItem2
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	ty, _ := item.GetString(co.PKEY_ItemTypeText)
//	println(ty)
//
// [GetString]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getstring
func (me *IShellItem2) GetString(pkey co.PKEY) (string, error) {
	guidPkey := PropertykeyFrom(pkey)
	var psz uintptr

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetString,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&psz)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(psz).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(psz)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetUInt32] method.
//
// [GetUInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint32
func (me *IShellItem2) GetUInt32(pkey co.PKEY) (uint32, error) {
	guidPkey := PropertykeyFrom(pkey)
	var ui uint32

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetUInt32,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&ui)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ui, nil
	} else {
		return 0, hr
	}
}

// [GetUInt64] method.
//
// [GetUInt64]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint64
func (me *IShellItem2) GetUInt64(pkey co.PKEY) (uint64, error) {
	guidPkey := PropertykeyFrom(pkey)
	var ull uint64

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetUInt64,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&ull)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ull, nil
	} else {
		return 0, hr
	}
}

// [Update] method.
//
// [Update]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-update
func (me *IShellItem2) Update(bc *IBindCtx) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetUInt64,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(bc.Ppvt())))
	return utl.HresultToError(ret)
}

type _IShellItem2Vt struct {
	_IShellItemVt
	GetPropertyStore                 uintptr
	GetPropertyStoreWithCreateObject uintptr
	GetPropertyStoreForKeys          uintptr
	GetPropertyDescriptionList       uintptr
	Update                           uintptr
	GetProperty                      uintptr
	GetCLSID                         uintptr
	GetFileTime                      uintptr
	GetInt32                         uintptr
	GetString                        uintptr
	GetUInt32                        uintptr
	GetUInt64                        uintptr
	GetBool                          uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IShellItemArray] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItemArray) IID() co.IID {
	return co.IID_IShellItemArray
}

// Returns the path names of each [IShellItem] object by calling
// [IShellItemArray.GetCount], [IShellItemArray.GetItemAt] and
// [IShellItem.GetDisplayName].
//
// Example:
//
//	var arr win.IShellItemArray // initialized somewhere
//
//	names, _ := arr.EnumDisplayNames(co.SIGDN_FILESYSPATH)
//	for _, fullPath := range names {
//		println(fullPath)
//	}
func (me *IShellItemArray) EnumDisplayNames(sigdnName co.SIGDN) ([]string, error) {
	localRel := NewOleReleaser()
	defer localRel.Release()

	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, count)

	for i := 0; i < count; i++ {
		shellItem, err := me.GetItemAt(localRel, i)
		if err != nil {
			return nil, err
		}

		name, err := shellItem.GetDisplayName(sigdnName)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

// Returns all [IShellItem] objects by calling [IShellItemArray.GetCount] and
// [IShellItemArray.GetItemAt].
//
// If you just want to retrieve the paths, prefer using
// [IShellItemArray.EnumDisplayNames].
//
// Example:
//
//	var arr win.IShellItemArray // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	items, _ := arr.EnumItems(rel)
//	for _, item := range items {
//		fullPath, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(fullPath)
//	}
func (me *IShellItemArray) EnumItems(releaser *OleReleaser) ([]*IShellItem, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	items := make([]*IShellItem, 0, count)

	for i := 0; i < count; i++ {
		shellItem, err := me.GetItemAt(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		items = append(items, shellItem)
	}
	return items, nil
}

// [GetCount] method.
//
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
func (me *IShellItemArray) GetCount() (int, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		(*_IShellItemArrayVt)(unsafe.Pointer(*me.Ppvt())).GetCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&count)),
		0)

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [GetItemAt] method.
//
// Panics if index is negative.
//
// [GetItemAt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) GetItemAt(releaser *OleReleaser, index int) (*IShellItem, error) {
	utl.PanicNeg(index)
	var ppvtQueried **_IUnknownVt

	ret, _, _ := syscall.SyscallN(
		(*_IShellItemArrayVt)(unsafe.Pointer(*me.Ppvt())).GetItemAt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IShellItem](ret, ppvtQueried, releaser)
}

type _IShellItemArrayVt struct {
	_IUnknownVt
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr
	GetItemAt                  uintptr
	EnumItems                  uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IShellLink] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IShellLink]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
type IShellLink struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellLink) IID() co.IID {
	return co.IID_IShellLink
}

// [GetArguments] method.
//
// [GetArguments]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getarguments
func (me *IShellLink) GetArguments() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(utl.INFOTIPSIZE) // arbitrary

	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetArguments,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [GetDescription] method.
//
// [GetDescription]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getdescription
func (me *IShellLink) GetDescription() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(utl.INFOTIPSIZE) // arbitrary

	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetDescription,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [GetHotkey] method.
//
// [GetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-gethotkey
func (me *IShellLink) GetHotkey() (co.HOTKEYF, error) {
	var key uint16
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetHotkey,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&key)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.HOTKEYF(key), nil
	} else {
		return co.HOTKEYF(0), hr
	}
}

// [GetIconLocation] method.
//
// [GetIconLocation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-geticonlocation
func (me *IShellLink) GetIconLocation() (path string, index int, hr error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	var iconIndex uint16

	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetIconLocation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len()-1)),
		uintptr(unsafe.Pointer(&iconIndex)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), int(iconIndex), nil
	} else {
		return "", 0, hr
	}
}

// [GetPath] method.
//
// [GetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
func (me *IShellLink) GetPath(fd *WIN32_FIND_DATA, flags co.SLGP) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetPath,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len()-1)),
		uintptr(unsafe.Pointer(fd)),
		uintptr(flags))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [GetShowCmd] method.
//
// [GetShowCmd]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getshowcmd
func (me *IShellLink) GetShowCmd() (co.SW, error) {
	var cmd co.SW
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetShowCmd,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cmd)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return cmd, nil
	} else {
		return co.SW(0), hr
	}
}

// [GetWorkingDirectory] method.
//
// [GetWorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getworkingdirectory
func (me *IShellLink) GetWorkingDirectory() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetWorkingDirectory,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len()-1)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [Resolve] method.
//
// [Resolve]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
func (me *IShellLink) Resolve(hWnd HWND, flags co.SLR) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).Resolve,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [SetArguments] method.
//
// [SetArguments]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setarguments
func (me *IShellLink) SetArguments(args string) error {
	var wArgs wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetArguments,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wArgs.AllowEmpty(args)))
	return utl.HresultToError(ret)
}

// [SetDescription] method.
//
// [SetDescription]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setdescription
func (me *IShellLink) SetDescription(descr string) error {
	var wDescr wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetDescription,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wDescr.AllowEmpty(descr)))
	return utl.HresultToError(ret)
}

// [SetHotkey] method.
//
// [SetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-sethotkey
func (me *IShellLink) SetHotkey(hotkey co.HOTKEYF) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetHotkey,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hotkey))
	return utl.HresultToError(ret)
}

// [SetIconLocation] method.
//
// [SetIconLocation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-seticonlocation
func (me *IShellLink) SetIconLocation(path string, index int) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetIconLocation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wPath.AllowEmpty(path)),
		uintptr(int32(index)))
	return utl.HresultToError(ret)
}

// [SetPath] method.
//
// [SetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setpath
func (me *IShellLink) SetPath(path string) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetPath,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wPath.AllowEmpty(path)))
	return utl.HresultToError(ret)
}

// [SetRelativePath] method.
//
// [SetRelativePath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setrelativepath
func (me *IShellLink) SetRelativePath(path string) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetRelativePath,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wPath.AllowEmpty(path)))
	return utl.HresultToError(ret)
}

// [SetShowCmd] method.
//
// [SetShowCmd]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setshowcmd
func (me *IShellLink) SetShowCmd(cmd co.SW) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetShowCmd,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(cmd))
	return utl.HresultToError(ret)
}

// [SetWorkingDirectory] method.
//
// [SetWorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setworkingdirectory
func (me *IShellLink) SetWorkingDirectory(path string) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetWorkingDirectory,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wPath.AllowEmpty(path)))
	return utl.HresultToError(ret)
}

type _IShellLinkVt struct {
	_IUnknownVt
	GetPath             uintptr
	GetIDList           uintptr
	SetIDList           uintptr
	GetDescription      uintptr
	SetDescription      uintptr
	GetWorkingDirectory uintptr
	SetWorkingDirectory uintptr
	GetArguments        uintptr
	SetArguments        uintptr
	GetHotkey           uintptr
	SetHotkey           uintptr
	GetShowCmd          uintptr
	SetShowCmd          uintptr
	GetIconLocation     uintptr
	SetIconLocation     uintptr
	SetRelativePath     uintptr
	Resolve             uintptr
	SetPath             uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IShellView] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IShellView]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellview
type IShellView struct{ IOleWindow }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellView) IID() co.IID {
	return co.IID_IShellView
}

// [DestroyViewWindow] method.
//
// [DestroyViewWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-destroyviewwindow
func (me *IShellView) DestroyViewWindow() error {
	return com_callNoParm(me,
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).DestroyViewWindow)
}

// [EnableModeless] method.
//
// [EnableModeless]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-enablemodeless
func (me *IShellView) EnableModeless(enable bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).EnableModeless,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(enable))
	return utl.HresultToError(ret)
}

// [Refresh] method.
//
// [Refresh]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-refresh
func (me *IShellView) Refresh() error {
	return com_callNoParm(me,
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).Refresh)
}

// [SaveViewState] method.
//
// [SaveViewState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-saveviewstate
func (me *IShellView) SaveViewState() error {
	return com_callNoParm(me,
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).SaveViewState)
}

// [TranslateAccelerator] method.
//
// [TranslateAccelerator]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-translateaccelerator
func (me *IShellView) TranslateAccelerator(msg *MSG) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).TranslateAccelerator,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(msg)))
	return utl.HresultToError(ret)
}

// [UIActivate] method.
//
// [UIActivate]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-uiactivate
func (me *IShellView) UIActivate(state co.SVUIA) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellViewVt)(unsafe.Pointer(*me.Ppvt())).UIActivate,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(state))
	return utl.HresultToError(ret)
}

type _IShellViewVt struct {
	_IOleWindowVt
	TranslateAccelerator  uintptr
	EnableModeless        uintptr
	UIActivate            uintptr
	Refresh               uintptr
	CreateViewWindow      uintptr
	DestroyViewWindow     uintptr
	GetCurrentInfo        uintptr
	AddPropertySheetPages uintptr
	SaveViewState         uintptr
	SelectItem            uintptr
	GetItemObject         uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [ITaskbarList] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var taskbl *win.ITaskbarList
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList) IID() co.IID {
	return co.IID_ITaskbarList
}

// [ActivateTab] method.
//
// [ActivateTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).ActivateTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [AddTab] method.
//
// [AddTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).AddTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [DeleteTab] method.
//
// [DeleteTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).DeleteTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [HrInit] method.
//
// [HrInit]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() error {
	return com_callNoParm(me,
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).HrInit)
}

// [SetActiveAlt] method.
//
// [SetActiveAlt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).SetActiveAlt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

type _ITaskbarListVt struct {
	_IUnknownVt
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [ITaskbarList2] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var taskbl *win.ITaskbarList2
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct{ ITaskbarList }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList2) IID() co.IID {
	return co.IID_ITaskbarList2
}

// [MarkFullscreenWindow] method.
//
// [MarkFullscreenWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(hwnd HWND, fullScreen bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList2Vt)(unsafe.Pointer(*me.Ppvt())).MarkFullscreenWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwnd),
		utl.BoolToUintptr(fullScreen))
	return utl.HresultToError(ret)
}

type _ITaskbarList2Vt struct {
	_ITaskbarListVt
	MarkFullscreenWindow uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [ITaskbarList3] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var taskbl *win.ITaskbarList3
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList3]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct{ ITaskbarList2 }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList3) IID() co.IID {
	return co.IID_ITaskbarList3
}

// [RegisterTab] method.
//
// [RegisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).RegisterTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(hwndMDI))
	return utl.HresultToError(ret)
}

// [SetOverlayIcon] method.
//
// [SetOverlayIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(hWnd HWND, hIcon HICON, description string) error {
	var wDescription wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetOverlayIcon,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(hIcon),
		uintptr(wDescription.AllowEmpty(description)))
	return utl.HresultToError(ret)
}

// [SetProgressState] method.
//
// [SetProgressState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hWnd HWND, flags co.TBPF) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetProgressState,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [SetProgressValue] method.
//
// Panics if completed or total is negative.
//
// [SetProgressValue]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(hWnd HWND, completed, total int) error {
	utl.PanicNeg(completed, total)
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetProgressValue,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(completed),
		uintptr(total))
	return utl.HresultToError(ret)
}

// [SetTabActive] method.
//
// [SetTabActive]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetTabActive,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(hwndMDI))
	return utl.HresultToError(ret)
}

// [SetTabOrder] method.
//
// [SetTabOrder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetTabOrder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(hwndInsertBefore))
	return utl.HresultToError(ret)
}

// [SetThumbnailClip] method.
//
// [SetThumbnailClip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hWnd HWND, rcClip *RECT) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetThumbnailClip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rcClip)))
	return utl.HresultToError(ret)
}

// [SetThumbnailTooltip] method.
//
// [SetThumbnailTooltip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hWnd HWND, tip string) error {
	var wTip wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).SetThumbnailTooltip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(wTip.EmptyIsNil(tip)))
	return utl.HresultToError(ret)
}

// [ThumbBarAddButtons] method.
//
// [ThumbBarAddButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(hWnd HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).ThumbBarAddButtons,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(uint32(len(buttons))),
		uintptr(unsafe.Pointer(unsafe.SliceData(buttons))))
	return utl.HresultToError(ret)
}

// [ThumbBarSetImageList] method.
//
// [ThumbBarSetImageList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(hWnd HWND, hImgl HIMAGELIST) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).ThumbBarSetImageList,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(hImgl))
	return utl.HresultToError(ret)
}

// [ThumbBarUpdateButtons] method.
//
// [ThumbBarUpdateButtons]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(hWnd HWND, buttons []THUMBBUTTON) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).ThumbBarUpdateButtons,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd),
		uintptr(uint32(len(buttons))),
		uintptr(unsafe.Pointer(unsafe.SliceData(buttons))))
	return utl.HresultToError(ret)
}

// [UnregisterTab] method.
//
// [UnregisterTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList3Vt)(unsafe.Pointer(*me.Ppvt())).UnregisterTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab))
	return utl.HresultToError(ret)
}

type _ITaskbarList3Vt struct {
	_ITaskbarList2Vt
	SetProgressValue      uintptr
	SetProgressState      uintptr
	RegisterTab           uintptr
	UnregisterTab         uintptr
	SetTabOrder           uintptr
	SetTabActive          uintptr
	ThumbBarAddButtons    uintptr
	ThumbBarUpdateButtons uintptr
	ThumbBarSetImageList  uintptr
	SetOverlayIcon        uintptr
	SetThumbnailTooltip   uintptr
	SetThumbnailClip      uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [ITaskbarList4] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var taskbl *win.ITaskbarList4
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList4]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct{ ITaskbarList3 }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList4) IID() co.IID {
	return co.IID_ITaskbarList4
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist4-settabproperties
func (me *ITaskbarList4) SetProperties(hwndTab HWND, flags co.STPFLAG) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList4Vt)(unsafe.Pointer(*me.Ppvt())).SetTabProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab),
		uintptr(flags))
	return utl.HresultToError(ret)
}

type _ITaskbarList4Vt struct {
	_ITaskbarList3Vt
	SetTabProperties uintptr
}
