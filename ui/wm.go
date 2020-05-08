package ui

import (
	"gowinui/api"
	c "gowinui/consts"
)

type wmBase struct { // raw window message parameters
	Msg    c.WM
	WParam api.WPARAM
	LParam api.LPARAM
}

type WmCommand struct {
	IsFromMenu        bool
	IsFromAccelerator bool
	IsFromControl     bool
	MenuId            c.ID
	AcceleratorId     c.ID
	ControlId         c.ID
	ControlNotifCode  uint16
	ControlHwnd       api.HWND
}

// type WmNotify struct {
// 	NmHdr *api.NMHDR
// }

//------------------------------------------------------------------------------

type WmActivate struct {
	Event          c.WA
	IsMinimized    bool
	PreviousWindow api.HWND
}

type WmCreate struct {
	CreateStruct *api.CREATESTRUCT
}

type WmInitMenuPopup struct {
	Hmenu           api.HMENU
	SourceItemIndex uint16
	IsWindowMenu    bool
}

type WmBaseBtn struct {
	HasCtrl      bool
	HasLeftBtn   bool
	HasMiddleBtn bool
	HasRightBtn  bool
	HasShift     bool
	HasXBtn1     bool
	HasXBtn2     bool
	Pos          *api.POINT
}
type WmLButtonDblClk struct{ WmBaseBtn }
type WmLButtonDown struct{ WmBaseBtn }
type WmLButtonUp struct{ WmBaseBtn }
type WmMButtonDblClk struct{ WmBaseBtn }
type WmMButtonDown struct{ WmBaseBtn }
type WmMButtonUp struct{ WmBaseBtn }
type WmMouseHover struct{ WmBaseBtn }
type WmMouseMove struct{ WmBaseBtn }
type WmRButtonDblClk struct{ WmBaseBtn }
type WmRButtonDown struct{ WmBaseBtn }
type WmRButtonUp struct{ WmBaseBtn }

type WmMove struct {
	Pos *api.POINT
}

type WmNcPaint struct {
	Hrgn api.HRGN
}

type WmSetFocus struct {
	UnfocusedWindow api.HWND
}

type WmSetFont struct {
	Hfont        api.HFONT
	ShouldRedraw bool
}

type WmSize struct {
	Request        c.SIZE_REQ
	ClientAreaSize *api.SIZE
}
