//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// Manages a group of native [RadioButton] controls.
//
// [radio buttons]: https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioGroup struct {
	radios []*RadioButton
	events EventsRadioGroup
}

// Creates the [RadioButton] controls with [win.CreateWindowEx].
//
// Panics if the number of radio buttons is zero.
//
// # Example
//
//	var wndOwner ui.Parent // initialized somewhere
//
//	radios := ui.NewRadioGroup(
//		wndOwner,
//		ui.OptsRadioButton().
//			Text("&First").
//			Position(260, 125),
//		ui.OptsRadioButton().
//			Text("&Second").
//			Position(260, 145).
//			Selected(true),
//	)
func NewRadioGroup(parent Parent, allOpts ...*VarOptsRadioButton) *RadioGroup {
	if len(allOpts) == 0 {
		panic("Cannot create a RadioGroup without radio buttons.")
	}

	me := &RadioGroup{
		radios: make([]*RadioButton, 0, len(allOpts)),
	}

	ctrlIds := make([]uint16, 0, len(allOpts))
	for idx, opts := range allOpts {
		setUniqueCtrlId(&opts.ctrlId)
		ctrlIds = append(ctrlIds, opts.ctrlId)

		if idx == 0 { // first radio of the group?
			opts.wndStyle |= co.WS_TABSTOP | co.WS_GROUP
		}

		me.radios = append(me.radios, &RadioButton{
			_BaseCtrl: newBaseCtrl(opts.ctrlId),
			events:    EventsButton{opts.ctrlId, &parent.base().userEvents},
			index:     uint(idx),
		})
	}
	me.events = EventsRadioGroup{me, &parent.base().userEvents}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		for idx, opts := range allOpts {
			if opts.size.Cx == 0 && opts.size.Cy == 0 {
				opts.size, _ = calcTextBoundBoxWithCheck(utl.RemoveAccelAmpersands(opts.text))
			}
			me.radios[idx].createWindow(opts.wndExStyle, "BUTTON", opts.text,
				opts.wndStyle|co.WS(opts.ctrlStyle), opts.position, opts.size, parent, true)
			parent.base().layout.Add(parent, me.radios[idx].hWnd, opts.layout)
			if opts.selected {
				me.radios[idx].Select()
			}
		}
	})

	return me
}

// Instantiates the [RadioButton] controls to be loaded from a dialog resource
// with [win.HWND.GetDlgItem].
//
// Panics if the number of radio buttons is zero.
func NewRadioGroupDlg(parent Parent, layout LAY, ctrlIds ...uint16) *RadioGroup {
	if len(ctrlIds) == 0 {
		panic("Cannot create a RadioGroup without radio buttons.")
	}

	me := &RadioGroup{
		radios: make([]*RadioButton, 0, len(ctrlIds)),
	}

	for idx, ctrlId := range ctrlIds {
		me.radios = append(me.radios, &RadioButton{
			_BaseCtrl: newBaseCtrl(ctrlId),
			events:    EventsButton{ctrlId, &parent.base().userEvents},
			index:     uint(idx),
		})
	}
	me.events = EventsRadioGroup{me, &parent.base().userEvents}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		for _, radio := range me.radios {
			radio.assignDialog(parent)
			parent.base().layout.Add(parent, radio.hWnd, layout)
		}
	})

	return me
}

// Exposes all the control notifications the can be handled for all
// [RadioButton] controls at once.
//
// Panics if called after the controls have been created.
func (me *RadioGroup) On() *EventsRadioGroup {
	if me.radios[0].hWnd != 0 {
		panic("Cannot add event handling after the control has been created.")
	}
	return &me.events
}

// Returns the number of [RadioButton] controls.
func (me *RadioGroup) Count() uint {
	return uint(len(me.radios))
}

// Enables or disables all [RadioButton] controls at once with
// [win.HWND.EnableWindow].
//
// Returns the same object, so further operations can be chained.
func (me *RadioGroup) Enable(enable bool) *RadioGroup {
	for _, radio := range me.radios {
		radio.hWnd.EnableWindow(enable)
	}
	return me
}

// Returns the [RadioButton] at the given index.
func (me *RadioGroup) Get(index uint) *RadioButton {
	return me.radios[index]
}

// Returns the [RadioButton] with the given control ID, or nil if none.
func (me *RadioGroup) GetById(ctrlId uint16) *RadioButton {
	for _, radio := range me.radios {
		if radio.CtrlId() == ctrlId {
			return radio
		}
	}
	return nil
}

// Returns the selected [RadioButton], or nil if none.
func (me *RadioGroup) Selected() *RadioButton {
	for _, radio := range me.radios {
		if radio.IsSelected() {
			return radio
		}
	}
	return nil
}

// Native [radio button] group events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [radio button]: https://learn.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type EventsRadioGroup struct {
	radioGroup   *RadioGroup
	parentEvents *EventsWindow
}

// [BN_CLICKED] message handler.
//
// [BN_CLICKED]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *EventsRadioGroup) BnClicked(fun func(radio *RadioButton)) {
	for _, radio := range me.radioGroup.radios {
		radio := radio
		me.parentEvents.WmCommand(radio.ctrlId, co.BN_CLICKED, func() {
			fun(radio)
		})
	}
}

// [BN_DBLCLK] message handler.
//
// [BN_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-dblclk
func (me *EventsRadioGroup) BnDblClk(fun func(radio *RadioButton)) {
	for _, radio := range me.radioGroup.radios {
		radio := radio
		me.parentEvents.WmCommand(radio.ctrlId, co.BN_DBLCLK, func() {
			fun(radio)
		})
	}
}

// [BN_KILLFOCUS] message handler.
//
// [BN_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-killfocus
func (me *EventsRadioGroup) BnKillFocus(fun func(radio *RadioButton)) {
	for _, radio := range me.radioGroup.radios {
		radio := radio
		me.parentEvents.WmCommand(radio.ctrlId, co.BN_KILLFOCUS, func() {
			fun(radio)
		})
	}
}

// [BN_SETFOCUS] message handler.
//
// [BN_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/controls/bn-setfocus
func (me *EventsRadioGroup) BnSetFocus(fun func(radio *RadioButton)) {
	for _, radio := range me.radioGroup.radios {
		radio := radio
		me.parentEvents.WmCommand(radio.ctrlId, co.BN_SETFOCUS, func() {
			fun(radio)
		})
	}
}
