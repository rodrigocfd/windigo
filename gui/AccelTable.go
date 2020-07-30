/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unicode"
	"wingows/co"
	"wingows/win"
)

// Helps building an accelerator table.
type AccelTable struct {
	accels []win.ACCEL
	hAccel win.HACCEL
}

// Adds a new character accelerator, with a specific command ID.
// Call Build() after all commands are added.
func (me *AccelTable) AddChar(
	character rune, modifiers co.ACCELF, cmdId int32) *AccelTable {

	if me.hAccel != 0 {
		panic("Cannot add character after accelerator table was built.")
	}

	me.accels = append(me.accels, win.ACCEL{
		FVirt: modifiers | co.ACCELF_VIRTKEY,
		Key:   co.VK(unicode.ToUpper(character)),
		Cmd:   uint16(cmdId),
	})
	return me
}

// Adds a new key accelerator, with a specific command ID.
// Call Build() after all commands are added.
func (me *AccelTable) AddKey(
	vKey co.VK, modifiers co.ACCELF, cmdId int32) *AccelTable {

	if me.hAccel != 0 {
		panic("Cannot add key after accelerator table was built.")
	}

	me.accels = append(me.accels, win.ACCEL{
		FVirt: modifiers | co.ACCELF_VIRTKEY,
		Key:   vKey,
		Cmd:   uint16(cmdId),
	})
	return me
}

// Builds the HACCEL from the ACCEL array, making it ready to use.
func (me *AccelTable) Build() *AccelTable {
	if me.hAccel == 0 && len(me.accels) > 0 { // build only once
		me.hAccel = win.CreateAcceleratorTable(me.accels)
	}
	return me
}

// Accelerator tables must be destroyed manually.
func (me *AccelTable) Destroy() {
	if me.hAccel != 0 {
		me.hAccel.DestroyAcceleratorTable()
		me.hAccel = 0
	}
}

// Returns the HACCEL handle.
func (me *AccelTable) Haccel() win.HACCEL {
	return me.hAccel
}
