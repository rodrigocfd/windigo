package ui

import (
	"fmt"
	"wingows/api"
	c "wingows/consts"
)

// Base to all child control types.
type nativeControlBase struct {
	ctrlIdGuard
	hwnd api.HWND
}

func makeNativeControlBase(ctrlId c.ID) nativeControlBase {
	return nativeControlBase{
		ctrlIdGuard: makeCtrlIdGuard(ctrlId),
	}
}

// Returns the underlying HWND handle of this native control.
func (me *nativeControlBase) Hwnd() api.HWND {
	return me.hwnd
}

func (me *nativeControlBase) create(exStyle c.WS_EX, className, title string,
	style c.WS, x, y int32, width, height uint32, parent Window) {

	if me.hwnd != 0 {
		panic(fmt.Sprintf("Trying to create %s twice.", className))
	}
	me.hwnd = api.CreateWindowEx(exStyle, className, title, style,
		x, y, width, height, parent.Hwnd(), api.HMENU(me.ctrlIdGuard.CtrlId()),
		parent.Hwnd().GetInstance(), nil)
}
