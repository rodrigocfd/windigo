/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
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

//------------------------------------------------------------------------------

// When the parent window resizes, changes size and/or position of several
// children automatically.
type Resizer struct {
	parent Parent
	ctrls  []_ReszCtrl
	szOrig win.SIZE // original parent client area size
}

type _ReszCtrl struct {
	hChild Control
	rcOrig win.RECT
	doHorz RESZ
	doVert RESZ
}

// Constructor.
func NewResizer(parent Parent) *Resizer {
	return &Resizer{
		parent: parent,
	}
}

// Adds child controls and their behavior when the parent is resized.
func (me *Resizer) Add(
	horzBehavior, vertBehavior RESZ, ctrls ...Control) *Resizer {

	if len(me.ctrls) == 0 { // first one being added
		rcParent := me.parent.Hwnd().GetClientRect()
		me.szOrig = win.SIZE{Cx: rcParent.Right, Cy: rcParent.Bottom} // cache
	}

	for _, ctrl := range ctrls {
		me.ctrls = append(me.ctrls, _ReszCtrl{
			hChild: ctrl,
			rcOrig: *ctrl.Hwnd().GetWindowRect(),
			doHorz: horzBehavior,
			doVert: vertBehavior,
		})
		me.parent.Hwnd().ScreenToClientRc(&me.ctrls[len(me.ctrls)-1].rcOrig) // client coordinates relative to parent
	}
	return me
}

// Call during WM_SIZE processing to adjust all child controls at once.
func (me *Resizer) AdjustToParent(p WmSize) {
	if len(me.ctrls) == 0 || p.Request() == co.SIZE_MINIMIZED {
		return // no need to resize if window is minimized
	}

	hdwp := win.BeginDeferWindowPos(int32(len(me.ctrls)))
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

		cx := c.rcOrig.Right - c.rcOrig.Left // keep original width
		if c.doHorz == RESZ_RESIZE {
			cx = szParent.Cx - me.szOrig.Cx + c.rcOrig.Right - c.rcOrig.Left
		}

		cy := c.rcOrig.Bottom - c.rcOrig.Top // keep original height
		if c.doVert == RESZ_RESIZE {
			cy = szParent.Cy - me.szOrig.Cy + c.rcOrig.Bottom - c.rcOrig.Top
		}

		hdwp.DeferWindowPos(c.hChild.Hwnd(), win.HWND(0), x, y, cx, cy, uFlags)
	}
}
