/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// A single column of a list view control.
type ListViewColumn struct {
	owner *ListView
	index uint
}

// Resizes the column to fill the remaining space.
func (me *ListViewColumn) FillRoom() *ListViewColumn {
	numCols := me.owner.ColumnCount()
	cxUsed := uint(0)

	for i := uint(0); i < numCols; i++ {
		if i != me.index {
			cxUsed += me.owner.Column(i).Width() // retrieve cx of each column, but us
		}
	}

	rc := me.owner.Hwnd().GetClientRect() // list view client area
	me.SetWidth(uint(rc.Right) - cxUsed)  // fill available space
	return me
}

// Returns the index of this column.
func (me *ListViewColumn) Index() uint {
	return me.index
}

// Sets the text of this column with LVM_SETCOLUMN.
func (me *ListViewColumn) SetText(text string) *ListViewColumn {
	textBuf := win.StrToSlice(text)
	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
		PszText:  &textBuf[0],
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN failed to set text \"%s\".", text))
	}
	return me
}

// Sends LVM_SETCOLUMNWIDTH.
func (me *ListViewColumn) SetWidth(width uint) *ListViewColumn {
	me.owner.sendLvmMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(me.index), win.LPARAM(width))
	return me
}

// Retrieves the text of this column with LVM_GETCOLUMN.
func (me *ListViewColumn) Text() string {
	buf := [128]uint16{} // arbitrary
	lvc := win.LVCOLUMN{
		ISubItem:   int32(me.index),
		Mask:       co.LVCF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.sendLvmMessage(co.LVM_GETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic("LVM_GETCOLUMN failed to get text.")
	}
	return syscall.UTF16ToString(buf[:])
}

// Sends LVM_GETCOLUMNWIDTH.
func (me *ListViewColumn) Width() uint {
	cx := me.owner.sendLvmMessage(co.LVM_GETCOLUMNWIDTH, win.WPARAM(me.index), 0)
	if cx == 0 {
		panic("LVM_GETCOLUMNWIDTH failed.")
	}
	return uint(cx)
}
