//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single item of a ListView.
type ListViewItem struct {
	lv    ListView
	index uint32
}

// Deletes the item.
func (me ListViewItem) Delete() {
	ret := me.lv.Hwnd().SendMessage(co.LVM_DELETEITEM, win.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM %d failed.", me.index))
	}
}

// Makes sure the item is visible, scrolling the ListView if needed.
func (me ListViewItem) EnsureVisible() {
	if me.lv.View() == co.LV_VIEW_DETAILS {
		// In details view, LVM_ENSUREVISIBLE won't center the item vertically.
		// That's what we do here.
		rc := me.lv.Hwnd().GetClientRect()
		cyList := rc.Bottom // total height of the listview

		lvii := win.LVITEMINDEX{
			IItem: int32(me.lv.Hwnd().SendMessage(co.LVM_GETTOPINDEX, 0, 0)), // 1st visible item
		}
		rc = win.RECT{
			Left: int32(co.LVIR_BOUNDS),
		}

		ret := me.lv.Hwnd().SendMessage(co.LVM_GETITEMINDEXRECT,
			win.WPARAM(unsafe.Pointer(&lvii)), win.LPARAM(unsafe.Pointer(&rc)))
		if ret == 0 {
			panic(fmt.Sprintf("LVM_GETITEMINDEXRECT %d failed.", lvii.IItem))
		}
		cyItem := rc.Bottom - rc.Top // height of a single item
		xTop := rc.Top               // topmost X of 1st visible item

		lvii = win.LVITEMINDEX{
			IItem: int32(me.index),
		}
		rc = win.RECT{}

		ret = me.lv.Hwnd().SendMessage(co.LVM_GETITEMINDEXRECT,
			win.WPARAM(unsafe.Pointer(&lvii)), win.LPARAM(unsafe.Pointer(&rc)))
		if ret == 0 {
			panic(fmt.Sprintf("LVM_GETITEMINDEXRECT %d failed.", lvii.IItem))
		}
		xUs := rc.Top // our current X

		if xUs < xTop || xUs > xTop+cyList { // if we're not visible
			me.lv.Scroll(0, int(xUs-xTop-cyList/2+cyItem*2))
		}

	} else {
		ret := me.lv.Hwnd().SendMessage(co.LVM_ENSUREVISIBLE,
			win.WPARAM(me.index), win.LPARAM(1)) // always entirely visible
		if ret == 0 {
			panic(fmt.Sprintf("LVM_ENSUREVISIBLE %d failed.", me.index))
		}
	}
}

// Sets the item as the focused one.
func (me ListViewItem) Focus() {
	lvi := win.LVITEM{
		State:     co.LVIS_FOCUSED,
		StateMask: co.LVIS_FOCUSED,
	}

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if int(ret) == -1 {
		panic(fmt.Sprintf("LVM_SETITEMSTATE %d failed.", me.index))
	}
}

// Returns the zero-based index of the item.
func (me ListViewItem) Index() int {
	return int(me.index)
}

// Tells whether the item is currently selected.
func (me ListViewItem) IsSelected() bool {
	return co.LVIS(
		me.lv.Hwnd().SendMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(me.index), win.LPARAM(co.LVIS_SELECTED)),
	) == co.LVIS_SELECTED
}

// Tells whether the item is currently visible.
func (me ListViewItem) IsVisible() bool {
	return me.lv.Hwnd().SendMessage(co.LVM_ISITEMVISIBLE,
		win.WPARAM(me.index), 0) != 0
}

// Retrieves the custom data associated with the item.
func (me ListViewItem) LParam() win.LPARAM {
	lvi := win.LVITEM{
		IItem: int32(me.index),
		Mask:  co.LVIF_PARAM,
	}

	ret := me.lv.Hwnd().SendMessage(co.LVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETITEM %d failed.", me.index))
	}

	return lvi.LParam
}

// Retrieves the coordinates of the rectangle surrounding the item.
func (me ListViewItem) Rect(portion co.LVIR) win.RECT {
	rcItem := win.RECT{
		Left: int32(portion),
	}

	ret := me.lv.Hwnd().SendMessage(co.LVM_GETITEMRECT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&rcItem)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETITEMRECT %d failed.", me.index))
	}
	return rcItem // coordinates relative to the ListView
}

// Selects the item.
func (me ListViewItem) Select(isSelected bool) {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMSTATE %d failed.", me.index))
	}
}

// Sets the custom data associated with the item.
func (me ListViewItem) SetLParam(lp win.LPARAM) {
	lvi := win.LVITEM{
		IItem:  int32(me.index),
		Mask:   co.LVIF_PARAM,
		LParam: lp,
	}

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEM %d failed.", me.index))
	}
}

// Sets the text of the item.
func (me ListViewItem) SetText(columnIndex int, text string) {
	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
	}
	lvi.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETITEMTEXT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT %d/%d failed \"%s\".",
			me.index, columnIndex, text))
	}
}

// Retrieves the text of the item.
func (me ListViewItem) Text(columnIndex int) string {
	const BLOCK int = 64 // arbitrary
	bufSz := BLOCK
	var buf []uint16

	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
	}

	for {
		buf = make([]uint16, bufSz)
		lvi.SetPszText(buf)

		nChars := int(
			me.lv.Hwnd().SendMessage(co.LVM_GETITEMTEXT,
				win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvi))),
		)

		if nChars+1 < bufSz { // to break, must have at least 1 char gap
			break
		}

		bufSz += BLOCK // increase buffer size to try again
	}

	return win.Str.FromNativeSlice(buf)
}

// Sends an LVM_UPDATE message to the item.
func (me ListViewItem) Update() {
	if me.lv.Hwnd().SendMessage(co.LVM_UPDATE, win.WPARAM(me.index), 0) == 0 {
		panic(fmt.Sprintf("LVM_UPDATE %d failed.", me.index))
	}
}
