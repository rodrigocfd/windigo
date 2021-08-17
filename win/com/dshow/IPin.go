package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
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

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// ⚠️ You must defer IPin.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connect
func (me *IPin) Connect(pmt *AM_MEDIA_TYPE) (IPin, error) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).Connect, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)),
		uintptr(unsafe.Pointer(pmt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IPin{
			win.IUnknown{Ppv: ppQueried},
		}, nil
	} else {
		return IPin{}, hr
	}
}

// ⚠️ You must defer IPin.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectedto
func (me *IPin) ConnectedTo() (IPin, error) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).ConnectedTo, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IPin{
			win.IUnknown{Ppv: ppQueried},
		}, nil
	} else {
		return IPin{}, hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectionmediatype
func (me *IPin) ConnectionMediaType(pmt *AM_MEDIA_TYPE) error {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).ConnectionMediaType, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pmt)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
func (me *IPin) Disconnect() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).Disconnect, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
func (me *IPin) EndFlush() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndFlush, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
func (me *IPin) EndOfStream() {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).EndOfStream, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// ⚠️ You must defer IEnumMediaTypes.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-enummediatypes
func (me *IPin) EnumMediaTypes() (IEnumMediaTypes, error) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).EnumMediaTypes, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IEnumMediaTypes{
			win.IUnknown{Ppv: ppQueried},
		}, nil
	} else {
		return IEnumMediaTypes{}, hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-queryaccept
func (me *IPin) QueryAccept(pmt *AM_MEDIA_TYPE) (bool, error) {
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).QueryAccept, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(pmt)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK || hr == errco.S_FALSE {
		return hr == errco.S_OK, nil
	} else {
		return false, hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-querydirection
func (me *IPin) QueryDirection() dshowco.PIN_DIRECTION {
	var pPinDir dshowco.PIN_DIRECTION
	ret, _, _ := syscall.Syscall(
		(*_IPinVtbl)(unsafe.Pointer(*me.Ppv)).QueryDirection, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pPinDir)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pPinDir
	} else {
		panic(hr)
	}
}
