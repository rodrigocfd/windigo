//go:build windows

package co

// [NMTVASYNCDRAW] dwRetFlags, don't seem to be defined anywhere, values are unconfirmed.
//
// [NMTVASYNCDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type ADRF uint32

const (
	ADRF_DRAWSYNC     ADRF = 0
	ADRF_DRAWNOTHING  ADRF = 1
	ADRF_DRAWFALLBACK ADRF = 2
	ADRF_DRAWIMAGE    ADRF = 3
)

// Toolbar button [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-control-and-button-styles
type BTNS uint8

const (
	BTNS_BUTTON        BTNS = BTNS(TBSTYLE_BUTTON)
	BTNS_SEP           BTNS = BTNS(TBSTYLE_SEP)
	BTNS_CHECK         BTNS = BTNS(TBSTYLE_CHECK)
	BTNS_GROUP         BTNS = BTNS(TBSTYLE_GROUP)
	BTNS_CHECKGROUP    BTNS = BTNS(TBSTYLE_CHECKGROUP)
	BTNS_DROPDOWN      BTNS = BTNS(TBSTYLE_DROPDOWN)
	BTNS_AUTOSIZE      BTNS = BTNS(TBSTYLE_AUTOSIZE)
	BTNS_NOPREFIX      BTNS = BTNS(TBSTYLE_NOPREFIX)
	BTNS_SHOWTEXT      BTNS = 0x0040
	BTNS_WHOLEDROPDOWN BTNS = 0x0080
)

// [NMCUSTOMDRAW] dwDrawStage.
//
// [NMCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type CDDS uint32

const (
	CDDS_PREPAINT      CDDS = 0x0000_0001
	CDDS_POSTPAINT     CDDS = 0x0000_0002
	CDDS_PREERASE      CDDS = 0x0000_0003
	CDDS_POSTERASE     CDDS = 0x0000_0004
	CDDS_ITEM          CDDS = 0x0001_0000
	CDDS_ITEMPREPAINT  CDDS = CDDS_ITEM | CDDS_PREPAINT
	CDDS_ITEMPOSTPAINT CDDS = CDDS_ITEM | CDDS_POSTPAINT
	CDDS_ITEMPREERASE  CDDS = CDDS_ITEM | CDDS_PREERASE
	CDDS_ITEMPOSTERASE CDDS = CDDS_ITEM | CDDS_POSTERASE
	CDDS_SUBITEM       CDDS = 0x0002_0000
)

// [NMCUSTOMDRAW] uItemState.
//
// [NMCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
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

// [NM_CUSTOMDRAW] return value.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type CDRF uint32

const (
	CDRF_DODEFAULT         CDRF = 0x0000_0000
	CDRF_NEWFONT           CDRF = 0x0000_0002
	CDRF_SKIPDEFAULT       CDRF = 0x0000_0004
	CDRF_DOERASE           CDRF = 0x0000_0008
	CDRF_SKIPPOSTPAINT     CDRF = 0x0000_0100
	CDRF_NOTIFYPOSTPAINT   CDRF = 0x0000_0010
	CDRF_NOTIFYITEMDRAW    CDRF = 0x0000_0020
	CDRF_NOTIFYSUBITEMDRAW CDRF = 0x0000_0020
	CDRF_NOTIFYPOSTERASE   CDRF = 0x0000_0040
)

// Clipboard [formats].
//
// [formats]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/standard-clipboard-formats
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

// DateTimePicker control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/date-and-time-picker-control-styles
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

// [NMLVEMPTYMARKUP] dwFlags.
//
// [NMLVEMPTYMARKUP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvemptymarkup
type EMF uint32

const (
	EMF_NULL     EMF = 0x0000_0000
	EMF_CENTERED EMF = 0x0000_0001
)

// [DTM_SETSYSTEMTIME] action.
//
// [DTM_SETSYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/controls/dtm-setsystemtime
type GDT uint32

const (
	GDT_VALID GDT = 0
	GDT_NONE  GDT = 1
)

// [NMBCHOTITEM] and [NMTBHOTITEM] dwFlags, [NMTBWRAPHOTITEM] iReason.
//
// [NMBCHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmbchotitem
// [NMTBHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbhotitem
// [NMTBWRAPHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-wraphotitem
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

// [INITCOMMONCONTROLSEX] dwIcc.
//
// [INITCOMMONCONTROLSEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-initcommoncontrolsex
type ICC uint32

const (
	ICC_ANIMATE_CLASS      ICC = 0x0000_0080 // Load animate control class.
	ICC_BAR_CLASSES        ICC = 0x0000_0004 // Load toolbar, status bar, trackbar, and tooltip control classes.
	ICC_COOL_CLASSES       ICC = 0x0000_0400 // Load rebar control class.
	ICC_DATE_CLASSES       ICC = 0x0000_0100 // Load date and time picker control class.
	ICC_HOTKEY_CLASS       ICC = 0x0000_0040 // Load hot key control class.
	ICC_INTERNET_CLASSES   ICC = 0x0000_0800 // Load IP address class.
	ICC_LINK_CLASS         ICC = 0x0000_8000 // Load a hyperlink control class.
	ICC_LISTVIEW_CLASSES   ICC = 0x0000_0001 // Load list-view and header control classes.
	ICC_NATIVEFNTCTL_CLASS ICC = 0x0000_2000 // Load a native font control class.
	ICC_PAGESCROLLER_CLASS ICC = 0x0000_1000 // Load pager control class.
	ICC_PROGRESS_CLASS     ICC = 0x0000_0020 // Load progress bar control class.
	ICC_STANDARD_CLASSES   ICC = 0x0000_4000 // Load one of the intrinsic User32 control classes. The user controls include button, edit, static, listbox, combobox, and scroll bar.
	ICC_TAB_CLASSES        ICC = 0x0000_0008 // Load tab and tooltip control classes.
	ICC_TREEVIEW_CLASSES   ICC = 0x0000_0002 // Load tree-view and tooltip control classes.
	ICC_UPDOWN_CLASS       ICC = 0x0000_0010 // Load up-down control class.
	ICC_USEREX_CLASSES     ICC = 0x0000_0200 // Load ComboBoxEx class.
	ICC_WIN95_CLASSES      ICC = 0x0000_00ff // Load animate control, header, hot key, list-view, progress bar, status bar, tab, tooltip, toolbar, trackbar, tree-view, and up-down control classes.
)

// [ImageList_Create] flags.
//
// [ImageList_Create]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
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

// [ImageList_Draw] flags.
//
// [ImageList_Draw]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_draw
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

// ImageList state [flags].
//
// [flags]: https://learn.microsoft.com/en-us/windows/win32/controls/imageliststateflags
type ILS uint32

const (
	ILS_NORMAL   ILS = 0x0000_0000
	ILS_GLOW     ILS = 0x0000_0001
	ILS_SHADOW   ILS = 0x0000_0002
	ILS_SATURATE ILS = 0x0000_0004
	ILS_ALPHA    ILS = 0x0000_0008
)

// [LITEM] mask.
//
// [LITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LIF uint32

const (
	LIF_ITEMINDEX LIF = 0x0000_0001
	LIF_STATE     LIF = 0x0000_0002
	LIF_ITEMID    LIF = 0x0000_0004
	LIF_URL       LIF = 0x0000_0008
)

// [LITEM] state.
//
// [LITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LIS uint32

const (
	LIS_FOCUSED       LIS = 0x0000_0001
	LIS_ENABLED       LIS = 0x0000_0002
	LIS_VISITED       LIS = 0x0000_0004
	LIS_HOTTRACK      LIS = 0x0000_0008
	LIS_DEFAULTCOLORS LIS = 0x0000_0010
)

// [LVM_GETVIEW] return value.
//
// [LVM_GETVIEW]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getview
type LV_VIEW uint32

const (
	LV_VIEW_ICON      LV_VIEW = 0x0000
	LV_VIEW_DETAILS   LV_VIEW = 0x0001
	LV_VIEW_SMALLICON LV_VIEW = 0x0002
	LV_VIEW_LIST      LV_VIEW = 0x0003
	LV_VIEW_TILE      LV_VIEW = 0x0004
)

// [NMLVCUSTOMDRAW] dwItemType.
//
// [NMLVCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
type LVCDI uint32

const (
	LVCDI_ITEM     LVCDI = 0x0000_0000
	LVCDI_GROUP    LVCDI = 0x0000_0001
	LVCDI_TEMSLIST LVCDI = 0x0000_0002
)

// [LVCOLUMN] mask.
//
// {LVCOLUMN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
type LVCF uint32

const (
	LVCF_DEFAULTWIDTH LVCF = 0x0080
	LVCF_FMT          LVCF = 0x0001
	LVCF_IDEALWIDTH   LVCF = 0x0100
	LVCF_IMAGE        LVCF = 0x0010
	LVCF_MINWIDTH     LVCF = 0x0040
	LVCF_ORDER        LVCF = 0x0020
	LVCF_SUBITEM      LVCF = 0x0008
	LVCF_TEXT         LVCF = 0x0004
	LVCF_WIDTH        LVCF = 0x0002
)

// [LVCOLUMN] fmt.
//
// [LVCOLUMN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
type LVCFMT_C int32

const (
	LVCFMT_C_LEFT            LVCFMT_C = 0x0000
	LVCFMT_C_RIGHT           LVCFMT_C = 0x0001
	LVCFMT_C_CENTER          LVCFMT_C = 0x0002
	LVCFMT_C_JUSTIFYMASK     LVCFMT_C = 0x0003
	LVCFMT_C_IMAGE           LVCFMT_C = 0x0800
	LVCFMT_C_BITMAP_ON_RIGHT LVCFMT_C = 0x1000
	LVCFMT_C_COL_HAS_IMAGES  LVCFMT_C = 0x8000
	LVCFMT_C_FIXED_WIDTH     LVCFMT_C = 0x0_0100
	LVCFMT_C_NO_DPI_SCALE    LVCFMT_C = 0x4_0000
	LVCFMT_C_FIXED_RATIO     LVCFMT_C = 0x8_0000
	LVCFMT_C_SPLITBUTTON     LVCFMT_C = 0x100_0000
)

// [LVITEM] piColFmt.
//
// [LVITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVCFMT_I int32

const (
	LVCFMT_I_LINE_BREAK         LVCFMT_I = 0x10_0000
	LVCFMT_I_FILL               LVCFMT_I = 0x20_0000
	LVCFMT_I_WRAP               LVCFMT_I = 0x40_0000
	LVCFMT_I_NO_TITLE           LVCFMT_I = 0x80_0000
	LVCFMT_I_TILE_PLACEMENTMASK LVCFMT_I = LVCFMT_I_LINE_BREAK | LVCFMT_I_FILL
)

// [LVFINDINFO] flags.
//
// [LVFINDINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvfindinfow
type LVFI uint32

const (
	LVFI_PARAM     LVFI = 0x0001
	LVFI_STRING    LVFI = 0x0002
	LVFI_SUBSTRING LVFI = 0x0004
	LVFI_PARTIAL   LVFI = 0x0008
	LVFI_WRAP      LVFI = 0x0020
	LVFI_NEARESTXY LVFI = 0x0040
)

// [NMLVCUSTOMDRAW] uAlign.
//
// [NMLVCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
type LVGA_HEADER uint32

const (
	LVGA_HEADER_LEFT   LVGA_HEADER = 0x0000_0001
	LVGA_HEADER_CENTER LVGA_HEADER = 0x0000_0002
	LVGA_HEADER_RIGHT  LVGA_HEADER = 0x0000_0004
)

// [NMLVGETINFOTIP] dwFlags.
//
// [NMLVGETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvgetinfotipw
type LVGIT uint32

const (
	LVGIT_ZERO     LVGIT = 0x0000
	LVGIT_UNFOLDED LVGIT = 0x0001
)

// [LVITEM] iGroupId.
//
// [LVITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVI_GROUPID int32

const (
	LVI_GROUPID_I_GROUPIDCALLBACK LVI_GROUPID = -1
	LVI_GROUPID_I_GROUPIDNONE     LVI_GROUPID = -2
)

// [LVITEM] mask.
//
// [LVITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVIF uint32

const (
	LVIF_COLFMT      LVIF = 0x0001_0000
	LVIF_COLUMNS     LVIF = 0x0000_0200
	LVIF_GROUPID     LVIF = 0x0000_0100
	LVIF_IMAGE       LVIF = 0x0000_0002
	LVIF_INDENT      LVIF = 0x0000_0010
	LVIF_NORECOMPUTE LVIF = 0x0000_0800
	LVIF_PARAM       LVIF = 0x0000_0004
	LVIF_STATE       LVIF = 0x0000_0008
	LVIF_TEXT        LVIF = 0x0000_0001
)

// [LVM_GETITEMRECT] portion.
//
// [LVM_GETITEMRECT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemrect
type LVIR uint32

const (
	LVIR_BOUNDS       LVIR = 0
	LVIR_ICON         LVIR = 1
	LVIR_LABEL        LVIR = 2
	LVIR_SELECTBOUNDS LVIR = 3
)

// ListView item [states].
//
// [states]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-item-states
type LVIS uint32

const (
	LVIS_NONE           LVIS = 0
	LVIS_FOCUSED        LVIS = 0x0001
	LVIS_SELECTED       LVIS = 0x0002
	LVIS_CUT            LVIS = 0x0004
	LVIS_DROPHILITED    LVIS = 0x0008
	LVIS_GLOW           LVIS = 0x0010
	LVIS_ACTIVATING     LVIS = 0x0020
	LVIS_OVERLAYMASK    LVIS = 0x0f00
	LVIS_STATEIMAGEMASK LVIS = 0xf000
)

// [NMITEMACTIVATE] uKeyFlags.
//
// [NMITEMACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmitemactivate
type LVKF uint32

const (
	LVKF_ALT     LVKF = 0x0001
	LVKF_CONTROL LVKF = 0x0002
	LVKF_SHIFT   LVKF = 0x0004
)

// [LVHITTESTINFO] flags.
//
// [LVHITTESTINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvhittestinfo
type LVHT uint32

const (
	LVHT_NOWHERE             LVHT = 0x0000_0001
	LVHT_ONITEMICON          LVHT = 0x0000_0002
	LVHT_ONITEMLABEL         LVHT = 0x0000_0004
	LVHT_ONITEMSTATEICON     LVHT = 0x0000_0008
	LVHT_ONITEM              LVHT = LVHT_ONITEMICON | LVHT_ONITEMLABEL | LVHT_ONITEMSTATEICON
	LVHT_ABOVE               LVHT = 0x0000_0008
	LVHT_BELOW               LVHT = 0x0000_0010
	LVHT_TORIGHT             LVHT = 0x0000_0020
	LVHT_TOLEFT              LVHT = 0x0000_0040
	LVHT_EX_GROUP_HEADER     LVHT = 0x1000_0000
	LVHT_EX_GROUP_FOOTER     LVHT = 0x2000_0000
	LVHT_EX_GROUP_COLLAPSE   LVHT = 0x4000_0000
	LVHT_EX_GROUP_BACKGROUND LVHT = 0x8000_0000
	LVHT_EX_GROUP_STATEICON  LVHT = 0x0100_0000
	LVHT_EX_GROUP_SUBSETLINK LVHT = 0x0200_0000
	LVHT_EX_GROUP            LVHT = LVHT_EX_GROUP_BACKGROUND | LVHT_EX_GROUP_COLLAPSE | LVHT_EX_GROUP_FOOTER | LVHT_EX_GROUP_HEADER | LVHT_EX_GROUP_STATEICON | LVHT_EX_GROUP_SUBSETLINK
	LVHT_EX_ONCONTENTS       LVHT = 0x0400_0000
	LVHT_EX_FOOTER           LVHT = 0x0800_0000
)

// [LVM_GETNEXTITEM] item relationship.
//
// [LVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
type LVNI uint32

const (
	LVNI_ALL           LVNI = 0x0000
	LVNI_FOCUSED       LVNI = 0x0001
	LVNI_SELECTED      LVNI = 0x0002
	LVNI_CUT           LVNI = 0x0004
	LVNI_DROPHILITED   LVNI = 0x0008
	LVNI_STATEMASK     LVNI = LVNI_FOCUSED | LVNI_SELECTED | LVNI_CUT | LVNI_DROPHILITED
	LVNI_VISIBLEORDER  LVNI = 0x0010
	LVNI_PREVIOUS      LVNI = 0x0020
	LVNI_VISIBLEONLY   LVNI = 0x0040
	LVNI_SAMEGROUPONLY LVNI = 0x0080
	LVNI_ABOVE         LVNI = 0x0100
	LVNI_BELOW         LVNI = 0x0200
	LVNI_TOLEFT        LVNI = 0x0400
	LVNI_TORIGHT       LVNI = 0x0800
	LVNI_DIRECTIONMASK LVNI = LVNI_ABOVE | LVNI_BELOW | LVNI_TOLEFT | LVNI_TORIGHT
)

// ListView control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-window-styles
type LVS WS

const (
	LVS_ALIGNLEFT       LVS = 0x0800
	LVS_ALIGNMASK       LVS = 0x0c00
	LVS_ALIGNTOP        LVS = 0x0000
	LVS_AUTOARRANGE     LVS = 0x0100
	LVS_EDITLABELS      LVS = 0x0200
	LVS_ICON            LVS = 0x0000
	LVS_LIST            LVS = 0x0003
	LVS_NOCOLUMNHEADER  LVS = 0x4000
	LVS_NOLABELWRAP     LVS = 0x0080
	LVS_NOSCROLL        LVS = 0x2000
	LVS_NOSORTHEADER    LVS = 0x8000
	LVS_OWNERDATA       LVS = 0x1000
	LVS_OWNERDRAWFIXED  LVS = 0x0400
	LVS_REPORT          LVS = 0x0001
	LVS_SHAREIMAGELISTS LVS = 0x0040
	LVS_SHOWSELALWAYS   LVS = 0x0008
	LVS_SINGLESEL       LVS = 0x0004
	LVS_SMALLICON       LVS = 0x0002
	LVS_SORTASCENDING   LVS = 0x0010
	LVS_SORTDESCENDING  LVS = 0x0020
	LVS_TYPEMASK        LVS = 0x0003
	LVS_TYPESTYLEMASK   LVS = 0xfc00
)

// ListView extended control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/extended-list-view-styles
type LVS_EX WS_EX

const (
	LVS_EX_NONE                  LVS_EX = 0
	LVS_EX_AUTOAUTOARRANGE       LVS_EX = 0x0100_0000
	LVS_EX_AUTOCHECKSELECT       LVS_EX = 0x0800_0000
	LVS_EX_AUTOSIZECOLUMNS       LVS_EX = 0x1000_0000
	LVS_EX_BORDERSELECT          LVS_EX = 0x0000_8000
	LVS_EX_CHECKBOXES            LVS_EX = 0x0000_0004
	LVS_EX_COLUMNOVERFLOW        LVS_EX = 0x8000_0000
	LVS_EX_COLUMNSNAPPOINTS      LVS_EX = 0x4000_0000
	LVS_EX_DOUBLEBUFFER          LVS_EX = 0x0001_0000
	LVS_EX_FLATSB                LVS_EX = 0x0000_0100
	LVS_EX_FULLROWSELECT         LVS_EX = 0x0000_0020
	LVS_EX_GRIDLINES             LVS_EX = 0x0000_0001
	LVS_EX_HEADERDRAGDROP        LVS_EX = 0x0000_0010
	LVS_EX_HEADERINALLVIEWS      LVS_EX = 0x0200_0000
	LVS_EX_HIDELABELS            LVS_EX = 0x0002_0000
	LVS_EX_INFOTIP               LVS_EX = 0x0000_0400
	LVS_EX_JUSTIFYCOLUMNS        LVS_EX = 0x0020_0000
	LVS_EX_LABELTIP              LVS_EX = 0x0000_4000
	LVS_EX_MULTIWORKAREAS        LVS_EX = 0x0000_2000
	LVS_EX_ONECLICKACTIVATE      LVS_EX = 0x0000_0040
	LVS_EX_REGIONAL              LVS_EX = 0x0000_0200
	LVS_EX_SIMPLESELECT          LVS_EX = 0x0010_0000
	LVS_EX_SINGLEROW             LVS_EX = 0x0004_0000
	LVS_EX_SNAPTOGRID            LVS_EX = 0x0008_0000
	LVS_EX_SUBITEMIMAGES         LVS_EX = 0x0000_0002
	LVS_EX_TRACKSELECT           LVS_EX = 0x0000_0008
	LVS_EX_TRANSPARENTBKGND      LVS_EX = 0x0040_0000
	LVS_EX_TRANSPARENTSHADOWTEXT LVS_EX = 0x0080_0000
	LVS_EX_TWOCLICKACTIVATE      LVS_EX = 0x0000_0080
	LVS_EX_UNDERLINECOLD         LVS_EX = 0x0000_1000
	LVS_EX_UNDERLINEHOT          LVS_EX = 0x0000_0800
)

// [LVM_GETIMAGELIST] type.
//
// [LVM_GETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getimagelist
type LVSIL uint8

const (
	LVSIL_NORMAL      LVSIL = 0
	LVSIL_SMALL       LVSIL = 1
	LVSIL_STATE       LVSIL = 2
	LVSIL_GROUPHEADER LVSIL = 3
)

// SysLink control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/syslink-control-styles
type LWS WS

const (
	LWS_TRANSPARENT    LWS = 0x0001
	LWS_IGNORERETURN   LWS = 0x0002
	LWS_NOPREFIX       LWS = 0x0004
	LWS_USEVISUALSTYLE LWS = 0x0008
	LWS_USECUSTOMTEXT  LWS = 0x0010
	LWS_RIGHT          LWS = 0x0020
)

// [NMVIEWCHANGE] dwOldView/dwNewView.
//
// [NMVIEWCHANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmviewchange
type MCMV uint32

const (
	MCMV_MONTH   MCMV = 0
	MCMV_YEAR    MCMV = 1
	MCMV_DECADE  MCMV = 2
	MCMV_CENTURY MCMV = 3
)

// MonthCalendar control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/month-calendar-control-styles
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

// ProgressBar control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/progress-bar-control-styles
type PBS WS

const (
	PBS_SMOOTH        PBS = 0x01
	PBS_VERTICAL      PBS = 0x04
	PBS_MARQUEE       PBS = 0x08
	PBS_SMOOTHREVERSE PBS = 0x10
)

// [PBM_SETSTATE] state.
//
// [PBM_SETSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/pbm-setstate
type PBST uint32

const (
	PBST_NORMAL PBST = 0x0001
	PBST_ERROR  PBST = 0x0002
	PBST_PAUSED PBST = 0x0003
)

// StatusBar [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/status-bar-styles
type SBARS WS

const (
	SBARS_SIZEGRIP SBARS = 0x0100 // The status bar control will include a sizing grip at the right end of the status bar. A sizing grip is similar to a sizing border; it is a rectangular area that the user can click and drag to resize the parent window.
	SBARS_TOOLTIPS SBARS = 0x0800 // Use this style to enable tooltips.
)

// [TBN_DROPDOWN] return values.
//
// [TBN_DROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-dropdown
type TBDDRET uint8

const (
	TBDDRET_DEFAULT      TBDDRET = 0
	TBDDRET_NODEFAULT    TBDDRET = 1
	TBDDRET_TREATPRESSED TBDDRET = 2
)

// [TBBUTTONINFO] dwMask.
//
// [TBBUTTONINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tbbuttoninfow
type TBIF uint32

const (
	TBIF_IMAGE   TBIF = 0x0000_0001
	TBIF_TEXT    TBIF = 0x0000_0002
	TBIF_STATE   TBIF = 0x0000_0004
	TBIF_STYLE   TBIF = 0x0000_0008
	TBIF_LPARAM  TBIF = 0x0000_0010
	TBIF_COMMAND TBIF = 0x0000_0020
	TBIF_SIZE    TBIF = 0x0000_0040
	TBIF_BYINDEX TBIF = 0x8000_0000
)

// [NMTBDISPINFO] dwMask.
//
// [NMTBDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbdispinfow
type TBNF uint32

const (
	TBNF_IMAGE      TBNF = 0x1
	TBNF_TEXT       TBNF = 0x2
	TBNF_DI_SETITEM TBNF = 0x1000_0000
)

// [TBN_INITCUSTOMIZE] and [TBN_RESET] return value.
//
// [TBN_INITCUSTOMIZE]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-initcustomize
// [TBN_RESET]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-reset
type TBNRF uint32

const (
	TBNRF_NONE         TBNRF = 0
	TBNRF_HIDEHELP     TBNRF = 0x0000_0001
	TBNRF_ENDCUSTOMIZE TBNRF = 0x0000_0002
)

// Trackbar control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/trackbar-control-styles
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

// Toolbar control [state].
//
// [state]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-button-states
type TBSTATE uint8

const (
	TBSTATE_CHECKED       TBSTATE = 0x01
	TBSTATE_PRESSED       TBSTATE = 0x02
	TBSTATE_ENABLED       TBSTATE = 0x04
	TBSTATE_HIDDEN        TBSTATE = 0x08
	TBSTATE_INDETERMINATE TBSTATE = 0x10
	TBSTATE_WRAP          TBSTATE = 0x20
	TBSTATE_ELLIPSES      TBSTATE = 0x40
	TBSTATE_MARKED        TBSTATE = 0x80
)

// Toolbar control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-control-and-button-styles
type TBSTYLE WS

const (
	TBSTYLE_BUTTON       TBSTYLE = 0x0000
	TBSTYLE_SEP          TBSTYLE = 0x0001
	TBSTYLE_CHECK        TBSTYLE = 0x0002
	TBSTYLE_GROUP        TBSTYLE = 0x0004
	TBSTYLE_CHECKGROUP   TBSTYLE = TBSTYLE_GROUP | TBSTYLE_CHECK
	TBSTYLE_DROPDOWN     TBSTYLE = 0x0008
	TBSTYLE_AUTOSIZE     TBSTYLE = 0x0010
	TBSTYLE_NOPREFIX     TBSTYLE = 0x0020
	TBSTYLE_TOOLTIPS     TBSTYLE = 0x0100
	TBSTYLE_WRAPABLE     TBSTYLE = 0x0200
	TBSTYLE_ALTDRAG      TBSTYLE = 0x0400
	TBSTYLE_FLAT         TBSTYLE = 0x0800
	TBSTYLE_LIST         TBSTYLE = 0x1000
	TBSTYLE_CUSTOMERASE  TBSTYLE = 0x2000
	TBSTYLE_REGISTERDROP TBSTYLE = 0x4000
	TBSTYLE_TRANSPARENT  TBSTYLE = 0x8000
)

// Toolbar control extended [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/toolbar-extended-styles
type TBSTYLE_EX uint32

const (
	TBSTYLE_EX_NONE               TBSTYLE_EX = 0
	TBSTYLE_EX_DRAWDDARROWS       TBSTYLE_EX = 0x0000_0001
	TBSTYLE_EX_MIXEDBUTTONS       TBSTYLE_EX = 0x0000_0008
	TBSTYLE_EX_HIDECLIPPEDBUTTONS TBSTYLE_EX = 0x0000_0010
	TBSTYLE_EX_MULTICOLUMN        TBSTYLE_EX = 0x0000_0002
	TBSTYLE_EX_VERTICAL           TBSTYLE_EX = 0x0000_0004
	TBSTYLE_EX_DOUBLEBUFFER       TBSTYLE_EX = 0x0000_0080
)

// [TaskDialog] pszIcon. Originally with TD prefix and ICON suffix.
//
// [TaskDialog]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
type TD_ICON uint16

const (
	TD_ICON_WARNING     TD_ICON = 0xffff
	TD_ICON_ERROR       TD_ICON = 0xfffe
	TD_ICON_INFORMATION TD_ICON = 0xfffd
	TD_ICON_SHIELD      TD_ICON = 0xfffc
)

// [TaskDialog] dwCommonButtons. Originally has BUTTON suffix.
//
// [TaskDialog]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
type TDCBF int32

const (
	TDCBF_OK     TDCBF = 0x0001
	TDCBF_YES    TDCBF = 0x0002
	TDCBF_NO     TDCBF = 0x0004
	TDCBF_CANCEL TDCBF = 0x0008
	TDCBF_RETRY  TDCBF = 0x0010
	TDCBF_CLOSE  TDCBF = 0x0020
)

// [TASKDIALOGCONFIG] dwFlags.
//
// [TASKDIALOGCONFIG]: https://learn.microsoft.com/en-us/windows/win32/api/Commctrl/ns-commctrl-taskdialogconfig
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
	TDF_NO_SET_FOREGROUND           TDF = 0x0001_0000
	TDF_SIZE_TO_CONTENT             TDF = 0x0100_0000
)

// [EDITBALLOONTIP] ttiIcon.
//
// [EDITBALLOONTIP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-editballoontip
type TTI int32

const (
	TTI_ERROR         TTI = 3
	TTI_INFO          TTI = 1
	TTI_NONE          TTI = 0
	TTI_WARNING       TTI = 2
	TTI_INFO_LARGE    TTI = 4
	TTI_WARNING_LARGE TTI = 5
	TTI_ERROR_LARGE   TTI = 6
)

// [TVM_EXPAND] action flag.
//
// [TVM_EXPAND]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-expand
type TVE uint32

const (
	TVE_COLLAPSE      TVE = 0x0001
	TVE_EXPAND        TVE = 0x0002
	TVE_TOGGLE        TVE = 0x0003
	TVE_EXPANDPARTIAL TVE = 0x4000
	TVE_COLLAPSERESET TVE = 0x8000
)

// [TVM_GETNEXTITEM] item to retrieve.
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
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

// [TVITEMTEX] cChildren.
//
// [TVITEMTEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVI_CHILDREN int32

const (
	TVI_CHILDREN_ZERO     TVI_CHILDREN = 0
	TVI_CHILDREN_ONE      TVI_CHILDREN = 1
	TVI_CHILDREN_CALLBACK TVI_CHILDREN = -1
	TVI_CHILDREN_AUTO     TVI_CHILDREN = -2
)

// [TVITEMTEX] mask.
//
// [TVITEMTEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
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

// [TVITEMTEX] state.
//
// [TVITEMTEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
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

// [TVITEMTEX] uStateEx.
//
// [TVITEMTEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVIS_EX uint32

const (
	TVIS_EX_FLAT     TVIS_EX = 0x0001
	TVIS_EX_DISABLED TVIS_EX = 0x0002
	TVIS_EX_ALL      TVIS_EX = 0x0002
)

// [TVN_SINGLEEXPAND] return value.
//
// [TVN_SINGLEEXPAND]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-singleexpand
type TVNRET uintptr

const (
	TVNRET_DEFAULT TVNRET = 0
	TVNRET_SKIPOLD TVNRET = 1
	TVNRET_SKIPNEW TVNRET = 2
)

// TreeView control [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-control-window-styles
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

// TreeView control extended [styles].
//
// [styles]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-control-window-extended-styles
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
