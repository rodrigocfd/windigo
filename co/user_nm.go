//go:build windows

package co

// [WM_COMMAND] notification codes.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
type CMD uint16

const (
	CMD_MENU        CMD = 0 // Message originated from a menu.
	CMD_ACCELERATOR CMD = 1 // Message originated from an accelerator.
)

// [WM_NOTIFY] common control notifications.
//
// [WM_NOTIFY]: https://learn.microsoft.com/en-us/windows/win32/controls/common-control-reference#notifications
type NM int32

const (
	_NM_FIRST NM = 0

	NM_OUTOFMEMORY          = _NM_FIRST - 1
	NM_CLICK                = _NM_FIRST - 2
	NM_DBLCLK               = _NM_FIRST - 3
	NM_RETURN               = _NM_FIRST - 4
	NM_RCLICK               = _NM_FIRST - 5
	NM_RDBLCLK              = _NM_FIRST - 6
	NM_SETFOCUS             = _NM_FIRST - 7
	NM_KILLFOCUS            = _NM_FIRST - 8
	NM_CUSTOMDRAW           = _NM_FIRST - 12
	NM_HOVER                = _NM_FIRST - 13
	NM_NCHITTEST            = _NM_FIRST - 14
	NM_KEYDOWN              = _NM_FIRST - 15
	NM_RELEASEDCAPTURE      = _NM_FIRST - 16
	NM_SETCURSOR            = _NM_FIRST - 17
	NM_CHAR                 = _NM_FIRST - 18
	NM_TOOLTIPSCREATED      = _NM_FIRST - 19
	NM_LDOWN                = _NM_FIRST - 20
	NM_RDOWN                = _NM_FIRST - 21
	NM_THEMECHANGED         = _NM_FIRST - 22
	NM_FONTCHANGED          = _NM_FIRST - 23
	NM_CUSTOMTEXT           = _NM_FIRST - 24
	NM_TVSTATEIMAGECHANGING = _NM_FIRST - 24
)

// Button control [notifications] (BCN, BN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
const (
	_BCN_FIRST NM = -1250

	BCN_HOTITEMCHANGE = _BCN_FIRST + 0x0001
	BCN_DROPDOWN      = _BCN_FIRST + 0x0002

	BN_CLICKED       CMD = 0
	BN_PAINT         CMD = 1
	BN_HILITE        CMD = 2
	BN_UNHILITE      CMD = 3
	BN_DISABLE       CMD = 4
	BN_DOUBLECLICKED CMD = 5
	BN_PUSHED            = BN_HILITE
	BN_UNPUSHED          = BN_UNHILITE
	BN_DBLCLK            = BN_DOUBLECLICKED
	BN_SETFOCUS      CMD = 6
	BN_KILLFOCUS     CMD = 7
)

// ComboBox control [notifications] (CBN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-combobox-control-reference-notifications
const (
	CBN_ERRSPACE     CMD = 0xffff
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

// ComboBoxEx control [notifications] (CBEN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-comboboxex-control-reference-notifications
const (
	_CBEN_FIRST NM = -800

	CBEN_INSERTITEM  = _CBEN_FIRST - 1
	CBEN_DELETEITEM  = _CBEN_FIRST - 2
	CBEN_BEGINEDIT   = _CBEN_FIRST - 4
	CBEN_ENDEDIT     = _CBEN_FIRST - 6
	CBEN_GETDISPINFO = _CBEN_FIRST - 7
	CBEN_DRAGBEGIN   = _CBEN_FIRST - 9
)

// DateTimePicker control [notifications] (DTN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-date-and-time-picker-control-reference-notifications
const (
	_DTN_FIRST  NM = -740
	_DTN_FIRST2 NM = -753

	DTN_CLOSEUP        = _DTN_FIRST2 - 0
	DTN_DROPDOWN       = _DTN_FIRST2 - 1
	DTN_DATETIMECHANGE = _DTN_FIRST2 - 6
	DTN_FORMATQUERY    = _DTN_FIRST - 2
	DTN_FORMAT         = _DTN_FIRST - 3
	DTN_WMKEYDOWN      = _DTN_FIRST - 4
	DTN_USERSTRING     = _DTN_FIRST - 5
)

// Edit control [notifications] (EN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-notifications
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

// Header control [notifications] (HDN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-header-control-reference-notifications
const (
	_HDN_FIRST NM = -300

	HDN_ITEMCHANGING       = _HDN_FIRST - 20
	HDN_ITEMCHANGED        = _HDN_FIRST - 21
	HDN_ITEMCLICK          = _HDN_FIRST - 22
	HDN_ITEMDBLCLICK       = _HDN_FIRST - 23
	HDN_DIVIDERDBLCLICK    = _HDN_FIRST - 25
	HDN_BEGINTRACK         = _HDN_FIRST - 26
	HDN_ENDTRACK           = _HDN_FIRST - 27
	HDN_TRACK              = _HDN_FIRST - 28
	HDN_GETDISPINFO        = _HDN_FIRST - 29
	HDN_BEGINDRAG          = _HDN_FIRST - 10
	HDN_ENDDRAG            = _HDN_FIRST - 11
	HDN_FILTERCHANGE       = _HDN_FIRST - 12
	HDN_FILTERBTNCLICK     = _HDN_FIRST - 13
	HDN_BEGINFILTEREDIT    = _HDN_FIRST - 14
	HDN_ENDFILTEREDIT      = _HDN_FIRST - 15
	HDN_ITEMSTATEICONCLICK = _HDN_FIRST - 16
	HDN_ITEMKEYDOWN        = _HDN_FIRST - 17
	HDN_DROPDOWN           = _HDN_FIRST - 18
	HDN_OVERFLOWCLICK      = _HDN_FIRST - 19
)

// IpAddress [notifications] (IPN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-ip-address-control-reference-notifications
const (
	_IPN_FIRST NM = -860

	IPN_FIELDCHANGED = _IPN_FIRST - 0
)

// ListView control [notifications] (LVN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-list-view-control-reference-notifications
const (
	_LVN_FIRST NM = -100

	LVN_ITEMCHANGING        = _LVN_FIRST - 0
	LVN_ITEMCHANGED         = _LVN_FIRST - 1
	LVN_INSERTITEM          = _LVN_FIRST - 2
	LVN_DELETEITEM          = _LVN_FIRST - 3
	LVN_DELETEALLITEMS      = _LVN_FIRST - 4
	LVN_BEGINLABELEDIT      = _LVN_FIRST - 75
	LVN_ENDLABELEDIT        = _LVN_FIRST - 76
	LVN_COLUMNCLICK         = _LVN_FIRST - 8
	LVN_BEGINDRAG           = _LVN_FIRST - 9
	LVN_BEGINRDRAG          = _LVN_FIRST - 11
	LVN_ODCACHEHINT         = _LVN_FIRST - 13
	LVN_ODFINDITEM          = _LVN_FIRST - 79
	LVN_ITEMACTIVATE        = _LVN_FIRST - 14
	LVN_ODSTATECHANGED      = _LVN_FIRST - 15
	LVN_HOTTRACK            = _LVN_FIRST - 21
	LVN_GETDISPINFO         = _LVN_FIRST - 77
	LVN_SETDISPINFO         = _LVN_FIRST - 78
	LVN_KEYDOWN             = _LVN_FIRST - 55
	LVN_MARQUEEBEGIN        = _LVN_FIRST - 56
	LVN_GETINFOTIP          = _LVN_FIRST - 58
	LVN_INCREMENTALSEARCH   = _LVN_FIRST - 63
	LVN_COLUMNDROPDOWN      = _LVN_FIRST - 64
	LVN_COLUMNOVERFLOWCLICK = _LVN_FIRST - 66
	LVN_BEGINSCROLL         = _LVN_FIRST - 80
	LVN_ENDSCROLL           = _LVN_FIRST - 81
	LVN_LINKCLICK           = _LVN_FIRST - 84
	LVN_GETEMPTYMARKUP      = _LVN_FIRST - 87
)

// MonthCalendar control [notifications] (MCN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-month-calendar-control-reference-notifications
const (
	_MCN_FIRST NM = -746

	MCN_SELCHANGE   = _MCN_FIRST - 3
	MCN_GETDAYSTATE = _MCN_FIRST - 1
	MCN_SELECT      = _MCN_FIRST
	MCN_VIEWCHANGE  = _MCN_FIRST - 4
)

// Rebar control [notifications] (RBN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-rebar-control-reference-notifications
const (
	_RBN_FIRST NM = -831

	RBN_HEIGHTCHANGE  = _RBN_FIRST - 0
	RBN_GETOBJECT     = _RBN_FIRST - 1
	RBN_LAYOUTCHANGED = _RBN_FIRST - 2
	RBN_AUTOSIZE      = _RBN_FIRST - 3
	RBN_BEGINDRAG     = _RBN_FIRST - 4
	RBN_ENDDRAG       = _RBN_FIRST - 5
	RBN_DELETINGBAND  = _RBN_FIRST - 6
	RBN_DELETEDBAND   = _RBN_FIRST - 7
	RBN_CHILDSIZE     = _RBN_FIRST - 8
	RBN_CHEVRONPUSHED = _RBN_FIRST - 10
	RBN_SPLITTERDRAG  = _RBN_FIRST - 11
	RBN_MINMAX        = _RBN_FIRST - 21
	RBN_AUTOBREAK     = _RBN_FIRST - 22
)

// StatusBar control [notifications] (SBN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-status-bars-reference-notifications
const (
	_SBN_FIRST NM = -880

	SBN_SIMPLEMODECHANGE = _SBN_FIRST - 0
)

// Static control [notifications] (STN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-static-control-reference-notifications
const (
	STN_CLICKED CMD = 0
	STN_DBLCLK  CMD = 1
	STN_ENABLE  CMD = 2
	STN_DISABLE CMD = 3
)

// Toolbar control [notifications] (TBN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-toolbar-control-reference-notifications
const (
	_TBN_FIRST NM = -700

	TBN_BEGINDRAG       = _TBN_FIRST - 1
	TBN_ENDDRAG         = _TBN_FIRST - 2
	TBN_BEGINADJUST     = _TBN_FIRST - 3
	TBN_ENDADJUST       = _TBN_FIRST - 4
	TBN_RESET           = _TBN_FIRST - 5
	TBN_QUERYINSERT     = _TBN_FIRST - 6
	TBN_QUERYDELETE     = _TBN_FIRST - 7
	TBN_TOOLBARCHANGE   = _TBN_FIRST - 8
	TBN_CUSTHELP        = _TBN_FIRST - 9
	TBN_DROPDOWN        = _TBN_FIRST - 10
	TBN_GETOBJECT       = _TBN_FIRST - 12
	TBN_HOTITEMCHANGE   = _TBN_FIRST - 13
	TBN_DRAGOUT         = _TBN_FIRST - 14
	TBN_DELETINGBUTTON  = _TBN_FIRST - 15
	TBN_GETDISPINFO     = _TBN_FIRST - 17
	TBN_GETINFOTIP      = _TBN_FIRST - 19
	TBN_GETBUTTONINFO   = _TBN_FIRST - 20
	TBN_RESTORE         = _TBN_FIRST - 21
	TBN_SAVE            = _TBN_FIRST - 22
	TBN_INITCUSTOMIZE   = _TBN_FIRST - 23
	TBN_WRAPHOTITEM     = _TBN_FIRST - 24
	TBN_DUPACCELERATOR  = _TBN_FIRST - 25
	TBN_WRAPACCELERATOR = _TBN_FIRST - 26
	TBN_DRAGOVER        = _TBN_FIRST - 27
	TBN_MAPACCELERATOR  = _TBN_FIRST - 28
)

// Tab control [notifications] (TCN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-tab-control-reference-notifications
const (
	_TCN_FIRST NM = -550

	TCN_KEYDOWN     = _TCN_FIRST - 0
	TCN_SELCHANGE   = _TCN_FIRST - 1
	TCN_SELCHANGING = _TCN_FIRST - 2
	TCN_GETOBJECT   = _TCN_FIRST - 3
	TCN_FOCUSCHANGE = _TCN_FIRST - 4
)

// Trackbar control [notifications] (TRBN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-trackbar-control-reference-notifications
const (
	_TRBN_FIRST NM = -1501

	TRBN_THUMBPOSCHANGING = _TRBN_FIRST - 1
)

// Tooltip control [notifications] (TTN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-tooltip-control-reference-notifications
const (
	_TTN_FIRST NM = -520

	TTN_GETDISPINFO = _TTN_FIRST - 10
	TTN_SHOW        = _TTN_FIRST - 1
	TTN_POP         = _TTN_FIRST - 2
	TTN_LINKCLICK   = _TTN_FIRST - 3
	TTN_NEEDTEXT    = TTN_GETDISPINFO
)

// TreeView control [notifications] (TVN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-tree-view-control-reference-notifications
const (
	_TVN_FIRST NM = -400

	TVN_SELCHANGING    = _TVN_FIRST - 50
	TVN_SELCHANGED     = _TVN_FIRST - 51
	TVN_GETDISPINFO    = _TVN_FIRST - 52
	TVN_SETDISPINFO    = _TVN_FIRST - 53
	TVN_ITEMEXPANDING  = _TVN_FIRST - 54
	TVN_ITEMEXPANDED   = _TVN_FIRST - 55
	TVN_BEGINDRAG      = _TVN_FIRST - 56
	TVN_BEGINRDRAG     = _TVN_FIRST - 57
	TVN_DELETEITEM     = _TVN_FIRST - 58
	TVN_BEGINLABELEDIT = _TVN_FIRST - 59
	TVN_ENDLABELEDIT   = _TVN_FIRST - 60
	TVN_KEYDOWN        = _TVN_FIRST - 12
	TVN_GETINFOTIP     = _TVN_FIRST - 14
	TVN_SINGLEEXPAND   = _TVN_FIRST - 15
	TVN_ITEMCHANGING   = _TVN_FIRST - 17
	TVN_ITEMCHANGED    = _TVN_FIRST - 19
	TVN_ASYNCDRAW      = _TVN_FIRST - 20
)

// UpDown control [notifications] (UDN).
//
// [notifications]: https://learn.microsoft.com/en-us/windows/win32/controls/bumper-up-down-control-reference-notifications
const (
	_UDN_FIRST NM = -721

	UDN_DELTAPOS = _UDN_FIRST - 1
)
