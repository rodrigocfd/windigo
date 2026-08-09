//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IOleWindow] COM interface.
//
// [IOleWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-iolewindow
type IOleWindow struct{ win.IUnknown }

type _IOleWindowVt struct {
	utl.IUnknownVt
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IOleWindow) IID() *co.IID {
	return &cosh.IID_IOleWindow
}

// [ContextSensitiveHelp] method.
//
// [ContextSensitiveHelp]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-iolewindow-contextsensitivehelp
func (me *IOleWindow) ContextSensitiveHelp() (bool, error) {
	var bVal win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IOleWindowVt](me.Ppvt()).ContextSensitiveHelp,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&bVal)))
	return utl.HresultToBoolError(int32(bVal), ret)
}

// [GetWindow] method.
//
// [GetWindow]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-iolewindow-getwindow
func (me *IOleWindow) GetWindow() (win.HWND, error) {
	return utl.OleCallReturnStruct[win.HWND](me,
		utl.Vt[_IOleWindowVt](me.Ppvt()).GetWindow)
}
