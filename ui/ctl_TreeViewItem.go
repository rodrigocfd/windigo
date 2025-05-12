//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// An item from a [tree view].
//
// [tree view]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-controls
type TreeViewItem struct {
	owner *TreeView
	hItem win.HTREEITEM
}

// Adds a child to this item with [TVM_INSERTITEM], returning the new item.
//
// The iconIndex is the zero-based index of the icon previously inserted into
// the control's image list, or -1 for no icon.
//
// Panics on error.
//
// [TVM_INSERTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-insertitem
func (me TreeViewItem) AddChild(text string, iconIndex int) TreeViewItem {
	text16 := wstr.NewBufWith[wstr.Stack20](text, wstr.ALLOW_EMPTY)
	mask := co.TVIF_TEXT
	if iconIndex != -1 {
		mask |= co.TVIF_IMAGE
	}

	tvi := win.TVINSERTSTRUCT{
		HParent:      me.hItem,
		HInsertAfter: win.HTREEITEM_LAST,
		Itemex: win.TVITEMEX{
			Mask:   mask,
			IImage: int32(iconIndex),
		},
	}
	tvi.Itemex.SetPszText(text16.HotSlice())

	hItemRet, err := me.owner.hWnd.SendMessage(co.TVM_INSERTITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if hItemRet == 0 || err != nil {
		panic(fmt.Sprintf("TVM_INSERTITEM \"%s\" failed.", text))
	}
	hItem := win.HTREEITEM(hItemRet)

	return TreeViewItem{me.owner, hItem}
}

// Returns the child items with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me TreeViewItem) Children() []TreeViewItem {
	hItem, _ := me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_CHILD), win.LPARAM(me.hItem)) // retrieve first child
	hasSibling := hItem != 0 // has first child?

	items := make([]TreeViewItem, 0)

	for hasSibling {
		items = append(items, me.owner.Items.Get(win.HTREEITEM(hItem)))

		hItem, _ = me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
			win.WPARAM(co.TVGN_NEXT), win.LPARAM(hItem)) // retrieve next siblings
		hasSibling = hItem != 0
	}

	return items
}

// Returns the user-custom data stored for this item, or nil if none.
//
// # Example
//
//	type Person struct {
//		Name string
//	}
//
//	var item ui.TreeViewItem // initialized somewhere
//
//	item.SetData(&Person{Name: "foo"})
//
//	if person := item.Data().(*Person); person != nil {
//		println(person.Name)
//	}
func (me TreeViewItem) Data() interface{} {
	if data, ok := me.owner.itemsData[me.hItem]; ok {
		return data
	}
	return nil
}

// Deletes the item and all its children with [TVM_DELETEITEM].
//
// Panics on error.
//
// [TVM_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-deleteitem
func (me TreeViewItem) Delete() {
	ret, _ := me.owner.hWnd.SendMessage(co.TVM_DELETEITEM, 0, win.LPARAM(me.hItem))
	if ret == 0 {
		panic("TVM_DELETEITEM failed.")
	}
}

// Makes sure the item is visible with [TVM_ENSUREVISIBLE], scrolling the
// control if needed.
//
// Returns the same item, so further operations can be chained.
//
// [TVM_ENSUREVISIBLE]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-ensurevisible
func (me TreeViewItem) EnsureVisible() TreeViewItem {
	me.owner.hWnd.SendMessage(co.TVM_ENSUREVISIBLE, 0, win.LPARAM(me.hItem))
	return me
}

// Expands the item with [TVM_EXPAND].
//
// Returns the same item, so further operations can be chained.
//
// [TVM_EXPAND]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-expand
func (me TreeViewItem) Expand(doExpand bool) TreeViewItem {
	flag := co.TVE_COLLAPSE
	if doExpand {
		flag = co.TVE_EXPAND
	}
	me.owner.hWnd.SendMessage(co.TVM_EXPAND, win.WPARAM(flag), win.LPARAM(me.hItem))
	return me
}

// Returns the unique handle that identifies item.
func (me TreeViewItem) Htreeitem() win.HTREEITEM {
	return me.hItem
}

// Retrieves the zero-based icon index with [TVM_GETITEM].
//
// Panics on error.
//
// [TVM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getitem
func (me TreeViewItem) IconIndex() int {
	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_IMAGE,
	}

	ret, err := me.owner.hWnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 || err != nil {
		panic("TVM_GETITEM failed.")
	}

	return int(tvi.IImage)
}

// Returns true if the item is currently expanded, with [TVM_GETITEMSTATE].
//
// [TVM_GETITEMSTATE]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getitemstate
func (me TreeViewItem) IsExpanded() bool {
	tvis, _ := me.owner.hWnd.SendMessage(co.TVM_GETITEMSTATE,
		win.WPARAM(me.hItem), win.LPARAM(co.TVIS_EXPANDED))
	return (co.TVIS(tvis) & co.TVIS_EXPANDED) != 0
}

// Returns true if the item has no parent.
func (me TreeViewItem) IsRoot() bool {
	_, hasParent := me.Parent()
	return !hasParent
}

// Retrieves the next sibling item, if any, with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me TreeViewItem) NextSibling() (TreeViewItem, bool) {
	hSibling, _ := me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_NEXT), win.LPARAM(me.hItem))
	if hSibling != 0 {
		return TreeViewItem{me.owner, win.HTREEITEM(hSibling)}, true
	}
	return TreeViewItem{}, false
}

// Retrieves the parent item, if any, with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me TreeViewItem) Parent() (TreeViewItem, bool) {
	hParent, _ := me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_PARENT), win.LPARAM(me.hItem))
	if hParent != 0 {
		return TreeViewItem{me.owner, win.HTREEITEM(hParent)}, true
	}
	return TreeViewItem{}, false
}

// Retrieves the previous item, if any, with [TVM_GETNEXTITEM].
//
// [TVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getnextitem
func (me TreeViewItem) PrevSibling() (TreeViewItem, bool) {
	hSibling, _ := me.owner.hWnd.SendMessage(co.TVM_GETNEXTITEM,
		win.WPARAM(co.TVGN_PREVIOUS), win.LPARAM(me.hItem))
	if hSibling != 0 {
		return TreeViewItem{me.owner, win.HTREEITEM(hSibling)}, true
	}
	return TreeViewItem{}, false
}

// Stores user-custom data for this item.
//
// # Example
//
//	type Person struct {
//		Name string
//	}
//
//	var item ui.TreeViewItem // initialized somewhere
//
//	item.SetData(&Person{Name: "foo"})
//
//	if person := item.Data().(*Person); person != nil {
//		println(person.Name)
//	}
func (me TreeViewItem) SetData(data interface{}) {
	me.owner.itemsData[me.hItem] = data
}

// Sets the zero-based icon index with [TVM_SETITEM].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [TVM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-setitem
func (me TreeViewItem) SetIconIndex(iconIndex int) TreeViewItem {
	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_IMAGE,
	}

	ret, err := me.owner.hWnd.SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 || err != nil {
		panic("TVM_SETITEM failed.")
	}

	return me
}

// Sets the text of the item with [TVM_SETITEM].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [TVM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-setitem
func (me TreeViewItem) SetText(text string) TreeViewItem {
	text16 := wstr.NewBufWith[wstr.Stack20](text, wstr.ALLOW_EMPTY)
	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_TEXT,
	}
	tvi.SetPszText(text16.HotSlice())

	ret, err := me.owner.hWnd.SendMessage(co.TVM_SETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 || err != nil {
		panic(fmt.Sprintf("TVM_SETITEM failed \"%s\".", text))
	}
	return me
}

// Retrieves the text of the item with [TVM_GETITEM].
//
// Panics on error.
//
// [TVM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tvm-getitem
func (me TreeViewItem) Text() string {
	var buf [utl.MAX_PATH]uint16 // arbitrary

	tvi := win.TVITEMEX{
		HItem: me.hItem,
		Mask:  co.TVIF_TEXT,
	}
	tvi.SetPszText(buf[:])

	ret, err := me.owner.hWnd.SendMessage(co.TVM_GETITEM,
		0, win.LPARAM(unsafe.Pointer(&tvi)))
	if ret == 0 || err != nil {
		panic("TVM_GETITEM failed.")
	}
	return wstr.WstrSliceToStr(buf[:])
}
