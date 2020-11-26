/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"windigo/co"
)

// A Win32 error, usually retrieved with GetLastError(). Contains the error code
// and the name of the function which triggered it.
//
// Implements error interface.
//
// https://docs.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-getlasterror
type WinError struct {
	code         co.ERROR
	functionName string
}

// Creates a new WinError.
func NewWinError(code co.ERROR, functionName string) *WinError {
	return &WinError{
		code:         code,
		functionName: functionName,
	}
}

// Returns the error code.
func (e *WinError) Code() co.ERROR {
	return e.code
}

// Returns the name of the function which triggered the error.
func (e *WinError) FunctionName() string {
	return e.functionName
}

// Calls FormatMessage() and returns a full error description.
//
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-formatmessagew
func (e *WinError) Error() string {
	return fmt.Sprintf("%s [%d 0x%02x] %s",
		e.functionName,
		uint32(e.code),
		uint32(e.code),
		syscall.Errno(e.code).Error(),
	)
}
