//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded comctl procedures.
var comctlCache [30]*syscall.Proc

// Loads comctl procedures.
func Comctl(procId PROC_COMCTL) uintptr {
	return LoadProc(SYSDLL_comctl32, comctlCache[:], comctlProcStr, uint64(procId)).Addr()
}

type PROC_COMCTL uint64 // Procedure identifiers for comctl.

// Auto-generated comctl procedure identifier: cache index | str start | str past-end.
const (
	PROC_ImageList_DragMove           PROC_COMCTL = 0 | (9 << 16) | (27 << 32)
	PROC_ImageList_DragShowNolock     PROC_COMCTL = 1 | (28 << 16) | (52 << 32)
	PROC_ImageList_DrawIndirect       PROC_COMCTL = 2 | (53 << 16) | (75 << 32)
	PROC_ImageList_EndDrag            PROC_COMCTL = 3 | (76 << 16) | (93 << 32)
	PROC_InitCommonControls           PROC_COMCTL = 4 | (94 << 16) | (112 << 32)
	PROC_InitCommonControlsEx         PROC_COMCTL = 5 | (113 << 16) | (133 << 32)
	PROC_InitMUILanguage              PROC_COMCTL = 6 | (134 << 16) | (149 << 32)
	PROC_TaskDialogIndirect           PROC_COMCTL = 7 | (150 << 16) | (168 << 32)
	PROC_ImageList_Create             PROC_COMCTL = 8 | (183 << 16) | (199 << 32)
	PROC_ImageList_Add                PROC_COMCTL = 9 | (200 << 16) | (213 << 32)
	PROC_ImageList_AddMasked          PROC_COMCTL = 10 | (214 << 16) | (233 << 32)
	PROC_ImageList_BeginDrag          PROC_COMCTL = 11 | (234 << 16) | (253 << 32)
	PROC_ImageList_Destroy            PROC_COMCTL = 12 | (254 << 16) | (271 << 32)
	PROC_ImageList_DrawEx             PROC_COMCTL = 13 | (272 << 16) | (288 << 32)
	PROC_ImageList_Duplicate          PROC_COMCTL = 14 | (289 << 16) | (308 << 32)
	PROC_ImageList_GetBkColor         PROC_COMCTL = 15 | (309 << 16) | (329 << 32)
	PROC_ImageList_GetIconSize        PROC_COMCTL = 16 | (330 << 16) | (351 << 32)
	PROC_ImageList_GetImageCount      PROC_COMCTL = 17 | (352 << 16) | (375 << 32)
	PROC_ImageList_GetImageInfo       PROC_COMCTL = 18 | (376 << 16) | (398 << 32)
	PROC_ImageList_Remove             PROC_COMCTL = 19 | (399 << 16) | (415 << 32)
	PROC_ImageList_ReplaceIcon        PROC_COMCTL = 20 | (416 << 16) | (437 << 32)
	PROC_ImageList_SetDragCursorImage PROC_COMCTL = 21 | (438 << 16) | (466 << 32)
	PROC_ImageList_SetIconSize        PROC_COMCTL = 22 | (467 << 16) | (488 << 32)
	PROC_ImageList_SetImageCount      PROC_COMCTL = 23 | (489 << 16) | (512 << 32)
	PROC_ImageList_SetOverlayImage    PROC_COMCTL = 24 | (513 << 16) | (538 << 32)
	PROC_DefSubclassProc              PROC_COMCTL = 25 | (547 << 16) | (562 << 32)
	PROC_ImageList_DragEnter          PROC_COMCTL = 26 | (563 << 16) | (582 << 32)
	PROC_ImageList_DragLeave          PROC_COMCTL = 27 | (583 << 16) | (602 << 32)
	PROC_RemoveWindowSubclass         PROC_COMCTL = 28 | (603 << 16) | (623 << 32)
	PROC_SetWindowSubclass            PROC_COMCTL = 29 | (624 << 16) | (641 << 32)
)

// Declaration of comctl procedure names.
const comctlProcStr = `
--funcs
ImageList_DragMove
ImageList_DragShowNolock
ImageList_DrawIndirect
ImageList_EndDrag
InitCommonControls
InitCommonControlsEx
InitMUILanguage
TaskDialogIndirect

--himagelist
ImageList_Create
ImageList_Add
ImageList_AddMasked
ImageList_BeginDrag
ImageList_Destroy
ImageList_DrawEx
ImageList_Duplicate
ImageList_GetBkColor
ImageList_GetIconSize
ImageList_GetImageCount
ImageList_GetImageInfo
ImageList_Remove
ImageList_ReplaceIcon
ImageList_SetDragCursorImage
ImageList_SetIconSize
ImageList_SetImageCount
ImageList_SetOverlayImage

--hwnd
DefSubclassProc
ImageList_DragEnter
ImageList_DragLeave
RemoveWindowSubclass
SetWindowSubclass
`
