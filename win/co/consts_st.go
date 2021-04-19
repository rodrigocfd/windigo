package co

// Status bar control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-status-bars-reference-messages
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

// WM_HSCROLL, WM_HSCROLL, WM_HSCROLLCLIPBOARD and WM_VSCROLLCLIPBOARD request.
// Originally with SB prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll
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

// GetScrollInfo() nBar, among others. Originally has SB prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getscrollinfo
type SB_TYPE int32

const (
	SB_TYPE_HORZ SB_TYPE = 0
	SB_TYPE_VERT SB_TYPE = 1
	SB_TYPE_CTL  SB_TYPE = 2
)

// StatusBar styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/status-bar-styles
type SBARS WS

const (
	SBARS_SIZEGRIP SBARS = 0x0100
	SBARS_TOOLTIPS SBARS = 0x0800
)

// StatusBar control notifications, sent via WM_NOTIFY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-status-bars-reference-notifications
const (
	_SBN_FIRST NM = -880

	SBN_SIMPLEMODECHANGE NM = _SBN_FIRST - 0
)

// WM_SYSCOMMAND type of requested command.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
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

// CreateFileMapping() flProtect.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type SEC uint32

const (
	SEC_NONE                   SEC = 0
	SEC_PARTITION_OWNER_HANDLE SEC = 0x00040000
	SEC_64K_PAGES              SEC = 0x00080000
	SEC_FILE                   SEC = 0x00800000
	SEC_IMAGE                  SEC = 0x01000000
	SEC_PROTECTED_IMAGE        SEC = 0x02000000
	SEC_RESERVE                SEC = 0x04000000
	SEC_COMMIT                 SEC = 0x08000000
	SEC_NOCACHE                SEC = 0x10000000
	SEC_WRITECOMBINE           SEC = 0x40000000
	SEC_LARGE_PAGES            SEC = 0x80000000
	SEC_IMAGE_NO_EXECUTE       SEC = SEC_IMAGE | SEC_NOCACHE
)

// CreateFile() dwFlagsAndAttributes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type SECURITY uint32

const (
	SECURITY_NONE             SECURITY = 0
	SECURITY_ANONYMOUS        SECURITY = 0 << 16
	SECURITY_IDENTIFICATION   SECURITY = 1 << 16
	SECURITY_IMPERSONATION    SECURITY = 2 << 16
	SECURITY_DELEGATION       SECURITY = 3 << 16
	SECURITY_CONTEXT_TRACKING SECURITY = 0x00040000
	SECURITY_EFFECTIVE_ONLY   SECURITY = 0x00080000
)

// SHFILEINFO dwAttributes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SFGAO uint32

const (
	_DROPEFFECT_NONE   SFGAO = 0
	_DROPEFFECT_COPY   SFGAO = 1
	_DROPEFFECT_MOVE   SFGAO = 2
	_DROPEFFECT_LINK   SFGAO = 4
	_DROPEFFECT_SCROLL SFGAO = 0x80000000

	SFGAO_CANCOPY         SFGAO = _DROPEFFECT_COPY
	SFGAO_CANMOVE         SFGAO = _DROPEFFECT_MOVE
	SFGAO_CANLINK         SFGAO = _DROPEFFECT_LINK
	SFGAO_STORAGE         SFGAO = 0x00000008
	SFGAO_CANRENAME       SFGAO = 0x00000010
	SFGAO_CANDELETE       SFGAO = 0x00000020
	SFGAO_HASPROPSHEET    SFGAO = 0x00000040
	SFGAO_DROPTARGET      SFGAO = 0x00000100
	SFGAO_CAPABILITYMASK  SFGAO = 0x00000177
	SFGAO_PLACEHOLDER     SFGAO = 0x00000800
	SFGAO_SYSTEM          SFGAO = 0x00001000
	SFGAO_ENCRYPTED       SFGAO = 0x00002000
	SFGAO_ISSLOW          SFGAO = 0x00004000
	SFGAO_GHOSTED         SFGAO = 0x00008000
	SFGAO_LINK            SFGAO = 0x00010000
	SFGAO_SHARE           SFGAO = 0x00020000
	SFGAO_READONLY        SFGAO = 0x00040000
	SFGAO_HIDDEN          SFGAO = 0x00080000
	SFGAO_DISPLAYATTRMASK SFGAO = 0x000fc000
	SFGAO_FILESYSANCESTOR SFGAO = 0x10000000
	SFGAO_FOLDER          SFGAO = 0x20000000
	SFGAO_FILESYSTEM      SFGAO = 0x40000000
	SFGAO_HASSUBFOLDER    SFGAO = 0x80000000
	SFGAO_CONTENTSMASK    SFGAO = 0x80000000
	SFGAO_VALIDATE        SFGAO = 0x01000000
	SFGAO_REMOVABLE       SFGAO = 0x02000000
	SFGAO_COMPRESSED      SFGAO = 0x04000000
	SFGAO_BROWSABLE       SFGAO = 0x08000000
	SFGAO_NONENUMERATED   SFGAO = 0x00100000
	SFGAO_NEWCONTENT      SFGAO = 0x00200000
	SFGAO_CANMONIKER      SFGAO = 0x00400000
	SFGAO_HASSTORAGE      SFGAO = 0x00400000
	SFGAO_STREAM          SFGAO = 0x00400000
	SFGAO_STORAGEANCESTOR SFGAO = 0x00800000
	SFGAO_STORAGECAPMASK  SFGAO = 0x70c50008
	SFGAO_PKEYSFGAOMASK   SFGAO = 0x81044000
)

// SHGetFileInfo() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
type SHGFI uint32

const (
	SHGFI_NONE              SHGFI = 0
	SHGFI_ICON              SHGFI = 0x000000100
	SHGFI_DISPLAYNAME       SHGFI = 0x000000200
	SHGFI_TYPENAME          SHGFI = 0x000000400
	SHGFI_ATTRIBUTES        SHGFI = 0x000000800
	SHGFI_ICONLOCATION      SHGFI = 0x000001000
	SHGFI_EXETYPE           SHGFI = 0x000002000
	SHGFI_SYSICONINDEX      SHGFI = 0x000004000
	SHGFI_LINKOVERLAY       SHGFI = 0x000008000
	SHGFI_SELECTED          SHGFI = 0x000010000
	SHGFI_ATTR_SPECIFIED    SHGFI = 0x000020000
	SHGFI_LARGEICON         SHGFI = 0x000000000
	SHGFI_SMALLICON         SHGFI = 0x000000001
	SHGFI_OPENICON          SHGFI = 0x000000002
	SHGFI_SHELLICONSIZE     SHGFI = 0x000000004
	SHGFI_PIDL              SHGFI = 0x000000008
	SHGFI_USEFILEATTRIBUTES SHGFI = 0x000000010
	SHGFI_ADDOVERLAYS       SHGFI = 0x000000020
	SHGFI_OVERLAYINDEX      SHGFI = 0x000000040
)

// SCROLLINFO fMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-scrollinfo
type SIF uint32

const (
	SIF_RANGE           SIF = 0x0001
	SIF_PAGE            SIF = 0x0002
	SIF_POS             SIF = 0x0004
	SIF_DISABLENOSCROLL SIF = 0x0008
	SIF_TRACKPOS        SIF = 0x0010
	SIF_ALL             SIF = SIF_RANGE | SIF_PAGE | SIF_POS | SIF_TRACKPOS
)

// WM_SIZE request.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-size
type SIZE_REQ int32

const (
	SIZE_REQ_RESTORED  SIZE_REQ = 0 // The window has been resized, but neither the SIZE_REQ_MINIMIZED nor SIZE_REQ_MAXIMIZED value applies.
	SIZE_REQ_MINIMIZED SIZE_REQ = 1 // The window has been minimized.
	SIZE_REQ_MAXIMIZED SIZE_REQ = 2 // The window has been maximized.
	SIZE_REQ_MAXSHOW   SIZE_REQ = 3 // Message is sent to all pop-up windows when some other window has been restored to its former size.
	SIZE_REQ_MAXHIDE   SIZE_REQ = 4 // Message is sent to all pop-up windows when some other window is maximized.
)

// GetSystemMetrics() nIndex.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
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

// SystemParametersInfo() uiAction.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
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

// SystemParametersInfo() fWinIni.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
type SPIF uint32

const (
	SPIF_UPDATEINIFILE    SPIF = 1
	SPIF_SENDWININICHANGE SPIF = 2
	SPIF_SENDCHANGE       SPIF = SPIF_SENDWININICHANGE
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/secauthz/standard-access-rights
type STANDARD_RIGHTS uint32

const (
	STANDARD_RIGHTS_REQUIRED STANDARD_RIGHTS = 0x000f0000
	STANDARD_RIGHTS_READ     STANDARD_RIGHTS = STANDARD_RIGHTS(ACCESS_RIGHTS_READ_CONTROL)
	STANDARD_RIGHTS_WRITE    STANDARD_RIGHTS = STANDARD_RIGHTS(ACCESS_RIGHTS_READ_CONTROL)
	STANDARD_RIGHTS_EXECUTE  STANDARD_RIGHTS = STANDARD_RIGHTS(ACCESS_RIGHTS_READ_CONTROL)
	STANDARD_RIGHTS_ALL      STANDARD_RIGHTS = 0x001f0000
)

// Static control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/static-control-styles
type SS WS

const (
	SS_LEFT            SS = 0x00000000
	SS_CENTER          SS = 0x00000001
	SS_RIGHT           SS = 0x00000002
	SS_ICON            SS = 0x00000003
	SS_BLACKRECT       SS = 0x00000004
	SS_GRAYRECT        SS = 0x00000005
	SS_WHITERECT       SS = 0x00000006
	SS_BLACKFRAME      SS = 0x00000007
	SS_GRAYFRAME       SS = 0x00000008
	SS_WHITEFRAME      SS = 0x00000009
	SS_USERITEM        SS = 0x0000000a
	SS_SIMPLE          SS = 0x0000000b
	SS_LEFTNOWORDWRAP  SS = 0x0000000c
	SS_OWNERDRAW       SS = 0x0000000d
	SS_BITMAP          SS = 0x0000000e
	SS_ENHMETAFILE     SS = 0x0000000f
	SS_ETCHEDHORZ      SS = 0x00000010
	SS_ETCHEDVERT      SS = 0x00000011
	SS_ETCHEDFRAME     SS = 0x00000012
	SS_TYPEMASK        SS = 0x0000001f
	SS_REALSIZECONTROL SS = 0x00000040
	SS_NOPREFIX        SS = 0x00000080
	SS_NOTIFY          SS = 0x00000100
	SS_CENTERIMAGE     SS = 0x00000200
	SS_RIGHTJUST       SS = 0x00000400
	SS_REALSIZEIMAGE   SS = 0x00000800
	SS_SUNKEN          SS = 0x00001000
	SS_EDITCONTROL     SS = 0x00002000
	SS_ENDELLIPSIS     SS = 0x00004000
	SS_PATHELLIPSIS    SS = 0x00008000
	SS_WORDELLIPSIS    SS = 0x0000c000
	SS_ELLIPSISMASK    SS = 0x0000c000
)

// Static control notifications, sent via WM_COMMAND.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-static-control-reference-notifications
const (
	STN_CLICKED CMD = 0
	STN_DBLCLK  CMD = 1
	STN_ENABLE  CMD = 2
	STN_DISABLE CMD = 3
)

// ShowWindow() nCmdShow.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
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

// SetWindowPos(), DeferWindowPos() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
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

// SetTextAlign() align. Includes values with VTA prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
type TA uint32

const (
	TA_NOUPDATECP TA = 0
	TA_UPDATECP   TA = 1
	TA_LEFT       TA = 0
	TA_RIGHT      TA = 2
	TA_CENTER     TA = 6
	TA_TOP        TA = 0
	TA_BOTTOM     TA = 8
	TA_BASELINE   TA = 24
	TA_RTLREADING TA = 256
)

// Trackbar's WM_HSCROLL and WM_VSCROLL request.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll--trackbar-
type TB uint16

const (
	TB_LINEUP        TB = 0
	TB_LINEDOWN      TB = 1
	TB_PAGEUP        TB = 2
	TB_PAGEDOWN      TB = 3
	TB_THUMBPOSITION TB = 4
	TB_THUMBTRACK    TB = 5
	TB_TOP           TB = 6
	TB_BOTTOM        TB = 7
	TB_ENDTRACK      TB = 8
)

// Trackbar control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-messages
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

// Trackbar control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/trackbar-control-styles
type TBS WS

const (
	TBS_AUTOTICKS        TBS = 0x1
	TBS_VERT             TBS = 0x2
	TBS_HORZ             TBS = 0x0
	TBS_TOP              TBS = 0x4
	TBS_BOTTOM           TBS = 0x0
	TBS_LEFT             TBS = 0x4
	TBS_RIGHT            TBS = 0x0
	TBS_BOTH             TBS = 0x8
	TBS_NOTICKS          TBS = 0x10
	TBS_ENABLESELRANGE   TBS = 0x20
	TBS_FIXEDLENGTH      TBS = 0x40
	TBS_NOTHUMB          TBS = 0x80
	TBS_TOOLTIPS         TBS = 0x100
	TBS_REVERSED         TBS = 0x200
	TBS_DOWNISLEFT       TBS = 0x400
	TBS_NOTIFYBEFOREMOVE TBS = 0x800
	TBS_TRANSPARENTBKGND TBS = 0x1000
)

// GetTimeZoneInformation() return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformation
type TIME_ZONE_ID uint32

const (
	TIME_ZONE_ID_UNKNOWN  TIME_ZONE_ID = 0
	TIME_ZONE_ID_STANDARD TIME_ZONE_ID = 1
	TIME_ZONE_ID_DAYLIGHT TIME_ZONE_ID = 2
)

// Trackbar control notifications, sent via WM_NOTIFY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-notifications
const (
	_TRBN_FIRST NM = -1501

	TRBN_THUMBPOSCHANGING NM = _TRBN_FIRST - 1
)

// TrackPopupMenu() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
type TPM uint32

const (
	TPM_LEFTBUTTON      TPM = 0x0000
	TPM_RIGHTBUTTON     TPM = 0x0002
	TPM_LEFTALIGN       TPM = 0x0000
	TPM_CENTERALIGN     TPM = 0x0004
	TPM_RIGHTALIGN      TPM = 0x0008
	TPM_TOPALIGN        TPM = 0x0000
	TPM_VCENTERALIGN    TPM = 0x0010
	TPM_BOTTOMALIGN     TPM = 0x0020
	TPM_HORIZONTAL      TPM = 0x0000
	TPM_VERTICAL        TPM = 0x0040
	TPM_NONOTIFY        TPM = 0x0080
	TPM_RETURNCMD       TPM = 0x0100
	TPM_RECURSE         TPM = 0x0001
	TPM_HORPOSANIMATION TPM = 0x0400
	TPM_HORNEGANIMATION TPM = 0x0800
	TPM_VERPOSANIMATION TPM = 0x1000
	TPM_VERNEGANIMATION TPM = 0x2000
	TPM_NOANIMATION     TPM = 0x4000
	TPM_LAYOUTRTL       TPM = 0x8000
	TPM_WORKAREA        TPM = 0x10000
)

// TVM_EXPAND action flag.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvm-expand
type TVE uint32

const (
	TVE_COLLAPSE      TVE = 0x0001
	TVE_EXPAND        TVE = 0x0002
	TVE_TOGGLE        TVE = 0x0003
	TVE_EXPANDPARTIAL TVE = 0x4000
	TVE_COLLAPSERESET TVE = 0x8000
)

// TVM_GETNEXTITEM item to retrieve.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
type TVGN uint32

const (
	TVGN_ROOT            TVGN = 0x0000
	TVGN_NEXT            TVGN = 0x0001
	TVGN_PREVIOUS        TVGN = 0x0002
	TVGN_PARENT          TVGN = 0x0003
	TVGN_CHILD           TVGN = 0x0004
	TVGN_FIRSTVISIBLE    TVGN = 0x0005
	TVGN_NEXTVISIBLE     TVGN = 0x0006
	TVGN_PREVIOUSVISIBLE TVGN = 0x0007
	TVGN_DROPHILITE      TVGN = 0x0008
	TVGN_CARET           TVGN = 0x0009
	TVGN_LASTVISIBLE     TVGN = 0x000a
	TVGN_NEXTSELECTED    TVGN = 0x000b
)

// TVITEMTEX cChildren.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVI_CHILDREN int32

const (
	TVI_CHILDREN_ZERO     TVI_CHILDREN = 0
	TVI_CHILDREN_ONE      TVI_CHILDREN = 1
	TVI_CHILDREN_CALLBACK TVI_CHILDREN = -1
	TVI_CHILDREN_AUTO     TVI_CHILDREN = -2
)

// TVITEMTEX mask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVIF uint32

const (
	TVIF_TEXT          TVIF = 0x0001
	TVIF_IMAGE         TVIF = 0x0002
	TVIF_PARAM         TVIF = 0x0004
	TVIF_STATE         TVIF = 0x0008
	TVIF_HANDLE        TVIF = 0x0010
	TVIF_SELECTEDIMAGE TVIF = 0x0020
	TVIF_CHILDREN      TVIF = 0x0040
	TVIF_INTEGRAL      TVIF = 0x0080
	TVIF_STATEEX       TVIF = 0x0100
	TVIF_EXPANDEDIMAGE TVIF = 0x0200
)

// TVITEMTEX state.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVIS uint32

const (
	TVIS_SELECTED       TVIS = 0x0002
	TVIS_CUT            TVIS = 0x0004
	TVIS_DROPHILITED    TVIS = 0x0008
	TVIS_BOLD           TVIS = 0x0010
	TVIS_EXPANDED       TVIS = 0x0020
	TVIS_EXPANDEDONCE   TVIS = 0x0040
	TVIS_EXPANDPARTIAL  TVIS = 0x0080
	TVIS_OVERLAYMASK    TVIS = 0x0f00
	TVIS_STATEIMAGEMASK TVIS = 0xf000
	TVIS_USERMASK       TVIS = 0xf000
)

// TVITEMTEX uStateEx.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVIS_EX uint32

const (
	TVIS_EX_FLAT     TVIS_EX = 0x0001
	TVIS_EX_DISABLED TVIS_EX = 0x0002
	TVIS_EX_ALL      TVIS_EX = 0x0002
)

// TreeView control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-tree-view-control-reference-messages
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

// TreeView control notifications, sent via WM_NOTIFY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-tree-view-control-reference-notifications
const (
	_TVN_FIRST NM = -400

	TVN_SELCHANGING    NM = _TVN_FIRST - 50
	TVN_SELCHANGED     NM = _TVN_FIRST - 51
	TVN_GETDISPINFO    NM = _TVN_FIRST - 52
	TVN_SETDISPINFO    NM = _TVN_FIRST - 53
	TVN_ITEMEXPANDING  NM = _TVN_FIRST - 54
	TVN_ITEMEXPANDED   NM = _TVN_FIRST - 55
	TVN_BEGINDRAG      NM = _TVN_FIRST - 56
	TVN_BEGINRDRAG     NM = _TVN_FIRST - 57
	TVN_DELETEITEM     NM = _TVN_FIRST - 58
	TVN_BEGINLABELEDIT NM = _TVN_FIRST - 59
	TVN_ENDLABELEDIT   NM = _TVN_FIRST - 60
	TVN_KEYDOWN        NM = _TVN_FIRST - 12
	TVN_GETINFOTIP     NM = _TVN_FIRST - 14
	TVN_SINGLEEXPAND   NM = _TVN_FIRST - 15
	TVN_ITEMCHANGING   NM = _TVN_FIRST - 17
	TVN_ITEMCHANGED    NM = _TVN_FIRST - 19
	TVN_ASYNCDRAW      NM = _TVN_FIRST - 20
)

// TVN_SINGLEEXPAND return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tvn-singleexpand
type TVNRET uintptr

const (
	TVNRET_DEFAULT TVNRET = 0
	TVNRET_SKIPOLD TVNRET = 1
	TVNRET_SKIPNEW TVNRET = 2
)

// TreeView control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-control-window-styles
type TVS WS

const (
	TVS_HASBUTTONS      TVS = 0x0001
	TVS_HASLINES        TVS = 0x0002
	TVS_LINESATROOT     TVS = 0x0004
	TVS_EDITLABELS      TVS = 0x0008
	TVS_DISABLEDRAGDROP TVS = 0x0010
	TVS_SHOWSELALWAYS   TVS = 0x0020
	TVS_RTLREADING      TVS = 0x0040
	TVS_NOTOOLTIPS      TVS = 0x0080
	TVS_CHECKBOXES      TVS = 0x0100
	TVS_TRACKSELECT     TVS = 0x0200
	TVS_SINGLEEXPAND    TVS = 0x0400
	TVS_INFOTIP         TVS = 0x0800
	TVS_FULLROWSELECT   TVS = 0x1000
	TVS_NOSCROLL        TVS = 0x2000
	TVS_NONEVENHEIGHT   TVS = 0x4000
	TVS_NOHSCROLL       TVS = 0x8000
)

// TreeView control extended styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-control-window-extended-styles
type TVS_EX WS_EX

const (
	TVS_EX_NONE                TVS_EX = 0
	TVS_EX_NOSINGLECOLLAPSE    TVS_EX = 0x0001
	TVS_EX_MULTISELECT         TVS_EX = 0x0002
	TVS_EX_DOUBLEBUFFER        TVS_EX = 0x0004
	TVS_EX_NOINDENTSTATE       TVS_EX = 0x0008
	TVS_EX_RICHTOOLTIP         TVS_EX = 0x0010
	TVS_EX_AUTOHSCROLL         TVS_EX = 0x0020
	TVS_EX_FADEINOUTEXPANDOS   TVS_EX = 0x0040
	TVS_EX_PARTIALCHECKBOXES   TVS_EX = 0x0080
	TVS_EX_EXCLUSIONCHECKBOXES TVS_EX = 0x0100
	TVS_EX_DIMMEDCHECKBOXES    TVS_EX = 0x0200
	TVS_EX_DRAWIMAGEASYNC      TVS_EX = 0x0400
)
