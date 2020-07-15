/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// IBaseFilter > IMediaFilter > IPersist > IUnknown.
type IBaseFilter struct {
	IMediaFilter
}

type iBaseFilterVtbl struct {
	iMediaFilterVtbl
	EnumPins        uintptr
	FindPin         uintptr
	QueryFilterInfo uintptr
	JoinFilterGraph uintptr
	QueryVendorInfo uintptr
}
