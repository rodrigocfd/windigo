//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// The items collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [Header].
type CollectionHeaderItems struct {
	owner *Header
}

// Adds a new item with its width, using [HDM_INSERTITEM], and returns the new
// item.
//
// Panics on error.
//
// [HDM_INSERTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-insertitem
func (me *CollectionHeaderItems) Add(text string, width int) HeaderItem {
	hdi := win.HDITEM{
		Mask: co.HDI_TEXT | co.HDI_WIDTH,
		Cxy:  int32(width),
	}

	var wText wstr.BufEncoder
	hdi.SetPszText(wText.Slice(text))

	newIdxRet, err := me.owner.hWnd.SendMessage(co.HDM_INSERTITEM,
		0xffff, win.LPARAM(unsafe.Pointer(&hdi)))
	newIdx := int(newIdxRet)
	if err != nil || newIdx == -1 {
		panic(fmt.Sprintf("HDM_INSERTITEM \"%s\" failed.", text))
	}

	return me.Get(newIdx)
}

// Returns all items.
func (me *CollectionHeaderItems) All() []HeaderItem {
	nItems := me.Count()
	items := make([]HeaderItem, 0, nItems)
	for i := uint(0); i < nItems; i++ {
		items = append(items, me.Get(int(i)))
	}
	return items
}

// Sends [HDM_GETORDERARRAY] to retrieve the items in the current order.
//
// [HDM_GETORDERARRAY]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getorderarray
func (me *CollectionHeaderItems) AllOrdered() []HeaderItem {
	nItems := me.Count()
	indexes := make([]int32, nItems)

	me.owner.hWnd.SendMessage(co.HDM_GETORDERARRAY,
		win.WPARAM(nItems), win.LPARAM(unsafe.Pointer(&indexes[0])))

	items := make([]HeaderItem, 0, nItems)
	for _, index := range indexes {
		items = append(items, me.Get(int(index)))
	}
	return items
}

// Retrieves the number of items with [HDM_GETITEMCOUNT].
//
// Panics on error.
//
// [HDM_GETITEMCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitemcount
func (me *CollectionHeaderItems) Count() uint {
	countRet, err := me.owner.hWnd.SendMessage(co.HDM_GETITEMCOUNT, 0, 0)
	count := int(countRet)
	if err != nil || count == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint(count)
}

// Returns the item at the given index.
func (me *CollectionHeaderItems) Get(index int) HeaderItem {
	return HeaderItem{
		owner: me.owner,
		index: int32(index),
	}
}

// Sends [HDM_ORDERTOINDEX] to retrieve the item at the given order.
//
// [HDM_ORDERTOINDEX]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-ordertoindex
func (me *CollectionHeaderItems) GetByOrder(order uint) HeaderItem {
	idx, _ := me.owner.hWnd.SendMessage(co.HDM_ORDERTOINDEX, win.WPARAM(order), 0)
	return me.Get(int(idx))
}

// Returns the last item.
func (me *CollectionHeaderItems) Last() HeaderItem {
	return me.Get(int(me.Count()) - 1)
}

// Sends a [HDM_SETORDERARRAY] to reorder the items with the given order.
//
// [HDM_SETORDERARRAY]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setorderarray
func (me *CollectionHeaderItems) Reorder(indexes []int) {
	buf := make([]int32, 0, len(indexes))
	for _, index := range indexes {
		buf = append(buf, int32(index))
	}

	me.owner.hWnd.SendMessage(co.HDM_SETORDERARRAY,
		win.WPARAM(len(buf)), win.LPARAM(unsafe.Pointer(&buf[0])))
}
