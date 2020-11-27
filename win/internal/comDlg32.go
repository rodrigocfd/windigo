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
	dllComDlg32 = syscall.NewLazyDLL("comdlg32.dll")

	CommDlgExtendedError = dllComDlg32.NewProc("CommDlgExtendedError")
	GetOpenFileName      = dllComDlg32.NewProc("GetOpenFileNameW")
	GetSaveFileName      = dllComDlg32.NewProc("GetSaveFileNameW")
)
