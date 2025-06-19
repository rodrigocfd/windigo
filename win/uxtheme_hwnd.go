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
	ret, _, _ := syscall.SyscallN(dll.Uxtheme(dll.PROC_IsThemeDialogTextureEnabled),
		uintptr(hWnd))
	return ret != 0
}

// [OpenThemeData] function.
//
// ⚠️ You must defer [HTHEME.CloseThemeData].
//
// [OpenThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hWnd HWND) OpenThemeData(classNames string) (HTHEME, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pClassNames := wbuf.PtrEmptyIsNil(classNames)

	ret, _, _ := syscall.SyscallN(dll.Uxtheme(dll.PROC_OpenThemeData),
		uintptr(hWnd),
		uintptr(pClassNames))
	if ret == 0 {
		return HTHEME(0), co.HRESULT_E_FAIL
	}
	return HTHEME(ret), nil
}
