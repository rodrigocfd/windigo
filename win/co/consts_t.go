package co

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
	TBS_AUTOTICKS        TBS = 0x1    // The trackbar control has a tick mark for each increment in its range of values.
	TBS_VERT             TBS = 0x2    // The trackbar control is oriented vertically.
	TBS_HORZ             TBS = 0x0    // The trackbar control is oriented horizontally. This is the default orientation.
	TBS_TOP              TBS = 0x4    // The trackbar control displays tick marks above the control. This style is valid only with TBS_HORZ.
	TBS_BOTTOM           TBS = 0x0    // The trackbar control displays tick marks below the control. This style is valid only with TBS_HORZ.
	TBS_LEFT             TBS = 0x4    // The trackbar control displays tick marks to the left of the control. This style is valid only with TBS_VERT.
	TBS_RIGHT            TBS = 0x0    // The trackbar control displays tick marks to the right of the control. This style is valid only with TBS_VERT.
	TBS_BOTH             TBS = 0x8    // The trackbar control displays tick marks on both sides of the control. This will be both top and bottom when used with TBS_HORZ or both left and right if used with TBS_VERT.
	TBS_NOTICKS          TBS = 0x10   // The trackbar control does not display any tick marks.
	TBS_ENABLESELRANGE   TBS = 0x20   // The trackbar control displays a selection range only. The tick marks at the starting and ending positions of a selection range are displayed as triangles (instead of vertical dashes), and the selection range is highlighted.
	TBS_FIXEDLENGTH      TBS = 0x40   // The trackbar control allows the size of the slider to be changed with the TBM_SETTHUMBLENGTH message.
	TBS_NOTHUMB          TBS = 0x80   // The trackbar control does not display a slider.
	TBS_TOOLTIPS         TBS = 0x100  // The trackbar control supports tooltips. When a trackbar control is created using this style, it automatically creates a default tooltip control that displays the slider's current position. You can change where the tooltips are displayed by using the TBM_SETTIPSIDE message.
	TBS_REVERSED         TBS = 0x200  // This style bit is used for "reversed" trackbars, where a smaller number indicates "higher" and a larger number indicates "lower." It has no effect on the control; it is simply a label that can be checked to determine whether a trackbar is normal or reversed.
	TBS_DOWNISLEFT       TBS = 0x400  // By default, the trackbar control uses down equal to right and up equal to left. Use the TBS_DOWNISLEFT style to reverse the default, making down equal left and up equal right.
	TBS_NOTIFYBEFOREMOVE TBS = 0x800  // Trackbar should notify parent before repositioning the slider due to user action (enables snapping).
	TBS_TRANSPARENTBKGND TBS = 0x1000 // Background is painted by the parent via the WM_PRINTCLIENT message.
)

// TaskDialog() pszIcon; originally with TD prefix and ICON suffix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
type TD_ICON uint16

const (
	TD_ICON_WARNING     TD_ICON = 0xffff
	TD_ICON_ERROR       TD_ICON = 0xfffe
	TD_ICON_INFORMATION TD_ICON = 0xfffd
	TD_ICON_SHIELD      TD_ICON = 0xfffc
)

// TaskDialog() dwCommonButtons. Originally has BUTTON suffix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
type TDCBF int32

const (
	TDCBF_OK     TDCBF = 0x0001
	TDCBF_YES    TDCBF = 0x0002
	TDCBF_NO     TDCBF = 0x0004
	TDCBF_CANCEL TDCBF = 0x0008
	TDCBF_RETRY  TDCBF = 0x0010
	TDCBF_CLOSE  TDCBF = 0x0020
)

// TASKDIALOGCONFIG dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/Commctrl/ns-commctrl-taskdialogconfig
type TDF int32

const (
	TDF_ENABLE_HYPERLINKS           TDF = 0x0001
	TDF_USE_HICON_MAIN              TDF = 0x0002
	TDF_USE_HICON_FOOTER            TDF = 0x0004
	TDF_ALLOW_DIALOG_CANCELLATION   TDF = 0x0008
	TDF_USE_COMMAND_LINKS           TDF = 0x0010
	TDF_USE_COMMAND_LINKS_NO_ICON   TDF = 0x0020
	TDF_EXPAND_FOOTER_AREA          TDF = 0x0040
	TDF_EXPANDED_BY_DEFAULT         TDF = 0x0080
	TDF_VERIFICATION_FLAG_CHECKED   TDF = 0x0100
	TDF_SHOW_PROGRESS_BAR           TDF = 0x0200
	TDF_SHOW_MARQUEE_PROGRESS_BAR   TDF = 0x0400
	TDF_CALLBACK_TIMER              TDF = 0x0800
	TDF_POSITION_RELATIVE_TO_WINDOW TDF = 0x1000
	TDF_RTL_LAYOUT                  TDF = 0x2000
	TDF_NO_DEFAULT_RADIO_BUTTON     TDF = 0x4000
	TDF_CAN_BE_MINIMIZED            TDF = 0x8000
	TDF_NO_SET_FOREGROUND           TDF = 0x00010000
	TDF_SIZE_TO_CONTENT             TDF = 0x01000000
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
