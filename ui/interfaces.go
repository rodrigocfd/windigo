/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/win"
)

type (
	// Any window with a HWND handle.
	Window interface {
		Hwnd() win.HWND
	}

	// Any child control with HWND and ID.
	Control interface {
		Window
		Id() int
	}
)

type (
	// The position of a window or child control.
	Pos struct {
		X, Y int
	}

	// The size of a window or child control.
	Size struct {
		Cx, Cy uint
	}
)
