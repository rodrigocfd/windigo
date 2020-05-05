package ui

import (
	"unsafe"
	"winffi/api"
	c "winffi/consts"
	"winffi/parm"
)

// Main application window.
type WindowMain struct {
	windowBase
	Setup windowMainSetup
}

func MakeWindowMain() WindowMain {
	return WindowMain{
		windowBase: makeWindowBase(),
		Setup:      makeWindowMainSetup(),
	}
}

func (me *WindowMain) RunAsMain() {
	api.InitCommonControls()
	hInst := api.GetModuleHandle("")
	me.registerClass(me.Setup.genWndclassex(hInst))

	me.On.WmNcDestroy(func(p parm.WmNcDestroy) {
		api.PostQuitMessage(0)
	})

	me.createWindow(hInst)
	me.runMainLoop()
}

func (me *WindowMain) createWindow(hInst api.HINSTANCE) {
	globalUiFont.CreateUi()

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN)
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	hwnd := api.CreateWindowEx(me.Setup.ExStyle,
		me.Setup.ClassName, me.Setup.Title, me.Setup.Style,
		cxScreen/2-int32(me.Setup.Width)/2, // center window on screen
		cyScreen/2-int32(me.Setup.Height)/2,
		me.Setup.Width, me.Setup.Height,
		api.HWND(0), api.HMENU(0), hInst,
		unsafe.Pointer(&me.windowBase)) // pass pointer to windowBase object

	hwnd.ShowWindow(me.Setup.CmdShow)
	hwnd.UpdateWindow()
}

func (me *WindowMain) runMainLoop() {
	defer globalUiFont.Destroy()

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

		if me.hwnd.IsDialogMessage(&msg) { // processes all keyboard actions for our child controls
			continue
		}

		msg.TranslateMessage()
		msg.DispatchMessage()
	}
}

func (me *WindowMain) isModelessMsg() bool {
	return false
}
