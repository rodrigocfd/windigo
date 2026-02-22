//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type _StatusBarPartData struct {
	sizePixels   int
	resizeWeight int
}

func (me *_StatusBarPartData) IsFixedWidth() bool {
	return me.resizeWeight == 0
}

// The parts collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [StatusBar].
type StatusBarPartCollection struct {
	owner      *StatusBar
	partsData  []_StatusBarPartData
	rightEdges []int32 // buffer to speed up ResizeToFitParent() calls
}

func (me *StatusBarPartCollection) addParts(parts []_StatusBarOptPart) {
	for _, part := range parts {
		if part.width < 0 {
			panic("StatusBar part width cannot be negative.")
		} else if part.flex < 0 {
			panic("StatusBar part flex cannot be negative.")
		}

		me.partsData = append(me.partsData, _StatusBarPartData{
			sizePixels:   part.width,
			resizeWeight: part.flex,
		})
	}

	hParent, _ := me.owner.Hwnd().GetParent()
	rc, _ := hParent.GetClientRect()
	me.resizeToFitParent(WmSize{ // force the creation of the parts, so we can set text
		Raw: Wm{
			WParam: win.WPARAM(co.SIZE_REQ_RESTORED),
			LParam: win.MAKELPARAM(uint16(rc.Right-rc.Left), 0),
		},
	})

	for i, part := range parts {
		me.owner.Parts.Get(i).SetText(part.text)
	}
}

func (me *StatusBarPartCollection) resizeToFitParent(parm WmSize) {
	if parm.Request() == co.SIZE_REQ_MINIMIZED || me.owner.hWnd == 0 {
		return
	}
	me.owner.hWnd.SendMessage(co.WM_SIZE, 0, 0) // tell status bar to fit parent

	if len(me.partsData) == 0 {
		return // no parts added, nothing else to do
	}

	cx := int(parm.ClientAreaSize().Cx) // available width

	totalWeight := 0 // total weight of all variable-width parts
	cxVariable := cx // total width to be divided among variable-width parts
	for i := range me.partsData {
		if me.partsData[i].IsFixedWidth() {
			cxVariable -= me.partsData[i].sizePixels
		} else {
			totalWeight += me.partsData[i].resizeWeight
		}
	}

	cxTotal := cx
	for i := len(me.partsData) - 1; i >= 0; i-- { // fill right edges array with the right edge of each part
		me.rightEdges[i] = int32(cxTotal)
		if me.partsData[i].IsFixedWidth() {
			cxTotal -= me.partsData[i].sizePixels
		} else {
			cxTotal -= (cxVariable / totalWeight) * me.partsData[i].resizeWeight
		}
	}
	me.owner.hWnd.SendMessage(co.SB_SETPARTS,
		win.WPARAM(int32(len(me.rightEdges))),
		win.LPARAM(unsafe.Pointer(&me.rightEdges[0])))
}

// Returns all parts.
func (me *StatusBarPartCollection) All() []StatusBarPart {
	nParts := me.Count()
	parts := make([]StatusBarPart, 0, nParts)
	for i := 0; i < nParts; i++ {
		parts = append(parts, me.Get(i))
	}
	return parts
}

// Returns the number of parts.
func (me *StatusBarPartCollection) Count() int {
	return len(me.partsData)
}

// Returns the part at the given index.
func (me *StatusBarPartCollection) Get(index int) StatusBarPart {
	return StatusBarPart{
		owner: me.owner,
		index: int32(index),
	}
}

// Returns the last part.
func (me *StatusBarPartCollection) Last() StatusBarPart {
	return me.Get(me.Count() - 1)
}
