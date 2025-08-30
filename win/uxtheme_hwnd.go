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
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_IsThemeDialogTextureEnabled, "IsThemeDialogTextureEnabled"),
		uintptr(hWnd))
	return ret != 0
}

var _IsThemeDialogTextureEnabled *syscall.Proc

// [OpenThemeData] function.
//
// ⚠️ You must defer [HTHEME.CloseThemeData].
//
// [OpenThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	var wClassNames wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_OpenThemeData, "OpenThemeData"),
		uintptr(hWnd),
		uintptr(wClassNames.EmptyIsNil(classNames)))
	if ret == 0 {
		return HTHEME(0), co.HRESULT_E_FAIL
	}
	return HTHEME(ret), nil
}

var _OpenThemeData *syscall.Proc
