//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [ITaskbarList] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var taskbl *win.ITaskbarList
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList) IID() co.IID {
	return co.IID_ITaskbarList
}

// [ActivateTab] method.
//
// [ActivateTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).ActivateTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

// [AddTab] method.
//
// [AddTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).AddTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

// [DeleteTab] method.
//
// [DeleteTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).DeleteTab,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

// [HrInit] method.
//
// [HrInit]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).HrInit,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SetActiveAlt] method.
//
// [SetActiveAlt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskbarListVt)(unsafe.Pointer(*me.Ppvt())).SetActiveAlt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

type _ITaskbarListVt struct {
	_IUnknownVt
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}
