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
	me.addMsg(c.WM_ACTIVATE, func(p WmBase) uintptr {
		userFunc(WmActivate{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmClose(userFunc func()) {
	me.addMsg(c.WM_CLOSE, func(p WmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmCreate(userFunc func(p WmCreate) int32) {
	me.addMsg(c.WM_CREATE, func(p WmBase) uintptr {
		return uintptr(userFunc(WmCreate{base: WmBase(p)}))
	})
}

func (me *windowMsg) WmDestroy(userFunc func()) {
	me.addMsg(c.WM_DESTROY, func(p WmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(c.WM_DROPFILES, func(p WmBase) uintptr {
		userFunc(WmDropFiles{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(c.WM_INITMENUPOPUP, func(p WmBase) uintptr {
		userFunc(WmInitMenuPopup{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(c.WM_LBUTTONDBLCLK, func(p WmBase) uintptr {
		userFunc(WmLButtonDblClk{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(c.WM_LBUTTONDOWN, func(p WmBase) uintptr {
		userFunc(WmLButtonDown{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(c.WM_LBUTTONUP, func(p WmBase) uintptr {
		userFunc(WmLButtonUp{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(c.WM_MBUTTONDBLCLK, func(p WmBase) uintptr {
		userFunc(WmMButtonDblClk{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(c.WM_MBUTTONDOWN, func(p WmBase) uintptr {
		userFunc(WmMButtonDown{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(c.WM_MBUTTONUP, func(p WmBase) uintptr {
		userFunc(WmMButtonUp{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(c.WM_MOUSEHOVER, func(p WmBase) uintptr {
		userFunc(WmMouseHover{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(c.WM_MOUSEMOVE, func(p WmBase) uintptr {
		userFunc(WmMouseMove{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(c.WM_RBUTTONDBLCLK, func(p WmBase) uintptr {
		userFunc(WmRButtonDblClk{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(c.WM_RBUTTONDOWN, func(p WmBase) uintptr {
		userFunc(WmRButtonDown{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(c.WM_RBUTTONUP, func(p WmBase) uintptr {
		userFunc(WmRButtonUp{wmBaseBtn: wmBaseBtn{base: WmBase(p)}})
		return 0
	})
}

func (me *windowMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(c.WM_MOUSELEAVE, func(p WmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(c.WM_MOVE, func(p WmBase) uintptr {
		userFunc(WmMove{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(c.WM_NCDESTROY, func(p WmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(c.WM_NCPAINT, func(p WmBase) uintptr {
		userFunc(WmNcPaint{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmPaint(userFunc func()) {
	me.addMsg(c.WM_PAINT, func(p WmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(c.WM_SETFOCUS, func(p WmBase) uintptr {
		userFunc(WmSetFocus{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(c.WM_SETFONT, func(p WmBase) uintptr {
		userFunc(WmSetFont{base: WmBase(p)})
		return 0
	})
}

func (me *windowMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(c.WM_SIZE, func(p WmBase) uintptr {
		userFunc(WmSize{base: WmBase(p)})
		return 0
	})
}
