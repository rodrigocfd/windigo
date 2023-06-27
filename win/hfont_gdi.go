//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [font].
//
// [font]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hfont
type HFONT HGDIOBJ

// [CreateFontIndirect] function.
//
// ⚠️ You must defer HFONT.DeleteObject().
//
// [CreateFontIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) HFONT {
	ret, _, err := syscall.SyscallN(proc.CreateFontIndirect.Addr(),
		uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hFont HFONT) DeleteObject() error {
	return HGDIOBJ(hFont).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hFont HFONT) GetObject(lf *LOGFONT) {
	ret, _, err := syscall.SyscallN(proc.GetObject.Addr(),
		uintptr(hFont), unsafe.Sizeof(*lf), uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
