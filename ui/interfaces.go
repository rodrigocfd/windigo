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
	// Any window.
	Window interface {
		Hwnd() win.HWND
	}

	// Any window which can have child controls.
	Parent interface {
		Window
		On() *_EventsWmCmdNfy
	}

	// Any child control.
	Control interface {
		Window
		CtrlId() int
	}
)

type (
	// The position of a window or child control.
	Pos struct {
		X, Y int
	}

	// The size of a window or child control.
	Size struct {
		Cx, Cy int
	}
)

func (me Pos) equals(other Pos) bool   { return me.X == other.X && me.Y == other.Y }
func (me Size) equals(other Size) bool { return me.Cx == other.Cy && me.Cy == other.Cy }
