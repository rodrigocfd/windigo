package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog struct{ IFileDialog }

// Constructs a COM object from the base IUnknown.
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
func NewIFileSaveDialog(base win.IUnknown) IFileSaveDialog {
	return IFileSaveDialog{IFileDialog: NewIFileDialog(base)}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(si *IShellItem) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileSaveDialog)(unsafe.Pointer(*me.Ptr())).SetSaveAsItem, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
