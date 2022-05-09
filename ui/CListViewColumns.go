//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ListViewColumns struct {
	lv ListView
}

func (me *_ListViewColumns) new(ctrl ListView) {
	me.lv = ctrl
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

		lvc.Cx = colWidth.Cx
		lvc.SetPszText(win.Str.ToNativeSlice(titles[i]))

		newIdx := int(
			me.lv.Hwnd().SendMessage(co.LVM_INSERTCOLUMN,
				0xffff, win.LPARAM(unsafe.Pointer(&lvc))),
		)
		if newIdx == -1 {
			panic(fmt.Sprintf("LVM_INSERTCOLUMN \"%s\" failed.", titles[i]))
		}
	}
}

// Retrieves the number of columns.
func (me *_ListViewColumns) Count() int {
	hHeader := win.HWND(me.lv.Hwnd().SendMessage(co.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := int(hHeader.SendMessage(co.HDM_GETITEMCOUNT, 0, 0))
	if count == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return count
}

// Returns the column at the given index.
//
// Note that this method is dumb: no validation is made, the given index is
// simply kept. If the index is invalid (or becomes invalid), subsequent
// operations on the ListViewColumn will fail.
func (me *_ListViewColumns) Get(index int) ListViewColumn {
	return ListViewColumn{lv: me.lv, index: uint32(index)}
}
