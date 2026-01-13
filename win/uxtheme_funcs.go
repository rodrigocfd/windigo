//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// [IsAppThemed] function.
//
// [IsAppThemed]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isappthemed
func IsAppThemed() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Uxtheme.Load(&_uxtheme_IsAppThemed, "IsAppThemed"))
	return ret != 0
}

var _uxtheme_IsAppThemed *syscall.Proc

// [IsCompositionActive] function.
//
// [IsCompositionActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-iscompositionactive
func IsCompositionActive() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Uxtheme.Load(&_uxtheme_IsCompositionActive, "IsCompositionActive"))
	return ret != 0
}

var _uxtheme_IsCompositionActive *syscall.Proc

// [IsThemeActive] function.
//
// [IsThemeActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Uxtheme.Load(&_uxtheme_IsThemeActive, "IsThemeActive"))
	return ret != 0
}

var _uxtheme_IsThemeActive *syscall.Proc
