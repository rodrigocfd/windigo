/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IFilterGraph struct{ _IUnknown }

	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
	//
	// IFilterGraph > IUnknown.
	IFilterGraph struct{ _IFilterGraph }

	_IFilterGraphVtbl struct {
		_IUnknownVtbl
		AddFilter            uintptr
		RemoveFilter         uintptr
		EnumFilters          uintptr
		FindFilterByName     uintptr
		ConnectDirect        uintptr
		Reconnect            uintptr
		Disconnect           uintptr
		SetDefaultSyncSource uintptr
	}
)
