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
	// In the main entry-point function, we first instantiate our MyWindow
	// object.
	my := MyWindow{}

	// Then we create and run it.
	my.RunBabyRun()
}

// This struct represents our main window, and all the state variables that
// belong to it.
// Using a struct to isolate this data helps us to build a scalable application,
// as we create more and more windows and controls.
type MyWindow struct {
	// Our window object contains ui.WindowMain, responsible for all window
	// creation and management.
	core ui.WindowMain
}

func (me *MyWindow) RunBabyRun() {
	// The Setup() method gives you access to the values that will be used in
	// window class registration and window creation. These values are used once,
	// and further changes on them will have no effect. All values are optional.
	me.core.Setup().Title = "My first window"

	// This method is responsible to start the process of window registration,
	// window creation, and main loop kick-off. It should be the last thing you
	// do when creating your main window.
	me.core.RunAsMain()
}
