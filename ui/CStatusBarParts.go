package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _StatusBarPartData struct {
	sizePixels   int
	resizeWeight int
}

func (me *_StatusBarPartData) IsFixedWidth() bool {
	return me.resizeWeight == 0
}

//------------------------------------------------------------------------------

type _StatusBarParts struct {
	pHwnd           *win.HWND
	partsData       []_StatusBarPartData
	rightEdges      []int32 // buffer to speed up ResizeToFitParent() calls
	initialParentCx int     // cache used when adding parts
}

func (me *_StatusBarParts) new(ctrl *_NativeControlBase) {
	me.pHwnd = &ctrl.hWnd
	me.partsData = make([]_StatusBarPartData, 0, 5)
	me.rightEdges = make([]int32, 0, 5)
	me.initialParentCx = 0
}

func (me *_StatusBarParts) cacheInitialParentCx() {
	if me.initialParentCx == 0 { // not cached yet?
		rc := me.pHwnd.GetClientRect()
		me.initialParentCx = int(rc.Right) // initial width of parent's client area
	}
}

func (me *_StatusBarParts) resizeToFitParent(parm wm.Size) {
	if parm.Request() == co.SIZE_REQ_MINIMIZED || *me.pHwnd == 0 {
		return
	}

	cx := int(parm.ClientAreaSize().Cx)    // available width
	me.pHwnd.SendMessage(co.WM_SIZE, 0, 0) // tell status bar to fit parent

	// Find the space to be divided among variable-width parts,
	// and total weight of variable-width parts.
	totalWeight := 0
	cxVariable := cx
	for i := range me.partsData {
		if me.partsData[i].IsFixedWidth() {
			cxVariable -= me.partsData[i].sizePixels
		} else {
			totalWeight += me.partsData[i].resizeWeight
		}
	}

	// Fill right edges array with the right edge of each part.
	cxTotal := cx
	for i := len(me.partsData) - 1; i >= 0; i-- {
		me.rightEdges[i] = int32(cxTotal)
		if me.partsData[i].IsFixedWidth() {
			cxTotal -= me.partsData[i].sizePixels
		} else {
			cxTotal -= (cxVariable / totalWeight) * me.partsData[i].resizeWeight
		}
	}
	me.pHwnd.SendMessage(co.SB_SETPARTS,
		win.WPARAM(len(me.rightEdges)),
		win.LPARAM(unsafe.Pointer(&me.rightEdges[0])))
}

// Adds one or more fixed-width parts.
//
// Widths will be adjusted to the current system DPI.
func (me *_StatusBarParts) AddFixed(widths ...int) {
	me.cacheInitialParentCx()

	for _, width := range widths {
		if width < 0 {
			panic(fmt.Sprintf("Width of part can't be negative: %d.", width))
		}

		size := win.SIZE{Cx: int32(width), Cy: 0}
		_MultiplyDpi(nil, &size)

		me.partsData = append(me.partsData, _StatusBarPartData{
			sizePixels: int(size.Cx),
		})
		me.rightEdges = append(me.rightEdges, 0)
	}

	me.resizeToFitParent(wm.Size{
		Msg: wm.Any{
			WParam: win.WPARAM(co.SIZE_REQ_RESTORED),
			LParam: win.MAKELPARAM(uint16(me.initialParentCx), 0),
		},
	})
}

// Adds one or more resizable parts.
//
// How resizeWeight works:
//
// - Suppose you have 3 parts, respectively with weights of 1, 1 and 2.
//
// - If available client area is 400px, respective part widths will be 100, 100 and 200px.
func (me *_StatusBarParts) AddResizable(resizeWeights ...int) {
	me.cacheInitialParentCx()

	for _, resizeWeight := range resizeWeights {
		if resizeWeight <= 0 {
			panic(fmt.Sprintf("Resize weight must be equal or greater than 1: %d.", resizeWeight))
		}

		me.partsData = append(me.partsData, _StatusBarPartData{
			resizeWeight: resizeWeight,
		})
		me.rightEdges = append(me.rightEdges, 0)
	}

	me.resizeToFitParent(wm.Size{
		Msg: wm.Any{
			WParam: win.WPARAM(co.SIZE_REQ_RESTORED),
			LParam: win.MAKELPARAM(uint16(me.initialParentCx), 0),
		},
	})
}

// Retrieves the texts of all parts at once.
func (me *_StatusBarParts) AllTexts() []string {
	texts := make([]string, 0, me.Count())
	for i := 0; i < me.Count(); i++ {
		texts = append(texts, me.Text(i))
	}
	return texts
}

// Returns the number of parts.
func (me *_StatusBarParts) Count() int {
	return len(me.partsData)
}

// Retrieves the HICON of the part.
//
// The icon is shared, the StatusBar doesn't own it.
func (me *_StatusBarParts) Icon(index int) win.HICON {
	return win.HICON(
		me.pHwnd.SendMessage(co.SB_GETICON, win.WPARAM(index), 0),
	)
}

// Sets the texts of all parts at once.
//
// Panics if the number of texts is greater than the number of parts.
func (me *_StatusBarParts) SetAllTexts(texts ...string) {
	if len(texts) > len(me.partsData) {
		panic(
			fmt.Sprintf("Number of texts (%d) is greater than the number of parts (%d).",
				len(texts), len(me.partsData)))
	}

	for i, text := range texts {
		me.SetText(i, text)
	}
}

// Puts the HICON on the part.
//
// The icon is shared, the StatusBar doesn't own it.
func (me *_StatusBarParts) SetIcon(index int, hIcon win.HICON) {
	me.pHwnd.SendMessage(co.SB_SETICON, win.WPARAM(index), win.LPARAM(hIcon))
}

// Sets the text of the part.
func (me *_StatusBarParts) SetText(index int, text string) {
	ret := me.pHwnd.SendMessage(co.SB_SETTEXT,
		win.MAKEWPARAM(win.MAKEWORD(uint8(index), 0), 0),
		win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(text))))
	if ret == 0 {
		panic(fmt.Sprintf("SB_SETTEXT %d failed \"%s\".", index, text))
	}
}

// Retrieves the text of the part.
func (me *_StatusBarParts) Text(index int) string {
	len := uint16(
		me.pHwnd.SendMessage(co.SB_GETTEXTLENGTH, win.WPARAM(index), 0),
	)
	if len == 0 {
		return ""
	}

	buf := make([]uint16, len+1)
	me.pHwnd.SendMessage(co.SB_GETTEXT,
		win.WPARAM(index), win.LPARAM(unsafe.Pointer(&buf[0])))
	return win.Str.FromUint16Slice(buf)
}
