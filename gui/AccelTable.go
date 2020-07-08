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
	accels map[string]win.ACCEL
	built  bool
}

// Adds a new accelerator, with an auto-generated command ID.
// If passing a character code, use uppercase.
func (me *AccelTable) Add(textId string,
	vKey co.VK, flags co.ACCELF) *AccelTable {

	newCmdId := controlId{}
	return me.AddWithCmdId(textId, vKey, flags, newCmdId.Id())
}

// Adds a new accelerator, with a specific command ID.
// If passing a character code, use uppercase.
func (me *AccelTable) AddWithCmdId(textId string,
	vKey co.VK, flags co.ACCELF, cmdId int32) *AccelTable {

	if me.accels == nil {
		me.accels = make(map[string]win.ACCEL)
	}
	me.accels[textId] = win.ACCEL{
		FVirt: flags | co.ACCELF_VIRTKEY,
		Key:   vKey,
		Cmd:   uint16(cmdId),
	}
	return me
}

// Builds the accelerator table resource and returns its handle.
// It must be freed with DestroyAcceleratorTable().
func (me *AccelTable) Build() win.HACCEL {
	if me.built {
		panic("Accelerator table already built.")
	}
	me.built = true

	accelSlice := make([]win.ACCEL, 0, len(me.accels))
	for _, accel := range me.accels { // make slice from map values
		accelSlice = append(accelSlice, accel)
	}
	return win.CreateAcceleratorTable(accelSlice)
}
