//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// [ACCEL] struct.
//
// [ACCEL]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
type ACCEL struct {
	FVirt co.ACCELF // Modifiers.
	Key   co.VK     // Virtual key code.
	Cmd   uint16    // LOWORD(wParam) value.
}

// [COMPAREITEMSTRUCT] struct.
//
// [COMPAREITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-compareitemstruct
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

// [CONVCONTEXT] struct.
//
// ⚠️ You must call SetCb() to initialize the struct.
//
// [CONVCONTEXT]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/ns-ddeml-convcontext
type CONVCONTEXT struct {
	cb         uint32
	wFlags     uint32
	WCountryID uint32
	iCodePage  int32
	dwLangId   uint32 // LANGID
	DwSecurity uint32
	Qos        SECURITY_QUALITY_OF_SERVICE
}

func (cc *CONVCONTEXT) SetCb() {
	cc.cb = uint32(unsafe.Sizeof(*cc))
	cc.iCodePage = _CP_WINUNICODE
	cc.Qos.SetLength()
}

func (cc *CONVCONTEXT) DwLangId() LANGID       { return LANGID(cc.dwLangId) }
func (cc *CONVCONTEXT) SetDwLangId(val LANGID) { cc.dwLangId = uint32(val) }

// [COPYDATASTRUCT] struct.
//
// [COPYDATASTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-copydatastruct
type COPYDATASTRUCT struct {
	DwData uintptr // ULONG_PTR
	CbData uint32
	LpData uintptr // PVOID
}

// [CREATESTRUCT] struct.
//
// [CREATESTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-createstructw
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

// [CWPRETSTRUCT] struct.
//
// [CWPRETSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-cwpretstruct
type CWPRETSTRUCT struct {
	LResult uintptr // LRESULT
	LParam  LPARAM
	WParam  WPARAM
	Message co.WM
	Hwnd    HWND
}

// [DELETEITEMSTRUCT] struct.
//
// [DELETEITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-deleteitemstruct
type DELETEITEMSTRUCT struct {
	CtlType  co.ODT_C
	CtlID    uint32
	ItemID   uint32
	HwndItem HWND
	ItemData uintptr // ULONG_PTR
}

// [DISPLAY_DEVICE] struct.
//
// ⚠️ You must call SetCb() to initialize the struct.
//
// [DISPLAY_DEVICE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-display_devicew
type DISPLAY_DEVICE struct {
	cb           uint32
	deviceName   [32]uint16
	deviceString [128]uint16
	StateFlags   co.DISPLAY_DEVICE
	deviceID     [128]uint16
	deviceKey    [128]int16
}

func (dd *DISPLAY_DEVICE) SetCb() { dd.cb = uint32(unsafe.Sizeof(*dd)) }

func (dd *DISPLAY_DEVICE) DeviceName() string { return Str.FromNativeSlice(dd.deviceName[:]) }
func (dd *DISPLAY_DEVICE) SetDeviceName(val string) {
	copy(dd.deviceName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dd.deviceName)-1)))
}

func (dd *DISPLAY_DEVICE) DeviceString() string { return Str.FromNativeSlice(dd.deviceString[:]) }
func (dd *DISPLAY_DEVICE) SetDeviceString(val string) {
	copy(dd.deviceString[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dd.deviceString)-1)))
}

// [DLGITEMTEMPLATE] struct.
//
// [DLGITEMTEMPLATE]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-dlgitemtemplate
type DLGITEMTEMPLATE struct {
	Style           co.WS
	DwExtendedStyle co.WS_EX
	X, Y, Cx, Cy    int16
	Id              uint16
}

// [DLGTEMPLATE] struct.
//
// [DLGTEMPLATE]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-dlgtemplate
type DLGTEMPLATE struct {
	Style           co.WS
	DwExtendedStyle co.WS_EX
	Cdit            uint16
	X, Y, Cx, Cy    int16
}

// [DRAWITEMSTRUCT] struct.
//
// [DRAWITEMSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
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

// [GUITHREADINFO] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [GUITHREADINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-guithreadinfo
type GUITHREADINFO struct {
	cbSize        uint32
	Flags         co.GUI
	HwndActive    HWND
	HwndFocus     HWND
	HwndCapture   HWND
	HwndMenuOwner HWND
	HwndMoveSize  HWND
	HwndCaret     HWND
	RcCaret       RECT
}

func (gti *GUITHREADINFO) SetCbSize() { gti.cbSize = uint32(unsafe.Sizeof(*gti)) }

// [HELPINFO] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [HELPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
type HELPINFO struct {
	cbSize       uint32
	IContextType co.HELPINFO
	ICtrlId      int32
	HItemHandle  HANDLE
	DwContextId  uintptr // DWORD_PTR
	MousePos     POINT
}

func (hi *HELPINFO) SetCbSize() { hi.cbSize = uint32(unsafe.Sizeof(*hi)) }

// [ICONINFO] struct.
//
// [ICONINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfo
type ICONINFO struct {
	fIcon    int32 // BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

func (ii *ICONINFO) FIcon() bool       { return ii.fIcon != 0 }
func (ii *ICONINFO) SetFIcon(val bool) { ii.fIcon = util.BoolToInt32(val) }

// [ICONINFOEX] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [ICONINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfoexw
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
func (iix *ICONINFOEX) SetFIcon(val bool) { iix.fIcon = util.BoolToInt32(val) }

func (iix *ICONINFOEX) SzModName() string { return Str.FromNativeSlice(iix.szModName[:]) }
func (iix *ICONINFOEX) SetSzModName(val string) {
	copy(iix.szModName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(iix.szModName)-1)))
}

func (iix *ICONINFOEX) SzResName() string { return Str.FromNativeSlice(iix.szResName[:]) }
func (iix *ICONINFOEX) SetSzResName(val string) {
	copy(iix.szResName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(iix.szResName)-1)))
}

// [MDINEXTMENU] struct.
//
// [MDINEXTMENU]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mdinextmenu
type MDINEXTMENU struct {
	HmenuIn   HMENU
	HmenuNext HMENU
	HwndNext  HWND
}

// [MENUGETOBJECTINFO] struct.
//
// [MENUGETOBJECTINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menugetobjectinfo
type MENUGETOBJECTINFO struct {
	DwFlags co.MNGOF
	UPos    uint32
	Hmenu   HMENU
	Riid    uintptr // PVOID
	PvObj   uintptr // PVOID
}

// [MENUINFO] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [MENUINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
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

// [MENUITEMINFO] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [MENUITEMINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
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

// [MINMAXINFO] struct.
//
// [MINMAXINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-minmaxinfo
type MINMAXINFO struct {
	ptReserved     POINT
	PtMaxSize      POINT
	PtMaxPosition  POINT
	PtMinTrackSize POINT
	PtMaxTrackSize POINT
}

// [MONITORINFOEX] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [MONITORINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-monitorinfoexw
type MONITORINFOEX struct {
	cbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   co.MONITORINFOF
	szDevice  [_CCHDEVICENAME]uint16
}

func (mi *MONITORINFOEX) SetCbSize() { mi.cbSize = uint32(unsafe.Sizeof(*mi)) }

func (dtf *MONITORINFOEX) SzDevice() string { return Str.FromNativeSlice(dtf.szDevice[:]) }
func (dtf *MONITORINFOEX) SetSzDevice(val string) {
	copy(dtf.szDevice[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtf.szDevice)-1)))
}

// [MSG] struct.
//
// [MSG]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type MSG struct {
	HWnd   HWND
	Msg    uint32
	WParam WPARAM
	LParam LPARAM
	Time   uint32
	Pt     POINT
}

// [NCCALCSIZE_PARAMS] struct.
//
// [NCCALCSIZE_PARAMS]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nccalcsize_params
type NCCALCSIZE_PARAMS struct {
	Rgrc  [3]RECT
	Lppos *WINDOWPOS
}

// [NMHDR] struct.
//
// [NMHDR]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nmhdr
type NMHDR struct {
	HWndFrom HWND
	IdFrom   uintptr // UINT_PTR, actually it's a simple control ID
	Code     uint32  // in fact it should be int32
}

// [NONCLIENTMETRICS] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [NONCLIENTMETRICS]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nonclientmetricsw
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

// [PAINTSTRUCT] struct.
//
// [PAINTSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-paintstruct
type PAINTSTRUCT struct {
	Hdc         HDC
	fErase      int32 // BOOL
	RcPaint     RECT
	fRestore    int32 // BOOL
	fIncUpdate  int32 // BOOL
	rgbReserved [32]byte
}

func (ps *PAINTSTRUCT) FErase() bool       { return ps.fErase != 0 }
func (ps *PAINTSTRUCT) SetFErase(val bool) { ps.fErase = util.BoolToInt32(val) }

// [POINT] struct.
//
// Basic point structure, with x and y coordinates.
//
// [POINT]: https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
type POINT struct {
	X, Y int32
}

// [POWERBROADCAST_SETTING] struct.
//
// [POWERBROADCAST_SETTING]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-powerbroadcast_setting
type POWERBROADCAST_SETTING struct {
	PowerSetting GUID
	DataLength   uint32
	data         [1]uint8
}

func (pbs *POWERBROADCAST_SETTING) Data(i int) *uint8 { return &pbs.data[i] }

// [RECT] struct.
//
// Basic rectangle structure, with left, top, right and bottom values.
//
// [RECT]: https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	Left, Top, Right, Bottom int32
}

// [SCROLLINFO] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [SCROLLINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-scrollinfo
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

// [SECURITY_QUALITY_OF_SERVICE] struct.
//
// ⚠️ You must call SetLength() to initialize the struct.
//
// [SECURITY_QUALITY_OF_SERVICE]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-security_quality_of_service
type SECURITY_QUALITY_OF_SERVICE struct {
	length              uint32
	ImpersonationLevel  co.SECURITY_IMPERSONATION_LEVEL
	contextTrackingMode uint8 // BOOLEAN
	effectiveOnly       uint8 // BOOLEAN
}

func (qs *SECURITY_QUALITY_OF_SERVICE) SetLength() { qs.length = uint32(unsafe.Sizeof(*qs)) }

func (qs *SECURITY_QUALITY_OF_SERVICE) ContextTrackingMode() bool { return qs.contextTrackingMode != 0 }
func (qs *SECURITY_QUALITY_OF_SERVICE) SetContextTrackingMode(val bool) {
	qs.contextTrackingMode = util.BoolToUint8(val)
}

func (qs *SECURITY_QUALITY_OF_SERVICE) EffectiveOnly() bool { return qs.effectiveOnly != 0 }
func (qs *SECURITY_QUALITY_OF_SERVICE) SetEffectiveOnly(val bool) {
	qs.effectiveOnly = util.BoolToUint8(val)
}

// [SIZE] struct.
//
// Basic area size structure, with cx and cy values.
//
// [SIZE]: https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-size
type SIZE struct {
	Cx, Cy int32
}

// [STYLESTRUCT] for WS styles.
//
// [STYLESTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-stylestruct
type STYLESTRUCT_WS struct {
	StyleOld co.WS
	StyleNew co.WS
}

// [STYLESTRUCT] for WS_EX styles.
//
// [STYLESTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-stylestruct
type STYLESTRUCT_WSEX struct {
	StyleOld co.WS_EX
	StyleNew co.WS_EX
}

// [TITLEBARINFOEX] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [TITLEBARINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-titlebarinfoex
type TITLEBARINFOEX struct {
	cbSize     uint32
	RcTitleBar RECT
	Rgstate    [_CCHILDREN_TITLEBAR + 1]uint32
	Rgrect     [_CCHILDREN_TITLEBAR + 1]RECT
}

func (tix *TITLEBARINFOEX) SetCbSize() { tix.cbSize = uint32(unsafe.Sizeof(*tix)) }

// [WINDOWPOS] struct.
//
// [WINDOWPOS]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-windowpos
type WINDOWPOS struct {
	HwndInsertAfter HWND
	Hwnd            HWND
	X, Y, Cx, Cy    int32
	Flags           co.SWP
}

// [WNDCLASSEX] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// [WNDCLASSEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
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
