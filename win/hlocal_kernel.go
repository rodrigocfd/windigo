//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a local [memory block].
//
// [memory block]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hlocal
type HLOCAL HANDLE

// [LocalAlloc] function.
//
// With co.LMEM_FIXED, the handle itself is the pointer to the memory block, and
// it can optionally be passed to unsafe.Slice() to create a slice over the
// memory block.
//
// With co.LMEM_MOVEABLE, you must call HLOCAL.LocalLock() to retrieve the
// pointer.
//
// ⚠️ You must defer HLOCAL.LocalFree().
//
// # Example:
//
//	hMem := win.LocalAlloc(co.LMEM_FIXED|co.LMEM_ZEROINIT, 50)
//	defer hMem.LocalFree()
//
//	sliceMem := hMem.LocalLock(50)
//	defer hMem.LocalUnlock()
//
// [LocalAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localalloc
func LocalAlloc(uFlags co.LMEM, numBytes int) HLOCAL {
	ret, _, err := syscall.SyscallN(proc.LocalAlloc.Addr(),
		uintptr(uFlags), uintptr(numBytes))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HLOCAL(ret)
}

// [LocalFlags] function.
//
// [LocalFlags]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localflags
func (hLocal HLOCAL) LocalFlags() co.LMEM {
	ret, _, err := syscall.SyscallN(proc.LocalFlags.Addr(),
		uintptr(hLocal))
	if ret == _LMEM_INVALID_HANDLE {
		panic(errco.ERROR(err))
	}
	return co.LMEM(ret)
}

// [LocalFree] function.
//
// This method is safe to be called if hLocal is zero.
//
// [LocalFree]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localfree
func (hLocal HLOCAL) LocalFree() error {
	ret, _, err := syscall.SyscallN(proc.LocalFree.Addr(),
		uintptr(hLocal))
	if ret != 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [LocalLock] function.
//
// If you called LocalAlloc() with co.LMEM_FIXED, technically you don't need to
// call this method, because the handle itself is the pointer to the memory
// block; however, this method is easier to use.
//
// Make sure that numBytes isn't greater than the memory block size, or you'll
// have a segfault. The safest way is simply call HLOCAL.LocalSize().
//
// ⚠️ You must defer HLOCAL.LocalUnlock(). After that, the slice must not be
// used.
//
// # Example:
//
//	hMem := win.LocalAlloc(co.LMEM_FIXED|co.LMEM_ZEROINIT, 50)
//	defer hMem.LocalFree()
//
//	sliceMem := hMem.LocalLock(hMem.LocalSize())
//	defer hMem.LocalUnlock()
//
// [LocalLock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-locallock
func (hLocal HLOCAL) LocalLock(numBytes int) []byte {
	ret, _, err := syscall.SyscallN(proc.LocalLock.Addr(),
		uintptr(hLocal))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), numBytes)
}

// [LocalReAlloc] function.
//
// ⚠️ You must defer HLOCAL.LocalFree().
//
// [LocalReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localrealloc
func (hLocal HLOCAL) LocalReAlloc(numBytes int, uFlags co.LMEM) HLOCAL {
	ret, _, err := syscall.SyscallN(proc.LocalReAlloc.Addr(),
		uintptr(hLocal), uintptr(numBytes), uintptr(uFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HLOCAL(ret)
}

// [LocalSize] function.
//
// [LocalSize]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localsize
func (hLocal HLOCAL) LocalSize() int {
	ret, _, err := syscall.SyscallN(proc.LocalSize.Addr(),
		uintptr(hLocal))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int(ret)
}

// [LocalUnlock] function.
//
// [LocalUnlock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localunlock
func (hLocal HLOCAL) LocalUnlock() error {
	ret, _, err := syscall.SyscallN(proc.LocalUnlock.Addr(),
		uintptr(hLocal))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}
