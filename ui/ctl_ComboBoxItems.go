//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()

	for _, text := range texts {
		pText := wbuf.PtrAllowEmpty(text)
		me.owner.hWnd.SendMessage(co.CB_ADDSTRING,
			0, win.LPARAM(pText))
		wbuf.Clear()
	}
}

// Returns all items with [CB_GETLBTEXT].
//
// [CB_GETLBTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getlbtext
func (me *CollectionComboBoxItems) All() []string {
	recvBuf := wstr.NewBufDecoder(wstr.BUF_MAX)
	defer recvBuf.Free()

	nItems := me.Count()
	items := make([]string, 0, nItems)

	for i := uint(0); i < nItems; i++ {
		// nChars, _ := me.owner.hWnd.SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(i), 0)
		// recvBuf.Resize(uint(nChars) + 1)

		me.owner.hWnd.SendMessage(co.CB_GETLBTEXT,
			win.WPARAM(i), win.LPARAM(recvBuf.UnsafePtr()))
		items = append(items, recvBuf.String())
	}

	return items
}

// Retrieves the number of items with [CB_GETCOUNT].
//
// [CB_GETCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/cb-getcount
func (me *CollectionComboBoxItems) Count() uint {
	n, _ := me.owner.hWnd.SendMessage(co.CB_GETCOUNT, 0, 0)
	return uint(n)
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
	recvBuf := wstr.NewBufDecoder(wstr.BUF_MAX)
	defer recvBuf.Free()

	// nChars, _ := me.owner.hWnd.SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(index), 0)
	// if int32(nChars) == -1 { // CB_ERR
	// 	panic(fmt.Sprintf("Invalid ComboBox index: %d", index))
	// }
	// recvBuf.Resize(uint(nChars) + 1)

	me.owner.hWnd.SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(index), win.LPARAM(recvBuf.UnsafePtr()))
	return recvBuf.String()
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
	me.owner.hWnd.SendMessage(co.CB_SETCURSEL, win.WPARAM(index), 0)
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
func (me *CollectionComboBoxItems) Text(index uint) string {
	// nChars, _ := me.owner.hWnd.SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(index), 0)
	// if int32(nChars) == -1 {
	// 	panic(fmt.Sprintf("CB_GETLBTEXTLEN failed at item %d.", index))
	// }

	recvBuf := wstr.NewBufDecoder(wstr.BUF_MAX)
	defer recvBuf.Free()

	me.owner.hWnd.SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(index), win.LPARAM(recvBuf.UnsafePtr()))
	return recvBuf.String()
}
