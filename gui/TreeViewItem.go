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
	tvi := win.TVINSERTSTRUCT{
		HParent:      me.hTreeItem,
		HInsertAfter: win.HTREEITEM(co.HTREEITEM_LAST),
		Itemex: win.TVITEMEX{
			Mask:    co.TVIF_TEXT,
			PszText: win.StrToUtf16Ptr(text),
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

func (me *TreeViewItem) Expand(flags co.TVE) *TreeViewItem {
	ret := me.owner.sendTvmMessage(co.TVM_EXPAND,
		win.WPARAM(flags), win.LPARAM(me.hTreeItem))
	if ret == 0 {
		panic("TVM_EXPAND failed.")
	}
	return me
}

func (me *TreeViewItem) FirstChild() *TreeViewItem {
	ret := me.owner.sendTvmMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_ROOT), win.LPARAM(me.hTreeItem))
	if ret == 0 {
		return nil
	}
	return &TreeViewItem{
		owner:     me.owner,
		hTreeItem: win.HTREEITEM(ret),
	}
}

func (me *TreeViewItem) HTreeItem() win.HTREEITEM {
	return me.hTreeItem
}

func (me *TreeViewItem) NextSibling() *TreeViewItem {
	ret := me.owner.sendTvmMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hTreeItem))
	if ret == 0 {
		return nil
	}
	return &TreeViewItem{
		owner:     me.owner,
		hTreeItem: win.HTREEITEM(ret),
	}
}

func (me *TreeViewItem) Owner() *TreeView {
	return me.owner
}

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

func (me *TreeViewItem) PrevSibling() *TreeViewItem {
	ret := me.owner.sendTvmMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_PREVIOUS), win.LPARAM(me.hTreeItem))
	if ret == 0 {
		return nil
	}
	return &TreeViewItem{
		owner:     me.owner,
		hTreeItem: win.HTREEITEM(ret),
	}
}

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

func (me *TreeViewItem) SetText(text string) *TreeViewItem {
	tvi := win.TVITEMEX{
		HItem:   me.hTreeItem,
		Mask:    co.TVIF_TEXT,
		PszText: win.StrToUtf16Ptr(text),
	}
	ret := me.owner.sendTvmMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
	return me
}

func (me *TreeViewItem) Text() string {
	buf := make([]uint16, 256) // arbitrary
	tvi := win.TVITEMEX{
		HItem:      me.hTreeItem,
		Mask:       co.TVIF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.sendTvmMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 {
		panic("TVM_GETITEM failed.")
	}
	return syscall.UTF16ToString(buf)
}
