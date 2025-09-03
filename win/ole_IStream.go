//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IStream] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-istream
type IStream struct{ ISequentialStream }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IStream) IID() co.IID {
	return co.IID_IStream
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-clone
func (me *IStream) Clone(releaser *OleReleaser) (*IStream, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IStream{ISequentialStream{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-commit
func (me *IStream) Commit(flags co.STGC) error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Commit,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(flags))
	return utl.ErrorAsHResult(ret)
}

// [CopyTo] method.
//
// Panics if numBytes is negative.
//
// [CopyTo]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-copyto
func (me *IStream) CopyTo(
	dest *IStream,
	numBytes int,
) (numBytesRead, numBytesWritten int, hr error) {
	utl.PanicNeg(numBytes)
	var read64, written64 uint64

	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).CopyTo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(dest.Ppvt())),
		uintptr(uint64(numBytes)),
		uintptr(unsafe.Pointer(&read64)),
		uintptr(unsafe.Pointer(&written64)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(read64), int(written64), nil
	} else {
		return 0, 0, hr
	}
}

// [LockRegion] method.
//
// Panics if offset or length is negative.
//
// ⚠️ You must defer [IStream.UnlockRegion].
//
// [LockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-lockregion
func (me *IStream) LockRegion(offset, length int, lockType co.LOCKTYPE) error {
	utl.PanicNeg(offset, length)
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).LockRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint64(offset)),
		uintptr(uint64(length)),
		uintptr(lockType))
	return utl.ErrorAsHResult(ret)
}

// [Revert] method.
//
// [Revert]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-revert
func (me *IStream) Revert() error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Revert,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [Seek] method.
//
// [Seek]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-seek
func (me *IStream) Seek(displacement int, origin co.STREAM_SEEK) (newOffset int, hr error) {
	var newOff64 uint64
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Seek,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int64(displacement)),
		uintptr(origin),
		uintptr(unsafe.Pointer(&newOff64)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(newOff64), nil
	} else {
		return 0, hr
	}
}

// [SetSize] method.
//
// Panics if newSize is negative.
//
// [SetSize]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-setsize
func (me *IStream) SetSize(newSize int) error {
	utl.PanicNeg(newSize)
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).SetSize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint64(newSize)))
	return utl.ErrorAsHResult(ret)
}

// [Stat] method.
//
// [Stat]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-stat
func (me *IStream) Stat(flag co.STATFLAG) (STATSTG, error) {
	var stg STATSTG
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Stat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&stg)),
		uintptr(flag))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stg, nil
	} else {
		return STATSTG{}, hr
	}
}

// [UnlockRegion] method.
//
// Paired with [IStream.LockRegion].
//
// Panics if offset or length is negative.
//
// [UnlockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-unlockregion
func (me *IStream) UnlockRegion(offset, length int, lockType co.LOCKTYPE) error {
	utl.PanicNeg(offset, length)
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).UnlockRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint64(offset)),
		uintptr(uint64(length)),
		uintptr(lockType))
	return utl.ErrorAsHResult(ret)
}

type _IStreamVt struct {
	_ISequentialStreamVt
	Seek         uintptr
	SetSize      uintptr
	CopyTo       uintptr
	Commit       uintptr
	Revert       uintptr
	LockRegion   uintptr
	UnlockRegion uintptr
	Stat         uintptr
	Clone        uintptr
}
