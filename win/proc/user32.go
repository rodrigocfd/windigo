/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllUser32 = syscall.NewLazyDLL("user32.dll")

	AdjustWindowRectEx            = dllUser32.NewProc("AdjustWindowRectEx")
	AppendMenu                    = dllUser32.NewProc("AppendMenuW")
	BeginDeferWindowPos           = dllUser32.NewProc("BeginDeferWindowPos")
	BeginPaint                    = dllUser32.NewProc("BeginPaint")
	CallNextHookEx                = dllUser32.NewProc("CallNextHookEx")
	CheckMenuItem                 = dllUser32.NewProc("CheckMenuItem")
	CheckMenuRadioItem            = dllUser32.NewProc("CheckMenuRadioItem")
	ClientToScreen                = dllUser32.NewProc("ClientToScreen")
	CopyAcceleratorTable          = dllUser32.NewProc("CopyAcceleratorTableW")
	CreateAcceleratorTable        = dllUser32.NewProc("CreateAcceleratorTableW")
	CreateMenu                    = dllUser32.NewProc("CreateMenu")
	CreatePopupMenu               = dllUser32.NewProc("CreatePopupMenu")
	CreateWindowEx                = dllUser32.NewProc("CreateWindowExW")
	DeferWindowPos                = dllUser32.NewProc("DeferWindowPos")
	DefWindowProc                 = dllUser32.NewProc("DefWindowProcW")
	DeleteMenu                    = dllUser32.NewProc("DeleteMenu")
	DestroyAcceleratorTable       = dllUser32.NewProc("DestroyAcceleratorTable")
	DestroyCaret                  = dllUser32.NewProc("DestroyCaret")
	DestroyCursor                 = dllUser32.NewProc("DestroyCursor")
	DestroyIcon                   = dllUser32.NewProc("DestroyIcon")
	DestroyMenu                   = dllUser32.NewProc("DestroyMenu")
	DestroyWindow                 = dllUser32.NewProc("DestroyWindow")
	DispatchMessage               = dllUser32.NewProc("DispatchMessageW")
	DrawMenuBar                   = dllUser32.NewProc("DrawMenuBar")
	EmptyClipboard                = dllUser32.NewProc("EmptyClipboard")
	EnableMenuItem                = dllUser32.NewProc("EnableMenuItem")
	EnableWindow                  = dllUser32.NewProc("EnableWindow")
	EndDeferWindowPos             = dllUser32.NewProc("EndDeferWindowPos")
	EndMenu                       = dllUser32.NewProc("EndMenu")
	EndPaint                      = dllUser32.NewProc("EndPaint")
	EnumChildWindows              = dllUser32.NewProc("EnumChildWindows")
	EnumDisplayMonitors           = dllUser32.NewProc("EnumDisplayMonitors")
	EnumWindows                   = dllUser32.NewProc("EnumWindows")
	GetAncestor                   = dllUser32.NewProc("GetAncestor")
	GetAsyncKeyState              = dllUser32.NewProc("GetAsyncKeyState")
	GetCaretPos                   = dllUser32.NewProc("GetCaretPos")
	GetClassInfoEx                = dllUser32.NewProc("GetClassInfoExW")
	GetClientRect                 = dllUser32.NewProc("GetClientRect")
	GetCursorPos                  = dllUser32.NewProc("GetCursorPos")
	GetDC                         = dllUser32.NewProc("GetDC")
	GetDlgCtrlID                  = dllUser32.NewProc("GetDlgCtrlID")
	GetDlgItem                    = dllUser32.NewProc("GetDlgItem")
	GetDpiForSystem               = dllUser32.NewProc("GetDpiForSystem")
	GetDpiForWindow               = dllUser32.NewProc("GetDpiForWindow")
	GetFocus                      = dllUser32.NewProc("GetFocus")
	GetForegroundWindow           = dllUser32.NewProc("GetForegroundWindow")
	GetMenu                       = dllUser32.NewProc("GetMenu")
	GetMenuDefaultItem            = dllUser32.NewProc("GetMenuDefaultItem")
	GetMenuInfo                   = dllUser32.NewProc("GetMenuInfo")
	GetMenuItemCount              = dllUser32.NewProc("GetMenuItemCount")
	GetMenuItemID                 = dllUser32.NewProc("GetMenuItemID")
	GetMenuItemInfo               = dllUser32.NewProc("GetMenuItemInfoW")
	GetMessage                    = dllUser32.NewProc("GetMessageW")
	GetMonitorInfo                = dllUser32.NewProc("GetMonitorInfoW")
	GetNextDlgTabItem             = dllUser32.NewProc("GetNextDlgTabItem")
	GetParent                     = dllUser32.NewProc("GetParent")
	GetPhysicalCursorPos          = dllUser32.NewProc("GetPhysicalCursorPos")
	GetSubMenu                    = dllUser32.NewProc("GetSubMenu")
	GetSystemMenu                 = dllUser32.NewProc("GetSystemMenu")
	GetSystemMetrics              = dllUser32.NewProc("GetSystemMetrics")
	GetWindow                     = dllUser32.NewProc("GetWindow")
	GetWindowDC                   = dllUser32.NewProc("GetWindowDC")
	GetWindowLongPtr              = dllUser32.NewProc("GetWindowLongPtrW")
	GetWindowRect                 = dllUser32.NewProc("GetWindowRect")
	GetWindowText                 = dllUser32.NewProc("GetWindowTextW")
	GetWindowTextLength           = dllUser32.NewProc("GetWindowTextLengthW")
	HideCaret                     = dllUser32.NewProc("HideCaret")
	InsertMenu                    = dllUser32.NewProc("InsertMenuW")
	InsertMenuItem                = dllUser32.NewProc("InsertMenuItemW")
	InvalidateRect                = dllUser32.NewProc("InvalidateRect")
	IsChild                       = dllUser32.NewProc("IsChild")
	IsDialogMessage               = dllUser32.NewProc("IsDialogMessageW")
	IsDlgButtonChecked            = dllUser32.NewProc("IsDlgButtonChecked")
	IsGUIThread                   = dllUser32.NewProc("IsGUIThread")
	IsWindow                      = dllUser32.NewProc("IsWindow")
	IsWindowEnabled               = dllUser32.NewProc("IsWindowEnabled")
	KillTimer                     = dllUser32.NewProc("KillTimer")
	LoadCursor                    = dllUser32.NewProc("LoadCursorW")
	LoadIcon                      = dllUser32.NewProc("LoadIconW")
	LoadMenu                      = dllUser32.NewProc("LoadMenuW")
	MenuItemFromPoint             = dllUser32.NewProc("MeunItemFromPoint")
	MessageBox                    = dllUser32.NewProc("MessageBoxW")
	MonitorFromPoint              = dllUser32.NewProc("MonitorFromPoint")
	MoveWindow                    = dllUser32.NewProc("MoveWindow")
	PostMessage                   = dllUser32.NewProc("PostMessageW")
	PostQuitMessage               = dllUser32.NewProc("PostQuitMessage")
	PostThreadMessage             = dllUser32.NewProc("PostThreadMessageW")
	RegisterClassEx               = dllUser32.NewProc("RegisterClassExW")
	RegisterWindowMessage         = dllUser32.NewProc("RegisterWindowMessageW")
	ReleaseDC                     = dllUser32.NewProc("ReleaseDC")
	ReplyMessage                  = dllUser32.NewProc("ReplyMessage")
	ScreenToClient                = dllUser32.NewProc("ScreenToClient")
	SendMessage                   = dllUser32.NewProc("SendMessageW")
	SetFocus                      = dllUser32.NewProc("SetFocus")
	SetForegroundWindow           = dllUser32.NewProc("SetForegroundWindow")
	SetMenuDefaultItem            = dllUser32.NewProc("SetMenuDefaultItem")
	SetMenuItemInfo               = dllUser32.NewProc("SetMenuItemInfoW")
	SetParent                     = dllUser32.NewProc("SetParent")
	SetProcessDPIAware            = dllUser32.NewProc("SetProcessDPIAware")
	SetProcessDpiAwarenessContext = dllUser32.NewProc("SetProcessDpiAwarenessContext")
	SetRect                       = dllUser32.NewProc("SetRect")
	SetTimer                      = dllUser32.NewProc("SetTimer")
	SetWindowLongPtr              = dllUser32.NewProc("SetWindowLongPtrW")
	SetWindowPos                  = dllUser32.NewProc("SetWindowPos")
	SetWindowsHookEx              = dllUser32.NewProc("SetWindowsHookExW")
	SetWindowText                 = dllUser32.NewProc("SetWindowTextW")
	ShowCaret                     = dllUser32.NewProc("ShowCaret")
	ShowWindow                    = dllUser32.NewProc("ShowWindow")
	SystemParametersInfo          = dllUser32.NewProc("SystemParametersInfoW")
	TrackPopupMenu                = dllUser32.NewProc("TrackPopupMenu")
	TranslateAccelerator          = dllUser32.NewProc("TranslateAcceleratorW")
	TranslateMessage              = dllUser32.NewProc("TranslateMessage")
	UnhookWindowsHookEx           = dllUser32.NewProc("UnhookWindowsHookEx")
	UpdateWindow                  = dllUser32.NewProc("UpdateWindow")
)
