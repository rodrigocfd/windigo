package co

// MessageBox() uType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
type MB uint32

const (
	MB_ABORTRETRYIGNORE  MB = 0x00000002
	MB_CANCELTRYCONTINUE MB = 0x00000006
	MB_HELP              MB = 0x00004000
	MB_OK                MB = 0x00000000
	MB_OKCANCEL          MB = 0x00000001
	MB_RETRYCANCEL       MB = 0x00000005
	MB_YESNO             MB = 0x00000004
	MB_YESNOCANCEL       MB = 0x00000003

	MB_ICONEXCLAMATION MB = 0x00000030
	MB_ICONWARNING     MB = 0x00000030
	MB_ICONINFORMATION MB = 0x00000040
	MB_ICONASTERISK    MB = 0x00000040
	MB_ICONQUESTION    MB = 0x00000020
	MB_ICONSTOP        MB = 0x00000010
	MB_ICONERROR       MB = 0x00000010
	MB_ICONHAND        MB = 0x00000010

	MB_DEFBUTTON1 MB = 0x00000000
	MB_DEFBUTTON2 MB = 0x00000100
	MB_DEFBUTTON3 MB = 0x00000200
	MB_DEFBUTTON4 MB = 0x00000300

	MB_APPLMODAL   MB = 0x00000000
	MB_SYSTEMMODAL MB = 0x00001000
	MB_TASKMODAL   MB = 0x00002000

	MB_DEFAULT_DESKTOP_ONLY MB = 0x00020000
	MB_RIGHT                MB = 0x00080000
	MB_RTLREADING           MB = 0x00100000
	MB_SETFOREGROUND        MB = 0x00010000
	MB_TOPMOST              MB = 0x00040000
	MB_SERVICE_NOTIFICATION MB = 0x00200000
)

// NMVIEWCHANGE dwOldView/dwNewView.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmviewchange
type MCMV uint32

const (
	MCMV_MONTH   MCMV = 0
	MCMV_YEAR    MCMV = 1
	MCMV_DECADE  MCMV = 2
	MCMV_CENTURY MCMV = 3
)

// MonthCalendar control notifications, sent via WM_NOTIFY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-month-calendar-control-reference-notifications
const (
	_MCN_FIRST NM = -746

	MCN_SELECT      NM = _MCN_FIRST
	MCN_GETDAYSTATE NM = _MCN_FIRST - 1
	MCN_SELCHANGE   NM = _MCN_FIRST - 3
	MCN_VIEWCHANGE  NM = _MCN_FIRST - 4
)

// MonthCalendar control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/month-calendar-control-styles
type MCS WS

const (
	MCS_NONE             MCS = 0
	MCS_DAYSTATE         MCS = 0x0001 // The month calendar sends MCN_GETDAYSTATE notifications to request information about which days should be displayed in bold.
	MCS_MULTISELECT      MCS = 0x0002 // The month calendar enables the user to select a range of dates within the control. By default, the maximum range is one week. You can change the maximum range that can be selected by using the MCM_SETMAXSELCOUNT message.
	MCS_WEEKNUMBERS      MCS = 0x0004 // The month calendar control displays week numbers (1-52) to the left of each row of days. Week 1 is defined as the first week that contains at least four days.
	MCS_NOTODAYCIRCLE    MCS = 0x0008 // The month calendar control does not circle the "today" date.
	MCS_NOTODAY          MCS = 0x0010 // The month calendar control does not display the "today" date at the bottom of the control.
	MCS_NOTRAILINGDATES  MCS = 0x0040 // Dates from the previous and next months are not displayed in the current month's calendar.
	MCS_SHORTDAYSOFWEEK  MCS = 0x0080 // Short day names are displayed in the header.
	MCS_NOSELCHANGEONNAV MCS = 0x0100 // The selection is not changed when the user navigates next or previous in the calendar. This allows the user to select a range larger than is visible.
)

// WM_MENUCHAR menu type. Originally with MF prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
type MFMC uint16

const (
	POPUP   MFMC = 0x00000010
	SYSMENU MFMC = 0x00002000
)

// CheckMenuItem() uCheck, among others.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuitem
type MF uint32

const (
	MF_INSERT          MF = 0x00000000
	MF_CHANGE          MF = 0x00000080
	MF_APPEND          MF = 0x00000100
	MF_DELETE          MF = 0x00000200
	MF_REMOVE          MF = 0x00001000
	MF_BYCOMMAND       MF = 0x00000000
	MF_BYPOSITION      MF = 0x00000400
	MF_SEPARATOR       MF = 0x00000800
	MF_ENABLED         MF = 0x00000000
	MF_GRAYED          MF = 0x00000001
	MF_DISABLED        MF = 0x00000002
	MF_UNCHECKED       MF = 0x00000000
	MF_CHECKED         MF = 0x00000008
	MF_USECHECKBITMAPS MF = 0x00000200
	MF_STRING          MF = 0x00000000
	MF_BITMAP          MF = 0x00000004
	MF_OWNERDRAW       MF = 0x00000100
	MF_POPUP           MF = 0x00000010
	MF_MENUBARBREAK    MF = 0x00000020
	MF_MENUBREAK       MF = 0x00000040
	MF_UNHILITE        MF = 0x00000000
	MF_HILITE          MF = 0x00000080
	MF_DEFAULT         MF = 0x00001000
	MF_SYSMENU         MF = 0x00002000
	MF_HELP            MF = 0x00004000
	MF_RIGHTJUSTIFY    MF = 0x00004000
	MF_MOUSESELECT     MF = 0x00008000
)

// MENUITEMINFO fState.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MFS uint32

const (
	MFS_GRAYED    MFS = 0x00000003
	MFS_DISABLED  MFS = MFS_GRAYED
	MFS_CHECKED   MFS = MFS(MF_CHECKED)
	MFS_HILITE    MFS = MFS(MF_HILITE)
	MFS_ENABLED   MFS = MFS(MF_ENABLED)
	MFS_UNCHECKED MFS = MFS(MF_UNCHECKED)
	MFS_UNHILITE  MFS = MFS(MF_UNHILITE)
	MFS_DEFAULT   MFS = MFS(MF_DEFAULT)
)

// MENUITEMINFO fType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MFT uint32

const (
	MFT_STRING       MFT = MFT(MF_STRING)
	MFT_BITMAP       MFT = MFT(MF_BITMAP)
	MFT_MENUBARBREAK MFT = MFT(MF_MENUBARBREAK)
	MFT_MENUBREAK    MFT = MFT(MF_MENUBREAK)
	MFT_OWNERDRAW    MFT = MFT(MF_OWNERDRAW)
	MFT_RADIOCHECK   MFT = 0x00000200
	MFT_SEPARATOR    MFT = MFT(MF_SEPARATOR)
	MFT_RIGHTORDER   MFT = 0x00002000
	MFT_RIGHTJUSTIFY MFT = MFT(MF_RIGHTJUSTIFY)
)

// MENUITEMINFO fMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MIIM uint32

const (
	MIIM_STATE      MIIM = 0x00000001
	MIIM_ID         MIIM = 0x00000002
	MIIM_SUBMENU    MIIM = 0x00000004
	MIIM_CHECKMARKS MIIM = 0x00000008
	MIIM_TYPE       MIIM = 0x00000010
	MIIM_DATA       MIIM = 0x00000020
	MIIM_STRING     MIIM = 0x00000040
	MIIM_BITMAP     MIIM = 0x00000080
	MIIM_FTYPE      MIIM = 0x00000100
)

// MENUINFO fMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
type MIM uint32

const (
	MIM_MAXHEIGHT       MIM = 0x00000001
	MIM_BACKGROUND      MIM = 0x00000002
	MIM_HELPID          MIM = 0x00000004
	MIM_MENUDATA        MIM = 0x00000008
	MIM_STYLE           MIM = 0x00000010
	MIM_APPLYTOSUBMENUS MIM = 0x80000000
)

// WM_LBUTTONDOWN virtual keys, among others
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
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

// WM_MENUCHAR return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
type MNC uint32

const (
	MNC_IGNORE  MNC = 0
	MNC_CLOSE   MNC = 1
	MNC_EXECUTE MNC = 2
	MNC_SELECT  MNC = 3
)

// WM_MENUDRAG return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
type MND uint32

const (
	MND_CONTINUE MND = 0
	MND_ENDMENU  MND = 1
)

// WM_MENUGETOBJECT return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
type MNGO uint32

const (
	MNGO_NOINTERFACE MNGO = 0x00000000
	MNGO_NOERROR     MNGO = 0x00000001
)

// MENUGETOBJECTINFO dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menugetobjectinfo
type MNGOF uint32

const (
	MNGOF_TOPGAP    MNGOF = 0x00000001
	MNGOF_BOTTOMGAP MNGOF = 0x00000002
)

// MENUINFO dwStyle.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
type MNS uint32

const (
	MNS_NOCHECK     MNS = 0x80000000
	MNS_MODELESS    MNS = 0x40000000
	MNS_DRAGDROP    MNS = 0x20000000
	MNS_AUTODISMISS MNS = 0x10000000
	MNS_NOTIFYBYPOS MNS = 0x08000000
	MNS_CHECKORBMP  MNS = 0x04000000
)

// WM_HOTKEY combined keys.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
type MOD uint16

const (
	MOD_ALT     MOD = 0x0001
	MOD_CONTROL MOD = 0x0002
	MOD_SHIFT   MOD = 0x0004
	MOD_WIN     MOD = 0x0008
)

// MonitorFromPoint() dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
type MONITOR uint32

const (
	MONITOR_DEFAULTTONULL    MONITOR = 0x00000000
	MONITOR_DEFAULTTOPRIMARY MONITOR = 0x00000001
	MONITOR_DEFAULTTONEAREST MONITOR = 0x00000002
)

// WM_ENTERIDLE displayed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
type MSGF uint32

const (
	MSGF_DIALOGBOX MSGF = 0
	MSGF_MENU      MSGF = 2
)

// NOTIFYICONDATA uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NIF uint32

const (
	NIF_MESSAGE  NIF = 0x00000001
	NIF_ICON     NIF = 0x00000002
	NIF_TIP      NIF = 0x00000004
	NIF_STATE    NIF = 0x00000008
	NIF_INFO     NIF = 0x00000010
	NIF_GUID     NIF = 0x00000020
	NIF_REALTIME NIF = 0x00000040
	NIF_SHOWTIP  NIF = 0x00000080
)

// NOTIFYICONDATA dwInfoFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NIIF uint32

const (
	NIIF_NONE               NIIF = 0x00000000
	NIIF_INFO               NIIF = 0x00000001
	NIIF_WARNING            NIIF = 0x00000002
	NIIF_ERROR              NIIF = 0x00000003
	NIIF_USER               NIIF = 0x00000004
	NIIF_NOSOUND            NIIF = 0x00000010
	NIIF_LARGE_ICON         NIIF = 0x00000020
	NIIF_RESPECT_QUIET_TIME NIIF = 0x00000080
)

// Shell_NotifyIcon dwMessage.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
type NIM uint32

const (
	NIM_ADD        NIM = 0x00000000
	NIM_MODIFY     NIM = 0x00000001
	NIM_DELETE     NIM = 0x00000002
	NIM_SETFOCUS   NIM = 0x00000003
	NIM_SETVERSION NIM = 0x00000004
)

// NOTIFYICONDATA dwState and dwStateMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NIS uint32

const (
	NIS_HIDDEN     NIS = 0x00000001
	NIS_SHAREDICON NIS = 0x00000002
)

// Common control notifications.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/common-control-reference#notifications
type NM int32

const (
	_NM_FIRST NM = 0

	NM_OUTOFMEMORY          NM = _NM_FIRST - 1
	NM_CLICK                NM = _NM_FIRST - 2
	NM_DBLCLK               NM = _NM_FIRST - 3
	NM_RETURN               NM = _NM_FIRST - 4
	NM_RCLICK               NM = _NM_FIRST - 5
	NM_RDBLCLK              NM = _NM_FIRST - 6
	NM_SETFOCUS             NM = _NM_FIRST - 7
	NM_KILLFOCUS            NM = _NM_FIRST - 8
	NM_CUSTOMDRAW           NM = _NM_FIRST - 12
	NM_HOVER                NM = _NM_FIRST - 13
	NM_NCHITTEST            NM = _NM_FIRST - 14
	NM_KEYDOWN              NM = _NM_FIRST - 15
	NM_RELEASEDCAPTURE      NM = _NM_FIRST - 16
	NM_SETCURSOR            NM = _NM_FIRST - 17
	NM_CHAR                 NM = _NM_FIRST - 18
	NM_TOOLTIPSCREATED      NM = _NM_FIRST - 19
	NM_LDOWN                NM = _NM_FIRST - 20
	NM_RDOWN                NM = _NM_FIRST - 21
	NM_THEMECHANGED         NM = _NM_FIRST - 22
	NM_FONTCHANGED          NM = _NM_FIRST - 23
	NM_CUSTOMTEXT           NM = _NM_FIRST - 24
	NM_TVSTATEIMAGECHANGING NM = _NM_FIRST - 24
)
