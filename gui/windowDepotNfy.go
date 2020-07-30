/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

type nfyHash struct { // custom hash for WM_NOTIFY messages
	IdFrom int32
	Code   co.NM
}

// Keeps all user common control notification handlers.
type windowDepotNfy struct {
	mapNfys map[nfyHash]func(p WmNotify) uintptr
}

func (me *windowDepotNfy) addNfy(idFrom int32, code co.NM,
	userFunc func(p WmNotify) uintptr) {

	if me.mapNfys == nil { // guard
		me.mapNfys = make(map[nfyHash]func(p WmNotify) uintptr, 16) // arbitrary capacity
	}
	me.mapNfys[nfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *windowDepotNfy) processMessage(msg co.WM, p Wm) (uintptr, bool) {
	if msg == co.WM_NOTIFY {
		pNfy := WmNotify(p)
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

func (me *windowDepotNfy) LvnBeginDrag(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginLabelEdit(listViewId int32, userFunc func(p *win.NMLVDISPINFO) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginRDrag(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINRDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnBeginScroll(listViewId int32, userFunc func(p *win.NMLVSCROLL)) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINSCROLL), func(p WmNotify) uintptr {
		userFunc((*win.NMLVSCROLL)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnClick(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_COLUMNCLICK), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnDropDown(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_COLUMNDROPDOWN), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnColumnOverflowClick(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_COLUMNOVERFLOWCLICK), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnDeleteAllItems(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_DELETEALLITEMS), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnDeleteItem(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnEndLabelEdit(listViewId int32, userFunc func(p *win.NMLVDISPINFO) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_ENDLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnEndScroll(listViewId int32, userFunc func(p *win.NMLVSCROLL)) {
	me.addNfy(listViewId, co.NM(co.LVN_ENDSCROLL), func(p WmNotify) uintptr {
		userFunc((*win.NMLVSCROLL)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnGetDispInfo(listViewId int32, userFunc func(p *win.NMLVDISPINFO)) {
	me.addNfy(listViewId, co.NM(co.LVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnGetEmptyMarkup(listViewId int32, userFunc func(p *win.NMLVEMPTYMARKUP) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_GETEMPTYMARKUP), func(p WmNotify) uintptr {
		if userFunc((*win.NMLVEMPTYMARKUP)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnGetInfoTip(listViewId int32, userFunc func(p *win.NMLVGETINFOTIP)) {
	me.addNfy(listViewId, co.NM(co.LVN_GETINFOTIP), func(p WmNotify) uintptr {
		userFunc((*win.NMLVGETINFOTIP)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnHotTrack(listViewId int32, userFunc func(p *win.NMLISTVIEW) int32) {
	me.addNfy(listViewId, co.NM(co.LVN_HOTTRACK), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnIncrementalSearch(listViewId int32, userFunc func(p *win.NMLVFINDITEM) int32) {
	me.addNfy(listViewId, co.NM(co.LVN_INCREMENTALSEARCH), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnInsertItem(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_INSERTITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemActivate(listViewId int32, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM(co.LVN_ITEMACTIVATE), func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemChanged(listViewId int32, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_ITEMCHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnItemChanging(listViewId int32, userFunc func(p *win.NMLISTVIEW) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_ITEMCHANGING), func(p WmNotify) uintptr {
		if userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) LvnKeyDown(listViewId int32, userFunc func(p *win.NMLVKEYDOWN)) {
	me.addNfy(listViewId, co.NM(co.LVN_KEYDOWN), func(p WmNotify) uintptr {
		userFunc((*win.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnLinkClick(listViewId int32, userFunc func(p *win.NMLVLINK)) {
	me.addNfy(listViewId, co.NM(co.LVN_LINKCLICK), func(p WmNotify) uintptr {
		userFunc((*win.NMLVLINK)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnMarqueeBegin(listViewId int32, userFunc func(p *win.NMHDR) uint32) {
	me.addNfy(listViewId, co.NM(co.LVN_MARQUEEBEGIN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnODCacheHint(listViewId int32, userFunc func(p *win.NMLVCACHEHINT)) {
	me.addNfy(listViewId, co.NM(co.LVN_ODCACHEHINT), func(p WmNotify) uintptr {
		userFunc((*win.NMLVCACHEHINT)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnODFindItem(listViewId int32, userFunc func(p *win.NMLVFINDITEM) int32) {
	me.addNfy(listViewId, co.NM(co.LVN_ODFINDITEM), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnODStateChanged(listViewId int32, userFunc func(p *win.NMLVODSTATECHANGE)) {
	me.addNfy(listViewId, co.NM(co.LVN_ODSTATECHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMLVODSTATECHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnSetDispInfo(listViewId int32, userFunc func(p *win.NMLVDISPINFO)) {
	me.addNfy(listViewId, co.NM(co.LVN_SETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//-------------------------------------------------------------- ListView NM ---

func (me *windowDepotNfy) LvnClick(listViewId int32, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnCustomDraw(listViewId int32, userFunc func(p *win.NMLVCUSTOMDRAW) co.CDRF) {
	me.addNfy(listViewId, co.NM_CUSTOMDRAW, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLVCUSTOMDRAW)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnDblClk(listViewId int32, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnHover(listViewId int32, userFunc func(p *win.NMHDR) uint32) {
	me.addNfy(listViewId, co.NM_HOVER, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) LvnKillFocus(listViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnRClick(listViewId int32, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnRDblClk(listViewId int32, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnReleasedCapture(listViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_RELEASEDCAPTURE, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnReturn(listViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_RETURN, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) LvnSetFocus(listViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------ StatusBar SBN ---

func (me *windowDepotNfy) SbnSimpleModeChange(statusBarId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(statusBarId, co.NM(co.SBN_SIMPLEMODECHANGE), func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- StatusBar NM ---

func (me *windowDepotNfy) SbnClick(statusBarId int32, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) SbnDblClk(statusBarId int32, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) SbnRClick(statusBarId int32, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) SbnRDblClk(statusBarId int32, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- TreeView TVN ---

func (me *windowDepotNfy) TvnAsyncDraw(treeViewId int32, userFunc func(p *win.NMTVASYNCDRAW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_ASYNCDRAW), func(p WmNotify) uintptr {
		userFunc((*win.NMTVASYNCDRAW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnBeginDrag(treeViewId int32, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_BEGINDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnBeginLabelEdit(treeViewId int32, userFunc func(p *win.NMTVDISPINFO) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_BEGINLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnBeginRDrag(treeViewId int32, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_BEGINRDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnDeleteItem(treeViewId int32, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnEndLabelEdit(treeViewId int32, userFunc func(p *win.NMTVDISPINFO) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_ENDLABELEDIT), func(p WmNotify) uintptr {
		if userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnGetDispInfo(treeViewId int32, userFunc func(p *win.NMTVDISPINFO)) {
	me.addNfy(treeViewId, co.NM(co.TVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnGetInfoTip(treeViewId int32, userFunc func(p *win.NMTVGETINFOTIP)) {
	me.addNfy(treeViewId, co.NM(co.TVN_GETINFOTIP), func(p WmNotify) uintptr {
		userFunc((*win.NMTVGETINFOTIP)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnItemChanged(treeViewId int32, userFunc func(p *win.NMTVITEMCHANGE)) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMCHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMTVITEMCHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnItemChanging(treeViewId int32, userFunc func(p *win.NMTVITEMCHANGE) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMCHANGING), func(p WmNotify) uintptr {
		if userFunc((*win.NMTVITEMCHANGE)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnItemExpanded(treeViewId int32, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMEXPANDED), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnItemExpanding(treeViewId int32, userFunc func(p *win.NMTREEVIEW) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMEXPANDING), func(p WmNotify) uintptr {
		if userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnKeyDown(treeViewId int32, userFunc func(p *win.NMTVKEYDOWN) int32) {
	me.addNfy(treeViewId, co.NM(co.TVN_KEYDOWN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMTVKEYDOWN)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) TvnSelChanged(treeViewId int32, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_SELCHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnSelChanging(treeViewId int32, userFunc func(p *win.NMTREEVIEW) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_SELCHANGING), func(p WmNotify) uintptr {
		if userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

func (me *windowDepotNfy) TvnSetDispInfo(treeViewId int32, userFunc func(p *win.NMTVDISPINFO)) {
	me.addNfy(treeViewId, co.NM(co.TVN_SETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnSingleExpand(treeViewId int32, userFunc func(p *win.NMTREEVIEW) co.TVNRET) {
	me.addNfy(treeViewId, co.NM(co.TVN_SINGLEEXPAND), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))))
	})
}

//--------------------------------------------------------------- TreView NM ---

func (me *windowDepotNfy) TvnClick(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnCustomDraw(treeViewId int32, userFunc func(p *win.NMTVCUSTOMDRAW) co.CDRF) {
	me.addNfy(treeViewId, co.NM_CUSTOMDRAW, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMTVCUSTOMDRAW)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) TvnDblClk(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnKillFocus(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnRClick(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnRDblClk(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnReturn(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_RETURN, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

func (me *windowDepotNfy) TvnSetCursor(treeViewId int32, userFunc func(p *win.NMMOUSE) int32) {
	me.addNfy(treeViewId, co.NM_SETCURSOR, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam))))
	})
}

func (me *windowDepotNfy) TvnSetFocus(treeViewId int32, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}
