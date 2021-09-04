package co

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

// Button control notifications (BCN, BN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
const (
	_BCN_FIRST NM = -1250

	BCN_HOTITEMCHANGE NM = _BCN_FIRST + 0x0001
	BCN_DROPDOWN      NM = _BCN_FIRST + 0x0002

	BN_CLICKED       CMD = 0
	BN_PAINT         CMD = 1
	BN_HILITE        CMD = 2
	BN_UNHILITE      CMD = 3
	BN_DISABLE       CMD = 4
	BN_DOUBLECLICKED CMD = 5
	BN_PUSHED        CMD = BN_HILITE
	BN_UNPUSHED      CMD = BN_UNHILITE
	BN_DBLCLK        CMD = BN_DOUBLECLICKED
	BN_SETFOCUS      CMD = 6
	BN_KILLFOCUS     CMD = 7
)

// ComboBox control notifications (CBN).
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

// DateTimePicker control notifications (DTN).
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

// Edit control notifications (EN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-notifications
const (
	EN_SETFOCUS     CMD = 0x0100
	EN_KILLFOCUS    CMD = 0x0200
	EN_CHANGE       CMD = 0x0300
	EN_UPDATE       CMD = 0x0400
	EN_ERRSPACE     CMD = 0x0500
	EN_MAXTEXT      CMD = 0x0501
	EN_HSCROLL      CMD = 0x0601
	EN_VSCROLL      CMD = 0x0602
	EN_ALIGN_LTR_EC CMD = 0x0700
	EN_ALIGN_RTL_EC CMD = 0x0701
	EN_BEFORE_PASTE CMD = 0x0800
	EN_AFTER_PASTE  CMD = 0x0801
)

// ListView control notifications (LVN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-list-view-control-reference-notifications
const (
	_LVN_FIRST NM = -100

	LVN_ITEMCHANGING        NM = _LVN_FIRST - 0
	LVN_ITEMCHANGED         NM = _LVN_FIRST - 1
	LVN_INSERTITEM          NM = _LVN_FIRST - 2
	LVN_DELETEITEM          NM = _LVN_FIRST - 3
	LVN_DELETEALLITEMS      NM = _LVN_FIRST - 4
	LVN_BEGINLABELEDIT      NM = _LVN_FIRST - 75
	LVN_ENDLABELEDIT        NM = _LVN_FIRST - 76
	LVN_COLUMNCLICK         NM = _LVN_FIRST - 8
	LVN_BEGINDRAG           NM = _LVN_FIRST - 9
	LVN_BEGINRDRAG          NM = _LVN_FIRST - 11
	LVN_ODCACHEHINT         NM = _LVN_FIRST - 13
	LVN_ODFINDITEM          NM = _LVN_FIRST - 79
	LVN_ITEMACTIVATE        NM = _LVN_FIRST - 14
	LVN_ODSTATECHANGED      NM = _LVN_FIRST - 15
	LVN_HOTTRACK            NM = _LVN_FIRST - 21
	LVN_GETDISPINFO         NM = _LVN_FIRST - 77
	LVN_SETDISPINFO         NM = _LVN_FIRST - 78
	LVN_KEYDOWN             NM = _LVN_FIRST - 55
	LVN_MARQUEEBEGIN        NM = _LVN_FIRST - 56
	LVN_GETINFOTIP          NM = _LVN_FIRST - 58
	LVN_INCREMENTALSEARCH   NM = _LVN_FIRST - 63
	LVN_COLUMNDROPDOWN      NM = _LVN_FIRST - 64
	LVN_COLUMNOVERFLOWCLICK NM = _LVN_FIRST - 66
	LVN_BEGINSCROLL         NM = _LVN_FIRST - 80
	LVN_ENDSCROLL           NM = _LVN_FIRST - 81
	LVN_LINKCLICK           NM = _LVN_FIRST - 84
	LVN_GETEMPTYMARKUP      NM = _LVN_FIRST - 87
)

// MonthCalendar control notifications (MCN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-month-calendar-control-reference-notifications
const (
	_MCN_FIRST NM = -746

	MCN_SELECT      NM = _MCN_FIRST
	MCN_GETDAYSTATE NM = _MCN_FIRST - 1
	MCN_SELCHANGE   NM = _MCN_FIRST - 3
	MCN_VIEWCHANGE  NM = _MCN_FIRST - 4
)

// StatusBar control notifications (SBN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-status-bars-reference-notifications
const (
	_SBN_FIRST NM = -880

	SBN_SIMPLEMODECHANGE NM = _SBN_FIRST - 0
)

// Static control notifications (STN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-static-control-reference-notifications
const (
	STN_CLICKED CMD = 0
	STN_DBLCLK  CMD = 1
	STN_ENABLE  CMD = 2
	STN_DISABLE CMD = 3
)

// Trackbar control notifications (TRBN).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-notifications
const (
	_TRBN_FIRST NM = -1501

	TRBN_THUMBPOSCHANGING NM = _TRBN_FIRST - 1
)

// TreeView control notifications (TVN).
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
