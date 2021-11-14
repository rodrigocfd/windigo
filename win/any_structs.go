package win

import (
	"encoding/binary"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
type ACCEL struct {
	FVirt co.ACCELF // Modifiers.
	Key   co.VK     // Virtual key code.
	Cmd   uint16    // LOWORD(wParam) value.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmap
type BITMAP struct {
	bmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       *byte
}

// ⚠️ You must call BmiHeader.SetBiSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfo
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}

// ⚠️ You must call SetBiSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BITMAPINFOHEADER struct {
	biSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   co.BI
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

func (bih *BITMAPINFOHEADER) SetBiSize() { bih.biSize = uint32(unsafe.Sizeof(*bih)) }

// ⚠️ You must call SetLStructSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-choosecolorw-r1
type CHOOSECOLOR struct {
	lStructSize    uint32
	HwndOwner      HWND
	HInstance      HWND
	RgbResult      COLORREF
	LpCustColors   *COLORREF // Slice must have 16 values.
	Flags          co.CC
	LCustData      uintptr // LPARAM
	LpfnHook       uintptr // LPCCHOOKPROC
	LpTemplateName *uint16
}

func (cc *CHOOSECOLOR) SetLStructSize() { cc.lStructSize = uint32(unsafe.Sizeof(*cc)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-compareitemstruct
type COMPAREITEMSTRUCT struct {
	CtlType    co.ODT_C
	CtlID      uint32
	HwndItem   HWND
	ItemID1    uint32
	ItemData1  uintptr // ULONG_PTR
	ItemID2    uint32
	ItemData2  uintptr // ULONG_PTR
	DwLocaleId uint32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-copydatastruct
type COPYDATASTRUCT struct {
	DwData uintptr // ULONG_PTR
	CbData uint32
	LpData uintptr // PVOID
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-createstructw
type CREATESTRUCT struct {
	LpCreateParams uintptr // LPVOID
	HInstance      HINSTANCE
	HMenu          HMENU
	HwndParent     HWND
	Cy, Cx         int32
	Y, X           int32
	Style          co.WS
	LpszName       *uint16
	LpszClass      *uint16
	ExStyle        co.WS_EX
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-deleteitemstruct
type DELETEITEMSTRUCT struct {
	CtlType  co.ODT_C
	CtlID    uint32
	ItemID   uint32
	HwndItem HWND
	ItemData uintptr // ULONG_PTR
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type DRAWITEMSTRUCT struct {
	CtlType    co.ODT
	CtlID      uint32
	ItemID     uint32
	ItemAction co.ODA
	ItemState  co.ODS
	HwndItem   HWND
	Hdc        HDC
	RcItem     RECT
	ItemData   uintptr // ULONG_PTR
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/ns-timezoneapi-dynamic_time_zone_information
type DYNAMIC_TIME_ZONE_INFORMATION struct {
	Bias                        int32
	standardName                [32]uint16
	StandardDate                SYSTEMTIME
	StandardBias                int32
	daylightName                [32]uint16
	DaylightDate                SYSTEMTIME
	DaylightBias                int32
	timeZoneKeyName             [128]uint16
	dynamicDaylightTimeDisabled uint8 // BOOLEAN
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) StandardName() string {
	return Str.FromNativeSlice(dtz.standardName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetStandardName(val string) {
	copy(dtz.standardName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtz.standardName)-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) DaylightName() string {
	return Str.FromNativeSlice(dtz.daylightName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	copy(dtz.daylightName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtz.daylightName)-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) TimeZoneKeyName() string {
	return Str.FromNativeSlice(dtz.timeZoneKeyName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetTimeZoneKeyName(val string) {
	copy(dtz.timeZoneKeyName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtz.timeZoneKeyName)-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) DynamicDaylightTimeDisabled() bool {
	return dtz.dynamicDaylightTimeDisabled != 0
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetDynamicDaylightTimeDisabled(val bool) {
	dtz.dynamicDaylightTimeDisabled = uint8(util.BoolToUintptr(val))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
type FILETIME struct {
	dwLowDateTime  uint32
	dwHighDateTime uint32
}

func (ft *FILETIME) EpochNano100() uint64 { return util.Make64(ft.dwLowDateTime, ft.dwHighDateTime) }
func (ft *FILETIME) SetEpochNano100(val uint64) {
	ft.dwLowDateTime, ft.dwHighDateTime = util.Break64(val)
}

func (ft *FILETIME) ToTime() time.Time {
	// https://stackoverflow.com/a/4135003/6923555
	return time.Unix(0, int64(util.Make64(ft.dwLowDateTime, ft.dwHighDateTime)-116_444_736_000_000_000)*100)
}
func (ft *FILETIME) FromTime(val time.Time) {
	ft.dwLowDateTime, ft.dwHighDateTime = util.Break64(
		uint64(val.UnixNano()/100 + 116_444_736_000_000_000),
	)
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
type HELPINFO struct {
	cbSize       uint32
	IContextType co.HELPINFO
	ICtrlId      int32
	HItemHandle  HANDLE
	DwContextId  uintptr // DWORD_PTR
	MousePos     POINT
}

func (hi *HELPINFO) SetCbSize() { hi.cbSize = uint32(unsafe.Sizeof(*hi)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfo
type ICONINFO struct {
	fIcon    int32 // BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

func (ii *ICONINFO) FIcon() bool       { return ii.fIcon != 0 }
func (ii *ICONINFO) SetFIcon(val bool) { ii.fIcon = int32(util.BoolToUintptr(val)) }

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfoexw
type ICONINFOEX struct {
	cbSize    uint32
	fIcon     int32 // BOOL
	XHotspot  uint32
	YHotspot  uint32
	HbmMask   HBITMAP
	HbmColor  HBITMAP
	WResID    uint16
	szModName [_MAX_PATH]uint16
	szResName [_MAX_PATH]uint16
}

func (iix *ICONINFOEX) SetCbSize() { iix.cbSize = uint32(unsafe.Sizeof(*iix)) }

func (iix *ICONINFOEX) FIcon() bool       { return iix.fIcon != 0 }
func (iix *ICONINFOEX) SetFIcon(val bool) { iix.fIcon = int32(util.BoolToUintptr(val)) }

func (iix *ICONINFOEX) SzModName() string { return Str.FromNativeSlice(iix.szModName[:]) }
func (iix *ICONINFOEX) SetSzModName(val string) {
	copy(iix.szModName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(iix.szModName)-1)))
}

func (iix *ICONINFOEX) SzResName() string { return Str.FromNativeSlice(iix.szResName[:]) }
func (iix *ICONINFOEX) SetSzResName(val string) {
	copy(iix.szResName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(iix.szResName)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logbrush
type LOGBRUSH struct {
	LbStyle co.BRS
	LbColor COLORREF
	LbHatch uintptr // ULONG_PTR
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         co.FW
	LfItalic         uint8
	LfUnderline      uint8
	LfStrikeOut      uint8
	LfCharSet        uint8
	LfOutPrecision   uint8
	LfClipPrecision  uint8
	LfQuality        uint8
	LfPitchAndFamily uint8
	lfFaceName       [_LF_FACESIZE]uint16
}

func (lf *LOGFONT) LfFaceName() string { return Str.FromNativeSlice(lf.lfFaceName[:]) }
func (lf *LOGFONT) SetLfFaceName(val string) {
	copy(lf.lfFaceName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(lf.lfFaceName)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logpen
type LOGPEN struct {
	LopnStyle co.PS
	LopnWidth POINT
	LopnColor COLORREF
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/ns-uxtheme-margins
type MARGINS struct {
	CxLeftWidth    int32
	CxRightWidth   int32
	CyTopHeight    int32
	CyBottomHeight int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mdinextmenu
type MDINEXTMENU struct {
	HmenuIn   HMENU
	HmenuNext HMENU
	HwndNext  HWND
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menugetobjectinfo
type MENUGETOBJECTINFO struct {
	DwFlags co.MNGOF
	UPos    uint32
	Hmenu   HMENU
	Riid    uintptr // PVOID
	PvObj   uintptr // PVOID
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
type MENUINFO struct {
	cbSize          uint32
	FMask           co.MIM
	DwStyle         co.MNS
	CyMax           uint32
	HbrBack         HBRUSH
	DwContextHelpID uint32
	DwMenuData      uintptr // ULONG_PTR
}

func (mi *MENUINFO) SetCbSize() { mi.cbSize = uint32(unsafe.Sizeof(*mi)) }

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MENUITEMINFO struct {
	cbSize        uint32
	FMask         co.MIIM
	FType         co.MFT
	FState        co.MFS
	WId           uint32
	HSubMenu      HMENU
	HBmpChecked   HBITMAP
	HBmpUnchecked HBITMAP
	DwItemData    uintptr // ULONG_PTR
	DwTypeData    *uint16 // content changes according to fType
	Cch           uint32
	HBmpItem      HBITMAP
}

func (mii *MENUITEMINFO) SetCbSize() { mii.cbSize = uint32(unsafe.Sizeof(*mii)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-minmaxinfo
type MINMAXINFO struct {
	ptReserved     POINT
	PtMaxSize      POINT
	PtMaxPosition  POINT
	PtMinTrackSize POINT
	PtMaxTrackSize POINT
}

// ⚠️ You must call SetDwSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-moduleentry32w
type MODULEENTRY32 struct {
	dwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   uintptr
	ModBaseSize   uint32
	HModule       HINSTANCE
	szModule      [_MAX_MODULE_NAME32 + 1]uint16
	szExePath     [_MAX_PATH]uint16
}

func (me *MODULEENTRY32) SetDwSize() { me.dwSize = uint32(unsafe.Sizeof(*me)) }

func (me *MODULEENTRY32) SzModule() string { return Str.FromNativeSlice(me.szModule[:]) }
func (me *MODULEENTRY32) SetSzModule(val string) {
	copy(me.szModule[:], Str.ToNativeSlice(Str.Substr(val, 0, len(me.szModule)-1)))
}

func (me *MODULEENTRY32) SzExePath() string { return Str.FromNativeSlice(me.szExePath[:]) }
func (me *MODULEENTRY32) SetSzExePath(val string) {
	copy(me.szExePath[:], Str.ToNativeSlice(Str.Substr(val, 0, len(me.szExePath)-1)))
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-monitorinfo
type MONITORINFO struct {
	cbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   co.MONITORINFOF
}

func (mi *MONITORINFO) SetCbSize() { mi.cbSize = uint32(unsafe.Sizeof(*mi)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type MSG struct {
	HWnd   HWND
	Msg    uint32
	WParam WPARAM
	LParam LPARAM
	Time   uint32
	Pt     POINT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nccalcsize_params
type NCCALCSIZE_PARAMS struct {
	Rgrc  [3]RECT
	Lppos *WINDOWPOS
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nmhdr
type NMHDR struct {
	HWndFrom HWND
	IdFrom   uintptr // UINT_PTR, actually it's a simple control ID
	Code     uint32  // in fact it should be int32
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nonclientmetricsw
type NONCLIENTMETRICS struct {
	cbSize             uint32
	IBorderWidth       int32
	IScrollWidth       int32
	IScrollHeight      int32
	ICaptionWidth      int32
	ICaptionHeight     int32
	LfCaptionFont      LOGFONT
	ISmCaptionWidth    int32
	ISmCaptionHeight   int32
	LfSmCaptionFont    LOGFONT
	IMenuWidth         int32
	IMenuHeight        int32
	LfMenuFont         LOGFONT
	LfStatusFont       LOGFONT
	LfMessageFont      LOGFONT
	IPaddedBorderWidth int32
}

func (ncm *NONCLIENTMETRICS) CbSize() uint32 { return ncm.cbSize }
func (ncm *NONCLIENTMETRICS) SetCbSize() {
	ncm.cbSize = uint32(unsafe.Sizeof(*ncm))
	if !IsWindowsVistaOrGreater() {
		ncm.cbSize -= uint32(unsafe.Sizeof(ncm.IBorderWidth))
	}
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NOTIFYICONDATA struct {
	cbSize           uint32
	Hwnd             HWND
	UID              uint32
	UFlags           co.NIF
	UCallbackMessage co.WM
	HIcon            HICON
	szTip            [128]uint16
	DwState          co.NIS
	DwStateMask      co.NIS
	szInfo           [256]uint16
	UTimeoutVersion  uint32 // union
	szInfoTitle      [64]uint16
	DwInfoFlags      co.NIIF
	GuidItem         GUID
	HBalloonIcon     HICON
}

func (nid *NOTIFYICONDATA) SetCbSize() { nid.cbSize = uint32(unsafe.Sizeof(*nid)) }

func (nid *NOTIFYICONDATA) SzTip() string { return Str.FromNativeSlice(nid.szTip[:]) }
func (nid *NOTIFYICONDATA) SetSzTip(val string) {
	copy(nid.szTip[:], Str.ToNativeSlice(Str.Substr(val, 0, len(nid.szTip)-1)))
}

func (nid *NOTIFYICONDATA) SzInfo() string { return Str.FromNativeSlice(nid.szInfo[:]) }
func (nid *NOTIFYICONDATA) SetSzInfo(val string) {
	copy(nid.szInfo[:], Str.ToNativeSlice(Str.Substr(val, 0, len(nid.szInfo)-1)))
}

func (nid *NOTIFYICONDATA) SzInfoTitle() string { return Str.FromNativeSlice(nid.szInfoTitle[:]) }
func (nid *NOTIFYICONDATA) SetSzInfoTitle(val string) {
	copy(nid.szInfoTitle[:], Str.ToNativeSlice(Str.Substr(val, 0, len(nid.szInfoTitle)-1)))
}

// ⚠️ You must call SetDwOsVersionInfoSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
type OSVERSIONINFOEX struct {
	DwOsVersionInfoSize uint32
	DwMajorVersion      uint32
	DwMinorVersion      uint32
	DwBuildNumber       uint32
	DWPlatformId        uint32
	szCSDVersion        [128]uint16
	WServicePackMajor   uint16
	WServicePackMinor   uint16
	WSuiteMask          co.VER_SUITE
	WProductType        uint8
	wReserved           uint8
}

func (osv *OSVERSIONINFOEX) SetDwOsVersionInfoSize() {
	osv.DwOsVersionInfoSize = uint32(unsafe.Sizeof(*osv))
}

func (osv *OSVERSIONINFOEX) SzCSDVersion() string { return Str.FromNativeSlice(osv.szCSDVersion[:]) }
func (osv *OSVERSIONINFOEX) SetSzCSDVersion(val string) {
	copy(osv.szCSDVersion[:], Str.ToNativeSlice(Str.Substr(val, 0, len(osv.szCSDVersion)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-paintstruct
type PAINTSTRUCT struct {
	Hdc         HDC
	fErase      int32 // BOOL
	RcPaint     RECT
	fRestore    int32 // BOOL
	fIncUpdate  int32 // BOOL
	rgbReserved [32]byte
}

func (ps *PAINTSTRUCT) FErase() bool       { return ps.fErase != 0 }
func (ps *PAINTSTRUCT) SetFErase(val bool) { ps.fErase = int32(util.BoolToUintptr(val)) }

// Basic point structure, with x and y coordinates.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
type POINT struct {
	X, Y int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-powerbroadcast_setting
type POWERBROADCAST_SETTING struct {
	PowerSetting GUID
	DataLength   uint32
	Data         [1]uint16
}

// ⚠️ You must call SetDwSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-processentry32w
type PROCESSENTRY32 struct {
	dwSize              uint32
	cntUsage            uint32
	Th32ProcessID       uint32
	th32DefaultHeapID   uintptr
	th32ModuleID        uint32
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      int32
	dwFlags             uint32
	szExeFile           [_MAX_PATH]uint16
}

func (pe *PROCESSENTRY32) SetDwSize() { pe.dwSize = uint32(unsafe.Sizeof(*pe)) }

func (me *PROCESSENTRY32) SzExeFile() string { return Str.FromNativeSlice(me.szExeFile[:]) }
func (me *PROCESSENTRY32) SetSzExeFile(val string) {
	copy(me.szExeFile[:], Str.ToNativeSlice(Str.Substr(val, 0, len(me.szExeFile)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-process_information
type PROCESS_INFORMATION struct {
	HProcess    HPROCESS
	HThread     HTHREAD
	DwProcessId uint32
	DwThreadId  uint32
}

// Basic rectangle structure, with left, top, right and bottom values.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	Left, Top, Right, Bottom int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-rgbquad
type RGBQUAD struct {
	data [4]byte
}

func (rq *RGBQUAD) Blue() uint8       { return *(*uint8)(unsafe.Pointer(&rq.data[0])) }
func (rq *RGBQUAD) SetBlue(val uint8) { *(*uint8)(unsafe.Pointer(&rq.data[0])) = val }

func (rq *RGBQUAD) Green() uint8       { return *(*uint8)(unsafe.Pointer(&rq.data[1])) }
func (rq *RGBQUAD) SetGreen(val uint8) { *(*uint8)(unsafe.Pointer(&rq.data[1])) = val }

func (rq *RGBQUAD) Red() uint8       { return *(*uint8)(unsafe.Pointer(&rq.data[2])) }
func (rq *RGBQUAD) SetRed(val uint8) { *(*uint8)(unsafe.Pointer(&rq.data[2])) = val }

func (rq *RGBQUAD) ToColorref() COLORREF { return RGB(rq.Red(), rq.Green(), rq.Blue()) }

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-scrollinfo
type SCROLLINFO struct {
	cbSize    uint32
	FMask     co.SIF
	NMin      uint32
	NMax      uint32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}

func (si *SCROLLINFO) SetCbSize() { si.cbSize = uint32(unsafe.Sizeof(*si)) }

// ⚠️ You must call SetNLength().
//
// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/aa379560(v=vs.85)
type SECURITY_ATTRIBUTES struct {
	nLength              uint32
	LpSecurityDescriptor uintptr // LPVOID
	bInheritHandle       int32   // BOOL
}

func (sa *SECURITY_ATTRIBUTES) SetNLength() { sa.nLength = uint32(unsafe.Sizeof(*sa)) }

func (sa *SECURITY_ATTRIBUTES) BInheritHandle() bool { return sa.bInheritHandle != 0 }
func (sa *SECURITY_ATTRIBUTES) SetBInheritHandle(val bool) {
	sa.bInheritHandle = int32(util.BoolToUintptr(val))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  co.SFGAO
	szDisplayName [_MAX_PATH]uint16
	szTypeName    [80]uint16
}

func (shf *SHFILEINFO) SzDisplayName() string { return Str.FromNativeSlice(shf.szDisplayName[:]) }
func (shf *SHFILEINFO) SetSzDisplayName(val string) {
	copy(shf.szDisplayName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(shf.szDisplayName)-1)))
}

func (shf *SHFILEINFO) SzTypeName() string { return Str.FromNativeSlice(shf.szTypeName[:]) }
func (shf *SHFILEINFO) SetSzTypeName(val string) {
	copy(shf.szTypeName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(shf.szTypeName)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-taskdialog_button
type TASKDIALOG_BUTTON struct {
	NButtonID     int32
	PszButtonText *uint16
}

// This struct is originally packed, so we must serialize it before using.
func (tdb *TASKDIALOG_BUTTON) serializePacked() []byte {
	var data [12]byte
	binary.LittleEndian.PutUint32(data[0:], uint32(tdb.NButtonID))
	binary.LittleEndian.PutUint64(data[4:], uint64(uintptr(unsafe.Pointer(tdb.PszButtonText))))
	return data[:]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-taskdialogconfig
type TASKDIALOGCONFIG struct {
	HwndParent              HWND
	HInstance               HINSTANCE
	DwFlags                 co.TDF
	DwCommonButtons         co.TDCBF
	PszWindowTitle          *uint16
	HMainIcon               TdcIcon // Union PCWSTR + HICON, but string resource won't be considered.
	PszMainInstruction      *uint16
	PszContent              *uint16
	PButtons                []TASKDIALOG_BUTTON
	NDefaultButton          int32
	PRadioButtons           []TASKDIALOG_BUTTON
	NDefaultRadioButton     int32
	PszVerificationText     *uint16
	PszExpandedInformation  *uint16
	PszExpandedControlText  *uint16
	PszCollapsedControlText *uint16
	HFooterIcon             TdcIcon // Union PCWSTR + HICON, but string resource won't be considered.
	PszFooter               *uint16
	PfCallback              uintptr
	LpCallbackData          uintptr
	CxWidth                 uint32
}

// This struct is originally packed, so we must serialize it before using.
func (td *TASKDIALOGCONFIG) serializePacked() ([]byte, *byte, *byte) { // pointers must be kept alive
	buf := make([]byte, 160) // sizeof(TASKDIALOGCONFIG) packed

	binary.LittleEndian.PutUint32(buf[0:], 160) // cbSize
	binary.LittleEndian.PutUint64(buf[4:], uint64(td.HwndParent))
	binary.LittleEndian.PutUint64(buf[12:], uint64(td.HInstance))
	binary.LittleEndian.PutUint32(buf[20:], uint32(td.DwFlags))
	binary.LittleEndian.PutUint32(buf[24:], uint32(td.DwCommonButtons))
	binary.LittleEndian.PutUint64(buf[28:], uint64(uintptr(unsafe.Pointer(td.PszWindowTitle))))
	binary.LittleEndian.PutUint64(buf[36:], uint64(variantTdcIcon(td.HMainIcon)))
	binary.LittleEndian.PutUint64(buf[44:], uint64(uintptr(unsafe.Pointer(td.PszMainInstruction))))
	binary.LittleEndian.PutUint64(buf[52:], uint64(uintptr(unsafe.Pointer(td.PszContent))))

	var pButtonsPtr *byte
	if len(td.PButtons) > 0 {
		pButtonsBuf := make([]byte, 0, len(td.PButtons)*12) // sizeof(TASKDIALOG_BUTTON) packed
		for i := range td.PButtons {
			pButtonsBuf = append(pButtonsBuf, td.PButtons[i].serializePacked()...)
		}
		pButtonsPtr = &pButtonsBuf[0]
		binary.LittleEndian.PutUint32(buf[60:], uint32(len(td.PButtons)))
		binary.LittleEndian.PutUint64(buf[64:], uint64(uintptr(unsafe.Pointer(&pButtonsBuf[0]))))
	}

	binary.LittleEndian.PutUint32(buf[72:], uint32(td.NDefaultButton))

	var pRadioButtonsPtr *byte
	if len(td.PRadioButtons) > 0 {
		pRadioButtonsBuf := make([]byte, 0, len(td.PRadioButtons)*12) // sizeof(TASKDIALOG_BUTTON) packed
		for i := range td.PRadioButtons {
			pRadioButtonsBuf = append(pRadioButtonsBuf, td.PRadioButtons[i].serializePacked()...)
		}
		pRadioButtonsPtr = &pRadioButtonsBuf[0]
		binary.LittleEndian.PutUint32(buf[76:], uint32(len(td.PRadioButtons)))
		binary.LittleEndian.PutUint64(buf[80:], uint64(uintptr(unsafe.Pointer(&pRadioButtonsBuf[0]))))
	}

	binary.LittleEndian.PutUint32(buf[88:], uint32(td.NDefaultRadioButton))
	binary.LittleEndian.PutUint64(buf[92:], uint64(uintptr(unsafe.Pointer(td.PszVerificationText))))
	binary.LittleEndian.PutUint64(buf[100:], uint64(uintptr(unsafe.Pointer(td.PszExpandedInformation))))
	binary.LittleEndian.PutUint64(buf[108:], uint64(uintptr(unsafe.Pointer(td.PszExpandedControlText))))
	binary.LittleEndian.PutUint64(buf[116:], uint64(uintptr(unsafe.Pointer(td.PszCollapsedControlText))))
	binary.LittleEndian.PutUint64(buf[124:], uint64(variantTdcIcon(td.HFooterIcon)))
	binary.LittleEndian.PutUint64(buf[132:], uint64(uintptr(unsafe.Pointer(td.PszFooter))))
	binary.LittleEndian.PutUint64(buf[140:], uint64(td.PfCallback))
	binary.LittleEndian.PutUint64(buf[148:], uint64(td.LpCallbackData))
	binary.LittleEndian.PutUint32(buf[156:], td.CxWidth)

	return buf, pButtonsPtr, pRadioButtonsPtr
}

// ⚠️ You must call SetDwSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-threadentry32
type THREADENTRY32 struct {
	dwSize             uint32
	cntUsage           uint32
	Th32ThreadID       uint32
	Th32OwnerProcessID uint32
	TpBasePri          int32
	tpDeltaPri         int32
	dwFlags            uint32
}

func (te *THREADENTRY32) SetDwSize() { te.dwSize = uint32(unsafe.Sizeof(*te)) }

// Basic area size structure, with cx and cy values.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-size
type SIZE struct {
	Cx, Cy int32
}

// ⚠️ You must call SetCb().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfow
type STARTUPINFO struct {
	cb              uint32
	lpReserved      *uint16
	LpDesktop       *uint16
	LpTitle         *uint16
	DwX             uint32
	DwY             uint32
	DwXSize         uint32
	DwYSize         uint32
	DwXCountChars   uint32
	DwYCountChars   uint32
	DwFillAttribute uint32
	DwFlags         co.STARTF
	WShowWindow     uint16 // co.SW, should be uint16.
	cbReserved2     uint16
	lpReserved2     *uint8
	HStdInput       uintptr
	HStdOutput      uintptr
	HStdError       uintptr
}

func (si *STARTUPINFO) SetCb() { si.cb = uint32(unsafe.Sizeof(*si)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type SYSTEM_INFO struct {
	WProcessorArchitecture      co.PROCESSOR_ARCHITECTURE
	wReserved                   uint16
	DwPageSize                  uint32
	LpMinimumApplicationAddress uintptr
	LpMaximumApplicationAddress uintptr
	DwActiveProcessorMask       uintptr
	DwNumberOfProcessors        uint32
	DwProcessorType             co.PROCESSOR
	DwAllocationGranularity     uint32
	WProcessorLevel             uint16
	WProcessorRevision          uint16
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime
type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

// Decomposes a time.Duration into this SYSTEMTIME fields.
func (st *SYSTEMTIME) FromDuration(dur time.Duration) {
	*st = SYSTEMTIME{}
	st.WHour = uint16(dur / time.Hour)
	st.WMinute = uint16((dur -
		time.Duration(st.WHour)*time.Hour) / time.Minute)
	st.WSecond = uint16((dur -
		time.Duration(st.WHour)*time.Hour -
		time.Duration(st.WMinute)*time.Minute) / time.Second)
	st.WMilliseconds = uint16((dur -
		time.Duration(st.WHour)*time.Hour -
		time.Duration(st.WMinute)*time.Minute -
		time.Duration(st.WSecond)*time.Second) / time.Millisecond)
}

// Fills this SYSTEMTIME with the value of a time.Time.
func (st *SYSTEMTIME) FromTime(val time.Time) {
	var ft FILETIME
	ft.FromTime(val)

	var stUtc SYSTEMTIME
	FileTimeToSystemTime(&ft, &stUtc)

	// When converted, SYSTEMTIME will receive UTC values, so we need to convert
	// the fields to current timezone.
	SystemTimeToTzSpecificLocalTime(nil, &stUtc, st)
}

// Converts this SYSTEMTIME to time.Time.
func (st *SYSTEMTIME) ToTime() time.Time {
	return time.Date(int(st.WYear),
		time.Month(st.WMonth), int(st.WDay),
		int(st.WHour), int(st.WMinute), int(st.WSecond),
		int(st.WMilliseconds)*1_000_000,
		time.Local)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type TEXTMETRIC struct {
	TmHeight           uint32
	TmAscent           uint32
	TmDescent          uint32
	TmInternalLeading  uint32
	TmExternalLeading  uint32
	TmAveCharWidth     uint32
	TmMaxCharWidth     uint32
	TmWeight           uint32
	TmOverhang         uint32
	TmDigitizedAspectX uint32
	TmDigitizedAspectY uint32
	TmFirstChar        uint16
	TmLastChar         uint16
	TmDefaultChar      uint16
	TmBreakChar        uint16
	TmItalic           uint8
	TmUnderlined       uint8
	TmStruckOut        uint8
	TmPitchAndFamily   uint8
	TmCharSet          co.CHARSET
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/ns-timezoneapi-time_zone_information
type TIME_ZONE_INFORMATION struct {
	Bias         int32
	standardName [32]uint16
	StandardDate SYSTEMTIME
	StandardBias int32
	daylightName [32]uint16
	DaylightDate SYSTEMTIME
	DaylightBias int32
}

func (tzi *TIME_ZONE_INFORMATION) StandardName() string {
	return Str.FromNativeSlice(tzi.standardName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetStandardName(val string) {
	copy(tzi.standardName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(tzi.standardName)-1)))
}

func (tzi *TIME_ZONE_INFORMATION) DaylightName() string {
	return Str.FromNativeSlice(tzi.daylightName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	copy(tzi.daylightName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(tzi.daylightName)-1)))
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-titlebarinfoex
type TITLEBARINFOEX struct {
	cbSize     uint32
	RcTitleBar RECT
	Rgstate    [_CCHILDREN_TITLEBAR + 1]uint32
	Rgrect     [_CCHILDREN_TITLEBAR + 1]RECT
}

func (tix *TITLEBARINFOEX) SetCbSize() { tix.cbSize = uint32(unsafe.Sizeof(*tix)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VS_FIXEDFILEINFO struct {
	DwSignature        uint32
	DwStrucVersion     uint32
	dwFileVersionMS    uint32
	dwFileVersionLS    uint32
	dwProductVersionMS uint32
	dwProductVersionLS uint32
	DwFileFlagsMask    co.VS_FF
	DwFileFlags        co.VS_FF
	DwFileOS           co.VOS
	DwFileType         co.VFT
	DwFileSubtype      co.VFT2
	dwFileDateMS       uint32
	dwFileDateLS       uint32
}

func (ffi *VS_FIXEDFILEINFO) FileVersion() (major, minor, patch, build uint16) {
	return HIWORD(ffi.dwFileVersionMS), LOWORD(ffi.dwFileVersionMS),
		HIWORD(ffi.dwFileVersionLS), LOWORD(ffi.dwFileVersionLS)
}
func (ffi *VS_FIXEDFILEINFO) SetFileVersion(major, minor, patch, build uint16) {
	ffi.dwFileVersionMS = MAKELONG(minor, major)
	ffi.dwFileVersionLS = MAKELONG(build, patch)
}

func (ffi *VS_FIXEDFILEINFO) ProductVersion() (major, minor, patch, build uint16) {
	return HIWORD(ffi.dwProductVersionMS), LOWORD(ffi.dwProductVersionMS),
		HIWORD(ffi.dwProductVersionLS), LOWORD(ffi.dwProductVersionLS)
}
func (ffi *VS_FIXEDFILEINFO) SetProductVersion(major, minor, patch, build uint16) {
	ffi.dwProductVersionMS = MAKELONG(minor, major)
	ffi.dwProductVersionLS = MAKELONG(build, patch)
}

func (ffi *VS_FIXEDFILEINFO) FileDate() uint64 {
	return util.Make64(ffi.dwFileDateLS, ffi.dwFileDateMS)
}
func (ffi *VS_FIXEDFILEINFO) SetFileDate(val uint64) {
	ffi.dwFileDateLS, ffi.dwFileDateMS = util.Break64(val)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-win32_find_dataw
type WIN32_FIND_DATA struct {
	DwFileAttributes    co.FILE_ATTRIBUTE
	FtCreationTime      FILETIME
	FtLastAccessTime    FILETIME
	FtLastWriteTime     FILETIME
	NFileSizeHigh       uint32
	NFileSizeLow        uint32
	dwReserved0         uint32
	dwReserved1         uint32
	cFileName           [_MAX_PATH]uint16
	cCAlternateFileName [14]uint16
	DwFileType          uint32
	DwCreatorType       uint32
	WFinderFlags        uint16
}

func (wfd *WIN32_FIND_DATA) CFileName() string { return Str.FromNativeSlice(wfd.cFileName[:]) }
func (wfd *WIN32_FIND_DATA) SetCFileName(val string) {
	copy(wfd.cFileName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(wfd.cFileName)-1)))
}

func (wfd *WIN32_FIND_DATA) CAlternateFileName() string {
	return Str.FromNativeSlice(wfd.cCAlternateFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCAlternateFileName(val string) {
	copy(wfd.cCAlternateFileName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(wfd.cCAlternateFileName)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-windowpos
type WINDOWPOS struct {
	HwndInsertAfter HWND
	Hwnd            HWND
	X, Y, Cx, Cy    int32
	Flags           co.SWP
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
type WNDCLASSEX struct {
	cbSize        uint32
	Style         co.CS
	LpfnWndProc   uintptr // WNDPROC
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  uintptr
	LpszClassName *uint16
	HIconSm       HICON
}

func (wcx *WNDCLASSEX) SetCbSize() { wcx.cbSize = uint32(unsafe.Sizeof(*wcx)) }
