package api

import (
	"fmt"
	"gowinui/api/proc"
	c "gowinui/consts"
	"syscall"
	"unsafe"
)

type HTHEME HANDLE

func (hTheme HTHEME) CloseThemeData() {
	hr, _, _ := syscall.Syscall(proc.CloseThemeData.Addr(), 0,
		uintptr(hTheme), 0, 0)
	if hr != 0 {
		panic(fmt.Sprintf("CloseThemeData failed: %d.\n", hr))
	}
}

func (hTheme HTHEME) DrawThemeBackground(hdc HDC,
	partId c.VS_STYLE_PART, stateId c.VS_STYLE_STATE,
	rect *RECT, clipRect *RECT) {

	hr, _, _ := syscall.Syscall6(proc.DrawThemeBackground.Addr(), 6,
		uintptr(hTheme), uintptr(hdc), uintptr(partId), uintptr(stateId),
		uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(clipRect)))
	if hr != 0 {
		panic(fmt.Sprintf("DrawThemeBackground failed: %d.\n", hr))
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

func (hWnd HWND) OpenThemeData(classNames string) HTHEME {
	ret, _, _ := syscall.Syscall(proc.OpenThemeData.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(StrToUtf16Ptr(classNames))),
		0)
	return HTHEME(ret) // zero if no match, never fails
}
