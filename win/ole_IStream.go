//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
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
// [CopyTo]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-copyto
func (me *IStream) CopyTo(
	dest *IStream,
	numBytes uint64,
) (numBytesRead, numBytesWritten uint64, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).CopyTo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(dest.Ppvt())),
		uintptr(numBytes),
		uintptr(unsafe.Pointer(&numBytesRead)),
		uintptr(unsafe.Pointer(&numBytesWritten)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		numBytes, numBytesWritten = 0, 0
	}
	return
}

// [LockRegion] method.
//
// ⚠️ You must defer [IStream.UnlockRegion].
//
// [LockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-lockregion
func (me *IStream) LockRegion(offset, length uint64, lockType co.LOCKTYPE) error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).LockRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(offset),
		uintptr(length),
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
func (me *IStream) Seek(displacement int64, origin co.STREAM_SEEK) (newOffset uint, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Seek,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(displacement),
		uintptr(origin),
		uintptr(unsafe.Pointer(&newOffset)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		newOffset = 0
	}
	return
}

// [SetSize] method.
//
// [SetSize]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-setsize
func (me *IStream) SetSize(newSize uint) error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).SetSize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(newSize))
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
// [UnlockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-unlockregion
func (me *IStream) UnlockRegion(offset, length uint64, lockType co.LOCKTYPE) error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).UnlockRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(offset),
		uintptr(length),
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
