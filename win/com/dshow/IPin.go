package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IPinVtbl struct {
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

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
type IPin struct {
	win.IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-beginflush
func (me *IPin) BeginFlush() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).BeginFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
func (me *IPin) Disconnect() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).Disconnect, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
func (me *IPin) EndFlush() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
func (me *IPin) EndOfStream() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndOfStream, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
}
