//go:build windows

package winsh

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IModalWindow] COM interface.
//
// [IModalWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct{ win.IUnknown }

type _IModalWindowVt struct {
	utl.IUnknownVt
	Show uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IModalWindow) IID() *co.IID {
	return &cosh.IID_IModalWindow
}

// [Show] method.
//
// Returns false if user cancelled.
//
// [Show]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-imodalwindow-show
func (me *IModalWindow) Show(hwndOwner win.HWND) (bool, error) {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IModalWindowVt](me.Ppvt()).Show,
		me.Ppvt(),
		uintptr(hwndOwner))

	if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
		return true, nil
	} else if wErr == co.ERROR_CANCELLED {
		return false, nil
	} else {
		return false, wErr.ToHresult()
	}
}
