package wm

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type Activate struct{ Msg Any }

func (p Activate) Event() co.WA                         { return co.WA(p.Msg.WParam.Lo16()) }
func (p Activate) IsMinimized() bool                    { return p.Msg.WParam.Hi16() != 0 }
func (p Activate) HwndActivatedOrDeactivated() win.HWND { return win.HWND(p.Msg.LParam) }

type ActivateApp struct{ Msg Any }

func (p ActivateApp) IsBeingActivated() bool { return p.Msg.WParam != 0 }
func (p ActivateApp) ThreadId() uint32       { return uint32(p.Msg.LParam) }

type AppCommand struct{ Msg Any }

func (p AppCommand) OwnerWindow() win.HWND     { return win.HWND(p.Msg.WParam) }
func (p AppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.Msg.LParam.Hi16() &^ 0xf000) }
func (p AppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.Msg.LParam.Hi16() & 0xf000) }
func (p AppCommand) Keys() co.MK               { return co.MK(p.Msg.LParam.Lo16()) }

type AskCbFormatName struct{ Msg Any }

func (p AskCbFormatName) BufferSize() int { return int(p.Msg.WParam) }
func (p AskCbFormatName) Buffer() *uint16 { return (*uint16)(unsafe.Pointer(p.Msg.LParam)) }

type CaptureChanged struct{ Msg Any }

func (p CaptureChanged) HwndGainingMouse() win.HWND { return win.HWND(p.Msg.LParam) }

type ChangeCbChain struct{ Msg Any }

func (p ChangeCbChain) WindowBeingRemoved() win.HWND { return win.HWND(p.Msg.WParam) }
func (p ChangeCbChain) NextWindow() win.HWND         { return win.HWND(p.Msg.LParam) }
func (p ChangeCbChain) IsLastWindow() bool           { return p.Msg.LParam == 0 }

type Char struct{ Msg Any }

func (p Char) CharCode() rune            { return rune(p.Msg.WParam) }
func (p Char) RepeatCount() uint         { return uint(p.Msg.LParam.Lo16()) }
func (p Char) ScanCode() uint            { return uint(p.Msg.LParam.Hi16Lo8()) }
func (p Char) IsExtendedKey() bool       { return (p.Msg.LParam.Hi16Hi8() & 0b0000_0001) != 0 }
func (p Char) HasAltKey() bool           { return (p.Msg.LParam.Hi16Hi8() & 0b0010_0000) != 0 }
func (p Char) IsKeyDownBeforeSend() bool { return (p.Msg.LParam.Hi16Hi8() & 0b0100_0000) != 0 }
func (p Char) IsKeyBeingReleased() bool  { return (p.Msg.LParam.Hi16Hi8() & 0b1000_0000) != 0 }

type CharToItem struct{ Msg Any }

func (p CharToItem) CharCode() rune        { return rune(p.Msg.WParam.Lo16()) }
func (p CharToItem) CurrentCaretPos() int  { return int(p.Msg.WParam.Hi16()) }
func (p CharToItem) HwndListBox() win.HWND { return win.HWND(p.Msg.LParam) }

type Command struct{ Msg Any }

func (p Command) IsFromMenu() bool        { return p.Msg.WParam.Hi16() == 0 }
func (p Command) IsFromAccelerator() bool { return p.Msg.WParam.Hi16() == 1 }
func (p Command) IsFromControl() bool     { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p Command) MenuId() int             { return p.ControlId() }
func (p Command) AcceleratorId() int      { return p.ControlId() }
func (p Command) ControlId() int          { return int(p.Msg.WParam.Lo16()) }
func (p Command) ControlNotifCode() int   { return int(p.Msg.WParam.Hi16()) }
func (p Command) ControlHwnd() win.HWND   { return win.HWND(p.Msg.LParam) }

type CompareItem struct{ Msg Any }

func (p CompareItem) ControlId() int { return int(p.Msg.WParam) }
func (p CompareItem) CompareItemStruct() *win.COMPAREITEMSTRUCT {
	return (*win.COMPAREITEMSTRUCT)(unsafe.Pointer(p.Msg.LParam))
}

type ContextMenu struct{ Msg Any }

func (p ContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.Msg.WParam) }
func (p ContextMenu) CursorPos() win.POINT         { return p.Msg.LParam.MakePoint() }

type CopyData struct{ Msg Any }

func (p CopyData) HwndPassingData() win.HWND { return win.HWND(p.Msg.WParam) }
func (p CopyData) CopyDataStruct() *win.COPYDATASTRUCT {
	return (*win.COPYDATASTRUCT)(unsafe.Pointer(p.Msg.LParam))
}

type Create struct{ Msg Any }

func (p Create) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.Msg.LParam))
}

type CtlColor struct{ Msg Any }

func (p CtlColor) Hdc() win.HDC          { return win.HDC(p.Msg.WParam) }
func (p CtlColor) HwndControl() win.HWND { return win.HWND(p.Msg.LParam) }

type DeleteItem struct{ Msg Any }

func (p DeleteItem) ControlId() int { return int(p.Msg.WParam) }
func (p DeleteItem) DeleteItemStruct() *win.DELETEITEMSTRUCT {
	return (*win.DELETEITEMSTRUCT)(unsafe.Pointer(p.Msg.LParam))
}

type DevModeChange struct{ Msg Any }

func (p DevModeChange) DeviceName() string {
	return win.Str.FromUint16Ptr((*uint16)(unsafe.Pointer(p.Msg.LParam)))
}

type DisplayChange struct{ Msg Any }

func (p DisplayChange) BitsPerPixel() int { return int(p.Msg.WParam) }
func (p DisplayChange) Size() win.SIZE    { return p.Msg.LParam.MakeSize() }

type DrawItem struct{ Msg Any }

func (p DrawItem) ControlId() int   { return int(p.Msg.WParam) }
func (p DrawItem) IsFromMenu() bool { return p.Msg.WParam == 0 }
func (p DrawItem) DrawItemStruct() *win.DRAWITEMSTRUCT {
	return (*win.DRAWITEMSTRUCT)(unsafe.Pointer(p.Msg.LParam))
}

type DropFiles struct{ Msg Any }

func (p DropFiles) Hdrop() win.HDROP { return win.HDROP(p.Msg.WParam) }

type DwmColorizationColorChanged struct{ Msg Any }

func (p DwmColorizationColorChanged) Color() win.COLORREF      { return win.COLORREF(p.Msg.WParam) }
func (p DwmColorizationColorChanged) BlendedWithOpacity() bool { return p.Msg.LParam != 0 }

type DwmNcRenderingChanged struct{ Msg Any }

func (p DwmNcRenderingChanged) IsEnabled() bool { return p.Msg.WParam != 0 }

type DwmSendIconicThumbnail struct{ Msg Any }

func (p DwmSendIconicThumbnail) MaxCoords() win.POINT { return p.Msg.LParam.MakePoint() }

type DwmWindowMaximizedChange struct{ Msg Any }

func (p DwmWindowMaximizedChange) IsMaximized() bool { return p.Msg.WParam != 0 }

type Enable struct{ Msg Any }

func (p Enable) HasBeenEnabled() bool { return p.Msg.WParam != 0 }

type EndSession struct{ Msg Any }

func (p EndSession) IsSessionBeingEnded() bool { return p.Msg.WParam != 0 }
func (p EndSession) Event() co.ENDSESSION      { return co.ENDSESSION(p.Msg.LParam) }

type EnterIdle struct{ Msg Any }

func (p EnterIdle) Displayed() co.MSGF       { return co.MSGF(p.Msg.WParam) }
func (p EnterIdle) DialogOrWindow() win.HWND { return win.HWND(p.Msg.LParam) }

type EnterMenuLoop struct{ Msg Any }

func (p EnterMenuLoop) IsTrackPopupMenu() bool { return p.Msg.WParam != 0 }

type EraseBkgnd struct{ Msg Any }

func (p EraseBkgnd) Hdc() win.HDC { return win.HDC(p.Msg.WParam) }

type ExitMenuLoop struct{ Msg Any }

func (p ExitMenuLoop) IsShortcutMenu() bool { return p.Msg.WParam != 0 }

type GetDlgCode struct{ Msg Any }

func (p GetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.Msg.WParam) }
func (p GetDlgCode) IsQuery() bool         { return p.Msg.LParam == 0 }
func (p GetDlgCode) Message() *win.MSG     { return (*win.MSG)(unsafe.Pointer(p.Msg.LParam)) }
func (p GetDlgCode) HasAlt() bool          { return (win.GetAsyncKeyState(co.VK_MENU) & 0x8000) != 0 }
func (p GetDlgCode) HasCtrl() bool         { return (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0 }
func (p GetDlgCode) HasShift() bool        { return (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0 }

type GetIcon struct{ Msg Any }

func (p GetIcon) Type() co.ICON_SZ { return co.ICON_SZ(p.Msg.WParam) }
func (p GetIcon) Dpi() uint32      { return uint32(p.Msg.LParam) }

type GetMinMaxInfo struct{ Msg Any }

func (p GetMinMaxInfo) Info() *win.MINMAXINFO { return (*win.MINMAXINFO)(unsafe.Pointer(p.Msg.LParam)) }

type GetTitleBarInfoEx struct{ Msg Any }

func (p GetTitleBarInfoEx) Info() *win.TITLEBARINFOEX {
	return (*win.TITLEBARINFOEX)(unsafe.Pointer(p.Msg.LParam))
}

type Help struct{ Msg Any }

func (p Help) Info() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.Msg.LParam)) }

type HotKey struct{ Msg Any }

func (p HotKey) HotKey() co.IDHOT      { return co.IDHOT(p.Msg.WParam) }
func (p HotKey) OtherKeys() co.MOD     { return co.MOD(p.Msg.LParam.Lo16()) }
func (p HotKey) VirtualKeyCode() co.VK { return co.VK(p.Msg.LParam.Hi16()) }

type HScroll struct{ Msg Any }

func (p HScroll) ScrollBoxPos() uint      { return uint(p.Msg.WParam.Hi16()) }
func (p HScroll) Request() co.SB_REQ      { return co.SB_REQ(p.Msg.WParam.Lo16()) }
func (p HScroll) HwndScrollbar() win.HWND { return win.HWND(p.Msg.LParam) }

type HScrollClipboard struct{ Msg Any }

func (p HScrollClipboard) ScrollBoxPos() uint      { return uint(p.Msg.WParam.Hi16()) }
func (p HScrollClipboard) Request() co.SB_REQ      { return co.SB_REQ(p.Msg.WParam.Lo16()) }
func (p HScrollClipboard) HwndScrollbar() win.HWND { return win.HWND(p.Msg.LParam) }

type InitDialog struct{ Msg Any }

func (p InitDialog) HwndFocused() win.HWND { return win.HWND(p.Msg.WParam) }

type InitMenuPopup struct{ Msg Any }

func (p InitMenuPopup) Hmenu() win.HMENU   { return win.HMENU(p.Msg.WParam) }
func (p InitMenuPopup) Pos() int           { return int(p.Msg.LParam.Lo16()) }
func (p InitMenuPopup) IsWindowMenu() bool { return p.Msg.LParam.Hi16() != 0 }

type Key struct{ Msg Any }

func (p Key) VirtualKeyCode() co.VK     { return co.VK(p.Msg.WParam) }
func (p Key) RepeatCount() uint         { return uint(p.Msg.LParam.Lo16()) }
func (p Key) ScanCode() uint            { return uint(p.Msg.LParam.Hi16Lo8()) }
func (p Key) IsExtendedKey() bool       { return (p.Msg.LParam.Hi16Hi8() & 0b0000_0001) != 0 }
func (p Key) HasAltKey() bool           { return (p.Msg.LParam.Hi16Hi8() & 0b0010_0000) != 0 }
func (p Key) IsKeyDownBeforeSend() bool { return (p.Msg.LParam.Hi16Hi8() & 0b0100_0000) != 0 }
func (p Key) IsReleasingKey() bool      { return (p.Msg.LParam.Hi16Hi8() & 0b1000_0000) != 0 }

type KillFocus struct{ Msg Any }

func (p KillFocus) HwndReceivingFocus() win.HWND { return win.HWND(p.Msg.LParam) }

type Menu struct{ Msg Any }

func (p Menu) ItemIndex() uint  { return uint(p.Msg.WParam) }
func (p Menu) Hmenu() win.HMENU { return win.HMENU(p.Msg.LParam) }

type MenuChar struct{ Msg Any }

func (p MenuChar) CharCode() rune          { return rune(p.Msg.WParam.Lo16()) }
func (p MenuChar) ActiveMenuType() co.MFMC { return co.MFMC(p.Msg.WParam.Hi16()) }
func (p MenuChar) ActiveMenu() win.HMENU   { return win.HMENU(p.Msg.LParam) }

type MenuGetObject struct{ Msg Any }

func (p MenuGetObject) Info() *win.MENUGETOBJECTINFO {
	return (*win.MENUGETOBJECTINFO)(unsafe.Pointer(p.Msg.LParam))
}

type MenuSelect struct{ Msg Any }

func (p MenuSelect) Item() int        { return int(p.Msg.WParam.Lo16()) }
func (p MenuSelect) Flags() co.MF     { return co.MF(p.Msg.WParam.Hi16()) }
func (p MenuSelect) Hmenu() win.HMENU { return win.HMENU(p.Msg.LParam) }

type Mouse struct{ Msg Any }

func (p Mouse) VirtualKeys() co.MK { return co.MK(p.Msg.WParam.Lo16()) }
func (p Mouse) HasCtrl() bool      { return (p.VirtualKeys() & co.MK_CONTROL) != 0 }
func (p Mouse) HasShift() bool     { return (p.VirtualKeys() & co.MK_SHIFT) != 0 }
func (p Mouse) IsLeftBtn() bool    { return (p.VirtualKeys() & co.MK_LBUTTON) != 0 }
func (p Mouse) IsMiddleBtn() bool  { return (p.VirtualKeys() & co.MK_MBUTTON) != 0 }
func (p Mouse) IsRightBtn() bool   { return (p.VirtualKeys() & co.MK_RBUTTON) != 0 }
func (p Mouse) IsXBtn1() bool      { return (p.VirtualKeys() & co.MK_XBUTTON1) != 0 }
func (p Mouse) IsXBtn2() bool      { return (p.VirtualKeys() & co.MK_XBUTTON2) != 0 }
func (p Mouse) Pos() win.POINT     { return p.Msg.LParam.MakePoint() }

type Move struct{ Msg Any }

func (p Move) ClientAreaPos() win.POINT { return p.Msg.LParam.MakePoint() }

type Moving struct{ Msg Any }

func (p Moving) WindowPos() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Msg.LParam)) }

type NcActivate struct{ Msg Any }

func (p NcActivate) IsActive() bool            { return p.Msg.WParam != 0 }
func (p NcActivate) IsVisualStyleActive() bool { return p.Msg.LParam == 0 }
func (p NcActivate) UpdatedRegion() win.HRGN   { return win.HRGN(p.Msg.LParam) }

type NcCalcSize struct{ Msg Any }

func (p NcCalcSize) ShouldIndicateValidPart() bool { return p.Msg.WParam != 0 }
func (p NcCalcSize) NcCalcSizeParams() *win.NCCALCSIZE_PARAMS {
	return (*win.NCCALCSIZE_PARAMS)(unsafe.Pointer(p.Msg.LParam))
}
func (p NcCalcSize) Rect() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Msg.LParam)) }

type NcHitTest struct{ Msg Any }

func (p NcHitTest) CursorPos() win.POINT { return p.Msg.LParam.MakePoint() }

type NcMouse struct{ Msg Any }

func (p NcMouse) HitTest() co.HT { return co.HT(p.Msg.WParam) }
func (p NcMouse) Pos() win.POINT { return p.Msg.LParam.MakePoint() }

type NcMouseX struct{ Msg Any }

func (p NcMouseX) HitTest() co.HT { return co.HT(p.Msg.WParam.Lo16()) }
func (p NcMouseX) IsXBtn1() bool  { return p.Msg.WParam.Hi16() == 0x0001 }
func (p NcMouseX) IsXBtn2() bool  { return p.Msg.WParam.Hi16() == 0x0002 }
func (p NcMouseX) Pos() win.POINT { return p.Msg.LParam.MakePoint() }

type NextMenu struct{ Msg Any }

func (p NextMenu) VirtualKeyCode() co.VK { return co.VK(p.Msg.WParam) }
func (p NextMenu) MdiNextMenu() *win.MDINEXTMENU {
	return (*win.MDINEXTMENU)(unsafe.Pointer(p.Msg.LParam))
}

type NcPaint struct{ Msg Any }

func (p NcPaint) UpdatedHrgn() win.HRGN { return win.HRGN(p.Msg.WParam) }

type PaintClipboard struct{ Msg Any }

func (p PaintClipboard) CbViewerWindow() win.HWND { return win.HWND(p.Msg.WParam) }
func (p PaintClipboard) PaintStruct() *win.PAINTSTRUCT {
	return (*win.PAINTSTRUCT)(unsafe.Pointer(p.Msg.LParam))
}

type PowerBroadcast struct{ Msg Any }

func (p PowerBroadcast) Event() co.PBT { return co.PBT(p.Msg.WParam) }
func (p PowerBroadcast) PowerBroadcastSetting() *win.POWERBROADCAST_SETTING {
	if p.Event() == co.PBT_POWERSETTINGCHANGE {
		return (*win.POWERBROADCAST_SETTING)(unsafe.Pointer(p.Msg.LParam))
	}
	return nil
}

type Print struct{ Msg Any }

func (p Print) Hdc() win.HDC           { return win.HDC(p.Msg.WParam) }
func (p Print) DrawingOptions() co.PRF { return co.PRF(p.Msg.LParam) }

type RenderFormat struct{ Msg Any }

func (p RenderFormat) ClipboardFormat() co.CF { return co.CF(p.Msg.WParam) }

type SetFocus struct{ Msg Any }

func (p SetFocus) HwndLosingFocus() win.HWND { return win.HWND(p.Msg.WParam) }

type SetFont struct{ Msg Any }

func (p SetFont) Hfont() win.HFONT   { return win.HFONT(p.Msg.WParam) }
func (p SetFont) ShouldRedraw() bool { return p.Msg.LParam == 1 }

type SetIcon struct{ Msg Any }

func (p SetIcon) Size() co.ICON_SZ { return co.ICON_SZ(p.Msg.WParam) }
func (p SetIcon) Hicon() win.HICON { return win.HICON(p.Msg.LParam) }

type Size struct{ Msg Any }

func (p Size) Request() co.SIZE_REQ     { return co.SIZE_REQ(p.Msg.WParam) }
func (p Size) ClientAreaSize() win.SIZE { return p.Msg.LParam.MakeSize() }

type SizeClipboard struct{ Msg Any }

func (p SizeClipboard) CbViewerWindow() win.HWND { return win.HWND(p.Msg.WParam) }
func (p SizeClipboard) NewDimensions() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.Msg.LParam)) }

type SysCommand struct{ Msg Any }

func (p SysCommand) RequestCommand() co.SC { return co.SC(p.Msg.WParam) }
func (p SysCommand) CursorPos() win.POINT  { return p.Msg.LParam.MakePoint() }

type UnInitMenuPopup struct{ Msg Any }

func (p UnInitMenuPopup) Hmenu() win.HMENU { return win.HMENU(p.Msg.WParam) }

type VScroll struct{ Msg Any }

func (p VScroll) ScrollBoxPos() uint      { return uint(p.Msg.WParam.Hi16()) }
func (p VScroll) Request() co.SB_REQ      { return co.SB_REQ(p.Msg.WParam.Lo16()) }
func (p VScroll) HwndScrollbar() win.HWND { return win.HWND(p.Msg.LParam) }

type VScrollClipboard struct{ Msg Any }

func (p VScrollClipboard) ScrollBoxPos() uint      { return uint(p.Msg.WParam.Hi16()) }
func (p VScrollClipboard) Request() co.SB_REQ      { return co.SB_REQ(p.Msg.WParam.Lo16()) }
func (p VScrollClipboard) HwndScrollbar() win.HWND { return win.HWND(p.Msg.LParam) }
