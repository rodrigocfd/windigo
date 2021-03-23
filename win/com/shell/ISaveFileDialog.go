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

// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIFileSaveDialog(dwClsContext co.CLSCTX) IFileSaveDialog {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0xc0b4e2f3, 0xba21, 0x4773, 0x8dba, 0x335ec946eb8b), // CLSID_FileSaveDialog
		nil,
		dwClsContext,
		win.NewGuid(0x84bccd23, 0x5fde, 0x4cdb, 0xaea4, 0xaf64b83d78ab), // IID_IFileSaveDialog
	)
	if err != nil {
		panic(err)
	}
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
