package ui

import (
	"fmt"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// Custom hash for WM_NOTIFY messages.
type nfyHash struct {
	IdFrom c.ID
	Code   c.NM
}

// Keeps all user message handlers.
type windowMsg struct {
	msgs        map[c.WM]func(p wmBase) uintptr
	cmds        map[c.ID]func(p *WmCommand)
	nfys        map[nfyHash]func(p wmBase) uintptr
	loopStarted bool
}

func makeWindowMsg() windowMsg {
	msgs := make(map[c.WM]func(p wmBase) uintptr)
	cmds := make(map[c.ID]func(p *WmCommand))
	nfys := make(map[nfyHash]func(p wmBase) uintptr)

	return windowMsg{
		msgs:        msgs,
		cmds:        cmds,
		nfys:        nfys,
		loopStarted: false,
	}
}

func (me *windowMsg) addMsg(msg c.WM, userFunc func(p wmBase) uintptr) {
	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add message 0x%04x after application loop started.", msg))
	}
	me.msgs[msg] = userFunc
}

func (me *windowMsg) addCmd(cmd c.ID, userFunc func(p *WmCommand)) {
	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add command message %d after application loop started.", cmd))
	}
	me.cmds[cmd] = userFunc
}

func (me *windowMsg) addNfy(idFrom c.ID, code c.NM,
	userFunc func(p wmBase) uintptr) {

	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add motify message %d/%d after application loop started.",
			idFrom, code))
	}
	me.nfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowMsg) processMessage(p wmBase) (uintptr, bool) {
	switch p.Msg {
	case c.WM_COMMAND:
		cmdId := c.ID(api.LoWord(uint32(p.WParam)))
		if userFunc, hasCmd := me.cmds[cmdId]; hasCmd {
			userFunc(newWmCommand(p))
			return 0, true // always return zero
		}
	case c.WM_NOTIFY:
		nmHdr := (*api.NMHDR)(unsafe.Pointer(p.LParam))
		hash := nfyHash{
			IdFrom: c.ID(nmHdr.IdFrom),
			Code:   c.NM(nmHdr.Code),
		}
		if userFunc, hasNfy := me.nfys[hash]; hasNfy {
			return userFunc(p), true
		}
	default:
		if userFunc, hasMsg := me.msgs[p.Msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}
