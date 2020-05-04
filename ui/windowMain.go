package ui

import (
	a "winffi/api"
)

type windowMain struct {
	windowBase
}

// NewWindowMain creates a new main application window.
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
