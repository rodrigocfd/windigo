package parm

import (
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

type WmActivate Raw

func Event(p *WmActivate) c.WA              { return c.WA(api.LoWord(uint32(p.LParam))) }
func IsMinimized(p *WmActivate) bool        { return api.HiWord(uint32(p.LParam)) != 0 }
func PreviousWindow(p *WmActivate) api.HWND { return api.HWND(p.LParam) }

type WmClose Raw

type WmCommand Raw

func (p *WmCommand) IsFromMenu() bool         { return api.HiWord(uint32(p.WParam)) == 0 }
func (p *WmCommand) IsFromAccelerator() bool  { return api.HiWord(uint32(p.WParam)) == 1 }
func (p *WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p *WmCommand) MenuId() c.ID             { return p.ControlId() }
func (p *WmCommand) AcceleratorId() c.ID      { return p.ControlId() }
func (p *WmCommand) ControlId() c.ID          { return c.ID(api.LoWord(uint32(p.WParam))) }
func (p *WmCommand) ControlNotifCode() uint16 { return api.HiWord(uint32(p.WParam)) }
func (p *WmCommand) ControlHwnd() api.HWND    { return api.HWND(p.LParam) }

type WmCreate Raw

func (p *WmCreate) CreateStruct() *api.CREATESTRUCT {
	return (*api.CREATESTRUCT)(unsafe.Pointer(p.LParam))
}

type WmDestroy Raw

type WmNcDestroy Raw

type WmNotify Raw

func (p *WmNotify) NmHdr() *api.NMHDR { return (*api.NMHDR)(unsafe.Pointer(p.LParam)) }

type WmSize Raw

func Request(p *WmSize) c.SIZE { return c.SIZE(p.WParam) }
func Size(p *WmSize) api.SIZE {
	return api.SIZE{Cx: int32(api.LoWord(uint32(p.LParam))), Cy: int32(api.HiWord(uint32(p.LParam)))}
}
