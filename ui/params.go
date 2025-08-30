//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
)

// Raw [message] parameters to any message: [WPARAM] and [LPARAM].
//
// [message]: https://learn.microsoft.com/en-us/windows/win32/learnwin32/window-messages
// [WPARAM]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
// [LPARAM]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type Wm struct {
	Msg    co.WM      // Window message ID.
	WParam win.WPARAM // First parameter.
	LParam win.LPARAM // Second parameter.
}

// [WM_ACTIVATE] parameters.
//
// [WM_ACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-activate
type WmActivate struct{ Raw Wm }

func (p WmActivate) Event() co.WA                         { return co.WA(p.Raw.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                    { return p.Raw.WParam.HiWord() != 0 }
func (p WmActivate) HwndActivatedOrDeactivated() win.HWND { return win.HWND(p.Raw.LParam) }

// [WM_ACTIVATEAPP] parameters.
//
// [WM_ACTIVATEAPP]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-activateapp
type WmActivateApp struct{ Raw Wm }

func (p WmActivateApp) IsBeingActivated() bool { return p.Raw.WParam != 0 }
func (p WmActivateApp) ThreadId() uint32       { return uint32(p.Raw.LParam) }

// [WM_APPCOMMAND] parameters.
//
// [WM_APPCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
type WmAppCommand struct{ Raw Wm }

func (p WmAppCommand) OwnerWindow() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND {
	return co.APPCOMMAND(p.Raw.LParam.HiWord() &^ 0xf000)
}
func (p WmAppCommand) UDevice() co.FAPPCOMMAND { return co.FAPPCOMMAND(p.Raw.LParam.HiWord() & 0xf000) }
func (p WmAppCommand) Keys() co.MK             { return co.MK(p.Raw.LParam.LoWord()) }

// [WM_ASKCBFORMATNAME] parameters.
//
// [WM_ASKCBFORMATNAME]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-askcbformatname
type WmAskCbFormatName struct{ Raw Wm }

func (p WmAskCbFormatName) Buffer() []uint16 {
	return unsafe.Slice((*uint16)(unsafe.Pointer(p.Raw.LParam)), p.Raw.WParam)
}

// [WM_CAPTURECHANGED] parameters.
//
// [WM_CAPTURECHANGED]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-capturechanged
type WmCaptureChanged struct{ Raw Wm }

func (p WmCaptureChanged) HwndGainingMouse() win.HWND { return win.HWND(p.Raw.LParam) }

// [WM_CHANGECBCHAIN] parameters.
//
// [WM_CHANGECBCHAIN]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-changecbchain
type WmChangeCbChain struct{ Raw Wm }

func (p WmChangeCbChain) WindowBeingRemoved() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmChangeCbChain) NextWindow() win.HWND         { return win.HWND(p.Raw.LParam) }
func (p WmChangeCbChain) IsLastWindow() bool           { return p.Raw.LParam == 0 }

// Parameters for:
//   - [WM_CHAR]
//   - [WM_DEADCHAR]
//   - [WM_SYSCHAR]
//   - [WM_SYSDEADCHAR]
//
// [WM_CHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-char
// [WM_DEADCHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
// [WM_SYSCHAR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-syschar
// [WM_SYSDEADCHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-sysdeadchar
type WmChar struct{ Raw Wm }

func (p WmChar) CharCode() rune      { return rune(p.Raw.WParam) }
func (p WmChar) RepeatCount() uint   { return uint(p.Raw.LParam.LoWord()) }
func (p WmChar) ScanCode() uint8     { return win.LOBYTE(p.Raw.LParam.HiWord()) }
func (p WmChar) IsExtendedKey() bool { return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 0) }
func (p WmChar) HasAltKey() bool     { return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 5) }
func (p WmChar) IsKeyDownBeforeSend() bool {
	return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 6)
}
func (p WmChar) IsKeyBeingReleased() bool {
	return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 7)
}

// [WM_CHARTOITEM] parameters.
//
// [WM_CHARTOITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-chartoitem
type WmCharToItem struct{ Raw Wm }

func (p WmCharToItem) CharCode() rune        { return rune(p.Raw.WParam.LoWord()) }
func (p WmCharToItem) CurrentCaretPos() int  { return int(p.Raw.WParam.HiWord()) }
func (p WmCharToItem) HwndListBox() win.HWND { return win.HWND(p.Raw.LParam) }

// [WM_COMMAND] parameters.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
type WmCommand struct{ Raw Wm }

func (p WmCommand) IsFromMenu() bool        { return p.Raw.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool { return p.Raw.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool     { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() int             { return p.ControlId() }
func (p WmCommand) AcceleratorId() int      { return p.ControlId() }
func (p WmCommand) ControlId() int          { return int(p.Raw.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() int   { return int(p.Raw.WParam.HiWord()) }
func (p WmCommand) ControlHwnd() win.HWND   { return win.HWND(p.Raw.LParam) }

// [WM_COMPAREITEM] parameters.
//
// [WM_COMPAREITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-compareitem
type WmCompareItem struct{ Raw Wm }

func (p WmCompareItem) ControlId() int { return int(p.Raw.WParam) }
func (p WmCompareItem) CompareItemStruct() *win.COMPAREITEMSTRUCT {
	return (*win.COMPAREITEMSTRUCT)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_CONTEXTMENU] parameters.
//
// [WM_CONTEXTMENU]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-contextmenu
type WmContextMenu struct{ Raw Wm }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.Raw.LParam.MakePoint() }

// [WM_COPYDATA] parameters.
//
// [WM_COPYDATA]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-copydata
type WmCopyData struct{ Raw Wm }

func (p WmCopyData) HwndPassingData() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmCopyData) CopyDataStruct() *win.COPYDATASTRUCT {
	return (*win.COPYDATASTRUCT)(unsafe.Pointer(p.Raw.LParam))
}

// Parameters for:
//   - [WM_CREATE]
//   - [WM_NCCREATE]
//
// Sent only to windows created with [CreateWindowEx]; dialog windows will
// receive [WM_INITDIALOG] instead.
//
// [WM_CREATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-create
// [WM_NCCREATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-nccreate
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
// [WM_INITDIALOG]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-initdialog
type WmCreate struct{ Raw Wm }

func (p WmCreate) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.Raw.LParam))
}

// Parameters for:
//   - [WM_CTLCOLORBTN]
//   - [WM_CTLCOLORDLG]
//   - [WM_CTLCOLOREDIT]
//   - [WM_CTLCOLORLISTBOX]
//   - [WM_CTLCOLORSCROLLBAR]
//   - [WM_CTLCOLORSTATIC]
//
// [WM_CTLCOLORBTN]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorbtn
// [WM_CTLCOLORDLG]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-ctlcolordlg
// [WM_CTLCOLOREDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcoloredit
// [WM_CTLCOLORLISTBOX]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorlistbox
// [WM_CTLCOLORSCROLLBAR]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorscrollbar
// [WM_CTLCOLORSTATIC]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
type WmCtlColor struct{ Raw Wm }

func (p WmCtlColor) Hdc() win.HDC          { return win.HDC(p.Raw.WParam) }
func (p WmCtlColor) HwndControl() win.HWND { return win.HWND(p.Raw.LParam) }

// [WM_DELETEITEM] parameters.
//
// [WM_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-deleteitem
type WmDeleteItem struct{ Raw Wm }

func (p WmDeleteItem) ControlId() int { return int(p.Raw.WParam) }
func (p WmDeleteItem) DeleteItemStruct() *win.DELETEITEMSTRUCT {
	return (*win.DELETEITEMSTRUCT)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_DEVICECHANGE] parameters.
//
// [WM_DEVICECHANGE]: https://learn.microsoft.com/en-us/windows/win32/devio/wm-devicechange
type WmDeviceChange struct{ Raw Wm }

func (p WmDeviceChange) Event() co.DBT             { return co.DBT(p.Raw.WParam) }
func (p WmDeviceChange) EventData() unsafe.Pointer { return unsafe.Pointer(p.Raw.LParam) }

// [WM_DEVMODECHANGE] parameters.
//
// [WM_DEVMODECHANGE]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-devmodechange
type WmDevModeChange struct{ Raw Wm }

func (p WmDevModeChange) DeviceName() string {
	return wstr.DecodePtr((*uint16)(unsafe.Pointer(p.Raw.LParam)))
}

// [WM_DISPLAYCHANGE] parameters.
//
// [WM_DISPLAYCHANGE]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-displaychange
type WmDisplayChange struct{ Raw Wm }

func (p WmDisplayChange) BitsPerPixel() int { return int(p.Raw.WParam) }
func (p WmDisplayChange) Size() win.SIZE    { return p.Raw.LParam.MakeSize() }

// [WM_DRAWITEM] parameters.
//
// [WM_DRAWITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-drawitem
type WmDrawItem struct{ Raw Wm }

func (p WmDrawItem) ControlId() int   { return int(p.Raw.WParam) }
func (p WmDrawItem) IsFromMenu() bool { return p.Raw.WParam == 0 }
func (p WmDrawItem) DrawItemStruct() *win.DRAWITEMSTRUCT {
	return (*win.DRAWITEMSTRUCT)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_DROPFILES] parameters.
//
// [WM_DROPFILES]: https://learn.microsoft.com/en-us/windows/win32/shell/wm-dropfiles
type WmDropFiles struct{ Raw Wm }

func (p WmDropFiles) HDrop() win.HDROP { return win.HDROP(p.Raw.WParam) }

// [WM_DWMCOLORIZATIONCOLORCHANGED] parameters.
//
// [WM_DWMCOLORIZATIONCOLORCHANGED]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmcolorizationcolorchanged
type WmDwmColorizationColorChanged struct{ Raw Wm }

func (p WmDwmColorizationColorChanged) Color() win.COLORREF      { return win.COLORREF(p.Raw.WParam) }
func (p WmDwmColorizationColorChanged) BlendedWithOpacity() bool { return p.Raw.LParam != 0 }

// [WM_DWMNCRENDERINGCHANGED] parameters.
//
// [WM_DWMNCRENDERINGCHANGED]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmncrenderingchanged
type WmDwmNcRenderingChanged struct{ Raw Wm }

func (p WmDwmNcRenderingChanged) IsEnabled() bool { return p.Raw.WParam != 0 }

// [WM_DWMSENDICONICTHUMBNAIL] parameters.
//
// [WM_DWMSENDICONICTHUMBNAIL]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmsendiconicthumbnail
type WmDwmSendIconicThumbnail struct{ Raw Wm }

func (p WmDwmSendIconicThumbnail) MaxCoords() win.POINT { return p.Raw.LParam.MakePoint() }

// [WM_DWMWINDOWMAXIMIZEDCHANGE] parameters.
//
// [WM_DWMWINDOWMAXIMIZEDCHANGE]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmwindowmaximizedchange
type WmDwmWindowMaximizedChange struct{ Raw Wm }

func (p WmDwmWindowMaximizedChange) IsMaximized() bool { return p.Raw.WParam != 0 }

// [WM_ENABLE] parameters.
//
// [WM_ENABLE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-enable
type WmEnable struct{ Raw Wm }

func (p WmEnable) HasBeenEnabled() bool { return p.Raw.WParam != 0 }

// [WM_ENDSESSION] parameters.
//
// [WM_ENDSESSION]: https://learn.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
type WmEndSession struct{ Raw Wm }

func (p WmEndSession) IsSessionBeingEnded() bool { return p.Raw.WParam != 0 }
func (p WmEndSession) Event() co.ENDSESSION      { return co.ENDSESSION(p.Raw.LParam) }

// [WM_ENTERIDLE] parameters.
//
// [WM_ENTERIDLE]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
type WmEnterIdle struct{ Raw Wm }

func (p WmEnterIdle) Displayed() co.MSGF       { return co.MSGF(p.Raw.WParam) }
func (p WmEnterIdle) DialogOrWindow() win.HWND { return win.HWND(p.Raw.LParam) }

// [WM_ENTERMENULOOP] parameters.
//
// [WM_ENTERMENULOOP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-entermenuloop
type WmEnterMenuLoop struct{ Raw Wm }

func (p WmEnterMenuLoop) IsTrackPopupMenu() bool { return p.Raw.WParam != 0 }

// [WM_ERASEBKGND] parameters.
//
// [WM_ERASEBKGND]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-erasebkgnd
type WmEraseBkgnd struct{ Raw Wm }

func (p WmEraseBkgnd) Hdc() win.HDC { return win.HDC(p.Raw.WParam) }

// [WM_EXITMENULOOP] parameters.
//
// [WM_EXITMENULOOP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-exitmenuloop
type WmExitMenuLoop struct{ Raw Wm }

func (p WmExitMenuLoop) IsShortcutMenu() bool { return p.Raw.WParam != 0 }

// [WM_GETDLGCODE] parameters.
//
// [WM_GETDLGCODE]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
type WmGetDlgCode struct{ Raw Wm }

func (p WmGetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.Raw.WParam) }
func (p WmGetDlgCode) IsQuery() bool         { return p.Raw.LParam == 0 }
func (p WmGetDlgCode) Message() *win.MSG     { return (*win.MSG)(unsafe.Pointer(p.Raw.LParam)) }
func (p WmGetDlgCode) HasAlt() bool          { return (win.GetAsyncKeyState(co.VK_MENU) & 0x8000) != 0 }
func (p WmGetDlgCode) HasCtrl() bool         { return (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0 }
func (p WmGetDlgCode) HasShift() bool        { return (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0 }

// [WM_GETICON] parameters.
//
// [WM_GETICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-geticon
type WmGetIcon struct{ Raw Wm }

func (p WmGetIcon) Type() co.ICON_SZ { return co.ICON_SZ(p.Raw.WParam) }
func (p WmGetIcon) Dpi() uint32      { return uint32(p.Raw.LParam) }

// [WM_GETMINMAXINFO] parameters.
//
// [WM_GETMINMAXINFO]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-getminmaxinfo
type WmGetMinMaxInfo struct{ Raw Wm }

func (p WmGetMinMaxInfo) Info() *win.MINMAXINFO {
	return (*win.MINMAXINFO)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_GETTEXT] parameters.
//
// [WM_GETTEXT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-gettext
type WmGetText struct{ Raw Wm }

func (p WmGetText) MaxChars() uint  { return uint(p.Raw.WParam) }
func (p WmGetText) Buffer() *uint16 { return (*uint16)(unsafe.Pointer(p.Raw.LParam)) }

// [WM_GETTITLEBARINFOEX] parameters.
//
// [WM_GETTITLEBARINFOEX]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-gettitlebarinfoex
type WmGetTitleBarInfoEx struct{ Raw Wm }

func (p WmGetTitleBarInfoEx) Info() *win.TITLEBARINFOEX {
	return (*win.TITLEBARINFOEX)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_HELP] parameters.
//
// [WM_HELP]: https://learn.microsoft.com/en-us/windows/win32/shell/wm-help
type WmHelp struct{ Raw Wm }

func (p WmHelp) Info() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.Raw.LParam)) }

// [WM_HOTKEY] parameters.
//
// [WM_HOTKEY]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
type WmHotKey struct{ Raw Wm }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.Raw.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.Raw.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.Raw.LParam.HiWord()) }

// [WM_INITDIALOG] parameters.
//
// Sent only to dialog windows; those created with [CreateWindowEx] will receive
// [WM_CREATE] instead.
//
// [WM_INITDIALOG]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-initdialog
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
// [WM_CREATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-create
type WmInitDialog struct{ Raw Wm }

func (p WmInitDialog) HwndFocused() win.HWND { return win.HWND(p.Raw.WParam) }

// [WM_INITMENUPOPUP] parameters.
//
// [WM_INITMENUPOPUP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
type WmInitMenuPopup struct{ Raw Wm }

func (p WmInitMenuPopup) HMenu() win.HMENU   { return win.HMENU(p.Raw.WParam) }
func (p WmInitMenuPopup) Pos() int           { return int(p.Raw.LParam.LoWord()) }
func (p WmInitMenuPopup) IsWindowMenu() bool { return p.Raw.LParam.HiWord() != 0 }

// Parameters for:
//   - [WM_KEYDOWN]
//   - [WM_KEYUP]
//   - [WM_SYSKEYDOWN]
//   - [WM_SYSKEYUP]
//
// [WM_KEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-keydown
// [WM_KEYUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-keyup
// [WM_SYSKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-syskeydown
// [WM_SYSKEYUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-syskeyup
type WmKey struct{ Raw Wm }

func (p WmKey) VirtualKeyCode() co.VK { return co.VK(p.Raw.WParam) }
func (p WmKey) RepeatCount() uint     { return uint(p.Raw.LParam.LoWord()) }
func (p WmKey) ScanCode() uint8       { return win.LOBYTE(p.Raw.LParam.HiWord()) }
func (p WmKey) IsExtendedKey() bool   { return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 0) }
func (p WmKey) HasAltKey() bool       { return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 5) }
func (p WmKey) IsKeyDownBeforeSend() bool {
	return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 6)
}
func (p WmKey) IsReleasingKey() bool { return utl.BitIsSet(win.HIBYTE(p.Raw.LParam.HiWord()), 7) }

// [WM_KILLFOCUS] parameters.
//
// [WM_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-killfocus
type WmKillFocus struct{ Raw Wm }

func (p WmKillFocus) HwndReceivingFocus() win.HWND { return win.HWND(p.Raw.LParam) }

// Parameters for:
//   - [WM_MENUCOMMAND]
//   - [WM_MENUDRAG]
//   - [WM_MENURBUTTONUP]
//
// [WM_MENUCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
// [WM_MENUDRAG]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
// [WM_MENURBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menurbuttonup
type WmMenu struct{ Raw Wm }

func (p WmMenu) ItemIndex() uint  { return uint(p.Raw.WParam) }
func (p WmMenu) Hmenu() win.HMENU { return win.HMENU(p.Raw.LParam) }

// [WM_MENUCHAR] parameters.
//
// [WM_MENUCHAR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
type WmMenuChar struct{ Raw Wm }

func (p WmMenuChar) CharCode() rune          { return rune(p.Raw.WParam.LoWord()) }
func (p WmMenuChar) ActiveMenuType() co.MFMC { return co.MFMC(p.Raw.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU   { return win.HMENU(p.Raw.LParam) }

// [WM_MENUGETOBJECT] parameters.
//
// [WM_MENUGETOBJECT]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
type WmMenuGetObject struct{ Raw Wm }

func (p WmMenuGetObject) Info() *win.MENUGETOBJECTINFO {
	return (*win.MENUGETOBJECTINFO)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_MENUSELECT] parameters.
//
// [WM_MENUSELECT]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
type WmMenuSelect struct{ Raw Wm }

func (p WmMenuSelect) Item() int        { return int(p.Raw.WParam.LoWord()) }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.Raw.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.Raw.LParam) }

// Parameters for:
//   - [WM_LBUTTONDBLCLK]
//   - [WM_LBUTTONDOWN]
//   - [WM_LBUTTONUP]
//   - [WM_MBUTTONDBLCLK]
//   - [WM_MBUTTONDOWN]
//   - [WM_MBUTTONUP]
//   - [WM_MOUSEHOVER]
//   - [WM_MOUSEMOVE]
//   - [WM_RBUTTONDBLCLK]
//   - [WM_RBUTTONDOWN]
//   - [WM_RBUTTONUP]
//   - [WM_XBUTTONDBLCLK]
//   - [WM_XBUTTONDOWN]
//   - [WM_XBUTTONUP]
//
// [WM_LBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondblclk
// [WM_LBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
// [WM_LBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttonup
// [WM_MBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondblclk
// [WM_MBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondown
// [WM_MBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttonup
// [WM_MOUSEHOVER]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mousehover
// [WM_MOUSEMOVE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
// [WM_RBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondblclk
// [WM_RBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondown
// [WM_RBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttonup
// [WM_XBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondblclk
// [WM_XBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondown
// [WM_XBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttonup
type WmMouse struct{ Raw Wm }

func (p WmMouse) VirtualKeys() co.MK { return co.MK(p.Raw.WParam.LoWord()) }
func (p WmMouse) HasCtrl() bool      { return (p.VirtualKeys() & co.MK_CONTROL) != 0 }
func (p WmMouse) HasShift() bool     { return (p.VirtualKeys() & co.MK_SHIFT) != 0 }
func (p WmMouse) IsLeftBtn() bool    { return (p.VirtualKeys() & co.MK_LBUTTON) != 0 }
func (p WmMouse) IsMiddleBtn() bool  { return (p.VirtualKeys() & co.MK_MBUTTON) != 0 }
func (p WmMouse) IsRightBtn() bool   { return (p.VirtualKeys() & co.MK_RBUTTON) != 0 }
func (p WmMouse) IsXBtn1() bool      { return (p.VirtualKeys() & co.MK_XBUTTON1) != 0 }
func (p WmMouse) IsXBtn2() bool      { return (p.VirtualKeys() & co.MK_XBUTTON2) != 0 }
func (p WmMouse) Pos() win.POINT     { return p.Raw.LParam.MakePoint() }

// [WM_MOVE] parameters.
//
// [WM_MOVE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-move
type WmMove struct{ Raw Wm }

func (p WmMove) ClientAreaPos() win.POINT { return p.Raw.LParam.MakePoint() }

// [WM_MOVING] parameters.
//
// [WM_MOVING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-moving
type WmMoving struct{ Raw Wm }

func (p WmMoving) WindowPos() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Raw.LParam)) }

// [WM_NCACTIVATE] parameters.
//
// [WM_NCACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-ncactivate
type WmNcActivate struct{ Raw Wm }

func (p WmNcActivate) IsActive() bool            { return p.Raw.WParam != 0 }
func (p WmNcActivate) IsVisualStyleActive() bool { return p.Raw.LParam == 0 }
func (p WmNcActivate) UpdatedRegion() win.HRGN   { return win.HRGN(p.Raw.LParam) }

// [WM_NCCALCSIZE] parameters.
//
// [WM_NCCALCSIZE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
type WmNcCalcSize struct{ Raw Wm }

func (p WmNcCalcSize) ShouldIndicateValidPart() bool { return p.Raw.WParam != 0 }
func (p WmNcCalcSize) NcCalcSizeParams() *win.NCCALCSIZE_PARAMS {
	return (*win.NCCALCSIZE_PARAMS)(unsafe.Pointer(p.Raw.LParam))
}
func (p WmNcCalcSize) Rect() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Raw.LParam)) }

// [WM_NCHITTEST] parameters.
//
// [WM_NCHITTEST]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
type WmNcHitTest struct{ Raw Wm }

func (p WmNcHitTest) CursorPos() win.POINT { return p.Raw.LParam.MakePoint() }

// Parameters for:
//   - [WM_NCLBUTTONDBLCLK]
//   - [WM_NCLBUTTONDOWN]
//   - [WM_NCLBUTTONUP]
//   - [WM_NCMBUTTONDBLCLK]
//   - [WM_NCMBUTTONDOWN]
//   - [WM_NCMBUTTONUP]
//   - [WM_NCMOUSEHOVER]
//   - [WM_NCMOUSEMOVE]
//   - [WM_NCRBUTTONDBLCLK]
//   - [WM_NCRBUTTONDOWN]
//   - [WM_NCRBUTTONUP]
//
// [WM_NCLBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondblclk
// [WM_NCLBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondown
// [WM_NCLBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttonup
// [WM_NCMBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondblclk
// [WM_NCMBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondown
// [WM_NCMBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttonup
// [WM_NCMOUSEHOVER]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousehover
// [WM_NCMOUSEMOVE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousemove
// [WM_NCRBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondblclk
// [WM_NCRBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondown
// [WM_NCRBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttonup
type WmNcMouse struct{ Raw Wm }

func (p WmNcMouse) HitTest() co.HT { return co.HT(p.Raw.WParam) }
func (p WmNcMouse) Pos() win.POINT { return p.Raw.LParam.MakePoint() }

// Parameters for:
//   - [WM_NCXBUTTONDBLCLK]
//   - [WM_NCXBUTTONDOWN]
//   - [WM_NCXBUTTONUP]
//
// [WM_NCXBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondblclk
// [WM_NCXBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondown
// [WM_NCXBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttonup\
type WmNcMouseX struct{ Raw Wm }

func (p WmNcMouseX) HitTest() co.HT { return co.HT(p.Raw.WParam.LoWord()) }
func (p WmNcMouseX) IsXBtn1() bool  { return p.Raw.WParam.HiWord() == 0x0001 }
func (p WmNcMouseX) IsXBtn2() bool  { return p.Raw.WParam.HiWord() == 0x0002 }
func (p WmNcMouseX) Pos() win.POINT { return p.Raw.LParam.MakePoint() }

// [WM_NEXTDLGCTL] parameters.
//
// [WM_NEXTDLGCTL]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-nextdlgctl
type WmNextDlgCtl struct{ Raw Wm }

func (p WmNextDlgCtl) HwndFocus() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmNextDlgCtl) IsHwnd() bool        { return p.Raw.LParam.LoWord() != 0 }

// [WM_NEXTMENU] parameters.
//
// [WM_NEXTMENU]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-nextmenu
type WmNextMenu struct{ Raw Wm }

func (p WmNextMenu) VirtualKeyCode() co.VK { return co.VK(p.Raw.WParam) }
func (p WmNextMenu) MdiNextMenu() *win.MDINEXTMENU {
	return (*win.MDINEXTMENU)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_NCPAINT] parameters.
//
// [WM_NCPAINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-ncpaint
type WmNcPaint struct{ Raw Wm }

func (p WmNcPaint) UpdatedHrgn() win.HRGN { return win.HRGN(p.Raw.WParam) }

// [WM_PAINTCLIPBOARD] parameters.
//
// [WM_PAINTCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-paintclipboard
type WmPaintClipboard struct{ Raw Wm }

func (p WmPaintClipboard) CbViewerWindow() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmPaintClipboard) PaintStruct() *win.PAINTSTRUCT {
	return (*win.PAINTSTRUCT)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_PARENTNOTIFY] parameters.
//
// [WM_PARENTNOTIFY]: https://learn.microsoft.com/en-us/windows/win32/inputmsg/wm-parentnotify
type WmParentNotify struct{ Raw Wm }

func (p WmParentNotify) Event() co.WMPN    { return co.WMPN(p.Raw.WParam.LoWord()) }
func (p WmParentNotify) ChildId() uint16   { return p.Raw.WParam.HiWord() }
func (p WmParentNotify) Coords() win.POINT { return p.Raw.LParam.MakePoint() }

// [WM_POWERBROADCAST] parameters.
//
// [WM_POWERBROADCAST]: https://learn.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
type WmPowerBroadcast struct{ Raw Wm }

func (p WmPowerBroadcast) Event() co.PBT { return co.PBT(p.Raw.WParam) }
func (p WmPowerBroadcast) PowerBroadcastSetting() *win.POWERBROADCAST_SETTING {
	if p.Event() == co.PBT_POWERSETTINGCHANGE {
		return (*win.POWERBROADCAST_SETTING)(unsafe.Pointer(p.Raw.LParam))
	}
	return nil
}

// [WM_PRINT] parameters.
//
// [WM_PRINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-print
type WmPrint struct{ Raw Wm }

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.Raw.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.Raw.LParam) }

// [WM_RENDERFORMAT] parameters.
//
// [WM_RENDERFORMAT]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-renderformat
type WmRenderFormat struct{ Raw Wm }

func (p WmRenderFormat) ClipboardFormat() co.CF { return co.CF(p.Raw.WParam) }

// Parameters for:
//   - [WM_VSCROLL]
//   - [WM_HSCROLL]
//
// [WM_HSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-hscroll
// [WM_VSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-vscroll
type WmScroll struct{ Raw Wm }

func (p WmScroll) ScrollBoxPos() int       { return int(p.Raw.WParam.HiWord()) }
func (p WmScroll) Request() co.SB_REQ      { return co.SB_REQ(p.Raw.WParam.LoWord()) }
func (p WmScroll) HwndScrollbar() win.HWND { return win.HWND(p.Raw.LParam) }

// [WM_SETCURSOR] parameters.
//
// [WM_SETCURSOR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-setcursor
type WmSetCursor struct{ Raw Wm }

func (p WmSetCursor) Hwnd() win.HWND       { return win.HWND(p.Raw.WParam) }
func (p WmSetCursor) CursorPos() win.POINT { return p.Raw.LParam.MakePoint() }
func (p WmSetCursor) SrcMsg() co.WM        { return co.WM(p.Raw.LParam) }

// [WM_SETFOCUS] parameters.
//
// [WM_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-setfocus
type WmSetFocus struct{ Raw Wm }

func (p WmSetFocus) HwndLosingFocus() win.HWND { return win.HWND(p.Raw.WParam) }

// [WM_SETFONT] parameters.
//
// [WM_SETFONT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-setfont
type WmSetFont struct{ Raw Wm }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.Raw.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.Raw.LParam == 1 }

// [WM_SETICON] parameters.
//
// [WM_SETICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
type WmSetIcon struct{ Raw Wm }

func (p WmSetIcon) Size() co.ICON_SZ { return co.ICON_SZ(p.Raw.WParam) }
func (p WmSetIcon) Hicon() win.HICON { return win.HICON(p.Raw.LParam) }

// [WM_SETREDRAW] parameters.
//
// [WM_SETREDRAW]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-setredraw
type WmSetRedraw struct{ Raw Wm }

func (p WmSetRedraw) CanRedraw() bool { return p.Raw.WParam != 0 }

// [WM_SETTEXT] parameters.
//
// [WM_SETTEXT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-settext
type WmSetText struct{ Raw Wm }

func (p WmSetText) Text() *uint16 { return (*uint16)(unsafe.Pointer(p.Raw.LParam)) }

// [WM_SHOWWINDOW] parameters.
//
// [WM_SHOWWINDOW]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-showwindow
type WmShowWindow struct{ Raw Wm }

func (p WmShowWindow) IsBeingShown() bool { return p.Raw.WParam != 0 }
func (p WmShowWindow) Status() co.SWS     { return co.SWS(p.Raw.LParam) }

// [WM_SIZE] parameters.
//
// [WM_SIZE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-size
type WmSize struct{ Raw Wm }

func (p WmSize) Request() co.SIZE_REQ     { return co.SIZE_REQ(p.Raw.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.Raw.LParam.MakeSize() }

// [WM_SIZECLIPBOARD] parameters.
//
// [WM_SIZECLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-sizeclipboard
type WmSizeClipboard struct{ Raw Wm }

func (p WmSizeClipboard) CbViewerWindow() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmSizeClipboard) NewDimensions() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Raw.LParam)) }

// [WM_SIZING] parameters.
//
// [WM_SIZING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-sizing
type WmSizing struct{ Raw Wm }

func (p WmSizing) WindowEdge() co.WMSZ { return co.WMSZ(p.Raw.WParam) }
func (p WmSizing) DragRect() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Raw.LParam)) }

// Parameters for:
//   - [WM_STYLECHANGED]
//   - [WM_STYLECHANGING]
//
// [WM_STYLECHANGED]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-stylechanged
// [WM_STYLECHANGING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-stylechanging
type WmStyles struct{ Raw Wm }

func (p WmStyles) StylesWs() *win.STYLESTRUCT_WS {
	if co.GWLP(p.Raw.WParam) == co.GWLP_STYLE {
		return (*win.STYLESTRUCT_WS)(unsafe.Pointer(p.Raw.LParam))
	}
	return nil
}
func (p WmStyles) StylesWsEx() *win.STYLESTRUCT_WSEX {
	if co.GWLP(p.Raw.WParam) == co.GWLP_EXSTYLE {
		return (*win.STYLESTRUCT_WSEX)(unsafe.Pointer(p.Raw.LParam))
	}
	return nil
}

// [WM_SYSCOMMAND] parameters.
//
// [WM_SYSCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
type WmSysCommand struct{ Raw Wm }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.Raw.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.Raw.LParam.MakePoint() }

// [WM_UNINITMENUPOPUP] parameters.
//
// [WM_UNINITMENUPOPUP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-uninitmenupopup
type WmUnInitMenuPopup struct{ Raw Wm }

func (p WmUnInitMenuPopup) Hmenu() win.HMENU { return win.HMENU(p.Raw.WParam) }

// Parameters for:
//   - [WM_HSCROLLCLIPBOARD]
//   - [WM_VSCROLLCLIPBOARD]
//
// [WM_HSCROLLCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-hscrollclipboard
// [WM_VSCROLLCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-vscrollclipboard
type WmScrollClipboard struct{ Raw Wm }

func (p WmScrollClipboard) HwndClipboard() win.HWND { return win.HWND(p.Raw.WParam) }
func (p WmScrollClipboard) Request() co.SB_REQ      { return co.SB_REQ(p.Raw.WParam.LoWord()) }
func (p WmScrollClipboard) ScrollBoxPos() int       { return int(p.Raw.WParam.HiWord()) }

// Parameters for:
//   - [WM_WINDOWPOSCHANGED]
//   - [WM_WINDOWPOSCHANGING]
//
// [WM_WINDOWPOSCHANGED]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-windowposchanged
// [WM_WINDOWPOSCHANGING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-windowposchanging
type WmWindowPos struct{ Raw Wm }

func (p WmWindowPos) WindowPos() *win.WINDOWPOS {
	return (*win.WINDOWPOS)(unsafe.Pointer(p.Raw.LParam))
}

// [WM_WTSSESSION_CHANGE] parameters.
//
// [WM_WTSSESSION_CHANGE]: https://learn.microsoft.com/en-us/windows/win32/termserv/wm-wtssession-change
type WmWtsSessionChange struct{ Raw Wm }

func (p WmWtsSessionChange) State() co.WTS     { return co.WTS(p.Raw.WParam) }
func (p WmWtsSessionChange) SessionId() uint32 { return uint32(p.Raw.LParam) }
