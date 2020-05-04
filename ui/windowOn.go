package ui

import (
	c "winffi/consts"
)

type windowOn struct {
	msgs map[c.WM]func(p Param) uintptr
	cmds map[c.ID]func(p ParamCommand)
	nfys map[nfyMsgType]func(p ParamNotify) uintptr

	Wm windowOnWm
}

// Custom hash for WM_NOTIFY messages.
type nfyMsgType struct {
	IdFrom c.ID
	Code   uint32
}

// Constructor: must use.
func newWindowOn() windowOn {
	msgs := make(map[c.WM]func(p Param) uintptr)
	cmds := make(map[c.ID]func(p ParamCommand))
	nfys := make(map[nfyMsgType]func(p ParamNotify) uintptr)

	return windowOn{
		msgs: msgs,
		cmds: cmds,
		nfys: nfys,
		Wm:   newWindowOnWm(msgs, cmds),
	}
}
