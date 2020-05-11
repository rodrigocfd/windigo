package ui

import (
	c "wingows/consts"
)

var baseId = c.ID(1000) // arbitrary, taken from Visual Studio resource editor

// Ensures that control ID will be properly initialized before any call.
type ctrlIdGuard struct {
	ctrlId c.ID // defaults to zero
}

func makeCtrlIdGuard(initialId c.ID) ctrlIdGuard {
	return ctrlIdGuard{
		ctrlId: initialId, // properly initialized
	}
}

// Returns the ID of this child window control.
func (me *ctrlIdGuard) CtrlId() c.ID {
	if me.ctrlId == 0 { // not initialized yet?
		baseId += 1 // increments sequential global ID
		me.ctrlId = baseId
	}
	return me.ctrlId
}
