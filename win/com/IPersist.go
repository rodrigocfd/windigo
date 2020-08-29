/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IPersistImpl struct{ _IUnknownImpl }

	// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
	//
	// IPersist > IUnknown.
	IPersist struct{ _IPersistImpl }

	_IPersistVtbl struct {
		_IUnknownVtbl
		GetClassID uintptr
	}
)
