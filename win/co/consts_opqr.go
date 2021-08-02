package co

// DRAWITEMSTRUCT itemAction.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type ODA uint32

const (
	ODA_DRAWENTIRE ODA = 0x0001
	ODA_SELECT     ODA = 0x0002
	ODA_FOCUS      ODA = 0x0004
)

// DRAWITEMSTRUCT itemState.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type ODS uint32

const (
	ODS_SELECTED     ODS = 0x0001
	ODS_GRAYED       ODS = 0x0002
	ODS_DISABLED     ODS = 0x0004
	ODS_CHECKED      ODS = 0x0008
	ODS_FOCUS        ODS = 0x0010
	ODS_DEFAULT      ODS = 0x0020
	ODS_COMBOBOXEDIT ODS = 0x1000
	ODS_HOTLIGHT     ODS = 0x0040
	ODS_INACTIVE     ODS = 0x0080
	ODS_NOACCEL      ODS = 0x0100
	ODS_NOFOCUSRECT  ODS = 0x0200
)

// DRAWITEMSTRUCT CtlType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type ODT uint32

const (
	ODT_MENU     ODT = 1
	ODT_LISTBOX  ODT = 2
	ODT_COMBOBOX ODT = 3
	ODT_BUTTON   ODT = 4
	ODT_STATIC   ODT = 5
	ODT_TAB      ODT = 101
	ODT_LISTVIEW ODT = 102
)

// COMPAREITEMSTRUCT and DELETEITEMSTRUCT CtlType. Originally with ODT prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-compareitemstruct
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-deleteitemstruct
type ODT_C uint32

const (
	ODT_C_LISTBOX  ODT_C = ODT_C(ODT_LISTBOX)
	ODT_C_COMBOBOX ODT_C = ODT_C(ODT_COMBOBOX)
)

// Used in OpenFile().
//
// Behavior of the file opening.
type OPEN_FILE uint8

const (
	OPEN_FILE_READ_EXISTING     OPEN_FILE = iota // Open an existing file for read only.
	OPEN_FILE_RW_EXISTING                        // Open an existing file for read and write.
	OPEN_FILE_RW_OPEN_OR_CREATE                  // Open a file or create if it doesn't exist, for read and write.
)

// Used in OpenFileMapped().
//
// Behavior of the memory-mapped file opening.
type OPEN_FILEMAP uint8

const (
	OPEN_FILEMAP_MODE_READ OPEN_FILEMAP = iota // Open an existing file for read only.
	OPEN_FILEMAP_MODE_RW                       // Open an existing file for read and write.
)

// CreateFileMapping() flProtect.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type PAGE uint32

const (
	PAGE_NONE                   PAGE = 0
	PAGE_NOACCESS               PAGE = 0x01
	PAGE_READONLY               PAGE = 0x02
	PAGE_READWRITE              PAGE = 0x04
	PAGE_WRITECOPY              PAGE = 0x08
	PAGE_EXECUTE                PAGE = 0x10
	PAGE_EXECUTE_READ           PAGE = 0x20
	PAGE_EXECUTE_READWRITE      PAGE = 0x40
	PAGE_EXECUTE_WRITECOPY      PAGE = 0x80
	PAGE_GUARD                  PAGE = 0x100
	PAGE_NOCACHE                PAGE = 0x200
	PAGE_WRITECOMBINE           PAGE = 0x400
	PAGE_ENCLAVE_THREAD_CONTROL PAGE = 0x80000000
	PAGE_REVERT_TO_FILE_MAP     PAGE = 0x80000000
	PAGE_TARGETS_NO_UPDATE      PAGE = 0x40000000
	PAGE_TARGETS_INVALID        PAGE = 0x40000000
	PAGE_ENCLAVE_UNVALIDATED    PAGE = 0x20000000
	PAGE_ENCLAVE_DECOMMIT       PAGE = 0x10000000
)

// ProgressBar messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-progress-bar-control-reference-messages
const (
	PBM_SETRANGE    WM = WM_USER + 1
	PBM_SETPOS      WM = WM_USER + 2
	PBM_DELTAPOS    WM = WM_USER + 3
	PBM_SETSTEP     WM = WM_USER + 4
	PBM_STEPIT      WM = WM_USER + 5
	PBM_SETRANGE32  WM = WM_USER + 6
	PBM_GETRANGE    WM = WM_USER + 7
	PBM_GETPOS      WM = WM_USER + 8
	PBM_SETBARCOLOR WM = WM_USER + 9
	PBM_SETBKCOLOR  WM = CCM_SETBKCOLOR
	PBM_SETMARQUEE  WM = WM_USER + 10
	PBM_GETSTEP     WM = WM_USER + 13
	PBM_GETBKCOLOR  WM = WM_USER + 14
	PBM_GETBARCOLOR WM = WM_USER + 15
	PBM_SETSTATE    WM = WM_USER + 16
	PBM_GETSTATE    WM = WM_USER + 17
)

// ProgressBar control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/progress-bar-control-styles
type PBS WS

const (
	PBS_SMOOTH        PBS = 0x01
	PBS_VERTICAL      PBS = 0x04
	PBS_MARQUEE       PBS = 0x08
	PBS_SMOOTHREVERSE PBS = 0x10
)

// PBM_SETSTATE state.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/pbm-setstate
type PBST uint32

const (
	PBST_NORMAL PBST = 0x0001
	PBST_ERROR  PBST = 0x0002
	PBST_PAUSED PBST = 0x0003
)

// WM_POWERBROADCAST event.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
type PBT uint32

const (
	PBT_APMQUERYSUSPEND       PBT = 0x0000
	PBT_APMQUERYSTANDBY       PBT = 0x0001
	PBT_APMQUERYSUSPENDFAILED PBT = 0x0002
	PBT_APMQUERYSTANDBYFAILED PBT = 0x0003
	PBT_APMSUSPEND            PBT = 0x0004
	PBT_APMSTANDBY            PBT = 0x0005
	PBT_APMRESUMECRITICAL     PBT = 0x0006
	PBT_APMRESUMESUSPEND      PBT = 0x0007
	PBT_APMRESUMESTANDBY      PBT = 0x0008
	PBT_APMBATTERYLOW         PBT = 0x0009
	PBT_APMPOWERSTATUSCHANGE  PBT = 0x000a
	PBT_APMOEMEVENT           PBT = 0x000b
	PBT_APMRESUMEAUTOMATIC    PBT = 0x0012
	PBT_POWERSETTINGCHANGE    PBT = 0x8013
)

// PeekMessage() wRemoveMsg.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
type PM uint32

const (
	PM_NOREMOVE PM = 0x0000
	PM_REMOVE   PM = 0x0001
	PM_NOYIELD  PM = 0x0002

	PM_QS_INPUT       PM = PM(QS_INPUT << 16)
	PM_QS_PAINT       PM = PM(QS_PAINT << 16)
	PM_QS_POSTMESSAGE PM = PM((QS_POSTMESSAGE | QS_HOTKEY | QS_TIMER) << 16)
	PM_QS_SENDMESSAGE PM = PM(QS_SENDMESSAGE << 16)
)

// SetPolyFillMode() mode. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
type POLYF int32

const (
	POLYF_ALTERNATE POLYF = 1
	POLYF_WINDING   POLYF = 2
)

// WM_PRINT drawing options.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-print
type PRF uint32

const (
	PRF_CHECKVISIBLE PRF = 0x00000001
	PRF_NONCLIENT    PRF = 0x00000002
	PRF_CLIENT       PRF = 0x00000004
	PRF_ERASEBKGND   PRF = 0x00000008
	PRF_CHILDREN     PRF = 0x00000010
	PRF_OWNED        PRF = 0x00000020
)

// SYSTEM_INFO dwProcessorType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type PROCESSOR uint32

const (
	PROCESSOR_INTEL_386     PROCESSOR = 386
	PROCESSOR_INTEL_486     PROCESSOR = 486
	PROCESSOR_INTEL_PENTIUM PROCESSOR = 586
	PROCESSOR_INTEL_IA64    PROCESSOR = 2200
	PROCESSOR_AMD_X8664     PROCESSOR = 8664
	PROCESSOR_MIPS_R4000    PROCESSOR = 4000
	PROCESSOR_ALPHA_21064   PROCESSOR = 21064
	PROCESSOR_PPC_601       PROCESSOR = 601
	PROCESSOR_PPC_603       PROCESSOR = 603
	PROCESSOR_PPC_604       PROCESSOR = 604
	PROCESSOR_PPC_620       PROCESSOR = 620
	PROCESSOR_HITACHI_SH3   PROCESSOR = 10003
	PROCESSOR_HITACHI_SH3E  PROCESSOR = 10004
	PROCESSOR_HITACHI_SH4   PROCESSOR = 10005
	PROCESSOR_MOTOROLA_821  PROCESSOR = 821
	PROCESSOR_SHx_SH3       PROCESSOR = 103
	PROCESSOR_SHx_SH4       PROCESSOR = 104
	PROCESSOR_STRONGARM     PROCESSOR = 2577
	PROCESSOR_ARM720        PROCESSOR = 1824
	PROCESSOR_ARM820        PROCESSOR = 2080
	PROCESSOR_ARM920        PROCESSOR = 2336
	PROCESSOR_ARM_7TDMI     PROCESSOR = 70001
	PROCESSOR_OPTIL         PROCESSOR = 0x494f
)

// SYSTEM_INFO wProcessorArchitecture.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type PROCESSOR_ARCHITECTURE uint16

const (
	PROCESSOR_ARCHITECTURE_INTEL          PROCESSOR_ARCHITECTURE = 0
	PROCESSOR_ARCHITECTURE_MIPS           PROCESSOR_ARCHITECTURE = 1
	PROCESSOR_ARCHITECTURE_ALPHA          PROCESSOR_ARCHITECTURE = 2
	PROCESSOR_ARCHITECTURE_PPC            PROCESSOR_ARCHITECTURE = 3
	PROCESSOR_ARCHITECTURE_SHX            PROCESSOR_ARCHITECTURE = 4
	PROCESSOR_ARCHITECTURE_ARM            PROCESSOR_ARCHITECTURE = 5
	PROCESSOR_ARCHITECTURE_IA64           PROCESSOR_ARCHITECTURE = 6
	PROCESSOR_ARCHITECTURE_ALPHA64        PROCESSOR_ARCHITECTURE = 7
	PROCESSOR_ARCHITECTURE_MSIL           PROCESSOR_ARCHITECTURE = 8
	PROCESSOR_ARCHITECTURE_AMD64          PROCESSOR_ARCHITECTURE = 9
	PROCESSOR_ARCHITECTURE_IA32_ON_WIN64  PROCESSOR_ARCHITECTURE = 10
	PROCESSOR_ARCHITECTURE_NEUTRAL        PROCESSOR_ARCHITECTURE = 11
	PROCESSOR_ARCHITECTURE_ARM64          PROCESSOR_ARCHITECTURE = 12
	PROCESSOR_ARCHITECTURE_ARM32_ON_WIN64 PROCESSOR_ARCHITECTURE = 13
	PROCESSOR_ARCHITECTURE_IA32_ON_ARM64  PROCESSOR_ARCHITECTURE = 14
	PROCESSOR_ARCHITECTURE_UNKNOWN        PROCESSOR_ARCHITECTURE = 0xffff
)

// CreatePen() iStyle.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
type PS int32

const (
	PS_SOLID       PS = 0
	PS_DASH        PS = 1
	PS_DOT         PS = 2
	PS_DASHDOT     PS = 3
	PS_DASHDOTDOT  PS = 4
	PS_NULL        PS = 5
	PS_INSIDEFRAME PS = 6
)

// PolyDraw() aj.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
type PT uint8

const (
	PT_CLOSEFIGURE PT = 0x01
	PT_LINETO      PT = 0x02
	PT_BEZIERTO    PT = 0x04
	PT_MOVETO      PT = 0x06
)

// GetQueueStatus() flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
type QS uint32

const (
	QS_KEY            QS = 0x0001
	QS_MOUSEMOVE      QS = 0x0002
	QS_MOUSEBUTTON    QS = 0x0004
	QS_POSTMESSAGE    QS = 0x0008
	QS_TIMER          QS = 0x0010
	QS_PAINT          QS = 0x0020
	QS_SENDMESSAGE    QS = 0x0040
	QS_HOTKEY         QS = 0x0080
	QS_ALLPOSTMESSAGE QS = 0x0100
	QS_RAWINPUT       QS = 0x0400
	QS_TOUCH          QS = 0x0800
	QS_POINTER        QS = 0x1000
	QS_MOUSE          QS = QS_MOUSEMOVE | QS_MOUSEBUTTON
	QS_INPUT          QS = QS_MOUSE | QS_KEY | QS_RAWINPUT | QS_TOUCH | QS_POINTER
	QS_ALLINPUT       QS = QS_INPUT | QS_POSTMESSAGE | QS_TIMER | QS_PAINT | QS_HOTKEY | QS_SENDMESSAGE
)

// Registry value types.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/sysinfo/registry-value-types
type REG uint32

const (
	REG_NONE                REG = 0  // No value type.
	REG_SZ                  REG = 1  // Unicode nul terminated string.
	REG_EXPAND_SZ           REG = 2  // Unicode nul terminated string (with environment variable references).
	REG_BINARY              REG = 3  // Free form binary.
	REG_DWORD               REG = 4  // 32-bit number.
	REG_DWORD_LITTLE_ENDIAN REG = 4  // 32-bit number (same as REG_DWORD).
	REG_DWORD_BIG_ENDIAN    REG = 5  // 32-bit number.
	REG_LINK                REG = 6  // Symbolic Link (unicode).
	REG_MULTI_SZ            REG = 7  // Multiple Unicode strings.
	REG_QWORD               REG = 11 // 64-bit number.
	REG_QWORD_LITTLE_ENDIAN REG = 11 // 64-bit number (same as REG_QWORD).
)

// SelectObject() return value. Originally with REGION suffix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
type REGION uint32

const (
	REGION_NULL    REGION = 1
	REGION_SIMPLE  REGION = 2
	REGION_COMPLEX REGION = 3
)

// CombineRgn() iMode.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
type RGN int32

const (
	RGN_AND  RGN = 1
	RGN_OR   RGN = 2
	RGN_XOR  RGN = 3
	RGN_DIFF RGN = 4
	RGN_COPY RGN = 5
)

// BitBlt() rop, IMAGELISTDRAWPARAMS dwRop.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type ROP uint32

const (
	ROP_SRCCOPY        ROP = 0x00cc0020
	ROP_SRCPAINT       ROP = 0x00ee0086
	ROP_SRCAND         ROP = 0x008800c6
	ROP_SRCINVERT      ROP = 0x00660046
	ROP_SRCERASE       ROP = 0x00440328
	ROP_NOTSRCCOPY     ROP = 0x00330008
	ROP_NOTSRCERASE    ROP = 0x001100a6
	ROP_MERGECOPY      ROP = 0x00c000ca
	ROP_MERGEPAINT     ROP = 0x00bb0226
	ROP_PATCOPY        ROP = 0x00f00021
	ROP_PATPAINT       ROP = 0x00fb0a09
	ROP_PATINVERT      ROP = 0x005a0049
	ROP_DSTINVERT      ROP = 0x00550009
	ROP_BLACKNESS      ROP = 0x00000042
	ROP_WHITENESS      ROP = 0x00ff0062
	ROP_NOMIRRORBITMAP ROP = 0x80000000
	ROP_CAPTUREBLT     ROP = 0x40000000
)

// RegGetValue() dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
type RRF uint32

const (
	RRF_RT_REG_NONE      RRF = 0x00000001
	RRF_RT_REG_SZ        RRF = 0x00000002
	RRF_RT_REG_EXPAND_SZ RRF = 0x00000004
	RRF_RT_REG_BINARY    RRF = 0x00000008
	RRF_RT_REG_DWORD     RRF = 0x00000010
	RRF_RT_REG_MULTI_SZ  RRF = 0x00000020
	RRF_RT_REG_QWORD     RRF = 0x00000040
	RRF_RT_DWORD         RRF = RRF_RT_REG_BINARY | RRF_RT_REG_DWORD
	RRF_RT_QWORD         RRF = RRF_RT_REG_BINARY | RRF_RT_REG_QWORD
	RRF_RT_ANY           RRF = 0x0000ffff

	RRF_SUBKEY_WOW6464KEY RRF = 0x00010000
	RRF_SUBKEY_WOW6432KEY RRF = 0x00020000
	RRF_NOEXPAND          RRF = 0x10000000
	RRF_ZEROONFAILURE     RRF = 0x20000000
)
