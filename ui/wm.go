package ui

import (
	"wingows/api"
	c "wingows/consts"
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

type WmDropFiles struct {
	Hdrop api.HDROP
}

type WmInitMenuPopup struct {
	Hmenu           api.HMENU
	SourceItemIndex uint16
	IsWindowMenu    bool
}

type wmBaseBtn struct {
	HasCtrl      bool
	HasLeftBtn   bool
	HasMiddleBtn bool
	HasRightBtn  bool
	HasShift     bool
	HasXBtn1     bool
	HasXBtn2     bool
	Pos          *api.POINT
}
type WmLButtonDblClk struct{ wmBaseBtn }
type WmLButtonDown struct{ wmBaseBtn }
type WmLButtonUp struct{ wmBaseBtn }
type WmMButtonDblClk struct{ wmBaseBtn }
type WmMButtonDown struct{ wmBaseBtn }
type WmMButtonUp struct{ wmBaseBtn }
type WmMouseHover struct{ wmBaseBtn }
type WmMouseMove struct{ wmBaseBtn }
type WmRButtonDblClk struct{ wmBaseBtn }
type WmRButtonDown struct{ wmBaseBtn }
type WmRButtonUp struct{ wmBaseBtn }

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
