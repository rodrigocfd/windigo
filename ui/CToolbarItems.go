package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _ToolbarItems struct {
	pHwnd *win.HWND
}

func (me *_ToolbarItems) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
}

// Adds a button.
func (me *_ToolbarItems) AddButton(
	imgListIndex, iconIndex, cmdId int, text string) {

	tbb := win.TBBUTTON{}
	tbb.IdCommand = int32(cmdId)
	tbb.FsStyle = co.BTNS(co.TBSTYLE_AUTOSIZE)
	tbb.FsState = co.TBSTATE_ENABLED
	tbb.IString = win.Str.ToUint16Ptr(text)
	tbb.SetIBitmap(iconIndex, imgListIndex)

	if me.pHwnd.SendMessage(co.TB_ADDBUTTONS, 1, win.LPARAM(unsafe.Pointer(&tbb))) == 0 {
		panic(fmt.Sprintf("TB_ADDBUTTONS \"%s\" failed.", text))
	}
	me.pHwnd.SendMessage(co.TB_AUTOSIZE, 0, 0)
}

// Adds a separator.
func (me *_ToolbarItems) AddSeparator() {
	tbb := win.TBBUTTON{
		FsStyle: co.BTNS_SEP,
	}

	if me.pHwnd.SendMessage(co.TB_ADDBUTTONS, 1, win.LPARAM(unsafe.Pointer(&tbb))) == 0 {
		panic("TB_ADDBUTTONS failed for separator.")
	}
	me.pHwnd.SendMessage(co.TB_AUTOSIZE, 0, 0)
}

// Retrieves the number of buttons.
func (me *_ToolbarItems) Count() int {
	return int(me.pHwnd.SendMessage(co.TB_BUTTONCOUNT, 0, 0))
}

// Deletes a button.
func (me *_ToolbarItems) DeleteButton(index int) {
	if me.pHwnd.SendMessage(co.TB_DELETEBUTTON, win.WPARAM(index), 0) == 0 {
		panic(fmt.Sprintf("TB_DELETEBUTTON \"%d\" failed.", index))
	}
}

// Enables or disables a button.
func (me *_ToolbarItems) EnableButton(isEnabled bool, cmdId int) {
	if me.pHwnd.SendMessage(co.TB_ENABLEBUTTON,
		win.WPARAM(cmdId),
		win.MAKELPARAM(uint16(util.BoolToUintptr(isEnabled)), 0),
	) == 0 {
		panic(fmt.Sprintf("TB_ENABLEBUTTON \"%d\" failed.", cmdId))
	}
}

// Retrieves information about a button.
func (me *_ToolbarItems) GetButton(index int, info *win.TBBUTTON) {
	if me.pHwnd.SendMessage(co.TB_GETBUTTON,
		win.WPARAM(index), win.LPARAM(unsafe.Pointer(info)),
	) == 0 {
		panic(fmt.Sprintf("TB_GETBUTTON \"%d\" failed.", index))
	}
}

// Tells whether a button is enabled.
func (me *_ToolbarItems) IsButtonEnabled(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONENABLED, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is hidden.
func (me *_ToolbarItems) IsButtonHidden(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONHIDDEN, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is highlighted.
func (me *_ToolbarItems) IsButtonHighlighted(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONHIGHLIGHTED, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is indeterminate.
func (me *_ToolbarItems) IsButtonIndeterminate(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONINDETERMINATE, win.WPARAM(cmdId), 0) != 0
}

// Tells whether a button is pressed.
func (me *_ToolbarItems) IsButtonPressed(cmdId int) bool {
	return me.pHwnd.SendMessage(co.TB_ISBUTTONPRESSED, win.WPARAM(cmdId), 0) != 0
}
