/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"wingows/co"
	"wingows/win/proc"
)

// Simple conversion for syscalls.
func boolToUintptr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// Initializes the COM library.
//
// Must be freed with CoUninitialize().
func CoInitializeEx(dwCoInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(dwCoInit), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK && lerr != co.ERROR_S_FALSE {
		panic(lerr.Format("CoInitializeEx failed."))
	}
}

// Closes the COM library and frees the resources.
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
