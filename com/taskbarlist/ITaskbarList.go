/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package taskbarlist

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
	//
	// ITaskbarList > IUnknown.
	ITaskbarList struct{ win.IUnknown }

	ITaskbarListVtbl struct {
		win.IUnknownVtbl
		HrInit       uintptr
		AddTab       uintptr
		DeleteTab    uintptr
		ActivateTab  uintptr
		SetActiveAlt uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *ITaskbarList) CoCreateInstance(dwClsContext co.CLSCTX) *ITaskbarList {
	ppv, err := win.CoCreateInstance(
		win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a, 0x006097c9a090), // CLSID_TaskbarList
		nil,
		dwClsContext,
		win.NewGuid(0x56fdf342, 0xfd6d, 0x11d0, 0x958a, 0x006097c9a090)) // IID_ITaskbarList

	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "CoCreateInstance/ITaskbarList"))
	}
	me.Ppv = ppv
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() *ITaskbarList {
	syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).HrInit, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hwnd win.HWND) *ITaskbarList {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).AddTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.AddTab").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hwnd win.HWND) *ITaskbarList {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).DeleteTab, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.DeleteTab").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hwnd win.HWND) *ITaskbarList {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).ActivateTab, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.ActivateTab").Error())
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hwnd win.HWND) *ITaskbarList {
	ret, _, _ := syscall.Syscall(
		(*ITaskbarListVtbl)(unsafe.Pointer(*me.Ppv)).SetActiveAlt, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList.SetActiveAlt").Error())
	}
	return me
}
