package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to a GDI object.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
type HGDIOBJ HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() {
	ret, _, lerr := syscall.Syscall(proc.DeleteObject.Addr(), 1,
		uintptr(hGdiObj), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

//------------------------------------------------------------------------------

// A handle to a bitmap.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbitmap
type HBITMAP HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBmp HBITMAP) DeleteObject() {
	HGDIOBJ(hBmp).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a brush.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hbrush
type HBRUSH HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hBrush HBRUSH) DeleteObject() {
	HGDIOBJ(hBrush).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a font.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hfont
type HFONT HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hFont HFONT) DeleteObject() {
	HGDIOBJ(hFont).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a pen.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpen
type HPEN HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPen HPEN) DeleteObject() {
	HGDIOBJ(hPen).DeleteObject()
}

//------------------------------------------------------------------------------

// A handle to a region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
type HRGN HGDIOBJ

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hRgn HRGN) DeleteObject() {
	HGDIOBJ(hRgn).DeleteObject()
}

// Combines the two regions and stores the result in current region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(hrgnSrc1, hrgnSrc2 HRGN, iMode co.RGN) co.REGION {
	ret, _, lerr := syscall.Syscall6(proc.CombineRgn.Addr(), 4,
		uintptr(hRgn), uintptr(hrgnSrc1), uintptr(hrgnSrc2), uintptr(iMode), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
	return co.REGION(ret)
}
