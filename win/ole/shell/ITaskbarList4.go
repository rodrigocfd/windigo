//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// [ITaskbarList4] COM interface.
//
// Implements [ole.ComObj] and [ole.ComResource].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var taskbl *shell.ITaskbarList4
//	ole.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList4]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct{ ITaskbarList3 }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList4) IID() co.IID {
	return co.IID_ITaskbarList4
}

// [SetProperties] method.
//
// [SetProperties]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist4-settabproperties
func (me *ITaskbarList4) SetProperties(hwndTab win.HWND, flags co.STPFLAG) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarList4Vt)(unsafe.Pointer(*me.Ppvt())).SetTabProperties,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hwndTab), uintptr(flags))
	return utl.ErrorAsHResult(ret)
}

type _ITaskbarList4Vt struct {
	_ITaskbarList3Vt
	SetTabProperties uintptr
}
