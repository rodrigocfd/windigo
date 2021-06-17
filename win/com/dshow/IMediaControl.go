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
		uintptr(msTimeout), uintptr(unsafe.Pointer(&state)))

	err := errco.ERROR(ret)
	if err == errco.S_OK {
		return state, nil
	}
	return dshowco.FILTER_STATE(0), err
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-pause
func (me *IMediaControl) Pause() {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Pause, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if err := errco.ERROR(ret); err != errco.S_OK && err != errco.S_FALSE {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-renderfile
func (me *IMediaControl) RenderFile(strFilename string) {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).RenderFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(strFilename))), 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-run
func (me *IMediaControl) Run() {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Run, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if err := errco.ERROR(ret); err != errco.S_OK && err != errco.S_FALSE {
		panic(err)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stop
func (me *IMediaControl) Stop() {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Stop, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
}
