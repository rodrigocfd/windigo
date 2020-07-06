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

// Returns the column index of this subitem.
func (me *ListViewSubItem) Index() uint32 {
	return me.index
}

// Returns the ListViewItem to which this subitem belongs.
func (me *ListViewSubItem) OwnerItem() *ListViewItem {
	return me.item
}

// Sends LVM_SETITEMTEXT to change the text.
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

// Sends LVM_GETITEMTEXT to retrieve the text.
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
