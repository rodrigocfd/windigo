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
	dllOle32 = syscall.NewLazyDLL("ole32.dll")

	CoCreateInstance = dllOle32.NewProc("CoCreateInstance")
	CoInitializeEx   = dllOle32.NewProc("CoInitializeEx")
	CoUninitialize   = dllOle32.NewProc("CoUninitialize")
)
