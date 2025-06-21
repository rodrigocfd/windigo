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
		dll.Load(dll.UXTHEME, &_IsAppThemed, "IsAppThemed"))
	return ret != 0
}

var _IsAppThemed *syscall.Proc

// [IsCompositionActive] function.
//
// [IsCompositionActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-iscompositionactive
func IsCompositionActive() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_IsCompositionActive, "IsCompositionActive"))
	return ret != 0
}

var _IsCompositionActive *syscall.Proc

// [IsThemeActive] function.
//
// [IsThemeActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_IsThemeActive, "IsThemeActive"))
	return ret != 0
}

var _IsThemeActive *syscall.Proc
