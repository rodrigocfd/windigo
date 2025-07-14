//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IFileSaveDialog] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// # Example
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
		uintptr(ppvtOrNil(sink)))
	return utl.ErrorAsHResult(ret)
}

// [GetProperties] method.
//
// [GetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-getproperties
func (me *IFileSaveDialog) GetProperties(releaser *OleReleaser) (*IPropertyStore, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).GetProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IPropertyStore{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setproperties
func (me *IFileSaveDialog) SetProperties(store *IPropertyStore) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).SetProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(store.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SetSaveAsItem] method.
//
// [SetSaveAsItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(item *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileSaveDialogVt)(unsafe.Pointer(*me.Ppvt())).SetSaveAsItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

type _IFileSaveDialogVt struct {
	_IFileDialogVt
	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}
