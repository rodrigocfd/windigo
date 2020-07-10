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

type HTHEME HANDLE

func (hTheme HTHEME) CloseThemeData() {
	hr := hTheme.closeThemeDataNoPanic()
	if hr != 0 {
		panic(hr.Format("CloseThemeData failed."))
	}
}

func (hTheme HTHEME) closeThemeDataNoPanic() co.ERROR {
	return freeNoPanic(HANDLE(hTheme), proc.CloseThemeData)
}

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

func IsAppThemed() bool {
	ret, _, _ := syscall.Syscall(proc.IsAppThemed.Addr(), 0,
		0, 0, 0)
	return ret != 0
}

func IsThemeActive() bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeActive.Addr(), 0,
		0, 0, 0)
	return ret != 0
}
