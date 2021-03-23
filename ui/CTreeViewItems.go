package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _TreeViewItems struct {
	pHwnd *win.HWND
}

func (me *_TreeViewItems) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
}

// Adds a new child item.
func (me *_TreeViewItems) AddChild(
	hItemParent win.HTREEITEM, text string) win.HTREEITEM {

	textBuf := win.Str.ToUint16Slice(text)
	tvi := win.TVINSERTSTRUCT{
		HParent:      hItemParent,
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: &textBuf[0],
		},
	}

	ret := me.pHwnd.SendMessage(co.TVM_INSERTITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return win.HTREEITEM(ret)
}

// Adds a new root item.
func (me *_TreeViewItems) AddRoot(text string) win.HTREEITEM {
	textBuf := win.Str.ToUint16Slice(text)
	tvi := win.TVINSERTSTRUCT{
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: &textBuf[0],
		},
	}

	ret := me.pHwnd.SendMessage(co.TVM_INSERTITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return win.HTREEITEM(ret)
}

// Retrieves all child items of the given item.
func (me *_TreeViewItems) Children(hItemParent win.HTREEITEM) []win.HTREEITEM {
	hChildren := make([]win.HTREEITEM, 0)
	hItem := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CHILD), win.LPARAM(hItemParent)), // retrieve first child
	)
	hasSibling := hItem != 0 // has 1st child?

	for hasSibling {
		hChildren = append(hChildren, hItem)
		hItem, hasSibling = me.NextSibling(hItem)
	}

	return hChildren
}

// Retrieves the number of items.
func (me *_TreeViewItems) Count() int {
	return int(me.pHwnd.SendMessage(co.TVM_GETCOUNT, 0, 0))
}

// Deletes an item and all its children.
func (me *_TreeViewItems) Delete(hItem win.HTREEITEM) {
	if me.pHwnd.SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(hItem)) == 0 {
		panic("TVM_DELETEITEM failed.")
	}
}

// Deletes all items at once.
func (me *_TreeViewItems) DeleteAll() {
	if me.pHwnd.SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(win.HTREEITEM(0))) == 0 {
		panic("TVM_DELETEITEM for all items failed.")
	}
}

// Expand the item and scrolls the TreeView so it becomes visible.
func (me *_TreeViewItems) EnsureVisible(hItem win.HTREEITEM) {
	me.pHwnd.SendMessage(co.TVM_ENSUREVISIBLE, 0, win.LPARAM(hItem))
}

// Expands or collapses the item.
func (me *_TreeViewItems) Expand(hItem win.HTREEITEM, isExpanded bool) {
	flag := co.TVE_EXPAND
	if !isExpanded {
		flag = co.TVE_COLLAPSE
	}
	me.pHwnd.SendMessage(co.TVM_EXPAND, win.WPARAM(flag), win.LPARAM(hItem))
}

// Retrieves the first visible item, if any.
func (me *_TreeViewItems) FirstVisible() (win.HTREEITEM, bool) {
	hVisible := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_FIRSTVISIBLE), win.LPARAM(win.HTREEITEM(0))),
	)

	if hVisible == 0 {
		return win.HTREEITEM(0), false
	}
	return hVisible, true
}

// Tells if the item is currently expanded.
func (me *_TreeViewItems) IsExpanded(hItem win.HTREEITEM) bool {
	return (co.TVIS(
		me.pHwnd.SendMessage(co.TVM_GETITEMSTATE,
			win.WPARAM(hItem), win.LPARAM(co.TVIS_EXPANDED)),
	) & co.TVIS_EXPANDED) != 0
}

// Tells if the item is a root.
func (me *_TreeViewItems) IsRoot(hItem win.HTREEITEM) bool {
	_, hasParent := me.Parent(hItem)
	return !hasParent
}

// Retrieves the associated LPARAM with TVM_GETITEM.
func (me *_TreeViewItems) Param(hItem win.HTREEITEM) win.LPARAM {
	tvi := win.TVITEMEX{
		HItem: hItem,
		Mask:  co.TVIF_PARAM,
	}

	ret := me.pHwnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return tvi.LParam
}

// Retrieves the next sibling item, if any.
func (me *_TreeViewItems) NextSibling(
	hItem win.HTREEITEM) (win.HTREEITEM, bool) {

	hSibling := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_NEXT), win.LPARAM(hItem)),
	)

	if hSibling == 0 {
		return win.HTREEITEM(0), false
	}
	return hSibling, true
}

// Retrieves the parent item, if any. If no parent, the item is a root.
func (me *_TreeViewItems) Parent(hItem win.HTREEITEM) (win.HTREEITEM, bool) {
	hParent := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PARENT), win.LPARAM(hItem)),
	)

	if hParent == 0 {
		return win.HTREEITEM(0), false
	}
	return hParent, true
}

// Retrieves the previous sibling item, if any.
func (me *_TreeViewItems) PrevSibling(
	hItem win.HTREEITEM) (win.HTREEITEM, bool) {

	hSibling := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PREVIOUS), win.LPARAM(hItem)),
	)

	if hSibling == 0 {
		return win.HTREEITEM(0), false
	}
	return hSibling, true
}

// Retrieves all the root items.
func (me *_TreeViewItems) Roots() []win.HTREEITEM {
	return me.Children(win.HTREEITEM(0))
}

// Retrieves the selected item, if any.
func (me *_TreeViewItems) Selected() (win.HTREEITEM, bool) {
	hItem := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CARET), win.LPARAM(win.HTREEITEM(0))),
	)

	if hItem == 0 {
		return win.HTREEITEM(0), false
	}
	return hItem, true
}

// Sets the associated LPARAM.
func (me *_TreeViewItems) SetLParam(hItem win.HTREEITEM, lp win.LPARAM) {
	tvi := win.TVITEMEX{
		HItem:  hItem,
		Mask:   co.TVIF_PARAM,
		LParam: lp,
	}

	ret := me.pHwnd.SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_SETITEM failed.")
	}
}

// Sets the text of an item.
func (me *_TreeViewItems) SetText(hItem win.HTREEITEM, text string) {
	textBuf := win.Str.ToUint16Slice(text)
	tvi := win.TVITEMEX{
		HItem:   hItem,
		Mask:    co.TVIF_TEXT,
		PszText: &textBuf[0],
	}

	ret := me.pHwnd.SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
}

// Retrieves the text of an item.
func (me *_TreeViewItems) Text(hItem win.HTREEITEM) string {
	buf := [256]uint16{} // arbitrary
	tvi := win.TVITEMEX{
		HItem:      hItem,
		Mask:       co.TVIF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}

	ret := me.pHwnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return win.Str.FromUint16Slice(buf[:])
}

// Toggles the item, expanding or collapsing it.
func (me *_TreeViewItems) ToggleExpand(hItem win.HTREEITEM) {
	me.pHwnd.SendMessage(co.TVM_EXPAND,
		win.WPARAM(co.TVE_TOGGLE), win.LPARAM(hItem))
}
