/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"fmt"
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

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
//
// Must be freed with CoUninitialize().
func CoInitializeEx(dwCoInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(dwCoInit), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK && lerr != co.ERROR_S_FALSE {
		panic(fmt.Sprintf("CoInitializeEx failed. %s", lerr.Error()))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
