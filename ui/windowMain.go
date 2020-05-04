package ui

import (
	a "winffi/api"
)

// Main application window.
type windowMain struct {
	windowBase
}

// Creates a new main application window.
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
