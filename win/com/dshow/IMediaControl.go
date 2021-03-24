package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IMediaControlVtbl struct {
	_IDispatchVtbl
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
	IDispatch // Base IDispatch > IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-pause
func (me *IMediaControl) Pause() {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Pause, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK && lerr != err.S_FALSE {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-renderfile
func (me *IMediaControl) RenderFile(strFilename string) {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).RenderFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(strFilename))), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-run
func (me *IMediaControl) Run() {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Run, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK && lerr != err.S_FALSE {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stop
func (me *IMediaControl) Stop() {
	ret, _, _ := syscall.Syscall(
		(*_IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Stop, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
