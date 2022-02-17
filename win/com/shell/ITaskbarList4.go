package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct{ ITaskbarList3 }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList4.Release().
//
// Example:
//
//  taskbl4 := shell.NewITaskbarList4(
//      com.CoCreateInstance(
//          shellco.CLSID_TaskbarList, nil,
//          comco.CLSCTX_INPROC_SERVER,
//          shellco.IID_ITaskbarList4),
//  )
//  defer taskbl4.Release()
func NewITaskbarList4(base com.IUnknown) ITaskbarList4 {
	return ITaskbarList4{ITaskbarList3: NewITaskbarList3(base)}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist4-settabproperties
func (me *ITaskbarList4) SetProperties(
	hwndTab win.HWND, flags shellco.STPFLAG) {

	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList4)(unsafe.Pointer(*me.Ptr())).SetTabProperties, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
