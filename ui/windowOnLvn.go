package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"gowinui/parm"
	"unsafe"
)

func (me *windowOn) LvnDeleteAllItems(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_DELETEALLITEMS), func(p parm.WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnDeleteItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_DELETEITEM), func(p parm.WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnInsertItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_INSERTITEM), func(p parm.WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemChanged(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMCHANGED), func(p parm.WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemActivate(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMACTIVATE), func(p parm.WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnKeyDown(lv *ListView, userFunc func(p *api.NMLVKEYDOWN)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_KEYDOWN), func(p parm.WmNotify) uintptr {
		userFunc((*api.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}
