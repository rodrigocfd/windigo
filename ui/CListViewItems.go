package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ListViewItems struct {
	pHwnd *win.HWND
}

func (me *_ListViewItems) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
}

// Adds an item, specifying the texts under each column, returning its index.
func (me *_ListViewItems) Add(texts ...string) int {
	return me.AddWithIcon(-1, texts...)
}

// Adds an item, specifying its icon and the texts under each column, returning its index.
func (me *_ListViewItems) AddWithIcon(iconIndex int, texts ...string) int {
	textBuf := win.Str.ToUint16Slice(texts[0])
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT | co.LVIF_IMAGE,
		IItem:   0x0fff_ffff, // insert as last one
		PszText: &textBuf[0],
		IImage:  int32(iconIndex),
	}

	newIdx := int(
		me.pHwnd.SendMessage(co.LVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&lvi))),
	)
	if newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM \"%s\" failed.", texts[0]))
	}

	for i, text := range texts {
		if i > 0 {
			textBuf = win.Str.ToUint16Slice(text)
			lvi.ISubItem = int32(i)
			lvi.PszText = &textBuf[0]

			ret := me.pHwnd.SendMessage(co.LVM_SETITEMTEXT,
				win.WPARAM(newIdx), win.LPARAM(unsafe.Pointer(&lvi)))
			if ret == 0 {
				panic(fmt.Sprintf("LVM_SETITEMTEXT \"%s\" failed.", text))
			}
		}
	}

	return newIdx
}

// Retrieves the number of items.
func (me *_ListViewItems) Count() int {
	return int(me.pHwnd.SendMessage(co.LVM_GETITEMCOUNT, 0, 0))
}

// Deletes the items at the given indexes.
func (me *_ListViewItems) Delete(itemIndexes ...int) {
	for i := len(itemIndexes) - 1; i >= 0; i-- {
		ret := me.pHwnd.SendMessage(co.LVM_DELETEITEM,
			win.WPARAM(itemIndexes[i]), 0)
		if ret == 0 {
			panic(fmt.Sprintf("LVM_DELETEITEM %d failed.", itemIndexes[i]))
		}
	}
}

// Deletes all items at once.
func (me *_ListViewItems) DeleteAll() {
	me.pHwnd.SendMessage(co.LVM_DELETEALLITEMS, 0, 0)
}

// Scrolls the list view so the given item is visible
func (me *_ListViewItems) EnsureVisible(itemIndex int) {
	ret := me.pHwnd.SendMessage(co.LVM_ENSUREVISIBLE,
		win.WPARAM(itemIndex), win.LPARAM(1)) // always entirely visible
	if ret == 0 {
		panic(fmt.Sprintf("LVM_ENSUREVISIBLE %d failed.", itemIndex))
	}
}

// Retrieves the index of the focused item, if any.
func (me *_ListViewItems) Focused() (int, bool) {
	startIdx := -1
	idx := int(
		me.pHwnd.SendMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(startIdx), win.LPARAM(co.LVNI_FOCUSED)),
	)
	if idx == -1 {
		return -1, false
	}
	return idx, true
}

// Searches for an item with the given exact text, case-insensitive.
func (me *_ListViewItems) Find(text string) (int, bool) {
	textBuf := win.Str.ToUint16Slice(text)
	lvfi := win.LVFINDINFO{
		Flags: co.LVFI_STRING,
		Psz:   &textBuf[0],
	}

	wp := -1
	idx := int(
		me.pHwnd.SendMessage(co.LVM_FINDITEM,
			win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvfi))),
	)

	if idx == -1 {
		return -1, false // not found
	}
	return idx, true
}

// Sends LVM_HITTEST to determine the item at specified position, if any. Pos
// coordinates must be relative to list view.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-hittest
func (me *_ListViewItems) HitTest(pos win.POINT) *win.LVHITTESTINFO {
	lvhti := win.LVHITTESTINFO{
		Pt: pos,
	}

	wp := -1 // Vista: retrieve iGroup and iSubItem
	me.pHwnd.SendMessage(co.LVM_HITTEST,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvhti)))
	return &lvhti
}

// Retrieves information about an item.
func (me *_ListViewItems) Info(lvi *win.LVITEM) {
	ret := me.pHwnd.SendMessage(co.LVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETITEM %d failed.", lvi.IItem))
	}
}

// Retrieves whether the item is currently selected.
func (me *_ListViewItems) IsSelected(itemIndex int) bool {
	return co.LVIS(
		me.pHwnd.SendMessage(co.LVM_GETITEMSTATE,
			win.WPARAM(itemIndex), win.LPARAM(co.LVIS_SELECTED)),
	) == co.LVIS_SELECTED
}

// Retrieves whether the given item is visible.
func (me *_ListViewItems) IsVisible(itemIndex int) bool {
	return me.pHwnd.SendMessage(co.LVM_ISITEMVISIBLE, win.WPARAM(itemIndex), 0) != 0
}

// Retrieves the LPARAM associated to the item.
func (me *_ListViewItems) LParam(itemIndex int) win.LPARAM {
	lvi := win.LVITEM{
		IItem: int32(itemIndex),
		Mask:  co.LVIF_PARAM,
	}

	me.Info(&lvi)
	return lvi.LParam
}

// Retrieves bound coordinates of the item with LVM_GETITEMRECT. Coordinates are
// relative to list view.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemrect
func (me *_ListViewItems) Rect(itemIndex int, portion co.LVIR) win.RECT {
	rcItem := win.RECT{
		Left: int32(portion),
	}

	ret := me.pHwnd.SendMessage(co.LVM_GETITEMRECT,
		win.WPARAM(itemIndex), win.LPARAM(unsafe.Pointer(&rcItem)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETITEMRECT %d failed.", itemIndex))
	}
	return rcItem
}

// Retrieves the indexes of the selected items.
func (me *_ListViewItems) Selected() []int {
	indexes := make([]int, 0, me.SelectedCount())

	i := -1
	for {
		i = int(
			me.pHwnd.SendMessage(co.LVM_GETNEXTITEM,
				win.WPARAM(i), win.LPARAM(co.LVNI_SELECTED)),
		)
		if i == -1 {
			break
		}
		indexes = append(indexes, i)
	}

	return indexes
}

// Retrieves the number of selected items.
func (me *_ListViewItems) SelectedCount() int {
	return int(me.pHwnd.SendMessage(co.LVM_GETSELECTEDCOUNT, 0, 0))
}

// Sets the item as the currently focused one.
func (me *_ListViewItems) SetFocused(itemIndex int) {
	lvi := win.LVITEM{
		State:     co.LVIS_FOCUSED,
		StateMask: co.LVIS_FOCUSED,
	}

	ret := me.pHwnd.SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(itemIndex), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMSTATE %d failed.", itemIndex))
	}
}

// Sets information about an item.
func (me *_ListViewItems) SetInfo(lvi *win.LVITEM) {
	ret := me.pHwnd.SendMessage(co.LVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEM %d failed.", lvi.IItem))
	}
}

// Sets the LPARAM associated to the item.
func (me *_ListViewItems) SetLParam(itemIndex int, lParam win.LPARAM) {
	lvi := win.LVITEM{
		IItem:  int32(itemIndex),
		Mask:   co.LVIF_PARAM,
		LParam: lParam,
	}

	me.SetInfo(&lvi)
}

// Selects or deselects one or more items.
func (me *_ListViewItems) SetSelected(isSelected bool, itemIndexes ...int) {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}

	for _, index := range itemIndexes {
		ret := me.pHwnd.SendMessage(co.LVM_SETITEMSTATE,
			win.WPARAM(index), win.LPARAM(unsafe.Pointer(&lvi)))
		if ret == 0 {
			panic(fmt.Sprintf("LVM_SETITEMSTATE %d failed.", index))
		}
	}
}

// Selects or deselects all items at once.
func (me *_ListViewItems) SetSelectedAll(isSelected bool) {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}

	idx := -1
	ret := me.pHwnd.SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(idx), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
}

// Sets the text at the given column.
func (me *_ListViewItems) SetText(itemIndex, columnIndex int, text string) {
	textBuf := win.Str.ToUint16Slice(text)
	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
		PszText:  &textBuf[0],
	}

	ret := me.pHwnd.SendMessage(co.LVM_SETITEMTEXT,
		win.WPARAM(itemIndex), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT %d/%d failed \"%s\".",
			itemIndex, columnIndex, text))
	}
}

// Retrieves the text at the given column.
func (me *_ListViewItems) Text(itemIndex, columnIndex int) string {
	buf := [256]uint16{} // arbitrary
	lvi := win.LVITEM{
		ISubItem:   int32(columnIndex),
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}

	ret := me.pHwnd.SendMessage(co.LVM_GETITEMTEXT,
		win.WPARAM(itemIndex), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic(fmt.Sprintf("LVM_GETITEMTEXT %d/%d failed.",
			itemIndex, columnIndex))
	}
	return win.Str.FromUint16Slice(buf[:])
}

// Retrieves the index of the topmost visible item.
func (me *_ListViewItems) TopmostVisible() int {
	return int(me.pHwnd.SendMessage(co.LVM_GETTOPINDEX, 0, 0))
}
