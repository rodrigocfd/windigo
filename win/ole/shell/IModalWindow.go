//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IModalWindow] COM interface.
//
// [IModalWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ ole.IUnknown }

// Returns the unique COM interface identifier.
func (*IModalWindow) IID() co.IID {
	return co.IID_IModalWindow
}

// [Show] method.
//
// Returns false if user cancelled.
//
// [Show]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner win.HWND) (bool, error) {
	ret, _, _ := syscall.SyscallN(
		(*vt.IModalWindow)(unsafe.Pointer(*me.Ppvt())).Show,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndOwner))

	if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
		return true, nil
	} else if wErr == co.ERROR_CANCELLED {
		return false, nil
	} else {
		return false, wErr.HResult()
	}
}
