package ui

import (
	"unsafe"
	"winffi/api"
	c "winffi/consts"
	"winffi/parm"
)

// Main application window. Call NewWindowMain() to construct the object.
type WindowMain struct {
	windowBase
	Setup windowMainSetup
}

// Creates a new main application window.
func NewWindowMain() WindowMain {
	return WindowMain{
		windowBase: newWindowBase(),
		Setup:      newWindowMainSetup(),
	}
}

func (wnd *WindowMain) RunAsMain() {
	api.InitCommonControls()
	hInst := api.GetModuleHandle("")
	wnd.registerClass(wnd.Setup.genWndclassex(hInst))

	wnd.On.WmNcDestroy(func(p parm.WmNcDestroy) {
		api.PostQuitMessage(0)
	})

	wnd.createWindow(hInst)
	wnd.runMainLoop()
}

func (wnd *WindowMain) createWindow(hInst api.HINSTANCE) {
	globalUiFont.CreateUi()

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN)
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	hwnd := api.CreateWindowEx(wnd.Setup.ExStyle,
		wnd.Setup.ClassName, wnd.Setup.Title, wnd.Setup.Style,
		cxScreen/2-int32(wnd.Setup.Width)/2, // center window on screen
		cyScreen/2-int32(wnd.Setup.Height)/2,
		wnd.Setup.Width, wnd.Setup.Height,
		api.HWND(0), api.HMENU(0), hInst,
		unsafe.Pointer(&wnd.windowBase)) // pass pointer to windowBase object

	hwnd.ShowWindow(wnd.Setup.CmdShow)
	hwnd.UpdateWindow()
}

func (wnd *WindowMain) runMainLoop() {
	defer globalUiFont.Destroy()

	msg := api.MSG{}
	for {
		status := msg.GetMessage(api.HWND(0), 0, 0)
		if status == 0 {
			break // WM_QUIT was sent, gracefully terminate the program
		}

		if wnd.isModelessMsg() { // does this message belong to a modeless child (if any)?
			continue
		}

		// TODO haccel check !!!

		if wnd.hwnd.IsDialogMessage(&msg) { // processes all keyboard actions for our child controls
			continue
		}

		msg.TranslateMessage()
		msg.DispatchMessage()
	}
}

func (wnd *WindowMain) isModelessMsg() bool {
	return false
}
