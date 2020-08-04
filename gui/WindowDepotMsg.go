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
type _WindowDepotMsg struct {
	mapMsgs map[co.WM]func(p Wm) uintptr
	mapCmds map[int32]func(p WmCommand)
}

func (me *_WindowDepotMsg) addMsg(msg co.WM, userFunc func(p Wm) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p Wm) uintptr, 16) // arbitrary capacity
	}
	me.mapMsgs[msg] = userFunc
}

func (me *_WindowDepotMsg) addCmd(cmd int32, userFunc func(p WmCommand)) {
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[int32]func(p WmCommand), 4) // arbitrary capacity
	}
	me.mapCmds[cmd] = userFunc
}

func (me *_WindowDepotMsg) processMessage(msg co.WM, p Wm) (uintptr, bool) {
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
func (me *_WindowDepotMsg) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(co.WM_ACTIVATE, func(p Wm) uintptr {
		userFunc(WmActivate(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmActivateApp(userFunc func(p WmActivateApp)) {
	me.addMsg(co.WM_ACTIVATEAPP, func(p Wm) uintptr {
		userFunc(WmActivateApp(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsg(co.WM_APPCOMMAND, func(p Wm) uintptr {
		userFunc(WmAppCommand(p))
		return 1
	})
}

func (me *_WindowDepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p Wm) uintptr {
		userFunc(WmChar{_WmChar(Wm(p))})
		return 0
	})
}

// Warning: default handled in WindowModal.
func (me *_WindowDepotMsg) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *_WindowDepotMsg) WmCommand(cmd int32, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

func (me *_WindowDepotMsg) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p Wm) uintptr {
		userFunc(WmContextMenu(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmCreate(userFunc func(p WmCreate) int32) {
	me.addMsg(co.WM_CREATE, func(p Wm) uintptr {
		return uintptr(userFunc(WmCreate(p)))
	})
}

func (me *_WindowDepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p Wm) uintptr {
		userFunc(WmDeadChar{_WmChar(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *_WindowDepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p Wm) uintptr {
		userFunc(WmDropFiles(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmGetDlgCode(userFunc func(p WmGetDlgCode) co.DLGC) {
	me.addMsg(co.WM_GETDLGCODE, func(p Wm) uintptr {
		return uintptr(userFunc(WmGetDlgCode(p)))
	})
}

func (me *_WindowDepotMsg) WmHelp(userFunc func(p WmHelp)) {
	me.addMsg(co.WM_HELP, func(p Wm) uintptr {
		userFunc(WmHelp(p))
		return 1
	})
}

func (me *_WindowDepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p Wm) uintptr {
		userFunc(WmHotKey(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p Wm) uintptr {
		userFunc(WmInitMenuPopup(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmKeyDown(userFunc func(p WmKeyDown)) {
	me.addMsg(co.WM_KEYDOWN, func(p Wm) uintptr {
		userFunc(WmKeyDown{_WmKey(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmKeyUp(userFunc func(p WmKeyUp)) {
	me.addMsg(co.WM_KEYUP, func(p Wm) uintptr {
		userFunc(WmKeyUp{_WmKey(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p Wm) uintptr {
		userFunc(WmKillFocus(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmLButtonDblClk{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmLButtonDown{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(co.WM_LBUTTONUP, func(p Wm) uintptr {
		userFunc(WmLButtonUp{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMButtonDblClk{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMButtonDown{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(co.WM_MBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMButtonUp{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.addMsg(co.WM_MENUCHAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenuChar(p)))
	})
}

func (me *_WindowDepotMsg) WmMenuCommand(userFunc func(p WmMenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p Wm) uintptr {
		userFunc(WmMenuCommand(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p Wm) uintptr {
		userFunc(WmMenuSelect(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p Wm) uintptr {
		userFunc(WmMouseHover{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(co.WM_MOUSELEAVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *_WindowDepotMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p Wm) uintptr {
		userFunc(WmMouseMove{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(co.WM_MOVE, func(p Wm) uintptr {
		userFunc(WmMove(p))
		return 0
	})
}

// Warning: default handled in WindowMain.
func (me *_WindowDepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// Warning: default handled in WindowControl.
func (me *_WindowDepotMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(co.WM_NCPAINT, func(p Wm) uintptr {
		userFunc(WmNcPaint(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmPaint(userFunc func()) {
	me.addMsg(co.WM_PAINT, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

func (me *_WindowDepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(co.WM_PRINT, func(p Wm) uintptr {
		userFunc(WmPrint(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmRButtonDblClk{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmRButtonDown{_WmButton(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(co.WM_RBUTTONUP, func(p Wm) uintptr {
		userFunc(WmRButtonUp{_WmButton(Wm(p))})
		return 0
	})
}

// Warning: default handled in WindowMain and WindowModal.
func (me *_WindowDepotMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(co.WM_SETFOCUS, func(p Wm) uintptr {
		userFunc(WmSetFocus(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(co.WM_SETFONT, func(p Wm) uintptr {
		userFunc(WmSetFont(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(co.WM_SIZE, func(p Wm) uintptr {
		userFunc(WmSize(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmSysChar(userFunc func(p WmSysChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p Wm) uintptr {
		userFunc(WmSysChar{_WmChar(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p Wm) uintptr {
		userFunc(WmSysCommand(p))
		return 0
	})
}

func (me *_WindowDepotMsg) WmSysDeadChar(userFunc func(p WmSysDeadChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p Wm) uintptr {
		userFunc(WmSysDeadChar{_WmChar(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmSysKeyDown(userFunc func(p WmSysKeyDown)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p Wm) uintptr {
		userFunc(WmSysKeyDown{_WmKey(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmSysKeyUp(userFunc func(p WmSysKeyUp)) {
	me.addMsg(co.WM_SYSKEYUP, func(p Wm) uintptr {
		userFunc(WmSysKeyUp{_WmKey(Wm(p))})
		return 0
	})
}

func (me *_WindowDepotMsg) WmTimeChange(userFunc func()) {
	me.addMsg(co.WM_TIMECHANGE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}
