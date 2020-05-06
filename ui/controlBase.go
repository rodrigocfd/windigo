package ui

import (
	"fmt"
	"gowinui/api"
	c "gowinui/consts"
)

// Base to all child control types.
type controlBase struct {
	hwnd   api.HWND
	ctrlId c.ID
}

func makeControlBase() controlBase {
	return makeControlBaseWithId(NextAutoCtrlId())
}

func makeControlBaseWithId(ctrlId c.ID) controlBase {
	return controlBase{
		hwnd:   api.HWND(0),
		ctrlId: ctrlId,
	}
}

func (me *controlBase) create(exStyle c.WS_EX, className, title string,
	style c.WS, x, y int32, width, height uint32, parent Window) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}
	me.hwnd = api.CreateWindowEx(exStyle, className, title, style,
		x, y, width, height, parent.Hwnd(), api.HMENU(me.ctrlId),
		parent.Hwnd().GetInstance(), nil)
}

// Returns the control ID of this child window control.
func (me *controlBase) CtrlId() c.ID {
	return me.ctrlId
}

// Returns the underlying HWND handle of this window.
func (me *controlBase) Hwnd() api.HWND {
	return me.hwnd
}
