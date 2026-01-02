//go:build windows

package utl

import (
	"fmt"
	"syscall"

	"github.com/rodrigocfd/windigo/co"
)

// Panics if any of the numbers is negative.
func PanicNeg(nums ...int) {
	for _, n := range nums {
		if n < 0 {
			panic(fmt.Sprintf("Value cannot be negative: %d.", n))
		}
	}
}

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
func HresultToError(ret uintptr) error {
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return hr
	}
	return nil
}

// Error handling syntactic sugar for syscalls directly returning bool + HRESULT.
func HresultToBoolError(bVal int32, ret uintptr) (bool, error) {
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bVal != 0, nil
	} else {
		return false, hr
	}
}
