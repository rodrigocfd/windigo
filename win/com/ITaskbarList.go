/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"syscall"
	"wingows/co"
	"wingows/win"
)

type (
	_ITaskbarList struct{ _IUnknown }

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
	//
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

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *_ITaskbarList) ActivateTab(hwnd win.HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.ActivateTab, 1, me.uintptr, 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.ActivateTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *_ITaskbarList) SetActiveAlt(hwnd win.HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetActiveAlt, 1, me.uintptr, 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.SetActiveAlt").Error())
	}
}
