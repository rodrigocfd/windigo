package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IFileSaveDialogVtbl struct {
	_IFileDialogVtbl
	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog struct{ IFileDialog }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IFileSaveDialog.Release().
//
// Example:
//
//  fsd := shell.NewIFileSaveDialog(
//      win.CoCreateInstance(
//          shellco.CLSID_FileSaveDialog, nil,
//          co.CLSCTX_INPROC_SERVER,
//          shellco.IID_IFileSaveDialog),
//  )
//  defer fsd.Release()
func NewIFileSaveDialog(ptr win.IUnknownPtr) IFileSaveDialog {
	return IFileSaveDialog{
		IFileDialog: NewIFileDialog(ptr),
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(si *IShellItem) {
	ret, _, _ := syscall.Syscall(
		(*_IFileSaveDialogVtbl)(unsafe.Pointer(*me.Ptr())).SetSaveAsItem, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
