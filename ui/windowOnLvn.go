package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

func (me *windowOn) LvnDeleteAllItems(cid c.ID, userFunc func(p parm.LvnDeleteAllItems)) {
	me.nfys[nfyHash{IdFrom: cid, Code: c.WM(c.LVN_DELETEALLITEMS)}] = func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnDeleteAllItems(p))
		return 0
	}
}

func (me *windowOn) LvnItemChanged(cid c.ID, userFunc func(p parm.LvnItemChanged)) {
	me.nfys[nfyHash{IdFrom: cid, Code: c.WM(c.LVN_ITEMCHANGING)}] = func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnItemChanged(p))
		return 0
	}
}
