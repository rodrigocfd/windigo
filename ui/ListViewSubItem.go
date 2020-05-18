/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// A cell from a list view item row.
type ListViewSubItem struct {
	item  *ListViewItem
	index uint32
}

func newListViewSubItem(item *ListViewItem, index uint32) *ListViewSubItem {
	return &ListViewSubItem{
		item:  item,
		index: index,
	}
}

func (me *ListViewSubItem) Index() uint32 {
	return me.index
}

func (me *ListViewSubItem) SetText(text string) *ListViewSubItem {
	lvi := api.LVITEM{
		ISubItem: int32(me.index),
		PszText:  api.StrToUtf16Ptr(text),
	}
	ret := me.item.owner.sendLvmMessage(c.LVM_SETITEMTEXT,
		api.WPARAM(me.item.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT failed \"%s\".", text))
	}
	return me
}

func (me *ListViewSubItem) Text() string {
	buf := make([]uint16, 256) // arbitrary
	lvi := api.LVITEM{
		ISubItem:   int32(me.index),
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.item.owner.sendLvmMessage(c.LVM_GETITEMTEXT,
		api.WPARAM(me.item.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic("LVM_GETITEMTEXT failed.")
	}
	return syscall.UTF16ToString(buf)
}
