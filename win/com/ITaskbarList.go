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
	_ITaskbarListImpl struct{ _IUnknownImpl }

	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
	//
	// ITaskbarList > IUnknown.
	ITaskbarList struct{ _ITaskbarListImpl }

	_ITaskbarListVtbl struct {
		_IUnknownVtbl
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)

func (me *_ITaskbarListImpl) CoCreateInstance(dwClsContext co.CLSCTX) {
	me.coCreateInstancePtr(
		&co.CLSID_TaskbarList, dwClsContext, &co.IID_ITaskbarList)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *_ITaskbarListImpl) ActivateTab(hwnd win.HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.ActivateTab, 1, me.uintptr, 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.ActivateTab").Error())
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *_ITaskbarListImpl) SetActiveAlt(hwnd win.HWND) {
	vTbl := (*_ITaskbarListVtbl)(me.pVtbl())
	ret, _, _ := syscall.Syscall(vTbl.SetActiveAlt, 1, me.uintptr, 0, 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.SetActiveAlt").Error())
	}
}
