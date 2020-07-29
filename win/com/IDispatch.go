/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	baseIDispatch struct{ baseIUnknown }

	// IDispatch > IUnknown.
	IDispatch struct{ baseIDispatch }

	vtbIDispatch struct {
		vtbIUnknown
		GetTypeInfoCount uintptr
		GetTypeInfo      uintptr
		GetIDsOfNames    uintptr
		Invoke           uintptr
	}
)
