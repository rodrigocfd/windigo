package ui

import (
	"unsafe"
	a "winffi/api"
	c "winffi/consts"
)

type Param struct {
	Msg    c.WM
	Wparam a.WPARAM
	Lparam a.LPARAM
}

type ParamCommand Param

func (p *ParamCommand) IsFromMenu() bool         { return a.HiWord(uint32(p.Wparam)) == 0 }
func (p *ParamCommand) IsFromAccelerator() bool  { return a.HiWord(uint32(p.Wparam)) == 1 }
func (p *ParamCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p *ParamCommand) MenuId() uint16           { return p.ControlId() }
func (p *ParamCommand) AcceleratorId() uint16    { return p.ControlId() }
func (p *ParamCommand) ControlId() uint16        { return a.LoWord(uint32(p.Wparam)) }
func (p *ParamCommand) ControlNotifCode() uint16 { return a.HiWord(uint32(p.Wparam)) }
func (p *ParamCommand) ControlHwnd() a.HWND      { return a.HWND(p.Lparam) }

type ParamCreate Param

func (p *ParamCreate) Createstruct() *a.CREATESTRUCT {
	return (*a.CREATESTRUCT)(unsafe.Pointer(p.Lparam))
}

type ParamDestroy Param

type ParamNotify Param

func (p *ParamNotify) Nmhdr() *a.NMHDR { return (*a.NMHDR)(unsafe.Pointer(p.Lparam)) }
