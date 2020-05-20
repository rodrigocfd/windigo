/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	c "wingows/consts"
)

// All global variables in the package are kept here for tighter control.

var (
	globalUiFont = Font{} // WindowMain: created in RunAsMain(), freed in runMainLoop()

	globalBaseCtrlId = c.ID(1000) // controlIdGuard; arbitrary, taken from Visual Studio resource editor

	globalBaseSubclassId  = uint32(0)  // controlNativeBase; incremented at each subclass installed
	globalSubclassProcPtr = uintptr(0) // controlNativeBase; necessary for RemoveWindowSubclass

	globalDpiRatioX, globalDpiRatioY float32 = -1, -1 // multiplyByDpi()
)
