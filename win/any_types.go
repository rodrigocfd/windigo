package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/win/co"
)

// An atom.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// Specifies an RGB color.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/colorref
type COLORREF uint32

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rgb
func RGB(red, green, blue uint8) COLORREF {
	return COLORREF(uint32(red) | (uint32(green) << 8) | (uint32(blue) << 16))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrvalue
func (c COLORREF) Red() uint8 {
	return LOBYTE(LOWORD(uint32(c)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getgvalue
func (c COLORREF) Green() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 8))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbvalue
func (c COLORREF) Blue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 16))
}

func (c COLORREF) ToRgbquad() RGBQUAD {
	rq := RGBQUAD{}
	rq.SetBlue(c.Blue())
	rq.SetGreen(c.Green())
	rq.SetRed(c.Red())
	return rq
}

// A handle to an object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HANDLE syscall.Handle

// A handle to a tree view control item.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls#parent-and-child-items
type HTREEITEM HANDLE

// Language and sublanguage identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/language-identifiers
type LANGID uint16

// Predefined language identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/language-identifiers
const (
	LANGID_SYSTEM_DEFAULT LANGID = LANGID((uint16(co.SUBLANG_SYS_DEFAULT) << 10) | uint16(co.LANG_NEUTRAL))
	LANGID_USER_DEFAULT   LANGID = LANGID((uint16(co.SUBLANG_DEFAULT) << 10) | uint16(co.LANG_NEUTRAL))
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelangid
func MAKELANGID(lang co.LANG, subLang co.SUBLANG) LANGID {
	return LANGID((uint16(subLang) << 10) | uint16(lang))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-primarylangid
func (lid LANGID) Lang() co.LANG { return co.LANG(uint16(lid) & 0x3ff) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-sublangid
func (lid LANGID) SubLang() co.SUBLANG { return co.SUBLANG(uint16(lid) >> 10) }

// Locale identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/locale-identifiers
type LCID uint32

// Predefined locale identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/locale-identifiers
const (
	LCID_SYSTEM_DEFAULT LCID = LCID((uint32(co.SORT_DEFAULT) << 16) | uint32(LANGID_SYSTEM_DEFAULT))
	LCID_USER_DEFAULT   LCID = LCID((uint32(co.SORT_DEFAULT) << 16) | uint32(LANGID_USER_DEFAULT))
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelcid
func MAKELCID(langId LANGID, sortId co.SORT) LCID {
	return LCID((uint32(sortId) << 16) | uint32(langId))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-langidfromlcid
func (lcid LCID) LangId() LANGID { return LANGID(uint16(lcid)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-sortidfromlcid
func (lcid LCID) SortId() co.SORT { return co.SORT((uint32(lcid) >> 16) & 0xf) }

// First message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makewparam
func MAKEWPARAM(lo, hi uint16) WPARAM {
	return WPARAM(MAKELONG(lo, hi))
}

func (wp WPARAM) LoWord() uint16 { return LOWORD(uint32(wp)) }
func (wp WPARAM) HiWord() uint16 { return HIWORD(uint32(wp)) }

// Second message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makelparam
func MAKELPARAM(lo, hi uint16) LPARAM {
	return LPARAM(MAKELONG(lo, hi))
}

func (lp LPARAM) LoWord() uint16 { return LOWORD(uint32(lp)) }
func (lp LPARAM) HiWord() uint16 { return HIWORD(uint32(lp)) }

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
