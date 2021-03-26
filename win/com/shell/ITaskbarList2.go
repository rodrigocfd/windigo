package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
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
func CoCreateITaskbarList2(dwClsContext co.CLSCTX) (ITaskbarList2, error) {
	clsidTaskbarList := win.NewGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a, 0x006097c9a090)
	iidITaskbarList2 := win.NewGuid(0x602d4995, 0xb13a, 0x429b, 0xa66e, 0x1935e44f4317)

	iUnk, lerr := win.CoCreateInstance(
		clsidTaskbarList, nil, dwClsContext, iidITaskbarList2)
	if lerr != nil {
		return ITaskbarList2{}, lerr
	}
	return ITaskbarList2{
		ITaskbarList{IUnknown: iUnk},
	}, nil
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
