//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IsThemeDialogTextureEnabled] function.
//
// [IsThemeDialogTextureEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemedialogtextureenabled
func (hWnd HWND) IsThemeDialogTextureEnabled() bool {
	ret, _, _ := syscall.SyscallN(_IsThemeDialogTextureEnabled.Addr(),
		uintptr(hWnd))
	return ret != 0
}

var _IsThemeDialogTextureEnabled = dll.Uxtheme.NewProc("IsThemeDialogTextureEnabled")

// [OpenThemeData] function.
//
// ⚠️ You must defer [HTHEME.CloseThemeData].
//
// [OpenThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	classNames16 := wstr.NewBufWith[wstr.Stack20](classNames, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(_OpenThemeData.Addr(),
		uintptr(hWnd), uintptr(classNames16.UnsafePtr()))
	if ret == 0 {
		return HTHEME(0), co.HRESULT_E_FAIL
	}
	return HTHEME(ret), nil
}

var _OpenThemeData = dll.Uxtheme.NewProc("OpenThemeData")
