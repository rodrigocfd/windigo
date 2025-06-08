//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded user procedures.
var userCache [181]*syscall.Proc

// Loads user procedures.
func User(procId PROC_USER) uintptr {
	return LoadProc(SYSDLL_user32, userCache[:], userProcStr, uint64(procId)).Addr()
}

type PROC_USER uint64 // Procedure identifiers for user.

// Auto-generated user procedure identifier: cache index | str start | str past-end.
const (
	PROC_AdjustWindowRectEx         PROC_USER = 0 | (9 << 16) | (27 << 32)
	PROC_AllowSetForegroundWindow   PROC_USER = 1 | (28 << 16) | (52 << 32)
	PROC_AnyPopup                   PROC_USER = 2 | (53 << 16) | (61 << 32)
	PROC_BroadcastSystemMessageW    PROC_USER = 3 | (62 << 16) | (85 << 32)
	PROC_CreateIconFromResourceEx   PROC_USER = 4 | (86 << 16) | (110 << 32)
	PROC_DestroyCaret               PROC_USER = 5 | (111 << 16) | (123 << 32)
	PROC_DispatchMessageW           PROC_USER = 6 | (124 << 16) | (140 << 32)
	PROC_EndMenu                    PROC_USER = 7 | (141 << 16) | (148 << 32)
	PROC_EnumDisplayDevicesW        PROC_USER = 8 | (149 << 16) | (168 << 32)
	PROC_EnumThreadWindows          PROC_USER = 9 | (169 << 16) | (186 << 32)
	PROC_EnumWindows                PROC_USER = 10 | (187 << 16) | (198 << 32)
	PROC_GetAsyncKeyState           PROC_USER = 11 | (199 << 16) | (215 << 32)
	PROC_GetCaretPos                PROC_USER = 12 | (216 << 16) | (227 << 32)
	PROC_GetCursorInfo              PROC_USER = 13 | (228 << 16) | (241 << 32)
	PROC_GetCursorPos               PROC_USER = 14 | (242 << 16) | (254 << 32)
	PROC_GetDialogBaseUnits         PROC_USER = 15 | (255 << 16) | (273 << 32)
	PROC_GetGUIThreadInfo           PROC_USER = 16 | (274 << 16) | (290 << 32)
	PROC_GetInputState              PROC_USER = 17 | (291 << 16) | (304 << 32)
	PROC_GetMessageW                PROC_USER = 18 | (305 << 16) | (316 << 32)
	PROC_GetMessageExtraInfo        PROC_USER = 19 | (317 << 16) | (336 << 32)
	PROC_GetMessagePos              PROC_USER = 20 | (337 << 16) | (350 << 32)
	PROC_GetMessageTime             PROC_USER = 21 | (351 << 16) | (365 << 32)
	PROC_GetPhysicalCursorPos       PROC_USER = 22 | (366 << 16) | (386 << 32)
	PROC_GetProcessDefaultLayout    PROC_USER = 23 | (387 << 16) | (410 << 32)
	PROC_GetQueueStatus             PROC_USER = 24 | (411 << 16) | (425 << 32)
	PROC_GetSysColor                PROC_USER = 25 | (426 << 16) | (437 << 32)
	PROC_GetSystemMetrics           PROC_USER = 26 | (438 << 16) | (454 << 32)
	PROC_InflateRect                PROC_USER = 27 | (455 << 16) | (466 << 32)
	PROC_InSendMessage              PROC_USER = 28 | (467 << 16) | (480 << 32)
	PROC_InSendMessageEx            PROC_USER = 29 | (481 << 16) | (496 << 32)
	PROC_IsGUIThread                PROC_USER = 30 | (497 << 16) | (508 << 32)
	PROC_LockSetForegroundWindow    PROC_USER = 31 | (509 << 16) | (532 << 32)
	PROC_OffsetRect                 PROC_USER = 32 | (533 << 16) | (543 << 32)
	PROC_PeekMessageW               PROC_USER = 33 | (544 << 16) | (556 << 32)
	PROC_PostQuitMessage            PROC_USER = 34 | (557 << 16) | (572 << 32)
	PROC_PostThreadMessageW         PROC_USER = 35 | (573 << 16) | (591 << 32)
	PROC_RegisterClassExW           PROC_USER = 36 | (592 << 16) | (608 << 32)
	PROC_RegisterClipboardFormatW   PROC_USER = 37 | (609 << 16) | (633 << 32)
	PROC_RegisterWindowMessageW     PROC_USER = 38 | (634 << 16) | (656 << 32)
	PROC_ReplyMessage               PROC_USER = 39 | (657 << 16) | (669 << 32)
	PROC_SetCaretPos                PROC_USER = 40 | (670 << 16) | (681 << 32)
	PROC_SetCursorPos               PROC_USER = 41 | (682 << 16) | (694 << 32)
	PROC_SetMessageExtraInfo        PROC_USER = 42 | (695 << 16) | (714 << 32)
	PROC_SetProcessDefaultLayout    PROC_USER = 43 | (715 << 16) | (738 << 32)
	PROC_SetProcessDPIAware         PROC_USER = 44 | (739 << 16) | (757 << 32)
	PROC_ShowCursor                 PROC_USER = 45 | (758 << 16) | (768 << 32)
	PROC_SoundSentry                PROC_USER = 46 | (769 << 16) | (780 << 32)
	PROC_SystemParametersInfoW      PROC_USER = 47 | (781 << 16) | (802 << 32)
	PROC_TranslateMessage           PROC_USER = 48 | (803 << 16) | (819 << 32)
	PROC_UnregisterClassW           PROC_USER = 49 | (820 << 16) | (836 << 32)
	PROC_WaitMessage                PROC_USER = 50 | (837 << 16) | (848 << 32)
	PROC_CreateAcceleratorTableW    PROC_USER = 51 | (859 << 16) | (882 << 32)
	PROC_CopyAcceleratorTableW      PROC_USER = 52 | (883 << 16) | (904 << 32)
	PROC_DestroyAcceleratorTable    PROC_USER = 53 | (905 << 16) | (928 << 32)
	PROC_GetSysColorBrush           PROC_USER = 54 | (939 << 16) | (955 << 32)
	PROC_OpenClipboard              PROC_USER = 55 | (970 << 16) | (983 << 32)
	PROC_CloseClipboard             PROC_USER = 56 | (984 << 16) | (998 << 32)
	PROC_CountClipboardFormats      PROC_USER = 57 | (999 << 16) | (1020 << 32)
	PROC_EmptyClipboard             PROC_USER = 58 | (1021 << 16) | (1035 << 32)
	PROC_EnumClipboardFormats       PROC_USER = 59 | (1036 << 16) | (1056 << 32)
	PROC_GetClipboardData           PROC_USER = 60 | (1057 << 16) | (1073 << 32)
	PROC_GetClipboardFormatNameW    PROC_USER = 61 | (1074 << 16) | (1097 << 32)
	PROC_GetClipboardSequenceNumber PROC_USER = 62 | (1098 << 16) | (1124 << 32)
	PROC_IsClipboardFormatAvailable PROC_USER = 63 | (1125 << 16) | (1151 << 32)
	PROC_SetClipboardData           PROC_USER = 64 | (1152 << 16) | (1168 << 32)
	PROC_DestroyCursor              PROC_USER = 65 | (1180 << 16) | (1193 << 32)
	PROC_SetCursor                  PROC_USER = 66 | (1194 << 16) | (1203 << 32)
	PROC_DrawIcon                   PROC_USER = 67 | (1211 << 16) | (1219 << 32)
	PROC_DrawIconEx                 PROC_USER = 68 | (1220 << 16) | (1230 << 32)
	PROC_EnumDisplayMonitors        PROC_USER = 69 | (1231 << 16) | (1250 << 32)
	PROC_FrameRect                  PROC_USER = 70 | (1251 << 16) | (1260 << 32)
	PROC_InvertRect                 PROC_USER = 71 | (1261 << 16) | (1271 << 32)
	PROC_PaintDesktop               PROC_USER = 72 | (1272 << 16) | (1284 << 32)
	PROC_WindowFromDC               PROC_USER = 73 | (1285 << 16) | (1297 << 32)
	PROC_BeginDeferWindowPos        PROC_USER = 74 | (1306 << 16) | (1325 << 32)
	PROC_DeferWindowPos             PROC_USER = 75 | (1326 << 16) | (1340 << 32)
	PROC_EndDeferWindowPos          PROC_USER = 76 | (1341 << 16) | (1358 << 32)
	PROC_CreateIconIndirect         PROC_USER = 77 | (1368 << 16) | (1386 << 32)
	PROC_CopyIcon                   PROC_USER = 78 | (1387 << 16) | (1395 << 32)
	PROC_DestroyIcon                PROC_USER = 79 | (1396 << 16) | (1407 << 32)
	PROC_GetIconInfo                PROC_USER = 80 | (1408 << 16) | (1419 << 32)
	PROC_GetIconInfoExW             PROC_USER = 81 | (1420 << 16) | (1434 << 32)
	PROC_CreateDialogParamW         PROC_USER = 82 | (1448 << 16) | (1466 << 32)
	PROC_DialogBoxIndirectParamW    PROC_USER = 83 | (1467 << 16) | (1490 << 32)
	PROC_DialogBoxParamW            PROC_USER = 84 | (1491 << 16) | (1506 << 32)
	PROC_GetClassInfoExW            PROC_USER = 85 | (1507 << 16) | (1522 << 32)
	PROC_LoadAcceleratorsW          PROC_USER = 86 | (1523 << 16) | (1540 << 32)
	PROC_LoadCursorW                PROC_USER = 87 | (1541 << 16) | (1552 << 32)
	PROC_LoadIconW                  PROC_USER = 88 | (1553 << 16) | (1562 << 32)
	PROC_LoadImageW                 PROC_USER = 89 | (1563 << 16) | (1573 << 32)
	PROC_LoadMenuW                  PROC_USER = 90 | (1574 << 16) | (1583 << 32)
	PROC_CreateMenu                 PROC_USER = 91 | (1593 << 16) | (1603 << 32)
	PROC_CreatePopupMenu            PROC_USER = 92 | (1604 << 16) | (1619 << 32)
	PROC_CheckMenuItem              PROC_USER = 93 | (1620 << 16) | (1633 << 32)
	PROC_DeleteMenu                 PROC_USER = 94 | (1634 << 16) | (1644 << 32)
	PROC_DestroyMenu                PROC_USER = 95 | (1645 << 16) | (1656 << 32)
	PROC_EnableMenuItem             PROC_USER = 96 | (1657 << 16) | (1671 << 32)
	PROC_GetMenuDefaultItem         PROC_USER = 97 | (1672 << 16) | (1690 << 32)
	PROC_GetMenuItemID              PROC_USER = 98 | (1691 << 16) | (1704 << 32)
	PROC_GetMenuItemCount           PROC_USER = 99 | (1705 << 16) | (1721 << 32)
	PROC_GetMenuItemInfoW           PROC_USER = 100 | (1722 << 16) | (1738 << 32)
	PROC_GetSubMenu                 PROC_USER = 101 | (1739 << 16) | (1749 << 32)
	PROC_InsertMenuItemW            PROC_USER = 102 | (1750 << 16) | (1765 << 32)
	PROC_RemoveMenu                 PROC_USER = 103 | (1766 << 16) | (1776 << 32)
	PROC_SetMenuDefaultItem         PROC_USER = 104 | (1777 << 16) | (1795 << 32)
	PROC_SetMenuInfo                PROC_USER = 105 | (1796 << 16) | (1807 << 32)
	PROC_SetMenuItemBitmaps         PROC_USER = 106 | (1808 << 16) | (1826 << 32)
	PROC_SetMenuItemInfo            PROC_USER = 107 | (1827 << 16) | (1842 << 32)
	PROC_TrackPopupMenu             PROC_USER = 108 | (1843 << 16) | (1857 << 32)
	PROC_MonitorFromPoint           PROC_USER = 109 | (1870 << 16) | (1886 << 32)
	PROC_MonitorFromRect            PROC_USER = 110 | (1887 << 16) | (1902 << 32)
	PROC_SetUserObjectInformationW  PROC_USER = 111 | (1915 << 16) | (1940 << 32)
	PROC_GetClassLongW              PROC_USER = 112 | (1952 << 16) | (1965 << 32)
	PROC_GetWindowLongW             PROC_USER = 113 | (1966 << 16) | (1980 << 32)
	PROC_SetWindowLongW             PROC_USER = 114 | (1981 << 16) | (1995 << 32)
	PROC_GetClassLongPtrW           PROC_USER = 115 | (2006 << 16) | (2022 << 32)
	PROC_GetWindowLongPtrW          PROC_USER = 116 | (2023 << 16) | (2040 << 32)
	PROC_SetWindowLongPtrW          PROC_USER = 117 | (2041 << 16) | (2058 << 32)
	PROC_CreateWindowExW            PROC_USER = 118 | (2067 << 16) | (2082 << 32)
	PROC_FindWindowW                PROC_USER = 119 | (2083 << 16) | (2094 << 32)
	PROC_GetClipboardOwner          PROC_USER = 120 | (2095 << 16) | (2112 << 32)
	PROC_GetDesktopWindow           PROC_USER = 121 | (2113 << 16) | (2129 << 32)
	PROC_GetFocus                   PROC_USER = 122 | (2130 << 16) | (2138 << 32)
	PROC_GetForegroundWindow        PROC_USER = 123 | (2139 << 16) | (2158 << 32)
	PROC_GetOpenClipboardWindow     PROC_USER = 124 | (2159 << 16) | (2181 << 32)
	PROC_GetShellWindow             PROC_USER = 125 | (2182 << 16) | (2196 << 32)
	PROC_AnimateWindow              PROC_USER = 126 | (2197 << 16) | (2210 << 32)
	PROC_BeginPaint                 PROC_USER = 127 | (2211 << 16) | (2221 << 32)
	PROC_BringWindowToTop           PROC_USER = 128 | (2222 << 16) | (2238 << 32)
	PROC_ChildWindowFromPoint       PROC_USER = 129 | (2239 << 16) | (2259 << 32)
	PROC_ChildWindowFromPointEx     PROC_USER = 130 | (2260 << 16) | (2282 << 32)
	PROC_ClientToScreen             PROC_USER = 131 | (2283 << 16) | (2297 << 32)
	PROC_CloseWindow                PROC_USER = 132 | (2298 << 16) | (2309 << 32)
	PROC_DefDlgProcW                PROC_USER = 133 | (2310 << 16) | (2321 << 32)
	PROC_DefWindowProcW             PROC_USER = 134 | (2322 << 16) | (2336 << 32)
	PROC_DestroyWindow              PROC_USER = 135 | (2337 << 16) | (2350 << 32)
	PROC_DrawMenuBar                PROC_USER = 136 | (2351 << 16) | (2362 << 32)
	PROC_EnableWindow               PROC_USER = 137 | (2363 << 16) | (2375 << 32)
	PROC_EndDialog                  PROC_USER = 138 | (2376 << 16) | (2385 << 32)
	PROC_EndPaint                   PROC_USER = 139 | (2386 << 16) | (2394 << 32)
	PROC_EnumChildWindows           PROC_USER = 140 | (2395 << 16) | (2411 << 32)
	PROC_GetAncestor                PROC_USER = 141 | (2412 << 16) | (2423 << 32)
	PROC_GetClassNameW              PROC_USER = 142 | (2424 << 16) | (2437 << 32)
	PROC_GetClientRect              PROC_USER = 143 | (2438 << 16) | (2451 << 32)
	PROC_GetDC                      PROC_USER = 144 | (2452 << 16) | (2457 << 32)
	PROC_GetDCEx                    PROC_USER = 145 | (2458 << 16) | (2465 << 32)
	PROC_GetDlgCtrlID               PROC_USER = 146 | (2466 << 16) | (2478 << 32)
	PROC_GetDlgItem                 PROC_USER = 147 | (2479 << 16) | (2489 << 32)
	PROC_GetLastActivePopup         PROC_USER = 148 | (2490 << 16) | (2508 << 32)
	PROC_GetMenu                    PROC_USER = 149 | (2509 << 16) | (2516 << 32)
	PROC_GetNextDlgGroupItem        PROC_USER = 150 | (2517 << 16) | (2536 << 32)
	PROC_GetNextDlgTabItem          PROC_USER = 151 | (2537 << 16) | (2554 << 32)
	PROC_GetParent                  PROC_USER = 152 | (2555 << 16) | (2564 << 32)
	PROC_GetWindow                  PROC_USER = 153 | (2565 << 16) | (2574 << 32)
	PROC_GetWindowDC                PROC_USER = 154 | (2575 << 16) | (2586 << 32)
	PROC_GetWindowRect              PROC_USER = 155 | (2587 << 16) | (2600 << 32)
	PROC_GetWindowTextW             PROC_USER = 156 | (2601 << 16) | (2615 << 32)
	PROC_GetWindowTextLengthW       PROC_USER = 157 | (2616 << 16) | (2636 << 32)
	PROC_GetWindowThreadProcessId   PROC_USER = 158 | (2637 << 16) | (2661 << 32)
	PROC_InvalidateRect             PROC_USER = 159 | (2662 << 16) | (2676 << 32)
	PROC_IsChild                    PROC_USER = 160 | (2677 << 16) | (2684 << 32)
	PROC_IsDialogMessageW           PROC_USER = 161 | (2685 << 16) | (2701 << 32)
	PROC_IsWindow                   PROC_USER = 162 | (2702 << 16) | (2710 << 32)
	PROC_MapDialogRect              PROC_USER = 163 | (2711 << 16) | (2724 << 32)
	PROC_MessageBoxW                PROC_USER = 164 | (2725 << 16) | (2736 << 32)
	PROC_MonitorFromWindow          PROC_USER = 165 | (2737 << 16) | (2754 << 32)
	PROC_PostMessageW               PROC_USER = 166 | (2755 << 16) | (2767 << 32)
	PROC_RedrawWindow               PROC_USER = 167 | (2768 << 16) | (2780 << 32)
	PROC_ReleaseDC                  PROC_USER = 168 | (2781 << 16) | (2790 << 32)
	PROC_ScreenToClient             PROC_USER = 169 | (2791 << 16) | (2805 << 32)
	PROC_SendMessageW               PROC_USER = 170 | (2806 << 16) | (2818 << 32)
	PROC_SetFocus                   PROC_USER = 171 | (2819 << 16) | (2827 << 32)
	PROC_SetForegroundWindow        PROC_USER = 172 | (2828 << 16) | (2847 << 32)
	PROC_SetMenu                    PROC_USER = 173 | (2848 << 16) | (2855 << 32)
	PROC_SetWindowPos               PROC_USER = 174 | (2856 << 16) | (2868 << 32)
	PROC_SetWindowRgn               PROC_USER = 175 | (2869 << 16) | (2881 << 32)
	PROC_SetWindowTextW             PROC_USER = 176 | (2882 << 16) | (2896 << 32)
	PROC_ShowCaret                  PROC_USER = 177 | (2897 << 16) | (2906 << 32)
	PROC_ShowWindow                 PROC_USER = 178 | (2907 << 16) | (2917 << 32)
	PROC_TranslateAcceleratorW      PROC_USER = 179 | (2918 << 16) | (2939 << 32)
	PROC_UpdateWindow               PROC_USER = 180 | (2940 << 16) | (2952 << 32)
)

// Declaration of user procedure names.
const userProcStr = `
--funcs
AdjustWindowRectEx
AllowSetForegroundWindow
AnyPopup
BroadcastSystemMessageW
CreateIconFromResourceEx
DestroyCaret
DispatchMessageW
EndMenu
EnumDisplayDevicesW
EnumThreadWindows
EnumWindows
GetAsyncKeyState
GetCaretPos
GetCursorInfo
GetCursorPos
GetDialogBaseUnits
GetGUIThreadInfo
GetInputState
GetMessageW
GetMessageExtraInfo
GetMessagePos
GetMessageTime
GetPhysicalCursorPos
GetProcessDefaultLayout
GetQueueStatus
GetSysColor
GetSystemMetrics
InflateRect
InSendMessage
InSendMessageEx
IsGUIThread
LockSetForegroundWindow
OffsetRect
PeekMessageW
PostQuitMessage
PostThreadMessageW
RegisterClassExW
RegisterClipboardFormatW
RegisterWindowMessageW
ReplyMessage
SetCaretPos
SetCursorPos
SetMessageExtraInfo
SetProcessDefaultLayout
SetProcessDPIAware
ShowCursor
SoundSentry
SystemParametersInfoW
TranslateMessage
UnregisterClassW
WaitMessage

--haccel
CreateAcceleratorTableW
CopyAcceleratorTableW
DestroyAcceleratorTable

--hbrush
GetSysColorBrush

--hclipboard
OpenClipboard
CloseClipboard
CountClipboardFormats
EmptyClipboard
EnumClipboardFormats
GetClipboardData
GetClipboardFormatNameW
GetClipboardSequenceNumber
IsClipboardFormatAvailable
SetClipboardData

--hcursor
DestroyCursor
SetCursor

--hdc
DrawIcon
DrawIconEx
EnumDisplayMonitors
FrameRect
InvertRect
PaintDesktop
WindowFromDC

--hdwp
BeginDeferWindowPos
DeferWindowPos
EndDeferWindowPos

--hicon
CreateIconIndirect
CopyIcon
DestroyIcon
GetIconInfo
GetIconInfoExW

--hinstance
CreateDialogParamW
DialogBoxIndirectParamW
DialogBoxParamW
GetClassInfoExW
LoadAcceleratorsW
LoadCursorW
LoadIconW
LoadImageW
LoadMenuW

--hmenu
CreateMenu
CreatePopupMenu
CheckMenuItem
DeleteMenu
DestroyMenu
EnableMenuItem
GetMenuDefaultItem
GetMenuItemID
GetMenuItemCount
GetMenuItemInfoW
GetSubMenu
InsertMenuItemW
RemoveMenu
SetMenuDefaultItem
SetMenuInfo
SetMenuItemBitmaps
SetMenuItemInfo
TrackPopupMenu

--hmonitor
MonitorFromPoint
MonitorFromRect

--hprocess
SetUserObjectInformationW

--hwnd386
GetClassLongW
GetWindowLongW
SetWindowLongW

--hwnd64
GetClassLongPtrW
GetWindowLongPtrW
SetWindowLongPtrW

--hwnd
CreateWindowExW
FindWindowW
GetClipboardOwner
GetDesktopWindow
GetFocus
GetForegroundWindow
GetOpenClipboardWindow
GetShellWindow
AnimateWindow
BeginPaint
BringWindowToTop
ChildWindowFromPoint
ChildWindowFromPointEx
ClientToScreen
CloseWindow
DefDlgProcW
DefWindowProcW
DestroyWindow
DrawMenuBar
EnableWindow
EndDialog
EndPaint
EnumChildWindows
GetAncestor
GetClassNameW
GetClientRect
GetDC
GetDCEx
GetDlgCtrlID
GetDlgItem
GetLastActivePopup
GetMenu
GetNextDlgGroupItem
GetNextDlgTabItem
GetParent
GetWindow
GetWindowDC
GetWindowRect
GetWindowTextW
GetWindowTextLengthW
GetWindowThreadProcessId
InvalidateRect
IsChild
IsDialogMessageW
IsWindow
MapDialogRect
MessageBoxW
MonitorFromWindow
PostMessageW
RedrawWindow
ReleaseDC
ScreenToClient
SendMessageW
SetFocus
SetForegroundWindow
SetMenu
SetWindowPos
SetWindowRgn
SetWindowTextW
ShowCaret
ShowWindow
TranslateAcceleratorW
UpdateWindow
`
