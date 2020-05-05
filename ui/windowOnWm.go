package ui

import (
	c "winffi/consts"
	"winffi/parm"
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

func (me *windowOn) WmNcDestroy(userFunc func(p parm.WmNcDestroy)) {
	me.msgs[c.WM_NCDESTROY] = func(p parm.Raw) uintptr {
		userFunc(parm.WmNcDestroy(p))
		return 0
	}
}

func (me *windowOn) WmSize(userFunc func(p parm.WmSize)) {
	me.msgs[c.WM_SIZE] = func(p parm.Raw) uintptr {
		userFunc(parm.WmSize(p))
		return 0
	}
}
