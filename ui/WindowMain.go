package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Main application window.
type WindowMain struct {
	windowBase
	Setup windowMainSetup // Parameters that will be used to create the window.
}

func NewWindowMain() *WindowMain {
	me := WindowMain{
		windowBase: makeWindowBase(),
		Setup:      makeWindowMainSetup(),
	}

	me.windowBase.OnMsg.WmNcDestroy(func() { // default WM_NCDESTROY handling
		api.PostQuitMessage(0)
	})

	return &me
}

// Creates the main window and runs the main application loop.
func (me *WindowMain) RunAsMain() {
	if api.IsWindowsVistaOrGreater() {
		api.SetProcessDPIAware()
	}

	api.InitCommonControls()
	hInst := api.GetModuleHandle("")
	me.windowBase.registerClass(me.Setup.genWndClassEx(hInst))

	globalUiFont.CreateUi() // create global font to be applied everywhere

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN) // retrieve screen size
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	me.windowBase.createWindow(me.Setup.ExStyle, me.Setup.ClassName,
		me.Setup.Title, me.Setup.Style,
		cxScreen/2-int32(me.Setup.Width)/2, // center window on screen
		cyScreen/2-int32(me.Setup.Height)/2,
		me.Setup.Width, me.Setup.Height, nil, me.Setup.HMenu, hInst)

	me.windowBase.Hwnd().ShowWindow(me.Setup.CmdShow)
	me.windowBase.Hwnd().UpdateWindow()

	me.runMainLoop()
}

func (me *WindowMain) runMainLoop() {
	defer globalUiFont.Destroy()
	me.windowBase.OnMsg.loopStarted = true

	msg := api.MSG{}
	for {
		status := msg.GetMessage(api.HWND(0), 0, 0)
		if status == 0 {
			break // WM_QUIT was sent, gracefully terminate the program
		}

		if me.isModelessMsg() { // does this message belong to a modeless child (if any)?
			continue
		}

		// TODO haccel check !!!

		if me.windowBase.Hwnd().IsDialogMessage(&msg) { // processes all keyboard actions for our child controls
			continue
		}

		msg.TranslateMessage()
		msg.DispatchMessage()
	}
}

func (me *WindowMain) isModelessMsg() bool {
	return false
}
