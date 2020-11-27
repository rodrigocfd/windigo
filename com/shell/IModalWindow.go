/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
	IModalWindow struct{ win.IUnknown }

	IModalWindowVtbl struct {
		win.IUnknownVtbl
		Show uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*IModalWindowVtbl)(unsafe.Pointer(*me.Ppv)).Show, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndOwner), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IModalWindow.Show"))
	}
}
