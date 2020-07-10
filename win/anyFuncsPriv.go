/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"wingows/co"
)

func hiWord(value uint32) uint16 { return uint16(value >> 16 & 0xffff) }
func loWord(value uint32) uint16 { return uint16(value) }
func hiByte(value uint16) uint8  { return uint8(value >> 8 & 0xff) }
func loByte(value uint16) uint8  { return uint8(value) }

// Simple conversion for syscalls.
func boolToUintptr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// Calls the specific system call to release the handle, returning the error
// code without panic.
func freeNoPanic(h HANDLE, fun *syscall.LazyProc) co.ERROR {
	if h == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(fun.Addr(), 1, uintptr(h), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}
