/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// A single item node of a tree view control.
type TreeViewItem struct {
	owner     *TreeView
	hTreeItem win.HTREEITEM
}

// Adds a new child item; returns the newly inserted item.
func (me *TreeViewItem) AddChild(text string) *TreeViewItem {
	textBuf := win.StrToSlice(text)
	tvi := win.TVINSERTSTRUCT{
		HParent:      me.hTreeItem,
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: uintptr(unsafe.Pointer(&textBuf[0])),
		},
	}
	ret := me.owner.sendTvmMessage(co.TVM_INSERTITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_INSERTITEM failed \"%s\".", text))
	}
	return &TreeViewItem{
		owner:     me.owner,
		hTreeItem: win.HTREEITEM(ret),
	}
}

func (me *TreeViewItem) Children() []TreeViewItem {
	retNodes := make([]TreeViewItem, 0)
	node := me.FirstChild()
	for node != nil {
		retNodes = append(retNodes, *node)
		node = node.NextSibling()
	}
	return retNodes
}

// Expand or collapses the item with TVM_EXPAND.
func (me *TreeViewItem) Expand(flags co.TVE) *TreeViewItem {
	ret := me.owner.sendTvmMessage(co.TVM_EXPAND,
		win.WPARAM(flags), win.LPARAM(me.hTreeItem))
	if ret == 0 {
		panic("TVM_EXPAND failed.")
	}
	return me
}

// Sends TVM_GETNEXTITEM with TVGN_CHILD, returns nil if none.
func (me *TreeViewItem) FirstChild() *TreeViewItem {
	return me.NextItem(co.TVGN_CHILD)
}

func (me *TreeViewItem) HTreeItem() win.HTREEITEM {
	return me.hTreeItem
}

// Sends TVM_GETNEXTITEM, returns nil if none found.
func (me *TreeViewItem) NextItem(flags co.TVGN) *TreeViewItem {
	ret := me.owner.sendTvmMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(flags), win.LPARAM(me.hTreeItem))
	if ret == 0 {
		return nil
	}
	return &TreeViewItem{
		owner:     me.owner,
		hTreeItem: win.HTREEITEM(ret),
	}
}

// Sends TVM_GETNEXTITEM with TVGN_NEXT, returns nil if none.
func (me *TreeViewItem) NextSibling() *TreeViewItem {
	return me.NextItem(co.TVGN_NEXT)
}

// Returns the TreeView to which this item belongs.
func (me *TreeViewItem) Owner() *TreeView {
	return me.owner
}

// Retrieves the associated LPARAM with TVM_GETITEM.
func (me *TreeViewItem) Param() win.LPARAM {
	tvi := win.TVITEMEX{
		HItem: me.hTreeItem,
		Mask:  co.TVIF_PARAM,
	}
	ret := me.owner.sendTvmMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return tvi.LParam
}

// Sends TVM_GETNEXTITEM with TVGN_PARENT, returns nil if none.
func (me *TreeViewItem) Parent() *TreeViewItem {
	return me.NextItem(co.TVGN_PARENT)
}

// Sends TVM_GETNEXTITEM with TVGN_PREVIOUS, returns nil if none.
func (me *TreeViewItem) PrevSibling() *TreeViewItem {
	return me.NextItem(co.TVGN_PREVIOUS)
}

// Sets the associated LPARAM with TVM_SETITEM.
func (me *TreeViewItem) SetParam(lp win.LPARAM) *TreeViewItem {
	tvi := win.TVITEMEX{
		HItem:  me.hTreeItem,
		Mask:   co.TVIF_PARAM,
		LParam: lp,
	}
	ret := me.owner.sendTvmMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_SETITEM failed.")
	}
	return me
}

// Sets the text with TVM_SETITEM.
func (me *TreeViewItem) SetText(text string) *TreeViewItem {
	textBuf := win.StrToSlice(text)
	tvi := win.TVITEMEX{
		HItem:   me.hTreeItem,
		Mask:    co.TVIF_TEXT,
		PszText: uintptr(unsafe.Pointer(&textBuf[0])),
	}
	ret := me.owner.sendTvmMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
	return me
}

// Retrieves the text with TVM_GETITEM.
func (me *TreeViewItem) Text() string {
	buf := [256]uint16{} // arbitrary
	tvi := win.TVITEMEX{
		HItem:      me.hTreeItem,
		Mask:       co.TVIF_TEXT,
		PszText:    uintptr(unsafe.Pointer(&buf[0])),
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.sendTvmMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return syscall.UTF16ToString(buf[:])
}
