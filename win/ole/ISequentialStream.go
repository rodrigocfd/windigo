//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win/co"
)

// [ISequentialStream] COM interface.
//
// [ISequentialStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-isequentialstream
type ISequentialStream struct{ IUnknown }

// Returns the unique COM interface identifier.
func (*ISequentialStream) IID() co.IID {
	return co.IID_ISequentialStream
}

// [Read] method.
//
// If returned numBytesRead is lower than requested buffer size, it means
// the end of stream was reached.
//
// [Read]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-read
func (me *ISequentialStream) Read(buffer []byte) (numBytesRead uint32, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*vt.ISequentialStream)(unsafe.Pointer(*me.Ppvt())).Read,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&numBytesRead)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		numBytesRead = 0
	}
	return
}

// [Write] method.
//
// [Write]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-write
func (me *ISequentialStream) Write(data []byte) (numBytesWritten uint32, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*vt.ISequentialStream)(unsafe.Pointer(*me.Ppvt())).Write,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&numBytesWritten)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		numBytesWritten = 0
	}
	return
}
