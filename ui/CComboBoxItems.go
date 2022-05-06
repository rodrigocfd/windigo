package ui

import (
	"runtime"
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
		pText := win.Str.ToNativePtr(text)
		me.cmb.Hwnd().SendMessage(co.CB_ADDSTRING,
			0, win.LPARAM(unsafe.Pointer(pText)))
		runtime.KeepAlive(pText)
	}
}

// Retrieves the number of items.
func (me *_ComboBoxItems) Count() int {
	return int(me.cmb.Hwnd().SendMessage(co.CB_GETCOUNT, 0, 0))
}

// Deletes all items.
func (me *_ComboBoxItems) DeleteAll() {
	me.cmb.Hwnd().SendMessage(co.CB_RESETCONTENT, 0, 0)
}

// Returns the item at the given index.
//
// Note that this method is dumb: no validation is made, the given index is
// simply kept. If the index is invalid (or becomes invalid), subsequent
// operations on the ComboBoxItem will fail.
func (me *_ComboBoxItems) Get(index int) ComboBoxItem {
	return ComboBoxItem{cmb: me.cmb, index: uint32(index)}
}

// Retrieves the selected item, if any.
func (me *_ComboBoxItems) Selected() (ComboBoxItem, bool) {
	idx := int(me.cmb.Hwnd().SendMessage(co.CB_GETCURSEL, 0, 0))
	if idx == -1 {
		return me.Get(-1), false
	}
	return me.Get(idx), true
}

// Retrieves the selected text, if any.
func (me *_ComboBoxItems) SelectedText() (string, bool) {
	if selItem, hasSel := me.Selected(); hasSel {
		return selItem.Text(), true
	}
	return "", false
}
