//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [ITaskbarList] COM interface.
//
// [ITaskbarList]: https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList interface {
	com.IUnknown

	// [ActivateTab] COM method.
	//
	// [ActivateTab]: https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
	ActivateTab(hWnd win.HWND)

	// [AddTab] COM method.
	//
	// [AddTab]: https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
	AddTab(hWnd win.HWND)

	// [DeleteTab] COM method.
	//
	// [DeleteTab]: https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
	DeleteTab(hWnd win.HWND)

	// [HrInit] COM method.
	//
	// [HrInit]: https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
	HrInit()

	// [SetActiveAlt] COM method.
	//
	// [SetActiveAlt]: https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
	SetActiveAlt(hWnd win.HWND)
}

type _ITaskbarList struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList.Release().
//
// Example:
//
//	taskbl := shell.NewITaskbarList(
//		com.CoCreateInstance(
//			shellco.CLSID_TaskbarList, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			shellco.IID_ITaskbarList),
//	)
//	defer taskbl.Release()
func NewITaskbarList(base com.IUnknown) ITaskbarList {
	return &_ITaskbarList{IUnknown: base}
}

func (me *_ITaskbarList) ActivateTab(hWnd win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).ActivateTab,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) AddTab(hWnd win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).AddTab,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) DeleteTab(hWnd win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).DeleteTab,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) HrInit() {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).HrInit,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ITaskbarList) SetActiveAlt(hWnd win.HWND) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).SetActiveAlt,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
