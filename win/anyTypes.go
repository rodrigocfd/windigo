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
	ATOM      uint16
	COLORREF  uint32
	HANDLE    syscall.Handle
	HBITMAP   HGDIOBJ
	HGDIOBJ   HANDLE
	HRGN      HGDIOBJ
	HTREEITEM HANDLE
)
