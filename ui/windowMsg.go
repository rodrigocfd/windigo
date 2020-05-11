/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

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
	loopStarted bool // false by default, set once by WindowMain
}

func (me *windowMsg) addMsg(msg c.WM, userFunc func(p wmBase) uintptr) {
	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add message 0x%04x after application loop started.", msg))
	}
	if me.msgs == nil {
		me.msgs = make(map[c.WM]func(p wmBase) uintptr)
	}
	me.msgs[msg] = userFunc
}

func (me *windowMsg) addCmd(cmd c.ID, userFunc func(p *WmCommand)) {
	if me.loopStarted {
		panic(fmt.Sprintf(
			"Cannot add command message %d after application loop started.", cmd))
	}
	if me.cmds == nil {
		me.cmds = make(map[c.ID]func(p *WmCommand))
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
	if me.nfys == nil {
		me.nfys = make(map[nfyHash]func(p wmBase) uintptr)
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
