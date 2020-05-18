/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

// Manages window subclassing for controls.
type windowSubclass struct {
	wndMsg windowMsg
}

// Exposes all the window messages the can be handled through subclassing.
func (me *windowBase) OnSubclass() *windowMsg {
	return &me.wndMsg
}
