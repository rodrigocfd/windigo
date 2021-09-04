package co

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
	SBARS_SIZEGRIP SBARS = 0x0100 // The status bar control will include a sizing grip at the right end of the status bar. A sizing grip is similar to a sizing border; it is a rectangular area that the user can click and drag to resize the parent window.
	SBARS_TOOLTIPS SBARS = 0x0800 // Use this style to enable tooltips.
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

// STARTUPINFO dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfow
type STARTF uint32

const (
	STARTF_FORCEONFEEDBACK  STARTF = 0x00000040
	STARTF_FORCEOFFFEEDBACK STARTF = 0x00000080
	STARTF_PREVENTPINNING   STARTF = 0x00002000
	STARTF_RUNFULLSCREEN    STARTF = 0x00000020
	STARTF_TITLEISAPPID     STARTF = 0x00001000
	STARTF_TITLEISLINKNAME  STARTF = 0x00000800
	STARTF_UNTRUSTEDSOURCE  STARTF = 0x00008000
	STARTF_USECOUNTCHARS    STARTF = 0x00000008
	STARTF_USEFILLATTRIBUTE STARTF = 0x00000010
	STARTF_USEHOTKEY        STARTF = 0x00000200
	STARTF_USEPOSITION      STARTF = 0x00000004
	STARTF_USESHOWWINDOW    STARTF = 0x00000001
	STARTF_USESIZE          STARTF = 0x00000002
	STARTF_USESTDHANDLES    STARTF = 0x00000100
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
