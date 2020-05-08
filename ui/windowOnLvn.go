package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

func (me *windowOn) LvnDeleteAllItems(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_DELETEALLITEMS), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnDeleteItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_DELETEITEM), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnInsertItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_INSERTITEM), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemActivate(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMACTIVATE), func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemChanged(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_ITEMCHANGED), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnKeyDown(lv *ListView, userFunc func(p *api.NMLVKEYDOWN)) {
	me.addNfy(lv.CtrlId(), int32(c.LVN_KEYDOWN), func(p wmBase) uintptr {
		userFunc((*api.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}
