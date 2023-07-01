//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a [theme].
//
// [theme]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/
type HTHEME HANDLE

// [CloseThemeData] function.
//
// [CloseThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-closethemedata
func (hTheme HTHEME) CloseThemeData() {
	syscall.SyscallN(proc.CloseThemeData.Addr(),
		uintptr(hTheme))
}

// [DrawThemeBackground] function.
//
// [DrawThemeBackground]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-drawthemebackground
func (hTheme HTHEME) DrawThemeBackground(
	hdc HDC, partStateId co.VS, rc *RECT, clipRc *RECT) {

	ret, _, _ := syscall.SyscallN(proc.DrawThemeBackground.Addr(),
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(clipRc)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// [GetThemeColor] function.
//
// [GetThemeColor]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemecolor
func (hTheme HTHEME) GetThemeColor(partStateId co.VS, propId co.TMT) COLORREF {
	var color COLORREF
	ret, _, _ := syscall.SyscallN(proc.GetThemeColor.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&color)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return color
}

// [GetThemeInt] function.
//
// [GetThemeInt]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeint
func (hTheme HTHEME) GetThemeInt(partStateId co.VS, propId co.TMT) int32 {
	var intVal int32
	ret, _, _ := syscall.SyscallN(proc.GetThemeInt.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&intVal)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return intVal
}

// [GetThemeMetric] function.
//
// [GetThemeMetric]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthememetric
func (hTheme HTHEME) GetThemeMetric(
	hdc HDC, partStateId co.VS, propId co.TMT) int32 {

	var intVal int32
	ret, _, _ := syscall.SyscallN(proc.GetThemeMetric.Addr(),
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&intVal)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return intVal
}

// [GetThemePosition] function.
//
// [GetThemePosition]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeposition
func (hTheme HTHEME) GetThemePosition(partStateId co.VS, propId co.TMT) POINT {
	var pt POINT
	ret, _, _ := syscall.SyscallN(proc.GetThemePosition.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&pt)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return pt
}

// [GetThemePropertyOrigin] function.
//
// [GetThemePropertyOrigin]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemepropertyorigin
func (hTheme HTHEME) GetThemePropertyOrigin(
	partStateId co.VS, propId co.TMT) co.PROPERTYORIGIN {

	var origin co.PROPERTYORIGIN
	ret, _, _ := syscall.SyscallN(proc.GetThemePropertyOrigin.Addr(),
		uintptr(hTheme),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&origin)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return origin
}

// [GetThemeRect] function.
//
// [GetThemeRect]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemerect
func (hTheme HTHEME) GetThemeRect(partStateId co.VS, propId co.TMT) RECT {
	var rc RECT
	ret, _, _ := syscall.SyscallN(proc.GetThemeRect.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&rc)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return rc
}

// [GetThemeString] function.
//
// [GetThemeString]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemestring
func (hTheme HTHEME) GetThemeString(partStateId co.VS, propId co.TMT) string {
	var buf [256 + 1]uint16 // arbitrary
	ret, _, _ := syscall.SyscallN(proc.GetThemeRect.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return Str.FromNativeSlice(buf[:])
}

// [GetThemeSysColorBrush] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [GetThemeSysColorBrush]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesyscolorbrush
func (hTheme HTHEME) GetThemeSysColorBrush(colorId co.TMT) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.GetThemeSysColorBrush.Addr(),
		uintptr(hTheme), uintptr(colorId))
	if ret == 0 {
		panic(errco.ERROR(err)) // uncertain?
	}
	return HBRUSH(ret)
}

// [GetThemeSysFont] function.
//
// [GetThemeSysFont]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesysfont
func (hTheme HTHEME) GetThemeSysFont(fontId co.TMT, lf *LOGFONT) {
	ret, _, _ := syscall.SyscallN(proc.GetThemeSysFont.Addr(),
		uintptr(hTheme), uintptr(fontId), uintptr(unsafe.Pointer(lf)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// [GetThemeTextMetrics] function.
//
// [GetThemeTextMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemetextmetrics
func (hTheme HTHEME) GetThemeTextMetrics(
	hdc HDC, partStateId co.VS, tm *TEXTMETRIC) {

	ret, _, _ := syscall.SyscallN(proc.GetThemeTextMetrics.Addr(),
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(tm)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// [IsThemeBackgroundPartiallyTransparent] function.
//
// [IsThemeBackgroundPartiallyTransparent]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemebackgroundpartiallytransparent
func (hTheme HTHEME) IsThemeBackgroundPartiallyTransparent(partStateId co.VS) bool {
	ret, _, _ := syscall.SyscallN(proc.IsThemeBackgroundPartiallyTransparent.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()))
	return ret != 0
}

// [IsThemePartDefined] function.
//
// [IsThemePartDefined]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemepartdefined
func (hTheme HTHEME) IsThemePartDefined(partStateId co.VS) bool {
	ret, _, _ := syscall.SyscallN(proc.IsThemePartDefined.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()))
	return ret != 0
}
