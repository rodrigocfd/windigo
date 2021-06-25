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
