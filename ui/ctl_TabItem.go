//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// An item from a [tab].
//
// [tab]: https://learn.microsoft.com/en-us/windows/win32/controls/tab-controls
type TabItem struct {
	owner *Tab
	index int32
}

func (me TabItem) displayContent() {
	if len(me.owner.children) == 0 {
		return
	}

	for idx, child := range me.owner.children {
		if idx != int(me.index) {
			child.Content.Hwnd().ShowWindow(co.SW_HIDE) // hide all others
		}
	}

	hParent, _ := me.owner.hWnd.GetParent()
	rcTab, _ := me.owner.hWnd.GetWindowRect()
	hParent.ScreenToClientRc(&rcTab)
	me.owner.hWnd.SendMessage(co.TCM_ADJUSTRECT, 0, win.LPARAM(unsafe.Pointer(&rcTab))) // ideal child size
	me.owner.children[me.index].Content.Hwnd().
		SetWindowPos(win.HWND(0), int(rcTab.Left), int(rcTab.Top),
			uint(rcTab.Right-rcTab.Left), uint(rcTab.Bottom-rcTab.Top),
			co.SWP_NOZORDER|co.SWP_SHOWWINDOW)
}

// Returns the zero-based index of the item.
func (me TabItem) Index() int {
	return int(me.index)
}

// Selects this tab with [TCM_SETCURSEL].
//
// Returns the same item, so further operations can be chained.
//
// [TCM_SETCURSEL]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-setcursel
func (me TabItem) Select() TabItem {
	me.owner.hWnd.SendMessage(co.TCM_SETCURSEL, win.WPARAM(me.index), 0)
	me.displayContent() // because notification is not sent
	return me
}

// Sets the text with [TCM_SETITEM].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [TCM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-setitem
func (me TabItem) SetText(text string) TabItem {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()

	tci := win.TCITEM{
		Mask: co.TCIF_TEXT,
	}
	tci.SetPszText(wbuf.SliceAllowEmpty(text))

	ret, err := me.owner.hWnd.SendMessage(co.TCM_SETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&tci)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("TCM_SETITEM %d to \"%s\" failed.", me.index, text))
	}

	return me
}

// Retrieves the text with [TCM_GETITEM].
//
// Panics on error.
//
// [TCM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tcm-getitem
func (me TabItem) Text() string {
	recvBuf := wstr.NewBufDecoder(wstr.BUF_MAX)
	defer recvBuf.Free()

	tci := win.TCITEM{
		Mask: co.TCIF_TEXT,
	}
	tci.SetPszText(recvBuf.HotSlice())

	ret, err := me.owner.hWnd.SendMessage(co.TCM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&tci)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("TCM_GETITEM %d failed.", me.index))
	}

	return recvBuf.String()
}
