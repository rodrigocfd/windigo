package procs

import (
	"syscall"
)

var (
	dllUxTheme = syscall.NewLazyDLL("uxtheme.dll")

	CloseThemeData      = dllUxTheme.NewProc("CloseThemeData")
	DrawThemeBackground = dllUxTheme.NewProc("DrawThemeBackground")
	IsAppThemed         = dllUxTheme.NewProc("IsAppThemed")
	IsThemeActive       = dllUxTheme.NewProc("IsThemeActive")
	OpenThemeData       = dllUxTheme.NewProc("OpenThemeData")
)
