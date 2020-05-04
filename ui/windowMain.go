package ui

import (
	"winffi/api"
	"winffi/parm"
)

// Main application window.
type windowMain struct {
	windowBase
}

// Creates a new main application window.
func NewWindowMain(className string) *windowMain {
	us := windowMain{
		newWindowBase(),
	}
	us.Wcx.LpszClassName = api.ToUtf16Ptr(className)
	return &us
}

func (wnd *windowMain) RunAsMain() {
	api.InitCommonControls()

	hInst := api.GetModuleHandle("")
	wnd.registerClass(hInst)

	wnd.On.Wm.NcDestroy(func(p parm.WmNcDestroy) {
		api.PostQuitMessage(0)
	})

	wnd.createWindow(hInst)
	wnd.runMainLoop()
}

func (wnd *windowMain) createWindow(hInst api.HINSTANCE) {

}

func (wnd *windowMain) runMainLoop() {

}
