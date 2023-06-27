//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// An [atom].
//
// [atom]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// [GlobalAddAtom] function.
//
// ⚠️ You must defer ATOM.GlobalDeleteAtom().
//
// [GlobalAddAtom]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globaladdatomw
func GlobalAddAtom(s string) ATOM {
	ret, _, err := syscall.SyscallN(proc.GlobalAddAtom.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(s))))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return ATOM(ret)
}

// [GlobalDeleteAtom] function.
//
// [GlobalDeleteAtom]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globaldeleteatom
func (atom ATOM) GlobalDeleteAtom() error {
	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	_, _, err := syscall.SyscallN(proc.GlobalDeleteAtom.Addr(),
		uintptr(atom))
	if wErr := errco.ERROR(err); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// [GlobalGetAtomName] function.
//
// [GlobalGetAtomName]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalgetatomnamew
func (atom ATOM) GlobalGetAtomName() string {
	var buf [_MAX_PATH + 1]uint16 // arbitrary

	ret, _, err := syscall.SyscallN(proc.GlobalGetAtomName.Addr(),
		uintptr(atom), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

//------------------------------------------------------------------------------

// A [handle] to an object. This generic handle is used throughout the whole
// API, with different meanings.
//
// [handle]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HANDLE syscall.Handle

// A handle to an [event].
//
// [event]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-createeventw
type HEVENT HANDLE

// A handle to a [resource].
//
// [resource]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-findresourcew
type HRSRC HANDLE

// A handle to a [resource memory block].
//
// [resource memory block]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadresource
type HRSRCMEM HANDLE

//------------------------------------------------------------------------------

// Language and sublanguage [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifiers
type LANGID uint16

// Predefined language [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifiers
const (
	LANGID_SYSTEM_DEFAULT LANGID = LANGID((uint16(co.SUBLANG_SYS_DEFAULT) << 10) | uint16(co.LANG_NEUTRAL))
	LANGID_USER_DEFAULT   LANGID = LANGID((uint16(co.SUBLANG_DEFAULT) << 10) | uint16(co.LANG_NEUTRAL))
)

// [MAKELANGID] macro.
//
// [MAKELANGID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelangid
func MAKELANGID(lang co.LANG, subLang co.SUBLANG) LANGID {
	return LANGID((uint16(subLang) << 10) | uint16(lang))
}

// [PRIMARYLANGID] macro.
//
// [PRIMARYLANGID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-primarylangid
func (lid LANGID) Lang() co.LANG { return co.LANG(uint16(lid) & 0x3ff) }

// [SUBLANGID] macro.
//
// [SUBLANGID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-sublangid
func (lid LANGID) SubLang() co.SUBLANG { return co.SUBLANG(uint16(lid) >> 10) }

//------------------------------------------------------------------------------

// Locale [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/locale-identifiers
type LCID uint32

// Predefined locale [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/locale-identifiers
const (
	LCID_SYSTEM_DEFAULT LCID = LCID((uint32(co.SORT_DEFAULT) << 16) | uint32(LANGID_SYSTEM_DEFAULT))
	LCID_USER_DEFAULT   LCID = LCID((uint32(co.SORT_DEFAULT) << 16) | uint32(LANGID_USER_DEFAULT))
)

// [MAKELCID] macro.
//
// [MAKELCID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelcid
func MAKELCID(langId LANGID, sortId co.SORT) LCID {
	return LCID((uint32(sortId) << 16) | uint32(langId))
}

// [LANGIDFROMLCID] macro.
//
// [LANGIDFROMLCID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-langidfromlcid
func (lcid LCID) LangId() LANGID { return LANGID(uint16(lcid)) }

// [SORTIDFROMLCID] macro.
//
// [SORTIDFROMLCID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-sortidfromlcid
func (lcid LCID) SortId() co.SORT { return co.SORT((uint32(lcid) >> 16) & 0xf) }
