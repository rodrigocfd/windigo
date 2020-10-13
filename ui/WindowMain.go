/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"runtime/debug"
	"windigo/co"
	"windigo/win"
)

var (
	_globalUiFont = Font{} // created/freed in RunAsMain()
)

// Main application window.
//
// Allows message and notification handling.
type WindowMain struct {
	_WindowBase
	setup            _WindowSetupMain
	modelessChildren []win.HWND
	childPrevFocus   win.HWND // when window is inactivated
}

// Exposes parameters that will be used to create the window.
func (me *WindowMain) Setup() *_WindowSetupMain {
	if me.Hwnd() != 0 {
		panic("Cannot change setup after the window was created.")
	}
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the main window and runs the main application loop.
func (me *WindowMain) RunAsMain() int {
	defer func() {
		// Recover from any panic within GUI thread.
		// Panics in other threads can't be recovered.
		if r := recover(); r != nil {
			msg, ok := r.(string)
			if ok {
				msg = fmt.Sprintf("A panic has occurred:\n\n%s", msg)
			} else {
				msg = "An unspecified panic occurred."
			}
			win.HWND(0).MessageBox(msg+"\n\n"+string(debug.Stack()),
				"Panic", co.MB_ICONERROR)
		}

		// Free resources.
		me.setup.AcceleratorTable.Destroy()
		_globalUiFont.Destroy()
	}()

	if win.IsWindowsVistaOrGreater() {
		win.SetProcessDPIAware()
	}
	win.InitCommonControls()

	me.setup.initOnce() // guard
	hInst := win.GetModuleHandle("")
	me._WindowBase.registerClass(me.setup.genWndclassex(hInst))

	_globalUiFont.CreateUi() // create global font to be applied everywhere

	me.defaultMessageHandling()

	pos, size := me.setup.calcCoords()
	me._WindowBase.createWindow("WindowMain", me.setup.ExStyle,
		me.setup.ClassName, me.setup.Title, me.setup.Style,
		pos, size, nil, me.setup.MainMenu.Hmenu(), hInst)

	me.Hwnd().ShowWindow(me.setup.CmdShow)
	me.Hwnd().UpdateWindow()
	return me.runMainLoop()
}

func (me *WindowMain) defaultMessageHandling() {
	me.OnMsg().WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})

	me.OnMsg().WmSetFocus(func(hwndLosingFocus win.HWND) {
		if me.Hwnd() == win.GetFocus() {
			// If window receives focus, delegate to first child.
			me.Hwnd().
				GetNextDlgTabItem(win.HWND(0), false).
				SetFocus()
		}
	})

	me.OnMsg().WmActivate(func(p WmActivate) {
		// https://devblogs.microsoft.com/oldnewthing/20140521-00/?p=943
		if !p.IsMinimized() {
			if p.Event() == co.WA_INACTIVE {
				curFocus := win.GetFocus()
				if curFocus != 0 && me.Hwnd().IsChild(curFocus) {
					me.childPrevFocus = curFocus // save previously focused control
				}
			} else if me.childPrevFocus != 0 {
				me.childPrevFocus.SetFocus() // put focus back
			}
		}
	})
}

func (me *WindowMain) runMainLoop() int {
	msg := win.MSG{}
	for {
		if win.GetMessage(&msg, win.HWND(0), 0, 0) == 0 {
			// WM_QUIT was sent, gracefully terminate the program.
			// If it returned -1, it will simply panic.
			// WParam has the program exit code.
			// https://docs.microsoft.com/en-us/windows/win32/winmsg/using-messages-and-message-queues
			return int(msg.WParam)
		}

		if me.isModelessMsg(&msg) { // does this message belong to a modeless child (if any)?
			// http://www.winprog.org/tutorial/modeless_dialogs.html
			continue
		}

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		hTopLevel := msg.HWnd.GetAncestor(co.GA_ROOT)

		// If we have an accelerator table, try to translate the message.
		if me.setup.AcceleratorTable.Haccel() != 0 &&
			hTopLevel.TranslateAccelerator(
				me.setup.AcceleratorTable.Haccel(), &msg) {
			// Message translated, no further processing is done.
			continue
		}

		if hTopLevel.IsDialogMessage(&msg) {
			// Processed all keyboard actions for child controls.
			continue
		}

		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}

func (me *WindowMain) isModelessMsg(msg *win.MSG) bool {
	for _, hChild := range me.modelessChildren { // check all modeless HWNDs
		if hChild == 0 || !hChild.IsWindow() {
			continue // skip invalid HWND
		}
		if hChild.IsDialogMessage(msg) {
			return true // it was a message for this modeless, it was processed and we're done
		}
	}
	return false // the message wasn't for any of the modeless HWNDs
}
