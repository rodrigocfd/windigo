/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	baseIBaseFilter struct{ baseIMediaFilter }

	// IBaseFilter > IMediaFilter > IPersist > IUnknown.
	IBaseFilter struct{ baseIBaseFilter }

	vtbIBaseFilter struct {
		vtbIMediaFilter
		EnumPins        uintptr
		FindPin         uintptr
		QueryFilterInfo uintptr
		JoinFilterGraph uintptr
		QueryVendorInfo uintptr
	}
)
