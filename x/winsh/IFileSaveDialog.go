//go:build windows

package winsh

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IFileSaveDialog] COM interface.
//
// Example:
//
//	var hWnd win.HWND // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var fsd *winsh.IFileSaveDialog
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_FileSaveDialog,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&fsd,
//	)
//
//	_ = fsd.SetFileTypes([]winsh.COMDLG_FILTERSPEC{
//		{Name: "Text files", Spec: "*.txt"},
//		{Name: "All files", Spec: "*.*"},
//	})
//	_ = fsd.SetFileTypeIndex(1)
//
//	_ = fsd.SetFileName("default-file-name.txt")
//
//	if ok, _ := fsd.Show(hWnd); ok {
//		item, _ := fsd.GetResult(rel)
//		txtPath, _ := item.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//		println(txtPath)
//	}
//
// [IFileSaveDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog struct{ IFileDialog }

type _IFileSaveDialogVt struct {
	_IFileDialogVt
	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileSaveDialog) IID() *co.IID {
	return &cosh.IID_IFileSaveDialog
}

// [ApplyProperties] method.
//
// [ApplyProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-applyproperties
func (me *IFileSaveDialog) ApplyProperties(
	item *IShellItem,
	store *IPropertyStore,
	hwnd win.HWND,
	sink *IFileOperationProgressSink,
) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileSaveDialogVt](me.Ppvt()).ApplyProperties,
		me.Ppvt(),
		item.Ppvt(),
		store.Ppvt(),
		uintptr(hwnd),
		utl.OlePpvtOrNil(sink))
	return utl.HresultToError(ret)
}

// [GetProperties] method.
//
// [GetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-getproperties
func (me *IFileSaveDialog) GetProperties(releaser *win.OleReleaser) (*IPropertyStore, error) {
	return utl.OleNewFromCallWithoutParms[*IPropertyStore](me, releaser,
		utl.Vt[_IFileSaveDialogVt](me.Ppvt()).GetProperties)
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setproperties
func (me *IFileSaveDialog) SetProperties(store *IPropertyStore) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileSaveDialogVt](me.Ppvt()).SetProperties,
		me.Ppvt(),
		store.Ppvt())
	return utl.HresultToError(ret)
}

// [SetSaveAsItem] method.
//
// [SetSaveAsItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(item *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileSaveDialogVt](me.Ppvt()).SetSaveAsItem,
		me.Ppvt(),
		item.Ppvt())
	return utl.HresultToError(ret)
}
