/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type IDispatch struct {
	IUnknown
}

type iDispatchVtbl struct {
	iUnknownVtbl
	GetIDsOfNames    uintptr
	GetTypeInfo      uintptr
	GetTypeInfoCount uintptr
	Invoke           uintptr
}
