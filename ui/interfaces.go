package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	CtrlId() c.ID
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() api.HWND
}

//------------------------------------------------------------------------------

var baseId = c.ID(1000) // arbitrary, taken from Visual Studio resource editor

// Returns the next automatically incremented control ID.
func nextAutoCtrlId() c.ID {
	baseId += 1
	return baseId
}

// Enables or disables many controls at once.
func EnableControls(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}
