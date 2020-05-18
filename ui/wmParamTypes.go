/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// Raw window message parameters.
type wmBase struct {
	WParam api.WPARAM
	LParam api.LPARAM
}

type WmCommand struct{ base wmBase }

func (p WmCommand) IsFromMenu() bool         { return hiWordWp(p.base.WParam) == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return hiWordWp(p.base.WParam) == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() c.ID             { return p.ControlId() }
func (p WmCommand) AcceleratorId() c.ID      { return p.ControlId() }
func (p WmCommand) ControlId() c.ID          { return c.ID(loWordWp(p.base.WParam)) }
func (p WmCommand) ControlNotifCode() uint16 { return hiWordWp(p.base.WParam) }
func (p WmCommand) ControlHwnd() api.HWND    { return api.HWND(p.base.LParam) }

type WmNotify struct{ base wmBase }

func (p WmNotify) NmHdr() *api.NMHDR { return (*api.NMHDR)(unsafe.Pointer(p.base.LParam)) }

//------------------------------------------------------------------------------

type WmActivate struct{ base wmBase }

func (p WmActivate) State() c.WA                            { return c.WA(loWordWp(p.base.WParam)) }
func (p WmActivate) IsMinimized() bool                      { return hiWordWp(p.base.WParam) != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() api.HWND { return api.HWND(p.base.LParam) }

type WmActivateApp struct{ base wmBase }

func (p WmActivateApp) IsBeingActivated() bool { return p.base.WParam != 0 }
func (p WmActivateApp) ThreadId() uint32       { return uint32(p.base.LParam) }

type WmAppCommand struct{ base wmBase }

func (p WmAppCommand) OwnerWindow() api.HWND { return api.HWND(p.base.WParam) }
func (p WmAppCommand) AppCommand() c.APPCOMMAND {
	return c.APPCOMMAND(hiWordLp(p.base.LParam) &^ 0xF000)
}
func (p WmAppCommand) UDevice() c.FAPPCOMMAND { return c.FAPPCOMMAND(hiWordLp(p.base.LParam) & 0xF000) }
func (p WmAppCommand) Keys() c.MK             { return c.MK(loWordLp(p.base.LParam)) }

//------------------------------------------------------------------------------

type wmCharBase struct{ base wmBase }

func (p wmCharBase) CharCode() uint16    { return uint16(p.base.WParam) }
func (p wmCharBase) RepeatCount() uint16 { return api.LoWord(uint32(p.base.LParam)) }
func (p wmCharBase) ScanCode() uint8     { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p wmCharBase) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p wmCharBase) HasAltKey() bool { return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00100000) != 0 }
func (p wmCharBase) IsKeyDownBeforeSend() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b01000000) != 0
}
func (p wmCharBase) KeyBeingReleased() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b10000000) != 0
}

type WmChar struct{ wmCharBase } // inherit
type WmDeadChar struct{ wmCharBase }
type WmSysDeadChar struct{ wmCharBase }

//------------------------------------------------------------------------------

type WmCreate struct{ base wmBase }

func (p WmCreate) CreateStruct() *api.CREATESTRUCT {
	return (*api.CREATESTRUCT)(unsafe.Pointer(p.base.LParam))
}

type WmDropFiles struct{ base wmBase }

func (p WmDropFiles) Hdrop() api.HDROP { return api.HDROP(p.base.WParam) }

type WmHelp struct{ base wmBase }

func (p WmHelp) HelpInfo() *api.HELPINFO { return (*api.HELPINFO)(unsafe.Pointer(p.base.LParam)) }

type WmHotKey struct{ base wmBase }

func (p WmHotKey) HotKey() c.IDHOT      { return c.IDHOT(p.base.WParam) }
func (p WmHotKey) OtherKeys() c.MOD     { return c.MOD(loWordLp(p.base.LParam)) }
func (p WmHotKey) VirtualKeyCode() c.VK { return c.VK(hiWordLp(p.base.LParam)) }

type WmInitMenuPopup struct{ base wmBase }

func (p WmInitMenuPopup) Hmenu() api.HMENU        { return api.HMENU(p.base.WParam) }
func (p WmInitMenuPopup) SourceItemIndex() uint16 { return loWordLp(p.base.LParam) }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return hiWordLp(p.base.LParam) != 0 }

type WmKeyDown struct{ base wmBase }

func (p WmKeyDown) VirtualKeyCode() c.VK { return c.VK(p.base.WParam) }
func (p WmKeyDown) RepeatCount() uint16  { return api.LoWord(uint32(p.base.LParam)) }
func (p WmKeyDown) ScanCode() uint8      { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmKeyDown) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p WmKeyDown) IsKeyDownBeforeSend() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b01000000) != 0
}

type WmKeyUp struct{ base wmBase }

func (p WmKeyUp) VirtualKeyCode() c.VK { return c.VK(p.base.WParam) }
func (p WmKeyUp) ScanCode() uint8      { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmKeyUp) IsExtendedKey() bool  { return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0 }

type WmKillFocus struct{ base wmBase }

func (p WmKillFocus) WindowReceivingFocus() api.HWND { return api.HWND(p.base.WParam) }

//------------------------------------------------------------------------------

type wmLButtonDblClkBase struct{ base wmBase }

func (p wmLButtonDblClkBase) HasCtrl() bool      { return (c.MK(p.base.WParam) & c.MK_CONTROL) != 0 }
func (p wmLButtonDblClkBase) HasLeftBtn() bool   { return (c.MK(p.base.WParam) & c.MK_LBUTTON) != 0 }
func (p wmLButtonDblClkBase) HasMiddleBtn() bool { return (c.MK(p.base.WParam) & c.MK_MBUTTON) != 0 }
func (p wmLButtonDblClkBase) HasRightBtn() bool  { return (c.MK(p.base.WParam) & c.MK_RBUTTON) != 0 }
func (p wmLButtonDblClkBase) HasShift() bool     { return (c.MK(p.base.WParam) & c.MK_SHIFT) != 0 }
func (p wmLButtonDblClkBase) HasXBtn1() bool     { return (c.MK(p.base.WParam) & c.MK_XBUTTON1) != 0 }
func (p wmLButtonDblClkBase) HasXBtn2() bool     { return (c.MK(p.base.WParam) & c.MK_XBUTTON2) != 0 }
func (p wmLButtonDblClkBase) Pos() api.POINT     { return makePointLp(p.base.LParam) }

type WmLButtonDblClk struct{ wmLButtonDblClkBase } // inherit
type WmLButtonDown struct{ wmLButtonDblClkBase }
type WmLButtonUp struct{ wmLButtonDblClkBase }
type WmMButtonDblClk struct{ wmLButtonDblClkBase }
type WmMButtonDown struct{ wmLButtonDblClkBase }
type WmMButtonUp struct{ wmLButtonDblClkBase }
type WmMouseHover struct{ wmLButtonDblClkBase }
type WmMouseMove struct{ wmLButtonDblClkBase }
type WmRButtonDblClk struct{ wmLButtonDblClkBase }
type WmRButtonDown struct{ wmLButtonDblClkBase }
type WmRButtonUp struct{ wmLButtonDblClkBase }

//------------------------------------------------------------------------------

type WmMenuChar struct{ base wmBase }

func (p WmMenuChar) CharCode() uint16      { return loWordWp(p.base.WParam) }
func (p WmMenuChar) ActiveMenuType() c.MF  { return c.MF(hiWordWp(p.base.WParam)) }
func (p WmMenuChar) ActiveMenu() api.HMENU { return api.HMENU(p.base.LParam) }

type WmMove struct{ base wmBase }

func (p WmMove) Pos() api.POINT { return makePointLp(p.base.LParam) }

type WmNcPaint struct{ base wmBase }

func (p WmNcPaint) Hrgn() api.HRGN { return api.HRGN(p.base.WParam) }

type WmPrint struct{ base wmBase }

func (p WmPrint) Hdc() api.HDC          { return api.HDC(p.base.WParam) }
func (p WmPrint) DrawingOptions() c.PRF { return c.PRF(p.base.LParam) }

type WmSetFocus struct{ base wmBase }

func (p WmSetFocus) UnfocusedWindow() api.HWND { return api.HWND(p.base.WParam) }

type WmSetFont struct{ base wmBase }

func (p WmSetFont) Hfont() api.HFONT   { return api.HFONT(p.base.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.base.LParam == 1 }

type WmSize struct{ base wmBase }

func (p WmSize) Request() c.SIZE_REQ      { return c.SIZE_REQ(p.base.WParam) }
func (p WmSize) ClientAreaSize() api.SIZE { return makeSizeLp(p.base.LParam) }

type WmSysCommand struct{ base wmBase }

func (p WmSysCommand) RequestCommand() c.SC { return c.SC(p.base.WParam) }
func (p WmSysCommand) CursorPos() api.POINT { return makePointLp(p.base.LParam) }

type WmSysKeyDown struct{ base wmBase }

func (p WmSysKeyDown) VirtualKeyCode() c.VK { return c.VK(p.base.WParam) }
func (p WmSysKeyDown) RepeatCount() uint16  { return api.LoWord(uint32(p.base.LParam)) }
func (p WmSysKeyDown) ScanCode() uint8      { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmSysKeyDown) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p WmSysKeyDown) HasAltKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00100000) != 0
}
func (p WmSysKeyDown) IsKeyDownBeforeSend() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b01000000) != 0
}

type WmSysKeyUp struct{ base wmBase }

func (p WmSysKeyUp) VirtualKeyCode() c.VK { return c.VK(p.base.WParam) }
func (p WmSysKeyUp) ScanCode() uint8      { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmSysKeyUp) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p WmSysKeyUp) HasAltKey() bool { return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00100000) != 0 }

//------------------------------------------------------------------------------

func loWordWp(p api.WPARAM) uint16 { return api.LoWord(uint32(p)) }
func hiWordWp(p api.WPARAM) uint16 { return api.HiWord(uint32(p)) }
func loWordLp(p api.LPARAM) uint16 { return api.LoWord(uint32(p)) }
func hiWordLp(p api.LPARAM) uint16 { return api.HiWord(uint32(p)) }

func makePointLp(p api.LPARAM) api.POINT {
	return api.POINT{
		X: int32(api.LoWord(uint32(p))),
		Y: int32(api.HiWord(uint32(p))),
	}
}

func makeSizeLp(p api.LPARAM) api.SIZE {
	return api.SIZE{
		Cx: int32(api.LoWord(uint32(p))),
		Cy: int32(api.HiWord(uint32(p))),
	}
}
