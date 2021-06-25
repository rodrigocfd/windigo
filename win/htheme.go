package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a theme.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/
type HTHEME HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-closethemedata
func (hTheme HTHEME) CloseThemeData() {
	if hTheme != 0 {
		syscall.Syscall(proc.CloseThemeData.Addr(), 1,
			uintptr(hTheme), 0, 0)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-drawthemebackground
func (hTheme HTHEME) DrawThemeBackground(
	hdc HDC, partStateId co.VS, rect *RECT, clipRect *RECT) {

	hr, _, _ := syscall.Syscall6(proc.DrawThemeBackground.Addr(), 6,
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(clipRect)))
	if errco.ERROR(hr) != errco.S_OK {
		panic(errco.ERROR(hr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeposition
func (hTheme HTHEME) GetThemePosition(
	iPartStateId co.VS, iPropId co.TMT) POINT {

	pPoint := POINT{}
	hr, _, _ := syscall.Syscall6(proc.GetThemePosition.Addr(), 5,
		uintptr(hTheme), uintptr(iPartStateId.Part()), uintptr(iPartStateId.State()),
		uintptr(iPropId), uintptr(unsafe.Pointer(&pPoint)), 0)
	if errco.ERROR(hr) != errco.S_OK {
		panic(errco.ERROR(hr))
	}
	return pPoint
}

// ⚠️ You must defer DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesyscolorbrush
func (hTheme HTHEME) GetThemeSysColorBrush(iColorId co.TMT) HBRUSH {
	ret, _, err := syscall.Syscall(proc.GetThemeSysColorBrush.Addr(), 2,
		uintptr(hTheme), uintptr(iColorId), 0)
	if ret == 0 {
		panic(errco.ERROR(err)) // uncertain?
	}
	return HBRUSH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesysfont
func (hTheme HTHEME) GetThemeSysFont(iFontId co.TMT, plf *LOGFONT) {
	hr, _, _ := syscall.Syscall(proc.GetThemeSysFont.Addr(), 3,
		uintptr(hTheme), uintptr(iFontId), uintptr(unsafe.Pointer(plf)))
	if errco.ERROR(hr) != errco.S_OK {
		panic(errco.ERROR(hr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemetextmetrics
func (hTheme HTHEME) GetThemeTextMetrics(
	hdc HDC, iPartStateId co.VS, ptm *TEXTMETRIC) {

	hr, _, _ := syscall.Syscall6(proc.GetThemeTextMetrics.Addr(), 5,
		uintptr(hTheme), uintptr(hdc),
		uintptr(iPartStateId.Part()), uintptr(iPartStateId.State()),
		uintptr(unsafe.Pointer(ptm)), 0)
	if errco.ERROR(hr) != errco.S_OK {
		panic(errco.ERROR(hr))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemebackgroundpartiallytransparent
func (hTheme HTHEME) IsThemeBackgroundPartiallyTransparent(iPartStateId co.VS) bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeBackgroundPartiallyTransparent.Addr(), 3,
		uintptr(hTheme), uintptr(iPartStateId.Part()), uintptr(iPartStateId.State()))
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemepartdefined
func (hTheme HTHEME) IsThemePartDefined(iPartStateId co.VS) bool {
	ret, _, _ := syscall.Syscall(proc.IsThemePartDefined.Addr(), 3,
		uintptr(hTheme), uintptr(iPartStateId.Part()), uintptr(iPartStateId.State()))
	return ret != 0
}
