/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IBaseFilterImpl struct{ _IMediaFilterImpl }

	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
	//
	// IBaseFilter > IMediaFilter > IPersist > IUnknown.
	IBaseFilter struct{ _IBaseFilterImpl }

	_IBaseFilterVtbl struct {
		_IMediaFilterVtbl
		EnumPins        uintptr
		FindPin         uintptr
		QueryFilterInfo uintptr
		JoinFilterGraph uintptr
		QueryVendorInfo uintptr
	}
)
