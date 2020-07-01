/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
)

var (
	globalBaseCtrlId = co.ID(1000) // arbitrary, taken from Visual Studio resource editor
)

// Encapsulates the control ID and, if not initialized, uses an auto-incremented value.
type controlIdGuard struct {
	id co.ID // defaults to zero
}

// Optional; returns a ctrlIdGuard with a custom control ID, not using the
// default auto-incremented one.
func makeCtrlIdGuard(initialId co.ID) controlIdGuard {
	return controlIdGuard{
		id: initialId, // properly initialized
	}
}

// Returns the ID of this child window control.
// Will be initialized upon first call.
func (me *controlIdGuard) Id() co.ID {
	if me.id == 0 { // not initialized yet?
		globalBaseCtrlId++ // increments sequential global ID
		me.id = globalBaseCtrlId
	}
	return me.id
}
