package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single item of a TreeView.
type TreeViewItem struct {
	tv    TreeView
	hItem win.HTREEITEM
}

// Adds a child to this item.
func (me TreeViewItem) AddChild(text string) TreeViewItem {
	tvi := win.TVINSERTSTRUCT{
		HParent:      me.hItem,
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask: co.TVIF_TEXT,
		},
	}
	tvi.Itemex.SetPszText(win.Str.ToNativeSlice(text))

	hNewItem := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&tvi))),
	)
	if hNewItem == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return me.tv.Items().Get(hNewItem)
}

// Retrieves all the children of this item.
func (me TreeViewItem) Children() []TreeViewItem {
	hChildren := make([]TreeViewItem, 0)
	hItem := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CHILD), win.LPARAM(me.hItem)), // retrieve first child
	)
	hasSibling := hItem != 0 // has 1st child?

	for hasSibling {
		hChildren = append(hChildren, me.tv.Items().Get(hItem))

		hItem = win.HTREEITEM( // retrieve the next siblings
			me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
				win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem)),
		)
		hasSibling = hItem != 0
	}

	return hChildren
}

// Deletes the item and all its children.
func (me TreeViewItem) Delete() {
	if me.tv.Hwnd().SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(me.hItem)) == 0 {
		panic("TVM_DELETEITEM failed.")
	}
}

// Makes sure the item is visible, scrolling the TreeView if needed.
func (me TreeViewItem) EnsureVisible() {
	me.tv.Hwnd().SendMessage(co.TVM_ENSUREVISIBLE, 0, win.LPARAM(me.hItem))
}

// Expands the item.
func (me TreeViewItem) Expand(doExpand bool) {
	me.tv.Hwnd().SendMessage(co.TVM_EXPAND,
		win.WPARAM(util.Iif(doExpand, co.TVE_EXPAND, co.TVE_COLLAPSE).(co.TVE)),
		win.LPARAM(me.hItem))
}

// Returns the unique handle of the item.
func (me TreeViewItem) Htreeitem() win.HTREEITEM {
	return me.hItem
}

// Tells whether the item is currently expanded.
func (me TreeViewItem) IsExpanded() bool {
	return (co.TVIS(
		me.tv.Hwnd().SendMessage(co.TVM_GETITEMSTATE,
			win.WPARAM(me.hItem), win.LPARAM(co.TVIS_EXPANDED)),
	) & co.TVIS_EXPANDED) != 0
}

// Tells whether the item is a root item (has no parent).
func (me TreeViewItem) IsRoot() bool {
	_, hasParent := me.Parent()
	return !hasParent
}

// Retrieves the custom data associated with the item.
func (me TreeViewItem) LParam() win.LPARAM {
	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_PARAM,
	}

	ret := me.tv.Hwnd().SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return tvi.LParam
}

// Retrieves the next item, if any.
func (me TreeViewItem) NextSibling() (TreeViewItem, bool) {
	hSibling := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem)),
	)
	return me.tv.Items().Get(hSibling), hSibling != 0
}

// Retrieves the parent item, if any.
func (me TreeViewItem) Parent() (TreeViewItem, bool) {
	hParent := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PARENT), win.LPARAM(me.hItem)),
	)
	return me.tv.Items().Get(hParent), hParent != 0
}

// Retrieves the previous item, if any.
func (me TreeViewItem) PrevSibling() (TreeViewItem, bool) {
	hSibling := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PREVIOUS), win.LPARAM(me.hItem)),
	)
	return me.tv.Items().Get(hSibling), hSibling != 0
}

// Sets the custom data associated with the item.
func (me TreeViewItem) SetLParam(lp win.LPARAM) {
	tvi := win.TVITEMEX{
		HItem:  me.hItem,
		Mask:   co.TVIF_PARAM,
		LParam: lp,
	}

	ret := me.tv.Hwnd().SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_SETITEM failed.")
	}
}

// Sets the text of the item.
func (me TreeViewItem) SetText(text string) {
	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_TEXT,
	}
	tvi.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.tv.Hwnd().SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
}

// Retrieves the text of the item.
func (me TreeViewItem) Text() string {
	var buf [256]uint16 // arbitrary

	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_TEXT,
	}
	tvi.SetPszText(buf[:])

	ret := me.tv.Hwnd().SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return win.Str.FromNativeSlice(buf[:])
}
