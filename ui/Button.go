package ui

import (
	"gowinui/api"
	c "gowinui/consts"
)

// Button control.
type Button struct {
	controlBase
}

func NewButton() *Button {
	return &Button{
		controlBase: makeControlBase(),
	}
}

func NewButtonWithId(ctrlId c.ID) *Button {
	return &Button{
		controlBase: makeControlBaseWithId(ctrlId),
	}
}

func (me *Button) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *Button {

	me.controlBase.create(exStyles, "Button", text,
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

func (me *Button) Enable(enabled bool) *Button {
	me.controlBase.Hwnd().EnableWindow(enabled)
	return me
}

func (me *Button) IsEnabled() bool {
	return me.controlBase.Hwnd().IsWindowEnabled()
}

func (me *Button) SetFocus() api.HWND {
	return me.controlBase.Hwnd().SetFocus()
}

func (me *Button) SetText(text string) *Button {
	me.controlBase.Hwnd().SetWindowText(text)
	return me
}

func (me *Button) Text() string {
	return me.controlBase.Hwnd().GetWindowText()
}
