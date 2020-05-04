package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

// Allows user to add WM message handlers.
type windowOnWm struct {
	msgs map[c.WM]func(p parm.Raw) uintptr
	cmds map[c.ID]func(p parm.WmCommand)
}

// Constructor: must use.
func newWindowOnWm(msgs map[c.WM]func(p parm.Raw) uintptr,
	cmds map[c.ID]func(p parm.WmCommand)) windowOnWm {

	return windowOnWm{
		msgs: msgs,
		cmds: cmds,
	}
}

func (won *windowOnWm) Command(cmd c.ID, userFunc func(p parm.WmCommand)) {
	won.cmds[cmd] = userFunc
}

func (won *windowOnWm) Create(userFunc func(p parm.WmCreate) uintptr) {
	won.msgs[c.WM_CREATE] = func(p parm.Raw) uintptr {
		return userFunc(parm.WmCreate(p))
	}
}

func (won *windowOnWm) Destroy(userFunc func(p parm.WmDestroy)) {
	won.msgs[c.WM_DESTROY] = func(p parm.Raw) uintptr {
		userFunc(parm.WmDestroy(p))
		return 0
	}
}
