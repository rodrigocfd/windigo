//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded uxtheme procedures.
var uxthemeCache [19]*syscall.Proc

// Loads uxtheme procedures.
func Uxtheme(procId PROC_UXTHEME) uintptr {
	return LoadProc(SYSDLL_uxtheme, uxthemeCache[:], uxthemeProcStr, uint64(procId)).Addr()
}

type PROC_UXTHEME uint64 // Procedure identifiers for uxtheme.

// Auto-generated uxtheme procedure identifier: cache index | str start | str past-end.
const (
	PROC_IsAppThemed                           PROC_UXTHEME = 0 | (9 << 16) | (20 << 32)
	PROC_IsCompositionActive                   PROC_UXTHEME = 1 | (21 << 16) | (40 << 32)
	PROC_IsThemeActive                         PROC_UXTHEME = 2 | (41 << 16) | (54 << 32)
	PROC_CloseThemeData                        PROC_UXTHEME = 3 | (65 << 16) | (79 << 32)
	PROC_DrawThemeBackground                   PROC_UXTHEME = 4 | (80 << 16) | (99 << 32)
	PROC_GetThemeColor                         PROC_UXTHEME = 5 | (100 << 16) | (113 << 32)
	PROC_GetThemeInt                           PROC_UXTHEME = 6 | (114 << 16) | (125 << 32)
	PROC_GetThemeMetric                        PROC_UXTHEME = 7 | (126 << 16) | (140 << 32)
	PROC_GetThemePosition                      PROC_UXTHEME = 8 | (141 << 16) | (157 << 32)
	PROC_GetThemePropertyOrigin                PROC_UXTHEME = 9 | (158 << 16) | (180 << 32)
	PROC_GetThemeRect                          PROC_UXTHEME = 10 | (181 << 16) | (193 << 32)
	PROC_GetThemeString                        PROC_UXTHEME = 11 | (194 << 16) | (208 << 32)
	PROC_GetThemeSysColorBrush                 PROC_UXTHEME = 12 | (209 << 16) | (230 << 32)
	PROC_GetThemeSysFont                       PROC_UXTHEME = 13 | (231 << 16) | (246 << 32)
	PROC_GetThemeTextMetrics                   PROC_UXTHEME = 14 | (247 << 16) | (266 << 32)
	PROC_IsThemeBackgroundPartiallyTransparent PROC_UXTHEME = 15 | (267 << 16) | (304 << 32)
	PROC_IsThemePartDefined                    PROC_UXTHEME = 16 | (305 << 16) | (323 << 32)
	PROC_IsThemeDialogTextureEnabled           PROC_UXTHEME = 17 | (332 << 16) | (359 << 32)
	PROC_OpenThemeData                         PROC_UXTHEME = 18 | (360 << 16) | (373 << 32)
)

// Declaration of uxtheme procedure names.
const uxthemeProcStr = `
--funcs
IsAppThemed
IsCompositionActive
IsThemeActive

--htheme
CloseThemeData
DrawThemeBackground
GetThemeColor
GetThemeInt
GetThemeMetric
GetThemePosition
GetThemePropertyOrigin
GetThemeRect
GetThemeString
GetThemeSysColorBrush
GetThemeSysFont
GetThemeTextMetrics
IsThemeBackgroundPartiallyTransparent
IsThemePartDefined

--hwnd
IsThemeDialogTextureEnabled
OpenThemeData
`
