//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// The items collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [TreeView].
type CollectionTreeViewItems struct {
	owner *TreeView
}

// Adds a new root item with [TVM_INSERTITEM], returning the new item.
//
// The iconIndex is the zero-based index of the icon previously inserted into
// the control's image list, or -1 for no icon.
//
// [TVM_INSERTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-insertitem
func (me *CollectionTreeViewItems) AddRoot(text string, iconIndex int) TreeViewItem {
	return me.Get(win.HTREEITEM(0)).
		AddChild(text, iconIndex)
}

// Retrieves the total number of items in the control, with [TVM_GETCOUNT].
//
// [TVM_GETCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getcount
func (me *CollectionTreeViewItems) Count() int {
	c, _ := me.owner.hWnd.SendMessage(co.TVM_GETCOUNT, 0, 0)
	return int(c)
}

// Deletes all items at once with [TVM_DELETEITEM].
//
// Panics on error.
//
// [TVM_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-deleteitem
func (me *CollectionTreeViewItems) DeleteAll() {
	ret, err := me.owner.hWnd.SendMessage(co.TVM_DELETEITEM,
		0, win.LPARAM(win.HTREEITEM(0)))
	if ret == 0 || err != nil {
		panic("TVM_DELETEITEM for all items failed.")
	}
}

// Retrieves the first visible item, if any, with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me *CollectionTreeViewItems) FirstVisible() (TreeViewItem, bool) {
	hVisible, _ := me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_FIRSTVISIBLE), win.LPARAM(win.HTREEITEM(0)))
	if hVisible != 0 {
		return TreeViewItem{me.owner, win.HTREEITEM(hVisible)}, true
	}
	return TreeViewItem{}, false
}

// Returns the item with the given handle.
func (me *CollectionTreeViewItems) Get(hItem win.HTREEITEM) TreeViewItem {
	return TreeViewItem{me.owner, hItem}
}

// Returns the root items with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me *CollectionTreeViewItems) Roots() []TreeViewItem {
	roof := TreeViewItem{me.owner, win.HTREEITEM(0)}
	return roof.Children()
}

// Retrieves the selected item, if any, with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me *CollectionTreeViewItems) Selected() (TreeViewItem, bool) {
	hItem, _ := me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_CARET), win.LPARAM(win.HTREEITEM(0)))
	if hItem != 0 {
		return TreeViewItem{me.owner, win.HTREEITEM(hItem)}, true
	}
	return TreeViewItem{}, false
}
