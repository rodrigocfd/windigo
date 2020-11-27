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
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ns-shobjidl_core-thumbbutton
	THUMBBUTTON struct {
		DwMask  THB
		IId     uint32
		IBitmap uint32
		HIcon   win.HICON
		SzTip   [260]uint16
		DwFlags THBF
	}
)
