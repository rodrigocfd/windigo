/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// A single item row of a list view control.
type ListViewItem struct {
	owner *ListView
	index uint32
}

// Sends LVM_DELETEITEM for this item.
func (me *ListViewItem) Delete() {
	if me.index >= me.owner.ItemCount() { // index out of bounds: ignore
		return
	}
	ret := me.owner.sendLvmMessage(co.LVM_DELETEITEM,
		win.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM failed, index %d.", me.index))
	}
}

// Sends LVM_ENSUREVISIBLE for this item.
// Scrolls the list view if necessary.
func (me *ListViewItem) EnsureVisible() *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_ENSUREVISIBLE,
		win.WPARAM(me.index), win.LPARAM(1)) // always entirely visible
	if ret == 0 {
		panic("LVM_ENSUREVISIBLE failed.")
	}
	return me
}

// Retrieves the image index with LVM_GETITEM.
func (me *ListViewItem) IconIndex() uint32 {
	lvi := win.LVITEM{
		IItem: int32(me.index),
		Mask:  co.LVIF_IMAGE,
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return uint32(lvi.IImage)
}

func (me *ListViewItem) Index() uint32 {
	return me.index
}

// Returns nil if none.
func (me *ListViewItem) NextItem(relationship co.LVNI) *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_GETNEXTITEM,
		win.WPARAM(me.index), win.LPARAM(relationship))
	if int32(ret) == -1 {
		return nil
	}
	return &ListViewItem{
		owner: me.owner,
		index: uint32(ret),
	}
}

func (me *ListViewItem) Owner() *ListView {
	return me.owner
}

// Retrieves the LPARAM with LVM_GETITEM.
func (me *ListViewItem) Param() win.LPARAM {
	lvi := win.LVITEM{
		IItem: int32(me.index),
		Mask:  co.LVIF_PARAM,
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return lvi.LParam
}

// Sends LVM_GETITEMRECT for this item.
// Retrieved coordinates are relative to list view.
func (me *ListViewItem) Rect(portion co.LVIR) *win.RECT {
	rcItem := &win.RECT{
		Left: int32(portion),
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETITEMRECT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&rcItem)))
	if ret == 0 {
		panic("LVM_GETITEMRECT failed.")
	}
	return rcItem
}

// Sets the image index with LVM_SETITEM.
func (me *ListViewItem) SetIconIndex(index uint32) *ListViewItem {
	lvi := win.LVITEM{
		IItem:  int32(me.index),
		Mask:   co.LVIF_IMAGE,
		IImage: int32(index),
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return me
}

// Sets the LPARAM with LVM_SETITEM.
func (me *ListViewItem) SetParam(lParam win.LPARAM) *ListViewItem {
	lvi := win.LVITEM{
		IItem:  int32(me.index),
		Mask:   co.LVIF_PARAM,
		LParam: lParam,
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return me
}

func (me *ListViewItem) SetState(
	state co.LVIS, stateMask co.LVIS) *ListViewItem {

	lvi := win.LVITEM{
		State:     state,
		StateMask: stateMask,
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

func (me *ListViewItem) SetText(text string) *ListViewItem {
	me.SubItem(0).SetText(text)
	return me
}

func (me *ListViewItem) State() co.LVIS {
	return co.LVIS(
		me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_FOCUSED)),
	)
}

func (me *ListViewItem) SubItem(index uint32) *ListViewSubItem {
	numCols := me.owner.ColumnCount()
	if index >= numCols {
		panic("Trying to retrieve sub item with index out of bounds.")
	}
	return &ListViewSubItem{
		item:  me,
		index: index,
	}
}

func (me *ListViewItem) Text() string {
	return me.SubItem(0).Text()
}

// Sends LVM_UPDATE for this item.
func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_UPDATE, win.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}

// Sends LVM_ISITEMVISIBLE for this item.
func (me *ListViewItem) Visible() bool {
	return me.owner.sendLvmMessage(co.LVM_ISITEMVISIBLE,
		win.WPARAM(me.index), 0) != 0
}
