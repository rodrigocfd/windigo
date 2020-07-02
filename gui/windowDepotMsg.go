/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
	"wingows/gui/wm"
)

// Keeps all user message handlers.
type windowDepotMsg struct {
	mapMsgs map[co.WM]func(p wm.Base) uintptr
	mapCmds map[int32]func(p wm.Command)
}

func (me *windowDepotMsg) addMsg(msg co.WM, userFunc func(p wm.Base) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p wm.Base) uintptr, 16) // arbitrary capacity
	}
	me.mapMsgs[msg] = userFunc
}

func (me *windowDepotMsg) addCmd(cmd int32, userFunc func(p wm.Command)) {
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[int32]func(p wm.Command), 16) // arbitrary capacity
	}
	me.mapCmds[cmd] = userFunc
}

func (me *windowDepotMsg) processMessage(msg co.WM, p wm.Base) (uintptr, bool) {
	if msg == co.WM_COMMAND {
		pCmd := wm.Command{Base: p}
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
func (me *windowDepotMsg) WmActivate(userFunc func(p wm.Activate)) {
	me.addMsg(co.WM_ACTIVATE, func(p wm.Base) uintptr {
		userFunc(wm.Activate{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmActivateApp(userFunc func(p wm.ActivateApp)) {
	me.addMsg(co.WM_ACTIVATEAPP, func(p wm.Base) uintptr {
		userFunc(wm.ActivateApp{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmAppCommand(userFunc func(p wm.AppCommand)) {
	me.addMsg(co.WM_APPCOMMAND, func(p wm.Base) uintptr {
		userFunc(wm.AppCommand{Base: p})
		return 1
	})
}

func (me *windowDepotMsg) WmChar(userFunc func(p wm.Char)) {
	me.addMsg(co.WM_CHAR, func(p wm.Base) uintptr {
		userFunc(wm.Char{Base: p})
		return 0
	})
}

// Warning: default handled in WindowModal.
func (me *windowDepotMsg) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p wm.Base) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmCommand(cmd int32, userFunc func(p wm.Command)) {
	me.addCmd(cmd, userFunc)
}

func (me *windowDepotMsg) WmContextMenu(userFunc func(p wm.ContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p wm.Base) uintptr {
		userFunc(wm.ContextMenu{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmCreate(userFunc func(p wm.Create) int32) {
	me.addMsg(co.WM_CREATE, func(p wm.Base) uintptr {
		return uintptr(userFunc(wm.Create{Base: p}))
	})
}

func (me *windowDepotMsg) WmDeadChar(userFunc func(p wm.DeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p wm.Base) uintptr {
		userFunc(wm.DeadChar{wm.Char{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p wm.Base) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmDropFiles(userFunc func(p wm.DropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p wm.Base) uintptr {
		userFunc(wm.DropFiles{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmHelp(userFunc func(p wm.Help)) {
	me.addMsg(co.WM_HELP, func(p wm.Base) uintptr {
		userFunc(wm.Help{Base: p})
		return 1
	})
}

func (me *windowDepotMsg) WmHotKey(userFunc func(p wm.HotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p wm.Base) uintptr {
		userFunc(wm.HotKey{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmInitMenuPopup(userFunc func(p wm.InitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p wm.Base) uintptr {
		userFunc(wm.InitMenuPopup{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmKeyDown(userFunc func(p wm.KeyDown)) {
	me.addMsg(co.WM_KEYDOWN, func(p wm.Base) uintptr {
		userFunc(wm.KeyDown{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmKeyUp(userFunc func(p wm.KeyUp)) {
	me.addMsg(co.WM_KEYUP, func(p wm.Base) uintptr {
		userFunc(wm.KeyUp{wm.KeyDown{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmKillFocus(userFunc func(p wm.KillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p wm.Base) uintptr {
		userFunc(wm.KillFocus{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonDblClk(userFunc func(p wm.LButtonDblClk)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p wm.Base) uintptr {
		userFunc(wm.LButtonDblClk{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonDown(userFunc func(p wm.LButtonDown)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p wm.Base) uintptr {
		userFunc(wm.LButtonDown{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonUp(userFunc func(p wm.LButtonUp)) {
	me.addMsg(co.WM_LBUTTONUP, func(p wm.Base) uintptr {
		userFunc(wm.LButtonUp{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonDblClk(userFunc func(p wm.MButtonDblClk)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p wm.Base) uintptr {
		userFunc(wm.MButtonDblClk{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonDown(userFunc func(p wm.MButtonDown)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p wm.Base) uintptr {
		userFunc(wm.MButtonDown{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonUp(userFunc func(p wm.MButtonUp)) {
	me.addMsg(co.WM_MBUTTONUP, func(p wm.Base) uintptr {
		userFunc(wm.MButtonUp{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmMenuChar(userFunc func(p wm.MenuChar) co.MNC) {
	me.addMsg(co.WM_MENUCHAR, func(p wm.Base) uintptr {
		return uintptr(userFunc(wm.MenuChar{Base: p}))
	})
}

func (me *windowDepotMsg) WmMenuCommand(userFunc func(p wm.MenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p wm.Base) uintptr {
		userFunc(wm.MenuCommand{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmMenuSelect(userFunc func(p wm.MenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p wm.Base) uintptr {
		userFunc(wm.MenuSelect{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmMouseHover(userFunc func(p wm.MouseHover)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p wm.Base) uintptr {
		userFunc(wm.MouseHover{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(co.WM_MOUSELEAVE, func(p wm.Base) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmMouseMove(userFunc func(p wm.MouseMove)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p wm.Base) uintptr {
		userFunc(wm.MouseMove{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmMove(userFunc func(p wm.Move)) {
	me.addMsg(co.WM_MOVE, func(p wm.Base) uintptr {
		userFunc(wm.Move{Base: p})
		return 0
	})
}

// Warning: default handled in WindowMain.
func (me *windowDepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p wm.Base) uintptr {
		userFunc()
		return 0
	})
}

// Warning: default handled in WindowControl.
func (me *windowDepotMsg) WmNcPaint(userFunc func(p wm.NcPaint)) {
	me.addMsg(co.WM_NCPAINT, func(p wm.Base) uintptr {
		userFunc(wm.NcPaint{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmPaint(userFunc func()) {
	me.addMsg(co.WM_PAINT, func(p wm.Base) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmPrint(userFunc func(p wm.Print)) {
	me.addMsg(co.WM_PRINT, func(p wm.Base) uintptr {
		userFunc(wm.Print{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonDblClk(userFunc func(p wm.RButtonDblClk)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p wm.Base) uintptr {
		userFunc(wm.RButtonDblClk{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonDown(userFunc func(p wm.RButtonDown)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p wm.Base) uintptr {
		userFunc(wm.RButtonDown{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonUp(userFunc func(p wm.RButtonUp)) {
	me.addMsg(co.WM_RBUTTONUP, func(p wm.Base) uintptr {
		userFunc(wm.RButtonUp{wm.LButtonDblClk{Base: p}})
		return 0
	})
}

// Warning: default handled in WindowMain and WindowModal.
func (me *windowDepotMsg) WmSetFocus(userFunc func(p wm.SetFocus)) {
	me.addMsg(co.WM_SETFOCUS, func(p wm.Base) uintptr {
		userFunc(wm.SetFocus{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmSetFont(userFunc func(p wm.SetFont)) {
	me.addMsg(co.WM_SETFONT, func(p wm.Base) uintptr {
		userFunc(wm.SetFont{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmSize(userFunc func(p wm.Size)) {
	me.addMsg(co.WM_SIZE, func(p wm.Base) uintptr {
		userFunc(wm.Size{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmSysChar(userFunc func(p wm.SysChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p wm.Base) uintptr {
		userFunc(wm.SysChar{wm.Char{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmSysCommand(userFunc func(p wm.SysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p wm.Base) uintptr {
		userFunc(wm.SysCommand{Base: p})
		return 0
	})
}

func (me *windowDepotMsg) WmSysDeadChar(userFunc func(p wm.SysDeadChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p wm.Base) uintptr {
		userFunc(wm.SysDeadChar{wm.Char{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmSysKeyDown(userFunc func(p wm.SysKeyDown)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p wm.Base) uintptr {
		userFunc(wm.SysKeyDown{wm.KeyDown{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmSysKeyUp(userFunc func(p wm.SysKeyUp)) {
	me.addMsg(co.WM_SYSKEYUP, func(p wm.Base) uintptr {
		userFunc(wm.SysKeyUp{wm.KeyDown{Base: p}})
		return 0
	})
}

func (me *windowDepotMsg) WmTimeChange(userFunc func()) {
	me.addMsg(co.WM_TIMECHANGE, func(p wm.Base) uintptr {
		userFunc()
		return 0
	})
}
