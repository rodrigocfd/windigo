/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	c "wingows/consts"
)

// Keeps all user message handlers.
type windowDepotMsg struct {
	mapMsgs    map[c.WM]func(p wmBase) uintptr
	mapCmds    map[c.ID]func(p WmCommand)
	wasCreated bool // false by default, set by windowBase/controlNativeBase when the window is created
}

func (me *windowDepotMsg) addMsg(msg c.WM, userFunc func(p wmBase) uintptr) {
	if me.wasCreated {
		panic(fmt.Sprintf(
			"Cannot add message 0x%04x after the window was created.", msg))
	}
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[c.WM]func(p wmBase) uintptr)
	}
	me.mapMsgs[msg] = userFunc
}

func (me *windowDepotMsg) addCmd(cmd c.ID, userFunc func(p WmCommand)) {
	if me.wasCreated {
		panic(fmt.Sprintf(
			"Cannot add command message %d after the window was created.", cmd))
	}
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[c.ID]func(p WmCommand))
	}
	me.mapCmds[cmd] = userFunc
}

func (me *windowDepotMsg) processMessage(msg c.WM, p wmBase) (uintptr, bool) {
	if msg == c.WM_COMMAND {
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

func (me *windowDepotMsg) WmCommand(cmd c.ID, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

//------------------------------------------------------------------------------

func (me *windowDepotMsg) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(c.WM_ACTIVATE, func(p wmBase) uintptr {
		userFunc(WmActivate{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmActivateApp(userFunc func(p WmActivateApp)) {
	me.addMsg(c.WM_ACTIVATEAPP, func(p wmBase) uintptr {
		userFunc(WmActivateApp{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsg(c.WM_APPCOMMAND, func(p wmBase) uintptr {
		userFunc(WmAppCommand{base: wmBase(p)})
		return 1
	})
}

//------------------------------------------------------------------------------

func (me *windowDepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(c.WM_CHAR, func(p wmBase) uintptr {
		userFunc(WmChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(c.WM_DEADCHAR, func(p wmBase) uintptr {
		userFunc(WmDeadChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmSysDeadChar(userFunc func(p WmSysDeadChar)) {
	me.addMsg(c.WM_SYSDEADCHAR, func(p wmBase) uintptr {
		userFunc(WmSysDeadChar{wmCharBase: wmCharBase{base: wmBase(p)}})
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowDepotMsg) WmClose(userFunc func()) {
	me.addMsg(c.WM_CLOSE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmCreate(userFunc func(p WmCreate) int32) {
	me.addMsg(c.WM_CREATE, func(p wmBase) uintptr {
		return uintptr(userFunc(WmCreate{base: wmBase(p)}))
	})
}

func (me *windowDepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(c.WM_DESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(c.WM_DROPFILES, func(p wmBase) uintptr {
		userFunc(WmDropFiles{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmHelp(userFunc func(p WmHelp)) {
	me.addMsg(c.WM_HELP, func(p wmBase) uintptr {
		userFunc(WmHelp{base: wmBase(p)})
		return 1
	})
}

func (me *windowDepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(c.WM_HOTKEY, func(p wmBase) uintptr {
		userFunc(WmHotKey{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(c.WM_INITMENUPOPUP, func(p wmBase) uintptr {
		userFunc(WmInitMenuPopup{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmKeyDown(userFunc func(p WmKeyDown)) {
	me.addMsg(c.WM_KEYDOWN, func(p wmBase) uintptr {
		userFunc(WmKeyDown{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmKeyUp(userFunc func(p WmKeyUp)) {
	me.addMsg(c.WM_KEYUP, func(p wmBase) uintptr {
		userFunc(WmKeyUp{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(c.WM_KILLFOCUS, func(p wmBase) uintptr {
		userFunc(WmKillFocus{base: wmBase(p)})
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowDepotMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(c.WM_LBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmLButtonDblClk{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(c.WM_LBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmLButtonDown{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(c.WM_LBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmLButtonUp{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(c.WM_MBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmMButtonDblClk{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(c.WM_MBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmMButtonDown{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(c.WM_MBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmMButtonUp{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(c.WM_MOUSEHOVER, func(p wmBase) uintptr {
		userFunc(WmMouseHover{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(c.WM_MOUSEMOVE, func(p wmBase) uintptr {
		userFunc(WmMouseMove{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(c.WM_RBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmRButtonDblClk{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(c.WM_RBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmRButtonDown{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

func (me *windowDepotMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(c.WM_RBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmRButtonUp{wmLButtonDblClkBase: wmLButtonDblClkBase{base: wmBase(p)}})
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowDepotMsg) WmMenuChar(userFunc func(p WmMenuChar) c.MNC) {
	me.addMsg(c.WM_MENUCHAR, func(p wmBase) uintptr {
		return uintptr(userFunc(WmMenuChar{base: wmBase(p)}))
	})
}

func (me *windowDepotMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(c.WM_MOUSELEAVE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(c.WM_MOVE, func(p wmBase) uintptr {
		userFunc(WmMove{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(c.WM_NCDESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(c.WM_NCPAINT, func(p wmBase) uintptr {
		userFunc(WmNcPaint{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmPaint(userFunc func()) {
	me.addMsg(c.WM_PAINT, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowDepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(c.WM_PRINT, func(p wmBase) uintptr {
		userFunc(WmPrint{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(c.WM_SETFOCUS, func(p wmBase) uintptr {
		userFunc(WmSetFocus{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(c.WM_SETFONT, func(p wmBase) uintptr {
		userFunc(WmSetFont{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(c.WM_SIZE, func(p wmBase) uintptr {
		userFunc(WmSize{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(c.WM_SYSCOMMAND, func(p wmBase) uintptr {
		userFunc(WmSysCommand{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmSysKeyDown(userFunc func(p WmSysKeyDown)) {
	me.addMsg(c.WM_SYSKEYDOWN, func(p wmBase) uintptr {
		userFunc(WmSysKeyDown{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmSysKeyUp(userFunc func(p WmSysKeyUp)) {
	me.addMsg(c.WM_SYSKEYUP, func(p wmBase) uintptr {
		userFunc(WmSysKeyUp{base: wmBase(p)})
		return 0
	})
}

func (me *windowDepotMsg) WmTimeChange(userFunc func()) {
	me.addMsg(c.WM_TIMECHANGE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}
