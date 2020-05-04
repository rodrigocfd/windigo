package ui

import (
	"winffi/api"
	c "winffi/consts"
)

type IControl interface {
	IWindow
	CtrlId() c.ID
}

type IWindow interface {
	Hwnd() api.HWND
}
