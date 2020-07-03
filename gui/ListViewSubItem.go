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

// A cell from a list view item row.
type ListViewSubItem struct {
	item  *ListViewItem
	index uint32
}

func (me *ListViewSubItem) Index() uint32 {
	return me.index
}

func (me *ListViewSubItem) OwnerItem() *ListViewItem {
	return me.item
}

func (me *ListViewSubItem) SetText(text string) *ListViewSubItem {
	lvi := win.LVITEM{
		ISubItem: int32(me.index),
		PszText:  win.StrToPtr(text),
	}
	ret := me.item.owner.sendLvmMessage(co.LVM_SETITEMTEXT,
		win.WPARAM(me.item.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT failed \"%s\".", text))
	}
	return me
}

func (me *ListViewSubItem) Text() string {
	buf := make([]uint16, 256) // arbitrary
	lvi := win.LVITEM{
		ISubItem:   int32(me.index),
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.item.owner.sendLvmMessage(co.LVM_GETITEMTEXT,
		win.WPARAM(me.item.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic("LVM_GETITEMTEXT failed.")
	}
	return syscall.UTF16ToString(buf)
}
