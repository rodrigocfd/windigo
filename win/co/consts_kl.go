package co

// LITEM mask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LIF uint32

const (
	LIF_ITEMINDEX LIF = 0x00000001
	LIF_STATE     LIF = 0x00000002
	LIF_ITEMID    LIF = 0x00000004
	LIF_URL       LIF = 0x00000008
)

// LITEM state.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LIS uint32

const (
	LIS_FOCUSED       LIS = 0x00000001
	LIS_ENABLED       LIS = 0x00000002
	LIS_VISITED       LIS = 0x00000004
	LIS_HOTTRACK      LIS = 0x00000008
	LIS_DEFAULTCOLORS LIS = 0x00000010
)

// LoadImage fuLoad.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
type LR uint32

const (
	LR_DEFAULTCOLOR     LR = 0x00000000
	LR_MONOCHROME       LR = 0x00000001
	LR_COLOR            LR = 0x00000002
	LR_COPYRETURNORG    LR = 0x00000004
	LR_COPYDELETEORG    LR = 0x00000008
	LR_LOADFROMFILE     LR = 0x00000010
	LR_LOADTRANSPARENT  LR = 0x00000020
	LR_DEFAULTSIZE      LR = 0x00000040
	LR_VGACOLOR         LR = 0x00000080
	LR_LOADMAP3DCOLORS  LR = 0x00001000
	LR_CREATEDIBSECTION LR = 0x00002000
	LR_COPYFROMRESOURCE LR = 0x00004000
	LR_SHARED           LR = 0x00008000
)

// NMLVCUSTOMDRAW dwItemType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
type LVCDI uint32

const (
	LVCDI_ITEM     LVCDI = 0x00000000
	LVCDI_GROUP    LVCDI = 0x00000001
	LVCDI_TEMSLIST LVCDI = 0x00000002
)

// LVCOLUMN mask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
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

// LVCOLUMN fmt.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
type LVCFMT_C int32

const (
	LVCFMT_C_LEFT            LVCFMT_C = 0x0000
	LVCFMT_C_RIGHT           LVCFMT_C = 0x0001
	LVCFMT_C_CENTER          LVCFMT_C = 0x0002
	LVCFMT_C_JUSTIFYMASK     LVCFMT_C = 0x0003
	LVCFMT_C_IMAGE           LVCFMT_C = 0x0800
	LVCFMT_C_BITMAP_ON_RIGHT LVCFMT_C = 0x1000
	LVCFMT_C_COL_HAS_IMAGES  LVCFMT_C = 0x8000
	LVCFMT_C_FIXED_WIDTH     LVCFMT_C = 0x00100
	LVCFMT_C_NO_DPI_SCALE    LVCFMT_C = 0x40000
	LVCFMT_C_FIXED_RATIO     LVCFMT_C = 0x80000
	LVCFMT_C_SPLITBUTTON     LVCFMT_C = 0x1000000
)

// LVITEM piColFmt.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVCFMT_I int32

const (
	LVCFMT_I_LINE_BREAK         LVCFMT_I = 0x100000
	LVCFMT_I_FILL               LVCFMT_I = 0x200000
	LVCFMT_I_WRAP               LVCFMT_I = 0x400000
	LVCFMT_I_NO_TITLE           LVCFMT_I = 0x800000
	LVCFMT_I_TILE_PLACEMENTMASK LVCFMT_I = LVCFMT_I_LINE_BREAK | LVCFMT_I_FILL
)

// LVFINDINFO flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvfindinfow
type LVFI uint32

const (
	LVFI_PARAM     LVFI = 0x0001
	LVFI_STRING    LVFI = 0x0002
	LVFI_SUBSTRING LVFI = 0x0004
	LVFI_PARTIAL   LVFI = 0x0008
	LVFI_WRAP      LVFI = 0x0020
	LVFI_NEARESTXY LVFI = 0x0040
)

// NMLVCUSTOMDRAW uAlign.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
type LVGA_HEADER uint32

const (
	LVGA_HEADER_LEFT   LVGA_HEADER = 0x00000001
	LVGA_HEADER_CENTER LVGA_HEADER = 0x00000002
	LVGA_HEADER_RIGHT  LVGA_HEADER = 0x00000004
)

// NMLVGETINFOTIP dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvgetinfotipw
type LVGIT uint32

const (
	LVGIT_ZERO     LVGIT = 0x0000
	LVGIT_UNFOLDED LVGIT = 0x0001
)

// LVITEM iGroupId.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVI_GROUPID int32

const (
	LVI_GROUPID_I_GROUPIDCALLBACK LVI_GROUPID = -1
	LVI_GROUPID_I_GROUPIDNONE     LVI_GROUPID = -2
)

// LVITEM mask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVIF uint32

const (
	LVIF_COLFMT      LVIF = 0x00010000
	LVIF_COLUMNS     LVIF = 0x00000200
	LVIF_GROUPID     LVIF = 0x00000100
	LVIF_IMAGE       LVIF = 0x00000002
	LVIF_INDENT      LVIF = 0x00000010
	LVIF_NORECOMPUTE LVIF = 0x00000800
	LVIF_PARAM       LVIF = 0x00000004
	LVIF_STATE       LVIF = 0x00000008
	LVIF_TEXT        LVIF = 0x00000001
)

// LVM_GETITEMRECT portion.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemrect
type LVIR uint32

const (
	LVIR_BOUNDS       LVIR = 0
	LVIR_ICON         LVIR = 1
	LVIR_LABEL        LVIR = 2
	LVIR_SELECTBOUNDS LVIR = 3
)

// ListView item states.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/list-view-item-states
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

// NMITEMACTIVATE uKeyFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmitemactivate
type LVKF uint32

const (
	LVKF_ALT     LVKF = 0x0001
	LVKF_CONTROL LVKF = 0x0002
	LVKF_SHIFT   LVKF = 0x0004
)

// ListView control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-list-view-control-reference-messages
const (
	_LVM_FIRST WM = 0x1000

	LVM_GETBKCOLOR               WM = _LVM_FIRST + 0
	LVM_SETBKCOLOR               WM = _LVM_FIRST + 1
	LVM_GETIMAGELIST             WM = _LVM_FIRST + 2
	LVM_SETIMAGELIST             WM = _LVM_FIRST + 3
	LVM_GETITEMCOUNT             WM = _LVM_FIRST + 4
	LVM_DELETEITEM               WM = _LVM_FIRST + 8
	LVM_DELETEALLITEMS           WM = _LVM_FIRST + 9
	LVM_GETCALLBACKMASK          WM = _LVM_FIRST + 10
	LVM_SETCALLBACKMASK          WM = _LVM_FIRST + 11
	LVM_GETNEXTITEM              WM = _LVM_FIRST + 12
	LVM_GETITEMRECT              WM = _LVM_FIRST + 14
	LVM_SETITEMPOSITION          WM = _LVM_FIRST + 15
	LVM_GETITEMPOSITION          WM = _LVM_FIRST + 16
	LVM_HITTEST                  WM = _LVM_FIRST + 18
	LVM_ENSUREVISIBLE            WM = _LVM_FIRST + 19
	LVM_SCROLL                   WM = _LVM_FIRST + 20
	LVM_REDRAWITEMS              WM = _LVM_FIRST + 21
	LVM_ARRANGE                  WM = _LVM_FIRST + 22
	LVM_GETEDITCONTROL           WM = _LVM_FIRST + 24
	LVM_DELETECOLUMN             WM = _LVM_FIRST + 28
	LVM_GETCOLUMNWIDTH           WM = _LVM_FIRST + 29
	LVM_SETCOLUMNWIDTH           WM = _LVM_FIRST + 30
	LVM_GETHEADER                WM = _LVM_FIRST + 31
	LVM_CREATEDRAGIMAGE          WM = _LVM_FIRST + 33
	LVM_GETVIEWRECT              WM = _LVM_FIRST + 34
	LVM_GETTEXTCOLOR             WM = _LVM_FIRST + 35
	LVM_SETTEXTCOLOR             WM = _LVM_FIRST + 36
	LVM_GETTEXTBKCOLOR           WM = _LVM_FIRST + 37
	LVM_SETTEXTBKCOLOR           WM = _LVM_FIRST + 38
	LVM_GETTOPINDEX              WM = _LVM_FIRST + 39
	LVM_GETCOUNTPERPAGE          WM = _LVM_FIRST + 40
	LVM_GETORIGIN                WM = _LVM_FIRST + 41
	LVM_UPDATE                   WM = _LVM_FIRST + 42
	LVM_SETITEMSTATE             WM = _LVM_FIRST + 43
	LVM_GETITEMSTATE             WM = _LVM_FIRST + 44
	LVM_SETITEMCOUNT             WM = _LVM_FIRST + 47
	LVM_SORTITEMS                WM = _LVM_FIRST + 48
	LVM_SETITEMPOSITION32        WM = _LVM_FIRST + 49
	LVM_GETSELECTEDCOUNT         WM = _LVM_FIRST + 50
	LVM_GETITEMSPACING           WM = _LVM_FIRST + 51
	LVM_SETICONSPACING           WM = _LVM_FIRST + 53
	LVM_SETEXTENDEDLISTVIEWSTYLE WM = _LVM_FIRST + 54
	LVM_GETEXTENDEDLISTVIEWSTYLE WM = _LVM_FIRST + 55
	LVM_GETSUBITEMRECT           WM = _LVM_FIRST + 56
	LVM_SUBITEMHITTEST           WM = _LVM_FIRST + 57
	LVM_SETCOLUMNORDERARRAY      WM = _LVM_FIRST + 58
	LVM_GETCOLUMNORDERARRAY      WM = _LVM_FIRST + 59
	LVM_SETHOTITEM               WM = _LVM_FIRST + 60
	LVM_GETHOTITEM               WM = _LVM_FIRST + 61
	LVM_SETHOTCURSOR             WM = _LVM_FIRST + 62
	LVM_GETHOTCURSOR             WM = _LVM_FIRST + 63
	LVM_APPROXIMATEVIEWRECT      WM = _LVM_FIRST + 64
	LVM_SETWORKAREAS             WM = _LVM_FIRST + 65
	LVM_GETSELECTIONMARK         WM = _LVM_FIRST + 66
	LVM_SETSELECTIONMARK         WM = _LVM_FIRST + 67
	LVM_GETWORKAREAS             WM = _LVM_FIRST + 70
	LVM_SETHOVERTIME             WM = _LVM_FIRST + 71
	LVM_GETHOVERTIME             WM = _LVM_FIRST + 72
	LVM_GETNUMBEROFWORKAREAS     WM = _LVM_FIRST + 73
	LVM_SETTOOLTIPS              WM = _LVM_FIRST + 74
	LVM_GETITEM                  WM = _LVM_FIRST + 75
	LVM_SETITEM                  WM = _LVM_FIRST + 76
	LVM_INSERTITEM               WM = _LVM_FIRST + 77
	LVM_GETTOOLTIPS              WM = _LVM_FIRST + 78
	LVM_SORTITEMSEX              WM = _LVM_FIRST + 81
	LVM_FINDITEM                 WM = _LVM_FIRST + 83
	LVM_GETSTRINGWIDTH           WM = _LVM_FIRST + 87
	LVM_GETGROUPSTATE            WM = _LVM_FIRST + 92
	LVM_GETFOCUSEDGROUP          WM = _LVM_FIRST + 93
	LVM_GETCOLUMN                WM = _LVM_FIRST + 95
	LVM_SETCOLUMN                WM = _LVM_FIRST + 96
	LVM_INSERTCOLUMN             WM = _LVM_FIRST + 97
	LVM_GETGROUPRECT             WM = _LVM_FIRST + 98
	LVM_GETITEMTEXT              WM = _LVM_FIRST + 115
	LVM_SETITEMTEXT              WM = _LVM_FIRST + 116
	LVM_GETISEARCHSTRING         WM = _LVM_FIRST + 117
	LVM_EDITLABEL                WM = _LVM_FIRST + 118
	LVM_SETBKIMAGE               WM = _LVM_FIRST + 138
	LVM_GETBKIMAGE               WM = _LVM_FIRST + 139
	LVM_SETSELECTEDCOLUMN        WM = _LVM_FIRST + 140
	LVM_SETVIEW                  WM = _LVM_FIRST + 142
	LVM_GETVIEW                  WM = _LVM_FIRST + 143
	LVM_INSERTGROUP              WM = _LVM_FIRST + 145
	LVM_SETGROUPINFO             WM = _LVM_FIRST + 147
	LVM_GETGROUPINFO             WM = _LVM_FIRST + 149
	LVM_REMOVEGROUP              WM = _LVM_FIRST + 150
	LVM_MOVEGROUP                WM = _LVM_FIRST + 151
	LVM_GETGROUPCOUNT            WM = _LVM_FIRST + 152
	LVM_GETGROUPINFOBYINDEX      WM = _LVM_FIRST + 153
	LVM_MOVEITEMTOGROUP          WM = _LVM_FIRST + 154
	LVM_SETGROUPMETRICS          WM = _LVM_FIRST + 155
	LVM_GETGROUPMETRICS          WM = _LVM_FIRST + 156
	LVM_ENABLEGROUPVIEW          WM = _LVM_FIRST + 157
	LVM_SORTGROUPS               WM = _LVM_FIRST + 158
	LVM_INSERTGROUPSORTED        WM = _LVM_FIRST + 159
	LVM_REMOVEALLGROUPS          WM = _LVM_FIRST + 160
	LVM_HASGROUP                 WM = _LVM_FIRST + 161
	LVM_SETTILEVIEWINFO          WM = _LVM_FIRST + 162
	LVM_GETTILEVIEWINFO          WM = _LVM_FIRST + 163
	LVM_SETTILEINFO              WM = _LVM_FIRST + 164
	LVM_GETTILEINFO              WM = _LVM_FIRST + 165
	LVM_SETINSERTMARK            WM = _LVM_FIRST + 166
	LVM_GETINSERTMARK            WM = _LVM_FIRST + 167
	LVM_INSERTMARKHITTEST        WM = _LVM_FIRST + 168
	LVM_GETINSERTMARKRECT        WM = _LVM_FIRST + 169
	LVM_SETINSERTMARKCOLOR       WM = _LVM_FIRST + 170
	LVM_GETINSERTMARKCOLOR       WM = _LVM_FIRST + 171
	LVM_SETINFOTIP               WM = _LVM_FIRST + 173
	LVM_GETSELECTEDCOLUMN        WM = _LVM_FIRST + 174
	LVM_ISGROUPVIEWENABLED       WM = _LVM_FIRST + 175
	LVM_GETOUTLINECOLOR          WM = _LVM_FIRST + 176
	LVM_SETOUTLINECOLOR          WM = _LVM_FIRST + 177
	LVM_CANCELEDITLABEL          WM = _LVM_FIRST + 179
	LVM_MAPINDEXTOID             WM = _LVM_FIRST + 180
	LVM_MAPIDTOINDEX             WM = _LVM_FIRST + 181
	LVM_ISITEMVISIBLE            WM = _LVM_FIRST + 182
	LVM_GETEMPTYTEXT             WM = _LVM_FIRST + 204
	LVM_GETFOOTERRECT            WM = _LVM_FIRST + 205
	LVM_GETFOOTERINFO            WM = _LVM_FIRST + 206
	LVM_GETFOOTERITEMRECT        WM = _LVM_FIRST + 207
	LVM_GETFOOTERITEM            WM = _LVM_FIRST + 208
	LVM_GETITEMINDEXRECT         WM = _LVM_FIRST + 209
	LVM_SETITEMINDEXSTATE        WM = _LVM_FIRST + 210
	LVM_GETNEXTITEMINDEX         WM = _LVM_FIRST + 211
)

// ListView control notifications, sent via WM_NOTIFY.
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

// LVHITTESTINFO flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvhittestinfo
type LVHT uint32

const (
	LVHT_NOWHERE             LVHT = 0x00000001
	LVHT_ONITEMICON          LVHT = 0x00000002
	LVHT_ONITEMLABEL         LVHT = 0x00000004
	LVHT_ONITEMSTATEICON     LVHT = 0x00000008
	LVHT_ONITEM              LVHT = LVHT_ONITEMICON | LVHT_ONITEMLABEL | LVHT_ONITEMSTATEICON
	LVHT_ABOVE               LVHT = 0x00000008
	LVHT_BELOW               LVHT = 0x00000010
	LVHT_TORIGHT             LVHT = 0x00000020
	LVHT_TOLEFT              LVHT = 0x00000040
	LVHT_EX_GROUP_HEADER     LVHT = 0x10000000
	LVHT_EX_GROUP_FOOTER     LVHT = 0x20000000
	LVHT_EX_GROUP_COLLAPSE   LVHT = 0x40000000
	LVHT_EX_GROUP_BACKGROUND LVHT = 0x80000000
	LVHT_EX_GROUP_STATEICON  LVHT = 0x01000000
	LVHT_EX_GROUP_SUBSETLINK LVHT = 0x02000000
	LVHT_EX_GROUP            LVHT = LVHT_EX_GROUP_BACKGROUND | LVHT_EX_GROUP_COLLAPSE | LVHT_EX_GROUP_FOOTER | LVHT_EX_GROUP_HEADER | LVHT_EX_GROUP_STATEICON | LVHT_EX_GROUP_SUBSETLINK
	LVHT_EX_ONCONTENTS       LVHT = 0x04000000
	LVHT_EX_FOOTER           LVHT = 0x08000000
)

// LVM_GETNEXTITEM item relationship.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
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

// ListView control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/list-view-window-styles
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

// ListView extended control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/extended-list-view-styles
type LVS_EX WS_EX

const (
	LVS_EX_NONE                  LVS_EX = 0
	LVS_EX_AUTOAUTOARRANGE       LVS_EX = 0x01000000
	LVS_EX_AUTOCHECKSELECT       LVS_EX = 0x08000000
	LVS_EX_AUTOSIZECOLUMNS       LVS_EX = 0x10000000
	LVS_EX_BORDERSELECT          LVS_EX = 0x00008000
	LVS_EX_CHECKBOXES            LVS_EX = 0x00000004
	LVS_EX_COLUMNOVERFLOW        LVS_EX = 0x80000000
	LVS_EX_COLUMNSNAPPOINTS      LVS_EX = 0x40000000
	LVS_EX_DOUBLEBUFFER          LVS_EX = 0x00010000
	LVS_EX_FLATSB                LVS_EX = 0x00000100
	LVS_EX_FULLROWSELECT         LVS_EX = 0x00000020
	LVS_EX_GRIDLINES             LVS_EX = 0x00000001
	LVS_EX_HEADERDRAGDROP        LVS_EX = 0x00000010
	LVS_EX_HEADERINALLVIEWS      LVS_EX = 0x02000000
	LVS_EX_HIDELABELS            LVS_EX = 0x00020000
	LVS_EX_INFOTIP               LVS_EX = 0x00000400
	LVS_EX_JUSTIFYCOLUMNS        LVS_EX = 0x00200000
	LVS_EX_LABELTIP              LVS_EX = 0x00004000
	LVS_EX_MULTIWORKAREAS        LVS_EX = 0x00002000
	LVS_EX_ONECLICKACTIVATE      LVS_EX = 0x00000040
	LVS_EX_REGIONAL              LVS_EX = 0x00000200
	LVS_EX_SIMPLESELECT          LVS_EX = 0x00100000
	LVS_EX_SINGLEROW             LVS_EX = 0x00040000
	LVS_EX_SNAPTOGRID            LVS_EX = 0x00080000
	LVS_EX_SUBITEMIMAGES         LVS_EX = 0x00000002
	LVS_EX_TRACKSELECT           LVS_EX = 0x00000008
	LVS_EX_TRANSPARENTBKGND      LVS_EX = 0x00400000
	LVS_EX_TRANSPARENTSHADOWTEXT LVS_EX = 0x00800000
	LVS_EX_TWOCLICKACTIVATE      LVS_EX = 0x00000080
	LVS_EX_UNDERLINECOLD         LVS_EX = 0x00001000
	LVS_EX_UNDERLINEHOT          LVS_EX = 0x00000800
)

// SysLink control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/syslink-control-styles
type LWS WS

const (
	LWS_TRANSPARENT    LWS = 0x0001
	LWS_IGNORERETURN   LWS = 0x0002
	LWS_NOPREFIX       LWS = 0x0004
	LWS_USEVISUALSTYLE LWS = 0x0008
	LWS_USECUSTOMTEXT  LWS = 0x0010
	LWS_RIGHT          LWS = 0x0020
)
