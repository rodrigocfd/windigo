package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-imediacontrol
type IMediaControl struct{ autom.IDispatch }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMediaControl.Release().
//
// Example:
//
//  var gb dshow.IGraphBuilder // initialized somewhere
//
//  mc := dshow.NewIMediaControl(
//      gb.QueryInterface(dshowco.IID_IMediaControl),
//  )
//  defer mc.Release()
func NewIMediaControl(base com.IUnknown) IMediaControl {
	return IMediaControl{IDispatch: autom.NewIDispatch(base)}
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-getstate
func (me *IMediaControl) GetState(msTimeout int) (dshowco.FILTER_STATE, error) {
	var state dshowco.FILTER_STATE
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).GetState, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(msTimeout)), uintptr(unsafe.Pointer(&state)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return state, nil
	} else if hr == errco.VFW_S_STATE_INTERMEDIATE || hr == errco.VFW_S_CANT_CUE {
		return state, hr
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-pause
func (me *IMediaControl) Pause() bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).Pause, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-renderfile
func (me *IMediaControl) RenderFile(fileName string) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).RenderFile, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-run
func (me *IMediaControl) Run() bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).Run, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stop
func (me *IMediaControl) Stop() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).Stop, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stopwhenready
func (me *IMediaControl) StopWhenReady() bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).StopWhenReady, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
