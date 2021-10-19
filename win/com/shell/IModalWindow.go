package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IModalWindowVtbl struct {
	win.IUnknownVtbl
	Show uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ win.IUnknown }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IModalWindow.Release().
func NewIModalWindow(ptr win.IUnknownPtr) IModalWindow {
	return IModalWindow{
		IUnknown: win.NewIUnknown(ptr),
	}
}

// Returns false if user cancelled.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner win.HWND) bool {
	ret, _, _ := syscall.Syscall(
		(*_IModalWindowVtbl)(unsafe.Pointer(*me.Ptr())).Show, 2,
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
