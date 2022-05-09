//go:build windows

package ui

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// A single part of a StatusBar.
type StatusBarPart struct {
	sb    StatusBar
	index uint32
}

// Retrieves the HICON.
//
// The icon is shared, the StatusBar doesn't own it.
func (me StatusBarPart) Icon() win.HICON {
	return win.HICON(
		me.sb.Hwnd().SendMessage(co.SB_GETICON, win.WPARAM(me.index), 0),
	)
}

// Returns the zero-based index of the part.
func (me StatusBarPart) Index() int {
	return int(me.index)
}

// Puts the HICON.
//
// The icon is shared, the StatusBar doesn't own it.
func (me StatusBarPart) SetIcon(hIcon win.HICON) {
	me.sb.Hwnd().SendMessage(co.SB_SETICON,
		win.WPARAM(me.index), win.LPARAM(hIcon))
}

// Sets the text.
func (me StatusBarPart) SetText(text string) {
	pText := win.Str.ToNativePtr(text)
	ret := me.sb.Hwnd().SendMessage(co.SB_SETTEXT,
		win.MAKEWPARAM(win.MAKEWORD(uint8(me.index), 0), 0),
		win.LPARAM(unsafe.Pointer(pText)))
	runtime.KeepAlive(pText)
	if ret == 0 {
		panic(fmt.Sprintf("SB_SETTEXT %d failed \"%s\".", me.index, text))
	}
}

// Retrieves the text of the part.
func (me StatusBarPart) Text() string {
	len := uint16(
		me.sb.Hwnd().SendMessage(co.SB_GETTEXTLENGTH, win.WPARAM(me.index), 0),
	)
	if len == 0 {
		return ""
	}

	buf := make([]uint16, len+1)
	me.sb.Hwnd().SendMessage(co.SB_GETTEXT,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&buf[0])))
	return win.Str.FromNativeSlice(buf)
}
