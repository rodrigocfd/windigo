package ui

import (
	c "winffi/consts"
)

type windowOnWm struct {
	msgs map[c.WM]func(p Param) uintptr
	cmds map[c.ID]func(p ParamCommand)
}

// Constructor: must use.
func newWindowOnWm(msgs map[c.WM]func(p Param) uintptr,
	cmds map[c.ID]func(p ParamCommand)) windowOnWm {

	return windowOnWm{
		msgs: msgs,
		cmds: cmds,
	}
}

func (won *windowOnWm) Command(cmd c.ID, userFunc func(p ParamCommand)) {
	won.cmds[cmd] = userFunc
}

func (won *windowOnWm) Create(userFunc func(p ParamCreate) uintptr) {
	won.msgs[c.WM_CREATE] = func(p Param) uintptr {
		return userFunc(ParamCreate(p))
	}
}

func (won *windowOnWm) Destroy(userFunc func(p ParamDestroy)) {
	won.msgs[c.WM_DESTROY] = func(p Param) uintptr {
		userFunc(ParamDestroy(p))
		return 0
	}
}
