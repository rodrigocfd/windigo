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
	// IMediaFilter > IPersist > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
	IMediaFilter struct{ IPersist }

	IMediaFilterVtbl struct {
		IPersistVtbl
		Stop          uintptr
		Pause         uintptr
		Run           uintptr
		GetState      uintptr
		SetSyncSource uintptr
		GetSyncSource uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-pause
func (me *IMediaFilter) Pause() {
	ret, _, _ := syscall.Syscall(
		(*IMediaFilterVtbl)(unsafe.Pointer(*me.Ppv)).Pause, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK && lerr != co.ERROR_S_FALSE {
		panic(win.NewWinError(lerr, "IMediaFilter.Pause"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-stop
func (me *IMediaFilter) Stop() {
	ret, _, _ := syscall.Syscall(
		(*IMediaFilterVtbl)(unsafe.Pointer(*me.Ppv)).Stop, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK && lerr != co.ERROR_S_FALSE {
		panic(win.NewWinError(lerr, "IMediaFilter.Stop"))
	}
}
