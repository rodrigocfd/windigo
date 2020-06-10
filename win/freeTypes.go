/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

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

type WPARAM uintptr

func (wp WPARAM) LoWord() uint16 { return LoWord(uint32(wp)) }
func (wp WPARAM) HiWord() uint16 { return HiWord(uint32(wp)) }

type LPARAM uintptr

func (lp LPARAM) LoWord() uint16 { return LoWord(uint32(lp)) }
func (lp LPARAM) HiWord() uint16 { return HiWord(uint32(lp)) }
