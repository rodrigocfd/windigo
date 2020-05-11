package ui

import (
	"wingows/api"
	c "wingows/consts"
)

// Native edit control.
type Edit struct {
	nativeControlBase
}

func NewEdit() *Edit {
	return &Edit{
		nativeControlBase: makeNativeControlBase(),
	}
}

func NewEditWithId(ctrlId c.ID) *Edit {
	return &Edit{
		nativeControlBase: makeNativeControlBaseWithId(ctrlId),
	}
}

func (me *Edit) Create(parent Window, x, y int32, width, height uint32,
	initialText string, exStyles c.WS_EX, styles c.WS, editStyles c.ES) *Edit {

	me.nativeControlBase.create(exStyles, "Edit", initialText,
		styles|c.WS(editStyles), x, y, width, height, parent)
	globalUiFont.SetOnControl(me)
	return me
}

func (me *Edit) CreateMultiLine(parent Window, x, y int32,
	width, height uint32, initialText string) *Edit {

	return me.Create(parent, x, y, width, height, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_MULTILINE|c.ES_WANTRETURN)
}

func (me *Edit) CreatePassword(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return me.Create(parent, x, y, width, 21, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_AUTOHSCROLL|c.ES_PASSWORD)
}

func (me *Edit) CreateSimple(parent Window, x, y int32, width uint32,
	initialText string) *Edit {

	return me.Create(parent, x, y, width, 21, initialText,
		c.WS_EX_CLIENTEDGE,
		c.WS_CHILD|c.WS_GROUP|c.WS_TABSTOP|c.WS_VISIBLE,
		c.ES_AUTOHSCROLL)
}

func (me *Edit) Enable(enabled bool) *Edit {
	me.nativeControlBase.Hwnd().EnableWindow(enabled)
	return me
}

func (me *Edit) IsEnabled() bool {
	return me.nativeControlBase.Hwnd().IsWindowEnabled()
}

func (me *Edit) SetFocus() api.HWND {
	return me.nativeControlBase.Hwnd().SetFocus()
}

func (me *Edit) SetText(text string) *Edit {
	me.nativeControlBase.Hwnd().SetWindowText(text)
	return me
}

func (me *Edit) Text() string {
	return me.nativeControlBase.Hwnd().GetWindowText()
}
