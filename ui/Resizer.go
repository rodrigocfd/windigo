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

// Action to be done when resizing occurs.
type ResizerDo uint8

const (
	ResizerDo_Repos  ResizerDo = iota // Move left/top coordinates of the control.
	ResizerDo_Resize                  // Increase or decrease width/height.
	ResizerDo_Nothing
)

type ctrl struct {
	hChild Control
	rcOrig win.RECT
	doHorz ResizerDo
	doVert ResizerDo
}

type Resizer struct {
	ctrls  []ctrl
	szOrig win.SIZE
}

func (me *Resizer) Add(hCtrl Control, doHorz, doVert ResizerDo) *Resizer {
	hParent := hCtrl.Hwnd().GetParent()
	if len(me.ctrls) == 0 { // first control being added
		rc := hParent.GetClientRect()
		me.szOrig.Cx = rc.Right
		me.szOrig.Cy = rc.Bottom // save original size of parent
	}

	me.ctrls = append(me.ctrls, ctrl{
		hChild: hCtrl,
		rcOrig: *hCtrl.Hwnd().GetWindowRect(),
		doHorz: doHorz,
		doVert: doVert,
	})
	hParent.ScreenToClientRc(&me.ctrls[len(me.ctrls)-1].rcOrig) // client coordinates relative to parent
	return me
}

func (me *Resizer) AddMany(hCtrls []Control,
	doHorz, doVert ResizerDo) *Resizer {

	for i := range hCtrls {
		me.Add(hCtrls[i], doHorz, doVert)
	}
	return me
}

func (me *Resizer) Adjust(p WmSize) {
	if len(me.ctrls) == 0 || p.Request() == co.SIZE_MINIMIZED {
		return // no need to resize if window is minimized
	}

	hdwp := win.BeginDeferWindowPos(uint32(len(me.ctrls)))
	defer hdwp.EndDeferWindowPos()

	for i := range me.ctrls {
		c := me.ctrls[i]

		uFlags := co.SWP_NOZORDER
		if c.doHorz == ResizerDo_Repos && c.doVert == ResizerDo_Repos { // repos both horz and vert
			uFlags |= co.SWP_NOSIZE
		} else if c.doHorz == ResizerDo_Resize && c.doVert == ResizerDo_Resize { // resize both horz and vert
			uFlags |= co.SWP_NOMOVE
		}

		szParent := p.ClientAreaSize()
		x := func() int32 {
			if c.doHorz == ResizerDo_Repos {
				return szParent.Cx - me.szOrig.Cx + c.rcOrig.Left
			} else {
				return c.rcOrig.Left // keep original pos
			}
		}()
		y := func() int32 {
			if c.doVert == ResizerDo_Repos {
				return szParent.Cy - me.szOrig.Cy + c.rcOrig.Top
			} else {
				return c.rcOrig.Top // keep original pos
			}
		}()
		cx := func() uint32 {
			if c.doHorz == ResizerDo_Resize {
				return uint32(szParent.Cx - me.szOrig.Cx + c.rcOrig.Right - c.rcOrig.Left)
			} else {
				return uint32(c.rcOrig.Right - c.rcOrig.Left) // keep original width
			}
		}()
		cy := func() uint32 {
			if c.doVert == ResizerDo_Resize {
				return uint32(szParent.Cy - me.szOrig.Cy + c.rcOrig.Bottom - c.rcOrig.Top)
			} else {
				return uint32(c.rcOrig.Bottom - c.rcOrig.Top) // keep original height
			}
		}()

		hdwp.DeferWindowPos(c.hChild.Hwnd(), win.HWND(0), x, y, cx, cy, uFlags)
	}
}
