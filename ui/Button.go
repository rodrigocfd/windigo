/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	c "wingows/consts"
)

// Native button control.
// Can be default-initialized.
// Call one of the create methods during parent's WM_CREATE.
type Button struct {
	nativeControlBase
}

// Optional; returns a Button with a specific control ID.
func MakeButton(ctrlId c.ID) Button {
	return Button{
		nativeControlBase: makeNativeControlBase(ctrlId),
	}
}

func (me *Button) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *Button {

	me.nativeControlBase.create(exStyles, "Button", text,
		styles|c.WS(btnStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

func (me *Button) CreateSimple(parent Window, x, y int32,
	width uint32, text string) *Button {

	return me.Create(parent, x, y, width, 23, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.BS(0))
}

func (me *Button) CreateSimpleDef(parent Window, x, y int32,
	width uint32, text string) *Button {

	return me.Create(parent, x, y, width, 23, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.BS_DEFPUSHBUTTON)
}
