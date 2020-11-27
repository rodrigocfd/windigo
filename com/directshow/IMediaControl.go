/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// IMediaControl > IPersist > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-imediacontrol
	IMediaControl struct{ IPersist }

	IMediaControlVtbl struct {
		IPersistVtbl
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
)

// https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-pause
func (me *IMediaControl) Pause() {
	ret, _, _ := syscall.Syscall(
		(*IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Pause, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK && lerr != co.ERROR_S_FALSE {
		panic(win.NewWinError(lerr, "IMediaControl.Pause"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-renderfile
func (me *IMediaControl) RenderFile(strFilename string) {
	ret, _, _ := syscall.Syscall(
		(*IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).RenderFile, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(strFilename))), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMediaControl.RenderFile"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-run
func (me *IMediaControl) Run() {
	ret, _, _ := syscall.Syscall(
		(*IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Run, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK && lerr != co.ERROR_S_FALSE {
		panic(win.NewWinError(lerr, "IMediaControl.Run"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stop
func (me *IMediaControl) Stop() {
	ret, _, _ := syscall.Syscall(
		(*IMediaControlVtbl)(unsafe.Pointer(*me.Ppv)).Stop, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IMediaControl.Stop"))
	}
}
