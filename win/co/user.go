//go:build windows

package co

// [ACCELL] fVirt.
//
// [ACCELL]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
type ACCELF uint8

const (
	ACCELF_NONE    ACCELF = 0
	ACCELF_VIRTKEY ACCELF = 1
	ACCELF_SHIFT   ACCELF = 0x04
	ACCELF_CONTROL ACCELF = 0x08
	ACCELF_ALT     ACCELF = 0x10
)

// [DdeInitialize] afCmd. Includes the prefixes APPCLASS, APPCMD, CBF and MF.
//
// [DdeInitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeinitializew
type AFCMD uint32

const (
	AFCMD_APPCLASS_MONITOR  AFCMD = 0x0000_0001
	AFCMD_APPCLASS_STANDARD AFCMD = 0x0000_0000

	AFCMD_APPCMD_CLIENTONLY  AFCMD = 0x0000_0010
	AFCMD_APPCMD_FILTERINITS AFCMD = 0x0000_0020

	AFCMD_CBF_FAIL_ALLSVRXACTIONS  AFCMD = 0x0003_f000
	AFCMD_CBF_FAIL_ADVISES         AFCMD = 0x0000_4000
	AFCMD_CBF_FAIL_CONNECTIONS     AFCMD = 0x0000_2000
	AFCMD_CBF_FAIL_EXECUTES        AFCMD = 0x0000_8000
	AFCMD_CBF_FAIL_POKES           AFCMD = 0x0001_0000
	AFCMD_CBF_FAIL_REQUESTS        AFCMD = 0x0002_0000
	AFCMD_CBF_FAIL_SELFCONNECTIONS AFCMD = 0x0000_1000

	AFCMD_CBF_SKIP_ALLNOTIFICATIONS AFCMD = 0x003c_0000
	AFCMD_CBF_SKIP_CONNECT_CONFIRMS AFCMD = 0x0004_0000
	AFCMD_CBF_SKIP_DISCONNECTS      AFCMD = 0x0020_0000
	AFCMD_CBF_SKIP_REGISTRATIONS    AFCMD = 0x0008_0000
	AFCMD_CBF_SKIP_UNREGISTRATIONS  AFCMD = 0x0010_0000

	AFCMD_MF_CALLBACKS AFCMD = 0x0800_0000
	AFCMD_MF_CONV      AFCMD = 0x4000_0000
	AFCMD_MF_ERRORS    AFCMD = 0x1000_0000
	AFCMD_MF_HSZ_INFO  AFCMD = 0x0100_0000
	AFCMD_MF_LINKS     AFCMD = 0x2000_0000
	AFCMD_MF_POSTMSGS  AFCMD = 0x0400_0000
	AFCMD_MF_SENDMSGS  AFCMD = 0x02000000
)

// [WM_APPCOMMAND] command.
//
// [WM_APPCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
type APPCOMMAND int16

const (
	APPCOMMAND_BROWSER_BACKWARD                  APPCOMMAND = 1
	APPCOMMAND_BROWSER_FORWARD                   APPCOMMAND = 2
	APPCOMMAND_BROWSER_REFRESH                   APPCOMMAND = 3
	APPCOMMAND_BROWSER_STOP                      APPCOMMAND = 4
	APPCOMMAND_BROWSER_SEARCH                    APPCOMMAND = 5
	APPCOMMAND_BROWSER_FAVORITES                 APPCOMMAND = 6
	APPCOMMAND_BROWSER_HOME                      APPCOMMAND = 7
	APPCOMMAND_VOLUME_MUTE                       APPCOMMAND = 8
	APPCOMMAND_VOLUME_DOWN                       APPCOMMAND = 9
	APPCOMMAND_VOLUME_UP                         APPCOMMAND = 10
	APPCOMMAND_MEDIA_NEXTTRACK                   APPCOMMAND = 11
	APPCOMMAND_MEDIA_PREVIOUSTRACK               APPCOMMAND = 12
	APPCOMMAND_MEDIA_STOP                        APPCOMMAND = 13
	APPCOMMAND_MEDIA_PLAY_PAUSE                  APPCOMMAND = 14
	APPCOMMAND_LAUNCH_MAIL                       APPCOMMAND = 15
	APPCOMMAND_LAUNCH_MEDIA_SELECT               APPCOMMAND = 16
	APPCOMMAND_LAUNCH_APP1                       APPCOMMAND = 17
	APPCOMMAND_LAUNCH_APP2                       APPCOMMAND = 18
	APPCOMMAND_BASS_DOWN                         APPCOMMAND = 19
	APPCOMMAND_BASS_BOOST                        APPCOMMAND = 20
	APPCOMMAND_BASS_UP                           APPCOMMAND = 21
	APPCOMMAND_TREBLE_DOWN                       APPCOMMAND = 22
	APPCOMMAND_TREBLE_UP                         APPCOMMAND = 23
	APPCOMMAND_MICROPHONE_VOLUME_MUTE            APPCOMMAND = 24
	APPCOMMAND_MICROPHONE_VOLUME_DOWN            APPCOMMAND = 25
	APPCOMMAND_MICROPHONE_VOLUME_UP              APPCOMMAND = 26
	APPCOMMAND_HELP                              APPCOMMAND = 27
	APPCOMMAND_FIND                              APPCOMMAND = 28
	APPCOMMAND_NEW                               APPCOMMAND = 29
	APPCOMMAND_OPEN                              APPCOMMAND = 30
	APPCOMMAND_CLOSE                             APPCOMMAND = 31
	APPCOMMAND_SAVE                              APPCOMMAND = 32
	APPCOMMAND_PRINT                             APPCOMMAND = 33
	APPCOMMAND_UNDO                              APPCOMMAND = 34
	APPCOMMAND_REDO                              APPCOMMAND = 35
	APPCOMMAND_COPY                              APPCOMMAND = 36
	APPCOMMAND_CUT                               APPCOMMAND = 37
	APPCOMMAND_PASTE                             APPCOMMAND = 38
	APPCOMMAND_REPLY_TO_MAIL                     APPCOMMAND = 39
	APPCOMMAND_FORWARD_MAIL                      APPCOMMAND = 40
	APPCOMMAND_SEND_MAIL                         APPCOMMAND = 41
	APPCOMMAND_SPELL_CHECK                       APPCOMMAND = 42
	APPCOMMAND_DICTATE_OR_COMMAND_CONTROL_TOGGLE APPCOMMAND = 43
	APPCOMMAND_MIC_ON_OFF_TOGGLE                 APPCOMMAND = 44
	APPCOMMAND_CORRECTION_LIST                   APPCOMMAND = 45
	APPCOMMAND_MEDIA_PLAY                        APPCOMMAND = 46
	APPCOMMAND_MEDIA_PAUSE                       APPCOMMAND = 47
	APPCOMMAND_MEDIA_RECORD                      APPCOMMAND = 48
	APPCOMMAND_MEDIA_FAST_FORWARD                APPCOMMAND = 49
	APPCOMMAND_MEDIA_REWIND                      APPCOMMAND = 50
	APPCOMMAND_MEDIA_CHANNEL_UP                  APPCOMMAND = 51
	APPCOMMAND_MEDIA_CHANNEL_DOWN                APPCOMMAND = 52
	APPCOMMAND_DELETE                            APPCOMMAND = 53
	APPCOMMAND_DWM_FLIP3D                        APPCOMMAND = 54
)

// Button control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/button-styles
type BS WS

const (
	BS_PUSHBUTTON      BS = 0x0000_0000
	BS_DEFPUSHBUTTON   BS = 0x0000_0001
	BS_CHECKBOX        BS = 0x0000_0002
	BS_AUTOCHECKBOX    BS = 0x0000_0003
	BS_RADIOBUTTON     BS = 0x0000_0004
	BS_3STATE          BS = 0x0000_0005
	BS_AUTO3STATE      BS = 0x0000_0006
	BS_GROUPBOX        BS = 0x0000_0007
	BS_USERBUTTON      BS = 0x0000_0008
	BS_AUTORADIOBUTTON BS = 0x0000_0009
	BS_PUSHBOX         BS = 0x0000_000a
	BS_OWNERDRAW       BS = 0x0000_000b
	BS_TYPEMASK        BS = 0x0000_000f
	BS_LEFTTEXT        BS = 0x0000_0020
	BS_TEXT            BS = 0x0000_0000
	BS_ICON            BS = 0x0000_0040
	BS_BITMAP          BS = 0x0000_0080
	BS_LEFT            BS = 0x0000_0100
	BS_RIGHT           BS = 0x0000_0200
	BS_CENTER          BS = 0x0000_0300
	BS_TOP             BS = 0x0000_0400
	BS_BOTTOM          BS = 0x0000_0800
	BS_VCENTER         BS = 0x0000_0c00
	BS_PUSHLIKE        BS = 0x0000_1000
	BS_MULTILINE       BS = 0x0000_2000
	BS_NOTIFY          BS = 0x0000_4000 // Button will send BN_DISABLE, BN_PUSHED, BN_KILLFOCUS, BN_PAINT, BN_SETFOCUS, and BN_UNPUSHED notifications.
	BS_FLAT            BS = 0x0000_8000
	BS_RIGHTBUTTON     BS = BS_LEFTTEXT
)

// [BroadcastSystemMessage] lpInfo.
//
// [BroadcastSystemMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
type BSM uint32

const (
	BSM_ALLCOMPONENTS BSM = 0x00000000
	BSM_ALLDESKTOPS   BSM = 0x00000010
	BSM_APPLICATIONS  BSM = 0x00000008
)

// [BroadcastSystemMessage] flags.
//
// [BroadcastSystemMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
type BSF uint32

const (
	BSF_ALLOWSFW           BSF = 0x0000_0080
	BSF_FLUSHDISK          BSF = 0x0000_0004
	BSF_FORCEIFHUNG        BSF = 0x0000_0020
	BSF_IGNORECURRENTTASK  BSF = 0x0000_0002
	BSF_NOHANG             BSF = 0x0000_0008
	BSF_NOTIMEOUTIFNOTHUNG BSF = 0x0000_0040
	BSF_POSTMESSAGE        BSF = 0x0000_0010
	BSF_QUERY              BSF = 0x0000_0001
	BSF_SENDNOTIFYMESSAGE  BSF = 0x0000_0100
)

// [IsDlgButtonChecked] return value, among others.
//
// [IsDlgButtonChecked]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
type BST uint32

const (
	BST_UNCHECKED     BST = 0x0000
	BST_CHECKED       BST = 0x0001
	BST_INDETERMINATE BST = 0x0002
	BST_PUSHED        BST = 0x0004
	BST_FOCUS         BST = 0x0008
)

// ComboBox [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/combo-box-styles
type CBS WS

const (
	CBS_SIMPLE            CBS = 0x0001
	CBS_DROPDOWN          CBS = 0x0002
	CBS_DROPDOWNLIST      CBS = 0x0003
	CBS_OWNERDRAWFIXED    CBS = 0x0010
	CBS_OWNERDRAWVARIABLE CBS = 0x0020
	CBS_AUTOHSCROLL       CBS = 0x0040
	CBS_OEMCONVERT        CBS = 0x0080
	CBS_SORT              CBS = 0x0100
	CBS_HASSTRINGS        CBS = 0x0200
	CBS_NOINTEGRALHEIGHT  CBS = 0x0400
	CBS_DISABLENOSCROLL   CBS = 0x0800
	CBS_UPPERCASE         CBS = 0x2000
	CBS_LOWERCASE         CBS = 0x4000
)

// [GetSysColor] nIndex.
//
// [GetSysColor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
type COLOR uint32

const (
	COLOR_SCROLLBAR               COLOR = 0
	COLOR_BACKGROUND              COLOR = 1
	COLOR_ACTIVECAPTION           COLOR = 2
	COLOR_INACTIVECAPTION         COLOR = 3
	COLOR_MENU                    COLOR = 4
	COLOR_WINDOW                  COLOR = 5
	COLOR_WINDOWFRAME             COLOR = 6
	COLOR_MENUTEXT                COLOR = 7
	COLOR_WINDOWTEXT              COLOR = 8
	COLOR_CAPTIONTEXT             COLOR = 9
	COLOR_ACTIVEBORDER            COLOR = 10
	COLOR_INACTIVEBORDER          COLOR = 11
	COLOR_APPWORKSPACE            COLOR = 12
	COLOR_HIGHLIGHT               COLOR = 13
	COLOR_HIGHLIGHTTEXT           COLOR = 14
	COLOR_BTNFACE                 COLOR = 15
	COLOR_BTNSHADOW               COLOR = 16
	COLOR_GRAYTEXT                COLOR = 17
	COLOR_BTNTEXT                 COLOR = 18
	COLOR_INACTIVECAPTIONTEXT     COLOR = 19
	COLOR_BTNHIGHLIGHT            COLOR = 20
	COLOR_3DDKSHADOW              COLOR = 21
	COLOR_3DLIGHT                 COLOR = 22
	COLOR_INFOTEXT                COLOR = 23
	COLOR_INFOBK                  COLOR = 24
	COLOR_HOTLIGHT                COLOR = 26
	COLOR_GRADIENTACTIVECAPTION   COLOR = 27
	COLOR_GRADIENTINACTIVECAPTION COLOR = 28
	COLOR_MENUHILIGHT             COLOR = 29
	COLOR_MENUBAR                 COLOR = 30
	COLOR_DESKTOP                 COLOR = COLOR_BACKGROUND
	COLOR_3DFACE                  COLOR = COLOR_BTNFACE
	COLOR_3DSHADOW                COLOR = COLOR_BTNSHADOW
	COLOR_3DHIGHLIGHT             COLOR = COLOR_BTNHIGHLIGHT
	COLOR_3DHILIGHT               COLOR = COLOR_BTNHIGHLIGHT
	COLOR_BTNHILIGHT              COLOR = COLOR_BTNHIGHLIGHT
)

// Window class [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/winmsg/window-class-styles
type CS uint32

const (
	CS_VREDRAW         CS = 0x0001
	CS_HREDRAW         CS = 0x0002
	CS_DBLCLKS         CS = 0x0008
	CS_OWNDC           CS = 0x0020
	CS_CLASSDC         CS = 0x0040
	CS_PARENTDC        CS = 0x0080
	CS_NOCLOSE         CS = 0x0200
	CS_SAVEBITS        CS = 0x0800
	CS_BYTEALIGNCLIENT CS = 0x1000
	CS_BYTEALIGNWINDOW CS = 0x2000
	CS_GLOBALCLASS     CS = 0x4000
	CS_DROPSHADOW      CS = 0x0002_0000
)

// [ChildWindowFromPointEx] flags.
//
// [ChildWindowFromPointEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-childwindowfrompointex
type CWP uint32

const (
	CWP_ALL             CWP = 0x0000
	CWP_SKIPDISABLED    CWP = 0x0002
	CWP_SKIPINVISIBLE   CWP = 0x0001
	CWP_SKIPTRANSPARENT CWP = 0x0004
)

// [DrawIconEx] diFlags.
//
// [DrawIconEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawiconex
type DI uint32

const (
	DI_COMPAT      DI = 0x0004
	DI_DEFAULTSIZE DI = 0x0008
	DI_IMAGE       DI = 0x0002
	DI_MASK        DI = 0x0001
	DI_NOMIRROR    DI = 0x0010
	DI_NORMAL      DI = 0x0003
)

// [EnumDisplayDevices] flags.
//
// [EnumDisplayDevices]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-display_devicew
type DISPLAY_DEVICE uint32

const (
	DISPLAY_DEVICE_ATTACHED_TO_DESKTOP DISPLAY_DEVICE = 0x0000_0001
	DISPLAY_DEVICE_MULTI_DRIVER        DISPLAY_DEVICE = 0x0000_0002
	DISPLAY_DEVICE_PRIMARY_DEVICE      DISPLAY_DEVICE = 0x0000_0004
	DISPLAY_DEVICE_MIRRORING_DRIVER    DISPLAY_DEVICE = 0x0000_0008
	DISPLAY_DEVICE_VGA_COMPATIBLE      DISPLAY_DEVICE = 0x0000_0010
	DISPLAY_DEVICE_REMOVABLE           DISPLAY_DEVICE = 0x0000_0020
	DISPLAY_DEVICE_ACC_DRIVER          DISPLAY_DEVICE = 0x0000_0040
	DISPLAY_DEVICE_MODESPRUNED         DISPLAY_DEVICE = 0x0800_0000
	DISPLAY_DEVICE_RDPUDD              DISPLAY_DEVICE = 0x0100_0000
	DISPLAY_DEVICE_REMOTE              DISPLAY_DEVICE = 0x0400_0000
	DISPLAY_DEVICE_DISCONNECT          DISPLAY_DEVICE = 0x0200_0000
	DISPLAY_DEVICE_TS_COMPATIBLE       DISPLAY_DEVICE = 0x0020_0000
	DISPLAY_DEVICE_UNSAFE_MODES_ON     DISPLAY_DEVICE = 0x0008_0000

	DISPLAY_DEVICE_ACTIVE   DISPLAY_DEVICE = 0x0000_0001
	DISPLAY_DEVICE_ATTACHED DISPLAY_DEVICE = 0x0000_0002
)

// [WM_GETDLGCODE] return value.
//
// [WM_GETDLGCODE]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
type DLGC uint32

const (
	DLGC_NONE            DLGC = 0
	DLGC_WANTARROWS      DLGC = 0x0001
	DLGC_WANTTAB         DLGC = 0x0002
	DLGC_WANTALLKEYS     DLGC = 0x0004
	DLGC_WANTMESSAGE     DLGC = 0x0004
	DLGC_HASSETSEL       DLGC = 0x0008
	DLGC_DEFPUSHBUTTON   DLGC = 0x0010
	DLGC_UNDEFPUSHBUTTON DLGC = 0x0020
	DLGC_RADIOBUTTON     DLGC = 0x0040
	DLGC_WANTCHARS       DLGC = 0x0080
	DLGC_STATIC          DLGC = 0x0100
	DLGC_BUTTON          DLGC = 0x2000
)

// [DdeNameService] opts. Originally with DNS prefix.
//
// [DdeNameService]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddenameservice
type DDENS uint32

const (
	DDENS_REGISTER   DDENS = 0x0001
	DDENS_UNREGISTER DDENS = 0x0002
	DDENS_FILTERON   DDENS = 0x0003
	DDENS_FILTEROFF  DDENS = 0x0004
)

// [SetProcessDpiAwarenessContext] value.
//
// [SetProcessDpiAwarenessContext]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
type DPI_AWARE_CTX int32

const (
	DPI_AWARE_CTX_UNAWARE           DPI_AWARE_CTX = -1
	DPI_AWARE_CTX_SYSTEM_AWARE      DPI_AWARE_CTX = -2
	DPI_AWARE_CTX_PER_MON_AWARE     DPI_AWARE_CTX = -3
	DPI_AWARE_CTX_PER_MON_AWARE_V2  DPI_AWARE_CTX = -4
	DPI_AWARE_CTX_UNAWARE_GDISCALED DPI_AWARE_CTX = -5
)

// [EnumDisplayDevices] flags.
//
// [EnumDisplayDevices]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaydevicesw
type EDD uint32

const (
	EDD_NONE                      EDD = 0
	EDD_GET_DEVICE_INTERFACE_NAME EDD = 0x0000_0001
)

// Edit control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/edit-control-styles
type ES WS

const (
	ES_LEFT        ES = 0x0000
	ES_CENTER      ES = 0x0001
	ES_RIGHT       ES = 0x0002
	ES_MULTILINE   ES = 0x0004
	ES_UPPERCASE   ES = 0x0008
	ES_LOWERCASE   ES = 0x0010
	ES_PASSWORD    ES = 0x0020
	ES_AUTOVSCROLL ES = 0x0040
	ES_AUTOHSCROLL ES = 0x0080
	ES_NOHIDESEL   ES = 0x0100
	ES_OEMCONVERT  ES = 0x0400
	ES_READONLY    ES = 0x0800
	ES_WANTRETURN  ES = 0x1000
	ES_NUMBER      ES = 0x2000
)

// [WM_APPCOMMAND] input event.
//
// [WM_APPCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
type FAPPCOMMAND uint32

const (
	FAPPCOMMAND_MOUSE FAPPCOMMAND = 0x8000
	FAPPCOMMAND_KEY   FAPPCOMMAND = 0
	FAPPCOMMAND_OEM   FAPPCOMMAND = 0x1000
)

// [GetAncestor] gaFlags.
//
// [GetAncestor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
type GA uint32

const (
	// Retrieves the parent window. This does not include the owner as it does
	// with the win.HWND.GetParent() function.
	GA_PARENT GA = 1
	// Retrieves the root window by walking the chain of parent windows. Returns
	// the closest parent with WS_OVERLAPPED or WS_POPUP.
	//
	// https://groups.google.com/a/chromium.org/g/chromium-dev/c/Hirr_DkuZdw/m/N0pSoJBhAAAJ
	GA_ROOT GA = 2
	// Retrieves the owned root window by walking the chain of parent and owner
	// windows returned by win.HWND.GetParent().
	//
	// Returns the furthest parent with WS_OVERLAPPED or WS_POPUP which usually
	// is the main application window.
	//
	// https://groups.google.com/a/chromium.org/g/chromium-dev/c/Hirr_DkuZdw/m/N0pSoJBhAAAJ
	GA_ROOTOWNER GA = 3
)

// [GetClassLong] nIndex. Includes values with GCW prefix.
//
// [GetClassLong]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongw
type GCL int32

const (
	GCL_ATOM          GCL = -32 // Originally GCW_ATOM.
	GCL_CBCLSEXTRA    GCL = -20
	GCL_CBWNDEXTRA    GCL = -18
	GCL_HBRBACKGROUND GCL = -10
	GCL_HCURSOR       GCL = -12
	GCL_HICON         GCL = -14
	GCL_HICONSM       GCL = -34
	GCL_HMODULE       GCL = -16
	GCL_MENUNAME      GCL = -8
	GCL_STYLE         GCL = -26
	GCL_WNDPROC       GCL = -24
)

// [GetMenuDefaultItem] gmdiFlags.
//
// [GetMenuDefaultItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmenudefaultitem
type GMDI uint32

const (
	GMDI_GOINTOPOPUPS GMDI = 0x0002
	GMDI_USEDISABLED  GMDI = 0x0001
)

// [GUITHREADINFO] flags.
//
// [GUITHREADINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-guithreadinfo
type GUI uint32

const (
	GUI_CARETBLINKING  GUI = 0x0000_0001
	GUI_INMENUMODE     GUI = 0x0000_0004
	GUI_INMOVESIZE     GUI = 0x0000_0002
	GUI_POPUPMENUMODE  GUI = 0x0000_0010
	GUI_SYSTEMMENUMODE GUI = 0x0000_0008
)

// [GetWindow] uCmd.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
type GW uint32

const (
	GW_HWNDFIRST    GW = 0
	GW_HWNDLAST     GW = 1
	GW_HWNDNEXT     GW = 2
	GW_HWNDPREV     GW = 3
	GW_OWNER        GW = 4
	GW_CHILD        GW = 5
	GW_ENABLEDPOPUP GW = 6
)

// [SetWindowsHookEx] callback CBT hook codes.
//
// [SetWindowsHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
type HCBT int32

const (
	HCBT_MOVESIZE     HCBT = 0
	HCBT_MINMAX       HCBT = 1
	HCBT_QS           HCBT = 2
	HCBT_CREATEWND    HCBT = 3
	HCBT_DESTROYWND   HCBT = 4
	HCBT_ACTIVATE     HCBT = 5
	HCBT_CLICKSKIPPED HCBT = 6
	HCBT_KEYSKIPPED   HCBT = 7
	HCBT_SYSCOMMAND   HCBT = 8
	HCBT_SETFOCUS     HCBT = 9
)

// [HELPINFO] iContextType.
//
// [HELPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
type HELPINFO int32

const (
	HELPINFO_WINDOW   HELPINFO = 0x0001
	HELPINFO_MENUITEM HELPINFO = 0x0002
)

// [WM_NCHITTEST] return value.
//
// [WM_NCHITTEST]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
type HT int32

const (
	HT_ERROR       HT = -2
	HT_TRANSPARENT HT = -1
	HT_NOWHERE     HT = 0
	HT_CLIENT      HT = 1
	HT_CAPTION     HT = 2
	HT_SYSMENU     HT = 3
	HT_GROWBOX     HT = 4
	HT_SIZE        HT = HT_GROWBOX
	HT_MENU        HT = 5
	HT_HSCROLL     HT = 6
	HT_VSCROLL     HT = 7
	HT_MINBUTTON   HT = 8
	HT_MAXBUTTON   HT = 9
	HT_LEFT        HT = 10
	HT_RIGHT       HT = 11
	HT_TOP         HT = 12
	HT_TOPLEFT     HT = 13
	HT_TOPRIGHT    HT = 14
	HT_BOTTOM      HT = 15
	HT_BOTTOMLEFT  HT = 16
	HT_BOTTOMRIGHT HT = 17
	HT_BORDER      HT = 18
	HT_REDUCE      HT = HT_MINBUTTON
	HT_ZOOM        HT = HT_MAXBUTTON
	HT_SIZEFIRST   HT = HT_LEFT
	HT_SIZELAST    HT = HT_BOTTOMRIGHT
	HT_OBJECT      HT = 19
	HT_CLOSE       HT = 20
	HT_HELP        HT = 21
)

// [SetWindowPos] hwndInsertAfter. Can be converted to HWND.
//
// [SetWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
type HWND_IA int32

const (
	HWND_IA_NONE      HWND_IA = 0
	HWND_IA_BOTTOM    HWND_IA = 1
	HWND_IA_NOTOPMOST HWND_IA = -2
	HWND_IA_TOP       HWND_IA = 0
	HWND_IA_TOPMOST   HWND_IA = -1
)

// [WM_SETICON] icon size. Originally with ICON prefix.
//
// [WM_SETICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
type ICON_SZ uint8

const (
	ICON_SZ_SMALL  ICON_SZ = 0
	ICON_SZ_BIG    ICON_SZ = 1
	ICON_SZ_SMALL2 ICON_SZ = 2
)

// Dialog codes returned by [MessageBox].
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
type ID int32

const (
	ID_ABORT    ID = 3
	ID_CANCEL   ID = 2
	ID_CONTINUE ID = 11
	ID_IGNORE   ID = 5
	ID_NO       ID = 7
	ID_OK       ID = 1
	ID_RETRY    ID = 4
	ID_TRYAGAIN ID = 10
	ID_YES      ID = 6
)

// [LoadCursor] lpCursorName.
//
// [LoadCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
type IDC uintptr

const (
	IDC_ARROW       IDC = 32512
	IDC_IBEAM       IDC = 32513
	IDC_WAIT        IDC = 32514
	IDC_CROSS       IDC = 32515
	IDC_UPARROW     IDC = 32516
	IDC_SIZENWSE    IDC = 32642
	IDC_SIZENESW    IDC = 32643
	IDC_SIZEWE      IDC = 32644
	IDC_SIZENS      IDC = 32645
	IDC_SIZEALL     IDC = 32646
	IDC_NO          IDC = 32648
	IDC_HAND        IDC = 32649
	IDC_APPSTARTING IDC = 32650
	IDC_HELP        IDC = 32651
	IDC_PIN         IDC = 32671
	IDC_PERSON      IDC = 32672
)

// [WM_HOTKEY] identifier.
//
// [WM_HOTKEY]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
type IDHOT int32

const (
	IDHOT_SNAPWINDOW  IDHOT = -1
	IDHOT_SNAPDESKTOP IDHOT = -2
)

// [LoadIcon] lpIconName.
//
// [LoadIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
type IDI uintptr

const (
	IDI_APPLICATION IDI = 32512
	IDI_HAND        IDI = 32513
	IDI_QUESTION    IDI = 32514
	IDI_EXCLAMATION IDI = 32515
	IDI_ASTERISK    IDI = 32516
	IDI_WINLOGO     IDI = 32517
	IDI_SHIELD      IDI = 32518
	IDI_WARNING     IDI = IDI_EXCLAMATION
	IDI_ERROR       IDI = IDI_HAND
	IDI_INFORMATION IDI = IDI_ASTERISK
)

// [LoadImage] type.
//
// [LoadImage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
type IMAGE uint32

const (
	IMAGE_BITMAP      IMAGE = 0
	IMAGE_ICON        IMAGE = 1
	IMAGE_CURSOR      IMAGE = 2
	IMAGE_ENHMETAFILE IMAGE = 3
)

// [InSendMessageEx] return value.
//
// [InSendMessageEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
type ISMEX uint32

const (
	ISMEX_NOSEND   ISMEX = 0x0000_0000
	ISMEX_CALLBACK ISMEX = 0x0000_0004
	ISMEX_NOTIFY   ISMEX = 0x0000_0002
	ISMEX_REPLIED  ISMEX = 0x0000_0008
	ISMEX_SEND     ISMEX = 0x0000_0001
)

// [SetProcessDefaultLayout] dwDefaultLayout.
//
// [SetProcessDefaultLayout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
type LAYOUT uint32

const (
	LAYOUT_NORMAL LAYOUT = 0
	LAYOUT_RTL    LAYOUT = 0x0000_0001
)

// [LoadImage] fuLoad.
//
// [LoadImage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
type LR uint32

const (
	LR_DEFAULTCOLOR     LR = 0x0000_0000
	LR_MONOCHROME       LR = 0x0000_0001
	LR_COLOR            LR = 0x0000_0002
	LR_COPYRETURNORG    LR = 0x0000_0004
	LR_COPYDELETEORG    LR = 0x0000_0008
	LR_LOADFROMFILE     LR = 0x0000_0010
	LR_LOADTRANSPARENT  LR = 0x0000_0020
	LR_DEFAULTSIZE      LR = 0x0000_0040
	LR_VGACOLOR         LR = 0x0000_0080
	LR_LOADMAP3DCOLORS  LR = 0x0000_1000
	LR_CREATEDIBSECTION LR = 0x0000_2000
	LR_COPYFROMRESOURCE LR = 0x0000_4000
	LR_SHARED           LR = 0x0000_8000
)

// [LockSetForegroundWindow] uLockCode.
//
// [LockSetForegroundWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
type LSFW uint32

const (
	LSFW_LOCK   LSFW = 1
	LSFW_UNLOCK LSFW = 2
)

// [SetLayeredWindowAttributes] flags.
//
// [SetLayeredWindowAttributes]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setlayeredwindowattributes
type LWA uint32

const (
	LWA_ALPHA    LWA = 0x0000_0002
	LWA_COLORKEY LWA = 0x0000_0001
)

// [MessageBox] uType.
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
type MB uint32

const (
	MB_ABORTRETRYIGNORE  MB = 0x0000_0002
	MB_CANCELTRYCONTINUE MB = 0x0000_0006
	MB_HELP              MB = 0x0000_4000
	MB_OK                MB = 0x0000_0000
	MB_OKCANCEL          MB = 0x0000_0001
	MB_RETRYCANCEL       MB = 0x0000_0005
	MB_YESNO             MB = 0x0000_0004
	MB_YESNOCANCEL       MB = 0x0000_0003

	MB_ICONEXCLAMATION MB = 0x0000_0030
	MB_ICONWARNING     MB = 0x0000_0030
	MB_ICONINFORMATION MB = 0x0000_0040
	MB_ICONASTERISK    MB = 0x0000_0040
	MB_ICONQUESTION    MB = 0x0000_0020
	MB_ICONSTOP        MB = 0x0000_0010
	MB_ICONERROR       MB = 0x0000_0010
	MB_ICONHAND        MB = 0x0000_0010

	MB_DEFBUTTON1 MB = 0x0000_0000
	MB_DEFBUTTON2 MB = 0x0000_0100
	MB_DEFBUTTON3 MB = 0x0000_0200
	MB_DEFBUTTON4 MB = 0x0000_0300

	MB_APPLMODAL   MB = 0x0000_0000
	MB_SYSTEMMODAL MB = 0x0000_1000
	MB_TASKMODAL   MB = 0x0000_2000

	MB_DEFAULT_DESKTOP_ONLY MB = 0x0002_0000
	MB_RIGHT                MB = 0x0008_0000
	MB_RTLREADING           MB = 0x0010_0000
	MB_SETFOREGROUND        MB = 0x0001_0000
	MB_TOPMOST              MB = 0x0004_0000
	MB_SERVICE_NOTIFICATION MB = 0x0020_0000
)

// [WM_MENUCHAR] menu type. Originally with MF prefix.
//
// [WM_MENUCHAR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
type MFMC uint16

const (
	POPUP   MFMC = 0x0000_0010
	SYSMENU MFMC = 0x0000_2000
)

// [CheckMenuItem] uCheck, among others.
//
// [CheckMenuItem]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
type MF uint32

const (
	MF_INSERT          MF = 0x0000_0000
	MF_CHANGE          MF = 0x0000_0080
	MF_APPEND          MF = 0x0000_0100
	MF_DELETE          MF = 0x0000_0200
	MF_REMOVE          MF = 0x0000_1000
	MF_BYCOMMAND       MF = 0x0000_0000
	MF_BYPOSITION      MF = 0x0000_0400
	MF_SEPARATOR       MF = 0x0000_0800
	MF_ENABLED         MF = 0x0000_0000
	MF_GRAYED          MF = 0x0000_0001
	MF_DISABLED        MF = 0x0000_0002
	MF_UNCHECKED       MF = 0x0000_0000
	MF_CHECKED         MF = 0x0000_0008
	MF_USECHECKBITMAPS MF = 0x0000_0200
	MF_STRING          MF = 0x0000_0000
	MF_BITMAP          MF = 0x0000_0004
	MF_OWNERDRAW       MF = 0x0000_0100
	MF_POPUP           MF = 0x0000_0010
	MF_MENUBARBREAK    MF = 0x0000_0020
	MF_MENUBREAK       MF = 0x0000_0040
	MF_UNHILITE        MF = 0x0000_0000
	MF_HILITE          MF = 0x0000_0080
	MF_DEFAULT         MF = 0x0000_1000
	MF_SYSMENU         MF = 0x0000_2000
	MF_HELP            MF = 0x0000_4000
	MF_RIGHTJUSTIFY    MF = 0x0000_4000
	MF_MOUSESELECT     MF = 0x0000_8000
)

// [MENUITEMINFO] fState.
//
// [MENUITEMINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MFS uint32

const (
	MFS_GRAYED    MFS = 0x0000_0003
	MFS_DISABLED  MFS = MFS_GRAYED
	MFS_CHECKED   MFS = MFS(MF_CHECKED)
	MFS_HILITE    MFS = MFS(MF_HILITE)
	MFS_ENABLED   MFS = MFS(MF_ENABLED)
	MFS_UNCHECKED MFS = MFS(MF_UNCHECKED)
	MFS_UNHILITE  MFS = MFS(MF_UNHILITE)
	MFS_DEFAULT   MFS = MFS(MF_DEFAULT)
)

// [MENUITEMINFO] fType.
//
// [MENUITEMINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MFT uint32

const (
	MFT_STRING       MFT = MFT(MF_STRING)
	MFT_BITMAP       MFT = MFT(MF_BITMAP)
	MFT_MENUBARBREAK MFT = MFT(MF_MENUBARBREAK)
	MFT_MENUBREAK    MFT = MFT(MF_MENUBREAK)
	MFT_OWNERDRAW    MFT = MFT(MF_OWNERDRAW)
	MFT_RADIOCHECK   MFT = 0x0000_0200
	MFT_SEPARATOR    MFT = MFT(MF_SEPARATOR)
	MFT_RIGHTORDER   MFT = 0x0000_2000
	MFT_RIGHTJUSTIFY MFT = MFT(MF_RIGHTJUSTIFY)
)

// [MENUITEMINFO] fMask.
//
// [MENUITEMINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MIIM uint32

const (
	MIIM_STATE      MIIM = 0x0000_0001
	MIIM_ID         MIIM = 0x0000_0002
	MIIM_SUBMENU    MIIM = 0x0000_0004
	MIIM_CHECKMARKS MIIM = 0x0000_0008
	MIIM_TYPE       MIIM = 0x0000_0010
	MIIM_DATA       MIIM = 0x0000_0020
	MIIM_STRING     MIIM = 0x0000_0040
	MIIM_BITMAP     MIIM = 0x0000_0080
	MIIM_FTYPE      MIIM = 0x0000_0100
)

// [MENUINFO] fMask.
//
// [MENUINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
type MIM uint32

const (
	MIM_MAXHEIGHT       MIM = 0x0000_0001
	MIM_BACKGROUND      MIM = 0x0000_0002
	MIM_HELPID          MIM = 0x0000_0004
	MIM_MENUDATA        MIM = 0x0000_0008
	MIM_STYLE           MIM = 0x0000_0010
	MIM_APPLYTOSUBMENUS MIM = 0x8000_0000
)

// [WM_LBUTTONDOWN] virtual keys, among others
//
// [WM_LBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
type MK uint16

const (
	MK_LBUTTON  MK = 0x0001
	MK_RBUTTON  MK = 0x0002
	MK_SHIFT    MK = 0x0004
	MK_CONTROL  MK = 0x0008
	MK_MBUTTON  MK = 0x0010
	MK_XBUTTON1 MK = 0x0020
	MK_XBUTTON2 MK = 0x0040
)

// [MONITORINFO] dwFlags
//
// [MONITORINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-monitorinfo
type MONITORINFOF uint32

const (
	MONITORINFOF_OTHER   MONITORINFOF = 0
	MONITORINFOF_PRIMARY MONITORINFOF = 0x0000_0001
)

// [WM_MENUCHAR] return value.
//
// [WM_MENUCHAR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
type MNC uint32

const (
	MNC_IGNORE  MNC = 0
	MNC_CLOSE   MNC = 1
	MNC_EXECUTE MNC = 2
	MNC_SELECT  MNC = 3
)

// [WM_MENUDRAG] return value.
//
// [WM_MENUDRAG]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
type MND uint32

const (
	MND_CONTINUE MND = 0
	MND_ENDMENU  MND = 1
)

// [WM_MENUGETOBJECT] return value.
//
// [WM_MENUGETOBJECT]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
type MNGO uint32

const (
	MNGO_NOINTERFACE MNGO = 0x0000_0000
	MNGO_NOERROR     MNGO = 0x0000_0001
)

// [MENUGETOBJECTINFO] dwFlags.
//
// [MENUGETOBJECTINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menugetobjectinfo
type MNGOF uint32

const (
	MNGOF_TOPGAP    MNGOF = 0x0000_0001
	MNGOF_BOTTOMGAP MNGOF = 0x0000_0002
)

// [MENUINFO] dwStyle.
//
// [MENUINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
type MNS uint32

const (
	MNS_NOCHECK     MNS = 0x8000_0000
	MNS_MODELESS    MNS = 0x4000_0000
	MNS_DRAGDROP    MNS = 0x2000_0000
	MNS_AUTODISMISS MNS = 0x1000_0000
	MNS_NOTIFYBYPOS MNS = 0x0800_0000
	MNS_CHECKORBMP  MNS = 0x0400_0000
)

// [WM_HOTKEY] combined keys.
//
// [WM_HOTKEY]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
type MOD uint16

const (
	MOD_ALT     MOD = 0x0001
	MOD_CONTROL MOD = 0x0002
	MOD_SHIFT   MOD = 0x0004
	MOD_WIN     MOD = 0x0008
)

// [MonitorFromPoint] dwFlags.
//
// [MonitorFromPoint]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
type MONITOR uint32

const (
	MONITOR_DEFAULTTONULL    MONITOR = 0x0000_0000
	MONITOR_DEFAULTTOPRIMARY MONITOR = 0x0000_0001
	MONITOR_DEFAULTTONEAREST MONITOR = 0x0000_0002
)

// [WM_ENTERIDLE] displayed.
//
// [WM_ENTERIDLE]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
type MSGF uint32

const (
	MSGF_DIALOGBOX MSGF = 0
	MSGF_MENU      MSGF = 2
)

// [DRAWITEMSTRUCT] itemAction.
//
// [DRAWITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type ODA uint32

const (
	ODA_DRAWENTIRE ODA = 0x0001
	ODA_SELECT     ODA = 0x0002
	ODA_FOCUS      ODA = 0x0004
)

// [SetSystemCursor] id.
//
// [SetSystemCursor]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setsystemcursor
type OCR uint32

const (
	OCR_APPSTARTING OCR = 32650
	OCR_NORMAL      OCR = 32512
	OCR_CROSS       OCR = 32515
	OCR_HAND        OCR = 32649
	OCR_HELP        OCR = 32651
	OCR_IBEAM       OCR = 32513
	OCR_NO          OCR = 32648
	OCR_SIZEALL     OCR = 32646
	OCR_SIZENESW    OCR = 32643
	OCR_SIZENS      OCR = 32645
	OCR_SIZENWSE    OCR = 32642
	OCR_SIZEWE      OCR = 32644
	OCR_UP          OCR = 32516
	OCR_WAIT        OCR = 32514
)

// [DRAWITEMSTRUCT] itemState.
//
// [DRAWITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type ODS uint32

const (
	ODS_SELECTED     ODS = 0x0001
	ODS_GRAYED       ODS = 0x0002
	ODS_DISABLED     ODS = 0x0004
	ODS_CHECKED      ODS = 0x0008
	ODS_FOCUS        ODS = 0x0010
	ODS_DEFAULT      ODS = 0x0020
	ODS_COMBOBOXEDIT ODS = 0x1000
	ODS_HOTLIGHT     ODS = 0x0040
	ODS_INACTIVE     ODS = 0x0080
	ODS_NOACCEL      ODS = 0x0100
	ODS_NOFOCUSRECT  ODS = 0x0200
)

// [DRAWITEMSTRUCT] CtlType.
//
// [DRAWITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type ODT uint32

const (
	ODT_MENU     ODT = 1
	ODT_LISTBOX  ODT = 2
	ODT_COMBOBOX ODT = 3
	ODT_BUTTON   ODT = 4
	ODT_STATIC   ODT = 5
	ODT_TAB      ODT = 101
	ODT_LISTVIEW ODT = 102
)

// [COMPAREITEMSTRUCT] and [DELETEITEMSTRUCT] CtlType. Originally with ODT
// prefix.
//
// [COMPAREITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-compareitemstruct
// [DELETEITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-deleteitemstruct
type ODT_C uint32

const (
	ODT_C_LISTBOX  ODT_C = ODT_C(ODT_LISTBOX)
	ODT_C_COMBOBOX ODT_C = ODT_C(ODT_COMBOBOX)
)

// [PeekMessage] wRemoveMsg.
//
// [PeekMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
type PM uint32

const (
	PM_NOREMOVE PM = 0x0000
	PM_REMOVE   PM = 0x0001
	PM_NOYIELD  PM = 0x0002

	PM_QS_INPUT       PM = PM(QS_INPUT << 16)
	PM_QS_PAINT       PM = PM(QS_PAINT << 16)
	PM_QS_POSTMESSAGE PM = PM((QS_POSTMESSAGE | QS_HOTKEY | QS_TIMER) << 16)
	PM_QS_SENDMESSAGE PM = PM(QS_SENDMESSAGE << 16)
)

// [GetQueueStatus] flags.
//
// [GetQueueStatus]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
type QS uint32

const (
	QS_KEY            QS = 0x0001
	QS_MOUSEMOVE      QS = 0x0002
	QS_MOUSEBUTTON    QS = 0x0004
	QS_POSTMESSAGE    QS = 0x0008
	QS_TIMER          QS = 0x0010
	QS_PAINT          QS = 0x0020
	QS_SENDMESSAGE    QS = 0x0040
	QS_HOTKEY         QS = 0x0080
	QS_ALLPOSTMESSAGE QS = 0x0100
	QS_RAWINPUT       QS = 0x0400
	QS_TOUCH          QS = 0x0800
	QS_POINTER        QS = 0x1000
	QS_MOUSE          QS = QS_MOUSEMOVE | QS_MOUSEBUTTON
	QS_INPUT          QS = QS_MOUSE | QS_KEY | QS_RAWINPUT | QS_TOUCH | QS_POINTER
	QS_ALLINPUT       QS = QS_INPUT | QS_POSTMESSAGE | QS_TIMER | QS_PAINT | QS_HOTKEY | QS_SENDMESSAGE
)

// [WM_HSCROLL], [WM_VSCROLL], [WM_HSCROLLCLIPBOARD] and [WM_VSCROLLCLIPBOARD]
// request. Originally with SB prefix.
//
// [WM_HSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-hscroll
// [WM_VSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-vscroll
// [WM_HSCROLLCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-hscrollclipboard
// [WM_VSCROLLCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-vscrollclipboard
type SB_REQ uint16

const (
	SB_REQ_LINEUP        SB_REQ = 0
	SB_REQ_LINELEFT      SB_REQ = 0
	SB_REQ_LINEDOWN      SB_REQ = 1
	SB_REQ_LINERIGHT     SB_REQ = 1
	SB_REQ_PAGEUP        SB_REQ = 2
	SB_REQ_PAGELEFT      SB_REQ = 2
	SB_REQ_PAGEDOWN      SB_REQ = 3
	SB_REQ_PAGERIGHT     SB_REQ = 3
	SB_REQ_THUMBPOSITION SB_REQ = 4
	SB_REQ_THUMBTRACK    SB_REQ = 5
	SB_REQ_TOP           SB_REQ = 6
	SB_REQ_LEFT          SB_REQ = 6
	SB_REQ_BOTTOM        SB_REQ = 7
	SB_REQ_RIGHT         SB_REQ = 7
	SB_REQ_ENDSCROLL     SB_REQ = 8
)

// [GetScrollInfo] nBar, among others. Originally has SB prefix.
//
// [GetScrollInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getscrollinfo
type SB_TYPE int32

const (
	SB_TYPE_HORZ SB_TYPE = 0
	SB_TYPE_VERT SB_TYPE = 1
	SB_TYPE_CTL  SB_TYPE = 2
	SB_TYPE_BOTH SB_TYPE = 3
)

// [WM_SYSCOMMAND] type of requested command.
//
// [WM_SYSCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
type SC uint32

const (
	SC_SIZE         SC = 0xf000
	SC_MOVE         SC = 0xf010
	SC_MINIMIZE     SC = 0xf020
	SC_MAXIMIZE     SC = 0xf030
	SC_NEXTWINDOW   SC = 0xf040
	SC_PREVWINDOW   SC = 0xf050
	SC_CLOSE        SC = 0xf060
	SC_VSCROLL      SC = 0xf070
	SC_HSCROLL      SC = 0xf080
	SC_MOUSEMENU    SC = 0xf090
	SC_KEYMENU      SC = 0xf100
	SC_ARRANGE      SC = 0xf110
	SC_RESTORE      SC = 0xf120
	SC_TASKLIST     SC = 0xf130
	SC_SCREENSAVE   SC = 0xf140
	SC_HOTKEY       SC = 0xf150
	SC_DEFAULT      SC = 0xf160
	SC_MONITORPOWER SC = 0xf170
	SC_CONTEXTHELP  SC = 0xf180
	SC_SEPARATOR    SC = 0xf00f
)

// [SECURITY_IMPERSONATION_LEVEL] enumeration.
//
// [SECURITY_IMPERSONATION_LEVEL]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ne-winnt-security_impersonation_level
type SECURITY_IMPERSONATION_LEVEL uint32

const (
	SECURITY_IMPERSONATION_LEVEL_ANONYMOUS SECURITY_IMPERSONATION_LEVEL = iota
	SECURITY_IMPERSONATION_LEVEL_IDENTIFICATION
	SECURITY_IMPERSONATION_LEVEL_IMPERSONATION
	SECURITY_IMPERSONATION_LEVEL_DELAGATION
)

// [SCROLLINFO] fMask.
//
// [SCROLLINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-scrollinfo
type SIF uint32

const (
	SIF_RANGE           SIF = 0x0001
	SIF_PAGE            SIF = 0x0002
	SIF_POS             SIF = 0x0004
	SIF_DISABLENOSCROLL SIF = 0x0008
	SIF_TRACKPOS        SIF = 0x0010
	SIF_ALL             SIF = SIF_RANGE | SIF_PAGE | SIF_POS | SIF_TRACKPOS
)

// [WM_SIZE] request.
//
// [WM_SIZE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-size
type SIZE_REQ int32

const (
	SIZE_REQ_RESTORED  SIZE_REQ = 0 // The window has been resized, but neither the SIZE_REQ_MINIMIZED nor SIZE_REQ_MAXIMIZED value applies.
	SIZE_REQ_MINIMIZED SIZE_REQ = 1 // The window has been minimized.
	SIZE_REQ_MAXIMIZED SIZE_REQ = 2 // The window has been maximized.
	SIZE_REQ_MAXSHOW   SIZE_REQ = 3 // Message is sent to all pop-up windows when some other window has been restored to its former size.
	SIZE_REQ_MAXHIDE   SIZE_REQ = 4 // Message is sent to all pop-up windows when some other window is maximized.
)

// [GetSystemMetrics] nIndex.
//
// [GetSystemMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
type SM int32

const (
	SM_CXSCREEN                    SM = 0
	SM_CYSCREEN                    SM = 1
	SM_CXVSCROLL                   SM = 2
	SM_CYHSCROLL                   SM = 3
	SM_CYCAPTION                   SM = 4
	SM_CXBORDER                    SM = 5
	SM_CYBORDER                    SM = 6
	SM_CXDLGFRAME                  SM = 7
	SM_CYDLGFRAME                  SM = 8
	SM_CYVTHUMB                    SM = 9
	SM_CXHTHUMB                    SM = 10
	SM_CXICON                      SM = 11
	SM_CYICON                      SM = 12
	SM_CXCURSOR                    SM = 13
	SM_CYCURSOR                    SM = 14
	SM_CYMENU                      SM = 15
	SM_CXFULLSCREEN                SM = 16
	SM_CYFULLSCREEN                SM = 17
	SM_CYKANJIWINDOW               SM = 18
	SM_MOUSEPRESENT                SM = 19
	SM_CYVSCROLL                   SM = 20
	SM_CXHSCROLL                   SM = 21
	SM_DEBUG                       SM = 22
	SM_SWAPBUTTON                  SM = 23
	SM_RESERVED1                   SM = 24
	SM_RESERVED2                   SM = 25
	SM_RESERVED3                   SM = 26
	SM_RESERVED4                   SM = 27
	SM_CXMIN                       SM = 28
	SM_CYMIN                       SM = 29
	SM_CXSIZE                      SM = 30
	SM_CYSIZE                      SM = 31
	SM_CXFRAME                     SM = 32
	SM_CYFRAME                     SM = 33
	SM_CXMINTRACK                  SM = 34
	SM_CYMINTRACK                  SM = 35
	SM_CXDOUBLECLK                 SM = 36
	SM_CYDOUBLECLK                 SM = 37
	SM_CXICONSPACING               SM = 38
	SM_CYICONSPACING               SM = 39
	SM_MENUDROPALIGNMENT           SM = 40
	SM_PENWINDOWS                  SM = 41
	SM_DBCSENABLED                 SM = 42
	SM_CMOUSEBUTTONS               SM = 43
	SM_CXFIXEDFRAME                SM = SM_CXDLGFRAME
	SM_CYFIXEDFRAME                SM = SM_CYDLGFRAME
	SM_CXSIZEFRAME                 SM = SM_CXFRAME
	SM_CYSIZEFRAME                 SM = SM_CYFRAME
	SM_SECURE                      SM = 44
	SM_CXEDGE                      SM = 45
	SM_CYEDGE                      SM = 46
	SM_CXMINSPACING                SM = 47
	SM_CYMINSPACING                SM = 48
	SM_CXSMICON                    SM = 49
	SM_CYSMICON                    SM = 50
	SM_CYSMCAPTION                 SM = 51
	SM_CXSMSIZE                    SM = 52
	SM_CYSMSIZE                    SM = 53
	SM_CXMENUSIZE                  SM = 54
	SM_CYMENUSIZE                  SM = 55
	SM_ARRANGE                     SM = 56
	SM_CXMINIMIZED                 SM = 57
	SM_CYMINIMIZED                 SM = 58
	SM_CXMAXTRACK                  SM = 59
	SM_CYMAXTRACK                  SM = 60
	SM_CXMAXIMIZED                 SM = 61
	SM_CYMAXIMIZED                 SM = 62
	SM_NETWORK                     SM = 63
	SM_CLEANBOOT                   SM = 67
	SM_CXDRAG                      SM = 68
	SM_CYDRAG                      SM = 69
	SM_SHOWSOUNDS                  SM = 70
	SM_CXMENUCHECK                 SM = 71
	SM_CYMENUCHECK                 SM = 72
	SM_SLOWMACHINE                 SM = 73
	SM_MIDEASTENABLED              SM = 74
	SM_MOUSEWHEELPRESENT           SM = 75
	SM_XVIRTUALSCREEN              SM = 76
	SM_YVIRTUALSCREEN              SM = 77
	SM_CXVIRTUALSCREEN             SM = 78
	SM_CYVIRTUALSCREEN             SM = 79
	SM_CMONITORS                   SM = 80
	SM_SAMEDISPLAYFORMAT           SM = 81
	SM_IMMENABLED                  SM = 82
	SM_CXFOCUSBORDER               SM = 83
	SM_CYFOCUSBORDER               SM = 84
	SM_TABLETPC                    SM = 86
	SM_MEDIACENTER                 SM = 87
	SM_STARTER                     SM = 88
	SM_SERVERR2                    SM = 89
	SM_MOUSEHORIZONTALWHEELPRESENT SM = 91
	SM_CXPADDEDBORDER              SM = 92
	SM_DIGITIZER                   SM = 94
	SM_MAXIMUMTOUCHES              SM = 95
	SM_CMETRICS                    SM = 97
	SM_REMOTESESSION               SM = 0x1000
	SM_SHUTTINGDOWN                SM = 0x2000
	SM_REMOTECONTROL               SM = 0x2001
	SM_CARETBLINKINGENABLED        SM = 0x2002
	SM_CONVERTIBLESLATEMODE        SM = 0x2003
	SM_SYSTEMDOCKED                SM = 0x2004
)

// [SendMessageTimeout] flags.
//
// [SendMessageTimeout]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagetimeoutw
type SMTO uint32

const (
	SMTO_ABORTIFHUNG        SMTO = 0x0002
	SMTO_BLOCK              SMTO = 0x0001
	SMTO_NORMAL             SMTO = 0x0000
	SMTO_NOTIMEOUTIFNOTHUNG SMTO = 0x0008
	SMTO_ERRORONEXIT        SMTO = 0x0020
)

// [SystemParametersInfo] uiAction.
//
// [SystemParametersInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
type SPI uint32

const (
	SPI_GETBEEP                     SPI = 0x0001
	SPI_SETBEEP                     SPI = 0x0002
	SPI_GETMOUSE                    SPI = 0x0003
	SPI_SETMOUSE                    SPI = 0x0004
	SPI_GETBORDER                   SPI = 0x0005
	SPI_SETBORDER                   SPI = 0x0006
	SPI_GETKEYBOARDSPEED            SPI = 0x000a
	SPI_SETKEYBOARDSPEED            SPI = 0x000b
	SPI_LANGDRIVER                  SPI = 0x000c
	SPI_ICONHORIZONTALSPACING       SPI = 0x000d
	SPI_GETSCREENSAVETIMEOUT        SPI = 0x000e
	SPI_SETSCREENSAVETIMEOUT        SPI = 0x000f
	SPI_GETSCREENSAVEACTIVE         SPI = 0x0010
	SPI_SETSCREENSAVEACTIVE         SPI = 0x0011
	SPI_GETGRIDGRANULARITY          SPI = 0x0012
	SPI_SETGRIDGRANULARITY          SPI = 0x0013
	SPI_SETDESKWALLPAPER            SPI = 0x0014
	SPI_SETDESKPATTERN              SPI = 0x0015
	SPI_GETKEYBOARDDELAY            SPI = 0x0016
	SPI_SETKEYBOARDDELAY            SPI = 0x0017
	SPI_ICONVERTICALSPACING         SPI = 0x0018
	SPI_GETICONTITLEWRAP            SPI = 0x0019
	SPI_SETICONTITLEWRAP            SPI = 0x001a
	SPI_GETMENUDROPALIGNMENT        SPI = 0x001b
	SPI_SETMENUDROPALIGNMENT        SPI = 0x001c
	SPI_SETDOUBLECLKWIDTH           SPI = 0x001d
	SPI_SETDOUBLECLKHEIGHT          SPI = 0x001e
	SPI_GETICONTITLELOGFONT         SPI = 0x001f
	SPI_SETDOUBLECLICKTIME          SPI = 0x0020
	SPI_SETMOUSEBUTTONSWAP          SPI = 0x0021
	SPI_SETICONTITLELOGFONT         SPI = 0x0022
	SPI_GETFASTTASKSWITCH           SPI = 0x0023
	SPI_SETFASTTASKSWITCH           SPI = 0x0024
	SPI_SETDRAGFULLWINDOWS          SPI = 0x0025
	SPI_GETDRAGFULLWINDOWS          SPI = 0x0026
	SPI_GETNONCLIENTMETRICS         SPI = 0x0029
	SPI_SETNONCLIENTMETRICS         SPI = 0x002a
	SPI_GETMINIMIZEDMETRICS         SPI = 0x002b
	SPI_SETMINIMIZEDMETRICS         SPI = 0x002c
	SPI_GETICONMETRICS              SPI = 0x002d
	SPI_SETICONMETRICS              SPI = 0x002e
	SPI_SETWORKAREA                 SPI = 0x002f
	SPI_GETWORKAREA                 SPI = 0x0030
	SPI_SETPENWINDOWS               SPI = 0x0031
	SPI_GETHIGHCONTRAST             SPI = 0x0042
	SPI_SETHIGHCONTRAST             SPI = 0x0043
	SPI_GETKEYBOARDPREF             SPI = 0x0044
	SPI_SETKEYBOARDPREF             SPI = 0x0045
	SPI_GETSCREENREADER             SPI = 0x0046
	SPI_SETSCREENREADER             SPI = 0x0047
	SPI_GETANIMATION                SPI = 0x0048
	SPI_SETANIMATION                SPI = 0x0049
	SPI_GETFONTSMOOTHING            SPI = 0x004a
	SPI_SETFONTSMOOTHING            SPI = 0x004b
	SPI_SETDRAGWIDTH                SPI = 0x004c
	SPI_SETDRAGHEIGHT               SPI = 0x004d
	SPI_SETHANDHELD                 SPI = 0x004e
	SPI_GETLOWPOWERTIMEOUT          SPI = 0x004f
	SPI_GETPOWEROFFTIMEOUT          SPI = 0x0050
	SPI_SETLOWPOWERTIMEOUT          SPI = 0x0051
	SPI_SETPOWEROFFTIMEOUT          SPI = 0x0052
	SPI_GETLOWPOWERACTIVE           SPI = 0x0053
	SPI_GETPOWEROFFACTIVE           SPI = 0x0054
	SPI_SETLOWPOWERACTIVE           SPI = 0x0055
	SPI_SETPOWEROFFACTIVE           SPI = 0x0056
	SPI_SETCURSORS                  SPI = 0x0057
	SPI_SETICONS                    SPI = 0x0058
	SPI_GETDEFAULTINPUTLANG         SPI = 0x0059
	SPI_SETDEFAULTINPUTLANG         SPI = 0x005a
	SPI_SETLANGTOGGLE               SPI = 0x005b
	SPI_GETWINDOWSEXTENSION         SPI = 0x005c
	SPI_SETMOUSETRAILS              SPI = 0x005d
	SPI_GETMOUSETRAILS              SPI = 0x005e
	SPI_SETSCREENSAVERRUNNING       SPI = 0x0061
	SPI_SCREENSAVERRUNNING          SPI = SPI_SETSCREENSAVERRUNNING
	SPI_GETFILTERKEYS               SPI = 0x0032
	SPI_SETFILTERKEYS               SPI = 0x0033
	SPI_GETTOGGLEKEYS               SPI = 0x0034
	SPI_SETTOGGLEKEYS               SPI = 0x0035
	SPI_GETMOUSEKEYS                SPI = 0x0036
	SPI_SETMOUSEKEYS                SPI = 0x0037
	SPI_GETSHOWSOUNDS               SPI = 0x0038
	SPI_SETSHOWSOUNDS               SPI = 0x0039
	SPI_GETSTICKYKEYS               SPI = 0x003a
	SPI_SETSTICKYKEYS               SPI = 0x003b
	SPI_GETACCESSTIMEOUT            SPI = 0x003c
	SPI_SETACCESSTIMEOUT            SPI = 0x003d
	SPI_GETSERIALKEYS               SPI = 0x003e
	SPI_SETSERIALKEYS               SPI = 0x003f
	SPI_GETSOUNDSENTRY              SPI = 0x0040
	SPI_SETSOUNDSENTRY              SPI = 0x0041
	SPI_GETSNAPTODEFBUTTON          SPI = 0x005f
	SPI_SETSNAPTODEFBUTTON          SPI = 0x0060
	SPI_GETMOUSEHOVERWIDTH          SPI = 0x0062
	SPI_SETMOUSEHOVERWIDTH          SPI = 0x0063
	SPI_GETMOUSEHOVERHEIGHT         SPI = 0x0064
	SPI_SETMOUSEHOVERHEIGHT         SPI = 0x0065
	SPI_GETMOUSEHOVERTIME           SPI = 0x0066
	SPI_SETMOUSEHOVERTIME           SPI = 0x0067
	SPI_GETWHEELSCROLLLINES         SPI = 0x0068
	SPI_SETWHEELSCROLLLINES         SPI = 0x0069
	SPI_GETMENUSHOWDELAY            SPI = 0x006a
	SPI_SETMENUSHOWDELAY            SPI = 0x006b
	SPI_GETWHEELSCROLLCHARS         SPI = 0x006c
	SPI_SETWHEELSCROLLCHARS         SPI = 0x006d
	SPI_GETSHOWIMEUI                SPI = 0x006e
	SPI_SETSHOWIMEUI                SPI = 0x006f
	SPI_GETMOUSESPEED               SPI = 0x0070
	SPI_SETMOUSESPEED               SPI = 0x0071
	SPI_GETSCREENSAVERRUNNING       SPI = 0x0072
	SPI_GETDESKWALLPAPER            SPI = 0x0073
	SPI_GETAUDIODESCRIPTION         SPI = 0x0074
	SPI_SETAUDIODESCRIPTION         SPI = 0x0075
	SPI_GETSCREENSAVESECURE         SPI = 0x0076
	SPI_SETSCREENSAVESECURE         SPI = 0x0077
	SPI_GETHUNGAPPTIMEOUT           SPI = 0x0078
	SPI_SETHUNGAPPTIMEOUT           SPI = 0x0079
	SPI_GETWAITTOKILLTIMEOUT        SPI = 0x007a
	SPI_SETWAITTOKILLTIMEOUT        SPI = 0x007b
	SPI_GETWAITTOKILLSERVICETIMEOUT SPI = 0x007c
	SPI_SETWAITTOKILLSERVICETIMEOUT SPI = 0x007d
	SPI_GETMOUSEDOCKTHRESHOLD       SPI = 0x007e
	SPI_SETMOUSEDOCKTHRESHOLD       SPI = 0x007f
	SPI_GETPENDOCKTHRESHOLD         SPI = 0x0080
	SPI_SETPENDOCKTHRESHOLD         SPI = 0x0081
	SPI_GETWINARRANGING             SPI = 0x0082
	SPI_SETWINARRANGING             SPI = 0x0083
	SPI_GETMOUSEDRAGOUTTHRESHOLD    SPI = 0x0084
	SPI_SETMOUSEDRAGOUTTHRESHOLD    SPI = 0x0085
	SPI_GETPENDRAGOUTTHRESHOLD      SPI = 0x0086
	SPI_SETPENDRAGOUTTHRESHOLD      SPI = 0x0087
	SPI_GETMOUSESIDEMOVETHRESHOLD   SPI = 0x0088
	SPI_SETMOUSESIDEMOVETHRESHOLD   SPI = 0x0089
	SPI_GETPENSIDEMOVETHRESHOLD     SPI = 0x008a
	SPI_SETPENSIDEMOVETHRESHOLD     SPI = 0x008b
	SPI_GETDRAGFROMMAXIMIZE         SPI = 0x008c
	SPI_SETDRAGFROMMAXIMIZE         SPI = 0x008d
	SPI_GETSNAPSIZING               SPI = 0x008e
	SPI_SETSNAPSIZING               SPI = 0x008f
	SPI_GETDOCKMOVING               SPI = 0x0090
	SPI_SETDOCKMOVING               SPI = 0x0091
)

// [SystemParametersInfo] fWinIni.
//
// [SystemParametersInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
type SPIF uint32

const (
	SPIF_UPDATEINIFILE    SPIF = 1
	SPIF_SENDWININICHANGE SPIF = 2
	SPIF_SENDCHANGE       SPIF = SPIF_SENDWININICHANGE
)

// Static control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/static-control-styles
type SS WS

const (
	SS_LEFT            SS = 0x0000_0000
	SS_CENTER          SS = 0x0000_0001
	SS_RIGHT           SS = 0x0000_0002
	SS_ICON            SS = 0x0000_0003
	SS_BLACKRECT       SS = 0x0000_0004
	SS_GRAYRECT        SS = 0x0000_0005
	SS_WHITERECT       SS = 0x0000_0006
	SS_BLACKFRAME      SS = 0x0000_0007
	SS_GRAYFRAME       SS = 0x0000_0008
	SS_WHITEFRAME      SS = 0x0000_0009
	SS_USERITEM        SS = 0x0000_000a
	SS_SIMPLE          SS = 0x0000_000b
	SS_LEFTNOWORDWRAP  SS = 0x0000_000c
	SS_OWNERDRAW       SS = 0x0000_000d
	SS_BITMAP          SS = 0x0000_000e
	SS_ENHMETAFILE     SS = 0x0000_000f
	SS_ETCHEDHORZ      SS = 0x0000_0010
	SS_ETCHEDVERT      SS = 0x0000_0011
	SS_ETCHEDFRAME     SS = 0x0000_0012
	SS_TYPEMASK        SS = 0x0000_001f
	SS_REALSIZECONTROL SS = 0x0000_0040
	SS_NOPREFIX        SS = 0x0000_0080
	SS_NOTIFY          SS = 0x0000_0100
	SS_CENTERIMAGE     SS = 0x0000_0200
	SS_RIGHTJUST       SS = 0x0000_0400
	SS_REALSIZEIMAGE   SS = 0x0000_0800
	SS_SUNKEN          SS = 0x0000_1000
	SS_EDITCONTROL     SS = 0x0000_2000
	SS_ENDELLIPSIS     SS = 0x0000_4000
	SS_PATHELLIPSIS    SS = 0x0000_8000
	SS_WORDELLIPSIS    SS = 0x0000_c000
	SS_ELLIPSISMASK    SS = 0x0000_c000
)

// [ShowWindow] nCmdShow.
//
// [ShowWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
type SW int32

const (
	SW_HIDE            SW = 0
	SW_SHOWNORMAL      SW = 1
	SW_SHOWMINIMIZED   SW = 2
	SW_SHOWMAXIMIZED   SW = 3
	SW_MAXIMIZE        SW = 3
	SW_SHOWNOACTIVATE  SW = 4
	SW_SHOW            SW = 5
	SW_MINIMIZE        SW = 6
	SW_SHOWMINNOACTIVE SW = 7
	SW_SHOWNA          SW = 8
	SW_RESTORE         SW = 9
	SW_SHOWDEFAULT     SW = 10
	SW_FORCEMINIMIZE   SW = 11
)

// [SetWindowPos], [DeferWindowPos] uFlags.
//
// [SetWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
// [DeferWindowPos]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
type SWP uint32

const (
	SWP_NOSIZE         SWP = 0x0001
	SWP_NOMOVE         SWP = 0x0002
	SWP_NOZORDER       SWP = 0x0004
	SWP_NOREDRAW       SWP = 0x0008
	SWP_NOACTIVATE     SWP = 0x0010
	SWP_FRAMECHANGED   SWP = 0x0020
	SWP_SHOWWINDOW     SWP = 0x0040
	SWP_HIDEWINDOW     SWP = 0x0080
	SWP_NOCOPYBITS     SWP = 0x0100
	SWP_NOOWNERZORDER  SWP = 0x0200
	SWP_NOSENDCHANGING SWP = 0x0400
	SWP_DRAWFRAME      SWP = SWP_FRAMECHANGED
	SWP_NOREPOSITION   SWP = SWP_NOOWNERZORDER
	SWP_DEFERERASE     SWP = 0x2000
	SWP_ASYNCWINDOWPOS SWP = 0x4000
)

// [WM_SHOWWINDOW] return value. Originally has SW prefix.
//
// [WM_SHOWWINDOW]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-showwindow
type SWS uint8

const (
	SWS_OTHERUNZOOM   SWS = 4 // The window is being uncovered because a maximize window was restored or minimized.
	SWS_OTHERZOOM     SWS = 2 // The window is being covered by another window that has been maximized.
	SWS_PARENTCLOSING SWS = 1 // The window's owner window is being minimized.
	SWS_PARENTOPENING SWS = 3 // The window's owner window is being restored.
)

// Trackbar's [WM_HSCROLL] and [WM_VSCROLL] request. Originally has TB prefix.
//
// [WM_HSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-hscroll--trackbar-
// [WM_VSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-vscroll--trackbar-
type TB_REQ uint16

const (
	TB_REQ_LINEUP        TB_REQ = 0
	TB_REQ_LINEDOWN      TB_REQ = 1
	TB_REQ_PAGEUP        TB_REQ = 2
	TB_REQ_PAGEDOWN      TB_REQ = 3
	TB_REQ_THUMBPOSITION TB_REQ = 4
	TB_REQ_THUMBTRACK    TB_REQ = 5
	TB_REQ_TOP           TB_REQ = 6
	TB_REQ_BOTTOM        TB_REQ = 7
	TB_REQ_ENDTRACK      TB_REQ = 8
)

// [SetUserObjectInformation] nIndex.
//
// [SetUserObjectInformation]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setuserobjectinformationw
type UOI int32

const (
	UOI_FLAGS                           UOI = 1
	UOI_TIMERPROC_EXCEPTION_SUPPRESSION UOI = 7
)

// [Virtual key codes].
//
// [Virtual key codes]: https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
type VK uint16

const (
	VK_LBUTTON             VK = 0x01
	VK_RBUTTON             VK = 0x02
	VK_CANCEL              VK = 0x03
	VK_MBUTTON             VK = 0x04
	VK_XBUTTON1            VK = 0x05
	VK_XBUTTON2            VK = 0x06
	VK_BACK                VK = 0x08
	VK_TAB                 VK = 0x09
	VK_CLEAR               VK = 0x0c
	VK_RETURN              VK = 0x0d
	VK_SHIFT               VK = 0x10
	VK_CONTROL             VK = 0x11
	VK_MENU                VK = 0x12
	VK_PAUSE               VK = 0x13
	VK_CAPITAL             VK = 0x14
	VK_KANA                VK = 0x15
	VK_HANGEUL             VK = 0x15
	VK_HANGUL              VK = 0x15
	VK_JUNJA               VK = 0x17
	VK_FINAL               VK = 0x18
	VK_HANJA               VK = 0x19
	VK_KANJI               VK = 0x19
	VK_ESCAPE              VK = 0x1b
	VK_CONVERT             VK = 0x1c
	VK_NONCONVERT          VK = 0x1d
	VK_ACCEPT              VK = 0x1e
	VK_MODECHANGE          VK = 0x1f
	VK_SPACE               VK = 0x20
	VK_PRIOR               VK = 0x21
	VK_NEXT                VK = 0x22
	VK_END                 VK = 0x23
	VK_HOME                VK = 0x24
	VK_LEFT                VK = 0x25
	VK_UP                  VK = 0x26
	VK_RIGHT               VK = 0x27
	VK_DOWN                VK = 0x28
	VK_SELECT              VK = 0x29
	VK_PRINT               VK = 0x2a
	VK_EXECUTE             VK = 0x2b
	VK_SNAPSHOT            VK = 0x2c
	VK_INSERT              VK = 0x2d
	VK_DELETE              VK = 0x2e
	VK_HELP                VK = 0x2f
	VK_LWIN                VK = 0x5b
	VK_RWIN                VK = 0x5c
	VK_APPS                VK = 0x5d
	VK_SLEEP               VK = 0x5f
	VK_NUMPAD0             VK = 0x60
	VK_NUMPAD1             VK = 0x61
	VK_NUMPAD2             VK = 0x62
	VK_NUMPAD3             VK = 0x63
	VK_NUMPAD4             VK = 0x64
	VK_NUMPAD5             VK = 0x65
	VK_NUMPAD6             VK = 0x66
	VK_NUMPAD7             VK = 0x67
	VK_NUMPAD8             VK = 0x68
	VK_NUMPAD9             VK = 0x69
	VK_MULTIPLY            VK = 0x6a
	VK_ADD                 VK = 0x6b
	VK_SEPARATOR           VK = 0x6c
	VK_SUBTRACT            VK = 0x6d
	VK_DECIMAL             VK = 0x6e
	VK_DIVIDE              VK = 0x6f
	VK_F1                  VK = 0x70
	VK_F2                  VK = 0x71
	VK_F3                  VK = 0x72
	VK_F4                  VK = 0x73
	VK_F5                  VK = 0x74
	VK_F6                  VK = 0x75
	VK_F7                  VK = 0x76
	VK_F8                  VK = 0x77
	VK_F9                  VK = 0x78
	VK_F10                 VK = 0x79
	VK_F11                 VK = 0x7a
	VK_F12                 VK = 0x7b
	VK_F13                 VK = 0x7c
	VK_F14                 VK = 0x7d
	VK_F15                 VK = 0x7e
	VK_F16                 VK = 0x7f
	VK_F17                 VK = 0x80
	VK_F18                 VK = 0x81
	VK_F19                 VK = 0x82
	VK_F20                 VK = 0x83
	VK_F21                 VK = 0x84
	VK_F22                 VK = 0x85
	VK_F23                 VK = 0x86
	VK_F24                 VK = 0x87
	VK_NUMLOCK             VK = 0x90
	VK_SCROLL              VK = 0x91
	VK_OEM_NEC_EQUAL       VK = 0x92
	VK_OEM_FJ_JISHO        VK = 0x92
	VK_OEM_FJ_MASSHOU      VK = 0x93
	VK_OEM_FJ_TOUROKU      VK = 0x94
	VK_OEM_FJ_LOYA         VK = 0x95
	VK_OEM_FJ_ROYA         VK = 0x96
	VK_LSHIFT              VK = 0xa0
	VK_RSHIFT              VK = 0xa1
	VK_LCONTROL            VK = 0xa2
	VK_RCONTROL            VK = 0xa3
	VK_LMENU               VK = 0xa4
	VK_RMENU               VK = 0xa5
	VK_BROWSER_BACK        VK = 0xa6
	VK_BROWSER_FORWARD     VK = 0xa7
	VK_BROWSER_REFRESH     VK = 0xa8
	VK_BROWSER_STOP        VK = 0xa9
	VK_BROWSER_SEARCH      VK = 0xaa
	VK_BROWSER_FAVORITES   VK = 0xab
	VK_BROWSER_HOME        VK = 0xac
	VK_VOLUME_MUTE         VK = 0xad
	VK_VOLUME_DOWN         VK = 0xae
	VK_VOLUME_UP           VK = 0xaf
	VK_MEDIA_NEXT_TRACK    VK = 0xb0
	VK_MEDIA_PREV_TRACK    VK = 0xb1
	VK_MEDIA_STOP          VK = 0xb2
	VK_MEDIA_PLAY_PAUSE    VK = 0xb3
	VK_LAUNCH_MAIL         VK = 0xb4
	VK_LAUNCH_MEDIA_SELECT VK = 0xb5
	VK_LAUNCH_APP1         VK = 0xb6
	VK_LAUNCH_APP2         VK = 0xb7
	VK_OEM_1               VK = 0xba
	VK_OEM_PLUS            VK = 0xbb
	VK_OEM_COMMA           VK = 0xbc
	VK_OEM_MINUS           VK = 0xbd
	VK_OEM_PERIOD          VK = 0xbe
	VK_OEM_2               VK = 0xbf
	VK_OEM_3               VK = 0xc0
	VK_OEM_4               VK = 0xdb
	VK_OEM_5               VK = 0xdc
	VK_OEM_6               VK = 0xdd
	VK_OEM_7               VK = 0xde
	VK_OEM_8               VK = 0xdf
	VK_OEM_AX              VK = 0xe1
	VK_OEM_102             VK = 0xe2
	VK_ICO_HELP            VK = 0xe3
	VK_ICO_00              VK = 0xe4
	VK_PROCESSKEY          VK = 0xe5
	VK_ICO_CLEAR           VK = 0xe6
	VK_PACKET              VK = 0xe7
	VK_OEM_RESET           VK = 0xe9
	VK_OEM_JUMP            VK = 0xea
	VK_OEM_PA1             VK = 0xeb
	VK_OEM_PA2             VK = 0xec
	VK_OEM_PA3             VK = 0xed
	VK_OEM_WSCTRL          VK = 0xee
	VK_OEM_CUSEL           VK = 0xef
	VK_OEM_ATTN            VK = 0xf0
	VK_OEM_FINISH          VK = 0xf1
	VK_OEM_COPY            VK = 0xf2
	VK_OEM_AUTO            VK = 0xf3
	VK_OEM_ENLW            VK = 0xf4
	VK_OEM_BACKTAB         VK = 0xf5
	VK_ATTN                VK = 0xf6
	VK_CRSEL               VK = 0xf7
	VK_EXSEL               VK = 0xf8
	VK_EREOF               VK = 0xf9
	VK_PLAY                VK = 0xfa
	VK_ZOOM                VK = 0xfb
	VK_NONAME              VK = 0xfc
	VK_PA1                 VK = 0xfd
	VK_OEM_CLEAR           VK = 0xfe
)

// [WM_ACTIVATE] activation state.
//
// [WM_ACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-activate
type WA int32

const (
	WA_INACTIVE    WA = 0
	WA_ACTIVE      WA = 1
	WA_CLICKACTIVE WA = 2
)

// [SetWindowDisplayAffinity] dwAffinity
//
// [SetWindowDisplayAffinity]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowdisplayaffinity
type WDA uint32

const (
	WDA_NONE               WDA = 0x0000_0000
	WDA_MONITOR            WDA = 0x0000_0001
	WDA_EXCLUDEFROMCAPTURE WDA = 0x0000_0011
)

// [SetWindowsHookEx] idHook.
//
// [SetWindowsHookEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
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

// [WM_SIZING] window edge.
//
// [WM_SIZING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-sizing
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

// Window [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
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

// Extended window [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/winmsg/extended-window-styles
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

// [WM_NCCALCSIZE] return flags.
//
// [WM_NCCALCSIZE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
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

// Composes XTYP in [PFNCALLBACK].
//
// [PFNCALLBACK]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nc-ddeml-pfncallback
type XCLASS uint32

const (
	XCLASS_MASK         XCLASS = 0xfc00
	XCLASS_BOOL         XCLASS = 0x1000
	XCLASS_DATA         XCLASS = 0x2000
	XCLASS_FLAGS        XCLASS = 0x4000
	XCLASS_NOTIFICATION XCLASS = 0x8000
)

// [PFNCALLBACK] wType.
//
// [PFNCALLBACK]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nc-ddeml-pfncallback
type XTYP uint32

const (
	XTYP_ERROR           XTYP = 0x0000 | XTYP(XCLASS_NOTIFICATION) | XTYP(XTYPF_NOBLOCK)
	XTYP_ADVDATA         XTYP = 0x0010 | XTYP(XCLASS_FLAGS)
	XTYP_ADVREQ          XTYP = 0x0020 | XTYP(XCLASS_DATA) | XTYP(XTYPF_NOBLOCK)
	XTYP_ADVSTART        XTYP = 0x0030 | XTYP(XCLASS_BOOL)
	XTYP_ADVSTOP         XTYP = 0x0040 | XTYP(XCLASS_NOTIFICATION)
	XTYP_EXECUTE         XTYP = 0x0050 | XTYP(XCLASS_FLAGS)
	XTYP_CONNECT         XTYP = 0x0060 | XTYP(XCLASS_BOOL) | XTYP(XTYPF_NOBLOCK)
	XTYP_CONNECT_CONFIRM XTYP = 0x0070 | XTYP(XCLASS_NOTIFICATION) | XTYP(XTYPF_NOBLOCK)
	XTYP_XACT_COMPLETE   XTYP = 0x0080 | XTYP(XCLASS_NOTIFICATION)
	XTYP_POKE            XTYP = 0x0090 | XTYP(XCLASS_FLAGS)
	XTYP_REGISTER        XTYP = 0x00a0 | XTYP(XCLASS_NOTIFICATION) | XTYP(XTYPF_NOBLOCK)
	XTYP_REQUEST         XTYP = 0x00b0 | XTYP(XCLASS_DATA)
	XTYP_DISCONNECT      XTYP = 0x00c0 | XTYP(XCLASS_NOTIFICATION) | XTYP(XTYPF_NOBLOCK)
	XTYP_UNREGISTER      XTYP = 0x00d0 | XTYP(XCLASS_NOTIFICATION) | XTYP(XTYPF_NOBLOCK)
	XTYP_WILDCONNECT     XTYP = 0x00e0 | XTYP(XCLASS_DATA) | XTYP(XTYPF_NOBLOCK)

	XTYP_MASK  XTYP = 0x00f0
	XTYP_SHIFT XTYP = 4
)

// Composes XTYP in [PFNCALLBACK].
//
// [PFNCALLBACK]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nc-ddeml-pfncallback
type XTYPF uint32

const (
	XTYPF_NOBLOCK XTYPF = 0x0002
	XTYPF_NODATA  XTYPF = 0x0004
	XTYPF_ACKREQ  XTYPF = 0x0008
)
