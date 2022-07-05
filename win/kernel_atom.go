//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// An atom.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// ⚠️ You must defer ATOM.GlobalDeleteAtom().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globaladdatomw
func GlobalAddAtom(s string) ATOM {
	ret, _, err := syscall.Syscall(proc.GlobalAddAtom.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToNativePtr(s))),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return ATOM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globaldeleteatom
func (atom ATOM) GlobalDeleteAtom() {
	syscall.Syscall(proc.SetLastError.Addr(), 0,
		0, 0, 0)

	_, _, err := syscall.Syscall(proc.GlobalDeleteAtom.Addr(), 1,
		uintptr(atom), 0, 0)
	if wErr := errco.ERROR(err); wErr != errco.SUCCESS {
		panic(wErr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalgetatomnamew
func (atom ATOM) GlobalGetAtomName() string {
	buf := make([]uint16, _MAX_PATH+1) // arbitrary

	ret, _, err := syscall.Syscall(proc.GlobalGetAtomName.Addr(), 3,
		uintptr(atom), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf)
}
