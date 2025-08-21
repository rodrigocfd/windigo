//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _StatusBarPartData struct {
	sizePixels   uint
	resizeWeight uint
}

func (me *_StatusBarPartData) IsFixedWidth() bool {
	return me.resizeWeight == 0
}

// The parts collection.
//
// You cannot create this object directly, it will be created automatically
// by the owning [StatusBar].
type CollectionStatusBarParts struct {
	owner           *StatusBar
	partsData       []_StatusBarPartData
	rightEdges      []int32 // buffer to speed up ResizeToFitParent() calls
	initialParentCx uint    // cache used when adding parts
}

func (me *CollectionStatusBarParts) cacheInitialParentCx() {
	if me.initialParentCx == 0 { // not cached yet?
		rc, _ := me.owner.hWnd.GetClientRect()
		me.initialParentCx = uint(rc.Right) // initial width of parent's client area
	}
}

func (me *CollectionStatusBarParts) resizeToFitParent(parm WmSize) {
	if parm.Request() == co.SIZE_REQ_MINIMIZED || me.owner.hWnd == 0 {
		return
	}
	me.owner.hWnd.SendMessage(co.WM_SIZE, 0, 0) // tell status bar to fit parent

	if len(me.partsData) == 0 {
		return // no parts added, nothing else to do
	}

	cx := uint(parm.ClientAreaSize().Cx) // available width

	totalWeight := uint(0) // total weight of all variable-width parts
	cxVariable := int(cx)  // total width to be divided among variable-width parts
	for i := range me.partsData {
		if me.partsData[i].IsFixedWidth() {
			cxVariable -= int(me.partsData[i].sizePixels)
		} else {
			totalWeight += me.partsData[i].resizeWeight
		}
	}

	cxTotal := int(cx)
	for i := len(me.partsData) - 1; i >= 0; i-- { // fill right edges array with the right edge of each part
		me.rightEdges[i] = int32(cxTotal)
		if me.partsData[i].IsFixedWidth() {
			cxTotal -= int(me.partsData[i].sizePixels)
		} else {
			cxTotal -= (cxVariable / int(totalWeight)) * int(me.partsData[i].resizeWeight)
		}
	}
	me.owner.hWnd.SendMessage(co.SB_SETPARTS,
		win.WPARAM(len(me.rightEdges)),
		win.LPARAM(unsafe.Pointer(&me.rightEdges[0])))
}

// Adds a fixed part.
//
// Panics if width is zero.
//
// Example:
//
//	var sbar ui.StatusBar // initialized somewhere
//
//	sbar.Parts.AddFixed("Text", ui.DpiX(200))
func (me *CollectionStatusBarParts) AddFixed(text string, width int) {
	me.cacheInitialParentCx()

	if width == 0 {
		panic(fmt.Sprintf("Width must be equal or greater than 1: %d.", width))
	}

	me.partsData = append(me.partsData, _StatusBarPartData{
		sizePixels: uint(width),
	})
	me.rightEdges = append(me.rightEdges, 0)

	me.resizeToFitParent(WmSize{
		Raw: Wm{
			WParam: win.WPARAM(co.SIZE_REQ_RESTORED),
			LParam: win.MAKELPARAM(uint16(me.initialParentCx), 0),
		},
	})

	lastIndex := len(me.partsData) - 1
	me.Get(lastIndex).SetText(text)
}

// Adds a resizable part.
//
// How resizeWeight works:
//   - Suppose you have 3 parts, respectively with weights of 1, 1 and 2.
//   - If available client area is 400px, respective part widths will be 100, 100 and 200px.
//
// Panics if resizeWeight is zero.
//
// Example:
//
//	var sbar ui.StatusBar // initialized somewhere
//
//	sbar.Parts.AddResizable("Text", 1)
func (me *CollectionStatusBarParts) AddResizable(text string, resizeWeight uint) {
	me.cacheInitialParentCx()

	if resizeWeight == 0 {
		panic(fmt.Sprintf("Resize weight must be equal or greater than 1: %d.", resizeWeight))
	}

	me.partsData = append(me.partsData, _StatusBarPartData{
		resizeWeight: resizeWeight,
	})
	me.rightEdges = append(me.rightEdges, 0)

	me.resizeToFitParent(WmSize{
		Raw: Wm{
			WParam: win.WPARAM(co.SIZE_REQ_RESTORED),
			LParam: win.MAKELPARAM(uint16(me.initialParentCx), 0),
		},
	})

	lastIndex := len(me.partsData) - 1
	me.Get(lastIndex).SetText(text)
}

// Returns all parts.
func (me *CollectionStatusBarParts) All() []StatusBarPart {
	nParts := me.Count()
	parts := make([]StatusBarPart, 0, nParts)
	for i := uint(0); i < nParts; i++ {
		parts = append(parts, me.Get(int(i)))
	}
	return parts
}

// Returns the number of parts.
func (me *CollectionStatusBarParts) Count() uint {
	return uint(len(me.partsData))
}

// Returns the part at the given index.
func (me *CollectionStatusBarParts) Get(index int) StatusBarPart {
	return StatusBarPart{
		owner: me.owner,
		index: int32(index),
	}
}

// Returns the last part.
func (me *CollectionStatusBarParts) Last() StatusBarPart {
	return me.Get(int(me.Count()) - 1)
}
