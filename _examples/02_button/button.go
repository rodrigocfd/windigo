/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package main

import (
	"wingows/ui"
)

func main() {
	my := MyWindow{}
	my.RunBabyRun()
}

type MyWindow struct {
	core ui.WindowMain

	// The button object is also a member of MyWindow struct, because we want to
	// interact with it during all MyWindow lifetime.
	btnClickMe ui.Button
}

func (me *MyWindow) RunBabyRun() {
	me.core.Setup().Title = "My first window"

	me.core.OnMsg().WmCreate(func(p ui.WmCreate) int32 {
		// We create the button when processing WM_CREATE message.
		// The first argument is a pointer to the parent window.
		me.btnClickMe.CreateSimple(&me.core, 20, 20, 100, "&Click me")

		// We want our button to be focused when the window shows up.
		me.btnClickMe.Hwnd().SetFocus()

		return 0
	})

	// When a button is clicked, Windows sends a WM_COMMAND message to our
	// application. The first argument below is the control ID of our button,
	// this is how we filter the specific click of our button.
	// Notice how this closure receives an ui.WmCommand parameter and returns no
	// value, differently from WM_CREATE.
	me.core.OnMsg().WmCommand(me.btnClickMe.CtrlId(), func(p ui.WmCommand) {
		// When the button is clicked, the window will have a new title.
		me.core.Hwnd().SetWindowText("This is a new title")
	})

	me.core.RunAsMain()
}
