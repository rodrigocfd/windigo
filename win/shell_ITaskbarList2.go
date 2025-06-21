//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [ITaskbarList2] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// # Example
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var taskbl *win.ITaskbarList2
//	win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct{ ITaskbarList }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList2) IID() co.IID {
	return co.IID_ITaskbarList2
}

// [MarkFullscreenWindow] method.
//
// [MarkFullscreenWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(hwnd HWND, fullScreen bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList2Vt)(unsafe.Pointer(*me.Ppvt())).MarkFullscreenWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwnd),
		utl.BoolToUintptr(fullScreen))
	return utl.ErrorAsHResult(ret)
}

type _ITaskbarList2Vt struct {
	_ITaskbarListVt
	MarkFullscreenWindow uintptr
}
