//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded dwmapi procedures.
var dwmapiCache [10]*syscall.Proc

// Loads dwmapi procedures.
func Dwmapi(procId PROC_DWMAPI) uintptr {
	return LoadProc(SYSDLL_dwmapi, dwmapiCache[:], dwmapiProcStr, uint64(procId)).Addr()
}

type PROC_DWMAPI uint64 // Procedure identifiers for dwmapi.

// Auto-generated dwmapi procedure identifier: cache index | str start | str past-end.
const (
	PROC_DwmEnableMMCSS                PROC_DWMAPI = 0 | (9 << 16) | (23 << 32)
	PROC_DwmFlush                      PROC_DWMAPI = 1 | (24 << 16) | (32 << 32)
	PROC_DwmGetColorizationColor       PROC_DWMAPI = 2 | (33 << 16) | (56 << 32)
	PROC_DwmIsCompositionEnabled       PROC_DWMAPI = 3 | (57 << 16) | (80 << 32)
	PROC_DwmExtendFrameIntoClientArea  PROC_DWMAPI = 4 | (89 << 16) | (117 << 32)
	PROC_DwmGetWindowAttribute         PROC_DWMAPI = 5 | (118 << 16) | (139 << 32)
	PROC_DwmInvalidateIconicBitmaps    PROC_DWMAPI = 6 | (140 << 16) | (166 << 32)
	PROC_DwmSetIconicLivePreviewBitmap PROC_DWMAPI = 7 | (167 << 16) | (196 << 32)
	PROC_DwmSetIconicThumbnail         PROC_DWMAPI = 8 | (197 << 16) | (218 << 32)
	PROC_DwmSetWindowAttribute         PROC_DWMAPI = 9 | (219 << 16) | (240 << 32)
)

// Declaration of dwmapi procedure names.
const dwmapiProcStr = `
--funcs
DwmEnableMMCSS
DwmFlush
DwmGetColorizationColor
DwmIsCompositionEnabled

--hwnd
DwmExtendFrameIntoClientArea
DwmGetWindowAttribute
DwmInvalidateIconicBitmaps
DwmSetIconicLivePreviewBitmap
DwmSetIconicThumbnail
DwmSetWindowAttribute
`
