package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ToolbarButtons struct {
	pHwnd *win.HWND
}

func (me *_ToolbarButtons) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
}

// Adds a button.
func (me *_ToolbarButtons) Add(
	imgListIndex, iconIndex, cmdId int, text string) *_ToolbarButtons {

	tbb := win.TBBUTTON{}
	tbb.IdCommand = int32(cmdId)
	tbb.FsStyle = co.BTNS_AUTOSIZE
	tbb.FsState = co.TBSTATE_ENABLED
	tbb.IString = win.Str.ToNativePtr(text)
	tbb.SetIBitmap(iconIndex, imgListIndex)

	if me.pHwnd.SendMessage(co.TB_ADDBUTTONS, 1, win.LPARAM(unsafe.Pointer(&tbb))) == 0 {
		panic(fmt.Sprintf("TB_ADDBUTTONS \"%s\" failed.", text))
	}
	me.pHwnd.SendMessage(co.TB_AUTOSIZE, 0, 0)
	return me
}

// Adds a separator.
func (me *_ToolbarButtons) AddSeparator() *_ToolbarButtons {
	tbb := win.TBBUTTON{
		FsStyle: co.BTNS_SEP,
	}

	if me.pHwnd.SendMessage(co.TB_ADDBUTTONS, 1, win.LPARAM(unsafe.Pointer(&tbb))) == 0 {
		panic("TB_ADDBUTTONS failed for separator.")
	}
	me.pHwnd.SendMessage(co.TB_AUTOSIZE, 0, 0)
	return me
}

// Changes the button icon.
func (me *_ToolbarButtons) ChangeIcon(cmdId, iconIdex int) {
	if me.pHwnd.SendMessage(co.TB_CHANGEBITMAP,
		win.WPARAM(cmdId), win.LPARAM(iconIdex)) == 0 {
		panic(fmt.Sprintf("TB_CHANGEBITMAP \"%d\" failed.", cmdId))
	}
}

// Retrieves the number of buttons.
func (me *_ToolbarButtons) Count() int {
	return int(me.pHwnd.SendMessage(co.TB_BUTTONCOUNT, 0, 0))
}

// Deletes a button.
func (me *_ToolbarButtons) Delete(index int) {
	if me.pHwnd.SendMessage(co.TB_DELETEBUTTON, win.WPARAM(index), 0) == 0 {
		panic(fmt.Sprintf("TB_DELETEBUTTON \"%d\" failed.", index))
	}
}

// Enables or disables a button.
func (me *_ToolbarButtons) Enable(isEnabled bool, cmdId int) {
	if me.pHwnd.SendMessage(co.TB_ENABLEBUTTON,
		win.WPARAM(cmdId),
		win.MAKELPARAM(uint16(util.BoolToUintptr(isEnabled)), 0),
	) == 0 {
		panic(fmt.Sprintf("TB_ENABLEBUTTON \"%d\" failed.", cmdId))
	}
}

// Tells whether a button is enabled.
func (me *_ToolbarButtons) IsEnabled(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONENABLED, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is hidden.
func (me *_ToolbarButtons) IsHidden(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONHIDDEN, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is highlighted.
func (me *_ToolbarButtons) IsHighlighted(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONHIGHLIGHTED, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is indeterminate.
func (me *_ToolbarButtons) IsIndeterminate(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONINDETERMINATE, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is pressed.
func (me *_ToolbarButtons) IsPressed(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONPRESSED, win.WPARAM(cmdId), 0) != 0
}

// Sets the text of the button.
func (me *_ToolbarButtons) SetText(cmdId int, text string) {
	tbi := win.TBBUTTONINFO{}
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_TEXT
	tbi.SetPszText(win.Str.ToNativeSlice(text))

	if me.pHwnd.SendMessage(co.TB_SETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi))) == 0 {
		panic(fmt.Sprintf("TB_SETBUTTONINFO \"%d\" \"%s\" failed.", cmdId, text))
	}
}

// Retrieves the text of the button.
func (me *_ToolbarButtons) Text(cmdId int) string {
	buf := [60]uint16{} // arbitrary

	tbi := win.TBBUTTONINFO{}
	tbi.SetCbSize()
	tbi.DwMask = co.TBIF_TEXT
	tbi.SetPszText(buf[:])

	if int(me.pHwnd.SendMessage(co.TB_GETBUTTONINFO,
		win.WPARAM(cmdId), win.LPARAM(unsafe.Pointer(&tbi)))) == -1 {
		panic(fmt.Sprintf("TB_SETBUTTONINFO \"%d\" failed.", cmdId))
	}

	return win.Str.FromNativeSlice(buf[:])
}
