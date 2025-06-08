//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IOleWindow] COM interface.
//
// Implements [ole.ComObj] and [ole.ComResource].
//
// [IOleWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-iolewindow
type IOleWindow struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IOleWindow) IID() co.IID {
	return co.IID_IOleWindow
}

// [ContextSensitiveHelp] method.
//
// [ContextSensitiveHelp]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-iolewindow-contextsensitivehelp
func (me *IOleWindow) ContextSensitiveHelp() (bool, error) {
	var bVal int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IOleWindowVt)(unsafe.Pointer(*me.Ppvt())).ContextSensitiveHelp,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bVal)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bVal != 0, nil
	} else {
		return false, hr
	}
}

// [GetWindow] method.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-iolewindow-getwindow
func (me *IOleWindow) GetWindow() (win.HWND, error) {
	var hWnd win.HWND
	ret, _, _ := syscall.SyscallN(
		(*_IOleWindowVt)(unsafe.Pointer(*me.Ppvt())).GetWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hWnd)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hWnd, nil
	} else {
		return win.HWND(0), hr
	}
}

type _IOleWindowVt struct {
	ole.IUnknownVt
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
}
