/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"windigo/win"
)

type (
	// IFileDialog > IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
	IFileDialog struct{ win.IUnknown }

	IFileDialogVtbl struct {
		IModalWindowVtbl
		SetFileTypes        uintptr
		SetFileTypeIndex    uintptr
		GetFileTypeIndex    uintptr
		Advise              uintptr
		Unadvise            uintptr
		SetOptions          uintptr
		GetOptions          uintptr
		SetDefaultFolder    uintptr
		SetFolder           uintptr
		GetFolder           uintptr
		GetCurrentSelection uintptr
		SetFileName         uintptr
		GetFileName         uintptr
		SetTitle            uintptr
		SetOkButtonLabel    uintptr
		SetFileNameLabel    uintptr
		GetResult           uintptr
		AddPlace            uintptr
		SetDefaultExtension uintptr
		Close               uintptr
		SetClientGuid       uintptr
		ClearClientData     uintptr
		SetFilter           uintptr
	}
)
