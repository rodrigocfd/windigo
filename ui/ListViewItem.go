package ui

import (
	"fmt"
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// A single item row of a list view control.
type ListViewItem struct {
	owner *ListView
	index uint32
}

func NewListViewItem(owner *ListView, index uint32) *ListViewItem {
	return &ListViewItem{
		owner: owner,
		index: index,
	}
}

func (me *ListViewItem) Delete() {
	if me.index >= me.owner.ItemCount() { // index out of bounds: ignore
		return
	}

	ret := me.owner.sendLvmMessage(c.LVM_DELETEITEM,
		api.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM failed, index %d.\n", me.index))
	}
}

func (me *ListViewItem) Index() uint32 {
	return me.index
}

func (me *ListViewItem) IsCut() bool {
	sta := me.owner.sendLvmMessage(c.LVM_GETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(c.LVIS_CUT))
	return (c.LVIS(sta) & c.LVIS_CUT) != 0
}

func (me *ListViewItem) IsFocused() bool {
	sta := me.owner.sendLvmMessage(c.LVM_GETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(c.LVIS_FOCUSED))
	return (c.LVIS(sta) & c.LVIS_FOCUSED) != 0
}

func (me *ListViewItem) IsSelected() bool {
	sta := me.owner.sendLvmMessage(c.LVM_GETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(c.LVIS_SELECTED))
	return (c.LVIS(sta) & c.LVIS_SELECTED) != 0
}

func (me *ListViewItem) IsVisible() bool {
	return me.owner.sendLvmMessage(c.LVM_ISITEMVISIBLE,
		api.WPARAM(me.index), 0) != 0
}

func (me *ListViewItem) SetFocus() *ListViewItem {
	lvi := api.LVITEM{
		StateMask: c.LVIS_FOCUSED,
		State:     c.LVIS_FOCUSED,
	}
	ret := me.owner.sendLvmMessage(c.LVM_SETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_FOCUSED.")
	}
	return me
}

func (me *ListViewItem) SetSelected(selected bool) *ListViewItem {
	lvi := api.LVITEM{
		StateMask: c.LVIS_SELECTED,
	}
	if selected { // otherwise is zero
		lvi.State = c.LVIS_SELECTED
	}
	ret := me.owner.sendLvmMessage(c.LVM_SETITEMSTATE,
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
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
	return NewListViewSubItem(me, index)
}

func (me *ListViewItem) Text() string {
	return me.SubItem(0).Text()
}

func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.sendLvmMessage(c.LVM_UPDATE, api.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}
