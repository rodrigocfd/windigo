/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
	"wingows/co"
)

// Keeps all user message handlers.
type windowDepotMsg struct {
	mapMsgs map[co.WM]func(p wmBase) uintptr
	mapCmds map[co.ID]func(p WmCommand)
}

func (me *windowDepotMsg) addMsg(msg co.WM, userFunc func(p wmBase) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p wmBase) uintptr, 16) // arbitrary capacity
	}
	me.mapMsgs[msg] = userFunc
}

func (me *windowDepotMsg) addCmd(cmd co.ID, userFunc func(p WmCommand)) {
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[co.ID]func(p WmCommand), 16) // arbitrary capacity
	}
	me.mapCmds[cmd] = userFunc
}

func (me *windowDepotMsg) processMessage(msg co.WM, p wmBase) (uintptr, bool) {
	if msg == co.WM_COMMAND {
		pCmd := WmCommand{base: p}
		if userFunc, hasCmd := me.mapCmds[pCmd.ControlId()]; hasCmd {
			userFunc(pCmd)
			return 0, true // always return zero
		}
	} else {
		if userFunc, hasMsg := me.mapMsgs[msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}

//------------------------------------------------------------------------------

// Raw window message parameters.
type wmBase struct {
	WParam api.WPARAM
	LParam api.LPARAM
}

type WmCommand struct{ base wmBase }

func (p WmCommand) IsFromMenu() bool         { return hiWordWp(p.base.WParam) == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return hiWordWp(p.base.WParam) == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() co.ID            { return p.ControlId() }
func (p WmCommand) AcceleratorId() co.ID     { return p.ControlId() }
func (p WmCommand) ControlId() co.ID         { return co.ID(loWordWp(p.base.WParam)) }
func (p WmCommand) ControlNotifCode() uint16 { return hiWordWp(p.base.WParam) }
func (p WmCommand) ControlHwnd() api.HWND    { return api.HWND(p.base.LParam) }

func (me *windowDepotMsg) WmCommand(cmd co.ID, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

// Not directly handled, use the common control notifications instead.
type WmNotify struct{ base wmBase }

func (p WmNotify) NmHdr() *api.NMHDR { return (*api.NMHDR)(unsafe.Pointer(p.base.LParam)) }

//------------------------------------------------------------------------------

type WmActivate struct{ base wmBase }

func (p WmActivate) State() co.WA                           { return co.WA(loWordWp(p.base.WParam)) }
func (p WmActivate) IsMinimized() bool                      { return hiWordWp(p.base.WParam) != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() api.HWND { return api.HWND(p.base.LParam) }

func (me *windowDepotMsg) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(co.WM_ACTIVATE, func(p wmBase) uintptr {
		userFunc(WmActivate{base: wmBase(p)})
		return 0
	})
}

type WmActivateApp struct{ base wmBase }

func (p WmActivateApp) IsBeingActivated() bool { return p.base.WParam != 0 }
func (p WmActivateApp) ThreadId() uint32       { return uint32(p.base.LParam) }

func (me *windowDepotMsg) WmActivateApp(userFunc func(p WmActivateApp)) {
	me.addMsg(co.WM_ACTIVATEAPP, func(p wmBase) uintptr {
		userFunc(WmActivateApp{base: wmBase(p)})
		return 0
	})
}

type WmAppCommand struct{ base wmBase }

func (p WmAppCommand) OwnerWindow() api.HWND { return api.HWND(p.base.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND {
	return co.APPCOMMAND(hiWordLp(p.base.LParam) &^ 0xF000)
}
func (p WmAppCommand) UDevice() co.FAPPCOMMAND {
	return co.FAPPCOMMAND(hiWordLp(p.base.LParam) & 0xF000)
}
func (p WmAppCommand) Keys() co.MK { return co.MK(loWordLp(p.base.LParam)) }

func (me *windowDepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsg(co.WM_APPCOMMAND, func(p wmBase) uintptr {
		userFunc(WmAppCommand{base: wmBase(p)})
		return 1
	})
}

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

type WmChar struct{ wmCharBase }

func (me *windowDepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p wmBase) uintptr {
		userFunc(WmChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

type WmDeadChar struct{ wmCharBase }

func (me *windowDepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p wmBase) uintptr {
		userFunc(WmDeadChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

type WmSysChar struct{ wmCharBase }

func (me *windowDepotMsg) WmSysChar(userFunc func(p WmSysChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p wmBase) uintptr {
		userFunc(WmSysChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

type WmSysDeadChar struct{ wmCharBase }

func (me *windowDepotMsg) WmSysDeadChar(userFunc func(p WmSysDeadChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p wmBase) uintptr {
		userFunc(WmSysDeadChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowDepotMsg) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmContextMenu struct{ base wmBase }

func (p WmContextMenu) RightClickedWindow() api.HWND { return api.HWND(p.base.WParam) }
func (p WmContextMenu) CursorPos() api.POINT         { return makePointLp(p.base.LParam) }

func (me *windowDepotMsg) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p wmBase) uintptr {
		userFunc(WmContextMenu{base: wmBase(p)})
		return 0
	})
}

type WmCreate struct{ base wmBase }

func (p WmCreate) CreateStruct() *api.CREATESTRUCT {
	return (*api.CREATESTRUCT)(unsafe.Pointer(p.base.LParam))
}

func (me *windowDepotMsg) WmCreate(userFunc func(p WmCreate) int32) {
	me.addMsg(co.WM_CREATE, func(p wmBase) uintptr {
		return uintptr(userFunc(WmCreate{base: wmBase(p)}))
	})
}

func (me *windowDepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmDropFiles struct{ base wmBase }

func (p WmDropFiles) Hdrop() api.HDROP { return api.HDROP(p.base.WParam) }

func (me *windowDepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p wmBase) uintptr {
		userFunc(WmDropFiles{base: wmBase(p)})
		return 0
	})
}

type WmHelp struct{ base wmBase }

func (p WmHelp) HelpInfo() *api.HELPINFO { return (*api.HELPINFO)(unsafe.Pointer(p.base.LParam)) }

func (me *windowDepotMsg) WmHelp(userFunc func(p WmHelp)) {
	me.addMsg(co.WM_HELP, func(p wmBase) uintptr {
		userFunc(WmHelp{base: wmBase(p)})
		return 1
	})
}

type WmHotKey struct{ base wmBase }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.base.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(loWordLp(p.base.LParam)) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(hiWordLp(p.base.LParam)) }

func (me *windowDepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p wmBase) uintptr {
		userFunc(WmHotKey{base: wmBase(p)})
		return 0
	})
}

type WmInitMenuPopup struct{ base wmBase }

func (p WmInitMenuPopup) Hmenu() api.HMENU        { return api.HMENU(p.base.WParam) }
func (p WmInitMenuPopup) SourceItemIndex() uint16 { return loWordLp(p.base.LParam) }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return hiWordLp(p.base.LParam) != 0 }

func (me *windowDepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p wmBase) uintptr {
		userFunc(WmInitMenuPopup{base: wmBase(p)})
		return 0
	})
}

type WmKeyDown struct{ base wmBase }

func (p WmKeyDown) VirtualKeyCode() co.VK { return co.VK(p.base.WParam) }
func (p WmKeyDown) RepeatCount() uint16   { return api.LoWord(uint32(p.base.LParam)) }
func (p WmKeyDown) ScanCode() uint8       { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmKeyDown) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p WmKeyDown) IsKeyDownBeforeSend() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b01000000) != 0
}

func (me *windowDepotMsg) WmKeyDown(userFunc func(p WmKeyDown)) {
	me.addMsg(co.WM_KEYDOWN, func(p wmBase) uintptr {
		userFunc(WmKeyDown{base: wmBase(p)})
		return 0
	})
}

type WmKeyUp struct{ base wmBase }

func (p WmKeyUp) VirtualKeyCode() co.VK { return co.VK(p.base.WParam) }
func (p WmKeyUp) ScanCode() uint8       { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmKeyUp) IsExtendedKey() bool   { return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0 }

func (me *windowDepotMsg) WmKeyUp(userFunc func(p WmKeyUp)) {
	me.addMsg(co.WM_KEYUP, func(p wmBase) uintptr {
		userFunc(WmKeyUp{base: wmBase(p)})
		return 0
	})
}

type WmKillFocus struct{ base wmBase }

func (p WmKillFocus) WindowReceivingFocus() api.HWND { return api.HWND(p.base.WParam) }

func (me *windowDepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p wmBase) uintptr {
		userFunc(WmKillFocus{base: wmBase(p)})
		return 0
	})
}

//------------------------------------------------------------------------------

type wmLButtonDblClkBase struct{ base wmBase }

func (p wmLButtonDblClkBase) HasCtrl() bool    { return (co.MK(p.base.WParam) & co.MK_CONTROL) != 0 }
func (p wmLButtonDblClkBase) HasLeftBtn() bool { return (co.MK(p.base.WParam) & co.MK_LBUTTON) != 0 }
func (p wmLButtonDblClkBase) HasMiddleBtn() bool {
	return (co.MK(p.base.WParam) & co.MK_MBUTTON) != 0
}
func (p wmLButtonDblClkBase) HasRightBtn() bool { return (co.MK(p.base.WParam) & co.MK_RBUTTON) != 0 }
func (p wmLButtonDblClkBase) HasShift() bool    { return (co.MK(p.base.WParam) & co.MK_SHIFT) != 0 }
func (p wmLButtonDblClkBase) HasXBtn1() bool    { return (co.MK(p.base.WParam) & co.MK_XBUTTON1) != 0 }
func (p wmLButtonDblClkBase) HasXBtn2() bool    { return (co.MK(p.base.WParam) & co.MK_XBUTTON2) != 0 }
func (p wmLButtonDblClkBase) Pos() api.POINT    { return makePointLp(p.base.LParam) }

type WmLButtonDblClk struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmLButtonDblClk{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmLButtonDown struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmLButtonDown{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmLButtonUp struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(co.WM_LBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmLButtonUp{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmMButtonDblClk struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmMButtonDblClk{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmMButtonDown struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmMButtonDown{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmMButtonUp struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(co.WM_MBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmMButtonUp{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmMouseHover struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p wmBase) uintptr {
		userFunc(WmMouseHover{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmMouseMove struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p wmBase) uintptr {
		userFunc(WmMouseMove{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmRButtonDblClk struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmRButtonDblClk{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmRButtonDown struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmRButtonDown{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

type WmRButtonUp struct{ wmLButtonDblClkBase }

func (me *windowDepotMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(co.WM_RBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmRButtonUp{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

//------------------------------------------------------------------------------

type WmMenuChar struct{ base wmBase }

func (p WmMenuChar) CharCode() uint16      { return loWordWp(p.base.WParam) }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(hiWordWp(p.base.WParam)) }
func (p WmMenuChar) ActiveMenu() api.HMENU { return api.HMENU(p.base.LParam) }

func (me *windowDepotMsg) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.addMsg(co.WM_MENUCHAR, func(p wmBase) uintptr {
		return uintptr(userFunc(WmMenuChar{base: wmBase(p)}))
	})
}

type WmMenuCommand struct{ base wmBase }

func (p WmMenuCommand) ItemIndex() uint16 { return uint16(p.base.WParam) }
func (p WmMenuCommand) Hmenu() api.HMENU  { return api.HMENU(p.base.LParam) }

func (me *windowDepotMsg) WmMenuCommand(userFunc func(p WmMenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p wmBase) uintptr {
		userFunc(WmMenuCommand{base: wmBase(p)})
		return 0
	})
}

type WmMenuSelect struct{ base wmBase }

func (p WmMenuSelect) Item() uint16     { return loWordWp(p.base.WParam) }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(hiWordWp(p.base.WParam)) }
func (p WmMenuSelect) Hmenu() api.HMENU { return api.HMENU(p.base.LParam) }

func (me *windowDepotMsg) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p wmBase) uintptr {
		userFunc(WmMenuSelect{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(co.WM_MOUSELEAVE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmMove struct{ base wmBase }

func (p WmMove) Pos() api.POINT { return makePointLp(p.base.LParam) }

func (me *windowDepotMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(co.WM_MOVE, func(p wmBase) uintptr {
		userFunc(WmMove{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmNcPaint struct{ base wmBase }

func (p WmNcPaint) Hrgn() api.HRGN { return api.HRGN(p.base.WParam) }

func (me *windowDepotMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(co.WM_NCPAINT, func(p wmBase) uintptr {
		userFunc(WmNcPaint{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmPaint(userFunc func()) {
	me.addMsg(co.WM_PAINT, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmPrint struct{ base wmBase }

func (p WmPrint) Hdc() api.HDC           { return api.HDC(p.base.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.base.LParam) }

func (me *windowDepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(co.WM_PRINT, func(p wmBase) uintptr {
		userFunc(WmPrint{base: wmBase(p)})
		return 0
	})
}

type WmSetFocus struct{ base wmBase }

func (p WmSetFocus) UnfocusedWindow() api.HWND { return api.HWND(p.base.WParam) }

func (me *windowDepotMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(co.WM_SETFOCUS, func(p wmBase) uintptr {
		userFunc(WmSetFocus{base: wmBase(p)})
		return 0
	})
}

type WmSetFont struct{ base wmBase }

func (p WmSetFont) Hfont() api.HFONT   { return api.HFONT(p.base.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.base.LParam == 1 }

func (me *windowDepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(co.WM_SETFONT, func(p wmBase) uintptr {
		userFunc(WmSetFont{base: wmBase(p)})
		return 0
	})
}

type WmSize struct{ base wmBase }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.base.WParam) }
func (p WmSize) ClientAreaSize() api.SIZE { return makeSizeLp(p.base.LParam) }

func (me *windowDepotMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(co.WM_SIZE, func(p wmBase) uintptr {
		userFunc(WmSize{base: wmBase(p)})
		return 0
	})
}

type WmSysCommand struct{ base wmBase }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.base.WParam) }
func (p WmSysCommand) CursorPos() api.POINT  { return makePointLp(p.base.LParam) }

func (me *windowDepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p wmBase) uintptr {
		userFunc(WmSysCommand{base: wmBase(p)})
		return 0
	})
}

type WmSysKeyDown struct{ base wmBase }

func (p WmSysKeyDown) VirtualKeyCode() co.VK { return co.VK(p.base.WParam) }
func (p WmSysKeyDown) RepeatCount() uint16   { return api.LoWord(uint32(p.base.LParam)) }
func (p WmSysKeyDown) ScanCode() uint8       { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmSysKeyDown) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p WmSysKeyDown) HasAltKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00100000) != 0
}
func (p WmSysKeyDown) IsKeyDownBeforeSend() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b01000000) != 0
}

func (me *windowDepotMsg) WmSysKeyDown(userFunc func(p WmSysKeyDown)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p wmBase) uintptr {
		userFunc(WmSysKeyDown{base: wmBase(p)})
		return 0
	})
}

type WmSysKeyUp struct{ base wmBase }

func (p WmSysKeyUp) VirtualKeyCode() co.VK { return co.VK(p.base.WParam) }
func (p WmSysKeyUp) ScanCode() uint8       { return api.LoByte(api.HiWord(uint32(p.base.LParam))) }
func (p WmSysKeyUp) IsExtendedKey() bool {
	return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00000001) != 0
}
func (p WmSysKeyUp) HasAltKey() bool { return (api.HiByte(hiWordLp(p.base.LParam)) & 0b00100000) != 0 }

func (me *windowDepotMsg) WmSysKeyUp(userFunc func(p WmSysKeyUp)) {
	me.addMsg(co.WM_SYSKEYUP, func(p wmBase) uintptr {
		userFunc(WmSysKeyUp{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmTimeChange(userFunc func()) {
	me.addMsg(co.WM_TIMECHANGE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

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
