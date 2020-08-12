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
//
// Scrolls the list view if necessary.
func (me *ListViewItem) EnsureVisible() *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_ENSUREVISIBLE,
		win.WPARAM(me.index), win.LPARAM(1)) // always entirely visible
	if ret == 0 {
		panic("LVM_ENSUREVISIBLE failed.")
	}
	return me
}

// Sets the item as the currently focused one.
func (me *ListViewItem) Focus() *ListViewItem {
	lvi := win.LVITEM{
		State:     co.LVIS_FOCUSED,
		StateMask: co.LVIS_FOCUSED,
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

// Tells if the item is the currently focused one.
func (me *ListViewItem) Focused() bool {
	return co.LVIS(
		me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_FOCUSED)),
	) == co.LVIS_FOCUSED
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

// Returns the index of this item.
func (me *ListViewItem) Index() uint32 {
	return me.index
}

// Returns the ListView to which this item belongs.
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
//
// Retrieved coordinates are relative to list view.
func (me *ListViewItem) Rect(portion co.LVIR) *win.RECT {
	rcItem := &win.RECT{
		Left: int32(portion),
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETITEMRECT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(rcItem)))
	if ret == 0 {
		panic("LVM_GETITEMRECT failed.")
	}
	return rcItem
}

// Selects or unselects the item.
func (me *ListViewItem) Select(isSelected bool) *ListViewItem {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

// Tells if the item is currently selected.
func (me *ListViewItem) Selected() bool {
	return co.LVIS(
		me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_SELECTED)),
	) == co.LVIS_SELECTED
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

// Sends LVM_SETITEMTEXT to change text of first column.
func (me *ListViewItem) SetText(text string) *ListViewItem {
	me.SubItem(0).SetText(text)
	return me
}

// Returns the subitem at the given column index.
//
// Does not perform bound checking.
func (me *ListViewItem) SubItem(index uint32) *ListViewSubItem {
	return &ListViewSubItem{
		item:  me,
		index: index,
	}
}

// Sends LVM_GETITEMTEXT to retrieve the text of the first column.
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
