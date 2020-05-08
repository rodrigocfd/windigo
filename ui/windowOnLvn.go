package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

func (me *windowOn) LvnDeleteAllItems(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEALLITEMS), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnDeleteItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEITEM), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnInsertItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_INSERTITEM), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemActivate(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMACTIVATE), func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemChanged(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGED), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnItemChanging(lv *ListView, userFunc func(p *api.NMLISTVIEW) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGING), func(p wmBase) uintptr {
		if userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowOn) LvnKeyDown(lv *ListView, userFunc func(p *api.NMLVKEYDOWN)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_KEYDOWN), func(p wmBase) uintptr {
		userFunc((*api.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnMarqueeBegin(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_MARQUEEBEGIN), func(p wmBase) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

//------------------------------------------------------------------------------

func (me *windowOn) LvnClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_CLICK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_DBLCLK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnHover(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM_HOVER, func(p wmBase) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowOn) LvnKillFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_KILLFOCUS, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnRClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RCLICK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnRDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RDBLCLK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnReleasedCapture(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RELEASEDCAPTURE, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnReturn(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RETURN, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowOn) LvnSetFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_SETFOCUS, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}
