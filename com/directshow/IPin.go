/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
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
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
	//
	// IPin > IUnknown.
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

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
func (me *IPin) Disconnect() *IPin {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).Disconnect, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.Disconnect"))
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
func (me *IPin) EndOfStream() *IPin {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndOfStream, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.EndOfStream"))
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-beginflush
func (me *IPin) BeginFlush() *IPin {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).BeginFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.BeginFlush"))
	}
	return me
}

// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
func (me *IPin) EndFlush() *IPin {
	ret, _, _ := syscall.Syscall(
		(*IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IPin.EndFlush"))
	}
	return me
}
