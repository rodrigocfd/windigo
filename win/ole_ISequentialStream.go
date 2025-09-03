//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
)

// [ISequentialStream] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ISequentialStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-isequentialstream
type ISequentialStream struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ISequentialStream) IID() co.IID {
	return co.IID_ISequentialStream
}

// [Read] method.
//
// If returned numBytesRead is lower than requested buffer size, it means
// the end of stream was reached.
//
// [Read]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-read
func (me *ISequentialStream) Read(buffer []byte) (numBytesRead int, hr error) {
	var read32 uint32
	ret, _, _ := syscall.SyscallN(
		(*_ISequentialStreamVt)(unsafe.Pointer(*me.Ppvt())).Read,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(uint32(len(buffer))),
		uintptr(unsafe.Pointer(&read32)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(read32), nil
	} else {
		return 0, hr
	}
}

// [Write] method.
//
// [Write]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-write
func (me *ISequentialStream) Write(data []byte) (numBytesWritten int, hr error) {
	var written32 uint32
	ret, _, _ := syscall.SyscallN(
		(*_ISequentialStreamVt)(unsafe.Pointer(*me.Ppvt())).Write,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(uint32(len(data))),
		uintptr(unsafe.Pointer(&written32)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(written32), nil
	} else {
		return 0, hr
	}
	return
}

type _ISequentialStreamVt struct {
	_IUnknownVt
	Read  uintptr
	Write uintptr
}
