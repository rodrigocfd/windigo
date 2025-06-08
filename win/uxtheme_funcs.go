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
	ret, _, _ := syscall.SyscallN(dll.Uxtheme(dll.PROC_IsAppThemed))
	return ret != 0
}

// [IsCompositionActive] function.
//
// [IsCompositionActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-iscompositionactive
func IsCompositionActive() bool {
	ret, _, _ := syscall.SyscallN(dll.Uxtheme(dll.PROC_IsCompositionActive))
	return ret != 0
}

// [IsThemeActive] function.
//
// [IsThemeActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.SyscallN(dll.Uxtheme(dll.PROC_IsThemeActive))
	return ret != 0
}
