package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ListViewItems struct {
	pHwnd *win.HWND
}

func (me *_ListViewItems) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
}

func (me *_ListViewItems) getUnchecked(index int) ListViewItem {
	item := _ListViewItem{}
	item.new(me.pHwnd, index)
	return &item
}

// Adds an item, specifying the texts under each column, returning the new item.
func (me *_ListViewItems) Add(texts ...string) ListViewItem {
	return me.AddWithIcon(-1, texts...)
}

// Adds an item, specifying its icon and the texts under each column, returning the new item.
func (me *_ListViewItems) AddWithIcon(iconIndex int, texts ...string) ListViewItem {
	lvi := win.LVITEM{}
	lvi.Mask = co.LVIF_TEXT | co.LVIF_IMAGE
	lvi.IItem = 0x0fff_ffff // insert as last one
	lvi.IImage = int32(iconIndex)
	lvi.SetPszText(win.Str.ToNativeSlice(texts[0])) // first column is inserted right away

	newIdx := int(
		me.pHwnd.SendMessage(co.LVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&lvi))),
	)
	if newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM col %d, \"%s\" failed.", 0, texts[0]))
	}

	for i := 1; i < len(texts); i++ { // each subsequent column
		lvi.ISubItem = int32(i)
		lvi.SetPszText(win.Str.ToNativeSlice(texts[i]))

		ret := me.pHwnd.SendMessage(co.LVM_SETITEMTEXT,
			win.WPARAM(newIdx), win.LPARAM(unsafe.Pointer(&lvi)))
		if ret == 0 {
			panic(fmt.Sprintf("LVM_SETITEMTEXT col %d, \"%s\" failed.", i, texts[i]))
		}
	}

	return me.getUnchecked(newIdx)
}

// Retrieves all the items.
func (me *_ListViewItems) All() []ListViewItem {
	numItems := me.Count()
	items := make([]ListViewItem, 0, numItems)
	for i := 0; i < numItems; i++ {
		items = append(items, me.getUnchecked(i))
	}
	return items
}

// Retrieves the number of items.
func (me *_ListViewItems) Count() int {
	return int(me.pHwnd.SendMessage(co.LVM_GETITEMCOUNT, 0, 0))
}

// Deletes all items at once.
func (me *_ListViewItems) DeleteAll() {
	me.pHwnd.SendMessage(co.LVM_DELETEALLITEMS, 0, 0)
}

// Deletes all selected items at once.
func (me *_ListViewItems) DeleteSelected() {
	for {
		idx := -1 // always search the first one
		idx = int(
			me.pHwnd.SendMessage(co.LVM_GETNEXTITEM,
				win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED)),
		)
		if idx == -1 {
			break
		}

		if me.pHwnd.SendMessage(co.LVM_DELETEITEM, win.WPARAM(idx), 0) == 0 {
			panic(fmt.Sprintf("LVM_DELETEITEM %d failed.", idx))
		}
	}
}

// Retrieves the focused item, if any.
func (me *_ListViewItems) Focused() (ListViewItem, bool) {
	startIdx := -1
	idx := int(
		me.pHwnd.SendMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(startIdx), win.LPARAM(co.LVNI_FOCUSED)),
	)
	if idx == -1 {
		return nil, false
	}

	return me.getUnchecked(idx), true
}

// Sends LVM_FINDITEM to search for an item with the given exact text,
// case-insensitive.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/lvm-finditem
func (me *_ListViewItems) Find(text string) (ListViewItem, bool) {
	lvfi := win.LVFINDINFO{
		Flags: co.LVFI_STRING,
		Psz:   win.Str.ToNativePtr(text),
	}

	wp := -1
	idx := int(
		me.pHwnd.SendMessage(co.LVM_FINDITEM,
			win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvfi))),
	)
	if idx == -1 {
		return nil, false // not found
	}

	return me.getUnchecked(idx), true
}

// Returns the item at the given index; panics if the item does not exist.
func (me *_ListViewItems) Get(index int) ListViewItem {
	if index < 0 || index >= me.Count() {
		panic(fmt.Sprintf("Trying to retrieve item %d, there are %d.",
			index, me.Count()))
	}
	return me.getUnchecked(index)
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

// Retrieves the selected items.
func (me *_ListViewItems) Selected() []ListViewItem {
	items := make([]ListViewItem, 0, me.SelectedCount())

	idx := -1
	for {
		idx = int(
			me.pHwnd.SendMessage(co.LVM_GETNEXTITEM,
				win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED)),
		)
		if idx == -1 {
			break
		}
		items = append(items, me.getUnchecked(idx))
	}

	return items
}

// Retrieves the number of selected items.
func (me *_ListViewItems) SelectedCount() int {
	return int(me.pHwnd.SendMessage(co.LVM_GETSELECTEDCOUNT, 0, 0))
}

// Selects or deselects all items at once.
func (me *_ListViewItems) SetSelectedAll(doSelect bool) {
	lvi := win.LVITEM{
		State:     util.Iif(doSelect, co.LVIS_SELECTED, co.LVIS_NONE).(co.LVIS),
		StateMask: co.LVIS_SELECTED,
	}

	idx := -1
	ret := me.pHwnd.SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(idx), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
}

// Retrieves the topmost visible item, if any.
func (me *_ListViewItems) TopmostVisible() (ListViewItem, bool) {
	idx := int(me.pHwnd.SendMessage(co.LVM_GETTOPINDEX, 0, 0))
	if idx == -1 {
		return nil, false
	}

	return me.getUnchecked(idx), true
}
