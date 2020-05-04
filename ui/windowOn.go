package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

// Keeps all user message handlers.
type windowOn struct {
	msgs map[c.WM]func(p parm.Raw) uintptr
	cmds map[c.ID]func(p parm.WmCommand)
	nfys map[nfyMsgType]func(p parm.WmNotify) uintptr

	Wm windowOnWm
}

// Custom hash for WM_NOTIFY messages.
type nfyMsgType struct {
	IdFrom c.ID
	Code   uint32
}

// Constructor: must use.
func newWindowOn() windowOn {
	msgs := make(map[c.WM]func(p parm.Raw) uintptr)
	cmds := make(map[c.ID]func(p parm.WmCommand))
	nfys := make(map[nfyMsgType]func(p parm.WmNotify) uintptr)

	return windowOn{
		msgs: msgs,
		cmds: cmds,
		nfys: nfys,
		Wm:   newWindowOnWm(msgs, cmds),
	}
}
