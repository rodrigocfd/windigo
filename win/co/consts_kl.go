package co

// Registry key security and access rights
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/sysinfo/registry-key-security-and-access-rights
type KEY uint32

const (
	// Required to query the values of a registry key.
	KEY_QUERY_VALUE KEY = 0x0001
	// Required to create, delete, or set a registry value.
	KEY_SET_VALUE KEY = 0x0002
	// Required to create a subkey of a registry key.
	KEY_CREATE_SUB_KEY KEY = 0x0004
	// Required to enumerate the subkeys of a registry key.
	KEY_ENUMERATE_SUB_KEYS KEY = 0x0008
	// Required to request change notifications for a registry key or for
	// subkeys of a registry key.
	KEY_NOTIFY KEY = 0x0010
	// Reserved for system use.
	KEY_CREATE_LINK KEY = 0x0020
	// Indicates that an application on 64-bit Windows should operate on the
	// 32-bit registry view.
	KEY_WOW64_32KEY KEY = 0x0200
	// Indicates that an application on 64-bit Windows should operate on the
	// 64-bit registry view.
	KEY_WOW64_64KEY KEY = 0x0100
	// Undocumented flag.
	KEY_WOW64_RES KEY = 0x0300

	// Combines the STANDARD_RIGHTS_READ, KEY_QUERY_VALUE,
	// KEY_ENUMERATE_SUB_KEYS, and KEY_NOTIFY values.
	KEY_READ KEY = (KEY(STANDARD_RIGHTS_READ) | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Combines the STANDARD_RIGHTS_WRITE, KEY_SET_VALUE, and KEY_CREATE_SUB_KEY
	// access rights.
	KEY_WRITE KEY = (KEY(STANDARD_RIGHTS_WRITE) | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Equivalent to KEY_READ.
	KEY_EXECUTE KEY = KEY_READ & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Combines the STANDARD_RIGHTS_REQUIRED, KEY_QUERY_VALUE, KEY_SET_VALUE,
	// KEY_CREATE_SUB_KEY, KEY_ENUMERATE_SUB_KEYS, KEY_NOTIFY, and
	// KEY_CREATE_LINK access rights.
	KEY_ALL_ACCESS KEY = (KEY(STANDARD_RIGHTS_ALL) | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
)

// Language identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/language-identifier-constants-and-strings
type LANG uint16

const (
	LANG_NEUTRAL             LANG = 0x00
	LANG_INVARIANT           LANG = 0x7f
	LANG_AFRIKAANS           LANG = 0x36
	LANG_ALBANIAN            LANG = 0x1c
	LANG_ALSATIAN            LANG = 0x84
	LANG_AMHARIC             LANG = 0x5e
	LANG_ARABIC              LANG = 0x01
	LANG_ARMENIAN            LANG = 0x2b
	LANG_ASSAMESE            LANG = 0x4d
	LANG_AZERI               LANG = 0x2c
	LANG_AZERBAIJANI         LANG = 0x2c
	LANG_BANGLA              LANG = 0x45
	LANG_BASHKIR             LANG = 0x6d
	LANG_BASQUE              LANG = 0x2d
	LANG_BELARUSIAN          LANG = 0x23
	LANG_BENGALI             LANG = 0x45
	LANG_BRETON              LANG = 0x7e
	LANG_BOSNIAN             LANG = 0x1a
	LANG_BOSNIAN_NEUTRAL     LANG = 0x781a
	LANG_BULGARIAN           LANG = 0x02
	LANG_CATALAN             LANG = 0x03
	LANG_CENTRAL_KURDISH     LANG = 0x92
	LANG_CHEROKEE            LANG = 0x5c
	LANG_CHINESE             LANG = 0x04
	LANG_CHINESE_SIMPLIFIED  LANG = 0x04
	LANG_CHINESE_TRADITIONAL LANG = 0x7c04
	LANG_CORSICAN            LANG = 0x83
	LANG_CROATIAN            LANG = 0x1a
	LANG_CZECH               LANG = 0x05
	LANG_DANISH              LANG = 0x06
	LANG_DARI                LANG = 0x8c
	LANG_DIVEHI              LANG = 0x65
	LANG_DUTCH               LANG = 0x13
	LANG_ENGLISH             LANG = 0x09
	LANG_ESTONIAN            LANG = 0x25
	LANG_FAEROESE            LANG = 0x38
	LANG_FARSI               LANG = 0x29
	LANG_FILIPINO            LANG = 0x64
	LANG_FINNISH             LANG = 0x0b
	LANG_FRENCH              LANG = 0x0c
	LANG_FRISIAN             LANG = 0x62
	LANG_FULAH               LANG = 0x67
	LANG_GALICIAN            LANG = 0x56
	LANG_GEORGIAN            LANG = 0x37
	LANG_GERMAN              LANG = 0x07
	LANG_GREEK               LANG = 0x08
	LANG_GREENLANDIC         LANG = 0x6f
	LANG_GUJARATI            LANG = 0x47
	LANG_HAUSA               LANG = 0x68
	LANG_HAWAIIAN            LANG = 0x75
	LANG_HEBREW              LANG = 0x0d
	LANG_HINDI               LANG = 0x39
	LANG_HUNGARIAN           LANG = 0x0e
	LANG_ICELANDIC           LANG = 0x0f
	LANG_IGBO                LANG = 0x70
	LANG_INDONESIAN          LANG = 0x21
	LANG_INUKTITUT           LANG = 0x5d
	LANG_IRISH               LANG = 0x3c
	LANG_ITALIAN             LANG = 0x10
	LANG_JAPANESE            LANG = 0x11
	LANG_KANNADA             LANG = 0x4b
	LANG_KASHMIRI            LANG = 0x60
	LANG_KAZAK               LANG = 0x3f
	LANG_KHMER               LANG = 0x53
	LANG_KICHE               LANG = 0x86
	LANG_KINYARWANDA         LANG = 0x87
	LANG_KONKANI             LANG = 0x57
	LANG_KOREAN              LANG = 0x12
	LANG_KYRGYZ              LANG = 0x40
	LANG_LAO                 LANG = 0x54
	LANG_LATVIAN             LANG = 0x26
	LANG_LITHUANIAN          LANG = 0x27
	LANG_LOWER_SORBIAN       LANG = 0x2e
	LANG_LUXEMBOURGISH       LANG = 0x6e
	LANG_MACEDONIAN          LANG = 0x2f
	LANG_MALAY               LANG = 0x3e
	LANG_MALAYALAM           LANG = 0x4c
	LANG_MALTESE             LANG = 0x3a
	LANG_MANIPURI            LANG = 0x58
	LANG_MAORI               LANG = 0x81
	LANG_MAPUDUNGUN          LANG = 0x7a
	LANG_MARATHI             LANG = 0x4e
	LANG_MOHAWK              LANG = 0x7c
	LANG_MONGOLIAN           LANG = 0x50
	LANG_NEPALI              LANG = 0x61
	LANG_NORWEGIAN           LANG = 0x14
	LANG_OCCITAN             LANG = 0x82
	LANG_ODIA                LANG = 0x48
	LANG_ORIYA               LANG = 0x48
	LANG_PASHTO              LANG = 0x63
	LANG_PERSIAN             LANG = 0x29
	LANG_POLISH              LANG = 0x15
	LANG_PORTUGUESE          LANG = 0x16
	LANG_PULAR               LANG = 0x67
	LANG_PUNJABI             LANG = 0x46
	LANG_QUECHUA             LANG = 0x6b
	LANG_ROMANIAN            LANG = 0x18
	LANG_ROMANSH             LANG = 0x17
	LANG_RUSSIAN             LANG = 0x19
	LANG_SAKHA               LANG = 0x85
	LANG_SAMI                LANG = 0x3b
	LANG_SANSKRIT            LANG = 0x4f
	LANG_SCOTTISH_GAELIC     LANG = 0x91
	LANG_SERBIAN             LANG = 0x1a
	LANG_SERBIAN_NEUTRAL     LANG = 0x7c1a
	LANG_SINDHI              LANG = 0x59
	LANG_SINHALESE           LANG = 0x5b
	LANG_SLOVAK              LANG = 0x1b
	LANG_SLOVENIAN           LANG = 0x24
	LANG_SOTHO               LANG = 0x6c
	LANG_SPANISH             LANG = 0x0a
	LANG_SWAHILI             LANG = 0x41
	LANG_SWEDISH             LANG = 0x1d
	LANG_SYRIAC              LANG = 0x5a
	LANG_TAJIK               LANG = 0x28
	LANG_TAMAZIGHT           LANG = 0x5f
	LANG_TAMIL               LANG = 0x49
	LANG_TATAR               LANG = 0x44
	LANG_TELUGU              LANG = 0x4a
	LANG_THAI                LANG = 0x1e
	LANG_TIBETAN             LANG = 0x51
	LANG_TIGRIGNA            LANG = 0x73
	LANG_TIGRINYA            LANG = 0x73
	LANG_TSWANA              LANG = 0x32
	LANG_TURKISH             LANG = 0x1f
	LANG_TURKMEN             LANG = 0x42
	LANG_UIGHUR              LANG = 0x80
	LANG_UKRAINIAN           LANG = 0x22
	LANG_UPPER_SORBIAN       LANG = 0x2e
	LANG_URDU                LANG = 0x20
	LANG_UZBEK               LANG = 0x43
	LANG_VALENCIAN           LANG = 0x03
	LANG_VIETNAMESE          LANG = 0x2a
	LANG_WELSH               LANG = 0x52
	LANG_WOLOF               LANG = 0x88
	LANG_XHOSA               LANG = 0x34
	LANG_YAKUT               LANG = 0x85
	LANG_YI                  LANG = 0x78
	LANG_YORUBA              LANG = 0x6a
	LANG_ZULU                LANG = 0x35
)

// LITEM mask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LIF uint32

const (
	LIF_ITEMINDEX LIF = 0x0000_0001
	LIF_STATE     LIF = 0x0000_0002
	LIF_ITEMID    LIF = 0x0000_0004
	LIF_URL       LIF = 0x0000_0008
)

// LITEM state.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LIS uint32

const (
	LIS_FOCUSED       LIS = 0x0000_0001
	LIS_ENABLED       LIS = 0x0000_0002
	LIS_VISITED       LIS = 0x0000_0004
	LIS_HOTTRACK      LIS = 0x0000_0008
	LIS_DEFAULTCOLORS LIS = 0x0000_0010
)

// LoadImage fuLoad.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
type LR uint32

const (
	LR_DEFAULTCOLOR     LR = 0x0000_0000
	LR_MONOCHROME       LR = 0x0000_0001
	LR_COLOR            LR = 0x0000_0002
	LR_COPYRETURNORG    LR = 0x0000_0004
	LR_COPYDELETEORG    LR = 0x0000_0008
	LR_LOADFROMFILE     LR = 0x0000_0010
	LR_LOADTRANSPARENT  LR = 0x0000_0020
	LR_DEFAULTSIZE      LR = 0x0000_0040
	LR_VGACOLOR         LR = 0x0000_0080
	LR_LOADMAP3DCOLORS  LR = 0x0000_1000
	LR_CREATEDIBSECTION LR = 0x0000_2000
	LR_COPYFROMRESOURCE LR = 0x0000_4000
	LR_SHARED           LR = 0x0000_8000
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
	LVCDI_ITEM     LVCDI = 0x0000_0000
	LVCDI_GROUP    LVCDI = 0x0000_0001
	LVCDI_TEMSLIST LVCDI = 0x0000_0002
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
	LVCFMT_C_FIXED_WIDTH     LVCFMT_C = 0x0_0100
	LVCFMT_C_NO_DPI_SCALE    LVCFMT_C = 0x4_0000
	LVCFMT_C_FIXED_RATIO     LVCFMT_C = 0x8_0000
	LVCFMT_C_SPLITBUTTON     LVCFMT_C = 0x100_0000
)

// LVITEM piColFmt.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVCFMT_I int32

const (
	LVCFMT_I_LINE_BREAK         LVCFMT_I = 0x10_0000
	LVCFMT_I_FILL               LVCFMT_I = 0x20_0000
	LVCFMT_I_WRAP               LVCFMT_I = 0x40_0000
	LVCFMT_I_NO_TITLE           LVCFMT_I = 0x80_0000
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
	LVGA_HEADER_LEFT   LVGA_HEADER = 0x0000_0001
	LVGA_HEADER_CENTER LVGA_HEADER = 0x0000_0002
	LVGA_HEADER_RIGHT  LVGA_HEADER = 0x0000_0004
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
