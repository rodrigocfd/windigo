/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
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
//
// WinError implements error interface, and stores GetLastError() code and the
// name of the Win32 which originated the error.
func NewWinError(code co.ERROR, functionName string) *WinError {
	return &WinError{
		code:         code,
		functionName: functionName,
	}
}

// Returns the error code.
func (e *WinError) Code() co.ERROR { return e.code }

// Returns the name of the function which triggered the error.
func (e *WinError) FunctionName() string { return e.functionName }

// Returns the contained ERROR.
func (e *WinError) Unwrap() error { return e.code }

// Calls FormatMessage() and returns a full error description.
//
// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-formatmessagew
func (e *WinError) Error() string {
	return fmt.Sprintf("%s %s",
		e.functionName, e.Unwrap().Error()) // FormatMessage()
}
