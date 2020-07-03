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

func (me *ListViewItem) Index() uint32 {
	return me.index
}

func (me *ListViewItem) IsCut() bool {
	sta := me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(co.LVIS_CUT))
	return (co.LVIS(sta) & co.LVIS_CUT) != 0
}

func (me *ListViewItem) IsFocused() bool {
	sta := me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(co.LVIS_FOCUSED))
	return (co.LVIS(sta) & co.LVIS_FOCUSED) != 0
}

func (me *ListViewItem) IsSelected() bool {
	sta := me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(co.LVIS_SELECTED))
	return (co.LVIS(sta) & co.LVIS_SELECTED) != 0
}

func (me *ListViewItem) IsVisible() bool {
	return me.owner.sendLvmMessage(co.LVM_ISITEMVISIBLE,
		win.WPARAM(me.index), 0) != 0
}

func (me *ListViewItem) Owner() *ListView {
	return me.owner
}

func (me *ListViewItem) SetFocus() *ListViewItem {
	lvi := win.LVITEM{
		StateMask: co.LVIS_FOCUSED,
		State:     co.LVIS_FOCUSED,
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_FOCUSED.")
	}
	return me
}

func (me *ListViewItem) SetSelected(selected bool) *ListViewItem {
	lvi := win.LVITEM{
		StateMask: co.LVIS_SELECTED,
	}
	if selected { // otherwise remains zero
		lvi.State = co.LVIS_SELECTED
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_SELECTED.")
	}
	return me
}

func (me *ListViewItem) SetText(text string) *ListViewItem {
	me.SubItem(0).SetText(text)
	return me
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

func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_UPDATE, win.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}
