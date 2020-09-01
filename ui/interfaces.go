/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/win"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	Id() int
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() win.HWND
}
