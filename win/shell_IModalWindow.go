//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// [IModalWindow] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IModalWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ IUnknown }

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
func (me *IModalWindow) Show(hwndOwner HWND) (bool, error) {
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
	_IUnknownVt
	Show uintptr
}
