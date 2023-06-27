//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
)

// [IsAppThemed] function.
//
// [IsAppThemed]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isappthemed
func IsAppThemed() bool {
	ret, _, _ := syscall.SyscallN(proc.IsAppThemed.Addr())
	return ret != 0
}

// [IsCompositionActive] function.
//
// [IsCompositionActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-iscompositionactive
func IsCompositionActive() bool {
	ret, _, _ := syscall.SyscallN(proc.IsCompositionActive.Addr())
	return ret != 0
}

// [IsThemeActive] function.
//
// [IsThemeActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.SyscallN(proc.IsThemeActive.Addr())
	return ret != 0
}
