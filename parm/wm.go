package parm

import (
	"unsafe"
	"winffi/api"
	c "winffi/consts"
)

type Raw struct {
	Msg    c.WM
	Wparam api.WPARAM
	Lparam api.LPARAM
}

type WmCommand Raw

func (p *WmCommand) IsFromMenu() bool         { return api.HiWord(uint32(p.Wparam)) == 0 }
func (p *WmCommand) IsFromAccelerator() bool  { return api.HiWord(uint32(p.Wparam)) == 1 }
func (p *WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p *WmCommand) MenuId() uint16           { return p.ControlId() }
func (p *WmCommand) AcceleratorId() uint16    { return p.ControlId() }
func (p *WmCommand) ControlId() uint16        { return api.LoWord(uint32(p.Wparam)) }
func (p *WmCommand) ControlNotifCode() uint16 { return api.HiWord(uint32(p.Wparam)) }
func (p *WmCommand) ControlHwnd() api.HWND    { return api.HWND(p.Lparam) }

type WmCreate Raw

func (p *WmCreate) Createstruct() *api.CREATESTRUCT {
	return (*api.CREATESTRUCT)(unsafe.Pointer(p.Lparam))
}

type WmDestroy Raw

type WmNotify Raw

func (p *WmNotify) Nmhdr() *api.NMHDR { return (*api.NMHDR)(unsafe.Pointer(p.Lparam)) }
