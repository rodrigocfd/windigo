//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
)

// An item from a [list view].
//
// [list view]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type ListViewItem struct {
	owner *ListView
	index int32
}

// Deletes the item with [LVM_DELETEITEM].
//
// Panics on error.
//
// [LVM_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-deleteitem
func (me ListViewItem) Delete() {
	ret, err := me.owner.hWnd.SendMessage(co.LVM_DELETEITEM, win.WPARAM(int32(me.index)), 0)
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM %d failed.", me.index))
	}
}

// Returns the user-custom data stored for this item, or nil if none.
//
// Example:
//
//	type Person struct {
//		Name string
//	}
//
//	var item ui.ListViewItem // initialized somewhere
//
//	item.SetData(&Person{Name: "foo"})
//
//	if person := item.Data().(*Person); person != nil {
//		println(person.Name)
//	}
func (me ListViewItem) Data() interface{} {
	if data, ok := me.owner.itemsData[me.Uid()]; ok {
		return data
	}
	return nil
}

// Makes sure the item is visible with [LVM_ENSUREVISIBLE], scrolling the
// list view if needed.
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [LVM_ENSUREVISIBLE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-ensurevisible
func (me ListViewItem) EnsureVisible() ListViewItem {
	if me.owner.View() == co.LV_VIEW_DETAILS {
		// In details view, LVM_ENSUREVISIBLE won't center the item vertically.
		// That's what we do here.
		rc, _ := me.owner.hWnd.GetClientRect()
		cyList := rc.Bottom // total height of the list view

		idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETTOPINDEX, 0, 0) // 1st visible item
		lvii := win.LVITEMINDEX{
			IItem: int32(idxRet),
		}
		rc = win.RECT{
			Left: int32(co.LVIR_BOUNDS),
		}

		ret, err := me.owner.hWnd.SendMessage(co.LVM_GETITEMINDEXRECT,
			win.WPARAM(unsafe.Pointer(&lvii)), win.LPARAM(unsafe.Pointer(&rc)))
		if err != nil || ret == 0 {
			panic(fmt.Sprintf("LVM_GETITEMINDEXRECT %d failed.", lvii.IItem))
		}
		cyItem := rc.Bottom - rc.Top // height of a single item
		xTop := rc.Top               // topmost X of 1st visible item

		lvii = win.LVITEMINDEX{
			IItem: me.index,
		}
		rc = win.RECT{}

		ret, err = me.owner.hWnd.SendMessage(co.LVM_GETITEMINDEXRECT,
			win.WPARAM(unsafe.Pointer(&lvii)), win.LPARAM(unsafe.Pointer(&rc)))
		if err != nil || ret == 0 {
			panic(fmt.Sprintf("LVM_GETITEMINDEXRECT %d failed.", lvii.IItem))
		}
		xUs := rc.Top // our current X

		if xUs < xTop || xUs > xTop+cyList { // if we're not visible
			me.owner.Scroll(0, int(xUs-xTop-cyList/2+cyItem*2))
		}

	} else {
		ret, err := me.owner.hWnd.SendMessage(co.LVM_ENSUREVISIBLE,
			win.WPARAM(int32(me.index)), win.LPARAM(1)) // always entirely visible
		if err != nil || ret == 0 {
			panic(fmt.Sprintf("LVM_ENSUREVISIBLE %d failed.", me.index))
		}
	}
	return me
}

// Sets the item as the focused one with [LVM_SETITEMSTATE].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETITEMSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setitemstate
func (me ListViewItem) Focus() ListViewItem {
	lvi := win.LVITEM{
		State:     co.LVIS_FOCUSED,
		StateMask: co.LVIS_FOCUSED,
	}

	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(int32(me.index)), win.LPARAM(unsafe.Pointer(&lvi)))
	if err != nil || int32(ret) == -1 {
		panic(fmt.Sprintf("LVM_SETITEMSTATE %d failed.", me.index))
	}

	return me
}

// Retrieves the zero-based icon index with [LVM_GETITEM].
//
// Panics on error.
//
// [LVM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitem
func (me ListViewItem) IconIndex() int {
	lvi := win.LVITEM{
		IItem: me.index,
		Mask:  co.LVIF_IMAGE,
	}

	ret, err := me.owner.hWnd.SendMessage(co.LVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 || err != nil {
		panic("LVM_GETITEM failed.")
	}

	return int(lvi.IImage)
}

// Returns the zero-based index of the item.
func (me ListViewItem) Index() int {
	return int(me.index)
}

// Tells whether the item is currently selected with [LVM_GETITEMSTATE].
//
// [LVM_GETITEMSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemstate
func (me ListViewItem) IsSelected() bool {
	lvisRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETITEMSTATE,
		win.WPARAM(int32(me.index)), win.LPARAM(co.LVIS_SELECTED))
	return co.LVIS(lvisRet) == co.LVIS_SELECTED
}

// Tells whether the item is currently visible with [LVM_ISITEMVISIBLE].
//
// [LVM_ISITEMVISIBLE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-isitemvisible
func (me ListViewItem) IsVisible() bool {
	ret, _ := me.owner.hWnd.SendMessage(co.LVM_ISITEMVISIBLE,
		win.WPARAM(int32(me.index)), 0)
	return ret != 0
}

// Retrieves the coordinates of the rectangle surrounding the item with
// [LVM_GETITEMRECT].
//
// Panics on error.
//
// [LVM_GETITEMRECT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemrect
func (me ListViewItem) ItemRect(portion co.LVIR) win.RECT {
	rcItem := win.RECT{
		Left: int32(portion),
	}

	ret, err := me.owner.hWnd.SendMessage(co.LVM_GETITEMRECT,
		win.WPARAM(int32(me.index)), win.LPARAM(unsafe.Pointer(&rcItem)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_GETITEMRECT %d failed.", me.index))
	}
	return rcItem // coordinates relative to the ListView
}

// Returns the next item, if any.
func (me ListViewItem) Next() (ListViewItem, bool) {
	count := me.owner.Items.Count()
	if me.index == int32(count)-1 { // we are the last one
		return ListViewItem{}, false
	}
	return ListViewItem{owner: me.owner, index: me.index + 1}, true
}

// Returns the previous item, if any.
func (me ListViewItem) Prev() (ListViewItem, bool) {
	if me.index == 0 { // we are the first one
		return ListViewItem{}, false
	}
	return ListViewItem{owner: me.owner, index: me.index - 1}, true
}

// Selects or deselects the item with [LVM_SETITEMSTATE].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETITEMSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setitemstate
func (me ListViewItem) Select(isSelected bool) ListViewItem {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}

	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(int32(me.index)), win.LPARAM(unsafe.Pointer(&lvi)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMSTATE %d failed.", me.index))
	}

	return me
}

// Stores user-custom data for this item.
//
// Example:
//
//	type Person struct {
//		Name string
//	}
//
//	var item ui.ListViewItem // initialized somewhere
//
//	item.SetData(&Person{Name: "foo"})
//
//	if person := item.Data().(*Person); person != nil {
//		println(person.Name)
//	}
func (me ListViewItem) SetData(data interface{}) {
	me.owner.itemsData[me.Uid()] = data
}

// Sets the zero-based icon index with [LVM_SETITEM].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setitem
func (me ListViewItem) SetIconIndex(iconIndex int) ListViewItem {
	lvi := win.LVITEM{
		IItem: me.index,
		Mask:  co.LVIF_IMAGE,
	}

	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 || err != nil {
		panic("LVM_SETITEM failed.")
	}

	return me
}

// Sets the text of the item at the given column with [LVM_SETITEMTEXT].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETITEMTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setitemtext
func (me ListViewItem) SetText(columnIndex int, text string) ListViewItem {
	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
	}

	var wText wstr.BufEncoder
	lvi.SetPszText(wText.Slice(text))

	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETITEMTEXT,
		win.WPARAM(int32(me.index)), win.LPARAM(unsafe.Pointer(&lvi)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT %d/%d failed \"%s\".",
			me.index, columnIndex, text))
	}

	return me
}

// Retrieves the text of the item, with [LVM_GETITEMTEXT].
//
// [LVM_GETITEMTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemtext
func (me ListViewItem) Text(columnIndex int) string {
	lvi := win.LVITEM{
		ISubItem: int32(columnIndex),
	}

	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	for {
		lvi.SetPszText(wBuf.HotSlice())

		nCharsRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETITEMTEXT,
			win.WPARAM(int32(me.index)), win.LPARAM(unsafe.Pointer(&lvi)))
		nChars := int(nCharsRet)

		if nChars+1 < wBuf.Len() { // to break, must have at least 1 char gap
			break
		}

		wBuf.AllocAndZero(wBuf.Len() + 64) // increase buffer size to try again
	}

	return wBuf.String()
}

// Returns the unique ID associated to the item with [LVM_MAPINDEXTOID].
//
// [LVM_MAPINDEXTOID]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-mapindextoid
func (me ListViewItem) Uid() int {
	uidRet, _ := me.owner.hWnd.SendMessage(co.LVM_MAPINDEXTOID, win.WPARAM(int32(me.index)), 0)
	return int(uidRet)
}

// Redraws the item immediately with [LVM_UPDATE].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [LVM_UPDATE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-update
func (me ListViewItem) Update() ListViewItem {
	ret, err := me.owner.hWnd.SendMessage(co.LVM_UPDATE, win.WPARAM(int32(me.index)), 0)
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_UPDATE %d failed.", me.index))
	}

	return me
}
