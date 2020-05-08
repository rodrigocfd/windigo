package parm

import (
	"gowinui/api"
	c "gowinui/consts"
	"unsafe"
)

type WmActivate Raw

func Event(p WmActivate) c.WA              { return c.WA(api.LoWord(uint32(p.LParam))) }
func IsMinimized(p WmActivate) bool        { return api.HiWord(uint32(p.LParam)) != 0 }
func PreviousWindow(p WmActivate) api.HWND { return api.HWND(p.LParam) }

type WmClose Raw

type WmCommand Raw

func (p WmCommand) IsFromMenu() bool         { return api.HiWord(uint32(p.WParam)) == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return api.HiWord(uint32(p.WParam)) == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() c.ID             { return p.ControlId() }
func (p WmCommand) AcceleratorId() c.ID      { return p.ControlId() }
func (p WmCommand) ControlId() c.ID          { return c.ID(api.LoWord(uint32(p.WParam))) }
func (p WmCommand) ControlNotifCode() uint16 { return api.HiWord(uint32(p.WParam)) }
func (p WmCommand) ControlHwnd() api.HWND    { return api.HWND(p.LParam) }

type WmCreate Raw

func (p WmCreate) CreateStruct() *api.CREATESTRUCT {
	return (*api.CREATESTRUCT)(unsafe.Pointer(p.LParam))
}

type WmDestroy Raw

type WmInitMenuPopup Raw

func (p WmInitMenuPopup) Hmenu() api.HMENU        { return api.HMENU(p.WParam) }
func (p WmInitMenuPopup) SourceItemIndex() uint16 { return api.LoWord(uint32(p.LParam)) }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return api.HiWord(uint32(p.LParam)) != 0 }

type WmLButtonDblClk Raw

func (p WmLButtonDblClk) HasCtrl() bool      { return (c.MK(p.WParam) & c.MK_CONTROL) != 0 }
func (p WmLButtonDblClk) HasLeftBtn() bool   { return (c.MK(p.WParam) & c.MK_LBUTTON) != 0 }
func (p WmLButtonDblClk) HasMiddleBtn() bool { return (c.MK(p.WParam) & c.MK_MBUTTON) != 0 }
func (p WmLButtonDblClk) HasRightBtn() bool  { return (c.MK(p.WParam) & c.MK_RBUTTON) != 0 }
func (p WmLButtonDblClk) HasShift() bool     { return (c.MK(p.WParam) & c.MK_SHIFT) != 0 }
func (p WmLButtonDblClk) HasXBtn1() bool     { return (c.MK(p.WParam) & c.MK_XBUTTON1) != 0 }
func (p WmLButtonDblClk) HasXBtn2() bool     { return (c.MK(p.WParam) & c.MK_XBUTTON2) != 0 }
func (p WmLButtonDblClk) Pos() *api.POINT {
	return &api.POINT{X: int32(api.LoWord(uint32(p.LParam))), Y: int32(api.HiWord(uint32(p.LParam)))}
}

type WmLButtonDown struct{ WmLButtonDblClk } // inherit
type WmLButtonUp struct{ WmLButtonDblClk }
type WmMButtonDblClk struct{ WmLButtonDblClk }
type WmMButtonDown struct{ WmLButtonDblClk }
type WmMButtonUp struct{ WmLButtonDblClk }
type WmMouseHover struct{ WmLButtonDblClk }
type WmMouseMove struct{ WmLButtonDblClk }
type WmRButtonDblClk struct{ WmLButtonDblClk }
type WmRButtonDown struct{ WmLButtonDblClk }
type WmRButtonUp struct{ WmLButtonDblClk }

type WmMouseLeave Raw

type WmMove Raw

func (p WmMove) Pos() *api.POINT {
	return &api.POINT{X: int32(api.LoWord(uint32(p.LParam))), Y: int32(api.HiWord(uint32(p.LParam)))}
}

type WmNcDestroy Raw

type WmNcPaint Raw

func (p WmNcPaint) Hrgn() api.HRGN { return api.HRGN(p.LParam) }

type WmNotify Raw

func (p WmNotify) NmHdr() *api.NMHDR { return (*api.NMHDR)(unsafe.Pointer(p.LParam)) }

type WmPaint Raw

type WmSetFocus Raw

func (p WmSetFocus) unfocusedWindow() api.HWND { return api.HWND(p.WParam) }

type WmSetFont Raw

func (p WmSetFont) Hfont() api.HFONT   { return api.HFONT(p.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.LParam == 1 }

type WmSize Raw

func (p WmSize) Request() c.SIZE_REQ { return c.SIZE_REQ(p.WParam) }
func (p WmSize) SizeClientArea() *api.SIZE {
	return &api.SIZE{Cx: int32(api.LoWord(uint32(p.LParam))), Cy: int32(api.HiWord(uint32(p.LParam)))}
}
