/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
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

// Returns false if user cancelled.
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner win.HWND) bool {
	ret, _, _ := syscall.Syscall(
		(*IModalWindowVtbl)(unsafe.Pointer(*me.Ppv)).Show, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndOwner), 0)

	if lerr := co.ERROR(ret & 0xffff); lerr == co.ERROR_CANCELLED { // HRESULT_FROM_WIN32()
		return false
	} else if lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IModalWindow.Show"))
	}
	return true
}
