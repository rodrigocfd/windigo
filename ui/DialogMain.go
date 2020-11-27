/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Main application dialog.
type DialogMain struct {
	*_DialogBase
	mainLoop   *_WindowMainLoop
	iconId     int
	accelTblId int
}

// Constructor.
// You can pass zero to iconId and accelTblId.
func NewDialogMain(dlgId, iconId, accelTblId int) *DialogMain {
	me := DialogMain{
		_DialogBase: _NewDialogBase(dlgId),
		mainLoop:    _NewWindowMainLoop(),
		iconId:      iconId,
		accelTblId:  accelTblId,
	}

	me.defaultMessageHandling()
	return &me
}

// Creates the main dialog and runs the main application loop.
// Will block until the window is closed.
func (me *DialogMain) RunAsMain() int {
	defer func() {
		// Recover from any panic within GUI thread.
		// Panics in other threads can't be recovered.
		if r := recover(); r != nil {
			_global.TreatPanic(r)
		}
		_global.uiFont.Destroy() // free resource
	}()

	if win.IsWindowsVistaOrGreater() {
		win.SetProcessDPIAware()
	}
	win.InitCommonControls()

	hInst := win.GetModuleHandle("")
	me._DialogBase.createDialogParam(hInst, nil)

	hAccel := win.HACCEL(0)
	if me.accelTblId != 0 {
		// An accelerator table loaded from resource is automatically freed by the system.
		hAccel = hInst.LoadAccelerators(int32(me.accelTblId))
	}

	me.setIconIfAny(hInst)
	me.Hwnd().ShowWindow(co.SW_SHOW)
	return me.mainLoop.RunLoop(me.Hwnd(), hAccel)
}

// Adds the messages which have a default processing.
func (me *DialogMain) defaultMessageHandling() {
	me.On().WmClose(func() {
		me.Hwnd().DestroyWindow()
	})

	me.On().WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})
}

func (me *DialogMain) setIconIfAny(hInst win.HINSTANCE) {
	// If an icon ID was specified, load it from the resource.
	// Resource icons are automatically released by the system.
	if me.iconId != 0 {
		me.Hwnd().SendMessage(co.WM_SETICON,
			win.WPARAM(co.ICON_SZ_SMALL),
			win.LPARAM(
				hInst.LoadImage(int32(me.iconId), co.IMAGE_ICON,
					16, 16, co.LR_DEFAULTCOLOR),
			))

		me.Hwnd().SendMessage(co.WM_SETICON,
			win.WPARAM(co.ICON_SZ_BIG), win.LPARAM(
				hInst.LoadImage(int32(me.iconId), co.IMAGE_ICON,
					32, 32, co.LR_DEFAULTCOLOR),
			))
	}
}
