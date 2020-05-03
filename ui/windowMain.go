package ui

import (
	a "winffi/api"
)

type windowMain struct {
	windowBase
}

func NewWindowMain() *windowMain {
	return &windowMain{
		newWindowBase(),
	}
}

func (wnd *windowMain) RunAsMain() {
	a.InitCommonControls()

	hInst := a.GetModuleHandle("")
	wnd.registerClass(hInst)

}
