/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Main application window.
type WindowMain struct {
	windowBase
	setup windowMainSetup
}

// Exposes parameters that will be used to create the window.
func (me *WindowMain) Setup() *windowMainSetup {
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the main window and runs the main application loop.
func (me *WindowMain) RunAsMain() {
	if api.IsWindowsVistaOrGreater() {
		api.SetProcessDPIAware()
	}
	api.InitCommonControls()

	me.setup.initOnce() // guard
	hInst := api.GetModuleHandle("")
	me.windowBase.registerClass(me.setup.genWndClassEx(hInst))

	globalUiFont.CreateUi() // create global font to be applied everywhere

	me.windowBase.OnMsg().WmNcDestroy(func() { // default WM_NCDESTROY handling
		api.PostQuitMessage(0)
	})

	cxScreen := api.GetSystemMetrics(c.SM_CXSCREEN) // retrieve screen size
	cyScreen := api.GetSystemMetrics(c.SM_CYSCREEN)

	me.windowBase.createWindow("WindowMain", me.setup.ExStyle,
		me.setup.ClassName, me.setup.Title, me.setup.Style,
		cxScreen/2-int32(me.setup.Width)/2, // center window on screen
		cyScreen/2-int32(me.setup.Height)/2,
		me.setup.Width, me.setup.Height, nil, me.setup.HMenu, hInst)

	me.windowBase.Hwnd().ShowWindow(me.setup.CmdShow)
	me.windowBase.Hwnd().UpdateWindow()

	me.runMainLoop()
}

func (me *WindowMain) runMainLoop() {
	defer globalUiFont.Destroy()
	me.windowBase.wndMsg.loopStarted = true

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
