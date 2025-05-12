//go:build windows

package ui

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// The items collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [ListView].
type CollectionListViewItems struct {
	owner *ListView
}

// Adds one item with [LVM_INSERTITEM], then sets the texts under each
// subsequent column, returning the new item.
//
// Panics if no text is informed; panics on error.
//
// [LVM_INSERTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-insertitem
func (me *CollectionListViewItems) Add(texts ...string) ListViewItem {
	return me.AddWithIcon(-1, texts...)
}

// Returns all items.
func (me *CollectionListViewItems) All() []ListViewItem {
	nItems := me.Count()
	items := make([]ListViewItem, 0, nItems)
	for i := uint(0); i < nItems; i++ {
		items = append(items, me.Get(int(i)))
	}
	return items
}

// Adds one item with [LVM_INSERTITEM], then sets the texts under each
// subsequent column, returning the new item.
//
// The iconIndex is the zero-based index of the icon previously inserted into
// the control's image list, or -1 for no icon.
//
// Panics if no text is informed; panics on error.
//
// [LVM_INSERTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-insertitem
func (me *CollectionListViewItems) AddWithIcon(iconIndex int, texts ...string) ListViewItem {
	if len(texts) == 0 {
		panic("You must inform at least 1 text when adding a ListView item.")
	}

	text16 := wstr.NewBufWith[wstr.Stack20](texts[0], wstr.ALLOW_EMPTY)
	mask := co.LVIF_TEXT
	if iconIndex != -1 {
		mask |= co.LVIF_IMAGE
	}

	lvi := win.LVITEM{
		Mask:   mask,
		IItem:  0x0fff_ffff, // insert as last one
		IImage: int32(iconIndex),
	}
	lvi.SetPszText(text16.HotSlice()) // first column is inserted right away

	newIdxRet, err := me.owner.hWnd.SendMessage(co.LVM_INSERTITEM,
		0, win.LPARAM(unsafe.Pointer(&lvi)))
	newIdx := int(newIdxRet)
	if err != nil || newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM col %d, \"%s\" failed.", 0, texts[0]))
	}

	for i := 1; i < len(texts); i++ { // each subsequent column
		lvi.ISubItem = int32(i)
		text16.Set(texts[i], wstr.ALLOW_EMPTY)
		lvi.SetPszText(text16.HotSlice())

		ret, err := me.owner.hWnd.SendMessage(co.LVM_SETITEMTEXT,
			win.WPARAM(newIdx), win.LPARAM(unsafe.Pointer(&lvi)))
		if err != nil || ret == 0 {
			panic(fmt.Sprintf("LVM_SETITEMTEXT col %d, \"%s\" failed.", i, texts[i]))
		}
	}

	return ListViewItem{me.owner, int32(newIdx)}
}

// Retrieves the number of items with [LVM_GETITEMCOUNT].
//
// [LVM_GETITEMCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemcount
func (me *CollectionListViewItems) Count() uint {
	count, _ := me.owner.hWnd.SendMessage(co.LVM_GETITEMCOUNT, 0, 0)
	return uint(count)
}

// Deletes all items at once with [LVM_DELETEALLITEMS].
//
// [LVM_DELETEALLITEMS]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-deleteallitems
func (me *CollectionListViewItems) DeleteAll() {
	me.owner.hWnd.SendMessage(co.LVM_DELETEALLITEMS, 0, 0)
}

// Deletes all selected items at once by searching them with [LVM_GETNEXTITEM],
// then calling [LVM_DELETEITEM].
//
// Panics on error.
//
// [LVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
// [LVM_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-deleteitem
func (me *CollectionListViewItems) DeleteSelected() {
	for {
		idxBaseSearch := -1 // always search the first one
		idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(idxBaseSearch), win.LPARAM(co.LVNI_SELECTED))
		idx := int(idxRet)
		if idx == -1 {
			break
		}

		delRet, err := me.owner.hWnd.SendMessage(co.LVM_DELETEITEM, win.WPARAM(idx), 0)
		if err != nil || delRet == 0 {
			panic(fmt.Sprintf("LVM_DELETEITEM %d failed.", idx))
		}
	}
}

// Retrieves the focused item with [LVM_GETNEXTITEM], if any.
//
// [LVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
func (me *CollectionListViewItems) Focused() (ListViewItem, bool) {
	idxBaseSearch := -1
	idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETNEXTITEM,
		win.WPARAM(idxBaseSearch), win.LPARAM(co.LVNI_FOCUSED))
	idx := int(idxRet)
	if idx == -1 {
		return ListViewItem{}, false
	}

	return me.Get(idx), true
}

// Calls [LVM_FINDITEM] to search for the first item with the given exact text,
// case-insensitive.
//
// [LVM_FINDITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-finditem
func (me *CollectionListViewItems) Find(text string) (ListViewItem, bool) {
	text16 := wstr.NewBufWith[wstr.Stack20](text, wstr.ALLOW_EMPTY)
	lvfi := win.LVFINDINFO{
		Flags: co.LVFI_STRING,
		Psz:   text16.Ptr(),
	}

	idxBaseSearch := -1
	idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_FINDITEM,
		win.WPARAM(idxBaseSearch), win.LPARAM(unsafe.Pointer(&lvfi)))
	idx := int(idxRet)
	if idx == -1 {
		return ListViewItem{}, false // not found
	}

	return me.Get(idx), true
}

// Returns the item at the given index.
func (me *CollectionListViewItems) Get(index int) ListViewItem {
	return ListViewItem{me.owner, int32(index)}
}

// Calls [LVM_MAPIDTOINDEX] to return the item associated to the unique ID.
//
// [LVM_MAPIDTOINDEX]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-mapidtoindex
func (me *CollectionListViewItems) GetByUid(uid int) ListViewItem {
	idx, _ := me.owner.hWnd.SendMessage(co.LVM_MAPIDTOINDEX, win.WPARAM(uid), 0)
	return me.Get(int(idx))
}

// Retrieves the item below the given coordinates with [LVM_HITTEST], if any.
//
// The coordinates must be relative to the ListView.
//
// [LVM_HITTEST]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-hittest
func (me *CollectionListViewItems) HitTest(pos win.POINT) (ListViewItem, bool) {
	lvhti := win.LVHITTESTINFO{
		Pt: pos,
	}

	idxBaseSearch := -1 // Vista: retrieve iGroup and iSubItem
	me.owner.hWnd.SendMessage(co.LVM_HITTEST,
		win.WPARAM(idxBaseSearch), win.LPARAM(unsafe.Pointer(&lvhti)))

	if lvhti.IItem == -1 {
		return me.Get(-1), false
	}
	return me.Get(int(lvhti.IItem)), true
}

// Returns the last item.
func (me *CollectionListViewItems) Last() ListViewItem {
	return me.Get(int(me.Count()) - 1)
}

// Selects or deselects all items at once with [LVM_SETITEMSTATE].
//
// Panics on error.
//
// [LVM_SETITEMSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setitemstate
func (me *CollectionListViewItems) SelectAll(doSelect bool) {
	stylesRet, _ := me.owner.hWnd.GetWindowLongPtr(co.GWLP_STYLE)
	styles := co.LVS(stylesRet)
	if (styles & co.LVS_SINGLESEL) != 0 {
		return // single-sel list views cannot have all items selected
	}

	state := co.LVIS_NONE
	if doSelect {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}

	idxBaseSearch := -1
	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(idxBaseSearch), win.LPARAM(unsafe.Pointer(&lvi)))
	if err != nil || ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
}

// Returns the selected items with [LVM_GETNEXTITEM].
//
// [LVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
func (me *CollectionListViewItems) Selected() []ListViewItem {
	nSel := me.SelectedCount()
	items := make([]ListViewItem, 0, nSel)

	idx := -1
	for {
		idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED))
		idx = int(idxRet)
		if idx == -1 {
			break
		}
		items = append(items, me.Get(idx))
	}

	return items
}

// Retrieves the number of selected items with [LVM_GETSELECTEDCOUNT].
//
// [LVM_GETSELECTEDCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getselectedcount
func (me *CollectionListViewItems) SelectedCount() uint {
	ret, _ := me.owner.hWnd.SendMessage(co.LVM_GETSELECTEDCOUNT, 0, 0)
	return uint(ret)
}

// Sorts the items according to the callback with [LVM_SORTITEMSEX].
//
// # Example
//
//	var lv ui.ListView // initialized somewhere
//
//	lv.Items.Sort(func(itemA, itemB ui.ListViewItem) int {
//		return win.Str.Cmp(itemA.Text(0), itemB.Text(0))
//	})
//
// [LVM_SORTITEMSEX]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-sortitemsex
func (me *CollectionListViewItems) Sort(fun func(a, b ListViewItem) int) {
	listViewSortCallback()
	pPack := &_ListViewSortPack{lv: me.owner, f: fun}
	me.owner.hWnd.SendMessage(co.LVM_SORTITEMSEX,
		win.WPARAM(unsafe.Pointer(pPack)), win.LPARAM(listViewSortCallback()))
	runtime.KeepAlive(pPack)
}

type _ListViewSortPack struct {
	lv *ListView
	f  func(a, b ListViewItem) int
}

var _listViewSortCallback uintptr

func listViewSortCallback() uintptr {
	if _listViewSortCallback == 0 {
		_listViewSortCallback = syscall.NewCallback(
			func(idxA, idxB, lParam uintptr) uintptr {
				pPack := (*_ListViewSortPack)(unsafe.Pointer(lParam))
				itemA := pPack.lv.Items.Get(int(idxA))
				itemB := pPack.lv.Items.Get(int(idxB))
				return uintptr(pPack.f(itemA, itemB))
			},
		)
	}
	return _listViewSortCallback
}

// Retrieves the topmost visible item with [LVM_GETTOPINDEX], if any.
//
// [LVM_GETTOPINDEX]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-gettopindex
func (me *CollectionListViewItems) TopmostVisible() (ListViewItem, bool) {
	idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETTOPINDEX, 0, 0)
	idx := int(idxRet)
	if idx == -1 {
		return me.Get(-1), false
	}
	return me.Get(idx), true
}
