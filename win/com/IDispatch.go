/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// IDispatch > IUnknown.
type IDispatch struct {
	IUnknown
}

type iDispatchVtbl struct {
	iUnknownVtbl
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}
