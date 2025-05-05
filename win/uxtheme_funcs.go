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
	ret, _, _ := syscall.SyscallN(_IsAppThemed.Addr())
	return ret != 0
}

var _IsAppThemed = dll.Uxtheme.NewProc("IsAppThemed")

// [IsCompositionActive] function.
//
// [IsCompositionActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-iscompositionactive
func IsCompositionActive() bool {
	ret, _, _ := syscall.SyscallN(_IsCompositionActive.Addr())
	return ret != 0
}

var _IsCompositionActive = dll.Uxtheme.NewProc("IsCompositionActive")

// [IsThemeActive] function.
//
// [IsThemeActive]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemeactive
func IsThemeActive() bool {
	ret, _, _ := syscall.SyscallN(_IsThemeActive.Addr())
	return ret != 0
}

var _IsThemeActive = dll.Uxtheme.NewProc("IsThemeActive")
