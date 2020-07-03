/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package wm

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Raw window message parameters.
type Base struct {
	WParam win.WPARAM
	LParam win.LPARAM
}

type Activate struct{ Base }

func (p Activate) Event() co.WA                           { return co.WA(p.WParam.LoWord()) }
func (p Activate) IsMinimized() bool                      { return p.WParam.HiWord() != 0 }
func (p Activate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.LParam) }

type ActivateApp struct{ Base }

func (p ActivateApp) IsBeingActivated() bool { return p.WParam != 0 }
func (p ActivateApp) ThreadId() uint32       { return uint32(p.LParam) }

type AppCommand struct{ Base }

func (p AppCommand) OwnerWindow() win.HWND     { return win.HWND(p.WParam) }
func (p AppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.LParam.HiWord() &^ 0xF000) }
func (p AppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.LParam.HiWord() & 0xF000) }
func (p AppCommand) Keys() co.MK               { return co.MK(p.LParam.LoWord()) }

type Char struct{ Base }

func (p Char) CharCode() uint16          { return uint16(p.WParam) }
func (p Char) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p Char) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p Char) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p Char) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p Char) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }
func (p Char) KeyBeingReleased() bool    { return (p.LParam.HiByteHiWord() & 0b1000_0000) != 0 }

type Command struct{ Base }

func (p Command) IsFromMenu() bool         { return p.WParam.HiWord() == 0 }
func (p Command) IsFromAccelerator() bool  { return p.WParam.HiWord() == 1 }
func (p Command) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p Command) MenuId() int32            { return p.ControlId() }
func (p Command) AcceleratorId() int32     { return p.ControlId() }
func (p Command) ControlId() int32         { return int32(p.WParam.LoWord()) }
func (p Command) ControlNotifCode() uint16 { return p.WParam.HiWord() }
func (p Command) ControlHwnd() win.HWND    { return win.HWND(p.LParam) }

type ContextMenu struct{ Base }

func (p ContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.WParam) }
func (p ContextMenu) CursorPos() win.POINT         { return p.LParam.MakePoint() }

type Create struct{ Base }

func (p Create) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.LParam))
}

type DeadChar struct{ Char }

type DropFiles struct{ Base }

func (p DropFiles) Hdrop() win.HDROP { return win.HDROP(p.WParam) }

type GetDlgCode struct{ Base }

func (p GetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.LParam) }
func (p GetDlgCode) Msg() *win.MSG         { return (*win.MSG)(unsafe.Pointer(p.LParam)) }

type Help struct{ Base }

func (p Help) HelpInfo() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.LParam)) }

type HotKey struct{ Base }

func (p HotKey) HotKey() co.IDHOT      { return co.IDHOT(p.WParam) }
func (p HotKey) OtherKeys() co.MOD     { return co.MOD(p.LParam.LoWord()) }
func (p HotKey) VirtualKeyCode() co.VK { return co.VK(p.LParam.HiWord()) }

type InitMenuPopup struct{ Base }

func (p InitMenuPopup) Hmenu() win.HMENU        { return win.HMENU(p.WParam) }
func (p InitMenuPopup) SourceItemIndex() uint16 { return p.LParam.LoWord() }
func (p InitMenuPopup) IsWindowMenu() bool      { return p.LParam.HiWord() != 0 }

type KeyDown struct{ Base }

func (p KeyDown) VirtualKeyCode() co.VK     { return co.VK(p.WParam) }
func (p KeyDown) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p KeyDown) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p KeyDown) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p KeyDown) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p KeyDown) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }

type KeyUp struct{ KeyDown }

type KillFocus struct{ Base }

func (p KillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.WParam) }

type LButtonDblClk struct{ Base }

func (p LButtonDblClk) HasCtrl() bool      { return (co.MK(p.WParam) & co.MK_CONTROL) != 0 }
func (p LButtonDblClk) HasLeftBtn() bool   { return (co.MK(p.WParam) & co.MK_LBUTTON) != 0 }
func (p LButtonDblClk) HasMiddleBtn() bool { return (co.MK(p.WParam) & co.MK_MBUTTON) != 0 }
func (p LButtonDblClk) HasRightBtn() bool  { return (co.MK(p.WParam) & co.MK_RBUTTON) != 0 }
func (p LButtonDblClk) HasShift() bool     { return (co.MK(p.WParam) & co.MK_SHIFT) != 0 }
func (p LButtonDblClk) HasXBtn1() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON1) != 0 }
func (p LButtonDblClk) HasXBtn2() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON2) != 0 }
func (p LButtonDblClk) Pos() win.POINT     { return p.LParam.MakePoint() }

type LButtonDown struct{ LButtonDblClk }

type LButtonUp struct{ LButtonDblClk }

type MButtonDblClk struct{ LButtonDblClk }

type MButtonDown struct{ LButtonDblClk }

type MButtonUp struct{ LButtonDblClk }

type MenuChar struct{ Base }

func (p MenuChar) CharCode() uint16      { return p.WParam.LoWord() }
func (p MenuChar) ActiveMenuType() co.MF { return co.MF(p.WParam.HiWord()) }
func (p MenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.LParam) }

type MenuCommand struct{ Base }

func (p MenuCommand) ItemIndex() uint16 { return uint16(p.WParam) }
func (p MenuCommand) Hmenu() win.HMENU  { return win.HMENU(p.LParam) }

type MenuSelect struct{ Base }

func (p MenuSelect) Item() uint16     { return p.WParam.LoWord() }
func (p MenuSelect) Flags() co.MF     { return co.MF(p.WParam.HiWord()) }
func (p MenuSelect) Hmenu() win.HMENU { return win.HMENU(p.LParam) }

type MouseHover struct{ LButtonDblClk }

type MouseMove struct{ LButtonDblClk }

type Move struct{ Base }

func (p Move) Pos() win.POINT { return p.LParam.MakePoint() }

type NcPaint struct{ Base }

func (p NcPaint) Hrgn() win.HRGN { return win.HRGN(p.WParam) }

type Notify struct{ Base }

func (p Notify) NmHdr() *win.NMHDR { return (*win.NMHDR)(unsafe.Pointer(p.LParam)) }

type Print struct{ Base }

func (p Print) Hdc() win.HDC           { return win.HDC(p.WParam) }
func (p Print) DrawingOptions() co.PRF { return co.PRF(p.LParam) }

type RButtonDblClk struct{ LButtonDblClk }

type RButtonDown struct{ LButtonDblClk }

type RButtonUp struct{ LButtonDblClk }

type SetFocus struct{ Base }

func (p SetFocus) UnfocusedWindow() win.HWND { return win.HWND(p.WParam) }

type SetFont struct{ Base }

func (p SetFont) Hfont() win.HFONT   { return win.HFONT(p.WParam) }
func (p SetFont) ShouldRedraw() bool { return p.LParam == 1 }

type Size struct{ Base }

func (p Size) Request() co.SIZE         { return co.SIZE(p.WParam) }
func (p Size) ClientAreaSize() win.SIZE { return p.LParam.MakeSize() }

type SysChar struct{ Char }

type SysCommand struct{ Base }

func (p SysCommand) RequestCommand() co.SC { return co.SC(p.WParam) }
func (p SysCommand) CursorPos() win.POINT  { return p.LParam.MakePoint() }

type SysDeadChar struct{ Char }

type SysKeyDown struct{ KeyDown }

type SysKeyUp struct{ KeyDown }
