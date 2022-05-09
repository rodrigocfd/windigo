//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/win/co"
)

// An atom.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// A handle to an object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HANDLE syscall.Handle

// A handle to a tree view control item.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls#parent-and-child-items
type HTREEITEM HANDLE

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

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
