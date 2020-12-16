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
	shcoreDll = syscall.NewLazyDLL("shcore.dll")

	GetDpiForMonitor = shcoreDll.NewProc("GetDpiForMonitor")
)
