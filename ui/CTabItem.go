//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single item of a Tab.
type TabItem struct {
	t     Tab
	index uint32
}

// Deletes the item.
func (me TabItem) Delete() {
	ret := me.t.Hwnd().SendMessage(co.TCM_DELETEITEM, win.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("TCM_DELETEITEM %d failed.", me.index))
	}
}

// Sets the item as the focused one.
func (me TabItem) Focus() {
	me.t.Hwnd().SendMessage(co.TCM_SETCURFOCUS, win.WPARAM(me.index), 0)
}

// Returns the zero-based index of the item.
func (me TabItem) Index() int {
	return int(me.index)
}

// Tells whether the item is currently selected.
func (me TabItem) IsSelected() bool {
	return uint32(me.t.Hwnd().SendMessage(co.TCM_GETCURSEL, 0, 0)) == me.index
}

// Tells whether the item is currently visible.
// func (me TabItem) IsVisible() bool {
// 	return me.t.Hwnd().SendMessage(co.TCM_ISITEMVISIBLE,
// 		win.WPARAM(me.index), 0) != 0
// }

// Retrieves the custom data associated with the item.
func (me TabItem) LParam() win.LPARAM {
	tci := win.TCITEM{
		Mask: co.TCIF_PARAM,
	}

	ret := me.t.Hwnd().SendMessage(co.TCM_GETITEMA,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&tci)))
	if ret == 0 {
		panic(fmt.Sprintf("TCM_GETITEM %d failed.", me.index))
	}

	return tci.LParam
}

// Retrieves the coordinates of the rectangle surrounding the item.
func (me TabItem) Rect() win.RECT {
	rcItem := win.RECT{}

	ret := me.t.Hwnd().SendMessage(co.TCM_GETITEMRECT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&rcItem)))
	if ret == 0 {
		panic(fmt.Sprintf("TCM_GETITEMRECT %d failed.", me.index))
	}
	return rcItem // coordinates relative to the Tab
}

// Selects the item.
func (me TabItem) Select() {
	me.t.Hwnd().SendMessage(co.TCM_SETCURSEL, win.WPARAM(me.index), 0)
}

// Sets the custom data associated with the item.
func (me TabItem) SetLParam(lp win.LPARAM) {
	tci := win.TCITEM{
		Mask:   co.TCIF_PARAM,
		LParam: lp,
	}

	ret := me.t.Hwnd().SendMessage(co.TCM_SETITEMA,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&tci)))
	if ret == 0 {
		panic(fmt.Sprintf("TCM_SETITEM %d failed.", me.index))
	}
}

// Sets the text of the item.
func (me TabItem) SetText(text string) {
	tci := win.TCITEM{}
	tci.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.t.Hwnd().SendMessage(co.TCM_SETITEMA,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&tci)))
	if ret == 0 {
		panic(fmt.Sprintf("TCM_SETITEM %d failed \"%s\".",
			me.index, text))
	}
}

// Retrieves the text of the item.
func (me TabItem) Text() string {
	const BLOCK int = 64 // arbitrary
	bufSz := BLOCK
	var buf []uint16

	tci := win.TCITEM{
		Mask: co.TCIF_TEXT,
	}

	for {
		buf = make([]uint16, bufSz)
		tci.SetPszText(buf)

		nChars := int(
			me.t.Hwnd().SendMessage(co.TCM_GETITEMA,
				win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&tci))),
		)

		if nChars+1 < bufSz { // to break, must have at least 1 char gap
			fmt.Println(nChars, bufSz)
			break
		}

		bufSz += BLOCK // increase buffer size to try again
	}

	return win.Str.FromNativeSlice(buf)
}
