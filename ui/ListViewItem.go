/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// A single item row of a list view control.
type ListViewItem struct {
	owner *ListView
	index uint
}

// Deletes this item with LVM_DELETEITEM.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-deleteitem
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

// Scrolls the list view so this item becomes visible, with LVM_ENSUREVISIBLE.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-ensurevisible
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

// Retrieves the image index with LVM_GETITEM; or -1 if no image.
func (me *ListViewItem) IconIndex() int {
	lvi := win.LVITEM{
		IItem: int32(me.index),
		Mask:  co.LVIF_IMAGE,
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_GETITEM failed.")
	}
	return int(lvi.IImage)
}

// Returns the index of this item.
func (me *ListViewItem) Index() uint {
	return me.index
}

// Tells if the item is the currently focused one.
func (me *ListViewItem) IsFocused() bool {
	return co.LVIS(
		me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_FOCUSED)),
	) == co.LVIS_FOCUSED
}

// Tells if the item is currently selected.
func (me *ListViewItem) IsSelected() bool {
	return co.LVIS(
		me.owner.sendLvmMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_SELECTED)),
	) == co.LVIS_SELECTED
}

// Checks if this item is visible with LVM_ISITEMVISIBLE.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-isitemvisible
func (me *ListViewItem) IsVisible() bool {
	return me.owner.sendLvmMessage(co.LVM_ISITEMVISIBLE,
		win.WPARAM(me.index), 0) != 0
}

// Returns the ListView to which this item belongs.
func (me *ListViewItem) Owner() *ListView {
	return me.owner
}

// Retrieves the LPARAM associated to this item with LVM_GETITEM.
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

// Retrieves bound coordinates of the item with LVM_GETITEMRECT. Coordinates are
// relative to list view.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemrect
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

// Sets the image index with LVM_SETITEM; or -1 for no image.
func (me *ListViewItem) SetIconIndex(index int) *ListViewItem {
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

// Sets the LPARAM associated to this item with LVM_SETITEM.
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

// Sends LVM_SETITEMTEXT to change the text below any column.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setitemtext
func (me *ListViewItem) SetSubItemText(
	columnIndex uint, text string) *ListViewItem {

	textBuf := win.Str.ToUint16Slice(text)
	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
		PszText:  &textBuf[0],
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETITEMTEXT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT failed \"%s\".", text))
	}
	return me
}

// Sends LVM_SETITEMTEXT to change text of first column.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setitemtext
func (me *ListViewItem) SetText(text string) *ListViewItem {
	return me.SetSubItemText(0, text)
}

// Sends LVM_GETITEMTEXT to retrieve the text below any column.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemtext
func (me *ListViewItem) SubItemText(columnIndex uint) string {
	buf := [256]uint16{} // arbitrary
	lvi := win.LVITEM{
		ISubItem:   int32(columnIndex),
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETITEMTEXT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic("LVM_GETITEMTEXT failed.")
	}
	return syscall.UTF16ToString(buf[:])
}

// Sends LVM_GETITEMTEXT to retrieve the text of the first column.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemtext
func (me *ListViewItem) Text() string {
	return me.SubItemText(0)
}

// Updates the item with LVM_UPDATE.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-update
func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.sendLvmMessage(co.LVM_UPDATE, win.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}
