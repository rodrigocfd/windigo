/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/gui/wm"
	"wingows/win"
)

type nfyHash struct { // custom hash for WM_NOTIFY messages
	IdFrom int32
	Code   co.NM
}

// Keeps all user common control notification handlers.
type windowDepotNfy struct {
	mapNfys map[nfyHash]func(p wm.Notify) uintptr
}

func (me *windowDepotNfy) addNfy(idFrom int32, code co.NM,
	userFunc func(p wm.Notify) uintptr) {

	if me.mapNfys == nil { // guard
		me.mapNfys = make(map[nfyHash]func(p wm.Notify) uintptr, 16) // arbitrary capacity
	}
	me.mapNfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowDepotNfy) processMessage(msg co.WM, p wm.Base) (uintptr, bool) {
	if msg == co.WM_NOTIFY {
		pNfy := wm.Notify(p)
		hash := nfyHash{
			IdFrom: int32(pNfy.NmHdr().IdFrom),
			Code:   co.NM(pNfy.NmHdr().Code),
		}
		if userFunc, hasNfy := me.mapNfys[hash]; hasNfy {
			return userFunc(pNfy), true
		}
	}

	return 0, false // no user handler found
}

//------------------------------------------------------------- ListView LVN ---

func (me *windowDepotNfy) LvnBeginDrag(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_BEGINDRAG), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginLabelEdit(lv *ListView, userFunc func(p *win.NMLVDISPINFO) bool) {
	me.addNfy(lv.Id(), co.NM(co.LVN_BEGINLABELEDIT), func(p wm.Notify) uintptr {
		if userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginRDrag(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_BEGINRDRAG), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginScroll(lv *ListView, userFunc func(p *win.NMLVSCROLL)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_BEGINSCROLL), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVSCROLL)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnClick(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_COLUMNCLICK), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnDropDown(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_COLUMNDROPDOWN), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnOverflowClick(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_COLUMNOVERFLOWCLICK), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnDeleteAllItems(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_DELETEALLITEMS), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnDeleteItem(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_DELETEITEM), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnEndLabelEdit(lv *ListView, userFunc func(p *win.NMLVDISPINFO) bool) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ENDLABELEDIT), func(p wm.Notify) uintptr {
		if userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnEndScroll(lv *ListView, userFunc func(p *win.NMLVSCROLL)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ENDSCROLL), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVSCROLL)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnGetDispInfo(lv *ListView, userFunc func(p *win.NMLVDISPINFO)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_GETDISPINFO), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnGetEmptyMarkup(lv *ListView, userFunc func(p *win.NMLVEMPTYMARKUP) bool) {
	me.addNfy(lv.Id(), co.NM(co.LVN_GETEMPTYMARKUP), func(p wm.Notify) uintptr {
		if userFunc((*win.NMLVEMPTYMARKUP)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnGetInfoTip(lv *ListView, userFunc func(p *win.NMLVGETINFOTIP)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_GETINFOTIP), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVGETINFOTIP)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnHotTrack(lv *ListView, userFunc func(p *win.NMLISTVIEW) int32) {
	me.addNfy(lv.Id(), co.NM(co.LVN_HOTTRACK), func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnIncrementalSearch(lv *ListView, userFunc func(p *win.NMLVFINDITEM) int32) {
	me.addNfy(lv.Id(), co.NM(co.LVN_INCREMENTALSEARCH), func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnInsertItem(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_INSERTITEM), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemActivate(lv *ListView, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ITEMACTIVATE), func(p wm.Notify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemChanged(lv *ListView, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ITEMCHANGED), func(p wm.Notify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemChanging(lv *ListView, userFunc func(p *win.NMLISTVIEW) bool) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ITEMCHANGING), func(p wm.Notify) uintptr {
		if userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnKeyDown(lv *ListView, userFunc func(p *win.NMLVKEYDOWN)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_KEYDOWN), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnLinkClick(lv *ListView, userFunc func(p *win.NMLVLINK)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_LINKCLICK), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVLINK)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnMarqueeBegin(lv *ListView, userFunc func(p *win.NMHDR) uint32) {
	me.addNfy(lv.Id(), co.NM(co.LVN_MARQUEEBEGIN), func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnODCacheHint(lv *ListView, userFunc func(p *win.NMLVCACHEHINT)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ODCACHEHINT), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVCACHEHINT)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnODFindItem(lv *ListView, userFunc func(p *win.NMLVFINDITEM) int32) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ODFINDITEM), func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnODStateChanged(lv *ListView, userFunc func(p *win.NMLVODSTATECHANGE)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_ODSTATECHANGED), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVODSTATECHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnSetDispInfo(lv *ListView, userFunc func(p *win.NMLVDISPINFO)) {
	me.addNfy(lv.Id(), co.NM(co.LVN_SETDISPINFO), func(p wm.Notify) uintptr {
		userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//-------------------------------------------------------------- ListView NM ---

func (me *windowDepotNfy) LvnClick(lv *ListView, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(lv.Id(), co.NM_CLICK, func(p wm.Notify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnCustomDraw(lv *ListView, userFunc func(p *win.NMLVCUSTOMDRAW) co.CDRF) {
	me.addNfy(lv.Id(), co.NM_CUSTOMDRAW, func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMLVCUSTOMDRAW)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnDblClk(lv *ListView, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(lv.Id(), co.NM_DBLCLK, func(p wm.Notify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnHover(lv *ListView, userFunc func(p *win.NMHDR) uint32) {
	me.addNfy(lv.Id(), co.NM_HOVER, func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnKillFocus(lv *ListView, userFunc func(p *win.NMHDR)) {
	me.addNfy(lv.Id(), co.NM_KILLFOCUS, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnRClick(lv *ListView, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(lv.Id(), co.NM_RCLICK, func(p wm.Notify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnRDblClk(lv *ListView, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(lv.Id(), co.NM_RDBLCLK, func(p wm.Notify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnReleasedCapture(lv *ListView, userFunc func(p *win.NMHDR)) {
	me.addNfy(lv.Id(), co.NM_RELEASEDCAPTURE, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnReturn(lv *ListView, userFunc func(p *win.NMHDR)) {
	me.addNfy(lv.Id(), co.NM_RETURN, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnSetFocus(lv *ListView, userFunc func(p *win.NMHDR)) {
	me.addNfy(lv.Id(), co.NM_SETFOCUS, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------ StatusBar SBN ---

func (me *windowDepotNfy) SbnSimpleModeChange(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM(co.SBN_SIMPLEMODECHANGE), func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- StatusBar NM ---

func (me *windowDepotNfy) SbnClick(sb *StatusBar, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(sb.Id(), co.NM_CLICK, func(p wm.Notify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) SbnDblClk(sb *StatusBar, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(sb.Id(), co.NM_DBLCLK, func(p wm.Notify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) SbnRClick(sb *StatusBar, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(sb.Id(), co.NM_RCLICK, func(p wm.Notify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) SbnRDblClk(sb *StatusBar, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(sb.Id(), co.NM_RDBLCLK, func(p wm.Notify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- TreeView TVN ---

func (me *windowDepotNfy) TvnAsyncDraw(tv *TreeView, userFunc func(p *win.NMTVASYNCDRAW)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_ASYNCDRAW), func(p wm.Notify) uintptr {
		userFunc((*win.NMTVASYNCDRAW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnBeginDrag(tv *TreeView, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_BEGINDRAG), func(p wm.Notify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnBeginLabelEdit(tv *TreeView, userFunc func(p *win.NMTVDISPINFO) bool) {
	me.addNfy(tv.Id(), co.NM(co.TVN_BEGINLABELEDIT), func(p wm.Notify) uintptr {
		if userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnBeginRDrag(tv *TreeView, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_BEGINRDRAG), func(p wm.Notify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnDeleteItem(tv *TreeView, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_DELETEITEM), func(p wm.Notify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnEndLabelEdit(tv *TreeView, userFunc func(p *win.NMTVDISPINFO) bool) {
	me.addNfy(tv.Id(), co.NM(co.TVN_ENDLABELEDIT), func(p wm.Notify) uintptr {
		if userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnGetDispInfo(tv *TreeView, userFunc func(p *win.NMTVDISPINFO)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_GETDISPINFO), func(p wm.Notify) uintptr {
		userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnGetInfoTip(tv *TreeView, userFunc func(p *win.NMTVGETINFOTIP)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_GETINFOTIP), func(p wm.Notify) uintptr {
		userFunc((*win.NMTVGETINFOTIP)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnItemChanged(tv *TreeView, userFunc func(p *win.NMTVITEMCHANGE)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_ITEMCHANGED), func(p wm.Notify) uintptr {
		userFunc((*win.NMTVITEMCHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnItemChanging(tv *TreeView, userFunc func(p *win.NMTVITEMCHANGE) bool) {
	me.addNfy(tv.Id(), co.NM(co.TVN_ITEMCHANGING), func(p wm.Notify) uintptr {
		if userFunc((*win.NMTVITEMCHANGE)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnItemExpanded(tv *TreeView, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_ITEMEXPANDED), func(p wm.Notify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnItemExpanding(tv *TreeView, userFunc func(p *win.NMTREEVIEW) bool) {
	me.addNfy(tv.Id(), co.NM(co.TVN_ITEMEXPANDING), func(p wm.Notify) uintptr {
		if userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnKeyDown(tv *TreeView, userFunc func(p *win.NMTVKEYDOWN) int32) {
	me.addNfy(tv.Id(), co.NM(co.TVN_KEYDOWN), func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMTVKEYDOWN)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) TvnSelChanged(tv *TreeView, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_SELCHANGED), func(p wm.Notify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnSelChanging(tv *TreeView, userFunc func(p *win.NMTREEVIEW) bool) {
	me.addNfy(tv.Id(), co.NM(co.TVN_SELCHANGING), func(p wm.Notify) uintptr {
		if userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnSetDispInfo(tv *TreeView, userFunc func(p *win.NMTVDISPINFO)) {
	me.addNfy(tv.Id(), co.NM(co.TVN_SETDISPINFO), func(p wm.Notify) uintptr {
		userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnSingleExpand(tv *TreeView, userFunc func(p *win.NMTREEVIEW) co.TVNRET) {
	me.addNfy(tv.Id(), co.NM(co.TVN_SINGLEEXPAND), func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))))
	})
}

//--------------------------------------------------------------- TreView NM ---

func (me *windowDepotNfy) TvnClick(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_CLICK, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnCustomDraw(tv *TreeView, userFunc func(p *win.NMTVCUSTOMDRAW) co.CDRF) {
	me.addNfy(tv.Id(), co.NM_CUSTOMDRAW, func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMTVCUSTOMDRAW)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) TvnDblClk(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_DBLCLK, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnKillFocus(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_KILLFOCUS, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnRClick(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_RCLICK, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnRDblClk(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_RDBLCLK, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnReturn(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_RETURN, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnSetCursor(tv *TreeView, userFunc func(p *win.NMMOUSE) int32) {
	me.addNfy(tv.Id(), co.NM_SETCURSOR, func(p wm.Notify) uintptr {
		return uintptr(userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) TvnSetFocus(tv *TreeView, userFunc func(p *win.NMHDR)) {
	me.addNfy(tv.Id(), co.NM_SETFOCUS, func(p wm.Notify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}
