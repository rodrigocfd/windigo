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

// [ITaskbarList4] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	taskbl, _ := ole.CoCreateInstance[shell.ITaskbarList4](
//		rel,
//		co.CLSID_TaskbarList,
//		co.CLSCTX_INPROC_SERVER,
//	)
//
// [ITaskbarList4]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct{ ITaskbarList3 }

// Returns the unique COM interface identifier.
func (*ITaskbarList4) IID() co.IID {
	return co.IID_ITaskbarList4
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist4-settabproperties
func (me *ITaskbarList4) SetProperties(hwndTab win.HWND, flags co.STPFLAG) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.ITaskbarList4)(unsafe.Pointer(*me.Ppvt())).SetTabProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab), uintptr(flags))
	return util.ErrorAsHResult(ret)
}
