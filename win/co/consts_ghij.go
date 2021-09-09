package co

// GetAncestor() gaFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getancestor
type GA uint32

const (
	GA_PARENT    GA = 1
	GA_ROOT      GA = 2
	GA_ROOTOWNER GA = 3
)

// DTM_SETSYSTEMTIME action.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/dtm-setsystemtime
type GDT uint32

const (
	GDT_VALID GDT = 0
	GDT_NONE  GDT = 1
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/secauthz/generic-access-rights
type GENERIC uint32

const (
	GENERIC_READ    GENERIC = 0x8000_0000
	GENERIC_WRITE   GENERIC = 0x4000_0000
	GENERIC_EXECUTE GENERIC = 0x2000_0000
	GENERIC_ALL     GENERIC = 0x1000_0000
)

// GetClassLong() nIndex. Includes values with GCW prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclasslongw
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

// GetDeviceCaps() index. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
type GDC int32

const (
	GDC_DRIVERVERSION   GDC = 0
	GDC_TECHNOLOGY      GDC = 2
	GDC_HORZSIZE        GDC = 4
	GDC_VERTSIZE        GDC = 6
	GDC_HORZRES         GDC = 8
	GDC_VERTRES         GDC = 10
	GDC_BITSPIXEL       GDC = 12
	GDC_PLANES          GDC = 14
	GDC_NUMBRUSHES      GDC = 16
	GDC_NUMPENS         GDC = 18
	GDC_NUMMARKERS      GDC = 20
	GDC_NUMFONTS        GDC = 22
	GDC_NUMCOLORS       GDC = 24
	GDC_PDEVICESIZE     GDC = 26
	GDC_CURVECAPS       GDC = 28
	GDC_LINECAPS        GDC = 30
	GDC_POLYGONALCAPS   GDC = 32
	GDC_TEXTCAPS        GDC = 34
	GDC_CLIPCAPS        GDC = 36
	GDC_RASTERCAPS      GDC = 38
	GDC_ASPECTX         GDC = 40
	GDC_ASPECTY         GDC = 42
	GDC_ASPECTXY        GDC = 44
	GDC_LOGPIXELSX      GDC = 88
	GDC_LOGPIXELSY      GDC = 90
	GDC_SIZEPALETTE     GDC = 104
	GDC_NUMRESERVED     GDC = 106
	GDC_COLORRES        GDC = 108
	GDC_PHYSICALWIDTH   GDC = 110
	GDC_PHYSICALHEIGHT  GDC = 111
	GDC_PHYSICALOFFSETX GDC = 112
	GDC_PHYSICALOFFSETY GDC = 113
	GDC_SCALINGFACTORX  GDC = 114
	GDC_SCALINGFACTORY  GDC = 115
	GDC_VREFRESH        GDC = 116
	GDC_DESKTOPVERTRES  GDC = 117
	GDC_DESKTOPHORZRES  GDC = 118
	GDC_BLTALIGNMENT    GDC = 119
	GDC_SHADEBLENDCAPS  GDC = 120
	GDC_COLORMGMTCAPS   GDC = 121
)

// GlobalAlloc() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
type GMEM uint32

const (
	GMEM_FIXED    GMEM = 0x0000
	GMEM_MOVEABLE GMEM = 0x0002
	GMEM_ZEROINIT GMEM = 0x0040
	GMEM_MODIFY   GMEM = 0x0080
	GMEM_GHND     GMEM = GMEM_MOVEABLE | GMEM_ZEROINIT
	GMEM_GPTR     GMEM = GMEM_FIXED | GMEM_ZEROINIT
)

// GetWindow() uCmd.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindow
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

// Get/SetWindowLongPtr() nIndex. Also includes constants with GWL prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
type GWLP int32

const (
	GWLP_STYLE          GWLP = -16
	GWLP_EXSTYLE        GWLP = -20
	GWLP_WNDPROC        GWLP = -4
	GWLP_HINSTANCE      GWLP = -6
	GWLP_HWNDPARENT     GWLP = -8
	GWLP_USERDATA       GWLP = -21
	GWLP_ID             GWLP = -12
	GWLP_DWLP_DLGPROC   GWLP = 8 // sizeof(LRESULT) on x64
	GWLP_DWLP_MSGRESULT GWLP = 0
	GWLP_DWLP_USER      GWLP = GWLP_DWLP_DLGPROC + 8 // sizeof(LRESULT) on x64
)

// SetWindowsHookEx() callback CBT hook codes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowshookexw
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

// HELPINFO iContextType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
type HELPINFO int32

const (
	HELPINFO_WINDOW   HELPINFO = 0x0001
	HELPINFO_MENUITEM HELPINFO = 0x0002
)

// NMBCHOTITEM and NMTBHOTITEM dwFlags, NMTBWRAPHOTITEM iReason.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmbchotitem
type HICF uint32

const (
	HICF_OTHER          HICF = 0x0000_0000
	HICF_MOUSE          HICF = 0x0000_0001
	HICF_ARROWKEYS      HICF = 0x0000_0002
	HICF_ACCELERATOR    HICF = 0x0000_0004
	HICF_DUPACCEL       HICF = 0x0000_0008
	HICF_ENTERING       HICF = 0x0000_0010
	HICF_LEAVING        HICF = 0x0000_0020
	HICF_RESELECT       HICF = 0x0000_0040
	HICF_LMOUSE         HICF = 0x0000_0080
	HICF_TOGGLEDROPDOWN HICF = 0x0000_0100
)

// CreateHatchBrush() iHatch.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhatchbrush
type HS int32

const (
	HS_HORIZONTAL HS = 0 // Pattern: -----
	HS_VERTICAL   HS = 1 // Pattern: |||||
	HS_FDIAGONAL  HS = 2 // Pattern: \\\\\
	HS_BDIAGONAL  HS = 3 // Pattern: /////
	HS_CROSS      HS = 4 // Pattern: +++++
	HS_DIAGCROSS  HS = 5 // Pattern: xxxxx
)

// WM_NCHITTEST return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
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

// TVINSERTSTRUCT hInsertAfter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvinsertstructw
type HTREEITEM uintptr

const (
	HTREEITEM_ROOT  HTREEITEM = 0x1_0000
	HTREEITEM_FIRST HTREEITEM = 0x0_ffff
	HTREEITEM_LAST  HTREEITEM = 0x0_fffe
	HTREEITEM_SORT  HTREEITEM = 0x0_fffd
)

// SetWindowPos() hwndInsertAfter. Can be converted to HWND.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
type HWND_IA int32

const (
	HWND_IA_NONE      HWND_IA = 0
	HWND_IA_BOTTOM    HWND_IA = 1
	HWND_IA_NOTOPMOST HWND_IA = -2
	HWND_IA_TOP       HWND_IA = 0
	HWND_IA_TOPMOST   HWND_IA = -1
)

// WM_SETICON icon size. Originally with ICON prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
type ICON_SZ uint8

const (
	ICON_SZ_SMALL  ICON_SZ = 0
	ICON_SZ_BIG    ICON_SZ = 1
	ICON_SZ_SMALL2 ICON_SZ = 2
)

// Dialog codes returned by MessageBox().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
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

// LoadCursor() lpCursorName.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
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

// WM_HOTKEY identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
type IDHOT int32

const (
	IDHOT_SNAPWINDOW  IDHOT = -1
	IDHOT_SNAPDESKTOP IDHOT = -2
)

// LoadIcon() lpIconName.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
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

// ImageList_Create() flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
type ILC uint32

const (
	ILC_MASK             ILC = 0x0000_0001
	ILC_COLOR            ILC = 0x0000_0000
	ILC_COLORDDB         ILC = 0x0000_00fe
	ILC_COLOR4           ILC = 0x0000_0004
	ILC_COLOR8           ILC = 0x0000_0008
	ILC_COLOR16          ILC = 0x0000_0010
	ILC_COLOR24          ILC = 0x0000_0018
	ILC_COLOR32          ILC = 0x0000_0020
	ILC_PALETTE          ILC = 0x0000_0800
	ILC_MIRROR           ILC = 0x0000_2000
	ILC_PERITEMMIRROR    ILC = 0x0000_8000
	ILC_ORIGINALSIZE     ILC = 0x0001_0000
	ILC_HIGHQUALITYSCALE ILC = 0x0002_0000
)

// ImageList_Draw() flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_draw
type ILD uint32

const (
	ILD_NORMAL        ILD = 0x0000_0000
	ILD_TRANSPARENT   ILD = 0x0000_0001
	ILD_MASK          ILD = 0x0000_0010
	ILD_IMAGE         ILD = 0x0000_0020
	ILD_ROP           ILD = 0x0000_0040
	ILD_BLEND25       ILD = 0x0000_0002
	ILD_BLEND50       ILD = 0x0000_0004
	ILD_OVERLAYMASK   ILD = 0x0000_0f00
	ILD_PRESERVEALPHA ILD = 0x0000_1000
	ILD_SCALE         ILD = 0x0000_2000
	ILD_DPISCALE      ILD = 0x0000_4000
	ILD_ASYNC         ILD = 0x0000_8000
	ILD_SELECTED      ILD = ILD_BLEND50
	ILD_FOCUS         ILD = ILD_BLEND25
	ILD_BLEND         ILD = ILD_BLEND50
)

// ImageList state flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/imageliststateflags
type ILS uint32

const (
	ILS_NORMAL   ILS = 0x0000_0000
	ILS_GLOW     ILS = 0x0000_0001
	ILS_SHADOW   ILS = 0x0000_0002
	ILS_SATURATE ILS = 0x0000_0004
	ILS_ALPHA    ILS = 0x0000_0008
)

// LoadImage() type.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
type IMAGE uint32

const (
	IMAGE_BITMAP      IMAGE = 0
	IMAGE_ICON        IMAGE = 1
	IMAGE_CURSOR      IMAGE = 2
	IMAGE_ENHMETAFILE IMAGE = 3
)
