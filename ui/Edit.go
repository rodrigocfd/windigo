package ui

import (
	"winffi/api"
	c "winffi/consts"
)

// Edit control.
type Edit struct {
	hwnd api.HWND
	id   c.ID
}

func MakeEdit() Edit {
	return MakeEditWithId(NextAutoCtrlId())
}

func MakeEditWithId(ctrlId c.ID) Edit {
	return Edit{
		hwnd: api.HWND(0),
		id:   ctrlId,
	}
}

func (edit *Edit) Create(parent Window, x, y int32, width, height uint32,
	initialText string, exStyles c.WS_EX, styles c.WS, editStyles c.ES) *Edit {

	if edit.hwnd != 0 {
		panic("Trying to create Edit twice.")
	}
	edit.hwnd = api.CreateWindowEx(exStyles, "Edit", initialText,
		styles|c.WS(editStyles), x, y, width, height,
		parent.Hwnd(), api.HMENU(edit.id), parent.Hwnd().GetInstance(), nil)
	globalUiFont.SetOnControl(edit)
	return edit
}

func (edit *Edit) CreateMultiLine(parent Window, x, y int32,
	width, height uint32, initialText string) *Edit {

	return edit.Create(parent, x, y, width, height, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_MULTILINE|c.ES_WANTRETURN)
}

func (edit *Edit) CreatePassword(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return edit.Create(parent, x, y, width, 21, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_AUTOHSCROLL|c.ES_PASSWORD)
}

func (edit *Edit) CreateSimple(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return edit.Create(parent, x, y, width, 21, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_AUTOHSCROLL)
}

func (edit *Edit) CtrlId() c.ID {
	return edit.id
}

func (edit *Edit) Enable(enabled bool) *Edit {
	edit.hwnd.EnableWindow(enabled)
	return edit
}

func (edit *Edit) GetText() string {
	return edit.hwnd.GetWindowText()
}

func (edit *Edit) Hwnd() api.HWND {
	return edit.hwnd
}

func (edit *Edit) IsEnabled() bool {
	return edit.hwnd.IsWindowEnabled()
}

func (edit *Edit) SetFocus() api.HWND {
	return edit.hwnd.SetFocus()
}

func (edit *Edit) SetText(text string) *Edit {
	edit.hwnd.SetWindowText(text)
	return edit
}
