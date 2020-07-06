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
	dllShell32 = syscall.NewLazyDLL("shell32.dll")

	DragFinish     = dllShell32.NewProc("DragFinish")
	DragQueryFile  = dllShell32.NewProc("DragQueryFileW")
	DragQueryPoint = dllShell32.NewProc("DragQueryPoint")
	SHGetFileInfo  = dllShell32.NewProc("SHGetFileInfoW")
)
