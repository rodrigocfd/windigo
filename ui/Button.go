package ui

import (
	"winffi/api"
	c "winffi/consts"
)

// Button control.
type Button struct {
	hwnd api.HWND
	id   c.ID
}

func MakeButton() Button {
	return MakeButtonWithId(NextAutoCtrlId())
}

func MakeButtonWithId(ctrlId c.ID) Button {
	return Button{
		hwnd: api.HWND(0),
		id:   ctrlId,
	}
}

func (btn *Button) Create(parent Window, x, y int32, width, height uint32,
	text string, exStyles c.WS_EX, styles c.WS, btnStyles c.BS) *Button {

	if btn.hwnd != 0 {
		panic("Trying to create Button twice.")
	}
	btn.hwnd = api.CreateWindowEx(exStyles, "Button", text,
		styles|c.WS(btnStyles), x, y, width, height,
		parent.Hwnd(), api.HMENU(btn.id), parent.Hwnd().GetInstance(), nil)
	globalUiFont.SetOnControl(btn)
	return btn
}

func (btn *Button) CreateSimple(parent Window, x, y int32,
	width uint32, text string) *Button {

	return btn.Create(parent, x, y, width, 23, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.BS(0))
}

func (btn *Button) CreateSimpleDef(parent Window, x, y int32,
	width uint32, text string) *Button {

	return btn.Create(parent, x, y, width, 23, text,
		c.WS_EX(0), c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.BS_DEFPUSHBUTTON)
}

func (btn *Button) CtrlId() c.ID {
	return btn.id
}

func (btn *Button) Enable(enabled bool) *Button {
	btn.hwnd.EnableWindow(enabled)
	return btn
}

func (btn *Button) GetText() string {
	return btn.hwnd.GetWindowText()
}

func (btn *Button) Hwnd() api.HWND {
	return btn.hwnd
}

func (btn *Button) IsEnabled() bool {
	return btn.hwnd.IsWindowEnabled()
}

func (btn *Button) SetFocus() api.HWND {
	return btn.hwnd.SetFocus()
}

func (btn *Button) SetText(text string) *Button {
	btn.hwnd.SetWindowText(text)
	return btn
}
