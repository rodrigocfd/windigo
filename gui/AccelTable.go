/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
	"wingows/win"
)

// Helps building an accelerator table.
type AccelTable struct {
	accels []win.ACCEL
	built  bool
}

// Adds a new accelerator, with a specific command ID.
// If passing a character code, use uppercase.
func (me *AccelTable) Add(textId string,
	vKey co.VK, flags co.ACCELF, cmdId int32) *AccelTable {

	me.accels = append(me.accels, win.ACCEL{
		FVirt: flags | co.ACCELF_VIRTKEY,
		Key:   vKey,
		Cmd:   uint16(cmdId),
	})
	return me
}

// Builds the accelerator table resource and returns its handle.
// It must be freed with DestroyAcceleratorTable().
func (me *AccelTable) Build() win.HACCEL {
	if me.built {
		panic("Accelerator table already built.")
	}
	me.built = true

	if len(me.accels) == 0 { // no accelerators added
		return win.HACCEL(0)
	}
	return win.CreateAcceleratorTable(me.accels)
}
