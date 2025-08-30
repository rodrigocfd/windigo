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
// by the owning [Tab].
type CollectionTabItems struct {
	owner *Tab
}

func (me *CollectionTabItems) add(title string) TabItem {
	tci := win.TCITEM{
		Mask: co.TCIF_TEXT,
	}

	var wBuf wstr.BufEncoder
	tci.SetPszText(wBuf.Slice(title))

	newIdxRet, err := me.owner.hWnd.SendMessage(co.TCM_INSERTITEM,
		0x0fff_ffff, win.LPARAM(unsafe.Pointer(&tci)))
	newIdx := int(newIdxRet)
	if err != nil || newIdx == -1 {
		panic(fmt.Sprintf("TCM_INSERTITEM \"%s\" failed.", title))
	}

	return me.Get(newIdx)
}

// Retrieves the number of items with [TCM_GETITEMCOUNT].
//
// Panics on error.
//
// [TCM_GETITEMCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getitemcount
func (me *CollectionTabItems) Count() uint {
	countRet, err := me.owner.hWnd.SendMessage(co.TCM_GETITEMCOUNT, 0, 0)
	count := int(countRet)
	if err != nil || count == -1 {
		panic("TCM_GETITEMCOUNT failed.")
	}
	return uint(count)
}

// Retrieves the focused item with [TCM_GETCURFOCUS], if any
//
// [TCM_GETCURFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getcurfocus
func (me *CollectionTabItems) Focused() (TabItem, bool) {
	idxRet, _ := me.owner.hWnd.SendMessage(co.TCM_GETCURFOCUS, 0, 0)
	idx := int(idxRet)
	if idx == -1 {
		return TabItem{}, false
	}
	return me.Get(idx), true
}

// Returns the item at the given index.
func (me *CollectionTabItems) Get(index int) TabItem {
	return TabItem{
		owner: me.owner,
		index: int32(index),
	}
}

// Retrieves the selected item with [TCM_GETCURSEL], if any
//
// [TCM_GETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getcursel
func (me *CollectionTabItems) Selected() (TabItem, bool) {
	idxRet, _ := me.owner.hWnd.SendMessage(co.TCM_GETCURSEL, 0, 0)
	idx := int(idxRet)
	if idx == -1 {
		return TabItem{}, false
	}
	return me.Get(idx), true
}
