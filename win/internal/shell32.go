/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	shell32Dll = syscall.NewLazyDLL("shell32.dll")

	DragAcceptFiles             = shell32Dll.NewProc("DragAcceptFiles")
	DragFinish                  = shell32Dll.NewProc("DragFinish")
	DragQueryFile               = shell32Dll.NewProc("DragQueryFileW")
	DragQueryPoint              = shell32Dll.NewProc("DragQueryPoint")
	DuplicateIcon               = shell32Dll.NewProc("DuplicateIcon")
	SHCreateItemFromParsingName = shell32Dll.NewProc("SHCreateItemFromParsingName")
	SHGetFileInfo               = shell32Dll.NewProc("SHGetFileInfoW")
)
