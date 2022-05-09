//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single column of a ListView.
type ListViewColumn struct {
	lv    ListView
	index uint32
}

// Returns the zero-based index of the column.
func (me ListViewColumn) Index() int {
	return int(me.index)
}

// Retrieves all selected item texts under this column.
func (me ListViewColumn) SelectedTexts() []string {
	selItems := me.lv.Items().SelectedItems()
	selTexts := make([]string, 0, len(selItems))

	for _, selItem := range selItems {
		selTexts = append(selTexts, selItem.Text(int(me.index)))
	}
	return selTexts
}

// Sets the title.
func (me ListViewColumn) SetTitle(text string) {
	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
	}
	lvc.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN %d to \"%s\" failed.", me.index, text))
	}
}

// Sets the width. Will be adjusted to the current system DPI.
func (me ListViewColumn) SetWidth(width int) {
	colWidth := win.SIZE{Cx: int32(width), Cy: 0}
	_MultiplyDpi(nil, &colWidth)

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(me.index), win.LPARAM(colWidth.Cx))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMNWIDTH %d to %d failed.", me.index, width))
	}
}

// Resizes the column to fill the remaining space.
func (me ListViewColumn) SetWidthToFill() {
	numCols := me.lv.Columns().Count()
	cxUsed := 0

	for i := 0; i < numCols; i++ {
		if i != int(me.index) {
			cxUsed += me.lv.Columns().Get(i).Width() // retrieve cx of each column, but us
		}
	}

	rc := me.lv.Hwnd().GetClientRect() // list view client area
	fillWidth := int(rc.Right) - cxUsed

	ret := me.lv.Hwnd().SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(me.index), win.LPARAM(fillWidth))
	if ret == 0 {
		panic(fmt.Sprintf(
			"LVM_SETCOLUMNWIDTH %d to %d failed.", me.index, fillWidth))
	}
}

// Retrieves all item texts under this column.
func (me ListViewColumn) Texts() []string {
	items := me.lv.Items().All()
	texts := make([]string, 0, len(items))

	for _, item := range items {
		texts = append(texts, item.Text(int(me.index)))
	}
	return texts
}

// Retrieves the title of the column.
func (me ListViewColumn) Title() string {
	var titleBuf [128]uint16 // arbitrary

	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
	}
	lvc.SetPszText(titleBuf[:])

	ret := me.lv.Hwnd().SendMessage(co.LVM_GETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMN %d failed.", lvc.ISubItem))
	}

	return win.Str.FromNativeSlice(titleBuf[:])
}

// Retrieves the width of the column.
func (me ListViewColumn) Width() int {
	cx := int(
		me.lv.Hwnd().SendMessage(co.LVM_GETCOLUMNWIDTH,
			win.WPARAM(me.index), 0),
	)
	if cx == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMNWIDTH %d failed.", me.index))
	}
	return cx
}
