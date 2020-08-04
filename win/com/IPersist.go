/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IPersist struct{ _IUnknown }

	// IPersist > IUnknown.
	IPersist struct{ _IPersist }

	_IPersistVtbl struct {
		_IUnknownVtbl
		GetClassID uintptr
	}
)
