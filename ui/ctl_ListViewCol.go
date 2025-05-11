//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// An column from a [list view].
//
// [list view]: https://learn.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type ListViewCol struct {
	owner *ListView
	index int32
}

// Returns the text of each item under this column with [LVM_GETITEMTEXT].
//
// [LVM_GETITEMTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemtext
func (me ListViewCol) AllTexts() []string {
	nItems := me.owner.Items.Count()
	texts := make([]string, 0, nItems)
	for i := uint(0); i < nItems; i++ {
		item := me.owner.Items.Get(int(i))
		texts = append(texts, item.Text(me.Index()))
	}
	return texts
}

// Returns the zero-based index of the column.
func (me ListViewCol) Index() int {
	return int(me.index)
}

// Retrieves the text justification with [HDM_GETITEM].
//
// Possible values:
//   - HDF_LEFT
//   - HDF_CENTER
//   - HDF_RIGHT
//
// Panics if the list view has no header.
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me ListViewCol) Justification() co.HDF {
	if me.owner.Header() == nil {
		panic("This ListView has no header.")
	}

	return me.owner.header.Items.Get(int(me.index)).Justification()
}

// Sets the text justification with [LVM_GETHEADER] and [HDM_SETITEM].
//
// Possible values:
//   - HDF_LEFT
//   - HDF_CENTER
//   - HDF_RIGHT
//
// Returns the same column, so further operations can be chained.
//
// Panics if the list view has no header.
//
// [LVM_GETHEADER]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getheader
// [HDM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setitem
func (me ListViewCol) SetJustification(hdf co.HDF) ListViewCol {
	if me.owner.Header() == nil {
		panic("This ListView has no header.")
	}

	me.owner.header.Items.Get(int(me.index)).SetJustification(hdf)
	return me
}

// Returns the text of each selected item under this column with
// [LVM_GETNEXTITEM] and [LVM_GETITEMTEXT].
//
// [LVM_GETNEXTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
// [LVM_GETITEMTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getitemtext
func (me ListViewCol) SelectedTexts() []string {
	nSel := me.owner.Items.SelectedCount()
	texts := make([]string, 0, nSel)

	idx := -1
	for {
		idxRet, _ := me.owner.hWnd.SendMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED))
		idx = int(idxRet)
		if idx == -1 {
			break
		}

		item := me.owner.Items.Get(idx)
		texts = append(texts, item.Text(me.Index()))
	}

	return texts
}

// Sets the displayed sort arrow with [LVM_GETHEADER] and [HDM_SETITEM].
//
// Possible values:
//   - co.HDF_NONE
//   - co.HDF_SORTUP
//   - co.HDF_SORTDOWN
//
// Returns the same column, so further operations can be chained.
//
// Panics if the list view has no header.
//
// [LVM_GETHEADER]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getheader
// [HDM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setitem
func (me ListViewCol) SetSortArrow(hdf co.HDF) ListViewCol {
	if me.owner.Header() == nil {
		panic("This ListView has no header.")
	}

	me.owner.header.Items.Get(int(me.index)).SetSortArrow(hdf)
	return me
}

// Sets the title with [LVM_SETCOLUMN].
//
// Returns the same column, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETCOLUMN]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setcolumn
func (me ListViewCol) SetTitle(title string) ListViewCol {
	title16 := wstr.NewBufWith[wstr.Stack20](title, wstr.ALLOW_EMPTY)
	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
	}
	lvc.SetPszText(title16.HotSlice())

	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN %d to \"%s\" failed.", me.index, title))
	}

	return me
}

// Sets the width with [LVM_SETCOLUMNWIDTH].
//
// Returns the same column, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETCOLUMNWIDTH]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setcolumnwidth
func (me ListViewCol) SetWidth(width int) ListViewCol {
	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(me.index), win.LPARAM(width))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMNWIDTH %d to %d failed.", me.index, width))
	}

	return me
}

// Resizes the column with [LVM_SETCOLUMNWIDTH] to fill the remaining space.
//
// Returns the same column, so further operations can be chained.
//
// Panics on error.
//
// [LVM_SETCOLUMNWIDTH]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-setcolumnwidth
func (me ListViewCol) SetWidthToFill() ListViewCol {
	numCols := int(me.owner.Cols.Count())
	cxUsed := uint(0)

	for i := 0; i < numCols; i++ {
		if i != int(me.index) {
			cxUsed += uint(me.owner.Cols.Get(i).Width()) // retrieve cx of each column, but us
		}
	}

	rc, _ := me.owner.hWnd.GetClientRect() // list view client area
	fillWidth := uint(rc.Right) - cxUsed

	ret, err := me.owner.hWnd.SendMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(me.index), win.LPARAM(fillWidth))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf(
			"LVM_SETCOLUMNWIDTH %d to %d failed.", me.index, fillWidth))
	}

	return me
}

// Retrieves the displayed sort arrow with [HDM_GETITEM].
//
// Possible values:
//   - co.HDF_NONE
//   - co.HDF_SORTUP
//   - co.HDF_SORTDOWN
//
// Panics if the list view has no header.
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me ListViewCol) SortArrow() co.HDF {
	if me.owner.Header() == nil {
		panic("This ListView has no header.")
	}

	return me.owner.header.Items.Get(int(me.index)).SortArrow()
}

// Retrieves the title of the column with [LVM_GETCOLUMN].
//
// Panics on error.
//
// [LVM_GETCOLUMN]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getcolumn
func (me ListViewCol) Title() string {
	var titleBuf [64]uint16 // arbitrary

	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
	}
	lvc.SetPszText(titleBuf[:])

	ret, err := me.owner.hWnd.SendMessage(co.LVM_GETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMN %d failed.", me.index))
	}

	return wstr.WstrSliceToStr(titleBuf[:])
}

// Retrieves the width of the column with [LVM_GETCOLUMNWIDTH].
//
// Panics on error.
//
// [LVM_GETCOLUMNWIDTH]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-getcolumnwidth
func (me ListViewCol) Width() int {
	cx, err := me.owner.hWnd.SendMessage(co.LVM_GETCOLUMNWIDTH,
		win.WPARAM(me.index), 0)
	if err != nil || cx == 0 {
		panic(fmt.Sprintf("LVM_GETCOLUMNWIDTH %d failed.", me.index))
	}
	return int(cx)
}
