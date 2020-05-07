package ui

import (
	"fmt"
	c "gowinui/consts"
	"gowinui/parm"
)

// Custom hash for WM_NOTIFY messages.
type nfyHash struct {
	IdFrom c.ID
	Code   int32
}

// Keeps all user message handlers.
type windowOn struct {
	msgs        map[c.WM]func(p parm.Raw) uintptr
	cmds        map[c.ID]func(p parm.WmCommand)
	nfys        map[nfyHash]func(p parm.WmNotify) uintptr
	loopStarted bool
}

func makeWindowOn() windowOn {
	msgs := make(map[c.WM]func(p parm.Raw) uintptr)
	cmds := make(map[c.ID]func(p parm.WmCommand))
	nfys := make(map[nfyHash]func(p parm.WmNotify) uintptr)

	return windowOn{
		msgs:        msgs,
		cmds:        cmds,
		nfys:        nfys,
		loopStarted: false,
	}
}

func (me *windowOn) addMsg(msg c.WM, userFunc func(p parm.Raw) uintptr) {
	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add message 0x%04x after application loop started.", msg))
	}
	me.msgs[msg] = userFunc
}

func (me *windowOn) addNfy(idFrom c.ID, code int32,
	userFunc func(p parm.WmNotify) uintptr) {

	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add motify message %d/%d after application loop started.",
			idFrom, code))
	}
	me.nfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowOn) processMessage(p parm.Raw) (uintptr, bool) {
	switch p.Msg {
	case c.WM_COMMAND:
		paramCmd := parm.WmCommand(p)
		if userFunc, hasCmd := me.cmds[paramCmd.ControlId()]; hasCmd {
			userFunc(paramCmd)
			return 0, true
		}
	case c.WM_NOTIFY:
		paramNfy := parm.WmNotify(p)
		hash := nfyHash{
			IdFrom: c.ID(paramNfy.NmHdr().IdFrom),
			Code:   int32(paramNfy.NmHdr().Code),
		}
		if userFunc, hasNfy := me.nfys[hash]; hasNfy {
			return userFunc(paramNfy), true
		}
	default:
		if userFunc, hasMsg := me.msgs[p.Msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}
