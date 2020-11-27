/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
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
