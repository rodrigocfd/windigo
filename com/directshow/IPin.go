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
	// IPin > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
	IPin struct{ win.IUnknown }

	IPinVtbl struct {
		win.IUnknownVtbl
		Connect                  uintptr
		ReceiveConnection        uintptr
		Disconnect               uintptr
		ConnectedTo              uintptr
		ConnectionMediaType      uintptr
		QueryPinInfo             uintptr
		QueryDirection           uintptr
		QueryId                  uintptr
		QueryAccept              uintptr
		EnumMediaTypes           uintptr
		QueryInternalConnections uintptr
		EndOfStream              uintptr
		BeginFlush               uintptr
		EndFlush                 uintptr
		NewSegment               uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-beginflush
func (me *IPin) BeginFlush() {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).BeginFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.BeginFlush"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
func (me *IPin) Disconnect() {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).Disconnect, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.Disconnect"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
func (me *IPin) EndFlush() {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.EndFlush"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
func (me *IPin) EndOfStream() {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndOfStream, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.EndOfStream"))
	}
}
