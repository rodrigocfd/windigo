//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
)

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
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IFileOpenDialogVt)(unsafe.Pointer(*me.Ppvt())).GetResults,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItemArray{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetSelectedItems] method.
//
// [GetSelectedItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems(releaser *OleReleaser) (*IShellItemArray, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IFileOpenDialogVt)(unsafe.Pointer(*me.Ppvt())).GetSelectedItems,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItemArray{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IFileOpenDialogVt struct {
	_IFileDialogVt
	GetResults       uintptr
	GetSelectedItems uintptr
}
