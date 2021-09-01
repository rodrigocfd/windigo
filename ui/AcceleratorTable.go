package ui

import (
	"unicode"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Native accelerator table resource.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/learnwin32/accelerator-tables
type AcceleratorTable interface {
	isAcceleratorTable() // prevent public implementation

	// Adds a new character accelerator, with a specific command ID.
	AddChar(character rune, modifiers co.ACCELF, cmdId int) AcceleratorTable

	// Adds a new virtual key accelerator, with a specific command ID.
	AddKey(vKey co.VK, modifiers co.ACCELF, cmdId int) AcceleratorTable

	// Free the resources.
	Destroy()

	// Builds the accelerator table once, and returns the HACCEL handle.
	//
	// Further accelerator additions will panic after this call.
	Haccel() win.HACCEL
}

//------------------------------------------------------------------------------

type _AccelTable struct {
	accels []win.ACCEL
	hAccel win.HACCEL
}

// Creates a new AcceleratorTable.
//
// ⚠️ You must defer AcceleratorTable.Destroy().
func NewAcceleratorTable() AcceleratorTable {
	return &_AccelTable{
		accels: nil,
		hAccel: win.HACCEL(0),
	}
}

func (me *_AccelTable) isAcceleratorTable() {}

func (me *_AccelTable) AddChar(
	character rune, modifiers co.ACCELF, cmdId int) AcceleratorTable {

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

func (me *_AccelTable) AddKey(
	vKey co.VK, modifiers co.ACCELF, cmdId int) AcceleratorTable {

	if me.hAccel != 0 {
		panic("Cannot add virtual key after accelerator table was built.")
	}

	me.accels = append(me.accels, win.ACCEL{
		FVirt: modifiers | co.ACCELF_VIRTKEY,
		Key:   vKey,
		Cmd:   uint16(cmdId),
	})
	return me
}

func (me *_AccelTable) Destroy() {
	if me.hAccel != 0 {
		me.hAccel.DestroyAcceleratorTable()
		me.hAccel = win.HACCEL(0)
	}
}

func (me *_AccelTable) Haccel() win.HACCEL {
	if me.hAccel == 0 && len(me.accels) > 0 { // not created yet, and has items?
		me.hAccel = win.CreateAcceleratorTable(me.accels)
	}
	return me.hAccel
}
