/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/win"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	Id() int32
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() win.HWND
}
