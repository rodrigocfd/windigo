package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _TreeViewItems struct {
	tv TreeView
}

func (me *_TreeViewItems) new(ctrl TreeView) {
	me.tv = ctrl
}

// Adds a new root item, returning it.
func (me *_TreeViewItems) AddRoot(text string) TreeViewItem {
	return me.Get(win.HTREEITEM(0)).
		AddChild(text)
}

// Retrieves the number of items.
func (me *_TreeViewItems) Count() int {
	return int(me.tv.Hwnd().SendMessage(co.TVM_GETCOUNT, 0, 0))
}

// Deletes all items at once.
func (me *_TreeViewItems) DeleteAll() {
	if me.tv.Hwnd().SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(win.HTREEITEM(0))) == 0 {
		panic("TVM_DELETEITEM for all items failed.")
	}
}

// Retrieves the first visible item, if any.
func (me *_TreeViewItems) FirstVisible() (TreeViewItem, bool) {
	hVisible := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_FIRSTVISIBLE), win.LPARAM(win.HTREEITEM(0))),
	)
	return me.tv.Items().Get(hVisible), hVisible != 0
}

// Returns the item with the given handle.
//
// Note that this method is dumb: no validation is made, the given handle is
// simply kept. If the handle is invalid (or becomes invalid), subsequent
// operations on the TreeViewItem will fail.
func (me *_TreeViewItems) Get(hItem win.HTREEITEM) TreeViewItem {
	return TreeViewItem{tv: me.tv, hItem: hItem}
}

// Retrieves all the root items.
func (me *_TreeViewItems) Roots() []TreeViewItem {
	return me.Get(win.HTREEITEM(0)).Children()
}

// Retrieves the selected item, if any.
func (me *_TreeViewItems) Selected() (TreeViewItem, bool) {
	hItem := win.HTREEITEM(
		me.tv.Hwnd().SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CARET), win.LPARAM(win.HTREEITEM(0))),
	)
	return me.tv.Items().Get(hItem), hItem != 0
}
