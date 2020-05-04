package ui

import (
	"winffi/api"
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
	api.InitCommonControls()

	hInst := api.GetModuleHandle("")
	wnd.registerClass(hInst)

}
