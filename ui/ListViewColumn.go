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

func (lvc *ListViewColumn) GetText() string {
	buf := make([]uint16, 256) // arbitrary
	lvcol := api.LVCOLUMN{
		ISubItem:   int32(lvc.index),
		Mask:       c.LVCF_TEXT,
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := lvc.owner.Hwnd().SendMessage(c.WM(c.LVM_GETCOLUMN),
		api.WPARAM(lvc.index), api.LPARAM(unsafe.Pointer(&lvcol)))
	if ret < 0 {
		panic("LVM_GETCOLUMN failed to get text.")
	}
	return syscall.UTF16ToString(buf)
}

func (lvc *ListViewColumn) GetWidth() uint32 {
	cx := lvc.owner.Hwnd().SendMessage(c.WM(c.LVM_GETCOLUMNWIDTH),
		api.WPARAM(lvc.index), 0)
	if cx == 0 {
		panic("LVM_GETCOLUMNWIDTH failed.")
	}
	return uint32(cx)
}

func (lvc *ListViewColumn) Index() uint32 {
	return lvc.index
}

func (lvc *ListViewColumn) SetText(text string) *ListViewColumn {
	lvcol := api.LVCOLUMN{
		ISubItem: int32(lvc.index),
		Mask:     c.LVCF_TEXT,
		PszText:  api.StrToUtf16Ptr(text),
	}
	ret := lvc.owner.Hwnd().SendMessage(c.WM(c.LVM_SETCOLUMN),
		api.WPARAM(lvc.index), api.LPARAM(unsafe.Pointer(&lvcol)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETCOLUMN failed to set text \"%s\".", text))
	}
	return lvc
}

func (lvc *ListViewColumn) SetWidth(width uint32) *ListViewColumn {
	lvc.owner.Hwnd().SendMessage(c.WM(c.LVM_SETCOLUMNWIDTH),
		api.WPARAM(lvc.index), api.LPARAM(width))
	return lvc
}
