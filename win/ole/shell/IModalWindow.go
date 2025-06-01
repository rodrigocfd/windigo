//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IModalWindow] COM interface.
//
// Implements [ole.ComObj] and [ole.ComResource].
//
// [IModalWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
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
		(*_IModalWindowVt)(unsafe.Pointer(*me.Ppvt())).Show,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndOwner))

	if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
		return true, nil
	} else if wErr == co.ERROR_CANCELLED {
		return false, nil
	} else {
		return false, wErr.ToHresult()
	}
}

type _IModalWindowVt struct {
	ole.IUnknownVt
	Show uintptr
}
