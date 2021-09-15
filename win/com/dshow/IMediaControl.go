package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IMediaControlVtbl struct {
	win.IDispatchVtbl
	Run                    uintptr
	Pause                  uintptr
	Stop                   uintptr
	GetState               uintptr
	RenderFile             uintptr
	AddSourceFilter        uintptr
	GetFilterCollection    uintptr
	GetRegFilterCollection uintptr
	StopWhenReady          uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-imediacontrol
type IMediaControl struct {
	win.IDispatch // Base IDispatch > IUnknown.
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-getstate
func (me *IMediaControl) GetState(msTimeout int) (dshowco.FILTER_STATE, error) {
	state := dshowco.FILTER_STATE(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).GetState, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
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
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Pause, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
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
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).RenderFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-run
func (me *IMediaControl) Run() bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Run, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
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
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Stop, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stopwhenready
func (me *IMediaControl) StopWhenReady() bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).StopWhenReady, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
