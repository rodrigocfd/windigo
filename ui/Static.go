/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	c "wingows/consts"
)

// Native static control (label).
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Static struct {
	controlNativeBase
}

// Optional; returns a Button with a specific control ID.
func MakeStatic(ctrlId c.ID) Static {
	return Static{
		controlNativeBase: makeNativeControlBase(ctrlId),
	}
}

func (me *Static) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, staStyles c.SS) *Static {

	me.controlNativeBase.create(exStyles, "STATIC", text,
		styles|c.WS(staStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

func (me *Static) CreateSimple(parent Window, x, y int32, text string) *Static {
	cx, cy := me.calcIdealSize(parent, text)
	return me.Create(parent, x, y, cx, cy, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_VISIBLE, c.SS_LEFT|c.SS_WORDELLIPSIS)
}

func (me *Static) calcIdealSize(parent Window, text string) (uint32, uint32) {
	parentDc := parent.Hwnd().GetDC()
	cloneDc := parentDc.CreateCompatibleDC()
	prevFont := cloneDc.SelectObjectFont(globalUiFont.Hfont())
	bounds := cloneDc.GetTextExtentPoint32(text) // counting &, must remove!
	cloneDc.SelectObjectFont(prevFont)
	cloneDc.DeleteDC()
	parent.Hwnd().ReleaseDC(parentDc)
	return uint32(bounds.Cx), uint32(bounds.Cy)
}
