/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"wingows/co"
	"wingows/win"
)

// Used in Resizer.Add().
//
// Action to be done on a child control when the parent is resized.
type RESZ uint8

const (
	RESZ_REPOS   RESZ = iota // When parent resizes, change control left/top position.
	RESZ_RESIZE              // When parent resizes, change control width/height.
	RESZ_NOTHING             // When parent resizes, do nothing.
)

type _ReszCtrl struct {
	hChild Control
	rcOrig win.RECT
	doHorz RESZ
	doVert RESZ
}

// When the parent window resizes, changes size and/or position of several
// children automatically.
type Resizer struct {
	ctrls  []_ReszCtrl
	szOrig win.SIZE
}

// Adds a child control and its resizing behavior.
func (me *Resizer) Add(child Control, doHorz, doVert RESZ) *Resizer {
	hParent := child.Hwnd().GetParent()
	if len(me.ctrls) == 0 { // first control being added
		rc := hParent.GetClientRect()
		me.szOrig.Cx = rc.Right
		me.szOrig.Cy = rc.Bottom // save original size of parent
	}

	me.ctrls = append(me.ctrls, _ReszCtrl{
		hChild: child,
		rcOrig: *child.Hwnd().GetWindowRect(),
		doHorz: doHorz,
		doVert: doVert,
	})
	hParent.ScreenToClientRc(&me.ctrls[len(me.ctrls)-1].rcOrig) // client coordinates relative to parent
	return me
}

// Adds many child controls at once with an unique resizing behavior.
func (me *Resizer) AddMany(children []Control, doHorz, doVert RESZ) *Resizer {
	for _, child := range children {
		me.Add(child, doHorz, doVert)
	}
	return me
}

// Call during WM_SIZE processing to adjust all child controls at once.
func (me *Resizer) Adjust(p WmSize) {
	if len(me.ctrls) == 0 || p.Request() == co.SIZE_MINIMIZED {
		return // no need to resize if window is minimized
	}

	hdwp := win.BeginDeferWindowPos(uint32(len(me.ctrls)))
	defer hdwp.EndDeferWindowPos()

	for i := range me.ctrls {
		c := me.ctrls[i]

		uFlags := co.SWP_NOZORDER
		if c.doHorz == RESZ_REPOS && c.doVert == RESZ_REPOS { // repos both horz and vert
			uFlags |= co.SWP_NOSIZE
		} else if c.doHorz == RESZ_RESIZE && c.doVert == RESZ_RESIZE { // resize both horz and vert
			uFlags |= co.SWP_NOMOVE
		}

		szParent := p.ClientAreaSize()

		x := c.rcOrig.Left // keep original left pos
		if c.doHorz == RESZ_REPOS {
			x = szParent.Cx - me.szOrig.Cx + c.rcOrig.Left
		}

		y := c.rcOrig.Top // keep original top pos
		if c.doVert == RESZ_REPOS {
			y = szParent.Cy - me.szOrig.Cy + c.rcOrig.Top
		}

		cx := uint32(c.rcOrig.Right - c.rcOrig.Left) // keep original width
		if c.doHorz == RESZ_RESIZE {
			cx = uint32(szParent.Cx - me.szOrig.Cx + c.rcOrig.Right - c.rcOrig.Left)
		}

		cy := uint32(c.rcOrig.Bottom - c.rcOrig.Top) // keep original height
		if c.doVert == RESZ_RESIZE {
			cy = uint32(szParent.Cy - me.szOrig.Cy + c.rcOrig.Bottom - c.rcOrig.Top)
		}

		hdwp.DeferWindowPos(c.hChild.Hwnd(), win.HWND(0), x, y, cx, cy, uFlags)
	}
}
