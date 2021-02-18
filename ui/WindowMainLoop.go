/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Runs the main application loop.
// Keeps modeless children handles.
type _WindowMainLoop struct {
	modelessChildren []win.HWND
}

// Constructor.
func _NewWindowMainLoop() *_WindowMainLoop {
	return &_WindowMainLoop{}
}

// Adds a new modeless children to be processed.
func (me *_WindowMainLoop) AddModelessChild(hModeless win.HWND) {
	me.modelessChildren = append(me.modelessChildren, hModeless)
}

// Returns the modeless children currently being processed.
func (me *_WindowMainLoop) ModelessChildren() []win.HWND {
	return me.modelessChildren
}

// Runs the main application loop.
// Will block until the loop ends.
func (me *_WindowMainLoop) RunLoop(hWnd win.HWND, hAccel win.HACCEL) int {
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
		if hAccel != 0 && hTopLevel.TranslateAccelerator(hAccel, &msg) {
			continue // message translated, no further processing is done
		}

		if hTopLevel.IsDialogMessage(&msg) {
			// Processed all keyboard actions for child controls.
			continue
		}

		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}

// Checks if the message belongs to a modeless child, and if so, processes it.
func (me *_WindowMainLoop) isModelessMsg(msg *win.MSG) bool {
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
