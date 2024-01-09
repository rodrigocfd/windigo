//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _TabItems struct {
	t Tab
}

func (me *_TabItems) new(ctrl Tab) {
	me.t = ctrl
}

// Adds a tab, specifying the text for it, returning the new tab.
func (me *_TabItems) Add(text string) TabItem {
	return me.AddWithIcon(-1, text)
}

// Adds a tab, specifying its icon and the text for it, returning the new tab.
func (me *_TabItems) AddWithIcon(iconIndex int, text string) TabItem {
	tci := win.TCITEM{
		Mask:   co.TCIF_TEXT | co.TCIF_IMAGE,
		IImage: int32(iconIndex),
	}
	tci.SetPszText(win.Str.ToNativeSlice(text))

	newIdx := int(
		me.t.Hwnd().SendMessage(co.TCM_INSERTITEMA,
			win.WPARAM(me.Count()+1), win.LPARAM(unsafe.Pointer(&tci))),
	)
	if newIdx == -1 {
		panic(fmt.Sprintf("TCM_INSERTITEMA col %d, \"%s\" failed.", 0, text))
	}

	return me.Get(newIdx)
}

// Retrieves all the items.
func (me *_TabItems) All() []TabItem {
	numItems := me.Count()
	items := make([]TabItem, 0, numItems)
	for i := 0; i < numItems; i++ {
		items = append(items, me.Get(i))
	}
	return items
}

// Retrieves the number of items.
func (me *_TabItems) Count() int {
	return int(me.t.Hwnd().SendMessage(co.TCM_GETITEMCOUNT, 0, 0))
}

// Deletes all items at once.
func (me *_TabItems) DeleteAll() {
	me.t.Hwnd().SendMessage(co.TCM_DELETEALLITEMS, 0, 0)
}

// Retrieves the focused item, if any.
func (me *_TabItems) Focused() (TabItem, bool) {
	idx := int(
		me.t.Hwnd().SendMessage(co.TCM_GETCURFOCUS,
			win.WPARAM(0), win.LPARAM(0)),
	)
	if idx == -1 {
		return me.Get(-1), false
	}

	return me.Get(idx), true
}

// Returns the tab at the given index.
//
// Note that this method is dumb: no validation is made, the given index is
// simply kept. If the index is invalid (or becomes invalid), subsequent
// operations on the TabItem will fail.
func (me *_TabItems) Get(index int) TabItem {
	return TabItem{t: me.t, index: uint32(index)}
}

// Retrieves the rab below the given coordinates, if any.
//
// The coordinates must be relative to the Tab.
func (me *_TabItems) HitTest(pos win.POINT) (TabItem, bool) {
	tchti := win.TCHITTESTINFO{
		Pt: pos,
	}

	idx := int(me.t.Hwnd().SendMessage(co.TCM_HITTEST,
		win.WPARAM(0), win.LPARAM(unsafe.Pointer(&tchti))))

	if idx == -1 {
		return me.Get(-1), false
	}
	return me.Get(int(idx)), true
}

// Selects or deselects all tabs at once (if possible).
// func (me *_TabItems) SelectAll(doSelect bool) {
// 	tcStyles := co.TCS(me.t.Hwnd().GetWindowLongPtr(co.GWLP_STYLE))
// 	if (tcStyles & co.TCS_MULTISELECT) == 0 {
// 		return // single-sel list views cannot have all items selected
// 	}
//
// 	tci := win.TCITEM{
// 		State:     util.Iif(doSelect, co.TCIS_BUTTONPRESSED, 0).(co.TCIS),
// 		StateMask: co.TCIS_BUTTONPRESSED,
// 	}
//
// 	idx := -1
// 	ret := me.t.Hwnd().SendMessage(co.LVM_SETITEMSTATE,
// 		win.WPARAM(idx), win.LPARAM(unsafe.Pointer(&lvi)))
// 	if ret == 0 {
// 		panic("LVM_SETITEMSTATE failed.")
// 	}
// }

// Retrieves the current tab.
func (me *_TabItems) CurrentTab() TabItem {
	idx := int(
		me.t.Hwnd().SendMessage(co.TCM_GETCURSEL, 0, 0),
	)
	if idx == -1 {
		return me.Get(-1)
	}

	return me.Get(idx)
}

// Retrieves the topmost visible item, if any.
func (me *_TabItems) TopmostVisible() (TabItem, bool) {
	idx := int(me.t.Hwnd().SendMessage(co.LVM_GETTOPINDEX, 0, 0))
	if idx == -1 {
		return me.Get(-1), false
	}

	return me.Get(idx), true
}
