//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// The columns collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [ListView].
type CollectionListViewCols struct {
	owner *ListView
}

// Add a column with its width, using [LVM_INSERTCOLUMN], and returns the new
// column.
//
// Panics on error.
//
// Example:
//
//	var list ui.ListView // initialized somewhere
//
//	list.Cols.Add("Title", ui.DpiX(80))
//
// [LVM_INSERTCOLUMN]: https://learn.microsoft.com/en-us/windows/win32/controls/lvm-insertcolumn
func (me *CollectionListViewCols) Add(title string, width int) ListViewCol {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()

	lvc := win.LVCOLUMN{
		Mask: co.LVCF_TEXT | co.LVCF_WIDTH,
		Cx:   int32(width),
	}
	lvc.SetPszText(wbuf.SliceAllowEmpty(title))

	newIdxRet, err := me.owner.hWnd.SendMessage(co.LVM_INSERTCOLUMN,
		0xffff, win.LPARAM(unsafe.Pointer(&lvc)))
	newIdx := int(newIdxRet)
	if err != nil || newIdx == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN \"%s\" failed.", title))
	}

	return me.Get(newIdx)
}

// Returns all columns.
func (me *CollectionListViewCols) All() []ListViewCol {
	nCols := me.Count()
	cols := make([]ListViewCol, 0, nCols)
	for i := uint(0); i < nCols; i++ {
		cols = append(cols, me.Get(int(i)))
	}
	return cols
}

// Retrieves the number of columns with [HDM_GETITEMCOUNT].
//
// Panics if the list view has no header.
//
// [HDM_GETITEMCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitemcount
func (me *CollectionListViewCols) Count() uint {
	if me.owner.Header() == nil {
		panic("This ListView has no header.")
	}

	return me.owner.header.Items.Count()
}

// Returns the column at the given index.
func (me *CollectionListViewCols) Get(index int) ListViewCol {
	return ListViewCol{
		owner: me.owner,
		index: int32(index),
	}
}

// Returns the last column.
func (me *CollectionListViewCols) Last() ListViewCol {
	return me.Get(int(me.Count()) - 1)
}
