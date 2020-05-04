package ui

import (
	c "winffi/consts"
	"winffi/parm"
)

// Allows user to add LVN message handlers.
type windowOnLvn struct {
	nfys map[nfyHash]func(p parm.WmNotify) uintptr // this is just a pointer
}

// Constructor: must use.
func newWindowOnLvn(nfys map[nfyHash]func(p parm.WmNotify) uintptr) windowOnLvn {
	return windowOnLvn{
		nfys: nfys,
	}
}

func (won *windowOnLvn) DeleteAllItems(cid c.ID, userFunc func(p parm.LvnDeleteAllItems)) {
	won.nfys[nfyHash{IdFrom: cid, Code: c.WM(c.LVN_DELETEALLITEMS)}] = func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnDeleteAllItems(p))
		return 0
	}
}

func (won *windowOnLvn) ItemChanged(cid c.ID, userFunc func(p parm.LvnItemChanged)) {
	won.nfys[nfyHash{IdFrom: cid, Code: c.WM(c.LVN_ITEMCHANGING)}] = func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnItemChanged(p))
		return 0
	}
}
