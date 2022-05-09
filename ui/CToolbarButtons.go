//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ToolbarButtons struct {
	tb Toolbar
}

func (me *_ToolbarButtons) new(ctrl Toolbar) {
	me.tb = ctrl
}

// Adds a button.
func (me *_ToolbarButtons) Add(
	imgListIndex, iconIndex, cmdId int, text string) *_ToolbarButtons {

	tbb := win.TBBUTTON{
		IdCommand: int32(cmdId),
		FsStyle:   co.BTNS_AUTOSIZE,
		FsState:   co.TBSTATE_ENABLED,
		IString:   win.Str.ToNativePtr(text),
	}
	tbb.SetIBitmap(iconIndex, imgListIndex)

	ret := me.tb.Hwnd().SendMessage(co.TB_ADDBUTTONS,
		1, win.LPARAM(unsafe.Pointer(&tbb)))
	if ret == 0 {
		panic(fmt.Sprintf("TB_ADDBUTTONS \"%s\" failed.", text))
	}

	me.tb.Hwnd().SendMessage(co.TB_AUTOSIZE, 0, 0)
	return me
}

// Adds a separator.
func (me *_ToolbarButtons) AddSeparator() *_ToolbarButtons {
	tbb := win.TBBUTTON{
		FsStyle: co.BTNS_SEP,
	}

	ret := me.tb.Hwnd().SendMessage(co.TB_ADDBUTTONS,
		1, win.LPARAM(unsafe.Pointer(&tbb)))
	if ret == 0 {
		panic("TB_ADDBUTTONS failed for separator.")
	}
	me.tb.Hwnd().SendMessage(co.TB_AUTOSIZE, 0, 0)
	return me
}

// Retrieves the number of buttons.
func (me *_ToolbarButtons) Count() int {
	return int(me.tb.Hwnd().SendMessage(co.TB_BUTTONCOUNT, 0, 0))
}

// Deletes a button.
func (me *_ToolbarButtons) Delete(index int) {
	if me.tb.Hwnd().SendMessage(co.TB_DELETEBUTTON, win.WPARAM(index), 0) == 0 {
		panic(fmt.Sprintf("TB_DELETEBUTTON \"%d\" failed.", index))
	}
}

// Enables or disables a button.
func (me *_ToolbarButtons) Enable(isEnabled bool, cmdId int) {
	ret := me.tb.Hwnd().SendMessage(co.TB_ENABLEBUTTON,
		win.WPARAM(cmdId),
		win.MAKELPARAM(uint16(util.BoolToUintptr(isEnabled)), 0))
	if ret == 0 {
		panic(fmt.Sprintf("TB_ENABLEBUTTON \"%d\" failed.", cmdId))
	}
}

// Retrieves the icon index of the button.
func (me *_ToolbarButtons) Icon(cmdId int) int {
	var tbi win.TBBUTTONINFO
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_IMAGE

	ret := me.tb.Hwnd().SendMessage(co.TB_GETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))
	if int(ret) == -1 {
		panic(fmt.Sprintf("TB_GETBUTTONINFO \"%d\" failed.", cmdId))
	}

	return int(tbi.IImage)
}

// Tells whether a button is enabled.
func (me *_ToolbarButtons) IsEnabled(cmdId int) bool {
	return me.tb.Hwnd().
		SendMessage(co.TB_ISBUTTONENABLED, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is hidden.
func (me *_ToolbarButtons) IsHidden(cmdId int) bool {
	return me.tb.Hwnd().
		SendMessage(co.TB_ISBUTTONHIDDEN, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is highlighted.
func (me *_ToolbarButtons) IsHighlighted(cmdId int) bool {
	return me.tb.Hwnd().
		SendMessage(co.TB_ISBUTTONHIGHLIGHTED, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is indeterminate.
func (me *_ToolbarButtons) IsIndeterminate(cmdId int) bool {
	return me.tb.Hwnd().
		SendMessage(co.TB_ISBUTTONINDETERMINATE, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is pressed.
func (me *_ToolbarButtons) IsPressed(cmdId int) bool {
	return me.tb.Hwnd().
		SendMessage(co.TB_ISBUTTONPRESSED, win.WPARAM(cmdId), 0) != 0
}

// Retrieves the custom data associated with the button.
func (me *_ToolbarButtons) LParam(cmdId int) win.LPARAM {
	var tbi win.TBBUTTONINFO
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_LPARAM

	ret := me.tb.Hwnd().SendMessage(co.TB_GETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))
	if int(ret) == -1 {
		panic(fmt.Sprintf("TB_GETBUTTONINFO \"%d\" failed.", cmdId))
	}

	return tbi.LParam
}

// Sets the icon index of the button.
func (me *_ToolbarButtons) SetIcon(cmdId, iconIdex int) {
	var tbi win.TBBUTTONINFO
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_IMAGE
	tbi.IImage = int32(iconIdex)

	ret := me.tb.Hwnd().SendMessage(co.TB_SETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))
	if ret == 0 {
		panic(fmt.Sprintf("TB_SETBUTTONINFO \"%d\" failed.", cmdId))
	}
}

// Sets the custom data associated with the button.
func (me *_ToolbarButtons) SetLParam(cmdId, lp win.LPARAM) {
	var tbi win.TBBUTTONINFO
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_LPARAM
	tbi.LParam = lp

	ret := me.tb.Hwnd().SendMessage(co.TB_SETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))
	if ret == 0 {
		panic(fmt.Sprintf("TB_SETBUTTONINFO \"%d\" failed.", cmdId))
	}
}

// Sets the text of the button.
func (me *_ToolbarButtons) SetText(cmdId int, text string) {
	var tbi win.TBBUTTONINFO
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_TEXT
	tbi.SetPszText(win.Str.ToNativeSlice(text))

	ret := me.tb.Hwnd().SendMessage(co.TB_SETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))
	if ret == 0 {
		panic(fmt.Sprintf("TB_SETBUTTONINFO \"%d\" \"%s\" failed.", cmdId, text))
	}
}

// Retrieves the text of the button.
func (me *_ToolbarButtons) Text(cmdId int) string {
	var buf [60]uint16 // arbitrary

	var tbi win.TBBUTTONINFO
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_TEXT
	tbi.SetPszText(buf[:])

	ret := me.tb.Hwnd().SendMessage(co.TB_GETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))
	if int(ret) == -1 {
		panic(fmt.Sprintf("TB_GETBUTTONINFO \"%d\" failed.", cmdId))
	}

	return win.Str.FromNativeSlice(buf[:])
}
