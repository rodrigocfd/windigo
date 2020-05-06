package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

type ListViewItem struct {
	owner *ListView
	index uint32
}

func NewListViewItem(owner *ListView, index uint32) *ListViewItem {
	return &ListViewItem{
		owner: owner,
		index: index,
	}
}

func (me *ListViewItem) delete() {
	if me.index >= me.owner.ItemCount() { // index out of bounds: ignore
		return
	}

	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_DELETEITEM),
		api.WPARAM(me.index), 0)
	if ret == 0 {
		panic(fmt.Sprintf("LVM_DELETEITEM failed, index %d.\n", me.index))
	}
}

func (me *ListViewItem) Index() uint32 {
	return me.index
}

func (me *ListViewItem) IsCut() bool {
	sta := me.owner.Hwnd().SendMessage(c.WM(c.LVM_GETITEMSTATE),
		api.WPARAM(me.index), api.LPARAM(c.LVIS_CUT))
	return (c.LVIS(sta) & c.LVIS_CUT) != 0
}

func (me *ListViewItem) IsFocused() bool {
	sta := me.owner.Hwnd().SendMessage(c.WM(c.LVM_GETITEMSTATE),
		api.WPARAM(me.index), api.LPARAM(c.LVIS_FOCUSED))
	return (c.LVIS(sta) & c.LVIS_FOCUSED) != 0
}

func (me *ListViewItem) IsSelected() bool {
	sta := me.owner.Hwnd().SendMessage(c.WM(c.LVM_GETITEMSTATE),
		api.WPARAM(me.index), api.LPARAM(c.LVIS_SELECTED))
	return (c.LVIS(sta) & c.LVIS_SELECTED) != 0
}

func (me *ListViewItem) IsVisible() bool {
	return me.owner.Hwnd().SendMessage(c.WM(c.LVM_ISITEMVISIBLE),
		api.WPARAM(me.index), 0) != 0
}

func (me *ListViewItem) SetFocus() *ListViewItem {
	lvi := api.LVITEM{
		StateMask: c.LVIS_FOCUSED,
		State:     c.LVIS_FOCUSED,
	}
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_SETITEMSTATE),
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_FOCUSED.")
	}
	return me
}

func (me *ListViewItem) SetSelected(selected bool) *ListViewItem {
	lvi := api.LVITEM{
		StateMask: c.LVIS_SELECTED,
	}
	if selected { // otherwise is zero
		lvi.State = c.LVIS_SELECTED
	}
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_SETITEMSTATE),
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed for LVIS_SELECTED.")
	}
	return me
}

func (me *ListViewItem) SetText(text string) *ListViewItem {
	lvi := api.LVITEM{
		PszText: api.StrToUtf16Ptr(text),
	}
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_SETITEMTEXT),
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic(fmt.Sprintf("LVM_SETITEMTEXT failed \"%s\".", text))
	}
	return me
}

func (me *ListViewItem) Text() string {
	buf := make([]uint16, 256) // arbitrary
	lvi := api.LVITEM{
		PszText:    &buf[0],
		CchTextMax: int32(len(buf)),
	}
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_GETITEMTEXT),
		api.WPARAM(me.index), api.LPARAM(unsafe.Pointer(&lvi)))
	if ret < 0 {
		panic("LVM_GETITEMTEXT failed.")
	}
	return syscall.UTF16ToString(buf)
}

func (me *ListViewItem) Update() *ListViewItem {
	ret := me.owner.Hwnd().SendMessage(c.WM(c.LVM_UPDATE),
		api.WPARAM(me.index), 0)
	if ret == 0 {
		panic("LVM_UPDATE failed.")
	}
	return me
}
