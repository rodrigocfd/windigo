//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// [ITaskbarList2] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	taskbl, _ := ole.CoCreateInstance[shell.ITaskbarList2](
//		rel,
//		co.CLSID_TaskbarList,
//		co.CLSCTX_INPROC_SERVER,
//	)
//
// [ITaskbarList2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct{ ITaskbarList }

// Returns the unique COM interface identifier.
func (*ITaskbarList2) IID() co.IID {
	return co.IID_ITaskbarList2
}

// [MarkFullscreenWindow] method.
//
// [MarkFullscreenWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(hwnd win.HWND, fullScreen bool) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList2)(unsafe.Pointer(*me.Ppvt())).MarkFullscreenWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwnd), util.BoolToUintptr(fullScreen))
	return util.ErrorAsHResult(ret)
}
