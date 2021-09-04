package co

// Registry key security and access rights
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/sysinfo/registry-key-security-and-access-rights
type KEY uint32

const (
	KEY_QUERY_VALUE        KEY = 0x0001
	KEY_SET_VALUE          KEY = 0x0002
	KEY_CREATE_SUB_KEY     KEY = 0x0004
	KEY_ENUMERATE_SUB_KEYS KEY = 0x0008
	KEY_NOTIFY             KEY = 0x0010
	KEY_CREATE_LINK        KEY = 0x0020
	KEY_WOW64_32KEY        KEY = 0x0200
	KEY_WOW64_64KEY        KEY = 0x0100
	KEY_WOW64_RES          KEY = 0x0300

	KEY_READ       KEY = (KEY(ACCESS_RIGHTS_STANDARD_READ) | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) & ^KEY(ACCESS_RIGHTS_SYNCHRONIZE)
	KEY_WRITE      KEY = (KEY(ACCESS_RIGHTS_STANDARD_WRITE) | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) & ^KEY(ACCESS_RIGHTS_SYNCHRONIZE)
	KEY_EXECUTE    KEY = KEY_READ & ^KEY(ACCESS_RIGHTS_SYNCHRONIZE)
	KEY_ALL_ACCESS KEY = (KEY(ACCESS_RIGHTS_STANDARD_ALL) | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) & ^KEY(ACCESS_RIGHTS_SYNCHRONIZE)
)

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

// LVM_GETVIEW return value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getview
type LV_VIEW uint32

const (
	LV_VIEW_ICON      LV_VIEW = 0x0000
	LV_VIEW_DETAILS   LV_VIEW = 0x0001
	LV_VIEW_SMALLICON LV_VIEW = 0x0002
	LV_VIEW_LIST      LV_VIEW = 0x0003
	LV_VIEW_TILE      LV_VIEW = 0x0004
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

// LVM_GETIMAGELIST type.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getimagelist
type LVSIL uint8

const (
	LVSIL_NORMAL      LVSIL = 0
	LVSIL_SMALL       LVSIL = 1
	LVSIL_STATE       LVSIL = 2
	LVSIL_GROUPHEADER LVSIL = 3
)

// LockSetForegroundWindow() uLockCode.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
type LSFW uint32

const (
	LSFW_LOCK   LSFW = 1
	LSFW_UNLOCK LSFW = 2
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
