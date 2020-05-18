/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

func (me *windowMsg) LvnBeginDrag(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINDRAG), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnBeginLabelEdit(lv *ListView, userFunc func(p *api.NMLVDISPINFO) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowMsg) LvnBeginRDrag(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINRDRAG), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnBeginScroll(lv *ListView, userFunc func(p *api.NMLVSCROLL)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINSCROLL), func(p WmNotify) uintptr {
		userFunc((*api.NMLVSCROLL)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnColumnClick(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_COLUMNCLICK), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnColumnDropDown(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_COLUMNDROPDOWN), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

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

func (me *windowMsg) LvnEndLabelEdit(lv *ListView, userFunc func(p *api.NMLVDISPINFO) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ENDLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowMsg) LvnEndScroll(lv *ListView, userFunc func(p *api.NMLVSCROLL)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ENDSCROLL), func(p WmNotify) uintptr {
		userFunc((*api.NMLVSCROLL)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnGetDispInfo(lv *ListView, userFunc func(p *api.NMLVDISPINFO)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnGetEmptyMarkup(lv *ListView, userFunc func(p *api.NMLVEMPTYMARKUP) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETEMPTYMARKUP), func(p WmNotify) uintptr {
		if userFunc((*api.NMLVEMPTYMARKUP)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
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

func (me *windowMsg) LvnIncrementalSearch(lv *ListView, userFunc func(p *api.NMLVFINDITEM) int32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_INCREMENTALSEARCH), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMLVFINDITEM)(unsafe.Pointer(p.base.LParam))))
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

func (me *windowMsg) LvnLinkClick(lv *ListView, userFunc func(p *api.NMLVLINK)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_LINKCLICK), func(p WmNotify) uintptr {
		userFunc((*api.NMLVLINK)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnMarqueeBegin(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_MARQUEEBEGIN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowMsg) LvnODCacheHint(lv *ListView, userFunc func(p *api.NMLVCACHEHINT)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODCACHEHINT), func(p WmNotify) uintptr {
		userFunc((*api.NMLVCACHEHINT)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowMsg) LvnODFindItem(lv *ListView, userFunc func(p *api.NMLVFINDITEM) int32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODFINDITEM), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMLVFINDITEM)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowMsg) LvnODStateChanged(lv *ListView, userFunc func(p *api.NMLVODSTATECHANGE)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODSTATECHANGED), func(p WmNotify) uintptr {
		userFunc((*api.NMLVODSTATECHANGE)(unsafe.Pointer(p.base.LParam)))
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

func (me *windowMsg) LvnCustomDraw(lv *ListView, userFunc func(p *api.NMCUSTOMDRAW) c.CDRF) {
	me.addNfy(lv.CtrlId(), c.NM_CUSTOMDRAW, func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMCUSTOMDRAW)(unsafe.Pointer(p.base.LParam))))
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
