/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	c "wingows/consts"
)

func (me *windowMsg) WmCommand(cmd c.ID, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

//------------------------------------------------------------------------------

func (me *windowMsg) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(c.WM_ACTIVATE, func(p wmBase) uintptr {
		userFunc(WmActivate{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmClose(userFunc func()) {
	me.addMsg(c.WM_CLOSE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmCreate(userFunc func(p WmCreate) int32) {
	me.addMsg(c.WM_CREATE, func(p wmBase) uintptr {
		return uintptr(userFunc(WmCreate{base: wmBase(p)}))
	})
}

func (me *windowMsg) WmDestroy(userFunc func()) {
	me.addMsg(c.WM_DESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(c.WM_DROPFILES, func(p wmBase) uintptr {
		userFunc(WmDropFiles{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(c.WM_INITMENUPOPUP, func(p wmBase) uintptr {
		userFunc(WmInitMenuPopup{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(c.WM_LBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmLButtonDblClk{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(c.WM_LBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmLButtonDown{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(c.WM_LBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmLButtonUp{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(c.WM_MBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmMButtonDblClk{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(c.WM_MBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmMButtonDown{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(c.WM_MBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmMButtonUp{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(c.WM_MOUSEHOVER, func(p wmBase) uintptr {
		userFunc(WmMouseHover{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(c.WM_MOUSEMOVE, func(p wmBase) uintptr {
		userFunc(WmMouseMove{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(c.WM_RBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(WmRButtonDblClk{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(c.WM_RBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(WmRButtonDown{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(c.WM_RBUTTONUP, func(p wmBase) uintptr {
		userFunc(WmRButtonUp{wmBaseBtn: wmBaseBtn{base: wmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(c.WM_MOUSELEAVE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(c.WM_MOVE, func(p wmBase) uintptr {
		userFunc(WmMove{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(c.WM_NCDESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(c.WM_NCPAINT, func(p wmBase) uintptr {
		userFunc(WmNcPaint{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmPaint(userFunc func()) {
	me.addMsg(c.WM_PAINT, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(c.WM_SETFOCUS, func(p wmBase) uintptr {
		userFunc(WmSetFocus{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(c.WM_SETFONT, func(p wmBase) uintptr {
		userFunc(WmSetFont{base: wmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(c.WM_SIZE, func(p wmBase) uintptr {
		userFunc(WmSize{base: wmBase(p)})
		return 0
	})
}
