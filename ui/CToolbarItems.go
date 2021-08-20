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
	tbb.SetIBitmap(iconIndex, imgListIndex)
	tbb.SetIString(text)

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

// Enables or disables a button.
func (me *_ToolbarItems) EnableButton(isEnabled bool, cmdId int) {
	if me.pHwnd.SendMessage(co.TB_ENABLEBUTTON,
		win.WPARAM(cmdId),
		win.LPARAM(
			win.Bytes.Make32(uint16(util.BoolToUintptr(isEnabled)), 0),
		)) == 0 {
		panic(fmt.Sprintf("TB_ENABLEBUTTON \"%d\" failed.", cmdId))
	}
}
