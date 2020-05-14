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

func (me *windowMsg) LvnColumnOverflowClick(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_COLUMNOVERFLOWCLICK), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnDeleteAllItems(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEALLITEMS), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnDeleteItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnGetDispInfo(lv *ListView, userFunc func(p *api.NMLVDISPINFO)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnGetInfoTip(lv *ListView, userFunc func(p *api.NMLVGETINFOTIP)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETINFOTIP), func(p WmNotify) uintptr {
		userFunc((*api.NMLVGETINFOTIP)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnHotTrack(lv *ListView, userFunc func(p *api.NMLISTVIEW) int32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_HOTTRACK), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowMsg) LvnInsertItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_INSERTITEM), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnItemActivate(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMACTIVATE), func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnItemChanged(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGED), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnItemChanging(lv *ListView, userFunc func(p *api.NMLISTVIEW) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGING), func(p WmNotify) uintptr {
		if userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowMsg) LvnKeyDown(lv *ListView, userFunc func(p *api.NMLVKEYDOWN)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_KEYDOWN), func(p WmNotify) uintptr {
		userFunc((*api.NMLVKEYDOWN)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnMarqueeBegin(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_MARQUEEBEGIN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowMsg) LvnOdCacheHint(lv *ListView, userFunc func(p *api.NMLVCACHEHINT)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODCACHEHINT), func(p WmNotify) uintptr {
		userFunc((*api.NMLVCACHEHINT)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnSetDispInfo(lv *ListView, userFunc func(p *api.NMLVDISPINFO)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_SETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowMsg) LvnClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnHover(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM_HOVER, func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowMsg) LvnKillFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnRClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnRDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnReleasedCapture(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RELEASEDCAPTURE, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnReturn(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RETURN, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnSetFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}
