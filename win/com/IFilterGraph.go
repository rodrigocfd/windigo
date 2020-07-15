/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// IFilterGraph > IUnknown.
type IFilterGraph struct {
	IUnknown
}

type iFilterGraphVtbl struct {
	iUnknownVtbl
	AddFilter            uintptr
	RemoveFilter         uintptr
	EnumFilters          uintptr
	FindFilterByName     uintptr
	ConnectDirect        uintptr
	Reconnect            uintptr
	Disconnect           uintptr
	SetDefaultSyncSource uintptr
}
