//go:build windows

package proc

import (
	"syscall"
)

var (
	uxtheme = syscall.NewLazyDLL("uxtheme.dll")

	CloseThemeData                        = uxtheme.NewProc("CloseThemeData")
	DrawThemeBackground                   = uxtheme.NewProc("DrawThemeBackground")
	GetThemeColor                         = uxtheme.NewProc("GetThemeColor")
	GetThemeInt                           = uxtheme.NewProc("GetThemeInt")
	GetThemeMetric                        = uxtheme.NewProc("GetThemeMetric")
	GetThemePosition                      = uxtheme.NewProc("GetThemePosition")
	GetThemePropertyOrigin                = uxtheme.NewProc("GetThemePropertyOrigin")
	GetThemeRect                          = uxtheme.NewProc("GetThemeRect")
	GetThemeString                        = uxtheme.NewProc("GetThemeString")
	GetThemeSysColorBrush                 = uxtheme.NewProc("GetThemeSysColorBrush")
	GetThemeSysFont                       = uxtheme.NewProc("GetThemeSysFont")
	GetThemeTextMetrics                   = uxtheme.NewProc("GetThemeTextMetrics")
	IsAppThemed                           = uxtheme.NewProc("IsAppThemed")
	IsCompositionActive                   = uxtheme.NewProc("IsCompositionActive")
	IsThemeActive                         = uxtheme.NewProc("IsThemeActive")
	IsThemeBackgroundPartiallyTransparent = uxtheme.NewProc("IsThemeBackgroundPartiallyTransparent")
	IsThemeDialogTextureEnabled           = uxtheme.NewProc("IsThemeDialogTextureEnabled")
	IsThemePartDefined                    = uxtheme.NewProc("IsThemePartDefined")
	OpenThemeData                         = uxtheme.NewProc("OpenThemeData")
)
