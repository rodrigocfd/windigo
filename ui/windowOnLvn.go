package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

func (won *windowOn) LvnDeleteAllItems(cid c.ID, userFunc func(p parm.LvnDeleteAllItems)) {
	won.nfys[nfyHash{IdFrom: cid, Code: c.WM(c.LVN_DELETEALLITEMS)}] = func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnDeleteAllItems(p))
		return 0
	}
}

func (won *windowOn) LvnItemChanged(cid c.ID, userFunc func(p parm.LvnItemChanged)) {
	won.nfys[nfyHash{IdFrom: cid, Code: c.WM(c.LVN_ITEMCHANGING)}] = func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnItemChanged(p))
		return 0
	}
}
