package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

func newWmCommand(p wmBase) *WmCommand {
	return &WmCommand{
		IsFromMenu:        api.HiWord(uint32(p.WParam)) == 0,
		IsFromAccelerator: api.HiWord(uint32(p.WParam)) == 1,
		IsFromControl:     api.HiWord(uint32(p.WParam)) != 0 && api.HiWord(uint32(p.WParam)) != 1,
		MenuId:            c.ID(api.LoWord(uint32(p.WParam))),
		AcceleratorId:     c.ID(api.LoWord(uint32(p.WParam))),
		ControlId:         c.ID(api.LoWord(uint32(p.WParam))),
		ControlNotifCode:  api.HiWord(uint32(p.WParam)),
		ControlHwnd:       api.HWND(p.LParam),
	}
}

func (me *windowMsg) WmCommand(cmd c.ID, userFunc func(p *WmCommand)) {
	me.addCmd(cmd, userFunc)
}

//------------------------------------------------------------------------------

func (me *windowMsg) WmActivate(userFunc func(p *WmActivate)) {
	me.addMsg(c.WM_ACTIVATE, func(p wmBase) uintptr {
		userFunc(&WmActivate{
			Event:          c.WA(api.LoWord(uint32(p.LParam))),
			IsMinimized:    api.HiWord(uint32(p.LParam)) != 0,
			PreviousWindow: api.HWND(p.LParam),
		})
		return 0
	})
}

func (me *windowMsg) WmClose(userFunc func()) {
	me.addMsg(c.WM_CLOSE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmCreate(userFunc func(p *WmCreate) int32) {
	me.addMsg(c.WM_CREATE, func(p wmBase) uintptr {
		return uintptr(userFunc(&WmCreate{
			CreateStruct: (*api.CREATESTRUCT)(unsafe.Pointer(p.LParam)),
		}))
	})
}

func (me *windowMsg) WmDestroy(userFunc func()) {
	me.addMsg(c.WM_DESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmDropFiles(userFunc func(p *WmDropFiles)) {
	me.addMsg(c.WM_DROPFILES, func(p wmBase) uintptr {
		userFunc(&WmDropFiles{
			Hdrop: api.HDROP(p.WParam),
		})
		return 0
	})
}

func (me *windowMsg) WmInitMenuPopup(userFunc func(p *WmInitMenuPopup)) {
	me.addMsg(c.WM_INITMENUPOPUP, func(p wmBase) uintptr {
		userFunc(&WmInitMenuPopup{
			Hmenu:           api.HMENU(p.WParam),
			SourceItemIndex: api.LoWord(uint32(p.LParam)),
			IsWindowMenu:    api.HiWord(uint32(p.LParam)) != 0,
		})
		return 0
	})
}

//------------------------------------------------------------------------------

func makeWmBaseBtn(p wmBase) wmBaseBtn {
	return wmBaseBtn{
		HasCtrl:      (c.MK(p.WParam) & c.MK_CONTROL) != 0,
		HasLeftBtn:   (c.MK(p.WParam) & c.MK_LBUTTON) != 0,
		HasMiddleBtn: (c.MK(p.WParam) & c.MK_MBUTTON) != 0,
		HasRightBtn:  (c.MK(p.WParam) & c.MK_RBUTTON) != 0,
		HasShift:     (c.MK(p.WParam) & c.MK_SHIFT) != 0,
		HasXBtn1:     (c.MK(p.WParam) & c.MK_XBUTTON1) != 0,
		HasXBtn2:     (c.MK(p.WParam) & c.MK_XBUTTON2) != 0,
		Pos: &api.POINT{
			X: int32(api.LoWord(uint32(p.LParam))),
			Y: int32(api.HiWord(uint32(p.LParam))),
		},
	}
}

func (me *windowMsg) WmLButtonDblClk(userFunc func(p *WmLButtonDblClk)) {
	me.addMsg(c.WM_LBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(&WmLButtonDblClk{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmLButtonDown(userFunc func(p *WmLButtonDown)) {
	me.addMsg(c.WM_LBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(&WmLButtonDown{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmLButtonUp(userFunc func(p *WmLButtonUp)) {
	me.addMsg(c.WM_LBUTTONUP, func(p wmBase) uintptr {
		userFunc(&WmLButtonUp{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmMButtonDblClk(userFunc func(p *WmMButtonDblClk)) {
	me.addMsg(c.WM_MBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(&WmMButtonDblClk{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmMButtonDown(userFunc func(p *WmMButtonDown)) {
	me.addMsg(c.WM_MBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(&WmMButtonDown{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmMButtonUp(userFunc func(p *WmMButtonUp)) {
	me.addMsg(c.WM_MBUTTONUP, func(p wmBase) uintptr {
		userFunc(&WmMButtonUp{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmMouseHover(userFunc func(p *WmMouseHover)) {
	me.addMsg(c.WM_MOUSEHOVER, func(p wmBase) uintptr {
		userFunc(&WmMouseHover{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmMouseMove(userFunc func(p *WmMouseMove)) {
	me.addMsg(c.WM_MOUSEMOVE, func(p wmBase) uintptr {
		userFunc(&WmMouseMove{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmRButtonDblClk(userFunc func(p *WmRButtonDblClk)) {
	me.addMsg(c.WM_RBUTTONDBLCLK, func(p wmBase) uintptr {
		userFunc(&WmRButtonDblClk{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmRButtonDown(userFunc func(p *WmRButtonDown)) {
	me.addMsg(c.WM_RBUTTONDOWN, func(p wmBase) uintptr {
		userFunc(&WmRButtonDown{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

func (me *windowMsg) WmRButtonUp(userFunc func(p *WmRButtonUp)) {
	me.addMsg(c.WM_RBUTTONUP, func(p wmBase) uintptr {
		userFunc(&WmRButtonUp{
			wmBaseBtn: makeWmBaseBtn(p),
		})
		return 0
	})
}

//------------------------------------------------------------------------------

func (me *windowMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(c.WM_MOUSELEAVE, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmMove(userFunc func(p *WmMove)) {
	me.addMsg(c.WM_MOVE, func(p wmBase) uintptr {
		userFunc(&WmMove{
			Pos: &api.POINT{
				X: int32(api.LoWord(uint32(p.LParam))),
				Y: int32(api.HiWord(uint32(p.LParam))),
			},
		})
		return 0
	})
}

func (me *windowMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(c.WM_NCDESTROY, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmNcPaint(userFunc func(p *WmNcPaint)) {
	me.addMsg(c.WM_NCPAINT, func(p wmBase) uintptr {
		userFunc(&WmNcPaint{
			Hrgn: api.HRGN(p.WParam),
		})
		return 0
	})
}

func (me *windowMsg) WmPaint(userFunc func()) {
	me.addMsg(c.WM_PAINT, func(p wmBase) uintptr {
		userFunc()
		return 0
	})
}

func (me *windowMsg) WmSetFocus(userFunc func(p *WmSetFocus)) {
	me.addMsg(c.WM_SETFOCUS, func(p wmBase) uintptr {
		userFunc(&WmSetFocus{
			UnfocusedWindow: api.HWND(p.WParam),
		})
		return 0
	})
}

func (me *windowMsg) WmSetFont(userFunc func(p *WmSetFont)) {
	me.addMsg(c.WM_SETFONT, func(p wmBase) uintptr {
		userFunc(&WmSetFont{
			Hfont:        api.HFONT(p.WParam),
			ShouldRedraw: p.LParam == 1,
		})
		return 0
	})
}

func (me *windowMsg) WmSize(userFunc func(p *WmSize)) {
	me.addMsg(c.WM_SIZE, func(p wmBase) uintptr {
		userFunc(&WmSize{
			Request: c.SIZE_REQ(p.WParam),
			ClientAreaSize: &api.SIZE{
				Cx: int32(api.LoWord(uint32(p.LParam))),
				Cy: int32(api.HiWord(uint32(p.LParam))),
			},
		})
		return 0
	})
}
