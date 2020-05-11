/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package proc

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
