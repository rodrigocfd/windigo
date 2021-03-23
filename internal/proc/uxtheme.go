package proc

import (
	"syscall"
)

var (
	uxtheme = syscall.NewLazyDLL("uxtheme.dll")

	CloseThemeData      = uxtheme.NewProc("CloseThemeData")
	DrawThemeBackground = uxtheme.NewProc("DrawThemeBackground")
	IsAppThemed         = uxtheme.NewProc("IsAppThemed")
	IsThemeActive       = uxtheme.NewProc("IsThemeActive")
	OpenThemeData       = uxtheme.NewProc("OpenThemeData")
)
