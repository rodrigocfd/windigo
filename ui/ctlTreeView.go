/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native tree view control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/tree-view-controls
type TreeView struct {
	*_NativeControlBase
	events *_EventsTreeView
	items  *_TreeViewItemCollection
}

// Constructor. Optionally receives a control ID.
func NewTreeView(parent Parent, ctrlId ...int) *TreeView {
	base := _NewNativeControlBase(parent, ctrlId...)
	me := &TreeView{
		_NativeControlBase: base,
		events:             _NewEventsTreeView(base),
	}
	me.items = _NewTreeViewItemCollection(me)
	return me
}

// Calls CreateWindowEx(). With this method, you must also specify WS and WS_EX
// window styles.
//
// Position and size will be adjusted to the current system DPI.
func (me *TreeView) CreateWs(
	pos Pos, size Size,
	tvStyles co.TVS, tvExStyles co.TVS_EX,
	styles co.WS, exStyles co.WS_EX) *TreeView {

	_global.MultiplyDpi(&pos, &size)
	me._NativeControlBase.create("SysTreeView32", "", pos, size,
		co.WS(tvStyles)|styles, exStyles)

	if tvExStyles != co.TVS_EX_NONE {
		me.SetExtendedStyle(true, tvExStyles)
	}
	return me
}

// Calls CreateWindowEx() with WS_CHILD | WS_GROUP | WS_TABSTOP | WS_VISIBLE,
// and WS_EX_CLIENTEDGE.
//
// A typical TreeView has TVS_HASLINES | TVS_LINESATROOT | TVS_SHOWSELALWAYS | TVS_HASBUTTONS.
//
// Position and size will be adjusted to the current system DPI.
func (me *TreeView) Create(
	pos Pos, size Size, tvStyles co.TVS, tvExStyles co.TVS_EX) *TreeView {

	return me.CreateWs(pos, size, tvStyles, tvExStyles,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.WS_EX_CLIENTEDGE)
}

// Exposes all ListView notifications.
func (me *TreeView) On() *_EventsTreeView {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the TreeView was created.")
	}
	return me.events
}

// Retrieves extended styles.
func (me *TreeView) ExtendedStyle() co.TVS_EX {
	return co.TVS_EX(
		me.Hwnd().SendMessage(co.WM(co.TVM_GETEXTENDEDSTYLE), 0, 0),
	)
}

// Access to the items.
func (me *TreeView) Items() *_TreeViewItemCollection {
	return me.items
}

// Sets or unsets extended styles.
func (me *TreeView) SetExtendedStyle(isSet bool, exStyle co.TVS_EX) *TreeView {
	mask := exStyle
	if !isSet {
		mask = 0
	}
	me.Hwnd().SendMessage(co.WM(co.TVM_SETEXTENDEDSTYLE),
		win.WPARAM(mask), win.LPARAM(exStyle))
	return me
}

// Sends WM_SETREDRAW to enable or disable UI updates.
//
// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-setredraw
func (me *TreeView) SetRedraw(allowRedraw bool) *TreeView {
	me.Hwnd().SendMessage(co.WM_SETREDRAW,
		win.WPARAM(_global.BoolToUint32(allowRedraw)), 0)
	return me
}

//------------------------------------------------------------------------------

type _TreeViewItemCollection struct {
	ctrl *TreeView
}

// Constructor.
func _NewTreeViewItemCollection(ctrl *TreeView) *_TreeViewItemCollection {
	return &_TreeViewItemCollection{
		ctrl: ctrl,
	}
}

// Adds a new root item, returning it.
func (me *_TreeViewItemCollection) AddRoot(text string) *TreeViewItem {
	textBuf := win.Str.ToUint16Slice(text)

	tvi := win.TVINSERTSTRUCT{
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: &textBuf[0],
		},
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_INSERTITEM),
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return me.Get(win.HTREEITEM(ret))
}

// Retrieves the number of items.
func (me *_TreeViewItemCollection) Count() int {
	return int(
		me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_GETCOUNT), 0, 0),
	)
}

// Deletes all items at once.
func (me *_TreeViewItemCollection) DeleteAll() *TreeView {
	if me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_DELETEITEM), 0, 0) == 0 {
		panic("TVM_DELETEITEM failed.")
	}
	return me.ctrl
}

// Retrieves the item with the given handle.
func (me *_TreeViewItemCollection) Get(hTreeItem win.HTREEITEM) *TreeViewItem {
	return _NewTreeViewItem(me.ctrl, hTreeItem)
}

// Retrieves the first root item, or nil if none.
func (me *_TreeViewItemCollection) FirstRoot() *TreeViewItem {
	return me.Get(0).nextItem(co.TVGN_ROOT)
}

// Retrieves the first visible item, or nil if none.
func (me *_TreeViewItemCollection) FirstVisible() *TreeViewItem {
	return me.Get(0).nextItem(co.TVGN_FIRSTVISIBLE)
}

// Retrieves all root items, if any.
func (me *_TreeViewItemCollection) Roots() []*TreeViewItem {
	return me.Get(0).Children()
}

// Retrieves the currently selected item, or nil if none.
func (me *_TreeViewItemCollection) Selected() *TreeViewItem {
	return me.Get(0).nextItem(co.TVGN_CARET)
}

//------------------------------------------------------------------------------

// A single item of a tree view control.
type TreeViewItem struct {
	ctrl      *TreeView
	hTreeItem win.HTREEITEM
}

// Constructor.
func _NewTreeViewItem(ctrl *TreeView, hTreeItem win.HTREEITEM) *TreeViewItem {
	return &TreeViewItem{
		ctrl:      ctrl,
		hTreeItem: hTreeItem,
	}
}

// Adds a new child item, returning it.
func (me *TreeViewItem) AddChild(text string) *TreeViewItem {
	textBuf := win.Str.ToUint16Slice(text)

	tvi := win.TVINSERTSTRUCT{
		HParent:      me.hTreeItem,
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: &textBuf[0],
		},
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_INSERTITEM),
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return me.ctrl.Items().Get(win.HTREEITEM(ret))
}

// Retrieves all child items, if any.
func (me *TreeViewItem) Children() []*TreeViewItem {
	childNodes := make([]*TreeViewItem, 0)
	node := me.FirstChild()
	for node != nil {
		childNodes = append(childNodes, node)
	}
	return childNodes
}

// Deletes this item and all its children.
func (me *TreeViewItem) Delete() {
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_DELETEITEM),
		0, win.LPARAM(me.hTreeItem))
	if ret == 0 {
		panic("TVM_DELETEITEM failed.")
	}
}

// Expand nodes and scrolls the TreeView so this item becomes visible.
func (me *TreeViewItem) EnsureVisible() *TreeViewItem {
	me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_ENSUREVISIBLE),
		0, win.LPARAM(me.hTreeItem))
	return me
}

// Expands or collapses the item.
func (me *TreeViewItem) Expand(isExpanded bool) *TreeViewItem {
	flag := co.TVE_EXPAND
	if !isExpanded {
		flag = co.TVE_COLLAPSE
	}
	me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_EXPAND),
		win.WPARAM(flag), win.LPARAM(me.hTreeItem))
	return me
}

// Retrieves the first child item, or nil if none.
func (me *TreeViewItem) FirstChild() *TreeViewItem {
	return me.nextItem(co.TVGN_CHILD)
}

// Returns the underlying HTREEITEM handle of this item.
func (me *TreeViewItem) HTreeItem() win.HTREEITEM {
	return me.hTreeItem
}

// Tells if the item is currently expanded.
func (me *TreeViewItem) IsExpanded() bool {
	return (co.TVIS(
		me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_GETITEMSTATE),
			win.WPARAM(me.hTreeItem), win.LPARAM(co.TVIS_EXPANDED)),
	) & co.TVIS_EXPANDED) != 0
}

// Tells if the node is a root node.
func (me *TreeViewItem) IsRoot() bool {
	return me.Parent().hTreeItem == 0
}

// Retrieves the next sibling, or nil if none.
func (me *TreeViewItem) NextSibling() *TreeViewItem {
	return me.nextItem(co.TVGN_NEXT)
}

// Retrieves the associated LPARAM with TVM_GETITEM.
func (me *TreeViewItem) Param() win.LPARAM {
	tvi := win.TVITEMEX{
		HItem: me.hTreeItem,
		Mask:  co.TVIF_PARAM,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_GETITEM),
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return tvi.LParam
}

// Retrieves the parent item, or nil if none.
func (me *TreeViewItem) Parent() *TreeViewItem {
	return me.nextItem(co.TVGN_PARENT)
}

// Retrieves the previous sibling, or nil if none.
func (me *TreeViewItem) PrevSibling() *TreeViewItem {
	return me.nextItem(co.TVGN_PREVIOUS)
}

// Sets the associated LPARAM.
func (me *TreeViewItem) SetParam(lp win.LPARAM) *TreeViewItem {
	tvi := win.TVITEMEX{
		HItem:  me.hTreeItem,
		Mask:   co.TVIF_PARAM,
		LParam: lp,
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_SETITEM),
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_SETITEM failed.")
	}
	return me
}

// Sets the text.
func (me *TreeViewItem) SetText(text string) *TreeViewItem {
	textBuf := win.Str.ToUint16Slice(text)
	tvi := win.TVITEMEX{
		HItem:   me.hTreeItem,
		Mask:    co.TVIF_TEXT,
		PszText: &textBuf[0],
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_SETITEM),
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
	return me
}

// Retrieves the text.
func (me *TreeViewItem) Text() string {
	buf := [256]uint16{} // arbitrary
	tvi := win.TVITEMEX{
		HItem:      me.hTreeItem,
		Mask:       co.TVIF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_GETITEM),
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return syscall.UTF16ToString(buf[:])
}

// Toggles the node, expanded or collapsed.
func (me *TreeViewItem) ToggleExpand() *TreeViewItem {
	me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_EXPAND),
		win.WPARAM(co.TVE_TOGGLE), win.LPARAM(me.hTreeItem))
	return me
}

// Sends TVM_GETNEXTITEM, returns nil if none found.
func (me *TreeViewItem) nextItem(flags co.TVGN) *TreeViewItem {
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.TVM_GETNEXTITEM),
		win.WPARAM(flags), win.LPARAM(me.hTreeItem)) // HTREEITEM can be zero
	if ret == 0 {
		return nil
	}
	return me.ctrl.Items().Get(win.HTREEITEM(ret))
}

//------------------------------------------------------------------------------

// TreeView control notifications.
type _EventsTreeView struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsTreeView(ctrl *_NativeControlBase) *_EventsTreeView {
	return &_EventsTreeView{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-asyncdraw
func (me *_EventsTreeView) TvnAsyncDraw(userFunc func(p *win.NMTVASYNCDRAW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_ASYNCDRAW), func(p unsafe.Pointer) {
		userFunc((*win.NMTVASYNCDRAW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-begindrag
func (me *_EventsTreeView) TvnBeginDrag(userFunc func(p *win.NMTREEVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_BEGINDRAG), func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-beginlabeledit
func (me *_EventsTreeView) TvnBeginLabelEdit(userFunc func(p *win.NMTVDISPINFO) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_BEGINLABELEDIT), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMTVDISPINFO)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-beginrdrag
func (me *_EventsTreeView) TvnBeginRDrag(userFunc func(p *win.NMTREEVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_BEGINRDRAG), func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-deleteitem
func (me *_EventsTreeView) TvnDeleteItem(userFunc func(p *win.NMTREEVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_DELETEITEM), func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-endlabeledit
func (me *_EventsTreeView) TvnEndLabelEdit(userFunc func(p *win.NMTVDISPINFO) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_ENDLABELEDIT), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMTVDISPINFO)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-getdispinfo
func (me *_EventsTreeView) TvnGetDispInfo(userFunc func(p *win.NMTVDISPINFO)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_GETDISPINFO), func(p unsafe.Pointer) {
		userFunc((*win.NMTVDISPINFO)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-getinfotip
func (me *_EventsTreeView) TvnGetInfoTip(userFunc func(p *win.NMTVGETINFOTIP)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_GETINFOTIP), func(p unsafe.Pointer) {
		userFunc((*win.NMTVGETINFOTIP)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemchanged
func (me *_EventsTreeView) TvnItemChanged(userFunc func(p *win.NMTVITEMCHANGE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_ITEMCHANGED), func(p unsafe.Pointer) {
		userFunc((*win.NMTVITEMCHANGE)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemchanging
func (me *_EventsTreeView) TvnItemChanging(userFunc func(p *win.NMTVITEMCHANGE) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_ITEMCHANGING), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMTVITEMCHANGE)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanded
func (me *_EventsTreeView) TvnItemExpanded(userFunc func(p *win.NMTREEVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_ITEMEXPANDED), func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanding
func (me *_EventsTreeView) TvnItemExpanding(userFunc func(p *win.NMTREEVIEW) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_ITEMEXPANDING), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMTREEVIEW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-keydown
func (me *_EventsTreeView) TvnKeyDown(userFunc func(p *win.NMTVKEYDOWN) int) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_KEYDOWN), func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTVKEYDOWN)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-selchanged
func (me *_EventsTreeView) TvnSelChanged(userFunc func(p *win.NMTREEVIEW)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_SELCHANGED), func(p unsafe.Pointer) {
		userFunc((*win.NMTREEVIEW)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-selchanging
func (me *_EventsTreeView) TvnSelChanging(userFunc func(p *win.NMTREEVIEW) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_SELCHANGING), func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMTREEVIEW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-setdispinfo
func (me *_EventsTreeView) TvnSetDispInfo(userFunc func(p *win.NMTVDISPINFO)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.TVN_SETDISPINFO), func(p unsafe.Pointer) {
		userFunc((*win.NMTVDISPINFO)(p))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-singleexpand
func (me *_EventsTreeView) TvnSingleExpand(userFunc func(p *win.NMTREEVIEW) co.TVNRET) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM(co.TVN_SINGLEEXPAND), func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTREEVIEW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-tree-view
func (me *_EventsTreeView) NmClick(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_CLICK, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-tree-view
func (me *_EventsTreeView) NmCustomDraw(userFunc func(p *win.NMTVCUSTOMDRAW) co.CDRF) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_CUSTOMDRAW, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMTVCUSTOMDRAW)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-tree-view
func (me *_EventsTreeView) NmDblClk(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_DBLCLK, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-tree-view
func (me *_EventsTreeView) NmKillFocus(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_KILLFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-tree-view
func (me *_EventsTreeView) NmRClick(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RCLICK, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tree-view
func (me *_EventsTreeView) NmRDblClk(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RDBLCLK, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-tree-view-
func (me *_EventsTreeView) NmReturn(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_RETURN, func(_ unsafe.Pointer) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setcursor-tree-view-
func (me *_EventsTreeView) NmSetCursor(userFunc func(p *win.NMMOUSE) int) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_SETCURSOR, func(p unsafe.Pointer) uintptr {
		return uintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-tree-view-
func (me *_EventsTreeView) NmSetFocus(userFunc func()) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM_SETFOCUS, func(_ unsafe.Pointer) {
		userFunc()
	})
}
