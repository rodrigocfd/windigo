//go:build windows

package winsh

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [ITaskbarList4] COM interface.
//
// Implements [OleResource].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var taskbl *winsh.ITaskbarList4
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList4]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct{ ITaskbarList3 }

type _ITaskbarList4Vt struct {
	_ITaskbarList3Vt
	SetTabProperties uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList4) IID() *co.IID {
	return &cosh.IID_ITaskbarList4
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist4-settabproperties
func (me *ITaskbarList4) SetProperties(hwndTab win.HWND, flags cosh.STPFLAG) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList4Vt](me.Ppvt()).SetTabProperties,
		me.Ppvt(),
		uintptr(hwndTab),
		uintptr(flags))
	return utl.HresultToError(ret)
}
