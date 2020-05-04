package ui

import (
	c "winffi/consts"
)

var baseId = c.ID(1000) // arbitrary, taken from Visual Studio resource editor

// NextAutoCtrlId returns the next automatically incremented control ID.
func NextAutoCtrlId() c.ID {
	baseId += 1
	return baseId
}
