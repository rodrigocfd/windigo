package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-isequentialstream
type ISequentialStream interface {
	IUnknown

	// If returned numBytesRead is lower than requested buffer size, it means
	// the end of stream was reached.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-read
	Read(buffer []byte) (numBytesRead uint32)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-write
	Write(data []byte) (numBytesWritten uint32)
}

type _ISequentialStream struct{ IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ISequentialStream.Release().
func NewISequentialStream(base IUnknown) ISequentialStream {
	return &_ISequentialStream{IUnknown: base}
}

func (me *_ISequentialStream) Read(buffer []byte) (numBytesRead uint32) {
	ret, _, _ := syscall.Syscall6(
		(*comvt.ISequentialStream)(unsafe.Pointer(*me.Ptr())).Read, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&numBytesRead)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_ISequentialStream) Write(data []byte) (numBytesWritten uint32) {
	ret, _, _ := syscall.Syscall6(
		(*comvt.ISequentialStream)(unsafe.Pointer(*me.Ptr())).Write, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&numBytesWritten)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}
