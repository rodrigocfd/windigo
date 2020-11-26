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
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
	//
	// ITaskbarList2 > ITaskbarList > IUnknown.
	ITaskbarList2 struct{ ITaskbarList }

	ITaskbarList2Vtbl struct {
		ITaskbarListVtbl
		MarkFullscreenWindow uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func (me *ITaskbarList2) CoCreateInstance(dwClsContext co.CLSCTX) *ITaskbarList2 {
	ppv, err := win.CoCreateInstance(
		win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a, 0x006097c9a090), // CLSID_TaskbarList
		nil,
		dwClsContext,
		win.NewGuid(0x602d4995, 0xb13a, 0x429b, 0xa66e, 0x1935e44f4317)) // IID_ITaskbarList2

	if err != co.ERROR_S_OK {
		panic(win.NewWinError(err, "CoCreateInstance/ITaskbarList2"))
	}
	me.Ppv = ppv
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(
	hwnd win.HWND, fFullScreen bool) *ITaskbarList2 {

	ret, _, _ := syscall.Syscall(
		(*ITaskbarList2Vtbl)(unsafe.Pointer(*me.Ppv)).MarkFullscreenWindow, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), _BoolToUintptr(fFullScreen))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "ITaskbarList2.MarkFullscreenWindow").Error())
	}
	return me
}
