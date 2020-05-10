package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

// Main application window.
type WindowMain struct {
	windowBase
	Setup windowMainSetup
}

func NewWindowMain() *WindowMain {
	return &WindowMain{
		windowBase: makeWindowBase(),
		Setup:      makeWindowMainSetup(),
	}
}

// Creates the main window and runs the main application loop.
func (me *WindowMain) RunAsMain() {
	if api.IsWindowsVistaOrGreater() {
		api.SetProcessDPIAware()
	}

	api.InitCommonControls()
	hInst := api.GetModuleHandle("")
	me.windowBase.registerClass(me.Setup.genWndClassEx(hInst))

	me.windowBase.On.WmNcDestroy(func() { // default WM_NCDESTROY handling
		api.PostQuitMessage(0)
	})

	me.createWindow(hInst)
	me.runMainLoop()
}

func (me *WindowMain) createWindow(hInst api.HINSTANCE) {
	globalUiFont.CreateUi()

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN)
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	api.CreateWindowEx(me.Setup.ExStyle, // hwnd member is saved in WM_NCCREATE processing
		me.Setup.ClassName, me.Setup.Title, me.Setup.Style,
		cxScreen/2-int32(me.Setup.Width)/2, // center window on screen
		cyScreen/2-int32(me.Setup.Height)/2,
		me.Setup.Width, me.Setup.Height,
		api.HWND(0), api.HMENU(0), hInst,
		unsafe.Pointer(&me.windowBase)) // pass pointer to windowBase object

	me.windowBase.Hwnd().ShowWindow(me.Setup.CmdShow)
	me.windowBase.Hwnd().UpdateWindow()
}

func (me *WindowMain) SetTitle(title string) *WindowMain {
	me.windowBase.Hwnd().SetWindowText(title)
	return me
}

func (me *WindowMain) Title() string {
	return me.windowBase.Hwnd().GetWindowText()
}

func (me *WindowMain) runMainLoop() {
	defer globalUiFont.Destroy()
	me.windowBase.On.loopStarted = true

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
