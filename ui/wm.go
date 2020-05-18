/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

// Raw window message parameters.
type wmBase struct {
	WParam api.WPARAM
	LParam api.LPARAM
}

type WmCommand struct{ base wmBase }

func (p WmCommand) IsFromMenu() bool         { return api.HiWord(uint32(p.base.WParam)) == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return api.HiWord(uint32(p.base.WParam)) == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() c.ID             { return p.ControlId() }
func (p WmCommand) AcceleratorId() c.ID      { return p.ControlId() }
func (p WmCommand) ControlId() c.ID          { return c.ID(api.LoWord(uint32(p.base.WParam))) }
func (p WmCommand) ControlNotifCode() uint16 { return api.HiWord(uint32(p.base.WParam)) }
func (p WmCommand) ControlHwnd() api.HWND    { return api.HWND(p.base.LParam) }

type WmNotify struct{ base wmBase }

func (p WmNotify) NmHdr() *api.NMHDR { return (*api.NMHDR)(unsafe.Pointer(p.base.LParam)) }

//------------------------------------------------------------------------------

type WmActivate struct{ base wmBase }

func (p WmActivate) Event() c.WA              { return c.WA(api.LoWord(uint32(p.base.LParam))) }
func (p WmActivate) IsMinimized() bool        { return api.HiWord(uint32(p.base.LParam)) != 0 }
func (p WmActivate) PreviousWindow() api.HWND { return api.HWND(p.base.LParam) }

type WmCreate struct{ base wmBase }

func (p WmCreate) CreateStruct() *api.CREATESTRUCT {
	return (*api.CREATESTRUCT)(unsafe.Pointer(p.base.LParam))
}

type WmDropFiles struct{ base wmBase }

func (p WmDropFiles) Hdrop() api.HDROP { return api.HDROP(p.base.WParam) }

type WmInitMenuPopup struct{ base wmBase }

func (p WmInitMenuPopup) Hmenu() api.HMENU        { return api.HMENU(p.base.WParam) }
func (p WmInitMenuPopup) SourceItemIndex() uint16 { return api.LoWord(uint32(p.base.LParam)) }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return api.HiWord(uint32(p.base.LParam)) != 0 }

type wmBaseBtn struct{ base wmBase }

func (p wmBaseBtn) HasCtrl() bool      { return (c.MK(p.base.WParam) & c.MK_CONTROL) != 0 }
func (p wmBaseBtn) HasLeftBtn() bool   { return (c.MK(p.base.WParam) & c.MK_LBUTTON) != 0 }
func (p wmBaseBtn) HasMiddleBtn() bool { return (c.MK(p.base.WParam) & c.MK_MBUTTON) != 0 }
func (p wmBaseBtn) HasRightBtn() bool  { return (c.MK(p.base.WParam) & c.MK_RBUTTON) != 0 }
func (p wmBaseBtn) HasShift() bool     { return (c.MK(p.base.WParam) & c.MK_SHIFT) != 0 }
func (p wmBaseBtn) HasXBtn1() bool     { return (c.MK(p.base.WParam) & c.MK_XBUTTON1) != 0 }
func (p wmBaseBtn) HasXBtn2() bool     { return (c.MK(p.base.WParam) & c.MK_XBUTTON2) != 0 }
func (p wmBaseBtn) Pos() api.POINT     { return makePointLp(p.base.LParam) }

type WmLButtonDblClk struct{ wmBaseBtn } // inherit
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

type WmMove struct{ base wmBase }

func (p WmMove) Pos() api.POINT { return makePointLp(p.base.LParam) }

type WmNcPaint struct{ base wmBase }

func (p WmNcPaint) Hrgn() api.HRGN { return api.HRGN(p.base.WParam) }

type WmSetFocus struct{ base wmBase }

func (p WmSetFocus) UnfocusedWindow() api.HWND { return api.HWND(p.base.WParam) }

type WmSetFont struct{ base wmBase }

func (p WmSetFont) Hfont() api.HFONT   { return api.HFONT(p.base.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.base.LParam == 1 }

type WmSize struct{ base wmBase }

func (p WmSize) Request() c.SIZE_REQ      { return c.SIZE_REQ(p.base.WParam) }
func (p WmSize) ClientAreaSize() api.SIZE { return makeSizeLp(p.base.LParam) }

//------------------------------------------------------------------------------

func makePointLp(p api.LPARAM) api.POINT {
	return api.POINT{
		X: int32(api.LoWord(uint32(p))),
		Y: int32(api.HiWord(uint32(p))),
	}
}

func makeSizeLp(p api.LPARAM) api.SIZE {
	return api.SIZE{
		Cx: int32(api.LoWord(uint32(p))),
		Cy: int32(api.HiWord(uint32(p))),
	}
}
