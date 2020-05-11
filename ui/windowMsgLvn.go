/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

func (me *windowMsg) LvnDeleteAllItems(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEALLITEMS), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnDeleteItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEITEM), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnInsertItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_INSERTITEM), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnItemActivate(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMACTIVATE), func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnItemChanged(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGED), func(p wmBase) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnItemChanging(lv *ListView, userFunc func(p *api.NMLISTVIEW) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGING), func(p wmBase) uintptr {
		if userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowMsg) LvnKeyDown(lv *ListView, userFunc func(p *api.NMLVKEYDOWN)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_KEYDOWN), func(p wmBase) uintptr {
		userFunc((*api.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnMarqueeBegin(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_MARQUEEBEGIN), func(p wmBase) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

//------------------------------------------------------------------------------

func (me *windowMsg) LvnClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_CLICK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_DBLCLK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnHover(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM_HOVER, func(p wmBase) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowMsg) LvnKillFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_KILLFOCUS, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnRClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RCLICK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnRDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RDBLCLK, func(p wmBase) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnReleasedCapture(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RELEASEDCAPTURE, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnReturn(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RETURN, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnSetFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_SETFOCUS, func(p wmBase) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}
