package ui

import (
	c "gowinui/consts"
	"gowinui/parm"
)

func (me *windowOn) LvnDeleteAllItems(lv *ListView, userFunc func(p parm.LvnDeleteAllItems)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_DELETEALLITEMS), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnDeleteAllItems(p))
		return 0
	})
}

func (me *windowOn) LvnItemChanged(lv *ListView, userFunc func(p parm.LvnItemChanged)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMCHANGING), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnItemChanged(p))
		return 0
	})
}
