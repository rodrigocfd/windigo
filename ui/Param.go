package ui

import (
	"unsafe"
	a "winffi/api"
	c "winffi/consts"
)

// Param cracks a raw message data.
type Param struct {
	Msg    c.WM
	Wparam a.WPARAM
	Lparam a.LPARAM
}

// ParamCommand cracks WM_COMMAND data.
type ParamCommand Param

// ParamCreate cracks WM_CREATE data.
type ParamCreate Param

// Createstruct cracker.
func (p *ParamCreate) Createstruct() *a.CREATESTRUCT {
	return (*a.CREATESTRUCT)(unsafe.Pointer(p.Lparam))
}

// ParamDestroy cracks WM_DESTROY data.
type ParamDestroy Param

// ParamNotify cracks WM_NOFIFY data.
type ParamNotify Param
