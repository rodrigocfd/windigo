package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single item of a TreeView.
type TreeViewItem interface {
	AddChild(text string) TreeViewItem // Adds a child to this item.
	Children() []TreeViewItem          // Retrieves all the children of this item.
	Delete()                           // Deletes the item and all its children.
	EnsureVisible()                    // Makes sure the item is visible, scrolling the TreeView if needed.
	Expand(doExpand bool)              // Expands the item.
	Htreeitem() win.HTREEITEM          // Returns the unique handle of the item.
	IsExpanded() bool                  // Tells whether the item is currently expanded.
	IsRoot() bool                      // Tells whether the item is a root item (has no parent).
	LParam() win.LPARAM                // Retrieves the custom data associated with the item.
	NextSibling() (TreeViewItem, bool) // Retrieves the next item, if any.
	Parent() (TreeViewItem, bool)      // Retrieves the parent item, if any.
	PrevSibling() (TreeViewItem, bool) // Retrieves the previous item, if any.
	SetLParam(lp win.LPARAM)           // Sets the custom data associated with the item.
	SetText(text string)               // Sets the text of the item.
	Text() string                      // Retrieves the text of the item.
}

//------------------------------------------------------------------------------

type _TreeViewItem struct {
	tv    TreeView
	hItem win.HTREEITEM
}

func (me *_TreeViewItem) new(ctrl TreeView, hItem win.HTREEITEM) {
	me.tv = ctrl
	me.hItem = hItem
}

func (me *_TreeViewItem) AddChild(text string) TreeViewItem {
	tvi := win.TVINSERTSTRUCT{}
	tvi.HParent = me.hItem
	tvi.HInsertAfter = win.HTREEITEM(co.HTREEITEM_LAST)
	tvi.Itemex.Mask = co.TVIF_TEXT
	tvi.Itemex.SetPszText(win.Str.ToNativeSlice(text))

	hNewItem := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&tvi))),
	)
	if hNewItem == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}

	return &_TreeViewItem{
		tv:    me.tv,
		hItem: win.HTREEITEM(hNewItem),
	}
}

func (me *_TreeViewItem) Children() []TreeViewItem {
	hChildren := make([]TreeViewItem, 0)
	hItem := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CHILD), win.LPARAM(me.hItem)), // retrieve first child
	)
	hasSibling := hItem != 0 // has 1st child?

	for hasSibling {
		hChildren = append(hChildren, &_TreeViewItem{
			tv:    me.tv,
			hItem: hItem,
		})

		hItem = win.HTREEITEM( // retrieve the next siblings
			me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
				win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem)),
		)
		hasSibling = hItem != 0
	}

	return hChildren
}

func (me *_TreeViewItem) Delete() {
	if me.tv.Hwnd().SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(me.hItem)) == 0 {
		panic("TVM_DELETEITEM failed.")
	}
}

func (me *_TreeViewItem) EnsureVisible() {
	me.tv.Hwnd().SendMessage(co.TVM_ENSUREVISIBLE, 0, win.LPARAM(me.hItem))
}

func (me *_TreeViewItem) Expand(doExpand bool) {
	me.tv.Hwnd().SendMessage(co.TVM_EXPAND,
		win.WPARAM(util.Iif(doExpand, co.TVE_EXPAND, co.TVE_COLLAPSE).(co.TVE)),
		win.LPARAM(me.hItem))
}

func (me *_TreeViewItem) Htreeitem() win.HTREEITEM {
	return me.hItem
}

func (me *_TreeViewItem) IsExpanded() bool {
	return (co.TVIS(
		me.tv.Hwnd().SendMessage(co.TVM_GETITEMSTATE,
			win.WPARAM(me.hItem), win.LPARAM(co.TVIS_EXPANDED)),
	) & co.TVIS_EXPANDED) != 0
}

func (me *_TreeViewItem) IsRoot() bool {
	_, hasParent := me.Parent()
	return !hasParent
}

func (me *_TreeViewItem) LParam() win.LPARAM {
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

func (me *_TreeViewItem) NextSibling() (TreeViewItem, bool) {
	hSibling := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem)),
	)

	if hSibling == 0 {
		return nil, false
	}
	return &_TreeViewItem{
		tv:    me.tv,
		hItem: hSibling,
	}, true
}

func (me *_TreeViewItem) Parent() (TreeViewItem, bool) {
	hParent := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PARENT), win.LPARAM(me.hItem)),
	)

	if hParent == 0 {
		return nil, false
	}
	return &_TreeViewItem{
		tv:    me.tv,
		hItem: hParent,
	}, true
}

func (me *_TreeViewItem) PrevSibling() (TreeViewItem, bool) {
	hSibling := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_PREVIOUS), win.LPARAM(me.hItem)),
	)

	if hSibling == 0 {
		return nil, false
	}
	return &_TreeViewItem{
		tv:    me.tv,
		hItem: hSibling,
	}, true
}

func (me *_TreeViewItem) SetLParam(lp win.LPARAM) {
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

func (me *_TreeViewItem) SetText(text string) {
	tvi := win.TVITEMEX{}
	tvi.HItem = me.hItem
	tvi.Mask = co.TVIF_TEXT
	tvi.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.tv.Hwnd().SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
}

func (me *_TreeViewItem) Text() string {
	buf := [256]uint16{} // arbitrary

	tvi := win.TVITEMEX{}
	tvi.HItem = me.hItem
	tvi.Mask = co.TVIF_TEXT
	tvi.SetPszText(buf[:])

	ret := me.tv.Hwnd().SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return win.Str.FromNativeSlice(buf[:])
}
