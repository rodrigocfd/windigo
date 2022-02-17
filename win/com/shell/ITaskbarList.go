package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList.Release().
//
// Example:
//
//  taskbl := shell.NewITaskbarList(
//      com.CoCreateInstance(
//          shellco.CLSID_TaskbarList, nil,
//          comco.CLSCTX_INPROC_SERVER,
//          shellco.IID_ITaskbarList),
//  )
//  defer taskbl.Release()
func NewITaskbarList(base com.IUnknown) ITaskbarList {
	return ITaskbarList{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-activatetab
func (me *ITaskbarList) ActivateTab(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).ActivateTab, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-addtab
func (me *ITaskbarList) AddTab(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).AddTab, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-deletetab
func (me *ITaskbarList) DeleteTab(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).DeleteTab, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-hrinit
func (me *ITaskbarList) HrInit() {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).HrInit, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist-setactivealt
func (me *ITaskbarList) SetActiveAlt(hWnd win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList)(unsafe.Pointer(*me.Ptr())).SetActiveAlt, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
