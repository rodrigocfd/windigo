package ui

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

type windowMainSetup struct {
	ClassName  string
	ClassStyle c.CS
	HCursor    api.HCURSOR
	HBrushBg   api.HBRUSH

	Title   string
	Width   uint32
	Height  uint32
	Style   c.WS
	ExStyle c.WS_EX
	HMenu   api.HMENU

	CmdShow c.SW
}

func makeWindowMainSetup() windowMainSetup {
	return windowMainSetup{
		ClassStyle: c.CS_DBLCLKS,

		Width:   600, // arbitrary dimensions
		Height:  500,
		Style:   c.WS_CAPTION | c.WS_SYSMENU | c.WS_CLIPCHILDREN | c.WS_BORDER,
		ExStyle: c.WS_EX(0),

		CmdShow: c.SW_SHOW,
	}
}

func (me *windowMainSetup) genWndclassex(hInst api.HINSTANCE) *api.WNDCLASSEX {
	wcx := api.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.LpszClassName = api.StrToUtf16PtrBlankIsNil(me.ClassName)
	wcx.Style = me.ClassStyle

	if me.HCursor != 0 {
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = api.HINSTANCE(0).LoadCursor(c.IDC_ARROW)
	}

	if me.HBrushBg != 0 {
		wcx.HbrBackground = me.HBrushBg
	} else {
		wcx.HbrBackground = api.NewBrushFromSysColor(c.COLOR_BTNFACE)
	}

	return &wcx
}
