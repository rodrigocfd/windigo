package ui

import (
	"winffi/api"
	c "winffi/consts"
)

type Control interface {
	Window
	CtrlId() c.ID
}

type Window interface {
	Hwnd() api.HWND
}
