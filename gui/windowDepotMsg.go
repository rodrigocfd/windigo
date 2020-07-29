/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
)

// Keeps all user message handlers.
type windowDepotMsg struct {
	mapMsgs map[co.WM]func(p Wm) uintptr
	mapCmds map[int32]func(p WmCommand)
}

func (me *windowDepotMsg) addMsg(msg co.WM, userFunc func(p Wm) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p Wm) uintptr, 16) // arbitrary capacity
	}
	me.mapMsgs[msg] = userFunc
}

func (me *windowDepotMsg) addCmd(cmd int32, userFunc func(p WmCommand)) {
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[int32]func(p WmCommand), 16) // arbitrary capacity
	}
	me.mapCmds[cmd] = userFunc
}

func (me *windowDepotMsg) processMessage(msg co.WM, p Wm) (uintptr, bool) {
	if msg == co.WM_COMMAND {
		pCmd := WmCommand(p)
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

// Warning: default handled in WindowMain.
func (me *windowDepotMsg) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(co.WM_ACTIVATE, func(p Wm) uintptr {
		userFunc(WmActivate(p))
		return 0
	})
}

func (me *windowDepotMsg) WmActivateApp(userFunc func(p WmActivateApp)) {
	me.addMsg(co.WM_ACTIVATEAPP, func(p Wm) uintptr {
		userFunc(WmActivateApp(p))
		return 0
	})
}

func (me *windowDepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsg(co.WM_APPCOMMAND, func(p Wm) uintptr {
		userFunc(WmAppCommand(p))
		return 1
	})
}

func (me *windowDepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p Wm) uintptr {
		userFunc(WmChar{bChar(Wm(p))})
		return 0
	})
}

// Warning: default handled in WindowModal.
func (me *windowDepotMsg) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmCommand(cmd int32, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

func (me *windowDepotMsg) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p Wm) uintptr {
		userFunc(WmContextMenu(p))
		return 0
	})
}

func (me *windowDepotMsg) WmCreate(userFunc func(p WmCreate) int32) {
	me.addMsg(co.WM_CREATE, func(p Wm) uintptr {
		return uintptr(userFunc(WmCreate(p)))
	})
}

func (me *windowDepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p Wm) uintptr {
		userFunc(WmDeadChar{bChar(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p Wm) uintptr {
		userFunc(WmDropFiles(p))
		return 0
	})
}

func (me *windowDepotMsg) WmGetDlgCode(userFunc func(p WmGetDlgCode) co.DLGC) {
	me.addMsg(co.WM_GETDLGCODE, func(p Wm) uintptr {
		return uintptr(userFunc(WmGetDlgCode(p)))
	})
}

func (me *windowDepotMsg) WmHelp(userFunc func(p WmHelp)) {
	me.addMsg(co.WM_HELP, func(p Wm) uintptr {
		userFunc(WmHelp(p))
		return 1
	})
}

func (me *windowDepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p Wm) uintptr {
		userFunc(WmHotKey(p))
		return 0
	})
}

func (me *windowDepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p Wm) uintptr {
		userFunc(WmInitMenuPopup(p))
		return 0
	})
}

func (me *windowDepotMsg) WmKeyDown(userFunc func(p WmKeyDown)) {
	me.addMsg(co.WM_KEYDOWN, func(p Wm) uintptr {
		userFunc(WmKeyDown{bKeyUpDn(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmKeyUp(userFunc func(p WmKeyUp)) {
	me.addMsg(co.WM_KEYUP, func(p Wm) uintptr {
		userFunc(WmKeyUp{bKeyUpDn(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p Wm) uintptr {
		userFunc(WmKillFocus(p))
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmLButtonDblClk{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmLButtonDown{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(co.WM_LBUTTONUP, func(p Wm) uintptr {
		userFunc(WmLButtonUp{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMButtonDblClk{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMButtonDown{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(co.WM_MBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMButtonUp{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.addMsg(co.WM_MENUCHAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenuChar(p)))
	})
}

func (me *windowDepotMsg) WmMenuCommand(userFunc func(p WmMenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p Wm) uintptr {
		userFunc(WmMenuCommand(p))
		return 0
	})
}

func (me *windowDepotMsg) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p Wm) uintptr {
		userFunc(WmMenuSelect(p))
		return 0
	})
}

func (me *windowDepotMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p Wm) uintptr {
		userFunc(WmMouseHover{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(co.WM_MOUSELEAVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p Wm) uintptr {
		userFunc(WmMouseMove{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(co.WM_MOVE, func(p Wm) uintptr {
		userFunc(WmMove(p))
		return 0
	})
}

// Warning: default handled in WindowMain.
func (me *windowDepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// Warning: default handled in WindowControl.
func (me *windowDepotMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(co.WM_NCPAINT, func(p Wm) uintptr {
		userFunc(WmNcPaint(p))
		return 0
	})
}

func (me *windowDepotMsg) WmPaint(userFunc func()) {
	me.addMsg(co.WM_PAINT, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(co.WM_PRINT, func(p Wm) uintptr {
		userFunc(WmPrint(p))
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmRButtonDblClk{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmRButtonDown{bButtonClk(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(co.WM_RBUTTONUP, func(p Wm) uintptr {
		userFunc(WmRButtonUp{bButtonClk(Wm(p))})
		return 0
	})
}

// Warning: default handled in WindowMain and WindowModal.
func (me *windowDepotMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(co.WM_SETFOCUS, func(p Wm) uintptr {
		userFunc(WmSetFocus(p))
		return 0
	})
}

func (me *windowDepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(co.WM_SETFONT, func(p Wm) uintptr {
		userFunc(WmSetFont(p))
		return 0
	})
}

func (me *windowDepotMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(co.WM_SIZE, func(p Wm) uintptr {
		userFunc(WmSize(p))
		return 0
	})
}

func (me *windowDepotMsg) WmSysChar(userFunc func(p WmSysChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p Wm) uintptr {
		userFunc(WmSysChar{bChar(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p Wm) uintptr {
		userFunc(WmSysCommand(p))
		return 0
	})
}

func (me *windowDepotMsg) WmSysDeadChar(userFunc func(p WmSysDeadChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p Wm) uintptr {
		userFunc(WmSysDeadChar{bChar(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmSysKeyDown(userFunc func(p WmSysKeyDown)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p Wm) uintptr {
		userFunc(WmSysKeyDown{bKeyUpDn(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmSysKeyUp(userFunc func(p WmSysKeyUp)) {
	me.addMsg(co.WM_SYSKEYUP, func(p Wm) uintptr {
		userFunc(WmSysKeyUp{bKeyUpDn(Wm(p))})
		return 0
	})
}

func (me *windowDepotMsg) WmTimeChange(userFunc func()) {
	me.addMsg(co.WM_TIMECHANGE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}
