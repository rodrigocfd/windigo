package win

import (
	"syscall"
)

// An atom.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// A Boolean variable (should be TRUE or FALSE).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#bool
type BOOL int32

// Specifies an RGB color.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/colorref
type COLORREF uint32

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrvalue
func (c COLORREF) GetRValue() uint8 {
	return LOBYTE(LOWORD(uint32(c)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getgvalue
func (c COLORREF) GetGValue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 8))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbvalue
func (c COLORREF) GetBValue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 16))
}

// A handle to an object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HANDLE syscall.Handle

// A handle to a cursor.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// Handle to a display monitor.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hmonitor
type HMONITOR HANDLE

// A handle to a tree view control item.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls#parent-and-child-items
type HTREEITEM HANDLE

// A message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

func (wp WPARAM) LoWord() uint16      { return LOWORD(uint32(wp)) }
func (wp WPARAM) HiWord() uint16      { return HIWORD(uint32(wp)) }
func (wp WPARAM) LoByteLoWord() uint8 { return LOBYTE(wp.LoWord()) }
func (wp WPARAM) HiByteLoWord() uint8 { return HIBYTE(wp.LoWord()) }
func (wp WPARAM) LoByteHiWord() uint8 { return LOBYTE(wp.HiWord()) }
func (wp WPARAM) HiByteHiWord() uint8 { return HIBYTE(wp.HiWord()) }

// A message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

func (lp LPARAM) LoWord() uint16      { return LOWORD(uint32(lp)) }
func (lp LPARAM) HiWord() uint16      { return HIWORD(uint32(lp)) }
func (lp LPARAM) LoByteLoWord() uint8 { return LOBYTE(lp.LoWord()) }
func (lp LPARAM) HiByteLoWord() uint8 { return HIBYTE(lp.LoWord()) }
func (lp LPARAM) LoByteHiWord() uint8 { return LOBYTE(lp.HiWord()) }
func (lp LPARAM) HiByteHiWord() uint8 { return HIBYTE(lp.HiWord()) }

func (lp LPARAM) MakePoint() POINT {
	return POINT{
		X: int32(lp.LoWord()),
		Y: int32(lp.HiWord()),
	}
}

func (lp LPARAM) MakeSize() SIZE {
	return SIZE{
		Cx: int32(lp.LoWord()),
		Cy: int32(lp.HiWord()),
	}
}
