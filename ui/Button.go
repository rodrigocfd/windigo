package ui

import (
	"gowinui/api"
	c "gowinui/consts"
)

// Button control.
type Button struct {
	hwnd api.HWND
	id   c.ID
}

func NewButton() *Button {
	return NewButtonWithId(NextAutoCtrlId())
}

func NewButtonWithId(ctrlId c.ID) *Button {
	return &Button{
		hwnd: api.HWND(0),
		id:   ctrlId,
	}
}

func (me *Button) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *Button {

	if me.hwnd != 0 {
		panic("Trying to create Button twice.")
	}
	me.hwnd = api.CreateWindowEx(exStyles, "Button", text,
		styles|c.WS(btnStyles), x, y, width, height,
		parent.Hwnd(), api.HMENU(me.id), parent.Hwnd().GetInstance(), nil)
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

func (me *Button) CtrlId() c.ID {
	return me.id
}

func (me *Button) Enable(enabled bool) *Button {
	me.hwnd.EnableWindow(enabled)
	return me
}

func (me *Button) Hwnd() api.HWND {
	return me.hwnd
}

func (me *Button) IsEnabled() bool {
	return me.hwnd.IsWindowEnabled()
}

func (me *Button) SetFocus() api.HWND {
	return me.hwnd.SetFocus()
}

func (me *Button) SetText(text string) *Button {
	me.hwnd.SetWindowText(text)
	return me
}

func (me *Button) Text() string {
	return me.hwnd.GetWindowText()
}
