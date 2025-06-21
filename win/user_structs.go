//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
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

// [CURSORINFO] struct.
//
// ⚠️ You must call [CURSORINFO.SetCbSize] to initialize the struct.
//
// # Example
//
//	var ci win.CURSORINFO
//	ci.SetCbSize()
//
// [CURSORINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-cursorinfo
type CURSORINFO struct {
	cbSize      uint32
	Flags       co.CURSOR
	HCursor     HCURSOR
	PtScreenPos POINT
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (ci *CURSORINFO) SetCbSize() {
	ci.cbSize = uint32(unsafe.Sizeof(*ci))
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
// ⚠️ You must call [DISPLAY_DEVICE.SetCb] to initialize the struct.
//
// # Example
//
//	var dd win.DISPLAY_DEVICE
//	dd.SetCb()
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

// Sets the cb field to the size of the struct, correctly initializing it.
func (dd *DISPLAY_DEVICE) SetCb() {
	dd.cb = uint32(unsafe.Sizeof(*dd))
}

func (dd *DISPLAY_DEVICE) DeviceName() string {
	return wstr.WinSliceToGo(dd.deviceName[:])
}
func (dd *DISPLAY_DEVICE) SetDeviceName(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(dd.deviceName)-1)), dd.deviceName[:])
}

func (dd *DISPLAY_DEVICE) DeviceString() string {
	return wstr.WinSliceToGo(dd.deviceString[:])
}
func (dd *DISPLAY_DEVICE) SetDeviceString(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(dd.deviceString)-1)), dd.deviceString[:])
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
// ⚠️ You must call [GUITHREADINFO.SetCbSize] to initialize the struct.
//
// # Example
//
//	var gti win.GUITHREADINFO
//	gti.SetCbSize()
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

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (gti *GUITHREADINFO) SetCbSize() {
	gti.cbSize = uint32(unsafe.Sizeof(*gti))
}

// [HELPINFO] struct.
//
// ⚠️ You must call [HELPINFO.SetCbSize] to initialize the struct.
//
// # Example
//
//	var hi win.HELPINFO
//	hi.SetCbSize()
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

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (hi *HELPINFO) SetCbSize() {
	hi.cbSize = uint32(unsafe.Sizeof(*hi))
}

// [ICONINFO] struct.
//
// [ICONINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfo
type ICONINFO struct {
	FIcon    int32 // This is a BOOL value.
	XHotspot uint32
	YHotspot uint32
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

// [ICONINFOEX] struct.
//
// ⚠️ You must call [ICONINFOEX.SetCbSize] to initialize the struct.
//
// # Example
//
//	var iix win.ICONINFOEX
//	iix.SetCbSize()
//
// [ICONINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfoexw
type ICONINFOEX struct {
	cbSize    uint32
	FIcon     int32 // This is a BOOL value.
	XHotspot  uint32
	YHotspot  uint32
	HbmMask   HBITMAP
	HbmColor  HBITMAP
	WResID    uint16
	szModName [utl.MAX_PATH]uint16
	szResName [utl.MAX_PATH]uint16
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (iix *ICONINFOEX) SetCbSize() {
	iix.cbSize = uint32(unsafe.Sizeof(*iix))
}

func (iix *ICONINFOEX) SzModName() string {
	return wstr.WinSliceToGo(iix.szModName[:])
}
func (iix *ICONINFOEX) SetSzModName(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(iix.szModName)-1)), iix.szModName[:])
}

func (iix *ICONINFOEX) SzResName() string {
	return wstr.WinSliceToGo(iix.szResName[:])
}
func (iix *ICONINFOEX) SetSzResName(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(iix.szResName)-1)), iix.szResName[:])
}

// Second message [parameter].
//
// [parameter]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

// [MAKELPARAM] macro.
//
// [MAKELPARAM]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makelparam
func MAKELPARAM(lo, hi uint16) LPARAM {
	return LPARAM(MAKELONG(lo, hi))
}

func (lp LPARAM) LoWord() uint16 {
	return LOWORD(uint32(lp))
}
func (lp LPARAM) HiWord() uint16 {
	return HIWORD(uint32(lp))
}

func (lp LPARAM) MakePoint() POINT {
	return POINT{
		X: int32(lp.LoWord()),
		Y: int32(lp.HiWord()),
	}
}
func (lp LPARAM) MakeSize() SIZE {
	return SIZE{
		Cx: int32(lp.LoWord()),
		Cy: int32(lp.HiWord()),
	}
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
// ⚠️ You must call [MENUINFO.SetCbSize] to initialize the struct.
//
// # Example
//
//	var mi win.MENUINFO
//	mi.SetCbSize()
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

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (mi *MENUINFO) SetCbSize() {
	mi.cbSize = uint32(unsafe.Sizeof(*mi))
}

// [MENUITEMINFO] struct.
//
// ⚠️ You must call [MENUITEMINFO.SetCbSize] to initialize the struct.
//
// # Example
//
//	var mii win.MENUITEMINFO
//	mii.SetCbSize()
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

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (mii *MENUITEMINFO) SetCbSize() {
	mii.cbSize = uint32(unsafe.Sizeof(*mii))
}

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

// [MSG] struct.
//
// [MSG]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type MSG struct {
	HWnd   HWND
	Msg    co.WM
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

// [NONCLIENTMETRICS] struct.
//
// ⚠️ You must call [NONCLIENTMETRICS.SetCbSize] to initialize the struct.
//
// # Example
//
//	var ncm win.NONCLIENTMETRICS
//	ncm.SetCbSize()
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

func (ncm *NONCLIENTMETRICS) CbSize() uint32 {
	return ncm.cbSize
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
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
	FErase      int32 // This is a BOOL value.
	RcPaint     RECT
	fRestore    int32 // This is a BOOL value.
	fIncUpdate  int32 // This is a BOOL value.
	rgbReserved [32]byte
}

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

func (pbs *POWERBROADCAST_SETTING) Data(i int) *uint8 {
	return &pbs.data[i]
}

// [RECT] struct.
//
// Basic rectangle structure, with left, top, right and bottom values.
//
// [RECT]: https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	Left, Top, Right, Bottom int32
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

// [TITLEBARINFO] struct.
//
// ⚠️ You must call [TITLEBARINFO.SetCbSize] to initialize the struct.
//
// # Example
//
//	var ti win.TITLEBARINFO
//	ti.SetCbSize()
//
// [TITLEBARINFO]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-titlebarinfo
type TITLEBARINFO struct {
	cbSize     uint32
	RcTitleBar RECT
	Rgstate    [utl.CCHILDREN_TITLEBAR + 1]uint32
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (ti *TITLEBARINFO) SetCbSize() {
	ti.cbSize = uint32(unsafe.Sizeof(*ti))
}

// [TITLEBARINFOEX] struct.
//
// ⚠️ You must call [TITLEBARINFOEX.SetCbSize] to initialize the struct.
//
// # Example
//
//	var tix win.TITLEBARINFOEX
//	tix.SetCbSize()
//
// [TITLEBARINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-titlebarinfoex
type TITLEBARINFOEX struct {
	cbSize     uint32
	RcTitleBar RECT
	Rgstate    [utl.CCHILDREN_TITLEBAR + 1]uint32
	Rgrect     [utl.CCHILDREN_TITLEBAR + 1]RECT
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (tix *TITLEBARINFOEX) SetCbSize() {
	tix.cbSize = uint32(unsafe.Sizeof(*tix))
}

// [WINDOWPLACEMENT] struct.
//
// ⚠️ You must call [WINDOWPLACEMENT.SetLength] to initialize the struct.
//
// # Example
//
//	var wp win.WINDOWPLACEMENT
//	wp.SetLength()
//
// [WINDOWPLACEMENT]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-windowplacement
type WINDOWPLACEMENT struct {
	length           uint32
	Flags            co.WPF
	ShowCmd          co.SW
	PtMinPosition    POINT
	PtMaxPosition    POINT
	RcNormalPosition RECT
}

// Sets the length field to the size of the struct, correctly initializing it.
func (wp *WINDOWPLACEMENT) SetLength() {
	wp.length = uint32(unsafe.Sizeof(*wp))
}

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
// ⚠️ You must call [WNDCLASSEX.SetCbSize] to initialize the struct.
//
// # Example
//
//	var wcx win.WNDCLASSEX
//	wcx.SetCbSize()
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

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (wcx *WNDCLASSEX) SetCbSize() {
	wcx.cbSize = uint32(unsafe.Sizeof(*wcx))
}

// First message [parameter].
//
// [parameter]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

// [MAKEWPARAM] macro.
//
// [MAKEWPARAM]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makewparam
func MAKEWPARAM(lo, hi uint16) WPARAM {
	return WPARAM(MAKELONG(lo, hi))
}

func (wp WPARAM) LoWord() uint16 {
	return LOWORD(uint32(wp))
}
func (wp WPARAM) HiWord() uint16 {
	return HIWORD(uint32(wp))
}
