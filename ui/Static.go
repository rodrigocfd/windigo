/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"strings"
	"wingows/api"
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

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *Static) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, staStyles c.SS) *Static {

	x, y, width, height = multiplyByDpi(x, y, width, height)

	me.controlNativeBase.create(exStyles, "STATIC", text,
		styles|c.WS(staStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx(). Position will be adjusted to the current system DPI.
// The size will be calculated to fit the text exactly.
func (me *Static) CreateLText(parent Window, x, y int32, text string) *Static {
	staStyles := c.SS_LEFT
	x, y, _, _ = multiplyByDpi(x, y, 0, 0)
	cx, cy := me.calcIdealSize(parent.Hwnd(), text, staStyles)

	me.controlNativeBase.create(c.WS_EX(0), "STATIC", text,
		c.WS_CHILD|c.WS_GROUP|c.WS_VISIBLE|c.WS(staStyles), x, y, cx, cy, parent)
	globalUiFont.SetOnControl(me)
	return me
}

// Sets the text and resizes the static control to fit the text exactly.
func (me *Static) SetText(text string) {
	cx, cy := me.calcIdealSize(me.Hwnd().GetParent(), text,
		c.SS(me.Hwnd().GetStyle()))
	me.Hwnd().SetWindowPos(c.SWP_HWND(0), 0, 0, cx, cy,
		c.SWP_NOZORDER|c.SWP_NOMOVE)
	me.Hwnd().SetWindowText(text)
}

func (me *Static) calcIdealSize(hParent api.HWND, text string,
	staStyles c.SS) (uint32, uint32) {

	parentDc := hParent.GetDC()
	cloneDc := parentDc.CreateCompatibleDC()
	prevFont := cloneDc.SelectObjectFont(globalUiFont.Hfont()) // system font; already adjusted to current DPI

	if (staStyles & c.SS_NOPREFIX) == 0 {
		text = me.removeAmpersands(text)
	}

	bounds := cloneDc.GetTextExtentPoint32(text)
	cloneDc.SelectObjectFont(prevFont)
	cloneDc.DeleteDC()
	hParent.ReleaseDC(parentDc)

	return uint32(bounds.Cx), uint32(bounds.Cy)
}

func (me *Static) removeAmpersands(text string) string {
	buf := strings.Builder{}
	for i := 0; i < len(text)-1; i++ {
		if text[i] == '&' && text[i+1] != '&' {
			continue
		}
		buf.WriteByte(text[i])
	}
	if text[len(text)-1] != '&' {
		buf.WriteByte(text[len(text)-1])
	}
	return buf.String()
}
