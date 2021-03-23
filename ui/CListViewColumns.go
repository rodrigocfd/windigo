package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ListViewColumns struct {
	pHwnd *win.HWND
}

func (me *_ListViewColumns) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
}

// Adds one or more columns with their widths.
// Widths will be adjusted to the current system DPI.
func (me *_ListViewColumns) Add(widths []int, titles ...string) {
	if len(titles) != len(widths) {
		panic(fmt.Sprintf("Unmatching titles (%d) and widths (%d).",
			len(titles), len(widths)))
	}

	lvc := win.LVCOLUMN{
		Mask: co.LVCF_TEXT | co.LVCF_WIDTH,
	}

	for i := 0; i < len(titles); i++ {
		colWidth := win.SIZE{Cx: int32(widths[i]), Cy: 0}
		_MultiplyDpi(nil, &colWidth)

		textBuf := win.Str.ToUint16Slice(titles[i])
		lvc.PszText = &textBuf[0]
		lvc.Cx = colWidth.Cx

		newIdx := int(
			me.pHwnd.SendMessage(co.LVM_INSERTCOLUMN,
				0xffff, win.LPARAM(unsafe.Pointer(&lvc))),
		)
		if newIdx == -1 {
			panic(fmt.Sprintf("LVM_INSERTCOLUMN \"%s\" failed.", titles[i]))
		}
	}
}

// Retrieves the texts of all items at the given column.
func (me *_ListViewColumns) AllTexts(columnIndex int) []string {
	count := int(me.pHwnd.SendMessage(co.LVM_GETITEMCOUNT, 0, 0))
	texts := make([]string, 0, count)

	textBuf := [256]uint16{} // arbitrary
	lvi := win.LVITEM{
		ISubItem:   int32(columnIndex),
		PszText:    &textBuf[0],
		CchTextMax: int32(len(textBuf)),
	}

	for i := 0; i < count; i++ {
		ret := me.pHwnd.SendMessage(co.LVM_GETITEMTEXT,
			win.WPARAM(i), win.LPARAM(unsafe.Pointer(&lvi)))
		if ret < 0 {
			panic(fmt.Sprintf("LVM_GETITEMTEXT %d/%d failed.", i, columnIndex))
		}

		texts = append(texts, win.Str.FromUint16Slice(textBuf[:]))
	}

	return texts
}

// Retrieves the number of columns.
func (me *_ListViewColumns) Count() int {
	hHeader := win.HWND(me.pHwnd.SendMessage(co.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := int(hHeader.SendMessage(co.HDM_GETITEMCOUNT, 0, 0))
	if count == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return count
}

// Retrieves information about a column.
func (me *_ListViewColumns) Info(lvc *win.LVCOLUMN) {
	ret := me.pHwnd.SendMessage(co.LVM_GETCOLUMN,
		win.WPARAM(lvc.ISubItem), win.LPARAM(unsafe.Pointer(lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMN %d failed.", lvc.ISubItem))
	}
}

// Retrieves the texts of the selected items at the given column.
func (me *_ListViewColumns) SelectedTexts(columnIndex int) []string {
	selCount := int(me.pHwnd.SendMessage(co.LVM_GETSELECTEDCOUNT, 0, 0))
	texts := make([]string, 0, selCount)

	textBuf := [256]uint16{} // arbitrary
	lvi := win.LVITEM{
		ISubItem:   int32(columnIndex),
		PszText:    &textBuf[0],
		CchTextMax: int32(len(textBuf)),
	}

	i := -1
	for {
		i = int(
			me.pHwnd.SendMessage(co.LVM_GETNEXTITEM,
				win.WPARAM(i), win.LPARAM(co.LVNI_SELECTED)),
		)
		if i == -1 {
			break
		}

		ret := me.pHwnd.SendMessage(co.LVM_GETITEMTEXT,
			win.WPARAM(i), win.LPARAM(unsafe.Pointer(&lvi)))
		if ret < 0 {
			panic(fmt.Sprintf("LVM_GETITEMTEXT %d/%d failed.", i, columnIndex))
		}

		texts = append(texts, win.Str.FromUint16Slice(textBuf[:]))
	}

	return texts
}

// Sets the title of this column.
func (me *_ListViewColumns) SetTitle(columnIndex int, text string) {
	titleBuf := win.Str.ToUint16Slice(text)
	lvc := win.LVCOLUMN{
		ISubItem: int32(columnIndex),
		Mask:     co.LVCF_TEXT,
		PszText:  &titleBuf[0],
	}

	ret := me.pHwnd.SendMessage(co.LVM_SETCOLUMN,
		win.WPARAM(columnIndex), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN %d to \"%s\" failed.", columnIndex, text))
	}
}

// Sets the width of the column. Will be adjusted to the current system DPI.
func (me *_ListViewColumns) SetWidth(columnIndex, width int) {
	colWidth := win.SIZE{Cx: int32(width), Cy: 0}
	_MultiplyDpi(nil, &colWidth)

	ret := me.pHwnd.SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(columnIndex), win.LPARAM(colWidth.Cx))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMNWIDTH %d to %d failed.", columnIndex, width))
	}
}

// Resizes the column to fill the remaining space.
func (me *_ListViewColumns) SetWidthToFill(columnIndex int) {
	numCols := me.Count()
	cxUsed := 0

	for i := 0; i < numCols; i++ {
		if i != columnIndex {
			cxUsed += me.Width(i) // retrieve cx of each column, but us
		}
	}

	rc := me.pHwnd.GetClientRect() // list view client area
	me.pHwnd.SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(columnIndex), win.LPARAM(int(rc.Right)-cxUsed)) // fill available space
}

// Retrieves the title of the column at the given index.
func (me *_ListViewColumns) Title(columnIndex int) string {
	titleBuf := [128]uint16{} // arbitrary
	lvc := win.LVCOLUMN{
		ISubItem:   int32(columnIndex),
		Mask:       co.LVCF_TEXT,
		PszText:    &titleBuf[0],
		CchTextMax: int32(len(titleBuf)),
	}

	me.Info(&lvc)
	return win.Str.FromUint16Slice(titleBuf[:])
}

// Retrieves the width of the column.
func (me *_ListViewColumns) Width(columnIndex int) int {
	cx := int(
		me.pHwnd.SendMessage(co.LVM_GETCOLUMNWIDTH, win.WPARAM(columnIndex), 0),
	)
	if cx == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMNWIDTH %d failed.", columnIndex))
	}
	return cx
}
