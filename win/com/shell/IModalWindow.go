//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IModalWindow] COM interface.
//
// [IModalWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow interface {
	com.IUnknown

	// [Show] COM method.
	//
	// Returns false if user cancelled.
	//
	// [Show]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
	Show(hwndOwner win.HWND) bool
}

type _IModalWindow struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IModalWindow.Release().
func NewIModalWindow(base com.IUnknown) IModalWindow {
	return &_IModalWindow{IUnknown: base}
}

func (me *_IModalWindow) Show(hwndOwner win.HWND) bool {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IModalWindow)(unsafe.Pointer(*me.Ptr())).Show,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndOwner))

	if hr := errco.ERROR(ret & 0xffff); hr == errco.S_OK { // HRESULT_FROM_WIN32()
		return true
	} else if hr == errco.CANCELLED {
		return false
	} else {
		panic(hr)
	}
}
