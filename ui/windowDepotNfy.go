/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

type nfyHash struct { // custom hash for WM_NOTIFY messages
	IdFrom c.ID
	Code   c.NM
}

// Keeps all user common control notification handlers.
type windowDepotNfy struct {
	mapNfys    map[nfyHash]func(p WmNotify) uintptr
	wasCreated bool // false by default, set by windowBase/controlNativeBase when the window is created
}

func (me *windowDepotNfy) addNfy(idFrom c.ID, code c.NM,
	userFunc func(p WmNotify) uintptr) {

	if me.wasCreated {
		panic(fmt.Sprintf(
			"Cannot add motify message %d/%d after the window was created.",
			idFrom, code))
	}
	if me.mapNfys == nil { // guard
		me.mapNfys = make(map[nfyHash]func(p WmNotify) uintptr)
	}
	me.mapNfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowDepotNfy) processMessage(msg c.WM, p wmBase) (uintptr, bool) {
	if msg == c.WM_NOTIFY {
		pNfy := WmNotify{base: p}
		hash := nfyHash{
			IdFrom: c.ID(pNfy.NmHdr().IdFrom),
			Code:   c.NM(pNfy.NmHdr().Code),
		}
		if userFunc, hasNfy := me.mapNfys[hash]; hasNfy {
			return userFunc(pNfy), true
		}
	}

	return 0, false // no user handler found
}

//------------------------------------------------------------------------------

func (me *windowDepotNfy) LvnBeginDrag(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINDRAG), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginLabelEdit(lv *ListView, userFunc func(p *api.NMLVDISPINFO) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginRDrag(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINRDRAG), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginScroll(lv *ListView, userFunc func(p *api.NMLVSCROLL)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_BEGINSCROLL), func(p WmNotify) uintptr {
		userFunc((*api.NMLVSCROLL)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnClick(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_COLUMNCLICK), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnDropDown(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_COLUMNDROPDOWN), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnOverflowClick(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_COLUMNOVERFLOWCLICK), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnDeleteAllItems(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEALLITEMS), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnDeleteItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnEndLabelEdit(lv *ListView, userFunc func(p *api.NMLVDISPINFO) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ENDLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnEndScroll(lv *ListView, userFunc func(p *api.NMLVSCROLL)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ENDSCROLL), func(p WmNotify) uintptr {
		userFunc((*api.NMLVSCROLL)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnGetDispInfo(lv *ListView, userFunc func(p *api.NMLVDISPINFO)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnGetEmptyMarkup(lv *ListView, userFunc func(p *api.NMLVEMPTYMARKUP) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETEMPTYMARKUP), func(p WmNotify) uintptr {
		if userFunc((*api.NMLVEMPTYMARKUP)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnGetInfoTip(lv *ListView, userFunc func(p *api.NMLVGETINFOTIP)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_GETINFOTIP), func(p WmNotify) uintptr {
		userFunc((*api.NMLVGETINFOTIP)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnHotTrack(lv *ListView, userFunc func(p *api.NMLISTVIEW) int32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_HOTTRACK), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowDepotNfy) LvnIncrementalSearch(lv *ListView, userFunc func(p *api.NMLVFINDITEM) int32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_INCREMENTALSEARCH), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMLVFINDITEM)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowDepotNfy) LvnInsertItem(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_INSERTITEM), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemActivate(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMACTIVATE), func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemChanged(lv *ListView, userFunc func(p *api.NMLISTVIEW)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGED), func(p WmNotify) uintptr {
		userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemChanging(lv *ListView, userFunc func(p *api.NMLISTVIEW) bool) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ITEMCHANGING), func(p WmNotify) uintptr {
		if userFunc((*api.NMLISTVIEW)(unsafe.Pointer(p.base.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnKeyDown(lv *ListView, userFunc func(p *api.NMLVKEYDOWN)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_KEYDOWN), func(p WmNotify) uintptr {
		userFunc((*api.NMLVKEYDOWN)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnLinkClick(lv *ListView, userFunc func(p *api.NMLVLINK)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_LINKCLICK), func(p WmNotify) uintptr {
		userFunc((*api.NMLVLINK)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnMarqueeBegin(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_MARQUEEBEGIN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowDepotNfy) LvnODCacheHint(lv *ListView, userFunc func(p *api.NMLVCACHEHINT)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODCACHEHINT), func(p WmNotify) uintptr {
		userFunc((*api.NMLVCACHEHINT)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnODFindItem(lv *ListView, userFunc func(p *api.NMLVFINDITEM) int32) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODFINDITEM), func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMLVFINDITEM)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowDepotNfy) LvnODStateChanged(lv *ListView, userFunc func(p *api.NMLVODSTATECHANGE)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_ODSTATECHANGED), func(p WmNotify) uintptr {
		userFunc((*api.NMLVODSTATECHANGE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnSetDispInfo(lv *ListView, userFunc func(p *api.NMLVDISPINFO)) {
	me.addNfy(lv.CtrlId(), c.NM(c.LVN_SETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*api.NMLVDISPINFO)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowDepotNfy) LvnClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnCustomDraw(lv *ListView, userFunc func(p *api.NMCUSTOMDRAW) c.CDRF) {
	me.addNfy(lv.CtrlId(), c.NM_CUSTOMDRAW, func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMCUSTOMDRAW)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowDepotNfy) LvnDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnHover(lv *ListView, userFunc func(p *api.NMHDR) uint32) {
	me.addNfy(lv.CtrlId(), c.NM_HOVER, func(p WmNotify) uintptr {
		return uintptr(userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam))))
	})
}

func (me *windowDepotNfy) LvnKillFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnRClick(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnRDblClk(lv *ListView, userFunc func(p *api.NMITEMACTIVATE)) {
	me.addNfy(lv.CtrlId(), c.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*api.NMITEMACTIVATE)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnReleasedCapture(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RELEASEDCAPTURE, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnReturn(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_RETURN, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnSetFocus(lv *ListView, userFunc func(p *api.NMHDR)) {
	me.addNfy(lv.CtrlId(), c.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*api.NMHDR)(unsafe.Pointer(p.base.LParam)))
		return 0
	})
}
