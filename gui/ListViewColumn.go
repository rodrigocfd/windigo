/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

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
	index uint32
}

func (me *ListViewColumn) Index() uint32 {
	return me.index
}

func (me *ListViewColumn) SetText(text string) *ListViewColumn {
	lvc := win.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     co.LVCF_TEXT,
		PszText:  win.StrToPtr(text),
	}
	ret := me.owner.sendLvmMessage(co.LVM_SETCOLUMN,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN failed to set text \"%s\".", text))
	}
	return me
}

func (me *ListViewColumn) SetWidth(width uint32) *ListViewColumn {
	me.owner.sendLvmMessage(co.LVM_SETCOLUMNWIDTH,
		win.WPARAM(me.index), win.LPARAM(width))
	return me
}

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

func (me *ListViewColumn) Width() uint32 {
	cx := me.owner.sendLvmMessage(co.LVM_GETCOLUMNWIDTH, win.WPARAM(me.index), 0)
	if cx == 0 {
		panic("LVM_GETCOLUMNWIDTH failed.")
	}
	return uint32(cx)
}
