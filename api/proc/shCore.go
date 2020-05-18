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
	dllShCore = syscall.NewLazyDLL("shcore.dll")

	GetDpiForMonitor = dllShCore.NewProc("GetDpiForMonitor")
)
