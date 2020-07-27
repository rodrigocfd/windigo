/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

type controlUtilT struct{}

// General control utilities.
var ControlUtil controlUtilT

// Enables or disables many controls at once.
func (controlUtilT) Enable(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}

// Returns the index of the checked radio button within the group, or -1 if none
// is checked.
func (controlUtilT) CheckedRadio(radios []RadioButton) int32 {
	for i := range radios {
		if radios[i].IsChecked() {
			return int32(i)
		}
	}
	return -1
}

// Converts a RadioButton slice into a Control slice.
func (controlUtilT) RadioSlice(radios []RadioButton) []Control {
	ctrls := make([]Control, 0, len(radios))
	for i := range radios {
		ctrls = append(ctrls, &radios[i])
	}
	return ctrls
}
