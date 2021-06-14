package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/err"
)

type _ITaskbarList2Vtbl struct {
	_ITaskbarListVtbl
	MarkFullscreenWindow uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct {
	ITaskbarList // Base ITaskbarList > IUnknown.
}

// Calls IUnknown.CoCreateInstance() to return ITaskbarList2.
//
// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateITaskbarList2(dwClsContext co.CLSCTX) ITaskbarList2 {
	iUnk := win.CoCreateInstance(
		shellco.CLSID_TaskbarList, nil, dwClsContext,
		shellco.IID_ITaskbarList2)
	return ITaskbarList2{
		ITaskbarList{IUnknown: iUnk},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
func (me *ITaskbarList2) MarkFullscreenWindow(hwnd win.HWND, fFullScreen bool) {
	ret, _, _ := syscall.Syscall(
		(*_ITaskbarList2Vtbl)(unsafe.Pointer(*me.Ppv)).MarkFullscreenWindow, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hwnd), util.BoolToUintptr(fFullScreen))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
