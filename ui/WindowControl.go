package ui

import (
	c "gowinui/consts"
)

// Custom user control.
type WindowControl struct {
	controlBase
	Setup windowControlSetup
}

func NewWindowControl() *WindowControl {
	return &WindowControl{
		controlBase: makeControlBase(),
		Setup:       makeWindowControlSetup(),
	}
}

func NewWindowControlWithId(ctrlId c.ID) *WindowControl {
	return &WindowControl{
		controlBase: makeControlBaseWithId(ctrlId),
		Setup:       makeWindowControlSetup(),
	}
}

func (me *WindowControl) Create(parent Window, x, y int32,
	width, height uint32) *WindowControl {

	return me
}
