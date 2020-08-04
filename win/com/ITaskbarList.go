/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"wingows/co"
)

type (
	_ITaskbarList struct{ _IUnknown }

	// ITaskbarList > IUnknown.
	ITaskbarList struct{ _ITaskbarList }

	_ITaskbarListVtbl struct {
		_IUnknownVtbl
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)

func (me *_ITaskbarList) CoCreateInstance(dwClsContext co.CLSCTX) {
	me._IUnknown.coCreateInstance(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList)
}
