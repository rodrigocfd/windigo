/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"fmt"
	"wingows/co"
	"wingows/win"
)

// Manages a group of native radio buttons.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/button-types-and-styles#radio-buttons
type RadioGroup struct {
	radios []RadioButton
}

// Calls CreateWindowEx() to add a new radio button with BS_AUTORADIOBUTTON,
// and WS_GROUP if the first one.
//
// Position and size will be adjusted to the current system DPI.
func (me *RadioGroup) Add(
	parent Window, ctrlId, x, y int32, text string) *RadioGroup {

	me.radios = append(me.radios, RadioButton{})
	newRad := &me.radios[len(me.radios)-1]

	if len(me.radios) == 1 {
		newRad.CreateFirst(parent, ctrlId, x, y, text).
			SetCheck() // first one is checked by default
	} else {
		newRad.CreateSubsequent(parent, ctrlId, x, y, text)
	}

	return me
}

// Returns a slice of Control with all radio buttons.
func (me *RadioGroup) AsSlice() []Control {
	ctrls := make([]Control, 0, len(me.radios))
	for i := range me.radios {
		ctrls = append(ctrls, &me.radios[i])
	}
	return ctrls
}

// Returns the ID of the currently selected radio button, if any.
func (me *RadioGroup) CheckedId() (int32, bool) {
	for i := range me.radios {
		if me.radios[i].IsChecked() {
			return me.radios[i].Id(), true
		}
	}
	return 0, false
}

// Sets BST_UNCHECKED on all radio buttons.
func (me *RadioGroup) ClearChecks() *RadioGroup {
	for i := range me.radios {
		me.radios[i].Hwnd().
			SendMessage(co.WM(co.BM_SETCHECK),
				win.WPARAM(co.BST_UNCHECKED), 0)
	}
	return me
}

// Returns the radio button at the given index.
//
// Does not perform bound checking.
func (me *RadioGroup) RadioButton(index uint32) *RadioButton {
	return &me.radios[index]
}

// Sets the currently checked radio button.
func (me *RadioGroup) SetCheck(radioButtonId int32) *RadioGroup {
	me.ClearChecks()
	me.radioById(radioButtonId).SetCheck()
	return me
}

// Sets the currently checked radio button, and emulates the user click.
func (me *RadioGroup) SetCheckAndTrigger(radioButtonId int32) *RadioGroup {
	me.SetCheck(radioButtonId)
	me.radioById(radioButtonId).SetCheckAndTrigger()
	return me
}

func (me *RadioGroup) radioById(radioButtonId int32) *RadioButton {
	for i := range me.radios {
		if me.radios[i].Id() == radioButtonId {
			return &me.radios[i]
		}
	}
	panic(fmt.Sprintf("Radio with ID %d doesn't exist.", radioButtonId))
}
