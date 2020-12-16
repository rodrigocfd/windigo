/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	uxthemeDll = syscall.NewLazyDLL("uxtheme.dll")

	CloseThemeData      = uxthemeDll.NewProc("CloseThemeData")
	DrawThemeBackground = uxthemeDll.NewProc("DrawThemeBackground")
	IsAppThemed         = uxthemeDll.NewProc("IsAppThemed")
	IsThemeActive       = uxthemeDll.NewProc("IsThemeActive")
	OpenThemeData       = uxthemeDll.NewProc("OpenThemeData")
)
