/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package main

import (
	"fmt"
	"wingows/ui"
)

func main() {
	my := MyWindow{}
	my.RunBabyRun()
}

type MyWindow struct {
	core ui.WindowMain
}

func (me *MyWindow) RunBabyRun() {
	me.core.Setup().Title = "My first window"

	// Let's choose the size of our window.
	// If the user is running Windows in a high-DPI setting, these values will be
	// automatically recalculated to match it.
	me.core.Setup().Width = 410
	me.core.Setup().Height = 320

	// Window messages are sent to us by the operational system, this is how we
	// respond to events in a Win32 application.
	// Here we add a handle to WM_CREATE message, which is sent to us right
	// before our window appears on screen:
	// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-create
	// Notice we pass a closure, which will be called every time the message is
	// sent to us. The closure receives an ui.WmCreate parameter, which contains
	// the data sent by the message.
	// Also notice this closure must return an int32. Refer to the official
	// WM_CREATE documentation to see what it means.
	me.core.OnMsg().WmCreate(func(p ui.WmCreate) int32 {
		// The parameter p, an object of ui.WmCreate type, exposes methods that
		// allows us to retrieve the data sent by the message.
		// Here we retrieve the api.CREATESTRUCT pointer, which is a pointer to
		// the data used internally to create our window.
		createStr := p.CreateStruct()

		// The width and height of our window, at the time we created it, can be
		// retrieved from the api.CREATESTRUCT object.
		ourWidth := createStr.Cx
		ourHeight := createStr.Cy

		// Here we call Hwnd() method to retrieve the underlying HWND handle of
		// our window, which exposes all HWND-related Win32 functions.
		// Take care, these are just thin Go wrappers over raw C Win32 functions,
		// and many of them can do a lot of harm! So please refer to the official
		// documentation of each function.
		// To set a new title to the window, we call SetWindowText:
		// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowtextw
		me.core.Hwnd().SetWindowText(
			fmt.Sprintf("Our window is %d x %d", ourWidth, ourHeight))

		// According to WM_CREATE documentation, we must return 0 to continue the
		// normal window creation.
		return 0
	})

	// After adding the message handlers, don't forget to create the window!
	me.core.RunAsMain()
}
