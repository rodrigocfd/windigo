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
	dllGdi32 = syscall.NewLazyDLL("gdi32.dll")

	CreateFontIndirect = dllGdi32.NewProc("CreateFontIndirectW")
	DeleteObject       = dllGdi32.NewProc("DeleteObject")
	GetDeviceCaps      = dllGdi32.NewProc("GetDeviceCaps")
)
