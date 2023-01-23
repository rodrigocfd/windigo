//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a brush.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createbrushindirect
func CreateBrushIndirect(lb *LOGBRUSH) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreateBrushIndirect.Addr(),
		uintptr(unsafe.Pointer(lb)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhatchbrush
func CreateHatchBrush(hatch co.HS, color COLORREF) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreateHatchBrush.Addr(),
		uintptr(hatch), uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpatternbrush
func CreatePatternBrush(hBmp HBITMAP) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreatePatternBrush.Addr(),
		uintptr(hBmp))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
func CreateSolidBrush(color COLORREF) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.CreateSolidBrush.Addr(),
		uintptr(color))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// Not an actual Win32 function, just a tricky conversion to create a brush from
// a system color, particularly used when registering a window class.
func CreateSysColorBrush(sysColor co.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() error {
	return HGDIOBJ(hBrush).DeleteObject()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hBrush HBRUSH) GetObject(lb *LOGBRUSH) {
	ret, _, err := syscall.SyscallN(proc.GetObject.Addr(),
		uintptr(hBrush), unsafe.Sizeof(*lb), uintptr(unsafe.Pointer(lb)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
