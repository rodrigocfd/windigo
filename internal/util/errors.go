//go:build windows

package util

import (
	"syscall"

	"github.com/rodrigocfd/windigo/win/co"
)

// Error handling syntactic sugar for syscalls which call GetLastError().
func ZeroAsGetLastError(ret uintptr, err syscall.Errno) error {
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

// Error handling syntactic sugar for syscalls directly returning a standard
// Windows error.
func ZeroAsSysError(ret uintptr) error {
	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return wErr
	}
	return nil
}

// Error handling syntactic sugar for syscalls returning the standard
// co.ERROR_INVALID_PARAMETER.
func ZeroAsSysInvalidParm(ret uintptr) error {
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

// Error handling syntactic sugar for syscalls returning -1.
func Minus1AsSysInvalidParm(ret uintptr) error {
	if int32(ret) == -1 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

// Error handling syntactic sugar for syscalls directly returning an HRESULT.
func ErrorAsHResult(ret uintptr) error {
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return hr
	}
	return nil
}
