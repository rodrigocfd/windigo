//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IsThemeDialogTextureEnabled] function.
//
// [IsThemeDialogTextureEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemedialogtextureenabled
func (hWnd HWND) IsThemeDialogTextureEnabled() bool {
	ret, _, _ := syscall.SyscallN(proc.IsThemeDialogTextureEnabled.Addr(),
		uintptr(hWnd))
	return ret != 0
}

// [OpenThemeData] function.
//
// ⚠️ You must defer HTHEME.CloseThemeData().
//
// [OpenThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	ret, _, err := syscall.SyscallN(proc.OpenThemeData.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(Str.ToNativePtr(classNames))))
	if ret == 0 {
		return HTHEME(0), errco.ERROR(err)
	}
	return HTHEME(ret), nil
}
