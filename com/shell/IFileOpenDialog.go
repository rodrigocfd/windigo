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
	// IFileOpenDialog > IFileDialog > IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
	IFileOpenDialog struct{ win.IUnknown }

	IFileOpenDialogVtbl struct {
		win.IUnknownVtbl
		GetResults       uintptr
		GetSelectedItems uintptr
	}
)
