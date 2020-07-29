/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	baseITaskbarList struct{ baseIUnknown }

	// ITaskbarList > IUnknown.
	ITaskbarList struct{ baseITaskbarList }

	tvbITaskbarList struct {
		vtbIUnknown
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)
