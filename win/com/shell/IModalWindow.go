package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IModalWindow.Release().
func NewIModalWindow(base win.IUnknown) IModalWindow {
	return IModalWindow{IUnknown: base}
}

// Returns false if user cancelled.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner win.HWND) bool {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IModalWindow)(unsafe.Pointer(*me.Ptr())).Show, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndOwner), 0)

	if hr := errco.ERROR(ret & 0xffff); hr == errco.S_OK { // HRESULT_FROM_WIN32()
		return true
	} else if hr == errco.CANCELLED {
		return false
	} else {
		panic(hr)
	}
}
