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

func (me *_TreeViewItems) getUnchecked(hItem win.HTREEITEM) TreeViewItem {
	item := _TreeViewItem{}
	item.new(me.pHwnd, hItem)
	return &item
}

// Adds a new root item, returning it.
func (me *_TreeViewItems) AddRoot(text string) TreeViewItem {
	tvi := win.TVINSERTSTRUCT{}
	tvi.HInsertAfter = win.HTREEITEM(co.HTREEITEM_LAST)
	tvi.Itemex.Mask = co.TVIF_TEXT
	tvi.Itemex.SetPszText(win.Str.ToUint16Slice(text))

	hNewItem := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&tvi))),
	)
	if hNewItem == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}
	return me.getUnchecked(hNewItem)
}

// Retrieves the number of items.
func (me *_TreeViewItems) Count() int {
	return int(me.pHwnd.SendMessage(co.TVM_GETCOUNT, 0, 0))
}

// Deletes all items at once.
func (me *_TreeViewItems) DeleteAll() {
	if me.pHwnd.SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(win.HTREEITEM(0))) == 0 {
		panic("TVM_DELETEITEM for all items failed.")
	}
}

// Retrieves the first visible item, if any.
func (me *_TreeViewItems) FirstVisible() (TreeViewItem, bool) {
	hVisible := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_FIRSTVISIBLE), win.LPARAM(win.HTREEITEM(0))),
	)
	if hVisible == 0 {
		return nil, false
	}
	return me.getUnchecked(hVisible), true
}

// Returns the item with the given handle, if any.
func (me *_TreeViewItems) Get(hItem win.HTREEITEM) (TreeViewItem, bool) {
	tvi := win.TVITEMEX{
		HItem: hItem,
		Mask:  co.TVIF_HANDLE,
	}

	ret := me.pHwnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		return nil, false
	}
	return me.getUnchecked(hItem), true
}

// Retrieves all the root items.
func (me *_TreeViewItems) Roots() []TreeViewItem {
	return me.getUnchecked(win.HTREEITEM(0)).Children()
}

// Retrieves the selected item, if any.
func (me *_TreeViewItems) Selected() (TreeViewItem, bool) {
	hItem := win.HTREEITEM(
		me.pHwnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_CARET), win.LPARAM(win.HTREEITEM(0))),
	)

	if hItem == 0 {
		return nil, false
	}
	return me.getUnchecked(hItem), true
}
