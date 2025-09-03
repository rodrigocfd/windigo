//go:build windows

package ui

import (
	"fmt"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
)

// The items collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [ComboBox].
type CollectionComboBoxItems struct {
	owner *ComboBox
}

// Adds one or more items using [CB_ADDSTRING].
//
// [CB_ADDSTRING]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-addstring
func (me *CollectionComboBoxItems) Add(texts ...string) {
	var wText wstr.BufEncoder
	for _, text := range texts {
		me.owner.hWnd.SendMessage(co.CB_ADDSTRING,
			0, win.LPARAM(wText.AllowEmpty(text)))
	}
}

// Returns all items with [CB_GETLBTEXT].
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *CollectionComboBoxItems) All() []string {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	nItems := me.Count()
	items := make([]string, 0, nItems)

	for i := 0; i < nItems; i++ {
		// nChars, _ := me.owner.hWnd.SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(int32(i)), 0)
		// wBuf.Alloc(int(nChars) + 1)

		wBuf.Zero()
		me.owner.hWnd.SendMessage(co.CB_GETLBTEXT,
			win.WPARAM(int32(i)), win.LPARAM(wBuf.Ptr()))
		items = append(items, wBuf.String())
	}

	return items
}

// Retrieves the number of items with [CB_GETCOUNT].
//
// [CB_GETCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getcount
func (me *CollectionComboBoxItems) Count() int {
	n, _ := me.owner.hWnd.SendMessage(co.CB_GETCOUNT, 0, 0)
	return int(n)
}

// Deletes all items with [CB_RESETCONTENT].
//
// [CB_RESETCONTENT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-resetcontent
func (me *CollectionComboBoxItems) DeleteAll() {
	me.owner.hWnd.SendMessage(co.CB_RESETCONTENT, 0, 0)
}

// Returns the item at the given index with [CB_GETLBTEXT].
//
// Panics if the index is not valid.
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *CollectionComboBoxItems) Get(index int) string {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	nChars, _ := me.owner.hWnd.SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(int32(index)), 0)
	if int32(nChars) == utl.CB_ERR {
		panic(fmt.Sprintf("Invalid ComboBox index: %d", index))
	}
	wBuf.Alloc(int(nChars) + 1)

	me.owner.hWnd.SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(int32(index)), win.LPARAM(wBuf.Ptr()))
	return wBuf.String()
}

// Returns the last item with [CB_GETLBTEXT].
//
// Panics if empty.
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *CollectionComboBoxItems) Last() string {
	return me.Get(int(me.Count()) - 1)
}

// Selects the given item with [CB_SETCURSEL].
//
// If index is -1, selection is cleared.
//
// [CB_SETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-setcursel
func (me *CollectionComboBoxItems) Select(index int) {
	me.owner.hWnd.SendMessage(co.CB_SETCURSEL, win.WPARAM(int32(index)), 0)
}

// Retrieves the selected index with [CB_GETCURSEL].
//
// If no item is selected, returns -1.
//
// [CB_GETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getcursel
func (me *CollectionComboBoxItems) Selected() int {
	n, _ := me.owner.hWnd.SendMessage(co.CB_GETCURSEL, 0, 0)
	return int(n)
}

// Returns the string at the given position with [CB_GETLBTEXT].
//
// Panics on error.
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *CollectionComboBoxItems) Text(index int) string {
	// nChars, _ := me.owner.hWnd.SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(index), 0)
	// if int32(nChars) == -1 {
	// 	panic(fmt.Sprintf("CB_GETLBTEXTLEN failed at item %d.", index))
	// }

	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	me.owner.hWnd.SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(int32(index)), win.LPARAM(wBuf.Ptr()))
	return wBuf.String()
}
