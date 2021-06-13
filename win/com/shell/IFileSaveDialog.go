package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
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
type IFileSaveDialog struct {
	IFileDialog // Base IFileDialog > IModalWindow > IUnknown.
}

// Calls IUnknown.CoCreateInstance() to return IFileSaveDialog.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIFileSaveDialog(dwClsContext co.CLSCTX) IFileSaveDialog {
	iUnk := win.CoCreateInstance(
		CLSID.FileSaveDialog, nil, dwClsContext, IID.IFileSaveDialog)
	return IFileSaveDialog{
		IFileDialog{
			IModalWindow{IUnknown: iUnk},
		},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(psi *IShellItem) {
	ret, _, _ := syscall.Syscall(
		(*_IFileSaveDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetSaveAsItem, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(psi.Ppv)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
