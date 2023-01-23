//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog interface {
	IFileDialog

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
	SetSaveAsItem(si IShellItem)
}

type _IFileSaveDialog struct{ IFileDialog }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileSaveDialog.Release().
//
// Example:
//
//	fsd := shell.NewIFileSaveDialog(
//		com.CoCreateInstance(
//			shellco.CLSID_FileSaveDialog, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			shellco.IID_IFileSaveDialog),
//	)
//	defer fsd.Release()
func NewIFileSaveDialog(base com.IUnknown) IFileSaveDialog {
	return &_IFileSaveDialog{IFileDialog: NewIFileDialog(base)}
}

func (me *_IFileSaveDialog) SetSaveAsItem(si IShellItem) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileSaveDialog)(unsafe.Pointer(*me.Ptr())).SetSaveAsItem,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
