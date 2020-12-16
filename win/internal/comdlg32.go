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
	comdlg32Dll = syscall.NewLazyDLL("comdlg32.dll")

	CommDlgExtendedError = comdlg32Dll.NewProc("CommDlgExtendedError")
	GetOpenFileName      = comdlg32Dll.NewProc("GetOpenFileNameW")
	GetSaveFileName      = comdlg32Dll.NewProc("GetSaveFileNameW")
)
