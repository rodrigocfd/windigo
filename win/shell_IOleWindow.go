//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
)

// [IOleWindow] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IOleWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-iolewindow
type IOleWindow struct{ IUnknown }

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
func (me *IOleWindow) GetWindow() (HWND, error) {
	var hWnd HWND
	ret, _, _ := syscall.SyscallN(
		(*_IOleWindowVt)(unsafe.Pointer(*me.Ppvt())).GetWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hWnd)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hWnd, nil
	} else {
		return HWND(0), hr
	}
}

type _IOleWindowVt struct {
	_IUnknownVt
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
}
