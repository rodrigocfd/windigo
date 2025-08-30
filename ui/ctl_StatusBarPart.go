//go:build windows

package ui

import (
	"fmt"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// A part from a [status bar].
//
// [status bar]: https://learn.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBarPart struct {
	owner *StatusBar
	index int32
}

// Retrieves the icon with [SB_GETICON].
//
// The icon is shared, the StatusBar doesn't own it.
//
// [SB_GETICON]: https://learn.microsoft.com/en-us/windows/win32/controls/sb-geticon
func (me StatusBarPart) Icon() win.HICON {
	h, _ := me.owner.hWnd.SendMessage(co.SB_GETICON, win.WPARAM(me.index), 0)
	return win.HICON(h)
}

// Returns the zero-based index of the column.
func (me StatusBarPart) Index() int {
	return int(me.index)
}

// Sets the icon with [SB_SETICON].
//
// The icon is shared, the [StatusBar] doesn't own it.
//
// Returns the same part, so further operations can be chained.
//
// [SB_SETICON]: https://learn.microsoft.com/en-us/windows/win32/controls/sb-seticon
func (me StatusBarPart) SetIcon(hIcon win.HICON) StatusBarPart {
	me.owner.hWnd.SendMessage(co.SB_SETICON,
		win.WPARAM(me.index), win.LPARAM(hIcon))
	return me
}

// Sets the text with [SB_SETTEXT].
//
// Returns the same part, so further operations can be chained.
//
// Panics on error.
//
// [SB_SETTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/sb-settext
func (me StatusBarPart) SetText(text string) StatusBarPart {
	var wText wstr.BufEncoder
	ret, _ := me.owner.hWnd.SendMessage(co.SB_SETTEXT,
		win.MAKEWPARAM(win.MAKEWORD(uint8(me.index), 0), 0),
		win.LPARAM(wText.AllowEmpty(text)))
	if ret == 0 {
		panic(fmt.Sprintf("SB_SETTEXT %d failed \"%s\".", me.index, text))
	}

	return me
}

// Retrieves the text with [SB_GETTEXT].
//
// [SB_GETTEXT]: https://learn.microsoft.com/en-us/windows/win32/controls/sb-gettext
func (me StatusBarPart) Text() string {
	nLen, _ := me.owner.hWnd.SendMessage(co.SB_GETTEXTLENGTH, win.WPARAM(me.index), 0)
	len := uint(nLen)
	if len == 0 {
		return ""
	}

	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	me.owner.hWnd.SendMessage(co.SB_GETTEXT,
		win.WPARAM(me.index), win.LPARAM(wBuf.Ptr()))
	return wBuf.String()
}
