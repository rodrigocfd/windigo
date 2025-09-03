//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a
// [palette](https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpalette).
type HPALETTE HANDLE

// [AnimatePalette] function.
//
// [AnimatePalette]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-animatepalette
func (hPal HPALETTE) AnimatePalette(startIndex int, entries []PALETTEENTRY) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.GDI32, &_AnimatePalette, "AnimatePalette"),
		uintptr(hPal),
		uintptr(uint32(startIndex)),
		uintptr(uint32(len(entries))),
		uintptr(unsafe.Pointer(&entries[0])))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _AnimatePalette *syscall.Proc

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPal HPALETTE) DeleteObject() error {
	return HGDIOBJ(hPal).DeleteObject()
}
