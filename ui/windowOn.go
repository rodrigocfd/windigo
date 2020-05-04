package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

// Keeps all user message handlers.
type windowOn struct {
	msgs map[c.WM]func(p parm.Raw) uintptr
	cmds map[c.ID]func(p parm.WmCommand)
	nfys map[nfyHash]func(p parm.WmNotify) uintptr

	Wm  windowOnWm
	Lvn windowOnLvn
}

// Custom hash for WM_NOTIFY messages.
type nfyHash struct {
	IdFrom c.ID
	Code   c.WM
}

// Constructor: must use.
func newWindowOn() windowOn {
	msgs := make(map[c.WM]func(p parm.Raw) uintptr)
	cmds := make(map[c.ID]func(p parm.WmCommand))
	nfys := make(map[nfyHash]func(p parm.WmNotify) uintptr)

	return windowOn{
		msgs: msgs,
		cmds: cmds,
		nfys: nfys,

		Wm:  newWindowOnWm(msgs, cmds),
		Lvn: newWindowOnLvn(nfys),
	}
}

func (won *windowOn) processMessage(p parm.Raw) (uintptr, bool) {
	switch p.Msg {
	case c.WM_COMMAND:
		paramCmd := parm.WmCommand(p)
		if userFunc, hasCmd := won.cmds[paramCmd.ControlId()]; hasCmd {
			userFunc(paramCmd)
			return 0, true
		}
	case c.WM_NOTIFY:
		paramNfy := parm.WmNotify(p)
		hash := nfyHash{
			IdFrom: c.ID(paramNfy.NmHdr().IdFrom),
			Code:   c.WM(paramNfy.NmHdr().Code),
		}
		if userFunc, hasNfy := won.nfys[hash]; hasNfy {
			return userFunc(paramNfy), true
		}
	default:
		if userFunc, hasMsg := won.msgs[p.Msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}
