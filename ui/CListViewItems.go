package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ListViewItems struct {
	lv ListView
}

func (me *_ListViewItems) new(ctrl ListView) {
	me.lv = ctrl
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
		me.lv.Hwnd().SendMessage(co.LVM_INSERTITEM,
			0, win.LPARAM(unsafe.Pointer(&lvi))),
	)
	if newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM col %d, \"%s\" failed.", 0, texts[0]))
	}

	for i := 1; i < len(texts); i++ { // each subsequent column
		lvi.ISubItem = int32(i)
		lvi.SetPszText(win.Str.ToNativeSlice(texts[i]))

		ret := me.lv.Hwnd().SendMessage(co.LVM_SETITEMTEXT,
			win.WPARAM(newIdx), win.LPARAM(unsafe.Pointer(&lvi)))
		if ret == 0 {
			panic(fmt.Sprintf("LVM_SETITEMTEXT col %d, \"%s\" failed.", i, texts[i]))
		}
	}

	return me.Get(newIdx)
}

// Retrieves all the items.
func (me *_ListViewItems) All() []ListViewItem {
	numItems := me.Count()
	items := make([]ListViewItem, 0, numItems)
	for i := 0; i < numItems; i++ {
		items = append(items, me.Get(i))
	}
	return items
}

// Retrieves the number of items.
func (me *_ListViewItems) Count() int {
	return int(me.lv.Hwnd().SendMessage(co.LVM_GETITEMCOUNT, 0, 0))
}

// Deletes all items at once.
func (me *_ListViewItems) DeleteAll() {
	me.lv.Hwnd().SendMessage(co.LVM_DELETEALLITEMS, 0, 0)
}

// Deletes all selected items at once.
func (me *_ListViewItems) DeleteSelected() {
	for {
		idx := -1 // always search the first one
		idx = int(
			me.lv.Hwnd().SendMessage(co.LVM_GETNEXTITEM,
				win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED)),
		)
		if idx == -1 {
			break
		}

		if me.lv.Hwnd().SendMessage(co.LVM_DELETEITEM, win.WPARAM(idx), 0) == 0 {
			panic(fmt.Sprintf("LVM_DELETEITEM %d failed.", idx))
		}
	}
}

// Retrieves the focused item, if any.
func (me *_ListViewItems) Focused() (ListViewItem, bool) {
	startIdx := -1
	idx := int(
		me.lv.Hwnd().SendMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(startIdx), win.LPARAM(co.LVNI_FOCUSED)),
	)
	if idx == -1 {
		return nil, false
	}

	return me.Get(idx), true
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
		me.lv.Hwnd().SendMessage(co.LVM_FINDITEM,
			win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvfi))),
	)
	if idx == -1 {
		return nil, false // not found
	}

	return me.Get(idx), true
}

// Returns the item at the given index.
//
// Note that this method is dumb: no validation is made, the given index is
// simply kept. If the index is invalid (or becomes invalid), subsequent
// operations on the ListViewItem will fail.
func (me *_ListViewItems) Get(index int) ListViewItem {
	item := &_ListViewItem{}
	item.new(me.lv, index)
	return item
}

// Retrieves the item below the given coordinates, if any.
//
// The coordinates must be relative to the ListView.
func (me *_ListViewItems) HitTest(pos win.POINT) (ListViewItem, bool) {
	lvhti := win.LVHITTESTINFO{
		Pt: pos,
	}

	wp := -1 // Vista: retrieve iGroup and iSubItem
	me.lv.Hwnd().SendMessage(co.LVM_HITTEST,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvhti)))

	if lvhti.IItem == -1 {
		return nil, false
	}
	return me.Get(int(lvhti.IItem)), true
}

// Selects or deselects all items at once.
func (me *_ListViewItems) SelectAll(doSelect bool) {
	lvi := win.LVITEM{
		State:     util.Iif(doSelect, co.LVIS_SELECTED, co.LVIS_NONE).(co.LVIS),
		StateMask: co.LVIS_SELECTED,
	}

	idx := -1
	ret := me.lv.Hwnd().SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(idx), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
}

// Retrieves the selected items.
func (me *_ListViewItems) Selected() []ListViewItem {
	items := make([]ListViewItem, 0, me.SelectedCount())

	idx := -1
	for {
		idx = int(
			me.lv.Hwnd().SendMessage(co.LVM_GETNEXTITEM,
				win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED)),
		)
		if idx == -1 {
			break
		}
		items = append(items, me.Get(idx))
	}

	return items
}

// Retrieves the number of selected items.
func (me *_ListViewItems) SelectedCount() int {
	return int(me.lv.Hwnd().SendMessage(co.LVM_GETSELECTEDCOUNT, 0, 0))
}

// Retrieves the topmost visible item, if any.
func (me *_ListViewItems) TopmostVisible() (ListViewItem, bool) {
	idx := int(me.lv.Hwnd().SendMessage(co.LVM_GETTOPINDEX, 0, 0))
	if idx == -1 {
		return nil, false
	}

	return me.Get(idx), true
}
