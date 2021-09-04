package co

// CHOOSECOLOR Flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-choosecolorw-r1
type CC uint32

const (
	CC_RGBINIT              CC = 0x00000001
	CC_FULLOPEN             CC = 0x00000002
	CC_PREVENTFULLOPEN      CC = 0x00000004
	CC_SHOWHELP             CC = 0x00000008
	CC_ENABLEHOOK           CC = 0x00000010
	CC_ENABLETEMPLATE       CC = 0x00000020
	CC_ENABLETEMPLATEHANDLE CC = 0x00000040
	CC_SOLIDCOLOR           CC = 0x00000080
	CC_ANYCOLOR             CC = 0x00000100
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

// CreateProcess() dwCreationFlags.
type CREATE uint32

const (
	CREATE_NONE CREATE = 0

	CREATE_BREAKAWAY_FROM_JOB        CREATE = 0x01000000
	CREATE_DEFAULT_ERROR_MODE        CREATE = 0x04000000
	CREATE_NEW_CONSOLE               CREATE = 0x00000010
	CREATE_NEW_PROCESS_GROUP         CREATE = 0x00000200
	CREATE_NO_WINDOW                 CREATE = 0x08000000
	CREATE_PROTECTED_PROCESS         CREATE = 0x00040000
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL CREATE = 0x02000000
	CREATE_SECURE_PROCESS            CREATE = 0x00400000
	CREATE_SEPARATE_WOW_VDM          CREATE = 0x00000800
	CREATE_SHARED_WOW_VDM            CREATE = 0x00001000
	CREATE_SUSPENDED                 CREATE = 0x00000004
	CREATE_UNICODE_ENVIRONMENT       CREATE = 0x00000400

	CREATE_DEBUG_ONLY_THIS_PROCESS      CREATE = 0x00000002
	CREATE_DEBUG_PROCESS                CREATE = 0x00000001
	CREATE_DETACHED_PROCESS             CREATE = 0x00000008
	CREATE_EXTENDED_STARTUPINFO_PRESENT CREATE = 0x00080000
	CREATE_INHERIT_PARENT_AFFINITY      CREATE = 0x00010000
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

// CreateFile() dwCreationDisposition. Originally without prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type DISPOSITION uint32

const (
	DISPOSITION_CREATE_ALWAYS     DISPOSITION = 2
	DISPOSITION_CREATE_NEW        DISPOSITION = 1
	DISPOSITION_OPEN_ALWAYS       DISPOSITION = 4
	DISPOSITION_OPEN_EXISTING     DISPOSITION = 3
	DISPOSITION_TRUNCATE_EXISTING DISPOSITION = 5
)

// WM_GETDLGCODE return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
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

// DwmSetIconicLivePreviewBitmap() dwSITFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
type DWM_SIT uint32

const (
	DWM_SIT_NONE         DWM_SIT = 0
	DWM_SIT_DISPLAYFRAME DWM_SIT = 0x00000001
)
