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
	ole32Dll = syscall.NewLazyDLL("ole32.dll")

	CoCreateInstance = ole32Dll.NewProc("CoCreateInstance")
	CoInitializeEx   = ole32Dll.NewProc("CoInitializeEx")
	CoTaskMemFree    = ole32Dll.NewProc("CoTaskMemFree")
	CoUninitialize   = ole32Dll.NewProc("CoUninitialize")
)
