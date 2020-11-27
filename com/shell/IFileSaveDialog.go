/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package shell

import (
	"windigo/win"
)

type (
	// IFileSaveDialog > IFileDialog > IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
	IFileSaveDialog struct{ win.IUnknown }

	IFileSaveDialogVtbl struct {
		win.IUnknownVtbl
		SetSaveAsItem          uintptr
		SetProperties          uintptr
		SetCollectedProperties uintptr
		GetProperties          uintptr
		ApplyProperties        uintptr
	}
)
