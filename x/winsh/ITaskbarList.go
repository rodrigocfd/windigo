//go:build windows

package winsh

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [ITaskbarList] COM interface.
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
//	var taskbl *winsh.ITaskbarList
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [ITaskbarList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList struct{ win.IUnknown }

type _ITaskbarListVt struct {
	utl.IUnknownVt
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskbarList) IID() *co.IID {
	return &cosh.IID_ITaskbarList
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *ITaskbarList) AddRef(releaser *win.OleReleaser) *ITaskbarList {
	return utl.OleNewFromAddRef[*ITaskbarList](me, releaser)
}

// [ActivateTab] method.
//
// [ActivateTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarListVt](me.Ppvt()).ActivateTab,
		me.Ppvt(),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [AddTab] method.
//
// [AddTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarListVt](me.Ppvt()).AddTab,
		me.Ppvt(),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [DeleteTab] method.
//
// [DeleteTab]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarListVt](me.Ppvt()).DeleteTab,
		me.Ppvt(),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [HrInit] method.
//
// [HrInit]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_ITaskbarListVt](me.Ppvt()).HrInit)
}

// [SetActiveAlt] method.
//
// [SetActiveAlt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskbarListVt](me.Ppvt()).SetActiveAlt,
		me.Ppvt(),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}
