//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [brush].
//
// [brush]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// [CreateBrushIndirect] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [CreateBrushIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbrushindirect
func CreateBrushIndirect(lb *LOGBRUSH) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreateBrushIndirect.Addr(),
		uintptr(unsafe.Pointer(lb)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// [CreateHatchBrush] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [CreateHatchBrush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhatchbrush
func CreateHatchBrush(hatch co.HS, color COLORREF) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreateHatchBrush.Addr(),
		uintptr(hatch), uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// [CreatePatternBrush] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [CreatePatternBrush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpatternbrush
func CreatePatternBrush(hBmp HBITMAP) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreatePatternBrush.Addr(),
		uintptr(hBmp))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// [CreateSolidBrush] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [CreateSolidBrush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
func CreateSolidBrush(color COLORREF) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreateSolidBrush.Addr(),
		uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() error {
	return HGDIOBJ(hBrush).DeleteObject()
}

// [GetObject] function.
//
// [GetObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBrush HBRUSH) GetObject(lb *LOGBRUSH) {
	ret, _, err := syscall.SyscallN(proc.GetObject.Addr(),
		uintptr(hBrush), unsafe.Sizeof(*lb), uintptr(unsafe.Pointer(lb)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
