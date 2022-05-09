//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single item of a ComboBox.
type ComboBoxItem struct {
	cmb   ComboBox
	index uint32
}

// Deletes the item.
func (me *ComboBoxItem) Delete() {
	me.cmb.Hwnd().SendMessage(co.CB_DELETESTRING, win.WPARAM(me.index), 0)
}

// Returns the zero-based index of the item.
func (me ComboBoxItem) Index() int {
	return int(me.index)
}

// Selects the item
func (me ComboBoxItem) Select() {
	me.cmb.Hwnd().SendMessage(co.CB_SETCURSEL, win.WPARAM(me.index), 0)
}

// Retrieves the text of the item.
func (me ComboBoxItem) Text() string {
	nChars := me.cmb.Hwnd().SendMessage(
		co.CB_GETLBTEXTLEN, win.WPARAM(me.index), 0)
	if int(nChars) == -1 {
		panic(fmt.Sprintf("CB_GETLBTEXTLEN failed at item %d.", me.index))
	}

	textBuf := make([]uint16, nChars+1)
	me.cmb.Hwnd().SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&textBuf[0])))
	return win.Str.FromNativeSlice(textBuf)
}
