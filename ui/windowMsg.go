/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	"fmt"
	c "wingows/consts"
)

// Custom hash for WM_NOTIFY messages.
type nfyHash struct {
	IdFrom c.ID
	Code   c.NM
}

// Keeps all user message handlers.
type windowMsg struct {
	msgs       map[c.WM]func(p wmBase) uintptr
	cmds       map[c.ID]func(p WmCommand)
	nfys       map[nfyHash]func(p WmNotify) uintptr
	wasCreated bool // false by default, set before the first message is handled
}

func (me *windowMsg) addMsg(msg c.WM, userFunc func(p wmBase) uintptr) {
	if me.wasCreated {
		panic(fmt.Sprintf(
			"Cannot add message 0x%04x after the window was created.", msg))
	}
	if me.msgs == nil {
		me.msgs = make(map[c.WM]func(p wmBase) uintptr)
	}
	me.msgs[msg] = userFunc
}

func (me *windowMsg) addCmd(cmd c.ID, userFunc func(p WmCommand)) {
	if me.wasCreated {
		panic(fmt.Sprintf(
			"Cannot add command message %d after the window was created.", cmd))
	}
	if me.cmds == nil {
		me.cmds = make(map[c.ID]func(p WmCommand))
	}
	me.cmds[cmd] = userFunc
}

func (me *windowMsg) addNfy(idFrom c.ID, code c.NM,
	userFunc func(p WmNotify) uintptr) {

	if me.wasCreated {
		panic(fmt.Sprintf(
			"Cannot add motify message %d/%d after the window was created.",
			idFrom, code))
	}
	if me.nfys == nil {
		me.nfys = make(map[nfyHash]func(p WmNotify) uintptr)
	}
	me.nfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowMsg) processMessage(msg c.WM, p wmBase) (uintptr, bool) {
	me.wasCreated = true // no further messages can be added

	switch msg {
	case c.WM_COMMAND:
		pCmd := WmCommand{base: p}
		if userFunc, hasCmd := me.cmds[pCmd.ControlId()]; hasCmd {
			userFunc(pCmd)
			return 0, true // always return zero
		}
	case c.WM_NOTIFY:
		pNfy := WmNotify{base: p}
		hash := nfyHash{
			IdFrom: c.ID(pNfy.NmHdr().IdFrom),
			Code:   c.NM(pNfy.NmHdr().Code),
		}
		if userFunc, hasNfy := me.nfys[hash]; hasNfy {
			return userFunc(pNfy), true
		}
	default:
		if userFunc, hasMsg := me.msgs[msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}
