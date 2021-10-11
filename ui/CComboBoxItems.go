package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ComboBoxItems struct {
	cmb ComboBox
}

func (me *_ComboBoxItems) new(ctrl ComboBox) {
	me.cmb = ctrl
}

// Adds one or more items.
func (me *_ComboBoxItems) Add(texts ...string) {
	for _, text := range texts {
		me.cmb.Hwnd().SendMessage(co.CB_ADDSTRING, 0,
			win.LPARAM(unsafe.Pointer(win.Str.ToNativePtr(text))))
	}
}

// Retrieves the number of items.
func (me *_ComboBoxItems) Count() int {
	return int(me.cmb.Hwnd().SendMessage(co.CB_GETCOUNT, 0, 0))
}

// Deletes the item at the given index.
func (me *_ComboBoxItems) Delete(index int) {
	me.cmb.Hwnd().SendMessage(co.CB_DELETESTRING, win.WPARAM(index), 0)
}

// Deletes all items.
func (me *_ComboBoxItems) DeleteAll() {
	me.cmb.Hwnd().SendMessage(co.CB_RESETCONTENT, 0, 0)
}

// Retrieves the selected index, if any.
func (me *_ComboBoxItems) Selected() (int, bool) {
	idx := int(me.cmb.Hwnd().SendMessage(co.CB_GETCURSEL, 0, 0))
	if idx == -1 {
		return -1, false
	}
	return idx, true
}

// Retrieves the selected text, if any.
func (me *_ComboBoxItems) SelectedText() (string, bool) {
	if idx, hasSel := me.Selected(); hasSel {
		return me.Text(idx), true
	}
	return "", false
}

// Sets the selected item.
func (me *_ComboBoxItems) SetSelected(index int) {
	me.cmb.Hwnd().SendMessage(co.CB_SETCURSEL, win.WPARAM(index), 0)
}

// Retrieves the text of the item at the given index.
func (me *_ComboBoxItems) Text(index int) string {
	nChars := me.cmb.Hwnd().SendMessage(co.CB_GETLBTEXTLEN, win.WPARAM(index), 0)
	if int(nChars) == -1 {
		panic(fmt.Sprintf("CB_GETLBTEXTLEN failed at item %d.", index))
	}

	textBuf := make([]uint16, nChars+1)
	me.cmb.Hwnd().SendMessage(co.CB_GETLBTEXT,
		win.WPARAM(index), win.LPARAM(unsafe.Pointer(&textBuf[0])))
	return win.Str.FromNativeSlice(textBuf)
}
