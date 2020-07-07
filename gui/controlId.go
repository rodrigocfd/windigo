/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
)

var (
	globalBaseCtrlId int32 = 1000 // arbitrary, taken from Visual Studio resource editor
)

// Encapsulates the control ID and, if not initialized, uses an auto-incremented value.
type controlId struct {
	id int32 // defaults to zero
}

// Sets a custom ID. Must be done before using the ID, otherwise an automatic
// ID will be set.
func (me *controlId) SetId(customId int32) {
	if me.id != 0 {
		panic(fmt.Sprintf("Can't set control ID %d, already set.", customId))
	}
	me.id = customId
}

// Returns the ID of this child window control.
// If not custom set, will be initialized upon first call.
func (me *controlId) Id() int32 {
	if me.id == 0 { // not initialized yet?
		globalBaseCtrlId++ // increments sequential global ID
		me.id = globalBaseCtrlId
	}
	return me.id
}
