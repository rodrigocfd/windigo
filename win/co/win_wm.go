//go:build windows

package co

// Window [messages].
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/learnwin32/window-messages
type WM uint32

// Standard window [messages] (WM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/learnwin32/window-messages
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
	WM_MN_GETHMENU                    WM = 0x01e1
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

// Button control [messages] (BCM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-messages
const (
	_BCM_FIRST WM = 0x1600

	BCM_GETIDEALSIZE     WM = _BCM_FIRST + 1
	BCM_SETIMAGELIST     WM = _BCM_FIRST + 2
	BCM_GETIMAGELIST     WM = _BCM_FIRST + 3
	BCM_SETTEXTMARGIN    WM = _BCM_FIRST + 4
	BCM_GETTEXTMARGIN    WM = _BCM_FIRST + 5
	BCM_SETDROPDOWNSTATE WM = _BCM_FIRST + 6
	BCM_SETSPLITINFO     WM = _BCM_FIRST + 7
	BCM_GETSPLITINFO     WM = _BCM_FIRST + 8
	BCM_SETNOTE          WM = _BCM_FIRST + 9
	BCM_GETNOTE          WM = _BCM_FIRST + 10
	BCM_GETNOTELENGTH    WM = _BCM_FIRST + 11
	BCM_SETSHIELD        WM = _BCM_FIRST + 12
)

// Button control [messages] (BM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-messages
const (
	BM_GETCHECK     WM = 0x00f0
	BM_SETCHECK     WM = 0x00f1
	BM_GETSTATE     WM = 0x00f2
	BM_SETSTATE     WM = 0x00f3
	BM_SETSTYLE     WM = 0x00f4
	BM_CLICK        WM = 0x00f5
	BM_GETIMAGE     WM = 0x00f6
	BM_SETIMAGE     WM = 0x00f7
	BM_SETDONTCLICK WM = 0x00f8
)

// Common controls [messages] (CCM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-general-control-reference-messages
const (
	_CCM_FIRST WM = 0x2000

	CCM_SETBKCOLOR       WM = _CCM_FIRST + 1
	CCM_SETCOLORSCHEME   WM = _CCM_FIRST + 2
	CCM_GETCOLORSCHEME   WM = _CCM_FIRST + 3
	CCM_GETDROPTARGET    WM = _CCM_FIRST + 4
	CCM_SETUNICODEFORMAT WM = _CCM_FIRST + 5
	CCM_GETUNICODEFORMAT WM = _CCM_FIRST + 6
	CCM_SETVERSION       WM = _CCM_FIRST + 0x7
	CCM_GETVERSION       WM = _CCM_FIRST + 0x8
	CCM_SETNOTIFYWINDOW  WM = _CCM_FIRST + 0x9
	CCM_SETWINDOWTHEME   WM = _CCM_FIRST + 0xb
	CCM_DPISCALE         WM = _CCM_FIRST + 0xc
)

// ComboBox control [messages] (CB).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-combobox-control-reference-messages
const (
	CB_GETEDITSEL            WM = 0x0140
	CB_LIMITTEXT             WM = 0x0141
	CB_SETEDITSEL            WM = 0x0142
	CB_ADDSTRING             WM = 0x0143
	CB_DELETESTRING          WM = 0x0144
	CB_DIR                   WM = 0x0145
	CB_GETCOUNT              WM = 0x0146
	CB_GETCURSEL             WM = 0x0147
	CB_GETLBTEXT             WM = 0x0148
	CB_GETLBTEXTLEN          WM = 0x0149
	CB_INSERTSTRING          WM = 0x014a
	CB_RESETCONTENT          WM = 0x014b
	CB_FINDSTRING            WM = 0x014c
	CB_SELECTSTRING          WM = 0x014d
	CB_SETCURSEL             WM = 0x014e
	CB_SHOWDROPDOWN          WM = 0x014f
	CB_GETITEMDATA           WM = 0x0150
	CB_SETITEMDATA           WM = 0x0151
	CB_GETDROPPEDCONTROLRECT WM = 0x0152
	CB_SETITEMHEIGHT         WM = 0x0153
	CB_GETITEMHEIGHT         WM = 0x0154
	CB_SETEXTENDEDUI         WM = 0x0155
	CB_GETEXTENDEDUI         WM = 0x0156
	CB_GETDROPPEDSTATE       WM = 0x0157
	CB_FINDSTRINGEXACT       WM = 0x0158
	CB_SETLOCALE             WM = 0x0159
	CB_GETLOCALE             WM = 0x015a
	CB_GETTOPINDEX           WM = 0x015b
	CB_SETTOPINDEX           WM = 0x015c
	CB_GETHORIZONTALEXTENT   WM = 0x015d
	CB_SETHORIZONTALEXTENT   WM = 0x015e
	CB_GETDROPPEDWIDTH       WM = 0x015f
	CB_SETDROPPEDWIDTH       WM = 0x0160
	CB_INITSTORAGE           WM = 0x0161
	CB_GETCOMBOBOXINFO       WM = 0x0164
	CB_MSGMAX                WM = 0x0165
)

// DateTimePicker control [messages] (DTM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-date-and-time-picker-control-reference-messages
const (
	_DTM_FIRST WM = 0x1000

	DTM_GETSYSTEMTIME         WM = _DTM_FIRST + 1
	DTM_SETSYSTEMTIME         WM = _DTM_FIRST + 2
	DTM_GETRANGE              WM = _DTM_FIRST + 3
	DTM_SETRANGE              WM = _DTM_FIRST + 4
	DTM_SETFORMAT             WM = _DTM_FIRST + 50
	DTM_SETMCCOLOR            WM = _DTM_FIRST + 6
	DTM_GETMCCOLOR            WM = _DTM_FIRST + 7
	DTM_GETMONTHCAL           WM = _DTM_FIRST + 8
	DTM_SETMCFONT             WM = _DTM_FIRST + 9
	DTM_GETMCFONT             WM = _DTM_FIRST + 10
	DTM_SETMCSTYLE            WM = _DTM_FIRST + 11
	DTM_GETMCSTYLE            WM = _DTM_FIRST + 12
	DTM_CLOSEMONTHCAL         WM = _DTM_FIRST + 13
	DTM_GETDATETIMEPICKERINFO WM = _DTM_FIRST + 14
	DTM_GETIDEALSIZE          WM = _DTM_FIRST + 15
)

// Edit control [messages] (EM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-messages
const (
	EM_GETSEL              WM = 0x00b0
	EM_SETSEL              WM = 0x00b1
	EM_GETRECT             WM = 0x00b2
	EM_SETRECT             WM = 0x00b3
	EM_SETRECTNP           WM = 0x00b4
	EM_SCROLL              WM = 0x00b5
	EM_LINESCROLL          WM = 0x00b6
	EM_SCROLLCARET         WM = 0x00b7
	EM_GETMODIFY           WM = 0x00b8
	EM_SETMODIFY           WM = 0x00b9
	EM_GETLINECOUNT        WM = 0x00ba
	EM_LINEINDEX           WM = 0x00bb
	EM_SETHANDLE           WM = 0x00bc
	EM_GETHANDLE           WM = 0x00bd
	EM_GETTHUMB            WM = 0x00be
	EM_LINELENGTH          WM = 0x00c1
	EM_REPLACESEL          WM = 0x00c2
	EM_GETLINE             WM = 0x00c4
	EM_LIMITTEXT           WM = 0x00c5
	EM_CANUNDO             WM = 0x00c6
	EM_UNDO                WM = 0x00c7
	EM_FMTLINES            WM = 0x00c8
	EM_LINEFROMCHAR        WM = 0x00c9
	EM_SETTABSTOPS         WM = 0x00cb
	EM_SETPASSWORDCHAR     WM = 0x00cc
	EM_EMPTYUNDOBUFFER     WM = 0x00cd
	EM_GETFIRSTVISIBLELINE WM = 0x00ce
	EM_SETREADONLY         WM = 0x00cf
	EM_SETWORDBREAKPROC    WM = 0x00d0
	EM_GETWORDBREAKPROC    WM = 0x00d1
	EM_GETPASSWORDCHAR     WM = 0x00d2
	EM_SETMARGINS          WM = 0x00d3
	EM_GETMARGINS          WM = 0x00d4
	EM_SETLIMITTEXT        WM = EM_LIMITTEXT
	EM_GETLIMITTEXT        WM = 0x00d5
	EM_POSFROMCHAR         WM = 0x00d6
	EM_CHARFROMPOS         WM = 0x00d7
	EM_SETIMESTATUS        WM = 0x00d8
	EM_GETIMESTATUS        WM = 0x00d9

	_ECM_FIRST WM = 0x1500

	EM_SETCUEBANNER     WM = _ECM_FIRST + 1
	EM_GETCUEBANNER     WM = _ECM_FIRST + 2
	EM_SHOWBALLOONTIP   WM = _ECM_FIRST + 3
	EM_HIDEBALLOONTIP   WM = _ECM_FIRST + 4
	EM_SETHILITE        WM = _ECM_FIRST + 5
	EM_GETHILITE        WM = _ECM_FIRST + 6
	EM_NOSETFOCUS       WM = _ECM_FIRST + 7
	EM_TAKEFOCUS        WM = _ECM_FIRST + 8
	EM_SETEXTENDEDSTYLE WM = _ECM_FIRST + 10
	EM_GETEXTENDEDSTYLE WM = _ECM_FIRST + 11
	EM_SETENDOFLINE     WM = _ECM_FIRST + 12
	EM_GETENDOFLINE     WM = _ECM_FIRST + 13
	EM_ENABLESEARCHWEB  WM = _ECM_FIRST + 14
	EM_SEARCHWEB        WM = _ECM_FIRST + 15
	EM_SETCARETINDEX    WM = _ECM_FIRST + 17
	EM_GETCARETINDEX    WM = _ECM_FIRST + 18
	EM_GETZOOM          WM = WM_USER + 224
	EM_SETZOOM          WM = WM_USER + 225
	EM_FILELINEFROMCHAR WM = _ECM_FIRST + 19
	EM_FILELINEINDEX    WM = _ECM_FIRST + 20
	EM_FILELINELENGTH   WM = _ECM_FIRST + 21
	EM_GETFILELINE      WM = _ECM_FIRST + 22
	EM_GETFILELINECOUNT WM = _ECM_FIRST + 23
)

// Header control [messages] (HDM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-header-control-reference-messages
const (
	_HDM_FIRST WM = 0x1200

	HDM_GETITEMCOUNT WM = _HDM_FIRST + 0
	HDM_INSERTITEM   WM = _HDM_FIRST + 10
	HDM_DELETEITEM   WM = _HDM_FIRST + 2
	HDM_GETITEM      WM = _HDM_FIRST + 11
	HDM_SETITEM      WM = _HDM_FIRST + 12
	HDM_LAYOUT       WM = _HDM_FIRST + 5
)

// ListView control [messages] (LVM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-list-view-control-reference-messages
const (
	_LVM_FIRST WM = 0x1000

	LVM_GETBKCOLOR               WM = _LVM_FIRST + 0
	LVM_SETBKCOLOR               WM = _LVM_FIRST + 1
	LVM_GETIMAGELIST             WM = _LVM_FIRST + 2
	LVM_SETIMAGELIST             WM = _LVM_FIRST + 3
	LVM_GETITEMCOUNT             WM = _LVM_FIRST + 4
	LVM_DELETEITEM               WM = _LVM_FIRST + 8
	LVM_DELETEALLITEMS           WM = _LVM_FIRST + 9
	LVM_GETCALLBACKMASK          WM = _LVM_FIRST + 10
	LVM_SETCALLBACKMASK          WM = _LVM_FIRST + 11
	LVM_GETNEXTITEM              WM = _LVM_FIRST + 12
	LVM_GETITEMRECT              WM = _LVM_FIRST + 14
	LVM_SETITEMPOSITION          WM = _LVM_FIRST + 15
	LVM_GETITEMPOSITION          WM = _LVM_FIRST + 16
	LVM_HITTEST                  WM = _LVM_FIRST + 18
	LVM_ENSUREVISIBLE            WM = _LVM_FIRST + 19
	LVM_SCROLL                   WM = _LVM_FIRST + 20
	LVM_REDRAWITEMS              WM = _LVM_FIRST + 21
	LVM_ARRANGE                  WM = _LVM_FIRST + 22
	LVM_GETEDITCONTROL           WM = _LVM_FIRST + 24
	LVM_DELETECOLUMN             WM = _LVM_FIRST + 28
	LVM_GETCOLUMNWIDTH           WM = _LVM_FIRST + 29
	LVM_SETCOLUMNWIDTH           WM = _LVM_FIRST + 30
	LVM_GETHEADER                WM = _LVM_FIRST + 31
	LVM_CREATEDRAGIMAGE          WM = _LVM_FIRST + 33
	LVM_GETVIEWRECT              WM = _LVM_FIRST + 34
	LVM_GETTEXTCOLOR             WM = _LVM_FIRST + 35
	LVM_SETTEXTCOLOR             WM = _LVM_FIRST + 36
	LVM_GETTEXTBKCOLOR           WM = _LVM_FIRST + 37
	LVM_SETTEXTBKCOLOR           WM = _LVM_FIRST + 38
	LVM_GETTOPINDEX              WM = _LVM_FIRST + 39
	LVM_GETCOUNTPERPAGE          WM = _LVM_FIRST + 40
	LVM_GETORIGIN                WM = _LVM_FIRST + 41
	LVM_UPDATE                   WM = _LVM_FIRST + 42
	LVM_SETITEMSTATE             WM = _LVM_FIRST + 43
	LVM_GETITEMSTATE             WM = _LVM_FIRST + 44
	LVM_SETITEMCOUNT             WM = _LVM_FIRST + 47
	LVM_SORTITEMS                WM = _LVM_FIRST + 48
	LVM_SETITEMPOSITION32        WM = _LVM_FIRST + 49
	LVM_GETSELECTEDCOUNT         WM = _LVM_FIRST + 50
	LVM_GETITEMSPACING           WM = _LVM_FIRST + 51
	LVM_SETICONSPACING           WM = _LVM_FIRST + 53
	LVM_SETEXTENDEDLISTVIEWSTYLE WM = _LVM_FIRST + 54
	LVM_GETEXTENDEDLISTVIEWSTYLE WM = _LVM_FIRST + 55
	LVM_GETSUBITEMRECT           WM = _LVM_FIRST + 56
	LVM_SUBITEMHITTEST           WM = _LVM_FIRST + 57
	LVM_SETCOLUMNORDERARRAY      WM = _LVM_FIRST + 58
	LVM_GETCOLUMNORDERARRAY      WM = _LVM_FIRST + 59
	LVM_SETHOTITEM               WM = _LVM_FIRST + 60
	LVM_GETHOTITEM               WM = _LVM_FIRST + 61
	LVM_SETHOTCURSOR             WM = _LVM_FIRST + 62
	LVM_GETHOTCURSOR             WM = _LVM_FIRST + 63
	LVM_APPROXIMATEVIEWRECT      WM = _LVM_FIRST + 64
	LVM_SETWORKAREAS             WM = _LVM_FIRST + 65
	LVM_GETSELECTIONMARK         WM = _LVM_FIRST + 66
	LVM_SETSELECTIONMARK         WM = _LVM_FIRST + 67
	LVM_GETWORKAREAS             WM = _LVM_FIRST + 70
	LVM_SETHOVERTIME             WM = _LVM_FIRST + 71
	LVM_GETHOVERTIME             WM = _LVM_FIRST + 72
	LVM_GETNUMBEROFWORKAREAS     WM = _LVM_FIRST + 73
	LVM_SETTOOLTIPS              WM = _LVM_FIRST + 74
	LVM_GETITEM                  WM = _LVM_FIRST + 75
	LVM_SETITEM                  WM = _LVM_FIRST + 76
	LVM_INSERTITEM               WM = _LVM_FIRST + 77
	LVM_GETTOOLTIPS              WM = _LVM_FIRST + 78
	LVM_SORTITEMSEX              WM = _LVM_FIRST + 81
	LVM_FINDITEM                 WM = _LVM_FIRST + 83
	LVM_GETSTRINGWIDTH           WM = _LVM_FIRST + 87
	LVM_GETGROUPSTATE            WM = _LVM_FIRST + 92
	LVM_GETFOCUSEDGROUP          WM = _LVM_FIRST + 93
	LVM_GETCOLUMN                WM = _LVM_FIRST + 95
	LVM_SETCOLUMN                WM = _LVM_FIRST + 96
	LVM_INSERTCOLUMN             WM = _LVM_FIRST + 97
	LVM_GETGROUPRECT             WM = _LVM_FIRST + 98
	LVM_GETITEMTEXT              WM = _LVM_FIRST + 115
	LVM_SETITEMTEXT              WM = _LVM_FIRST + 116
	LVM_GETISEARCHSTRING         WM = _LVM_FIRST + 117
	LVM_EDITLABEL                WM = _LVM_FIRST + 118
	LVM_SETBKIMAGE               WM = _LVM_FIRST + 138
	LVM_GETBKIMAGE               WM = _LVM_FIRST + 139
	LVM_SETSELECTEDCOLUMN        WM = _LVM_FIRST + 140
	LVM_SETVIEW                  WM = _LVM_FIRST + 142
	LVM_GETVIEW                  WM = _LVM_FIRST + 143
	LVM_INSERTGROUP              WM = _LVM_FIRST + 145
	LVM_SETGROUPINFO             WM = _LVM_FIRST + 147
	LVM_GETGROUPINFO             WM = _LVM_FIRST + 149
	LVM_REMOVEGROUP              WM = _LVM_FIRST + 150
	LVM_MOVEGROUP                WM = _LVM_FIRST + 151
	LVM_GETGROUPCOUNT            WM = _LVM_FIRST + 152
	LVM_GETGROUPINFOBYINDEX      WM = _LVM_FIRST + 153
	LVM_MOVEITEMTOGROUP          WM = _LVM_FIRST + 154
	LVM_SETGROUPMETRICS          WM = _LVM_FIRST + 155
	LVM_GETGROUPMETRICS          WM = _LVM_FIRST + 156
	LVM_ENABLEGROUPVIEW          WM = _LVM_FIRST + 157
	LVM_SORTGROUPS               WM = _LVM_FIRST + 158
	LVM_INSERTGROUPSORTED        WM = _LVM_FIRST + 159
	LVM_REMOVEALLGROUPS          WM = _LVM_FIRST + 160
	LVM_HASGROUP                 WM = _LVM_FIRST + 161
	LVM_SETTILEVIEWINFO          WM = _LVM_FIRST + 162
	LVM_GETTILEVIEWINFO          WM = _LVM_FIRST + 163
	LVM_SETTILEINFO              WM = _LVM_FIRST + 164
	LVM_GETTILEINFO              WM = _LVM_FIRST + 165
	LVM_SETINSERTMARK            WM = _LVM_FIRST + 166
	LVM_GETINSERTMARK            WM = _LVM_FIRST + 167
	LVM_INSERTMARKHITTEST        WM = _LVM_FIRST + 168
	LVM_GETINSERTMARKRECT        WM = _LVM_FIRST + 169
	LVM_SETINSERTMARKCOLOR       WM = _LVM_FIRST + 170
	LVM_GETINSERTMARKCOLOR       WM = _LVM_FIRST + 171
	LVM_SETINFOTIP               WM = _LVM_FIRST + 173
	LVM_GETSELECTEDCOLUMN        WM = _LVM_FIRST + 174
	LVM_ISGROUPVIEWENABLED       WM = _LVM_FIRST + 175
	LVM_GETOUTLINECOLOR          WM = _LVM_FIRST + 176
	LVM_SETOUTLINECOLOR          WM = _LVM_FIRST + 177
	LVM_CANCELEDITLABEL          WM = _LVM_FIRST + 179
	LVM_MAPINDEXTOID             WM = _LVM_FIRST + 180
	LVM_MAPIDTOINDEX             WM = _LVM_FIRST + 181
	LVM_ISITEMVISIBLE            WM = _LVM_FIRST + 182
	LVM_GETEMPTYTEXT             WM = _LVM_FIRST + 204
	LVM_GETFOOTERRECT            WM = _LVM_FIRST + 205
	LVM_GETFOOTERINFO            WM = _LVM_FIRST + 206
	LVM_GETFOOTERITEMRECT        WM = _LVM_FIRST + 207
	LVM_GETFOOTERITEM            WM = _LVM_FIRST + 208
	LVM_GETITEMINDEXRECT         WM = _LVM_FIRST + 209
	LVM_SETITEMINDEXSTATE        WM = _LVM_FIRST + 210
	LVM_GETNEXTITEMINDEX         WM = _LVM_FIRST + 211
)

// MonthCalendar control [messages] (MCM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-month-calendar-control-reference-messages
const (
	_MCM_FIRST WM = 0x1000

	MCM_GETCURSEL           WM = _MCM_FIRST + 1
	MCM_SETCURSEL           WM = _MCM_FIRST + 2
	MCM_GETMAXSELCOUNT      WM = _MCM_FIRST + 3
	MCM_SETMAXSELCOUNT      WM = _MCM_FIRST + 4
	MCM_GETSELRANGE         WM = _MCM_FIRST + 5
	MCM_SETSELRANGE         WM = _MCM_FIRST + 6
	MCM_GETMONTHRANGE       WM = _MCM_FIRST + 7
	MCM_SETDAYSTATE         WM = _MCM_FIRST + 8
	MCM_GETMINREQRECT       WM = _MCM_FIRST + 9
	MCM_SETCOLOR            WM = _MCM_FIRST + 10
	MCM_GETCOLOR            WM = _MCM_FIRST + 11
	MCM_SETTODAY            WM = _MCM_FIRST + 12
	MCM_GETTODAY            WM = _MCM_FIRST + 13
	MCM_HITTEST             WM = _MCM_FIRST + 14
	MCM_SETFIRSTDAYOFWEEK   WM = _MCM_FIRST + 15
	MCM_GETFIRSTDAYOFWEEK   WM = _MCM_FIRST + 16
	MCM_GETRANGE            WM = _MCM_FIRST + 17
	MCM_SETRANGE            WM = _MCM_FIRST + 18
	MCM_GETMONTHDELTA       WM = _MCM_FIRST + 19
	MCM_SETMONTHDELTA       WM = _MCM_FIRST + 20
	MCM_GETMAXTODAYWIDTH    WM = _MCM_FIRST + 21
	MCM_SETUNICODEFORMAT    WM = CCM_SETUNICODEFORMAT
	MCM_GETUNICODEFORMAT    WM = CCM_GETUNICODEFORMAT
	MCM_GETCURRENTVIEW      WM = _MCM_FIRST + 22
	MCM_GETCALENDARCOUNT    WM = _MCM_FIRST + 23
	MCM_GETCALENDARGRIDINFO WM = _MCM_FIRST + 24
	MCM_GETCALID            WM = _MCM_FIRST + 27
	MCM_SETCALID            WM = _MCM_FIRST + 28
	MCM_SIZERECTTOMIN       WM = _MCM_FIRST + 29
	MCM_SETCALENDARBORDER   WM = _MCM_FIRST + 30
	MCM_GETCALENDARBORDER   WM = _MCM_FIRST + 31
	MCM_SETCURRENTVIEW      WM = _MCM_FIRST + 32
)

// ProgressBar control [messages] (PBM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-progress-bar-control-reference-messages
const (
	PBM_SETRANGE    WM = WM_USER + 1
	PBM_SETPOS      WM = WM_USER + 2
	PBM_DELTAPOS    WM = WM_USER + 3
	PBM_SETSTEP     WM = WM_USER + 4
	PBM_STEPIT      WM = WM_USER + 5
	PBM_SETRANGE32  WM = WM_USER + 6
	PBM_GETRANGE    WM = WM_USER + 7
	PBM_GETPOS      WM = WM_USER + 8
	PBM_SETBARCOLOR WM = WM_USER + 9
	PBM_SETBKCOLOR  WM = CCM_SETBKCOLOR
	PBM_SETMARQUEE  WM = WM_USER + 10
	PBM_GETSTEP     WM = WM_USER + 13
	PBM_GETBKCOLOR  WM = WM_USER + 14
	PBM_GETBARCOLOR WM = WM_USER + 15
	PBM_SETSTATE    WM = WM_USER + 16
	PBM_GETSTATE    WM = WM_USER + 17
)

// Status bar control [messages] (SB).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-status-bars-reference-messages
const (
	SB_SETTEXT          WM = WM_USER + 11
	SB_GETTEXT          WM = WM_USER + 13
	SB_GETTEXTLENGTH    WM = WM_USER + 12
	SB_SETPARTS         WM = WM_USER + 4
	SB_GETPARTS         WM = WM_USER + 6
	SB_GETBORDERS       WM = WM_USER + 7
	SB_SETMINHEIGHT     WM = WM_USER + 8
	SB_SIMPLE           WM = WM_USER + 9
	SB_GETRECT          WM = WM_USER + 10
	SB_ISSIMPLE         WM = WM_USER + 14
	SB_SETICON          WM = WM_USER + 15
	SB_SETTIPTEXT       WM = WM_USER + 17
	SB_GETTIPTEXT       WM = WM_USER + 19
	SB_GETICON          WM = WM_USER + 20
	SB_SETUNICODEFORMAT WM = CCM_SETUNICODEFORMAT
	SB_GETUNICODEFORMAT WM = CCM_GETUNICODEFORMAT
)

// Toolbar control [messages] (TB).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-toolbar-control-reference-messages
const (
	TB_ENABLEBUTTON          WM = WM_USER + 1
	TB_CHECKBUTTON           WM = WM_USER + 2
	TB_PRESSBUTTON           WM = WM_USER + 3
	TB_HIDEBUTTON            WM = WM_USER + 4
	TB_INDETERMINATE         WM = WM_USER + 5
	TB_MARKBUTTON            WM = WM_USER + 6
	TB_ISBUTTONENABLED       WM = WM_USER + 9
	TB_ISBUTTONCHECKED       WM = WM_USER + 10
	TB_ISBUTTONPRESSED       WM = WM_USER + 11
	TB_ISBUTTONHIDDEN        WM = WM_USER + 12
	TB_ISBUTTONINDETERMINATE WM = WM_USER + 13
	TB_ISBUTTONHIGHLIGHTED   WM = WM_USER + 14
	TB_SETSTATE              WM = WM_USER + 17
	TB_GETSTATE              WM = WM_USER + 18
	TB_ADDBITMAP             WM = WM_USER + 19
	TB_DELETEBUTTON          WM = WM_USER + 22
	TB_GETBUTTON             WM = WM_USER + 23
	TB_BUTTONCOUNT           WM = WM_USER + 24
	TB_COMMANDTOINDEX        WM = WM_USER + 25
	TB_SAVERESTORE           WM = WM_USER + 76
	TB_CUSTOMIZE             WM = WM_USER + 27
	TB_ADDSTRING             WM = WM_USER + 77
	TB_GETITEMRECT           WM = WM_USER + 29
	TB_BUTTONSTRUCTSIZE      WM = WM_USER + 30
	TB_SETBUTTONSIZE         WM = WM_USER + 31
	TB_SETBITMAPSIZE         WM = WM_USER + 32
	TB_AUTOSIZE              WM = WM_USER + 33
	TB_GETTOOLTIPS           WM = WM_USER + 35
	TB_SETTOOLTIPS           WM = WM_USER + 36
	TB_SETPARENT             WM = WM_USER + 37
	TB_SETROWS               WM = WM_USER + 39
	TB_GETROWS               WM = WM_USER + 40
	TB_SETCMDID              WM = WM_USER + 42
	TB_CHANGEBITMAP          WM = WM_USER + 43
	TB_GETBITMAP             WM = WM_USER + 44
	TB_GETBUTTONTEXT         WM = WM_USER + 75
	TB_REPLACEBITMAP         WM = WM_USER + 46
	TB_SETINDENT             WM = WM_USER + 47
	TB_SETIMAGELIST          WM = WM_USER + 48
	TB_GETIMAGELIST          WM = WM_USER + 49
	TB_LOADIMAGES            WM = WM_USER + 50
	TB_GETRECT               WM = WM_USER + 51
	TB_SETHOTIMAGELIST       WM = WM_USER + 52
	TB_GETHOTIMAGELIST       WM = WM_USER + 53
	TB_SETDISABLEDIMAGELIST  WM = WM_USER + 54
	TB_GETDISABLEDIMAGELIST  WM = WM_USER + 55
	TB_SETSTYLE              WM = WM_USER + 56
	TB_GETSTYLE              WM = WM_USER + 57
	TB_GETBUTTONSIZE         WM = WM_USER + 58
	TB_SETBUTTONWIDTH        WM = WM_USER + 59
	TB_SETMAXTEXTROWS        WM = WM_USER + 60
	TB_GETTEXTROWS           WM = WM_USER + 61
	TB_GETOBJECT             WM = WM_USER + 62
	TB_GETHOTITEM            WM = WM_USER + 71
	TB_SETHOTITEM            WM = WM_USER + 72
	TB_SETANCHORHIGHLIGHT    WM = WM_USER + 73
	TB_GETANCHORHIGHLIGHT    WM = WM_USER + 74
	TB_GETINSERTMARK         WM = WM_USER + 79
	TB_SETINSERTMARK         WM = WM_USER + 80
	TB_INSERTMARKHITTEST     WM = WM_USER + 81
	TB_MOVEBUTTON            WM = WM_USER + 82
	TB_GETMAXSIZE            WM = WM_USER + 83
	TB_SETEXTENDEDSTYLE      WM = WM_USER + 84
	TB_GETEXTENDEDSTYLE      WM = WM_USER + 85
	TB_GETPADDING            WM = WM_USER + 86
	TB_SETPADDING            WM = WM_USER + 87
	TB_SETINSERTMARKCOLOR    WM = WM_USER + 88
	TB_GETINSERTMARKCOLOR    WM = WM_USER + 89
	TB_SETCOLORSCHEME        WM = CCM_SETCOLORSCHEME
	TB_GETCOLORSCHEME        WM = CCM_GETCOLORSCHEME
	TB_SETUNICODEFORMAT      WM = CCM_SETUNICODEFORMAT
	TB_GETUNICODEFORMAT      WM = CCM_GETUNICODEFORMAT
	TB_MAPACCELERATOR        WM = WM_USER + 90
	TB_GETBITMAPFLAGS        WM = WM_USER + 41
	TB_GETBUTTONINFO         WM = WM_USER + 63
	TB_SETBUTTONINFO         WM = WM_USER + 64
	TB_INSERTBUTTON          WM = WM_USER + 67
	TB_ADDBUTTONS            WM = WM_USER + 68
	TB_HITTEST               WM = WM_USER + 69
	TB_SETDRAWTEXTFLAGS      WM = WM_USER + 70
	TB_GETSTRING             WM = WM_USER + 91
	TB_SETBOUNDINGSIZE       WM = WM_USER + 93
	TB_SETHOTITEM2           WM = WM_USER + 94
	TB_HASACCELERATOR        WM = WM_USER + 95
	TB_SETLISTGAP            WM = WM_USER + 96
	TB_GETIMAGELISTCOUNT     WM = WM_USER + 98
	TB_GETIDEALSIZE          WM = WM_USER + 99
	TB_GETMETRICS            WM = WM_USER + 101
	TB_SETMETRICS            WM = WM_USER + 102
	TB_GETITEMDROPDOWNRECT   WM = WM_USER + 103
	TB_SETPRESSEDIMAGELIST   WM = WM_USER + 104
	TB_GETPRESSEDIMAGELIST   WM = WM_USER + 105
	TB_SETWINDOWTHEME        WM = CCM_SETWINDOWTHEME
)

// Trackbar control [messages] (TBM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-messages
const (
	TBM_GETPOS           WM = WM_USER
	TBM_GETRANGEMIN      WM = WM_USER + 1
	TBM_GETRANGEMAX      WM = WM_USER + 2
	TBM_GETTIC           WM = WM_USER + 3
	TBM_SETTIC           WM = WM_USER + 4
	TBM_SETPOS           WM = WM_USER + 5
	TBM_SETRANGE         WM = WM_USER + 6
	TBM_SETRANGEMIN      WM = WM_USER + 7
	TBM_SETRANGEMAX      WM = WM_USER + 8
	TBM_CLEARTICS        WM = WM_USER + 9
	TBM_SETSEL           WM = WM_USER + 10
	TBM_SETSELSTART      WM = WM_USER + 11
	TBM_SETSELEND        WM = WM_USER + 12
	TBM_GETPTICS         WM = WM_USER + 14
	TBM_GETTICPOS        WM = WM_USER + 15
	TBM_GETNUMTICS       WM = WM_USER + 16
	TBM_GETSELSTART      WM = WM_USER + 17
	TBM_GETSELEND        WM = WM_USER + 18
	TBM_CLEARSEL         WM = WM_USER + 19
	TBM_SETTICFREQ       WM = WM_USER + 20
	TBM_SETPAGESIZE      WM = WM_USER + 21
	TBM_GETPAGESIZE      WM = WM_USER + 22
	TBM_SETLINESIZE      WM = WM_USER + 23
	TBM_GETLINESIZE      WM = WM_USER + 24
	TBM_GETTHUMBRECT     WM = WM_USER + 25
	TBM_GETCHANNELRECT   WM = WM_USER + 26
	TBM_SETTHUMBLENGTH   WM = WM_USER + 27
	TBM_GETTHUMBLENGTH   WM = WM_USER + 28
	TBM_SETTOOLTIPS      WM = WM_USER + 29
	TBM_GETTOOLTIPS      WM = WM_USER + 30
	TBM_SETTIPSIDE       WM = WM_USER + 31
	TBM_SETBUDDY         WM = WM_USER + 32
	TBM_GETBUDDY         WM = WM_USER + 33
	TBM_SETUNICODEFORMAT WM = CCM_SETUNICODEFORMAT
	TBM_GETUNICODEFORMAT WM = CCM_GETUNICODEFORMAT
)

// TreeView control [messages] (TVM).
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-tree-view-control-reference-messages
const (
	_TVM_FIRST WM = 0x1100

	TVM_INSERTITEM          WM = _TVM_FIRST + 50
	TVM_DELETEITEM          WM = _TVM_FIRST + 1
	TVM_EXPAND              WM = _TVM_FIRST + 2
	TVM_GETITEMRECT         WM = _TVM_FIRST + 4
	TVM_GETCOUNT            WM = _TVM_FIRST + 5
	TVM_GETINDENT           WM = _TVM_FIRST + 6
	TVM_SETINDENT           WM = _TVM_FIRST + 7
	TVM_GETIMAGELIST        WM = _TVM_FIRST + 8
	TVM_SETIMAGELIST        WM = _TVM_FIRST + 9
	TVM_GETNEXTITEM         WM = _TVM_FIRST + 10
	TVM_SELECTITEM          WM = _TVM_FIRST + 11
	TVM_GETITEM             WM = _TVM_FIRST + 62
	TVM_SETITEM             WM = _TVM_FIRST + 63
	TVM_EDITLABEL           WM = _TVM_FIRST + 65
	TVM_GETEDITCONTROL      WM = _TVM_FIRST + 15
	TVM_GETVISIBLECOUNT     WM = _TVM_FIRST + 16
	TVM_HITTEST             WM = _TVM_FIRST + 17
	TVM_CREATEDRAGIMAGE     WM = _TVM_FIRST + 18
	TVM_SORTCHILDREN        WM = _TVM_FIRST + 19
	TVM_ENSUREVISIBLE       WM = _TVM_FIRST + 20
	TVM_SORTCHILDRENCB      WM = _TVM_FIRST + 21
	TVM_ENDEDITLABELNOW     WM = _TVM_FIRST + 22
	TVM_GETISEARCHSTRING    WM = _TVM_FIRST + 64
	TVM_SETTOOLTIPS         WM = _TVM_FIRST + 24
	TVM_GETTOOLTIPS         WM = _TVM_FIRST + 25
	TVM_SETINSERTMARK       WM = _TVM_FIRST + 26
	TVM_SETUNICODEFORMAT    WM = CCM_SETUNICODEFORMAT
	TVM_GETUNICODEFORMAT    WM = CCM_GETUNICODEFORMAT
	TVM_SETITEMHEIGHT       WM = _TVM_FIRST + 27
	TVM_GETITEMHEIGHT       WM = _TVM_FIRST + 28
	TVM_SETBKCOLOR          WM = _TVM_FIRST + 29
	TVM_SETTEXTCOLOR        WM = _TVM_FIRST + 30
	TVM_GETBKCOLOR          WM = _TVM_FIRST + 31
	TVM_GETTEXTCOLOR        WM = _TVM_FIRST + 32
	TVM_SETSCROLLTIME       WM = _TVM_FIRST + 33
	TVM_GETSCROLLTIME       WM = _TVM_FIRST + 34
	TVM_SETINSERTMARKCOLOR  WM = _TVM_FIRST + 37
	TVM_GETINSERTMARKCOLOR  WM = _TVM_FIRST + 38
	TVM_SETBORDER           WM = _TVM_FIRST + 35
	TVM_GETITEMSTATE        WM = _TVM_FIRST + 39
	TVM_SETLINECOLOR        WM = _TVM_FIRST + 40
	TVM_GETLINECOLOR        WM = _TVM_FIRST + 41
	TVM_MAPACCIDTOHTREEITEM WM = _TVM_FIRST + 42
	TVM_MAPHTREEITEMTOACCID WM = _TVM_FIRST + 43
	TVM_SETEXTENDEDSTYLE    WM = _TVM_FIRST + 44
	TVM_GETEXTENDEDSTYLE    WM = _TVM_FIRST + 45
	TVM_SETAUTOSCROLLINFO   WM = _TVM_FIRST + 59
	TVM_SETHOT              WM = _TVM_FIRST + 58
	TVM_GETSELECTEDCOUNT    WM = _TVM_FIRST + 70
	TVM_SHOWINFOTIP         WM = _TVM_FIRST + 71
	TVM_GETITEMPARTRECT     WM = _TVM_FIRST + 72
)
