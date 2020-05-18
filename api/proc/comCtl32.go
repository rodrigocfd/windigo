/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllComCtl32 = syscall.NewLazyDLL("comctl32.dll")

	DefSubclassProc      = dllComCtl32.NewProc("DefSubclassProc")
	InitCommonControls   = dllComCtl32.NewProc("InitCommonControls")
	RemoveWindowSubclass = dllComCtl32.NewProc("RemoveWindowSubclass")
	SetWindowSubclass    = dllComCtl32.NewProc("SetWindowSubclass")
)
