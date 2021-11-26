package co

// WM_ACTIVATE activation state.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-activate
type WA int32

const (
	WA_INACTIVE    WA = 0
	WA_ACTIVE      WA = 1
	WA_CLICKACTIVE WA = 2
)

// WaitForSingleObject() return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
type WAIT uint32

const (
	WAIT_ABANDONED WAIT = 0x0000_0080
	WAIT_OBJECT_0  WAIT = 0x0000_0000
	WAIT_TIMEOUT   WAIT = 0x0000_0102
	WAIT_FAILED    WAIT = 0xffff_ffff
)

// SetWindowDisplayAffinity() dwAffinity
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowdisplayaffinity
type WDA uint32

const (
	WDA_NONE               WDA = 0x0000_0000
	WDA_MONITOR            WDA = 0x0000_0001
	WDA_EXCLUDEFROMCAPTURE WDA = 0x0000_0011
)

// SetWindowsHookEx() idHook.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
type WH int32

const (
	WH_MSGFILTER       WH = -1
	WH_JOURNALRECORD   WH = 0
	WH_JOURNALPLAYBACK WH = 1
	WH_KEYBOARD        WH = 2
	WH_GETMESSAGE      WH = 3
	WH_CALLWNDPROC     WH = 4
	WH_CBT             WH = 5
	WH_SYSMSGFILTER    WH = 6
	WH_MOUSE           WH = 7
	WH_DEBUG           WH = 9
	WH_SHELL           WH = 10
	WH_FOREGROUNDIDLE  WH = 11
	WH_CALLWNDPROCRET  WH = 12
	WH_KEYBOARD_LL     WH = 13
	WH_MOUSE_LL        WH = 14
)

// IsWindowsVersionOrGreater() values; originally _WIN32_WINNT.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/using-the-windows-headers
type WIN32_WINNT uint16

const (
	WIN32_WINNT_NT4          WIN32_WINNT = 0x0400
	WIN32_WINNT_WIN2K        WIN32_WINNT = 0x0500
	WIN32_WINNT_WINXP        WIN32_WINNT = 0x0501
	WIN32_WINNT_WS03         WIN32_WINNT = 0x0502
	WIN32_WINNT_WIN6         WIN32_WINNT = 0x0600
	WIN32_WINNT_VISTA        WIN32_WINNT = 0x0600
	WIN32_WINNT_WS08         WIN32_WINNT = 0x0600
	WIN32_WINNT_LONGHORN     WIN32_WINNT = 0x0600
	WIN32_WINNT_WIN7         WIN32_WINNT = 0x0601
	WIN32_WINNT_WIN8         WIN32_WINNT = 0x0602
	WIN32_WINNT_WINBLUE      WIN32_WINNT = 0x0603
	WIN32_WINNT_WINTHRESHOLD WIN32_WINNT = 0x0a00
	WIN32_WINNT_WIN10        WIN32_WINNT = 0x0a00
)

// WM_SIZING window edge.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-sizing
type WMSZ uint8

const (
	WMSZ_BOTTOM      WMSZ = 6
	WMSZ_BOTTOMLEFT  WMSZ = 7
	WMSZ_BOTTOMRIGHT WMSZ = 8
	WMSZ_LEFT        WMSZ = 1
	WMSZ_RIGHT       WMSZ = 2
	WMSZ_TOP         WMSZ = 3
	WMSZ_TOPLEFT     WMSZ = 4
	WMSZ_TOPRIGHT    WMSZ = 5
)

// Window styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/window-styles
type WS uint32

const (
	WS_NONE             WS = 0
	WS_OVERLAPPED       WS = 0x0000_0000         // The window is an overlapped window. An overlapped window has a title bar and a border. Same as the WS_TILED style.
	WS_POPUP            WS = 0x8000_0000         // The window is a pop-up window. This style cannot be used with the WS_CHILD style.
	WS_CHILD            WS = 0x4000_0000         // The window is a child window.
	WS_MINIMIZE         WS = 0x2000_0000         // The window is initially minimized.
	WS_VISIBLE          WS = 0x1000_0000         // The window is initially visible.
	WS_DISABLED         WS = 0x0800_0000         // The window is initially disabled.
	WS_CLIPSIBLINGS     WS = 0x0400_0000         // Clips child windows relative to each other.
	WS_CLIPCHILDREN     WS = 0x0200_0000         // Excludes the area occupied by child windows when drawing occurs within the parent window. This style is used when creating the parent window.
	WS_MAXIMIZE         WS = 0x0100_0000         // The window is initially maximized.
	WS_CAPTION          WS = 0x00c0_0000         // The window has a title bar (includes the WS_BORDER style).
	WS_BORDER           WS = 0x0080_0000         // The window has a thin-line border.
	WS_DLGFRAME         WS = 0x0040_0000         // The window has a border of a style typically used with dialog boxes. A window with this style cannot have a title bar.
	WS_VSCROLL          WS = 0x0020_0000         // The window has a vertical scroll bar.
	WS_HSCROLL          WS = 0x0010_0000         // The window has a horizontal scroll bar.
	WS_SYSMENU          WS = 0x0008_0000         // The window has a window menu on its title bar. The WS_CAPTION style must also be specified.
	WS_THICKFRAME       WS = 0x0004_0000         // The window has a sizing border. Same as the WS_SIZEBOX style.
	WS_GROUP            WS = 0x0002_0000         // The window is the first control of a group of controls.
	WS_TABSTOP          WS = 0x0001_0000         // The window is a control that can receive the keyboard focus when the user presses the TAB key.
	WS_MINIMIZEBOX      WS = 0x0002_0000         // The window has a minimize button.
	WS_MAXIMIZEBOX      WS = 0x0001_0000         // The window has a maximize button.
	WS_TILED            WS = WS_OVERLAPPED       // The window is an overlapped window. An overlapped window has a title bar and a border. Same as the WS_OVERLAPPED style.
	WS_ICONIC           WS = WS_MINIMIZE         // The window is initially minimized. Same as the WS_MINIMIZE style.
	WS_SIZEBOX          WS = WS_THICKFRAME       // The window has a sizing border. Same as the WS_THICKFRAME style.
	WS_TILEDWINDOW      WS = WS_OVERLAPPEDWINDOW // The window is an overlapped window. Same as the WS_OVERLAPPEDWINDOW style.
	WS_OVERLAPPEDWINDOW WS = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU |
		WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX // The window is an overlapped window. Same as the WS_TILEDWINDOW style.
	WS_POPUPWINDOW WS = WS_POPUP | WS_BORDER | WS_SYSMENU // The window is a pop-up window. The WS_CAPTION and WS_POPUPWINDOW styles must be combined to make the window menu visible.
	WS_CHILDWINDOW WS = WS_CHILD                          // Same as the WS_CHILD style.
)

// Extended window styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/extended-window-styles
type WS_EX uint32

const (
	WS_EX_NONE                WS_EX = 0
	WS_EX_DLGMODALFRAME       WS_EX = 0x0000_0001 // The window has a double border; the window can, optionally, be created with a title bar by specifying the WS_CAPTION style in the dwStyle parameter.
	WS_EX_NOPARENTNOTIFY      WS_EX = 0x0000_0004 // The child window created with this style does not send the WM_PARENTNOTIFY message to its parent window when it is created or destroyed.
	WS_EX_TOPMOST             WS_EX = 0x0000_0008 // The window should be placed above all non-topmost windows and should stay above them, even when the window is deactivated.
	WS_EX_ACCEPTFILES         WS_EX = 0x0000_0010 // The window accepts drag-drop files.
	WS_EX_TRANSPARENT         WS_EX = 0x0000_0020
	WS_EX_MDICHILD            WS_EX = 0x0000_0040 // The window is a MDI child window.
	WS_EX_TOOLWINDOW          WS_EX = 0x0000_0080 // The window is intended to be used as a floating toolbar.
	WS_EX_WINDOWEDGE          WS_EX = 0x0000_0100 // The window has a border with a raised edge.
	WS_EX_CLIENTEDGE          WS_EX = 0x0000_0200 // The window has a border with a sunken edge.
	WS_EX_CONTEXTHELP         WS_EX = 0x0000_0400
	WS_EX_RIGHT               WS_EX = 0x0000_1000
	WS_EX_LEFT                WS_EX = 0x0000_0000 // The window has generic left-aligned properties. This is the default.
	WS_EX_RTLREADING          WS_EX = 0x0000_2000
	WS_EX_LTRREADING          WS_EX = 0x0000_0000 // The window text is displayed using left-to-right reading-order properties. This is the default.
	WS_EX_LEFTSCROLLBAR       WS_EX = 0x0000_4000
	WS_EX_RIGHTSCROLLBAR      WS_EX = 0x0000_0000 // The vertical scroll bar (if present) is to the right of the client area. This is the default.
	WS_EX_CONTROLPARENT       WS_EX = 0x0001_0000
	WS_EX_STATICEDGE          WS_EX = 0x0002_0000 // The window has a three-dimensional border style intended to be used for items that do not accept user input.
	WS_EX_APPWINDOW           WS_EX = 0x0004_0000 // Forces a top-level window onto the taskbar when the window is visible.
	WS_EX_OVERLAPPEDWINDOW    WS_EX = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	WS_EX_PALETTEWINDOW       WS_EX = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST // The window is palette window, which is a modeless dialog box that presents an array of commands.
	WS_EX_LAYERED             WS_EX = 0x0008_0000
	WS_EX_NOINHERITLAYOUT     WS_EX = 0x0010_0000 // The window does not pass its window layout to its child windows.
	WS_EX_NOREDIRECTIONBITMAP WS_EX = 0x0020_0000
	WS_EX_LAYOUTRTL           WS_EX = 0x0040_0000
	WS_EX_COMPOSITED          WS_EX = 0x0200_0000
	WS_EX_NOACTIVATE          WS_EX = 0x0800_0000
)

// WM_NCCALCSIZE return flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
type WVR uint32

const (
	WVR_ZERO        WVR = 0
	WVR_ALIGNTOP    WVR = 0x0010
	WVR_ALIGNLEFT   WVR = 0x0020
	WVR_ALIGNBOTTOM WVR = 0x0040
	WVR_ALIGNRIGHT  WVR = 0x0080
	WVR_HREDRAW     WVR = 0x0100
	WVR_VREDRAW     WVR = 0x0200
	WVR_REDRAW      WVR = WVR_HREDRAW | WVR_VREDRAW
	WVR_VALIDRECTS  WVR = 0x0400
)
