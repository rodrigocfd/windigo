/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unicode"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Native accelerator table resource.
// https://docs.microsoft.com/en-us/windows/win32/learnwin32/accelerator-tables
type AccelTable struct {
	accels []win.ACCEL
	hAccel win.HACCEL
}

// Adds a new character accelerator, with a specific command ID.
func (me *AccelTable) AddChar(
	character rune, modifiers co.ACCELF, cmdId int) *AccelTable {

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
func (me *AccelTable) AddKey(
	vKey co.VK, modifiers co.ACCELF, cmdId int) *AccelTable {

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

// Calls DestroyAcceleratorTable() to free the resources.
func (me *AccelTable) Destroy() {
	if me.hAccel != 0 {
		me.hAccel.DestroyAcceleratorTable()
		me.hAccel = 0
	}
}

// Buils the accelerator table once, and returns the HACCEL handle.
//
// Further accelerator additions will panic after this call.
func (me *AccelTable) Haccel() win.HACCEL {
	if me.hAccel == 0 && len(me.accels) > 0 { // build resource once
		me.hAccel = win.CreateAcceleratorTable(me.accels)
	}
	return me.hAccel
}
