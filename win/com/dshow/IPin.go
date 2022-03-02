package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
type IPin interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-beginflush
	BeginFlush()

	// ⚠️ You must defer IPin.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connect
	Connect(mt *AM_MEDIA_TYPE) (IPin, error)

	// ⚠️ You must defer IPin.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectedto
	ConnectedTo() (IPin, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectionmediatype
	ConnectionMediaType(mt *AM_MEDIA_TYPE) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
	Disconnect()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
	EndFlush()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
	EndOfStream()

	// ⚠️ You must defer IEnumMediaTypes.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-enummediatypes
	EnumMediaTypes() (IEnumMediaTypes, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-queryaccept
	QueryAccept(mt *AM_MEDIA_TYPE) (bool, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-querydirection
	QueryDirection() dshowco.PIN_DIRECTION
}

type _IPin struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPin.Release().
func NewIPin(base com.IUnknown) IPin {
	return &_IPin{IUnknown: base}
}

func (me *_IPin) BeginFlush() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).BeginFlush, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) Connect(mt *AM_MEDIA_TYPE) (IPin, error) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).Connect, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(com.NewIUnknown(ppQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IPin) ConnectedTo() (IPin, error) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).ConnectedTo, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(com.NewIUnknown(ppQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IPin) ConnectionMediaType(mt *AM_MEDIA_TYPE) error {
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

func (me *_IPin) Disconnect() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).Disconnect, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) EndFlush() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EndFlush, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) EndOfStream() {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EndOfStream, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) EnumMediaTypes() (IEnumMediaTypes, error) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EnumMediaTypes, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumMediaTypes(com.NewIUnknown(ppQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IPin) QueryAccept(mt *AM_MEDIA_TYPE) (bool, error) {
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

func (me *_IPin) QueryDirection() dshowco.PIN_DIRECTION {
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
