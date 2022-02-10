package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a pen.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
func CreatePen(style co.PS, width int32, color COLORREF) HPEN {
	ret, _, err := syscall.Syscall(proc.CreatePen.Addr(), 3,
		uintptr(style), uintptr(width), uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpenindirect
func CreatePenIndirect(lp *LOGPEN) HPEN {
	ret, _, err := syscall.Syscall(proc.CreatePenIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lp)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() {
	HGDIOBJ(hPen).DeleteObject()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hPen HPEN) GetObject(lp *LOGPEN) {
	ret, _, err := syscall.Syscall(proc.GetObject.Addr(), 3,
		uintptr(hPen), unsafe.Sizeof(*lp), uintptr(unsafe.Pointer(lp)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
