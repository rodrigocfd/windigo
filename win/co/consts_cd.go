package co

// ComboBox control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-combobox-control-reference-messages
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

// ComboBox control notifications, sent via WM_COMMAND.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-combobox-control-reference-notifications
const (
	CBN_ERRSPACE     CMD = -1
	CBN_SELCHANGE    CMD = 1
	CBN_DBLCLK       CMD = 2
	CBN_SETFOCUS     CMD = 3
	CBN_KILLFOCUS    CMD = 4
	CBN_EDITCHANGE   CMD = 5
	CBN_EDITUPDATE   CMD = 6
	CBN_DROPDOWN     CMD = 7
	CBN_CLOSEUP      CMD = 8
	CBN_SELENDOK     CMD = 9
	CBN_SELENDCANCEL CMD = 10
)

// ComboBox styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/combo-box-styles
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

// Common controls messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-general-control-reference-messages
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

// NMCUSTOMDRAW dwDrawStage.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type CDDS uint32

const (
	CDDS_PREPAINT      CDDS = 0x00000001
	CDDS_POSTPAINT     CDDS = 0x00000002
	CDDS_PREERASE      CDDS = 0x00000003
	CDDS_POSTERASE     CDDS = 0x00000004
	CDDS_ITEM          CDDS = 0x00010000
	CDDS_ITEMPREPAINT  CDDS = CDDS_ITEM | CDDS_PREPAINT
	CDDS_ITEMPOSTPAINT CDDS = CDDS_ITEM | CDDS_POSTPAINT
	CDDS_ITEMPREERASE  CDDS = CDDS_ITEM | CDDS_PREERASE
	CDDS_ITEMPOSTERASE CDDS = CDDS_ITEM | CDDS_POSTERASE
	CDDS_SUBITEM       CDDS = 0x00020000
)

// NMCUSTOMDRAW uItemState.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type CDIS uint32

const (
	CDIS_SELECTED         CDIS = 0x0001
	CDIS_GRAYED           CDIS = 0x0002
	CDIS_DISABLED         CDIS = 0x0004
	CDIS_CHECKED          CDIS = 0x0008
	CDIS_FOCUS            CDIS = 0x0010
	CDIS_DEFAULT          CDIS = 0x0020
	CDIS_HOT              CDIS = 0x0040
	CDIS_MARKED           CDIS = 0x0080
	CDIS_INDETERMINATE    CDIS = 0x0100
	CDIS_SHOWKEYBOARDCUES CDIS = 0x0200
	CDIS_NEARHOT          CDIS = 0x0400
	CDIS_OTHERSIDEHOT     CDIS = 0x0800
	CDIS_DROPHILITED      CDIS = 0x1000
)

// NM_CUSTOMDRAW return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type CDRF uint32

const (
	CDRF_DODEFAULT         CDRF = 0x00000000
	CDRF_NEWFONT           CDRF = 0x00000002
	CDRF_SKIPDEFAULT       CDRF = 0x00000004
	CDRF_DOERASE           CDRF = 0x00000008
	CDRF_SKIPPOSTPAINT     CDRF = 0x00000100
	CDRF_NOTIFYPOSTPAINT   CDRF = 0x00000010
	CDRF_NOTIFYITEMDRAW    CDRF = 0x00000020
	CDRF_NOTIFYSUBITEMDRAW CDRF = 0x00000020
	CDRF_NOTIFYPOSTERASE   CDRF = 0x00000040
)

// Clipboard formats.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/standard-clipboard-formats
type CF uint16

const (
	CF_TEXT            CF = 1
	CF_BITMAP          CF = 2
	CF_METAFILEPICT    CF = 3
	CF_SYLK            CF = 4
	CF_DIF             CF = 5
	CF_TIFF            CF = 6
	CF_OEMTEXT         CF = 7
	CF_DIB             CF = 8
	CF_PALETTE         CF = 9
	CF_PENDATA         CF = 10
	CF_RIFF            CF = 11
	CF_WAVE            CF = 12
	CF_UNICODETEXT     CF = 13
	CF_ENHMETAFILE     CF = 14
	CF_HDROP           CF = 15
	CF_LOCALE          CF = 16
	CF_DIBV5           CF = 17
	CF_MAX             CF = 18
	CF_OWNERDISPLAY    CF = 0x0080
	CF_DSPTEXT         CF = 0x0081
	CF_DSPBITMAP       CF = 0x0082
	CF_DSPMETAFILEPICT CF = 0x0083
	CF_DSPENHMETAFILE  CF = 0x008e
	CF_PRIVATEFIRST    CF = 0x0200
	CF_PRIVATELAST     CF = 0x02ff
	CF_GDIOBJFIRST     CF = 0x0300
	CF_GDIOBJLAST      CF = 0x03ff
)

// TEXTMETRIC tmCharSet. Originally with _CHARSET suffix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type CHARSET uint8

const (
	CHARSET_ANSI        CHARSET = 0
	CHARSET_DEFAULT     CHARSET = 1
	CHARSET_SYMBOL      CHARSET = 2
	CHARSET_SHIFTJIS    CHARSET = 128
	CHARSET_HANGUL      CHARSET = 129
	CHARSET_GB2312      CHARSET = 134
	CHARSET_CHINESEBIG5 CHARSET = 136
	CHARSET_OEM         CHARSET = 255
	CHARSET_JOHAB       CHARSET = 130
	CHARSET_HEBREW      CHARSET = 177
	CHARSET_ARABIC      CHARSET = 178
	CHARSET_GREEK       CHARSET = 161
	CHARSET_TURKISH     CHARSET = 162
	CHARSET_VIETNAMESE  CHARSET = 163
	CHARSET_THAI        CHARSET = 222
	CHARSET_EASTEUROPE  CHARSET = 238
	CHARSET_RUSSIAN     CHARSET = 204
	CHARSET_MAC         CHARSET = 77
	CHARSET_BALTIC      CHARSET = 186
)

// CoCreateInstance() dwClsContext.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
type CLSCTX uint32

const (
	CLSCTX_INPROC_SERVER          CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER         CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER           CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16        CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER          CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16       CLSCTX = 0x20
	CLSCTX_NO_CODE_DOWNLOAD       CLSCTX = 0x400
	CLSCTX_NO_CUSTOM_MARSHAL      CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG         CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA            CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA             CLSCTX = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT   CLSCTX = 0x20000
	CLSCTX_ACTIVATE_X86_SERVER    CLSCTX = 0x40000
	CLSCTX_ACTIVATE_32_BIT_SERVER CLSCTX = CLSCTX_ACTIVATE_X86_SERVER
	CLSCTX_ACTIVATE_64_BIT_SERVER CLSCTX = 0x80000
	CLSCTX_ENABLE_CLOAKING        CLSCTX = 0x100000
	CLSCTX_APPCONTAINER           CLSCTX = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU     CLSCTX = 0x800000
	CLSCTX_ACTIVATE_ARM32_SERVER  CLSCTX = 0x2000000
	CLSCTX_PS_DLL                 CLSCTX = 0x80000000
	CLSCTX_ALL                    CLSCTX = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
	CLSCTX_SERVER                 CLSCTX = CLSCTX_INPROC_SERVER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

// WM_COMMAND notification codes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
type CMD int32

const (
	CMD_MENU        CMD = 0 // Message originated from a menu.
	CMD_ACCELERATOR CMD = 1 // Message originated from an accelerator.
)

// CoInitializeEx() dwCoInit.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
type COINIT uint32

const (
	COINIT_APARTMENTTHREADED COINIT = 0x2
	COINIT_MULTITHREADED     COINIT = 0x0
	COINIT_DISABLE_OLE1DDE   COINIT = 0x4
	COINIT_SPEED_OVER_MEMORY COINIT = 0x8
)

// GetSysColor() nIndex.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
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

// CreateFile() dwCreationDisposition. Originally with CREATE prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type CREATION_DISP uint32

const (
	CREATION_DISP_CREATE_ALWAYS     CREATION_DISP = 2
	CREATION_DISP_CREATE_NEW        CREATION_DISP = 1
	CREATION_DISP_OPEN_ALWAYS       CREATION_DISP = 4
	CREATION_DISP_OPEN_EXISTING     CREATION_DISP = 3
	CREATION_DISP_TRUNCATE_EXISTING CREATION_DISP = 5
)

// Window class styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/window-class-styles
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
	CS_DROPSHADOW      CS = 0x00020000
)

// WM_GETDLGCODE return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
type DLGC uint32

const (
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

// SetProcessDpiAwarenessContext() value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
type DPI_AWARE_CTX int32

const (
	DPI_AWARE_CTX_UNAWARE           DPI_AWARE_CTX = -1
	DPI_AWARE_CTX_SYSTEM_AWARE      DPI_AWARE_CTX = -2
	DPI_AWARE_CTX_PER_MON_AWARE     DPI_AWARE_CTX = -3
	DPI_AWARE_CTX_PER_MON_AWARE_V2  DPI_AWARE_CTX = -4
	DPI_AWARE_CTX_UNAWARE_GDISCALED DPI_AWARE_CTX = -5
)

// DateTimePicker control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-date-and-time-picker-control-reference-messages
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

// DateTimePicker control notifications, sent via WM_NOTIFY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-date-and-time-picker-control-reference-notifications
const (
	_DTN_FIRST  NM = -740
	_DTN_FIRST2 NM = -753

	DTN_CLOSEUP        NM = _DTN_FIRST2 - 0
	DTN_DROPDOWN       NM = _DTN_FIRST2 - 1
	DTN_DATETIMECHANGE NM = _DTN_FIRST2 - 6
	DTN_FORMATQUERY    NM = _DTN_FIRST - 2
	DTN_FORMAT         NM = _DTN_FIRST - 3
	DTN_WMKEYDOWN      NM = _DTN_FIRST - 4
	DTN_USERSTRING     NM = _DTN_FIRST - 5
)

// DateTimePicker control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-control-styles
type DTS WS

const (
	DTS_NONE                   DTS = 0
	DTS_UPDOWN                 DTS = 0x0001
	DTS_SHOWNONE               DTS = 0x0002
	DTS_SHORTDATEFORMAT        DTS = 0x0000
	DTS_LONGDATEFORMAT         DTS = 0x0004
	DTS_SHORTDATECENTURYFORMAT DTS = 0x000c
	DTS_TIMEFORMAT             DTS = 0x0009
	DTS_APPCANPARSE            DTS = 0x0010
	DTS_RIGHTALIGN             DTS = 0x0020
)
