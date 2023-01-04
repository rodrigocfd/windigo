//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemedialogtextureenabled
func (hWnd HWND) IsThemeDialogTextureEnabled() bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeDialogTextureEnabled.Addr(), 1,
		uintptr(hWnd), 0, 0)
	return ret != 0
}

// ⚠️ You must defer HTHEME.CloseThemeData().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	ret, _, err := syscall.Syscall(proc.OpenThemeData.Addr(), 2,
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(classNames))),
		0)
	if ret == 0 {
		return HTHEME(0), errco.ERROR(err)
	}
	return HTHEME(ret), nil
}
