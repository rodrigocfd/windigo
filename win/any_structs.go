package win

import (
	"fmt"
	"reflect"
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
	DynamicDaylightTimeDisabled uint8 // BOOLEAN
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) StandardName() string {
	return Str.FromUint16Slice(dtz.standardName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetStandardName(val string) {
	copy(dtz.standardName[:], Str.ToUint16Slice(Str.Substr(val, 0, 32-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) DaylightName() string {
	return Str.FromUint16Slice(dtz.daylightName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	copy(dtz.daylightName[:], Str.ToUint16Slice(Str.Substr(val, 0, 32-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) TimeZoneKeyName() string {
	return Str.FromUint16Slice(dtz.timeZoneKeyName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetTimeZoneKeyName(val string) {
	copy(dtz.timeZoneKeyName[:], Str.ToUint16Slice(Str.Substr(val, 0, 128-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

// Fills this FILETIME with the value of a time.Time.
func (ft *FILETIME) FromTime(t time.Time) {
	st := SYSTEMTIME{}
	st.FromTime(t)
	SystemTimeToFileTime(&st, ft)
}

// Converts this FILETIME to time.Time.
func (ft *FILETIME) ToTime() time.Time {
	st := SYSTEMTIME{}
	FileTimeToSystemTime(ft, &st)
	return st.ToTime()
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
	fIcon    BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

func (ii *ICONINFO) FIcon() bool       { return ii.fIcon != 0 }
func (ii *ICONINFO) SetFIcon(val bool) { ii.fIcon = BOOL(util.BoolToUintptr(val)) }

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfoexw
type ICONINFOEX struct {
	cbSize    uint32
	fIcon     BOOL
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
func (iix *ICONINFOEX) SetFIcon(val bool) { iix.fIcon = BOOL(util.BoolToUintptr(val)) }

func (iix *ICONINFOEX) SzModName() string { return Str.FromUint16Slice(iix.szModName[:]) }
func (iix *ICONINFOEX) SetSzModName(val string) {
	copy(iix.szModName[:], Str.ToUint16Slice(Str.Substr(val, 0, _MAX_PATH-1)))
}

func (iix *ICONINFOEX) SzResName() string { return Str.FromUint16Slice(iix.szResName[:]) }
func (iix *ICONINFOEX) SetSzResName(val string) {
	copy(iix.szResName[:], Str.ToUint16Slice(Str.Substr(val, 0, _MAX_PATH-1)))
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

func (lf *LOGFONT) LfFaceName() string { return Str.FromUint16Slice(lf.lfFaceName[:]) }
func (lf *LOGFONT) SetLfFaceName(val string) {
	copy(lf.lfFaceName[:], Str.ToUint16Slice(Str.Substr(val, 0, _LF_FACESIZE-1)))
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

func (nid *NOTIFYICONDATA) SzTip() string { return Str.FromUint16Slice(nid.szTip[:]) }
func (nid *NOTIFYICONDATA) SetSzTip(val string) {
	copy(nid.szTip[:], Str.ToUint16Slice(Str.Substr(val, 0, 128-1)))
}

func (nid *NOTIFYICONDATA) SzInfo() string { return Str.FromUint16Slice(nid.szInfo[:]) }
func (nid *NOTIFYICONDATA) SetSzInfo(val string) {
	copy(nid.szInfo[:], Str.ToUint16Slice(Str.Substr(val, 0, 256-1)))
}

func (nid *NOTIFYICONDATA) SzInfoTitle() string { return Str.FromUint16Slice(nid.szInfoTitle[:]) }
func (nid *NOTIFYICONDATA) SetSzInfoTitle(val string) {
	copy(nid.szInfoTitle[:], Str.ToUint16Slice(Str.Substr(val, 0, 64-1)))
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

func (osv *OSVERSIONINFOEX) SzCSDVersion() string { return Str.FromUint16Slice(osv.szCSDVersion[:]) }
func (osv *OSVERSIONINFOEX) SetSzCSDVersion(val string) {
	copy(osv.szCSDVersion[:], Str.ToUint16Slice(Str.Substr(val, 0, 128-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-paintstruct
type PAINTSTRUCT struct {
	Hdc         HDC
	fErase      BOOL
	RcPaint     RECT
	fRestore    BOOL
	fIncUpdate  BOOL
	rgbReserved [32]byte
}

func (ps *PAINTSTRUCT) FErase() bool       { return ps.fErase != 0 }
func (ps *PAINTSTRUCT) SetFErase(val bool) { ps.fErase = BOOL(util.BoolToUintptr(val)) }

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
	bInheritHandle       BOOL
}

func (sa *SECURITY_ATTRIBUTES) SetNLength() { sa.nLength = uint32(unsafe.Sizeof(*sa)) }

func (sa *SECURITY_ATTRIBUTES) BInheritHandle() bool { return sa.bInheritHandle != 0 }
func (sa *SECURITY_ATTRIBUTES) SetBInheritHandle(val bool) {
	sa.bInheritHandle = BOOL(util.BoolToUintptr(val))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  co.SFGAO
	szDisplayName [_MAX_PATH]uint16
	szTypeName    [80]uint16
}

func (shf *SHFILEINFO) SzDisplayName() string { return Str.FromUint16Slice(shf.szDisplayName[:]) }
func (shf *SHFILEINFO) SetSzDisplayName(val string) {
	copy(shf.szDisplayName[:], Str.ToUint16Slice(Str.Substr(val, 0, _MAX_PATH-1)))
}

func (shf *SHFILEINFO) SzTypeName() string { return Str.FromUint16Slice(shf.szTypeName[:]) }
func (shf *SHFILEINFO) SetSzTypeName(val string) {
	copy(shf.szTypeName[:], Str.ToUint16Slice(Str.Substr(val, 0, 80-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-taskdialog_button
type TASKDIALOG_BUTTON struct {
	data [12]byte
}

func (tdb *TASKDIALOG_BUTTON) NButtonID() int32       { return *(*int32)(unsafe.Pointer(&tdb.data[0])) }
func (tdb *TASKDIALOG_BUTTON) SetNButtonID(val int32) { *(*int32)(unsafe.Pointer(&tdb.data[0])) = val }

func (tdc *TASKDIALOG_BUTTON) PszButtonText() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[4])))
}
func (tdc *TASKDIALOG_BUTTON) SetPszButtonText(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[4])) = Str.ToUint16Ptr(val)
}

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-taskdialogconfig
type TASKDIALOGCONFIG struct {
	data [160]byte
}

func (tdc *TASKDIALOGCONFIG) SetCbSize() {
	*(*uint32)(unsafe.Pointer(&tdc.data[0])) = uint32(unsafe.Sizeof(*tdc))
}

func (tdc *TASKDIALOGCONFIG) HwndParent() HWND       { return *(*HWND)(unsafe.Pointer(&tdc.data[4])) }
func (tdc *TASKDIALOGCONFIG) SetHwndParent(val HWND) { *(*HWND)(unsafe.Pointer(&tdc.data[4])) = val }

func (tdc *TASKDIALOGCONFIG) HInstance() HINSTANCE {
	return *(*HINSTANCE)(unsafe.Pointer(&tdc.data[12]))
}
func (tdc *TASKDIALOGCONFIG) SetHInstance(val HINSTANCE) {
	*(*HINSTANCE)(unsafe.Pointer(&tdc.data[12])) = val
}

func (tdc *TASKDIALOGCONFIG) DwFlags() co.TDF       { return *(*co.TDF)(unsafe.Pointer(&tdc.data[20])) }
func (tdc *TASKDIALOGCONFIG) SetDwFlags(val co.TDF) { *(*co.TDF)(unsafe.Pointer(&tdc.data[20])) = val }

func (tdc *TASKDIALOGCONFIG) DwCommonButtons() co.TDCBF {
	return *(*co.TDCBF)(unsafe.Pointer(&tdc.data[24]))
}
func (tdc *TASKDIALOGCONFIG) SetDwCommonButtons(val co.TDCBF) {
	*(*co.TDCBF)(unsafe.Pointer(&tdc.data[24])) = val
}

func (tdc *TASKDIALOGCONFIG) PszWindowTitle() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[28])))
}
func (tdc *TASKDIALOGCONFIG) SetPszWindowTitle(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[28])) = Str.ToUint16Ptr(val)
}

// Return type can be uint16, TD_ICON or HICON.
func (tdc *TASKDIALOGCONFIG) HMainIcon() interface{} {
	raw := *(*uintptr)(unsafe.Pointer(&tdc.data[36]))
	if HIWORD(uint32(raw)) != 0 {
		return HICON(raw)
	} else if LOWORD(uint32(raw)) >= 0xfffc {
		return co.TD_ICON(raw)
	} else {
		return uint16(raw)
	}
}

// ⚠️ val must be uint16, TD_ICON or HMENU.
func (tdc *TASKDIALOGCONFIG) SetHMainIcon(val interface{}) {
	switch v := val.(type) {
	case uint16:
		*(*uintptr)(unsafe.Pointer(&tdc.data[36])) = uintptr(v)
	case co.TD_ICON:
		*(*uintptr)(unsafe.Pointer(&tdc.data[36])) = uintptr(v)
	case HMENU:
		*(*uintptr)(unsafe.Pointer(&tdc.data[36])) = uintptr(v)
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(val)))
	}
}

func (tdc *TASKDIALOGCONFIG) PszMainInstruction() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[44])))
}
func (tdc *TASKDIALOGCONFIG) SetPszMainInstruction(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[44])) = Str.ToUint16Ptr(val)
}

func (tdc *TASKDIALOGCONFIG) PszContent() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[52])))
}
func (tdc *TASKDIALOGCONFIG) SetPszContent(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[52])) = Str.ToUint16Ptr(val)
}

func (tdc *TASKDIALOGCONFIG) PButtons() []TASKDIALOG_BUTTON {
	cButtons := *(*uint32)(unsafe.Pointer(&tdc.data[60]))
	return unsafe.Slice(*(**TASKDIALOG_BUTTON)(unsafe.Pointer(&tdc.data[64])), cButtons)
}
func (tdc *TASKDIALOGCONFIG) SetPButtons(val []TASKDIALOG_BUTTON) {
	*(*uint32)(unsafe.Pointer(&tdc.data[60])) = uint32(len(val))
	*(**TASKDIALOG_BUTTON)(unsafe.Pointer(&tdc.data[64])) = &val[0]
}

func (tdc *TASKDIALOGCONFIG) NDefaultButton() int32 { return *(*int32)(unsafe.Pointer(&tdc.data[72])) }
func (tdc *TASKDIALOGCONFIG) SetNDefaultButton(val int32) {
	*(*int32)(unsafe.Pointer(&tdc.data[72])) = val
}

func (tdc *TASKDIALOGCONFIG) PRadioButtons() []TASKDIALOG_BUTTON {
	cButtons := *(*uint32)(unsafe.Pointer(&tdc.data[76]))
	return unsafe.Slice(*(**TASKDIALOG_BUTTON)(unsafe.Pointer(&tdc.data[80])), cButtons)
}
func (tdc *TASKDIALOGCONFIG) SetPRadioButtons(val []TASKDIALOG_BUTTON) {
	*(*uint32)(unsafe.Pointer(&tdc.data[76])) = uint32(len(val))
	*(**TASKDIALOG_BUTTON)(unsafe.Pointer(&tdc.data[80])) = &val[0]
}

func (tdc *TASKDIALOGCONFIG) NDefaultRadioButton() int32 {
	return *(*int32)(unsafe.Pointer(&tdc.data[88]))
}
func (tdc *TASKDIALOGCONFIG) SetNDefaultRadioButton(val int32) {
	*(*int32)(unsafe.Pointer(&tdc.data[88])) = val
}

func (tdc *TASKDIALOGCONFIG) PszVerificationText() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[92])))
}
func (tdc *TASKDIALOGCONFIG) SetPszVerificationText(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[92])) = Str.ToUint16Ptr(val)
}

func (tdc *TASKDIALOGCONFIG) PszExpandedInformation() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[100])))
}
func (tdc *TASKDIALOGCONFIG) SetPszExpandedInformation(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[100])) = Str.ToUint16Ptr(val)
}

func (tdc *TASKDIALOGCONFIG) PszExpandedControlText() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[108])))
}
func (tdc *TASKDIALOGCONFIG) SetPszExpandedControlText(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[108])) = Str.ToUint16Ptr(val)
}

func (tdc *TASKDIALOGCONFIG) PszCollapsedControlText() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[116])))
}
func (tdc *TASKDIALOGCONFIG) SetPszCollapsedControlText(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[116])) = Str.ToUint16Ptr(val)
}

// Return type can be uint16 or HICON.
func (tdc *TASKDIALOGCONFIG) HFooterIcon() interface{} {
	raw := *(*uintptr)(unsafe.Pointer(&tdc.data[124]))
	if HIWORD(uint32(raw)) != 0 {
		return HICON(raw)
	} else {
		return uint16(raw)
	}
}

// ⚠️ val must be uint16 or HICON.
func (tdc *TASKDIALOGCONFIG) SetHFooterIcon(val interface{}) {
	switch v := val.(type) {
	case uint16:
		*(*uintptr)(unsafe.Pointer(&tdc.data[124])) = uintptr(v)
	case HICON:
		*(*uintptr)(unsafe.Pointer(&tdc.data[124])) = uintptr(v)
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(val)))
	}
}

func (tdc *TASKDIALOGCONFIG) PszFooter() string {
	return Str.FromUint16Ptr(*(**uint16)(unsafe.Pointer(&tdc.data[132])))
}
func (tdc *TASKDIALOGCONFIG) SetPszFooter(val string) {
	*(**uint16)(unsafe.Pointer(&tdc.data[132])) = Str.ToUint16Ptr(val)
}

func (tdc *TASKDIALOGCONFIG) PfCallback() uintptr { return *(*uintptr)(unsafe.Pointer(&tdc.data[140])) }
func (tdc *TASKDIALOGCONFIG) SetPfCallback(val uintptr) {
	*(*uintptr)(unsafe.Pointer(&tdc.data[140])) = val
}

func (tdc *TASKDIALOGCONFIG) LpCallbackData() uintptr {
	return *(*uintptr)(unsafe.Pointer(&tdc.data[148]))
}
func (tdc *TASKDIALOGCONFIG) SetLpCallbackData(val uintptr) {
	*(*uintptr)(unsafe.Pointer(&tdc.data[148])) = val
}

func (tdc *TASKDIALOGCONFIG) CxWidth() uint32       { return *(*uint32)(unsafe.Pointer(&tdc.data[156])) }
func (tdc *TASKDIALOGCONFIG) SetCxWidth(val uint32) { *(*uint32)(unsafe.Pointer(&tdc.data[156])) = val }

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
func (st *SYSTEMTIME) FromTime(t time.Time) {
	// https://support.microsoft.com/en-ca/help/167296/how-to-convert-a-unix-time-t-to-a-win32-filetime-or-systemtime
	epoch := t.UnixNano()/100 + 116_444_736_000_000_000

	ft := FILETIME{}
	ft.DwLowDateTime = uint32(epoch & 0xffff_ffff)
	ft.DwHighDateTime = uint32(epoch >> 32)

	stUtc := SYSTEMTIME{}
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
	return Str.FromUint16Slice(tzi.standardName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetStandardName(val string) {
	copy(tzi.standardName[:], Str.ToUint16Slice(Str.Substr(val, 0, 32-1)))
}

func (tzi *TIME_ZONE_INFORMATION) DaylightName() string {
	return Str.FromUint16Slice(tzi.daylightName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	copy(tzi.daylightName[:], Str.ToUint16Slice(Str.Substr(val, 0, 32-1)))
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
	DwFileVersionMS    uint32
	DwFileVersionLS    uint32
	DwProductVersionMS uint32
	DwProductVersionLS uint32
	DwFileFlagsMask    co.VS_FF
	DwFileFlags        co.VS_FF
	DwFileOS           co.VOS
	DwFileType         co.VFT
	DwFileSubtype      co.VFT2
	DwFileDateMS       uint32
	DwFileDateLS       uint32
}

// Returns the parsed DwFileVersion fields.
func (vfi *VS_FIXEDFILEINFO) FileVersion() [4]uint16 {
	return [4]uint16{
		HIWORD(vfi.DwFileVersionMS), LOWORD(vfi.DwFileVersionMS),
		HIWORD(vfi.DwFileVersionLS), LOWORD(vfi.DwFileVersionLS),
	}
}

// Returns the parsed DwProductVersion fields.
func (vfi *VS_FIXEDFILEINFO) ProductVersion() [4]uint16 {
	return [4]uint16{
		HIWORD(vfi.DwProductVersionMS), LOWORD(vfi.DwProductVersionMS),
		HIWORD(vfi.DwProductVersionLS), LOWORD(vfi.DwProductVersionLS),
	}
}

// Returns the parsed DwFileDate fields.
func (vfi *VS_FIXEDFILEINFO) FileDate() uint64 {
	return util.Make64(vfi.DwFileDateLS, vfi.DwFileDateMS)
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

func (wfd *WIN32_FIND_DATA) CFileName() string { return Str.FromUint16Slice(wfd.cFileName[:]) }
func (wfd *WIN32_FIND_DATA) SetCFileName(val string) {
	copy(wfd.cFileName[:], Str.ToUint16Slice(Str.Substr(val, 0, _MAX_PATH-1)))
}

func (wfd *WIN32_FIND_DATA) CAlternateFileName() string {
	return Str.FromUint16Slice(wfd.cCAlternateFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCAlternateFileName(val string) {
	copy(wfd.cCAlternateFileName[:], Str.ToUint16Slice(Str.Substr(val, 0, 14-1)))
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
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

func (wcx *WNDCLASSEX) SetCbSize() { wcx.cbSize = uint32(unsafe.Sizeof(*wcx)) }
