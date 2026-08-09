//go:build windows

package winsh

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [ITaskbarList2] COM interface.
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
//	var taskbl *winsh.ITaskbarList2
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct{ ITaskbarList }

type _ITaskbarList2Vt struct {
	_ITaskbarListVt
	MarkFullscreenWindow uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList2) IID() *co.IID {
	return &cosh.IID_ITaskbarList2
}

// [MarkFullscreenWindow] method.
//
// [MarkFullscreenWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(hwnd win.HWND, fullScreen bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarList2Vt](me.Ppvt()).MarkFullscreenWindow,
		me.Ppvt(),
		uintptr(hwnd),
		utl.BoolToUintptr(fullScreen))
	return utl.HresultToError(ret)
}
