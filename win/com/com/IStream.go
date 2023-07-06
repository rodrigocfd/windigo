//go:build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IStream] COM interface.
//
// [IStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-istream
type IStream interface {
	ISequentialStream

	// [Commit] COM method.
	//
	// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-commit
	Commit(flags comco.STGC)

	// [CopyTo] COM method.
	//
	// [CopyTo]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-copyto
	CopyTo(dest IStream, numBytes uint64) (numBytesRead, numBytesWritten uint64)

	// [LockRegion] COM method.
	//
	// ⚠️ You must defer IStream.UnlockRegion().
	//
	// [LockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-lockregion
	LockRegion(offset, length uint64, lockType comco.LOCKTYPE)

	// [Revert] COM method.
	//
	// [Revert]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-revert
	Revert()

	// [Seek] COM method.
	//
	// [Seek]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-seek
	Seek(displacement int64, origin comco.STREAM_SEEK) (newOffset uint64)

	// [SetSize] COM method.
	//
	// [SetSize]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-setsize
	SetSize(newSize uint64)

	// [UnlockRegion] COM method.
	//
	// [UnlockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-unlockregion
	UnlockRegion(offset, length uint64, lockType comco.LOCKTYPE)
}

type _IStream struct{ ISequentialStream }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IStream.Release().
func NewIStream(base IUnknown) IStream {
	return &_IStream{ISequentialStream: NewISequentialStream(base)}
}

// [SHCreateMemStream] function.
//
// Creates an IStream projection over a slice, which must remain valid in
// memory throughout IStream's lifetime.
//
// ⚠️ You must defer IStream.Release().
//
// # Example
//
//	data := []byte{0x10, 0x11, 0x12}
//	defer runtime.KeepAlive(data)
//
//	stream := SHCreateMemStream(data)
//	defer stream.Release()
//
// [SHCreateMemStream]: https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func SHCreateMemStream(src []byte) IStream {
	ret, _, _ := syscall.SyscallN(proc.SHCreateMemStream.Addr(),
		uintptr(unsafe.Pointer(&src[0])), uintptr(len(src)))
	if ret == 0 {
		panic(errco.E_OUTOFMEMORY)
	}
	return NewIStream(NewIUnknown((**comvt.IUnknown)(unsafe.Pointer(ret))))
}

func (me *_IStream) Commit(flags comco.STGC) {
	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).Commit,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) CopyTo(
	dest IStream, numBytes uint64) (numBytesRead, numBytesWritten uint64) {

	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).CopyTo,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(dest.Ptr())),
		uintptr(numBytes),
		uintptr(unsafe.Pointer(&numBytesRead)),
		uintptr(unsafe.Pointer(&numBytesWritten)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IStream) LockRegion(offset, length uint64, lockType comco.LOCKTYPE) {
	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).LockRegion,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(offset), uintptr(length), uintptr(lockType))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) Revert() {
	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).Revert,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) Seek(
	displacement int64, origin comco.STREAM_SEEK) (newOffset uint64) {

	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).Seek,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(displacement), uintptr(origin),
		uintptr(unsafe.Pointer(&newOffset)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IStream) SetSize(newSize uint64) {
	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).SetSize,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(newSize))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) UnlockRegion(
	offset, length uint64, lockType comco.LOCKTYPE) {

	ret, _, _ := syscall.SyscallN(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).UnlockRegion,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(offset), uintptr(length), uintptr(lockType))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
