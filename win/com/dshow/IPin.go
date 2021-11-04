package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
type IPin struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPin.Release().
func NewIPin(base win.IUnknown) IPin {
	return IPin{IUnknown: base}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-beginflush
func (me *IPin) BeginFlush() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).BeginFlush, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// ⚠️ You must defer IPin.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connect
func (me *IPin) Connect(mt *AM_MEDIA_TYPE) (IPin, error) {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).Connect, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(ppQueried), nil
	} else {
		return IPin{}, hr
	}
}

// ⚠️ You must defer IPin.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectedto
func (me *IPin) ConnectedTo() (IPin, error) {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).ConnectedTo, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(ppQueried), nil
	} else {
		return IPin{}, hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectionmediatype
func (me *IPin) ConnectionMediaType(mt *AM_MEDIA_TYPE) error {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).ConnectionMediaType, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(mt)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
func (me *IPin) Disconnect() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).Disconnect, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
func (me *IPin) EndFlush() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EndFlush, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
func (me *IPin) EndOfStream() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EndOfStream, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// ⚠️ You must defer IEnumMediaTypes.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-enummediatypes
func (me *IPin) EnumMediaTypes() (IEnumMediaTypes, error) {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EnumMediaTypes, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumMediaTypes(ppQueried), nil
	} else {
		return IEnumMediaTypes{}, hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-queryaccept
func (me *IPin) QueryAccept(mt *AM_MEDIA_TYPE) (bool, error) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).QueryAccept, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(mt)), 0)

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
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).QueryDirection, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pPinDir)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pPinDir
	} else {
		panic(hr)
	}
}
