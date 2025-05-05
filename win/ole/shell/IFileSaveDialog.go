//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IFileSaveDialog] COM interface.
//
// # Example
//
//	var hWnd win.HWND // initialized somewhere
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	fsd, _ := ole.CoCreateInstance[shell.IFileSaveDialog](
//		rel,
//		co.CLSID_FileSaveDialog,
//		co.CLSCTX_INPROC_SERVER,
//	)
//
//	fsd.SetFileTypes([]shell.COMDLG_FILTERSPEC{
//		{Name: "Text files", Spec: "*.txt"},
//		{Name: "All files", Spec: "*.*"},
//	})
//	fsd.SetFileTypeIndex(1)
//
//	fsd.SetFileName("default-file-name.txt")
//
//	if ok, _ := fsd.Show(hWnd); ok {
//		item, _ := fsd.GetResult(rel)
//		txtPath, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(txtPath)
//	}
//
// [IFileSaveDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog struct{ IFileDialog }

// Returns the unique COM interface identifier.
func (*IFileSaveDialog) IID() co.IID {
	return co.IID_IFileSaveDialog
}

// [SetSaveAsItem] method.
//
// [SetSaveAsItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.IFileSaveDialog)(unsafe.Pointer(*me.Ppvt())).SetSaveAsItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())))
	return util.ErrorAsHResult(ret)
}
