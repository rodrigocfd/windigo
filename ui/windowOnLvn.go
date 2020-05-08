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
func (me *windowOn) LvnDeleteItem(lv *ListView, userFunc func(p parm.LvnDeleteItem)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_DELETEITEM), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnDeleteItem{LvnDeleteAllItems: parm.LvnDeleteAllItems(p)})
		return 0
	})
}
func (me *windowOn) LvnInsertItem(lv *ListView, userFunc func(p parm.LvnInsertItem)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_INSERTITEM), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnInsertItem{LvnDeleteAllItems: parm.LvnDeleteAllItems(p)})
		return 0
	})
}
func (me *windowOn) LvnItemChanged(lv *ListView, userFunc func(p parm.LvnItemChanged)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMCHANGED), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnItemChanged{LvnDeleteAllItems: parm.LvnDeleteAllItems(p)})
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowOn) LvnItemActivate(lv *ListView, userFunc func(p parm.LvnItemActivate)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMACTIVATE), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnItemActivate(p))
		return 0
	})
}

func (me *windowOn) LvnKeyDown(lv *ListView, userFunc func(p parm.LvnKeyDown)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_KEYDOWN), func(p parm.WmNotify) uintptr {
		userFunc(parm.LvnKeyDown(p))
		return 0
	})
}
