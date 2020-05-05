package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

func (won *windowOn) WmCommand(cmd c.ID, userFunc func(p parm.WmCommand)) {
	won.cmds[cmd] = userFunc
}

func (won *windowOn) WmCreate(userFunc func(p parm.WmCreate) uintptr) {
	won.msgs[c.WM_CREATE] = func(p parm.Raw) uintptr {
		return userFunc(parm.WmCreate(p))
	}
}

func (won *windowOn) WmDestroy(userFunc func(p parm.WmDestroy)) {
	won.msgs[c.WM_DESTROY] = func(p parm.Raw) uintptr {
		userFunc(parm.WmDestroy(p))
		return 0
	}
}

func (won *windowOn) WmNcDestroy(userFunc func(p parm.WmNcDestroy)) {
	won.msgs[c.WM_NCDESTROY] = func(p parm.Raw) uintptr {
		userFunc(parm.WmNcDestroy(p))
		return 0
	}
}
