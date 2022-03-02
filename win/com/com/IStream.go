package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-istream
type IStream interface {
	ISequentialStream

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-commit
	Commit(flags comco.STGC)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-copyto
	CopyTo(dest IStream, numBytes uint64) (numBytesRead, numBytesWritten uint64)

	// ⚠️ You must defer IStream.UnlockRegion().
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-lockregion
	LockRegion(offset, length uint64, lockType comco.LOCKTYPE)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-revert
	Revert()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-seek
	Seek(displacement int64, origin comco.STREAM_SEEK) (newOffset uint64)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-setsize
	SetSize(newSize uint64)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-unlockregion
	UnlockRegion(offset, length uint64, lockType comco.LOCKTYPE)
}

type _IStream struct{ ISequentialStream }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IStream.Release().
func NewIStream(base IUnknown) IStream {
	return &_IStream{ISequentialStream: NewISequentialStream(base)}
}

// Calls SHCreateMemStream() to create a new stream over a slice, which must
// remain in memory.
//
// ⚠️ You must defer IStream.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func NewIStreamFromSlice(src []byte) IStream {
	ret, _, _ := syscall.Syscall(proc.SHCreateMemStream.Addr(), 2,
		uintptr(unsafe.Pointer(&src[0])), uintptr(len(src)), 0)
	if ret == 0 {
		panic(errco.E_OUTOFMEMORY)
	}
	return NewIStream(NewIUnknown((**comvt.IUnknown)(unsafe.Pointer(ret))))
}

func (me *_IStream) Commit(flags comco.STGC) {
	ret, _, _ := syscall.Syscall(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).Commit, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(flags), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) CopyTo(
	dest IStream, numBytes uint64) (numBytesRead, numBytesWritten uint64) {

	ret, _, _ := syscall.Syscall6(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).CopyTo, 5,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(dest.Ptr())),
		uintptr(numBytes),
		uintptr(unsafe.Pointer(&numBytesRead)),
		uintptr(unsafe.Pointer(&numBytesWritten)),
		0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IStream) LockRegion(offset, length uint64, lockType comco.LOCKTYPE) {
	ret, _, _ := syscall.Syscall6(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).LockRegion, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(offset), uintptr(length), uintptr(lockType),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) Revert() {
	ret, _, _ := syscall.Syscall(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).Revert, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) Seek(
	displacement int64, origin comco.STREAM_SEEK) (newOffset uint64) {

	ret, _, _ := syscall.Syscall6(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).Seek, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(displacement), uintptr(origin),
		uintptr(unsafe.Pointer(&newOffset)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IStream) SetSize(newSize uint64) {
	ret, _, _ := syscall.Syscall(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).SetSize, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(newSize), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IStream) UnlockRegion(
	offset, length uint64, lockType comco.LOCKTYPE) {

	ret, _, _ := syscall.Syscall6(
		(*comvt.IStream)(unsafe.Pointer(*me.Ptr())).UnlockRegion, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(offset), uintptr(length), uintptr(lockType),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
