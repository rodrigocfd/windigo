/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package api

import "syscall"

type (
	ATOM     uint16
	COLORREF uint32
	HANDLE   syscall.Handle
	HBITMAP  HGDIOBJ
	HCURSOR  HANDLE
	HGDIOBJ  HANDLE
	HICON    HANDLE
	HRGN     HGDIOBJ
)

type (
	LPARAM uintptr
	WPARAM uintptr
)
