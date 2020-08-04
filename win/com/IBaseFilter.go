/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IBaseFilter struct{ _IMediaFilter }

	// IBaseFilter > IMediaFilter > IPersist > IUnknown.
	IBaseFilter struct{ _IBaseFilter }

	_IBaseFilterVtbl struct {
		_IMediaFilterVtbl
		EnumPins        uintptr
		FindPin         uintptr
		QueryFilterInfo uintptr
		JoinFilterGraph uintptr
		QueryVendorInfo uintptr
	}
)
