/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

type _ControlUtilT struct{}

// General child control utilities.
var ControlUtil _ControlUtilT

// Enables or disables many controls at once.
func (_ControlUtilT) Enable(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}
