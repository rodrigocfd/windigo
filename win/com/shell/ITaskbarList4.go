package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _ITaskbarList4Vtbl struct {
	_ITaskbarList3Vtbl
	SetTabProperties uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct {
	ITaskbarList3 // Base ITaskbarList3 > ITaskbarList2 > ITaskbarList > IUnknown.
}

// Calls CoCreateInstance(), typically with CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer ITaskbarList4.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func NewITaskbarList4(dwClsContext co.CLSCTX) ITaskbarList4 {
	iUnk := win.CoCreateInstance(
		shellco.CLSID_TaskbarList, nil, dwClsContext,
		shellco.IID_ITaskbarList4)
	return ITaskbarList4{
		ITaskbarList3{
			ITaskbarList2{
				ITaskbarList{IUnknown: iUnk},
			},
		},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist4-settabproperties
func (me *ITaskbarList4) SetProperties(
	hwndTab win.HWND, stpFlags shellco.STPFLAG) {

	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList4Vtbl)(unsafe.Pointer(*me.Ppv)).SetTabProperties, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwndTab), uintptr(stpFlags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
