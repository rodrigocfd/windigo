/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
)

// Native combo box control.
type ComboBox struct {
	controlNativeBase
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them. Position and size will be adjusted to
// the current system DPI.
func (me *ComboBox) Create(parent Window, x, y int32, width, height uint32,
	exStyles co.WS_EX, styles co.WS, cbStyles co.CBS) *ComboBox {

	x, y, width, height = globalDpi.multiply(x, y, width, height)

	me.controlNativeBase.create(exStyles, "COMBOBOX", "",
		styles|co.WS(cbStyles), x, y, width, height, parent)

	return me
}
