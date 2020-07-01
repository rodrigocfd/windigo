/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
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
	WParam win.WPARAM
	LParam win.LPARAM
}

type WmCommand struct{ base wmBase }

func (p WmCommand) IsFromMenu() bool         { return p.base.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return p.base.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() co.ID            { return p.ControlId() }
func (p WmCommand) AcceleratorId() co.ID     { return p.ControlId() }
func (p WmCommand) ControlId() co.ID         { return co.ID(p.base.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() uint16 { return p.base.WParam.HiWord() }
func (p WmCommand) ControlHwnd() win.HWND    { return win.HWND(p.base.LParam) }

func (me *windowDepotMsg) WmCommand(cmd co.ID, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

// Not directly handled, use the common control notifications instead.
type WmNotify struct{ base wmBase }

func (p WmNotify) NmHdr() *win.NMHDR { return (*win.NMHDR)(unsafe.Pointer(p.base.LParam)) }

//------------------------------------------------------------------------------

type WmActivate struct{ base wmBase }

func (p WmActivate) Event() co.WA                           { return co.WA(p.base.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                      { return p.base.WParam.HiWord() != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.base.LParam) }

// Warning: default handled in WindowMain.
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

func (p WmAppCommand) OwnerWindow() win.HWND { return win.HWND(p.base.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND {
	return co.APPCOMMAND(p.base.LParam.HiWord() &^ 0xF000)
}
func (p WmAppCommand) UDevice() co.FAPPCOMMAND {
	return co.FAPPCOMMAND(p.base.LParam.HiWord() & 0xF000)
}
func (p WmAppCommand) Keys() co.MK { return co.MK(p.base.LParam.LoWord()) }

func (me *windowDepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsg(co.WM_APPCOMMAND, func(p wmBase) uintptr {
		userFunc(WmAppCommand{base: wmBase(p)})
		return 1
	})
}

// Warning: default handled in WindowModal.
func (me *windowDepotMsg) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmContextMenu struct{ base wmBase }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.base.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.base.LParam.MakePoint() }

func (me *windowDepotMsg) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p wmBase) uintptr {
		userFunc(WmContextMenu{base: wmBase(p)})
		return 0
	})
}

type WmCreate struct{ base wmBase }

func (p WmCreate) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.base.LParam))
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

func (p WmDropFiles) Hdrop() win.HDROP { return win.HDROP(p.base.WParam) }

func (me *windowDepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p wmBase) uintptr {
		userFunc(WmDropFiles{base: wmBase(p)})
		return 0
	})
}

type WmHelp struct{ base wmBase }

func (p WmHelp) HelpInfo() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.base.LParam)) }

func (me *windowDepotMsg) WmHelp(userFunc func(p WmHelp)) {
	me.addMsg(co.WM_HELP, func(p wmBase) uintptr {
		userFunc(WmHelp{base: wmBase(p)})
		return 1
	})
}

type WmHotKey struct{ base wmBase }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.base.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.base.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.base.LParam.HiWord()) }

func (me *windowDepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p wmBase) uintptr {
		userFunc(WmHotKey{base: wmBase(p)})
		return 0
	})
}

type WmInitMenuPopup struct{ base wmBase }

func (p WmInitMenuPopup) Hmenu() win.HMENU        { return win.HMENU(p.base.WParam) }
func (p WmInitMenuPopup) SourceItemIndex() uint16 { return p.base.LParam.LoWord() }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return p.base.LParam.HiWord() != 0 }

func (me *windowDepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p wmBase) uintptr {
		userFunc(WmInitMenuPopup{base: wmBase(p)})
		return 0
	})
}

type WmKillFocus struct{ base wmBase }

func (p WmKillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.base.WParam) }

func (me *windowDepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p wmBase) uintptr {
		userFunc(WmKillFocus{base: wmBase(p)})
		return 0
	})
}

type WmMenuChar struct{ base wmBase }

func (p WmMenuChar) CharCode() uint16      { return p.base.WParam.LoWord() }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.base.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.base.LParam) }

func (me *windowDepotMsg) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.addMsg(co.WM_MENUCHAR, func(p wmBase) uintptr {
		return uintptr(userFunc(WmMenuChar{base: wmBase(p)}))
	})
}

type WmMenuCommand struct{ base wmBase }

func (p WmMenuCommand) ItemIndex() uint16 { return uint16(p.base.WParam) }
func (p WmMenuCommand) Hmenu() win.HMENU  { return win.HMENU(p.base.LParam) }

func (me *windowDepotMsg) WmMenuCommand(userFunc func(p WmMenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p wmBase) uintptr {
		userFunc(WmMenuCommand{base: wmBase(p)})
		return 0
	})
}

type WmMenuSelect struct{ base wmBase }

func (p WmMenuSelect) Item() uint16     { return p.base.WParam.LoWord() }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.base.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.base.LParam) }

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

func (p WmMove) Pos() win.POINT { return p.base.LParam.MakePoint() }

func (me *windowDepotMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(co.WM_MOVE, func(p wmBase) uintptr {
		userFunc(WmMove{base: wmBase(p)})
		return 0
	})
}

// Warning: default handled in WindowMain.
func (me *windowDepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

type WmNcPaint struct{ base wmBase }

func (p WmNcPaint) Hrgn() win.HRGN { return win.HRGN(p.base.WParam) }

// Warning: default handled in WindowControl.
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

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.base.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.base.LParam) }

func (me *windowDepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(co.WM_PRINT, func(p wmBase) uintptr {
		userFunc(WmPrint{base: wmBase(p)})
		return 0
	})
}

type WmSetFocus struct{ base wmBase }

func (p WmSetFocus) UnfocusedWindow() win.HWND { return win.HWND(p.base.WParam) }

// Warning: default handled in WindowMain and WindowModal.
func (me *windowDepotMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(co.WM_SETFOCUS, func(p wmBase) uintptr {
		userFunc(WmSetFocus{base: wmBase(p)})
		return 0
	})
}

type WmSetFont struct{ base wmBase }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.base.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.base.LParam == 1 }

func (me *windowDepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(co.WM_SETFONT, func(p wmBase) uintptr {
		userFunc(WmSetFont{base: wmBase(p)})
		return 0
	})
}

type WmSize struct{ base wmBase }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.base.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.base.LParam.MakeSize() }

func (me *windowDepotMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(co.WM_SIZE, func(p wmBase) uintptr {
		userFunc(WmSize{base: wmBase(p)})
		return 0
	})
}

type WmSysCommand struct{ base wmBase }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.base.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.base.LParam.MakePoint() }

func (me *windowDepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p wmBase) uintptr {
		userFunc(WmSysCommand{base: wmBase(p)})
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

type wmBaseButton struct{ base wmBase }

func (p wmBaseButton) HasCtrl() bool      { return (co.MK(p.base.WParam) & co.MK_CONTROL) != 0 }
func (p wmBaseButton) HasLeftBtn() bool   { return (co.MK(p.base.WParam) & co.MK_LBUTTON) != 0 }
func (p wmBaseButton) HasMiddleBtn() bool { return (co.MK(p.base.WParam) & co.MK_MBUTTON) != 0 }
func (p wmBaseButton) HasRightBtn() bool  { return (co.MK(p.base.WParam) & co.MK_RBUTTON) != 0 }
func (p wmBaseButton) HasShift() bool     { return (co.MK(p.base.WParam) & co.MK_SHIFT) != 0 }
func (p wmBaseButton) HasXBtn1() bool     { return (co.MK(p.base.WParam) & co.MK_XBUTTON1) != 0 }
func (p wmBaseButton) HasXBtn2() bool     { return (co.MK(p.base.WParam) & co.MK_XBUTTON2) != 0 }
func (p wmBaseButton) Pos() win.POINT     { return p.base.LParam.MakePoint() }

type WmLButtonDblClk struct{ wmBaseButton }

func (me *windowDepotMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmLButtonDblClk{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmLButtonDown struct{ wmBaseButton }

func (me *windowDepotMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmLButtonDown{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmLButtonUp struct{ wmBaseButton }

func (me *windowDepotMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(co.WM_LBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmLButtonUp{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmMButtonDblClk struct{ wmBaseButton }

func (me *windowDepotMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmMButtonDblClk{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmMButtonDown struct{ wmBaseButton }

func (me *windowDepotMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmMButtonDown{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmMButtonUp struct{ wmBaseButton }

func (me *windowDepotMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(co.WM_MBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmMButtonUp{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmMouseHover struct{ wmBaseButton }

func (me *windowDepotMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p wmBase) uintptr {
		userFunc(WmMouseHover{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmMouseMove struct{ wmBaseButton }

func (me *windowDepotMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p wmBase) uintptr {
		userFunc(WmMouseMove{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmRButtonDblClk struct{ wmBaseButton }

func (me *windowDepotMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmRButtonDblClk{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmRButtonDown struct{ wmBaseButton }

func (me *windowDepotMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmRButtonDown{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

type WmRButtonUp struct{ wmBaseButton }

func (me *windowDepotMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(co.WM_RBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmRButtonUp{wmBaseButton: wmBaseButton{base: wmBase(p)}})
		return 0
	})
}

//------------------------------------------------------------------------------

type wmBaseChar struct{ base wmBase }

func (p wmBaseChar) CharCode() uint16    { return uint16(p.base.WParam) }
func (p wmBaseChar) RepeatCount() uint16 { return p.base.LParam.LoWord() }
func (p wmBaseChar) ScanCode() uint8     { return p.base.LParam.LoByteHiWord() }
func (p wmBaseChar) IsExtendedKey() bool { return (p.base.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p wmBaseChar) HasAltKey() bool     { return (p.base.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p wmBaseChar) IsKeyDownBeforeSend() bool {
	return (p.base.LParam.HiByteHiWord() & 0b0100_0000) != 0
}
func (p wmBaseChar) KeyBeingReleased() bool { return (p.base.LParam.HiByteHiWord() & 0b1000_0000) != 0 }

type WmChar struct{ wmBaseChar }

func (me *windowDepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p wmBase) uintptr {
		userFunc(WmChar{wmBaseChar: wmBaseChar{base: wmBase(p)}})
		return 0
	})
}

type WmDeadChar struct{ wmBaseChar }

func (me *windowDepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p wmBase) uintptr {
		userFunc(WmDeadChar{wmBaseChar: wmBaseChar{base: wmBase(p)}})
		return 0
	})
}

type WmSysChar struct{ wmBaseChar }

func (me *windowDepotMsg) WmSysChar(userFunc func(p WmSysChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p wmBase) uintptr {
		userFunc(WmSysChar{wmBaseChar: wmBaseChar{base: wmBase(p)}})
		return 0
	})
}

type WmSysDeadChar struct{ wmBaseChar }

func (me *windowDepotMsg) WmSysDeadChar(userFunc func(p WmSysDeadChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p wmBase) uintptr {
		userFunc(WmSysDeadChar{wmBaseChar: wmBaseChar{base: wmBase(p)}})
		return 0
	})
}

//------------------------------------------------------------------------------

type wmBaseVirtKey struct{ base wmBase }

func (p wmBaseVirtKey) VirtualKeyCode() co.VK { return co.VK(p.base.WParam) }
func (p wmBaseVirtKey) RepeatCount() uint16   { return p.base.LParam.LoWord() }
func (p wmBaseVirtKey) ScanCode() uint8       { return p.base.LParam.LoByteHiWord() }
func (p wmBaseVirtKey) IsExtendedKey() bool   { return (p.base.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p wmBaseVirtKey) HasAltKey() bool       { return (p.base.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p wmBaseVirtKey) IsKeyDownBeforeSend() bool {
	return (p.base.LParam.HiByteHiWord() & 0b0100_0000) != 0
}

type WmKeyDown struct{ wmBaseVirtKey }

func (me *windowDepotMsg) WmKeyDown(userFunc func(p WmKeyDown)) {
	me.addMsg(co.WM_KEYDOWN, func(p wmBase) uintptr {
		userFunc(WmKeyDown{wmBaseVirtKey: wmBaseVirtKey{base: wmBase(p)}})
		return 0
	})
}

type WmKeyUp struct{ wmBaseVirtKey }

func (me *windowDepotMsg) WmKeyUp(userFunc func(p WmKeyUp)) {
	me.addMsg(co.WM_KEYUP, func(p wmBase) uintptr {
		userFunc(WmKeyUp{wmBaseVirtKey: wmBaseVirtKey{base: wmBase(p)}})
		return 0
	})
}

type WmSysKeyDown struct{ wmBaseVirtKey }

func (me *windowDepotMsg) WmSysKeyDown(userFunc func(p WmSysKeyDown)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p wmBase) uintptr {
		userFunc(WmSysKeyDown{wmBaseVirtKey: wmBaseVirtKey{base: wmBase(p)}})
		return 0
	})
}

type WmSysKeyUp struct{ wmBaseVirtKey }

func (me *windowDepotMsg) WmSysKeyUp(userFunc func(p WmSysKeyUp)) {
	me.addMsg(co.WM_SYSKEYUP, func(p wmBase) uintptr {
		userFunc(WmSysKeyUp{wmBaseVirtKey: wmBaseVirtKey{base: wmBase(p)}})
		return 0
	})
}
