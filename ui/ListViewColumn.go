package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

type ListViewColumn struct {
	owner *ListView
	index uint32
}

func MakeListViewColumn(owner *ListView, index uint32) ListViewColumn {
	return ListViewColumn{
		owner: owner,
		index: index,
	}
}

func (me *ListViewColumn) Index() uint32 {
	return me.index
}

func (me *ListViewColumn) SetText(text string) *ListViewColumn {
	lvc := api.LVCOLUMN{
		ISubItem: int32(me.index),
		Mask:     c.LVCF_TEXT,
		PszText:  api.StrToUtf16Ptr(text),
	}
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_SETCOLUMN),
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvc)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN failed to set text \"%s\".", text))
	}
	return me
}

func (me *ListViewColumn) SetWidth(width uint32) *ListViewColumn {
	me.owner.Hwnd().SendMessage(c.WM(c.LVM_SETCOLUMNWIDTH),
		api.WPARAM(me.index), api.LPARAM(width))
	return me
}

func (me *ListViewColumn) Text() string {
	buf := make([]uint16, 256) // arbitrary
	lvc := api.LVCOLUMN{
		ISubItem:   int32(me.index),
		Mask:       c.LVCF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_GETCOLUMN),
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvc)))
	if ret < 0 {
		panic("LVM_GETCOLUMN failed to get text.")
	}
	return syscall.UTF16ToString(buf)
}

func (me *ListViewColumn) Width() uint32 {
	cx := me.owner.Hwnd().SendMessage(c.WM(c.LVM_GETCOLUMNWIDTH),
		api.WPARAM(me.index), 0)
	if cx == 0 {
		panic("LVM_GETCOLUMNWIDTH failed.")
	}
	return uint32(cx)
}
