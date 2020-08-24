/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/
type HTHEME HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-closethemedata
func (hTheme HTHEME) CloseThemeData() {
	hr := hTheme.closeThemeDataNoPanic()
	if hr != 0 {
		panic(hr.Format("CloseThemeData failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-drawthemebackground
func (hTheme HTHEME) DrawThemeBackground(hdc HDC,
	partId co.VS_PART, stateId co.VS_STATE,
	rect *RECT, clipRect *RECT) {

	hr, _, _ := syscall.Syscall6(proc.DrawThemeBackground.Addr(), 6,
		uintptr(hTheme), uintptr(hdc), uintptr(partId), uintptr(stateId),
		uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(clipRect)))
	if hr != 0 {
		hTheme.closeThemeDataNoPanic() // free resource
		panic(co.ERROR(hr).Format("DrawThemeBackground failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isappthemed
func IsAppThemed() bool {
	ret, _, _ := syscall.Syscall(proc.IsAppThemed.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeActive.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

func (hTheme HTHEME) closeThemeDataNoPanic() co.ERROR {
	if hTheme == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(proc.CloseThemeData.Addr(), 1,
		uintptr(hTheme), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}
