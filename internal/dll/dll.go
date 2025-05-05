//go:build windows

package dll

import (
	"syscall"
)

var (
	Advapi32 = syscall.NewLazyDLL("advapi32.dll")
	Comctl32 = syscall.NewLazyDLL("comctl32.dll")
	Dwmapi   = syscall.NewLazyDLL("dwmapi.dll")
	Gdi32    = syscall.NewLazyDLL("gdi32.dll")
	Kernel32 = syscall.NewLazyDLL("kernel32.dll")
	Ole32    = syscall.NewLazyDLL("ole32.dll")
	Oleaut32 = syscall.NewLazyDLL("oleaut32.dll")
	Shell32  = syscall.NewLazyDLL("shell32.dll")
	Shlwapi  = syscall.NewLazyDLL("shlwapi")
	User32   = syscall.NewLazyDLL("user32.dll")
	Uxtheme  = syscall.NewLazyDLL("uxtheme.dll")
	Version  = syscall.NewLazyDLL("version.dll")
)
