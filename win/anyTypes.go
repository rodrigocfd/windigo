/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
	ATOM uint16

	// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#colorref
	COLORREF uint32

	// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
	HANDLE syscall.Handle

	// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbitmap
	HBITMAP HGDIOBJ

	// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
	HGDIOBJ HANDLE

	// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
	HRGN HGDIOBJ

	// https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls#parent-and-child-items
	HTREEITEM HANDLE
)

// Private constants.
const (
	_CCHDEVICENAME        = 32
	_CCHILDREN_TITLEBAR   = 5
	_CLR_INVALID          = 0xFFFF_FFFF
	_HGDI_ERROR           = 0xFFFF_FFFF
	_INVALID_FILE_SIZE    = 0xFFFF_FFFF
	_INVALID_HANDLE_VALUE = -1
	_L_MAX_URL_LENGTH     = 2048 + 32 + 3
	_LF_FACESIZE          = 32
	_MAX_LINKID_TEXT      = 48
	_MAX_PATH             = 260
)
