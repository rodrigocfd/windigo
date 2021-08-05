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
	return Bytes.Lo8(Bytes.Lo16(uint32(c)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getgvalue
func (c COLORREF) GetGValue() uint8 {
	return Bytes.Lo8(Bytes.Lo16(uint32(c) >> 8))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbvalue
func (c COLORREF) GetBValue() uint8 {
	return Bytes.Lo8(Bytes.Lo16(uint32(c) >> 16))
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

// First message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makewparam
func MAKEWPARAM(lo, hi uint16) WPARAM {
	return WPARAM(Bytes.Make32(lo, hi))
}

func (wp WPARAM) Lo16() uint16   { return Bytes.Lo16(uint32(wp)) }
func (wp WPARAM) Hi16() uint16   { return Bytes.Hi16(uint32(wp)) }
func (wp WPARAM) Lo8Lo16() uint8 { return Bytes.Lo8(wp.Lo16()) }
func (wp WPARAM) Hi8Lo16() uint8 { return Bytes.Hi8(wp.Lo16()) }
func (wp WPARAM) Lo8Hi16() uint8 { return Bytes.Lo8(wp.Hi16()) }
func (wp WPARAM) Hi8Hi16() uint8 { return Bytes.Hi8(wp.Hi16()) }

// Second message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makelparam
func MAKELPARAM(lo, hi uint16) LPARAM {
	return LPARAM(Bytes.Make32(lo, hi))
}

func (lp LPARAM) Lo16() uint16   { return Bytes.Lo16(uint32(lp)) }
func (lp LPARAM) Hi16() uint16   { return Bytes.Hi16(uint32(lp)) }
func (lp LPARAM) Lo8Lo16() uint8 { return Bytes.Lo8(lp.Lo16()) }
func (lp LPARAM) Hi8Lo16() uint8 { return Bytes.Hi8(lp.Lo16()) }
func (lp LPARAM) Lo8Hi16() uint8 { return Bytes.Lo8(lp.Hi16()) }
func (lp LPARAM) Hi8Hi16() uint8 { return Bytes.Hi8(lp.Hi16()) }

func (lp LPARAM) MakePoint() POINT {
	return POINT{
		X: int32(lp.Lo16()),
		Y: int32(lp.Hi16()),
	}
}

func (lp LPARAM) MakeSize() SIZE {
	return SIZE{
		Cx: int32(lp.Lo16()),
		Cy: int32(lp.Hi16()),
	}
}
