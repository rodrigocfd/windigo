package win

import (
	"time"

	"github.com/rodrigocfd/windigo/win/co"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
type ACCEL struct {
	FVirt co.ACCELF // Modifiers.
	Key   co.VK     // Virtual key code.
	Cmd   uint16    // LOWORD(wParam) value.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BITMAPINFOHEADER struct {
	BiSize          uint32
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-compareitemstruct
type COMPAREITEMSTRUCT struct {
	CtlType    co.ODT_I
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
	CtlType  co.ODT_I
	CtlID    uint32
	ItemID   uint32
	HwndItem HWND
	ItemData uintptr // ULONG_PTR
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
type DRAWITEMSTRUCT struct {
	CtlType    co.ODT_D
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
	StandardName                [32]uint16
	StandardDate                SYSTEMTIME
	StandardBias                int32
	DaylightName                [32]uint16
	DaylightDate                SYSTEMTIME
	DaylightBias                int32
	TimeZoneKeyName             [128]uint16
	DynamicDaylightTimeDisabled uint8 // BOOLEAN
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

// Can be created with NewGuid().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint64
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
type HELPINFO struct {
	CbSize       uint32
	IContextType co.HELPINFO
	ICtrlId      int32
	HItemHandle  HANDLE
	DwContextId  uintptr // DWORD_PTR
	MousePos     POINT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfo
type ICONINFO struct {
	FIcon    BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfoexw
type ICONINFOEX struct {
	CbSize    uint32
	FIcon     BOOL
	XHotspot  uint32
	YHotspot  uint32
	HbmMask   HBITMAP
	HbmColor  HBITMAP
	WResID    uint16
	SzModName [_MAX_PATH]uint16
	SzResName [_MAX_PATH]uint16
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
	LfFaceName       [_LF_FACESIZE]uint16
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
type MENUINFO struct {
	CbSize          uint32
	FMask           co.MIM
	DwStyle         co.MNS
	CyMax           uint32
	HbrBack         HBRUSH
	DwContextHelpID uint32
	DwMenuData      uintptr // ULONG_PTR
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
type MENUITEMINFO struct {
	CbSize        uint32
	FMask         co.MIIM
	FType         co.MFT
	FState        co.MFS
	WId           uint32
	HSubMenu      HMENU
	HBmpChecked   HBITMAP
	HBmpUnchecked HBITMAP
	DwItemData    uintptr // ULONG_PTR
	DwTypeData    uintptr // LPWSTR, content changes according to fType
	Cch           uint32
	HBmpItem      HBITMAP
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nonclientmetricsw
type NONCLIENTMETRICS struct {
	CbSize             uint32
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
type OSVERSIONINFOEX struct {
	DwOsVersionInfoSize uint32
	DwMajorVersion      uint32
	DwMinorVersion      uint32
	DwBuildNumber       uint32
	DWPlatformId        uint32
	SzCSDVersion        [128]uint16
	WServicePackMajor   uint16
	WServicePackMinor   uint16
	WSuiteMask          co.VER_SUITE
	WProductType        uint8
	WReserved           uint8
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-paintstruct
type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

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

// Basic rectangle structure, with left, top, right and bottom values.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	Left, Top, Right, Bottom int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-scrollinfo
type SCROLLINFO struct {
	CbSize    uint32
	FMask     co.SIF
	NMin      uint32
	NMax      uint32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/aa379560(v=vs.85)
type SECURITY_ATTRIBUTES struct {
	NLength              uint32
	LpSecurityDescriptor uintptr // LPVOID
	BInheritHandle       int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  co.SFGAO
	SzDisplayName [_MAX_PATH]uint16
	SzTypeName    [80]uint16
}

// Basic area size structure, with cx and cy values.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-size
type SIZE struct {
	Cx, Cy int32
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
	StandardName [32]uint16
	StandardDate SYSTEMTIME
	StandardBias int32
	DaylightName [32]uint16
	DaylightDate SYSTEMTIME
	DaylightBias int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-titlebarinfoex
type TITLEBARINFOEX struct {
	CbSize     uint32
	RcTitleBar RECT
	Rgstate    [_CCHILDREN_TITLEBAR + 1]uint32
	Rgrect     [_CCHILDREN_TITLEBAR + 1]RECT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-win32_find_dataw
type WIN32_FIND_DATA struct {
	DwFileAttributes   co.FILE_ATTRIBUTE
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      uint32
	NFileSizeLow       uint32
	DwReserved0        uint32
	DwReserved1        uint32
	CFileName          [_MAX_PATH]uint16
	CAlternateFileName [14]uint16
	DwFileType         uint32
	DwCreatorType      uint32
	WFinderFlags       uint16
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-windowpos
type WINDOWPOS struct {
	HwndInsertAfter HWND
	Hwnd            HWND
	X, Y, Cx, Cy    int32
	Flags           co.SWP
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
type WNDCLASSEX struct {
	CbSize        uint32
	Style         co.CS
	LpfnWndProc   uintptr // Pointer to WNDPROC callback.
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
