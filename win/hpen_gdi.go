//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [pen].
//
// [pen]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HGDIOBJ

// [CreatePen] function.
//
// ⚠️ You must defer HPEN.DeleteObject().
//
// [CreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
func CreatePen(style co.PS, width int32, color COLORREF) HPEN {
	ret, _, err := syscall.SyscallN(proc.CreatePen.Addr(),
		uintptr(style), uintptr(width), uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// [CreatePenIndirect] function.
//
// [CreatePenIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(lp *LOGPEN) HPEN {
	ret, _, err := syscall.SyscallN(proc.CreatePenIndirect.Addr(),
		uintptr(unsafe.Pointer(lp)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() error {
	return HGDIOBJ(hPen).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hPen HPEN) GetObject(lp *LOGPEN) {
	ret, _, err := syscall.SyscallN(proc.GetObject.Addr(),
		uintptr(hPen), unsafe.Sizeof(*lp), uintptr(unsafe.Pointer(lp)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
