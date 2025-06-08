//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded gdi procedures.
var gdiCache [118]*syscall.Proc

// Loads gdi procedures.
func Gdi(procId PROC_GDI) uintptr {
	return LoadProc(SYSDLL_gdi32, gdiCache[:], gdiProcStr, uint64(procId)).Addr()
}

type PROC_GDI uint64 // Procedure identifiers for gdi.

// Auto-generated gdi procedure identifier: cache index | str start | str past-end.
const (
	PROC_AddFontResourceExW     PROC_GDI = 0 | (9 << 16) | (27 << 32)
	PROC_GdiFlush               PROC_GDI = 1 | (28 << 16) | (36 << 32)
	PROC_CreateBitmap           PROC_GDI = 2 | (48 << 16) | (60 << 32)
	PROC_CreateBitmapIndirect   PROC_GDI = 3 | (61 << 16) | (81 << 32)
	PROC_CreateBrushIndirect    PROC_GDI = 4 | (92 << 16) | (111 << 32)
	PROC_CreatePatternBrush     PROC_GDI = 5 | (112 << 16) | (130 << 32)
	PROC_CreateDCW              PROC_GDI = 6 | (138 << 16) | (147 << 32)
	PROC_CreateICW              PROC_GDI = 7 | (148 << 16) | (157 << 32)
	PROC_AbortDoc               PROC_GDI = 8 | (158 << 16) | (166 << 32)
	PROC_AbortPath              PROC_GDI = 9 | (167 << 16) | (176 << 32)
	PROC_AlphaBlend             PROC_GDI = 10 | (177 << 16) | (187 << 32)
	PROC_AngleArc               PROC_GDI = 11 | (188 << 16) | (196 << 32)
	PROC_Arc                    PROC_GDI = 12 | (197 << 16) | (200 << 32)
	PROC_ArcTo                  PROC_GDI = 13 | (201 << 16) | (206 << 32)
	PROC_BeginPath              PROC_GDI = 14 | (207 << 16) | (216 << 32)
	PROC_BitBlt                 PROC_GDI = 15 | (217 << 16) | (223 << 32)
	PROC_CancelDC               PROC_GDI = 16 | (224 << 16) | (232 << 32)
	PROC_Chord                  PROC_GDI = 17 | (233 << 16) | (238 << 32)
	PROC_CloseFigure            PROC_GDI = 18 | (239 << 16) | (250 << 32)
	PROC_ChoosePixelFormat      PROC_GDI = 19 | (251 << 16) | (268 << 32)
	PROC_CreateCompatibleBitmap PROC_GDI = 20 | (269 << 16) | (291 << 32)
	PROC_CreateCompatibleDC     PROC_GDI = 21 | (292 << 16) | (310 << 32)
	PROC_CreateDIBSection       PROC_GDI = 22 | (311 << 16) | (327 << 32)
	PROC_CreateHalftonePalette  PROC_GDI = 23 | (328 << 16) | (349 << 32)
	PROC_DeleteDC               PROC_GDI = 24 | (350 << 16) | (358 << 32)
	PROC_Ellipse                PROC_GDI = 25 | (359 << 16) | (366 << 32)
	PROC_EndDoc                 PROC_GDI = 26 | (367 << 16) | (373 << 32)
	PROC_EndPage                PROC_GDI = 27 | (374 << 16) | (381 << 32)
	PROC_EndPath                PROC_GDI = 28 | (382 << 16) | (389 << 32)
	PROC_ExcludeClipRect        PROC_GDI = 29 | (390 << 16) | (405 << 32)
	PROC_FillPath               PROC_GDI = 30 | (406 << 16) | (414 << 32)
	PROC_FillRect               PROC_GDI = 31 | (415 << 16) | (423 << 32)
	PROC_FillRgn                PROC_GDI = 32 | (424 << 16) | (431 << 32)
	PROC_FlattenPath            PROC_GDI = 33 | (432 << 16) | (443 << 32)
	PROC_FrameRgn               PROC_GDI = 34 | (444 << 16) | (452 << 32)
	PROC_GetBkColor             PROC_GDI = 35 | (453 << 16) | (463 << 32)
	PROC_GetBkMode              PROC_GDI = 36 | (464 << 16) | (473 << 32)
	PROC_GetCurrentPositionEx   PROC_GDI = 37 | (474 << 16) | (494 << 32)
	PROC_GetDCBrushColor        PROC_GDI = 38 | (495 << 16) | (510 << 32)
	PROC_GetDCPenColor          PROC_GDI = 39 | (511 << 16) | (524 << 32)
	PROC_GetDeviceCaps          PROC_GDI = 40 | (525 << 16) | (538 << 32)
	PROC_GetDIBits              PROC_GDI = 41 | (539 << 16) | (548 << 32)
	PROC_GetPixel               PROC_GDI = 42 | (549 << 16) | (557 << 32)
	PROC_GetPolyFillMode        PROC_GDI = 43 | (558 << 16) | (573 << 32)
	PROC_GetTextColor           PROC_GDI = 44 | (574 << 16) | (586 << 32)
	PROC_GetTextExtentPoint32W  PROC_GDI = 45 | (587 << 16) | (608 << 32)
	PROC_GetTextFaceW           PROC_GDI = 46 | (609 << 16) | (621 << 32)
	PROC_GetTextMetricsW        PROC_GDI = 47 | (622 << 16) | (637 << 32)
	PROC_GetViewportExtEx       PROC_GDI = 48 | (638 << 16) | (654 << 32)
	PROC_GetViewportOrgEx       PROC_GDI = 49 | (655 << 16) | (671 << 32)
	PROC_GetWindowExtEx         PROC_GDI = 50 | (672 << 16) | (686 << 32)
	PROC_GetWindowOrgEx         PROC_GDI = 51 | (687 << 16) | (701 << 32)
	PROC_GradientFill           PROC_GDI = 52 | (702 << 16) | (714 << 32)
	PROC_IntersectClipRect      PROC_GDI = 53 | (715 << 16) | (732 << 32)
	PROC_InvertRgn              PROC_GDI = 54 | (733 << 16) | (742 << 32)
	PROC_LineTo                 PROC_GDI = 55 | (743 << 16) | (749 << 32)
	PROC_LPtoDP                 PROC_GDI = 56 | (750 << 16) | (756 << 32)
	PROC_MaskBlt                PROC_GDI = 57 | (757 << 16) | (764 << 32)
	PROC_MoveToEx               PROC_GDI = 58 | (765 << 16) | (773 << 32)
	PROC_PaintRgn               PROC_GDI = 59 | (774 << 16) | (782 << 32)
	PROC_PatBlt                 PROC_GDI = 60 | (783 << 16) | (789 << 32)
	PROC_PathToRegion           PROC_GDI = 61 | (790 << 16) | (802 << 32)
	PROC_Pie                    PROC_GDI = 62 | (803 << 16) | (806 << 32)
	PROC_PolyBezier             PROC_GDI = 63 | (807 << 16) | (817 << 32)
	PROC_PolyBezierTo           PROC_GDI = 64 | (818 << 16) | (830 << 32)
	PROC_PolyDraw               PROC_GDI = 65 | (831 << 16) | (839 << 32)
	PROC_Polygon                PROC_GDI = 66 | (840 << 16) | (847 << 32)
	PROC_Polyline               PROC_GDI = 67 | (848 << 16) | (856 << 32)
	PROC_PolylineTo             PROC_GDI = 68 | (857 << 16) | (867 << 32)
	PROC_PolyPolygon            PROC_GDI = 69 | (868 << 16) | (879 << 32)
	PROC_PolyPolyline           PROC_GDI = 70 | (880 << 16) | (892 << 32)
	PROC_PtVisible              PROC_GDI = 71 | (893 << 16) | (902 << 32)
	PROC_RealizePalette         PROC_GDI = 72 | (903 << 16) | (917 << 32)
	PROC_Rectangle              PROC_GDI = 73 | (918 << 16) | (927 << 32)
	PROC_ResetDCW               PROC_GDI = 74 | (928 << 16) | (936 << 32)
	PROC_RestoreDC              PROC_GDI = 75 | (937 << 16) | (946 << 32)
	PROC_RoundRect              PROC_GDI = 76 | (947 << 16) | (956 << 32)
	PROC_SaveDC                 PROC_GDI = 77 | (957 << 16) | (963 << 32)
	PROC_SelectClipPath         PROC_GDI = 78 | (964 << 16) | (978 << 32)
	PROC_SelectClipRgn          PROC_GDI = 79 | (979 << 16) | (992 << 32)
	PROC_SelectPalette          PROC_GDI = 80 | (993 << 16) | (1006 << 32)
	PROC_SetArcDirection        PROC_GDI = 81 | (1007 << 16) | (1022 << 32)
	PROC_SetBkColor             PROC_GDI = 82 | (1023 << 16) | (1033 << 32)
	PROC_SetBkMode              PROC_GDI = 83 | (1034 << 16) | (1043 << 32)
	PROC_SetBrushOrgEx          PROC_GDI = 84 | (1044 << 16) | (1057 << 32)
	PROC_SetPixel               PROC_GDI = 85 | (1058 << 16) | (1066 << 32)
	PROC_SetPixelFormat         PROC_GDI = 86 | (1067 << 16) | (1081 << 32)
	PROC_SetPolyFillMode        PROC_GDI = 87 | (1082 << 16) | (1097 << 32)
	PROC_SetStretchBltMode      PROC_GDI = 88 | (1098 << 16) | (1115 << 32)
	PROC_SetTextAlign           PROC_GDI = 89 | (1116 << 16) | (1128 << 32)
	PROC_SetTextColor           PROC_GDI = 90 | (1129 << 16) | (1141 << 32)
	PROC_SetViewportExtEx       PROC_GDI = 91 | (1142 << 16) | (1158 << 32)
	PROC_StartDocW              PROC_GDI = 92 | (1159 << 16) | (1168 << 32)
	PROC_StartPage              PROC_GDI = 93 | (1169 << 16) | (1178 << 32)
	PROC_StretchBlt             PROC_GDI = 94 | (1179 << 16) | (1189 << 32)
	PROC_StrokeAndFillPath      PROC_GDI = 95 | (1190 << 16) | (1207 << 32)
	PROC_StrokePath             PROC_GDI = 96 | (1208 << 16) | (1218 << 32)
	PROC_SwapBuffers            PROC_GDI = 97 | (1219 << 16) | (1230 << 32)
	PROC_TextOutW               PROC_GDI = 98 | (1231 << 16) | (1239 << 32)
	PROC_TransparentBlt         PROC_GDI = 99 | (1240 << 16) | (1254 << 32)
	PROC_WidenPath              PROC_GDI = 100 | (1255 << 16) | (1264 << 32)
	PROC_CreateFontW            PROC_GDI = 101 | (1274 << 16) | (1285 << 32)
	PROC_CreateFontIndirectW    PROC_GDI = 102 | (1286 << 16) | (1305 << 32)
	PROC_GetStockObject         PROC_GDI = 103 | (1317 << 16) | (1331 << 32)
	PROC_DeleteObject           PROC_GDI = 104 | (1332 << 16) | (1344 << 32)
	PROC_GetObjectW             PROC_GDI = 105 | (1345 << 16) | (1355 << 32)
	PROC_SelectObject           PROC_GDI = 106 | (1356 << 16) | (1368 << 32)
	PROC_CreatePen              PROC_GDI = 107 | (1377 << 16) | (1386 << 32)
	PROC_CreatePenIndirect      PROC_GDI = 108 | (1387 << 16) | (1404 << 32)
	PROC_ExtCreatePen           PROC_GDI = 109 | (1405 << 16) | (1417 << 32)
	PROC_CreateRectRgnIndirect  PROC_GDI = 110 | (1426 << 16) | (1447 << 32)
	PROC_CombineRgn             PROC_GDI = 111 | (1448 << 16) | (1458 << 32)
	PROC_EqualRgn               PROC_GDI = 112 | (1459 << 16) | (1467 << 32)
	PROC_GetRgnBox              PROC_GDI = 113 | (1468 << 16) | (1477 << 32)
	PROC_OffsetClipRgn          PROC_GDI = 114 | (1478 << 16) | (1491 << 32)
	PROC_OffsetRgn              PROC_GDI = 115 | (1492 << 16) | (1501 << 32)
	PROC_PtInRegion             PROC_GDI = 116 | (1502 << 16) | (1512 << 32)
	PROC_RectInRegion           PROC_GDI = 117 | (1513 << 16) | (1525 << 32)
)

// Declaration of gdi procedure names.
const gdiProcStr = `
--funcs
AddFontResourceExW
GdiFlush

--hbitmap
CreateBitmap
CreateBitmapIndirect

--hbrush
CreateBrushIndirect
CreatePatternBrush

--hdc
CreateDCW
CreateICW
AbortDoc
AbortPath
AlphaBlend
AngleArc
Arc
ArcTo
BeginPath
BitBlt
CancelDC
Chord
CloseFigure
ChoosePixelFormat
CreateCompatibleBitmap
CreateCompatibleDC
CreateDIBSection
CreateHalftonePalette
DeleteDC
Ellipse
EndDoc
EndPage
EndPath
ExcludeClipRect
FillPath
FillRect
FillRgn
FlattenPath
FrameRgn
GetBkColor
GetBkMode
GetCurrentPositionEx
GetDCBrushColor
GetDCPenColor
GetDeviceCaps
GetDIBits
GetPixel
GetPolyFillMode
GetTextColor
GetTextExtentPoint32W
GetTextFaceW
GetTextMetricsW
GetViewportExtEx
GetViewportOrgEx
GetWindowExtEx
GetWindowOrgEx
GradientFill
IntersectClipRect
InvertRgn
LineTo
LPtoDP
MaskBlt
MoveToEx
PaintRgn
PatBlt
PathToRegion
Pie
PolyBezier
PolyBezierTo
PolyDraw
Polygon
Polyline
PolylineTo
PolyPolygon
PolyPolyline
PtVisible
RealizePalette
Rectangle
ResetDCW
RestoreDC
RoundRect
SaveDC
SelectClipPath
SelectClipRgn
SelectPalette
SetArcDirection
SetBkColor
SetBkMode
SetBrushOrgEx
SetPixel
SetPixelFormat
SetPolyFillMode
SetStretchBltMode
SetTextAlign
SetTextColor
SetViewportExtEx
StartDocW
StartPage
StretchBlt
StrokeAndFillPath
StrokePath
SwapBuffers
TextOutW
TransparentBlt
WidenPath

--hfont
CreateFontW
CreateFontIndirectW

--hgdiobj
GetStockObject
DeleteObject
GetObjectW
SelectObject

--hpen
CreatePen
CreatePenIndirect
ExtCreatePen

--hrgn
CreateRectRgnIndirect
CombineRgn
EqualRgn
GetRgnBox
OffsetClipRgn
OffsetRgn
PtInRegion
RectInRegion
`
