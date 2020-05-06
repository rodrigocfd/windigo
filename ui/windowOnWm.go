package ui

import (
	c "gowinui/consts"
	"gowinui/parm"
)

func (me *windowOn) WmCommand(cmd c.ID, userFunc func(p parm.WmCommand)) {
	me.cmds[cmd] = userFunc
}

func (me *windowOn) WmActivate(userFunc func(p parm.WmActivate)) {
	me.msgs[c.WM_ACTIVATE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmActivate(p))
		return 0
	}
}

func (me *windowOn) WmClose(userFunc func(p parm.WmClose)) {
	me.msgs[c.WM_CLOSE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmClose(p))
		return 0
	}
}

func (me *windowOn) WmCreate(userFunc func(p parm.WmCreate) uintptr) {
	me.msgs[c.WM_CREATE] = func(p parm.Raw) uintptr {
		return userFunc(parm.WmCreate(p))
	}
}

func (me *windowOn) WmDestroy(userFunc func(p parm.WmDestroy)) {
	me.msgs[c.WM_DESTROY] = func(p parm.Raw) uintptr {
		userFunc(parm.WmDestroy(p))
		return 0
	}
}

func (me *windowOn) WmInitMenuPopup(userFunc func(p parm.WmInitMenuPopup)) {
	me.msgs[c.WM_INITMENUPOPUP] = func(p parm.Raw) uintptr {
		userFunc(parm.WmInitMenuPopup(p))
		return 0
	}
}

func (me *windowOn) WmLButtonDblClk(userFunc func(p parm.WmLButtonDblClk)) {
	me.msgs[c.WM_LBUTTONDBLCLK] = func(p parm.Raw) uintptr {
		userFunc(parm.WmLButtonDblClk(p))
		return 0
	}
}
func (me *windowOn) WmLButtonDown(userFunc func(p parm.WmLButtonDown)) {
	me.msgs[c.WM_LBUTTONDOWN] = func(p parm.Raw) uintptr {
		userFunc(parm.WmLButtonDown{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}
func (me *windowOn) WmLButtonUp(userFunc func(p parm.WmLButtonUp)) {
	me.msgs[c.WM_LBUTTONUP] = func(p parm.Raw) uintptr {
		userFunc(parm.WmLButtonUp{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}

func (me *windowOn) WmMButtonDblClk(userFunc func(p parm.WmMButtonDblClk)) {
	me.msgs[c.WM_MBUTTONDBLCLK] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMButtonDblClk{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}
func (me *windowOn) WmMButtonDown(userFunc func(p parm.WmMButtonDown)) {
	me.msgs[c.WM_MBUTTONDOWN] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMButtonDown{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}
func (me *windowOn) WmMButtonUp(userFunc func(p parm.WmMButtonUp)) {
	me.msgs[c.WM_MBUTTONUP] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMButtonUp{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}

func (me *windowOn) WmMouseHover(userFunc func(p parm.WmMouseHover)) {
	me.msgs[c.WM_MOUSEHOVER] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMouseHover{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}

func (me *windowOn) WmMouseMove(userFunc func(p parm.WmMouseMove)) {
	me.msgs[c.WM_MOUSEMOVE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMouseMove{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}

func (me *windowOn) WmRButtonDblClk(userFunc func(p parm.WmRButtonDblClk)) {
	me.msgs[c.WM_RBUTTONDBLCLK] = func(p parm.Raw) uintptr {
		userFunc(parm.WmRButtonDblClk{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}
func (me *windowOn) WmRButtonDown(userFunc func(p parm.WmRButtonDown)) {
	me.msgs[c.WM_RBUTTONDOWN] = func(p parm.Raw) uintptr {
		userFunc(parm.WmRButtonDown{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}
func (me *windowOn) WmRButtonUp(userFunc func(p parm.WmRButtonUp)) {
	me.msgs[c.WM_RBUTTONUP] = func(p parm.Raw) uintptr {
		userFunc(parm.WmRButtonUp{WmLButtonDblClk: parm.WmLButtonDblClk(p)})
		return 0
	}
}

func (me *windowOn) WmMouseLeave(userFunc func(p parm.WmMouseLeave)) {
	me.msgs[c.WM_MOUSELEAVE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMouseLeave(p))
		return 0
	}
}

func (me *windowOn) WmMove(userFunc func(p parm.WmMove)) {
	me.msgs[c.WM_MOVE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmMove(p))
		return 0
	}
}

func (me *windowOn) WmNcDestroy(userFunc func(p parm.WmNcDestroy)) {
	me.msgs[c.WM_NCDESTROY] = func(p parm.Raw) uintptr {
		userFunc(parm.WmNcDestroy(p))
		return 0
	}
}

func (me *windowOn) WmNcPaint(userFunc func(p parm.WmNcPaint)) {
	me.msgs[c.WM_NCPAINT] = func(p parm.Raw) uintptr {
		userFunc(parm.WmNcPaint(p))
		return 0
	}
}

func (me *windowOn) WmPaint(userFunc func(p parm.WmPaint)) {
	me.msgs[c.WM_PAINT] = func(p parm.Raw) uintptr {
		userFunc(parm.WmPaint(p))
		return 0
	}
}

func (me *windowOn) WmSetFocus(userFunc func(p parm.WmSetFocus)) {
	me.msgs[c.WM_SETFOCUS] = func(p parm.Raw) uintptr {
		userFunc(parm.WmSetFocus(p))
		return 0
	}
}

func (me *windowOn) WmSetFont(userFunc func(p parm.WmSetFont)) {
	me.msgs[c.WM_SETFONT] = func(p parm.Raw) uintptr {
		userFunc(parm.WmSetFont(p))
		return 0
	}
}

func (me *windowOn) WmSize(userFunc func(p parm.WmSize)) {
	me.msgs[c.WM_SIZE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmSize(p))
		return 0
	}
}
