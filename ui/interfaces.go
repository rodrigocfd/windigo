/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/win"
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
