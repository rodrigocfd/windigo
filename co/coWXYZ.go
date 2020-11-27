/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package co

// WM_ACTIVATE activation state.
type WA int32

const (
	WA_INACTIVE    WA = 0
	WA_ACTIVE      WA = 1
	WA_CLICKACTIVE WA = 2
)

// SetWindowsHookEx() idHook.
type WH int

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

// Window messages.
type WM uint32

const (
	WM_NULL                           WM = 0x0000
	WM_CREATE                         WM = 0x0001
	WM_DESTROY                        WM = 0x0002
	WM_MOVE                           WM = 0x0003
	WM_SIZE                           WM = 0x0005
	WM_ACTIVATE                       WM = 0x0006
	WM_SETFOCUS                       WM = 0x0007
	WM_KILLFOCUS                      WM = 0x0008
	WM_ENABLE                         WM = 0x000a
	WM_SETREDRAW                      WM = 0x000b
	WM_SETTEXT                        WM = 0x000c
	WM_GETTEXT                        WM = 0x000d
	WM_GETTEXTLENGTH                  WM = 0x000e
	WM_PAINT                          WM = 0x000f
	WM_CLOSE                          WM = 0x0010
	WM_QUERYENDSESSION                WM = 0x0011
	WM_QUERYOPEN                      WM = 0x0013
	WM_ENDSESSION                     WM = 0x0016
	WM_QUIT                           WM = 0x0012
	WM_ERASEBKGND                     WM = 0x0014
	WM_SYSCOLORCHANGE                 WM = 0x0015
	WM_SHOWWINDOW                     WM = 0x0018
	WM_WININICHANGE                   WM = 0x001a
	WM_DEVMODECHANGE                  WM = 0x001b
	WM_ACTIVATEAPP                    WM = 0x001c
	WM_FONTCHANGE                     WM = 0x001d
	WM_TIMECHANGE                     WM = 0x001e
	WM_CANCELMODE                     WM = 0x001f
	WM_SETCURSOR                      WM = 0x0020
	WM_MOUSEACTIVATE                  WM = 0x0021
	WM_CHILDACTIVATE                  WM = 0x0022
	WM_QUEUESYNC                      WM = 0x0023
	WM_GETMINMAXINFO                  WM = 0x0024
	WM_PAINTICON                      WM = 0x0026
	WM_ICONERASEBKGND                 WM = 0x0027
	WM_NEXTDLGCTL                     WM = 0x0028
	WM_SPOOLERSTATUS                  WM = 0x002a
	WM_DRAWITEM                       WM = 0x002b
	WM_MEASUREITEM                    WM = 0x002c
	WM_DELETEITEM                     WM = 0x002d
	WM_VKEYTOITEM                     WM = 0x002e
	WM_CHARTOITEM                     WM = 0x002f
	WM_SETFONT                        WM = 0x0030
	WM_GETFONT                        WM = 0x0031
	WM_SETHOTKEY                      WM = 0x0032
	WM_GETHOTKEY                      WM = 0x0033
	WM_QUERYDRAGICON                  WM = 0x0037
	WM_COMPAREITEM                    WM = 0x0039
	WM_GETOBJECT                      WM = 0x003d
	WM_COPYDATA                       WM = 0x004a
	WM_COMPACTING                     WM = 0x0041
	WM_COMMNOTIFY                     WM = 0x0044
	WM_WINDOWPOSCHANGING              WM = 0x0046
	WM_WINDOWPOSCHANGED               WM = 0x0047
	WM_POWER                          WM = 0x0048
	WM_NOTIFY                         WM = 0x004e
	WM_INPUTLANGCHANGEREQUEST         WM = 0x0050
	WM_INPUTLANGCHANGE                WM = 0x0051
	WM_TCARD                          WM = 0x0052
	WM_HELP                           WM = 0x0053
	WM_USERCHANGED                    WM = 0x0054
	WM_NOTIFYFORMAT                   WM = 0x0055
	WM_CONTEXTMENU                    WM = 0x007b
	WM_STYLECHANGING                  WM = 0x007c
	WM_STYLECHANGED                   WM = 0x007d
	WM_DISPLAYCHANGE                  WM = 0x007e
	WM_GETICON                        WM = 0x007f
	WM_SETICON                        WM = 0x0080
	WM_NCCREATE                       WM = 0x0081
	WM_NCDESTROY                      WM = 0x0082
	WM_NCCALCSIZE                     WM = 0x0083
	WM_NCHITTEST                      WM = 0x0084
	WM_NCPAINT                        WM = 0x0085
	WM_NCACTIVATE                     WM = 0x0086
	WM_GETDLGCODE                     WM = 0x0087
	WM_SYNCPAINT                      WM = 0x0088
	WM_NCMOUSEMOVE                    WM = 0x00a0
	WM_NCLBUTTONDOWN                  WM = 0x00a1
	WM_NCLBUTTONUP                    WM = 0x00a2
	WM_NCLBUTTONDBLCLK                WM = 0x00a3
	WM_NCRBUTTONDOWN                  WM = 0x00a4
	WM_NCRBUTTONUP                    WM = 0x00a5
	WM_NCRBUTTONDBLCLK                WM = 0x00a6
	WM_NCMBUTTONDOWN                  WM = 0x00a7
	WM_NCMBUTTONUP                    WM = 0x00a8
	WM_NCMBUTTONDBLCLK                WM = 0x00a9
	WM_NCXBUTTONDOWN                  WM = 0x00ab
	WM_NCXBUTTONUP                    WM = 0x00ac
	WM_NCXBUTTONDBLCLK                WM = 0x00ad
	WM_INPUT_DEVICE_CHANGE            WM = 0x00fe
	WM_INPUT                          WM = 0x00ff
	WM_KEYFIRST                       WM = 0x0100
	WM_KEYDOWN                        WM = 0x0100
	WM_KEYUP                          WM = 0x0101
	WM_CHAR                           WM = 0x0102
	WM_DEADCHAR                       WM = 0x0103
	WM_SYSKEYDOWN                     WM = 0x0104
	WM_SYSKEYUP                       WM = 0x0105
	WM_SYSCHAR                        WM = 0x0106
	WM_SYSDEADCHAR                    WM = 0x0107
	WM_UNICHAR                        WM = 0x0109
	WM_KEYLAST                        WM = 0x0109
	WM_IME_STARTCOMPOSITION           WM = 0x010d
	WM_IME_ENDCOMPOSITION             WM = 0x010e
	WM_IME_COMPOSITION                WM = 0x010f
	WM_IME_KEYLAST                    WM = 0x010f
	WM_INITDIALOG                     WM = 0x0110
	WM_COMMAND                        WM = 0x0111
	WM_SYSCOMMAND                     WM = 0x0112
	WM_TIMER                          WM = 0x0113
	WM_HSCROLL                        WM = 0x0114
	WM_VSCROLL                        WM = 0x0115
	WM_INITMENU                       WM = 0x0116
	WM_INITMENUPOPUP                  WM = 0x0117
	WM_GESTURE                        WM = 0x0119
	WM_GESTURENOTIFY                  WM = 0x011a
	WM_MENUSELECT                     WM = 0x011f
	WM_MENUCHAR                       WM = 0x0120
	WM_ENTERIDLE                      WM = 0x0121
	WM_MENURBUTTONUP                  WM = 0x0122
	WM_MENUDRAG                       WM = 0x0123
	WM_MENUGETOBJECT                  WM = 0x0124
	WM_UNINITMENUPOPUP                WM = 0x0125
	WM_MENUCOMMAND                    WM = 0x0126
	WM_CHANGEUISTATE                  WM = 0x0127
	WM_UPDATEUISTATE                  WM = 0x0128
	WM_QUERYUISTATE                   WM = 0x0129
	WM_CTLCOLORMSGBOX                 WM = 0x0132
	WM_CTLCOLOREDIT                   WM = 0x0133
	WM_CTLCOLORLISTBOX                WM = 0x0134
	WM_CTLCOLORBTN                    WM = 0x0135
	WM_CTLCOLORDLG                    WM = 0x0136
	WM_CTLCOLORSCROLLBAR              WM = 0x0137
	WM_CTLCOLORSTATIC                 WM = 0x0138
	MN_GETHMENU                       WM = 0x01e1
	WM_MOUSEFIRST                     WM = 0x0200
	WM_MOUSEMOVE                      WM = 0x0200
	WM_LBUTTONDOWN                    WM = 0x0201
	WM_LBUTTONUP                      WM = 0x0202
	WM_LBUTTONDBLCLK                  WM = 0x0203
	WM_RBUTTONDOWN                    WM = 0x0204
	WM_RBUTTONUP                      WM = 0x0205
	WM_RBUTTONDBLCLK                  WM = 0x0206
	WM_MBUTTONDOWN                    WM = 0x0207
	WM_MBUTTONUP                      WM = 0x0208
	WM_MBUTTONDBLCLK                  WM = 0x0209
	WM_MOUSEHWHEEL                    WM = 0x020e
	WM_XBUTTONDOWN                    WM = 0x020b
	WM_XBUTTONUP                      WM = 0x020c
	WM_XBUTTONDBLCLK                  WM = 0x020d
	WM_MOUSELAST                      WM = 0x020e
	WM_PARENTNOTIFY                   WM = 0x0210
	WM_ENTERMENULOOP                  WM = 0x0211
	WM_EXITMENULOOP                   WM = 0x0212
	WM_NEXTMENU                       WM = 0x0213
	WM_SIZING                         WM = 0x0214
	WM_CAPTURECHANGED                 WM = 0x0215
	WM_MOVING                         WM = 0x0216
	WM_POWERBROADCAST                 WM = 0x0218
	WM_DEVICECHANGE                   WM = 0x0219
	WM_MDICREATE                      WM = 0x0220
	WM_MDIDESTROY                     WM = 0x0221
	WM_MDIACTIVATE                    WM = 0x0222
	WM_MDIRESTORE                     WM = 0x0223
	WM_MDINEXT                        WM = 0x0224
	WM_MDIMAXIMIZE                    WM = 0x0225
	WM_MDITILE                        WM = 0x0226
	WM_MDICASCADE                     WM = 0x0227
	WM_MDIICONARRANGE                 WM = 0x0228
	WM_MDIGETACTIVE                   WM = 0x0229
	WM_MDISETMENU                     WM = 0x0230
	WM_ENTERSIZEMOVE                  WM = 0x0231
	WM_EXITSIZEMOVE                   WM = 0x0232
	WM_DROPFILES                      WM = 0x0233
	WM_MDIREFRESHMENU                 WM = 0x0234
	WM_POINTERDEVICECHANGE            WM = 0x0238
	WM_POINTERDEVICEINRANGE           WM = 0x0239
	WM_POINTERDEVICEOUTOFRANGE        WM = 0x023a
	WM_TOUCH                          WM = 0x0240
	WM_NCPOINTERUPDATE                WM = 0x0241
	WM_NCPOINTERDOWN                  WM = 0x0242
	WM_NCPOINTERUP                    WM = 0x0243
	WM_POINTERUPDATE                  WM = 0x0245
	WM_POINTERDOWN                    WM = 0x0246
	WM_POINTERUP                      WM = 0x0247
	WM_POINTERENTER                   WM = 0x0249
	WM_POINTERLEAVE                   WM = 0x024a
	WM_POINTERACTIVATE                WM = 0x024b
	WM_POINTERCAPTURECHANGED          WM = 0x024c
	WM_TOUCHHITTESTING                WM = 0x024d
	WM_POINTERWHEEL                   WM = 0x024e
	WM_POINTERHWHEEL                  WM = 0x024f
	WM_POINTERHITTEST                 WM = 0x0250 // Originally DM_POINTERHITTEST.
	WM_POINTERROUTEDTO                WM = 0x0251
	WM_POINTERROUTEDAWAY              WM = 0x0252
	WM_POINTERROUTEDRELEASED          WM = 0x0253
	WM_IME_SETCONTEXT                 WM = 0x0281
	WM_IME_NOTIFY                     WM = 0x0282
	WM_IME_CONTROL                    WM = 0x0283
	WM_IME_COMPOSITIONFULL            WM = 0x0284
	WM_IME_SELECT                     WM = 0x0285
	WM_IME_CHAR                       WM = 0x0286
	WM_IME_REQUEST                    WM = 0x0288
	WM_IME_KEYDOWN                    WM = 0x0290
	WM_IME_KEYUP                      WM = 0x0291
	WM_MOUSEHOVER                     WM = 0x02a1
	WM_MOUSELEAVE                     WM = 0x02a3
	WM_NCMOUSEHOVER                   WM = 0x02a0
	WM_NCMOUSELEAVE                   WM = 0x02a2
	WM_WTSSESSION_CHANGE              WM = 0x02b1
	WM_TABLET_FIRST                   WM = 0x02c0
	WM_TABLET_LAST                    WM = 0x02df
	WM_DPICHANGED                     WM = 0x02e0
	WM_DPICHANGED_BEFOREPARENT        WM = 0x02e2
	WM_DPICHANGED_AFTERPARENT         WM = 0x02e3
	WM_GETDPISCALEDSIZE               WM = 0x02e4
	WM_CUT                            WM = 0x0300
	WM_COPY                           WM = 0x0301
	WM_PASTE                          WM = 0x0302
	WM_CLEAR                          WM = 0x0303
	WM_UNDO                           WM = 0x0304
	WM_RENDERFORMAT                   WM = 0x0305
	WM_RENDERALLFORMATS               WM = 0x0306
	WM_DESTROYCLIPBOARD               WM = 0x0307
	WM_DRAWCLIPBOARD                  WM = 0x0308
	WM_PAINTCLIPBOARD                 WM = 0x0309
	WM_VSCROLLCLIPBOARD               WM = 0x030a
	WM_SIZECLIPBOARD                  WM = 0x030b
	WM_ASKCBFORMATNAME                WM = 0x030c
	WM_CHANGECBCHAIN                  WM = 0x030d
	WM_HSCROLLCLIPBOARD               WM = 0x030e
	WM_QUERYNEWPALETTE                WM = 0x030f
	WM_PALETTEISCHANGING              WM = 0x0310
	WM_PALETTECHANGED                 WM = 0x0311
	WM_HOTKEY                         WM = 0x0312
	WM_PRINT                          WM = 0x0317
	WM_PRINTCLIENT                    WM = 0x0318
	WM_APPCOMMAND                     WM = 0x0319
	WM_THEMECHANGED                   WM = 0x031a
	WM_CLIPBOARDUPDATE                WM = 0x031d
	WM_DWMCOMPOSITIONCHANGED          WM = 0x031e
	WM_DWMNCRENDERINGCHANGED          WM = 0x031f
	WM_DWMCOLORIZATIONCOLORCHANGED    WM = 0x0320
	WM_DWMWINDOWMAXIMIZEDCHANGE       WM = 0x0321
	WM_DWMSENDICONICTHUMBNAIL         WM = 0x0323
	WM_DWMSENDICONICLIVEPREVIEWBITMAP WM = 0x0326
	WM_GETTITLEBARINFOEX              WM = 0x033f
	WM_HANDHELDFIRST                  WM = 0x0358
	WM_HANDHELDLAST                   WM = 0x035f
	WM_AFXFIRST                       WM = 0x0360
	WM_AFXLAST                        WM = 0x037f
	WM_PENWINFIRST                    WM = 0x0380
	WM_PENWINLAST                     WM = 0x038f
	WM_APP                            WM = 0x8000
	WM_USER                           WM = 0x0400
)

// Window styles.
type WS uint32

const (
	WS_NONE             WS = 0
	WS_OVERLAPPED       WS = 0x00000000
	WS_POPUP            WS = 0x80000000
	WS_CHILD            WS = 0x40000000 // The window is a child window.
	WS_MINIMIZE         WS = 0x20000000 // The window is initially minimized.
	WS_VISIBLE          WS = 0x10000000 // The window is initially visible.
	WS_DISABLED         WS = 0x08000000 // The window is initially disabled.
	WS_CLIPSIBLINGS     WS = 0x04000000 // Clips child windows relative to each other.
	WS_CLIPCHILDREN     WS = 0x02000000 // Excludes the area occupied by child windows when drawing occurs within the parent window. This style is used when creating the parent window.
	WS_MAXIMIZE         WS = 0x01000000 // The window is initially maximized.
	WS_CAPTION          WS = 0x00c00000 // The window has a title bar (includes the WS_BORDER style).
	WS_BORDER           WS = 0x00800000
	WS_DLGFRAME         WS = 0x00400000
	WS_VSCROLL          WS = 0x00200000 // The window has a vertical scroll bar.
	WS_HSCROLL          WS = 0x00100000 // The window has a horizontal scroll bar.
	WS_SYSMENU          WS = 0x00080000 // The window has a window menu on its title bar. The WS_CAPTION style must also be specified.
	WS_THICKFRAME       WS = 0x00040000 // The window has a sizing border. Same as the WS_SIZEBOX style.
	WS_GROUP            WS = 0x00020000 // The window is the first control of a group of controls.
	WS_TABSTOP          WS = 0x00010000 // The window is a control that can receive the keyboard focus when the user presses the TAB key.
	WS_MINIMIZEBOX      WS = 0x00020000 // The window has a minimize button.
	WS_MAXIMIZEBOX      WS = 0x00010000 // The window has a maximize button.
	WS_TILED            WS = WS_OVERLAPPED
	WS_ICONIC           WS = WS_MINIMIZE   // The window is initially minimized. Same as the WS_MINIMIZE style.
	WS_SIZEBOX          WS = WS_THICKFRAME // The window has a sizing border. Same as the WS_THICKFRAME style.
	WS_TILEDWINDOW      WS = WS_OVERLAPPEDWINDOW
	WS_OVERLAPPEDWINDOW WS = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_POPUPWINDOW      WS = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW      WS = WS_CHILD // Same as the WS_CHILD style.
)

// Window extended styles.
type WS_EX uint32

const (
	WS_EX_NONE                WS_EX = 0
	WS_EX_DLGMODALFRAME       WS_EX = 0x00000001
	WS_EX_NOPARENTNOTIFY      WS_EX = 0x00000004 // The child window created with this style does not send the WM_PARENTNOTIFY message to its parent window when it is created or destroyed.
	WS_EX_TOPMOST             WS_EX = 0x00000008
	WS_EX_ACCEPTFILES         WS_EX = 0x00000010 // The window accepts drag-drop files.
	WS_EX_TRANSPARENT         WS_EX = 0x00000020
	WS_EX_MDICHILD            WS_EX = 0x00000040 // The window is a MDI child window.
	WS_EX_TOOLWINDOW          WS_EX = 0x00000080
	WS_EX_WINDOWEDGE          WS_EX = 0x00000100
	WS_EX_CLIENTEDGE          WS_EX = 0x00000200 // The window has a border with a sunken edge.
	WS_EX_CONTEXTHELP         WS_EX = 0x00000400
	WS_EX_RIGHT               WS_EX = 0x00001000
	WS_EX_LEFT                WS_EX = 0x00000000
	WS_EX_RTLREADING          WS_EX = 0x00002000
	WS_EX_LTRREADING          WS_EX = 0x00000000
	WS_EX_LEFTSCROLLBAR       WS_EX = 0x00004000
	WS_EX_RIGHTSCROLLBAR      WS_EX = 0x00000000
	WS_EX_CONTROLPARENT       WS_EX = 0x00010000
	WS_EX_STATICEDGE          WS_EX = 0x00020000
	WS_EX_APPWINDOW           WS_EX = 0x00040000 // Forces a top-level window onto the taskbar when the window is visible.
	WS_EX_OVERLAPPEDWINDOW    WS_EX = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	WS_EX_PALETTEWINDOW       WS_EX = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST
	WS_EX_LAYERED             WS_EX = 0x00080000
	WS_EX_NOINHERITLAYOUT     WS_EX = 0x00100000
	WS_EX_NOREDIRECTIONBITMAP WS_EX = 0x00200000
	WS_EX_LAYOUTRTL           WS_EX = 0x00400000
	WS_EX_COMPOSITED          WS_EX = 0x02000000
	WS_EX_NOACTIVATE          WS_EX = 0x08000000
)

// WM_NCCALCSIZE return flags.
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
