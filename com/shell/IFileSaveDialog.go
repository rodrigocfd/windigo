/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// IFileSaveDialog > IFileDialog > IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
	IFileSaveDialog struct{ IFileDialog }

	IFileSaveDialogVtbl struct {
		IFileDialogVtbl
		SetSaveAsItem          uintptr
		SetProperties          uintptr
		SetCollectedProperties uintptr
		GetProperties          uintptr
		ApplyProperties        uintptr
	}
)

// Typically uses CLSCTX_INPROC_SERVER.
//
// You must defer Release().
func CoCreateIFileSaveDialog(dwClsContext co.CLSCTX) *IFileSaveDialog {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0xc0b4e2f3, 0xba21, 0x4773, 0x8dba, 0x335ec946eb8b), // CLSID_FileSaveDialog
		nil,
		dwClsContext,
		win.NewGuid(0x84bccd23, 0x5fde, 0x4cdb, 0xaea4, 0xaf64b83d78ab), // IID_IFileSaveDialog
	)
	if err != nil {
		panic(err)
	}
	return &IFileSaveDialog{
		IFileDialog{
			IModalWindow{
				IUnknown: *iUnk,
			},
		},
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifilesavedialog-setsaveasitem
func (me *IFileSaveDialog) SetSaveAsItem(psi *IShellItem) {
	ret, _, _ := syscall.Syscall(
		(*IFileSaveDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetSaveAsItem, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(psi.Ppv)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileSaveDialog.SetSaveAsItem"))
	}
}
