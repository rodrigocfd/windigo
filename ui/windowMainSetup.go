package ui

import (
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

type windowMainSetup struct {
	ClassName  string
	ClassStyle c.CS
	Hcursor    api.HCURSOR
	HbrushBg   api.HBRUSH

	Title   string
	Width   uint32
	Height  uint32
	Style   c.WS
	ExStyle c.WS_EX
	hMenu   api.HMENU

	CmdShow c.SW
}

// Constructor: must use.
func newWindowMainSetup() windowMainSetup {
	return windowMainSetup{
		ClassStyle: c.CS_DBLCLKS,

		Width:   600,
		Height:  500,
		Style:   c.WS_CAPTION | c.WS_SYSMENU | c.WS_CLIPCHILDREN | c.WS_BORDER,
		ExStyle: c.WS_EX(0),

		CmdShow: c.SW_SHOW,
	}
}

func (s *windowMainSetup) checkInit() {
	if s.Width == 0 || s.Height == 0 {
		panic("Internal structures not initialized... did you use the class constructor?")
	}
}

func (s *windowMainSetup) genWndclassex(hInst api.HINSTANCE) *api.WNDCLASSEX {
	wcx := api.WNDCLASSEX{}

	wcx.Size = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.LpszClassName = api.ToUtf16PtrBlankIsNil(s.ClassName)
	wcx.Style = s.ClassStyle
	wcx.HCursor = s.Hcursor
	wcx.HbrBackground = s.HbrushBg

	return &wcx
}
