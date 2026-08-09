//go:build windows

package winsh

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IFileOpenDialog] COM interface.
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
//	var fod *winsh.IFileOpenDialog
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_FileOpenDialog,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&fod,
//	)
//
//	defOpts, _ := fod.GetOptions()
//	_ = fod.SetOptions(defOpts |
//		cosh.FOS_FORCEFILESYSTEM |
//		cosh.FOS_FILEMUSTEXIST,
//	)
//
//	_ = fod.SetFileTypes([]winsh.COMDLG_FILTERSPEC{
//		{Name: "Text files", Spec: "*.txt"},
//		{Name: "All files", Spec: "*.*"},
//	})
//	_ = fod.SetFileTypeIndex(1)
//
//	if ok, _ := fod.Show(hWnd); ok {
//		item, _ := fod.GetResult(rel)
//		fileName, _ := item.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//		println(fileName)
//	}
//
// [IFileOpenDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog struct{ IFileDialog }

type _IFileOpenDialogVt struct {
	_IFileDialogVt
	GetResults       uintptr
	GetSelectedItems uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOpenDialog) IID() *co.IID {
	return &cosh.IID_IFileOpenDialog
}

// [GetResults] method.
//
// Returns the selected items after user confirmation, for multi-selection
// dialogs – those with [cosh.FOS_ALLOWMULTISELECT] option.
//
// For single-selection dialogs, use [IFileDialog.GetResult].
//
// [GetResults]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults(releaser *win.OleReleaser) (*IShellItemArray, error) {
	return utl.OleNewFromCallWithoutParms[*IShellItemArray](me, releaser,
		utl.Vt[_IFileOpenDialogVt](me.Ppvt()).GetResults)
}

// [GetSelectedItems] method.
//
// [GetSelectedItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems(releaser *win.OleReleaser) (*IShellItemArray, error) {
	return utl.OleNewFromCallWithoutParms[*IShellItemArray](me, releaser,
		utl.Vt[_IFileOpenDialogVt](me.Ppvt()).GetSelectedItems)
}
