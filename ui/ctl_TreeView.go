//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native [tree view] control.
//
// [tree view]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-controls
type TreeView struct {
	_BaseCtrl
	events    EventsTreeView
	itemsData map[win.HTREEITEM]interface{} // data associated with each item; replaces LPARAM approach
	Items     CollectionTreeViewItems       // Methods to interact with the items collection.
}

// Creates a new [TreeView] with [win.CreateWindowEx].
func NewTreeView(parent Parent, opts *VarOptsTreeView) *TreeView {
	setUniqueCtrlId(&opts.ctrlId)
	me := &TreeView{
		_BaseCtrl: newBaseCtrl(opts.ctrlId),
		events:    EventsTreeView{opts.ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

	parent.base().beforeUserEvents.WmCreate(func(_ WmCreate) int {
		me.createWindow(opts.wndExStyle, "SysTreeView32", "",
			opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, false)
		if opts.ctrlExStyle != co.TVS_EX(0) {
			me.SetExtendedStyle(true, opts.ctrlExStyle)
		}
		parent.base().layout.Add(parent, me.hWnd, opts.layout)
		return 0 // ignored
	})

	me.defaultMessageHandlers(parent)
	return me
}

// Instantiates a new [TreeView] to be loaded from a dialog resource with
// [win.HWND.GetDlgItem].
func NewTreeViewDlg(parent Parent, ctrlId uint16, layout LAY) *TreeView {
	me := &TreeView{
		_BaseCtrl: newBaseCtrl(ctrlId),
		events:    EventsTreeView{ctrlId, &parent.base().userEvents},
	}
	me.Items.owner = me

	parent.base().beforeUserEvents.WmInitDialog(func(_ WmInitDialog) bool {
		me.assignDialog(parent)
		parent.base().layout.Add(parent, me.hWnd, layout)
		return true // ignored
	})

	me.defaultMessageHandlers(parent)
	return me
}

func (me *TreeView) defaultMessageHandlers(parent Parent) {
	parent.base().afterUserEvents.WmNotify(me.ctrlId, co.TVN_DELETEITEM, func(p unsafe.Pointer) uintptr {
		nmtv := (*win.NMTREEVIEW)(p)
		delete(me.itemsData, nmtv.ItemOld.HItem)
		return 0 // ignored
	})

	parent.base().afterUserEvents.WmDestroy(func() {
		kinds := []co.TVSIL{co.TVSIL_NORMAL, co.TVSIL_STATE}
		for _, kind := range kinds {
			h, _ := me.hWnd.SendMessage(co.TVM_GETIMAGELIST, win.WPARAM(kind), 0)
			if h != 0 {
				me.hWnd.SendMessage(co.TVM_SETIMAGELIST, win.WPARAM(kind), 0)
				win.HIMAGELIST(h).Destroy()
			}
		}
	})
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *TreeView) On() *EventsTreeView {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Retrieves the given image list with [TVM_GETIMAGELIST]. The image lists are
// lazy-initialized: the first time you call this method for a given image list,
// it will be created and assigned with [TVM_SETIMAGELIST].
//
// The image lists will be automatically destroyed.
//
// [TVM_GETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getimagelist
// [TVM_SETIMAGELIST]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-setimagelist
func (me *TreeView) ImageList(which co.TVSIL) win.HIMAGELIST {
	h, _ := me.hWnd.SendMessage(co.TVM_GETIMAGELIST, win.WPARAM(which), 0)
	hImg := win.HIMAGELIST(h)
	if hImg == win.HIMAGELIST(0) {
		hImg, _ = win.ImageListCreate(16, 16, co.ILC_COLOR32, 1, 1)
		me.hWnd.SendMessage(co.TVM_SETIMAGELIST, win.WPARAM(which), win.LPARAM(hImg))
	}
	return hImg
}

// Adds or removes extended styles with [TVM_SETEXTENDEDSTYLE].
//
// Returns the same object, so further operations can be chained.
//
// [TVM_SETEXTENDEDSTYLE]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-setextendedstyle
func (me *TreeView) SetExtendedStyle(doSet bool, style co.TVS_EX) *TreeView {
	affected := co.TVS_EX(0)
	if doSet {
		affected = style
	}
	me.hWnd.SendMessage(co.TVM_SETEXTENDEDSTYLE,
		win.WPARAM(affected), win.LPARAM(style))
	return me
}

// Options for [NewTreeView]; returned by [OptsTreeView].
type VarOptsTreeView struct {
	ctrlId      uint16
	layout      LAY
	position    win.POINT
	size        win.SIZE
	ctrlStyle   co.TVS
	ctrlExStyle co.TVS_EX
	wndStyle    co.WS
	wndExStyle  co.WS_EX
}

// Options for [NewTreeView].
func OptsTreeView() *VarOptsTreeView {
	return &VarOptsTreeView{
		size:       win.SIZE{Cx: int32(DpiX(120)), Cy: int32(DpiY(120))},
		ctrlStyle:  co.TVS_HASLINES | co.TVS_LINESATROOT | co.TVS_SHOWSELALWAYS | co.TVS_HASBUTTONS,
		wndStyle:   co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE,
		wndExStyle: co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE,
	}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsTreeView) CtrlId(id uint16) *VarOptsTreeView { o.ctrlId = id; return o }

// Horizontal and vertical behavior for the control layout, when the parent
// window is resized.
//
// Defaults to ui.LAY_NONE_NONE.
func (o *VarOptsTreeView) Layout(l LAY) *VarOptsTreeView { o.layout = l; return o }

// Position coordinates within parent window client area, in pixels, passed to
// [win.CreateWindowEx].
//
// Defaults to ui.Dpi(0, 0).
func (o *VarOptsTreeView) Position(x, y int) *VarOptsTreeView {
	o.position.X = int32(x)
	o.position.Y = int32(y)
	return o
}

// Control size in pixels, passed to [win.CreateWindowEx].
//
// Defaults to ui.Dpi(120, 120).
func (o *VarOptsTreeView) Size(cx int, cy int) *VarOptsTreeView {
	o.size.Cx = int32(cx)
	o.size.Cy = int32(cy)
	return o
}

// Tree view control [style], passed to [win.CreateWindowEx].
//
// Defaults to co.TVS_HASLINES | co.TVS_LINESATROOT | co.TVS_SHOWSELALWAYS | co.TVS_HASBUTTONS.
//
// [style]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-control-window-styles
func (o *VarOptsTreeView) CtrlStyle(s co.TVS) *VarOptsTreeView { o.ctrlStyle = s; return o }

// Tree view control [extended style].
//
// Defaults to co.TVS_EX_NONE.
//
// [extended style]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-control-window-extended-styles
func (o *VarOptsTreeView) CtrlExStyle(s co.TVS_EX) *VarOptsTreeView { o.ctrlExStyle = s; return o }

// Window style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_CHILD | co.WS_GROUP | co.WS_TABSTOP | co.WS_VISIBLE.
func (o *VarOptsTreeView) WndStyle(s co.WS) *VarOptsTreeView { o.wndStyle = s; return o }

// Window extended style, passed to [win.CreateWindowEx].
//
// Defaults to co.WS_EX_LEFT | co.WS_EX_CLIENTEDGE.
func (o *VarOptsTreeView) WndExStyle(s co.WS_EX) *VarOptsTreeView { o.wndExStyle = s; return o }

// Native [tree view] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [tree view]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-controls
type EventsTreeView struct {
	ctrlId       uint16
	parentEvents *EventsWindow
}

// [TVN_ASYNCDRAW] message handler.
//
// [TVN_ASYNCDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-asyncdraw
func (me *EventsTreeView) TvnAsyncDraw(fun func(p *win.NMTVASYNCDRAW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_ASYNCDRAW, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTVASYNCDRAW)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_BEGINDRAG] message handler.
//
// [TVN_BEGINDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-begindrag
func (me *EventsTreeView) TvnBeginDrag(fun func(p *win.NMTREEVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_BEGINDRAG, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTREEVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_BEGINLABELEDIT] message handler.
//
// [TVN_BEGINLABELEDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-beginlabeledit
func (me *EventsTreeView) TvnBeginLabelEdit(fun func(p *win.NMTVDISPINFO) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_BEGINLABELEDIT, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTVDISPINFO)(p)))
	})
}

// [TVN_BEGINRDRAG] message handler.
//
// [TVN_BEGINRDRAG]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-beginrdrag
func (me *EventsTreeView) TvnBeginRDrag(fun func(p *win.NMTREEVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_BEGINRDRAG, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTREEVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_DELETEITEM] message handler.
//
// [TVN_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-deleteitem
func (me *EventsTreeView) TvnDeleteItem(fun func(p *win.NMTREEVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_DELETEITEM, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTREEVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_ENDLABELEDIT] message handler.
//
// [TVN_ENDLABELEDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-endlabeledit
func (me *EventsTreeView) TvnEndLabelEdit(fun func(p *win.NMTVDISPINFO) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_ENDLABELEDIT, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTVDISPINFO)(p)))
	})
}

// [TVN_GETDISPINFO] message handler.
//
// [TVN_GETDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-getdispinfo
func (me *EventsTreeView) TvnGetDispInfo(fun func(p *win.NMTVDISPINFO)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_GETDISPINFO, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTVDISPINFO)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_GETINFOTIP] message handler.
//
// [TVN_GETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-getinfotip
func (me *EventsTreeView) TvnGetInfoTip(fun func(p *win.NMTVGETINFOTIP)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_GETINFOTIP, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTVGETINFOTIP)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_ITEMCHANGED] message handler.
//
// [TVN_ITEMCHANGED]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-itemchanged
func (me *EventsTreeView) TvnItemChanged(fun func(p *win.NMTVITEMCHANGE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_ITEMCHANGED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTVITEMCHANGE)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_ITEMCHANGING] message handler.
//
// [TVN_ITEMCHANGING]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-itemchanging
func (me *EventsTreeView) TvnItemChanging(fun func(p *win.NMTVITEMCHANGE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_ITEMCHANGING, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTVITEMCHANGE)(p)))
	})
}

// [TVN_ITEMEXPANDED] message handler.
//
// [TVN_ITEMEXPANDED]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanded
func (me *EventsTreeView) TvnItemExpanded(fun func(p *win.NMTREEVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_ITEMEXPANDED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTREEVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_ITEMEXPANDING] message handler.
//
// [TVN_ITEMEXPANDING]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanding
func (me *EventsTreeView) TvnItemExpanding(fun func(p *win.NMTREEVIEW) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_ITEMEXPANDING, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTREEVIEW)(p)))
	})
}

// [TVN_KEYDOWN] message handler.
//
// [TVN_KEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-keydown
func (me *EventsTreeView) TvnKeyDown(fun func(p *win.NMTVKEYDOWN) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_KEYDOWN, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTVKEYDOWN)(p)))
	})
}

// [TVN_SELCHANGED] message handler.
//
// [TVN_SELCHANGED]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-selchanged
func (me *EventsTreeView) TvnSelChanged(fun func(p *win.NMTREEVIEW)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_SELCHANGED, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTREEVIEW)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_SELCHANGING] message handler.
//
// [TVN_SELCHANGING]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-selchanging
func (me *EventsTreeView) TvnSelChanging(fun func(p *win.NMTREEVIEW) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_SELCHANGING, func(p unsafe.Pointer) uintptr {
		return wutil.BoolToUintptr(fun((*win.NMTREEVIEW)(p)))
	})
}

// [TVN_SETDISPINFO] message handler.
//
// [TVN_SETDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-setdispinfo
func (me *EventsTreeView) TvnSetDispInfo(fun func(p *win.NMTVDISPINFO)) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_SETDISPINFO, func(p unsafe.Pointer) uintptr {
		fun((*win.NMTVDISPINFO)(p))
		return me.parentEvents.defProcVal
	})
}

// [TVN_SINGLEEXPAND] message handler.
//
// [TVN_SINGLEEXPAND]: https://learn.microsoft.com/en-us/windows/win32/controls/tvn-singleexpand
func (me *EventsTreeView) TvnSingleExpand(fun func(p *win.NMTREEVIEW) co.TVNRET) {
	me.parentEvents.WmNotify(me.ctrlId, co.TVN_SINGLEEXPAND, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTREEVIEW)(p)))
	})
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-tree-view
func (me *EventsTreeView) NmClick(fun func() int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(_ unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [NM_CUSTOMDRAW] message handler.
//
// [NM_CUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-customdraw-tree-view
func (me *EventsTreeView) NmCustomDraw(fun func(p *win.NMTVCUSTOMDRAW) co.CDRF) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMTVCUSTOMDRAW)(p)))
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-dblclk-tree-view
func (me *EventsTreeView) NmDblClk(fun func() int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_DBLCLK, func(_ unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [NM_KILLFOCUS] message handler.
//
// [NM_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-killfocus-tree-view
func (me *EventsTreeView) NmKillFocus(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_KILLFOCUS, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-tree-view
func (me *EventsTreeView) NmRClick(fun func() int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(_ unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tree-view
func (me *EventsTreeView) NmRDblClk(fun func() int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RDBLCLK, func(_ unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [NM_RETURN] message handler.
//
// [NM_RETURN]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-return-tree-view-
func (me *EventsTreeView) NmReturn(fun func() int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RETURN, func(_ unsafe.Pointer) uintptr {
		return uintptr(fun())
	})
}

// [NM_SETCURSOR] message handler.
//
// [NM_SETCURSOR]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-setcursor-tree-view-
func (me *EventsTreeView) NmSetCursor(fun func(p *win.NMMOUSE) int) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_SETCURSOR, func(p unsafe.Pointer) uintptr {
		return uintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_SETFOCUS] message handler.
//
// [NM_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-setfocus-tree-view-
func (me *EventsTreeView) NmSetFocus(fun func()) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_SETFOCUS, func(_ unsafe.Pointer) uintptr {
		fun()
		return me.parentEvents.defProcVal
	})
}
