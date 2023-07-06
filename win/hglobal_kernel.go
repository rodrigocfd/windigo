//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a global [memory block].
//
// [memory block]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hglobal
type HGLOBAL HANDLE

// [GlobalAlloc] function.
//
// With co.GMEM_FIXED, the handle itself is the pointer to the memory block, and
// it can optionally be passed to unsafe.Slice() to create a slice over the
// memory block.
//
// With co.GMEM_MOVEABLE, you must call HGLOBAL.GlobalLock() to retrieve the
// pointer.
//
// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// # Example
//
//	hMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 50)
//	defer hMem.GlobalFree()
//
//	sliceMem := hMem.GlobalLock(50)
//	defer hMem.GlobalUnlock()
//
// [GlobalAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(uFlags co.GMEM, numBytes int) HGLOBAL {
	ret, _, err := syscall.SyscallN(proc.GlobalAlloc.Addr(),
		uintptr(uFlags), uintptr(numBytes))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// [GlobalFlags] function.
//
// [GlobalFlags]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalflags
func (hGlobal HGLOBAL) GlobalFlags() co.GMEM {
	ret, _, err := syscall.SyscallN(proc.GlobalFlags.Addr(),
		uintptr(hGlobal))
	if ret == _GMEM_INVALID_HANDLE {
		panic(errco.ERROR(err))
	}
	return co.GMEM(ret)
}

// [GlobalFree] function.
//
// This method is safe to be called if hGlobal is zero.
//
// [GlobalFree]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
func (hGlobal HGLOBAL) GlobalFree() error {
	ret, _, err := syscall.SyscallN(proc.GlobalFree.Addr(),
		uintptr(hGlobal))
	if ret != 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [GlobalLock] function.
//
// If you called GlobalAlloc() with co.GMEM_FIXED, technically you don't need to
// call this method, because the handle itself is the pointer to the memory
// block; however, this method is easier to use.
//
// Make sure that numBytes isn't greater than the memory block size, or you'll
// have a segfault. The safest way is simply call HGLOBAL.GlobalSize().
//
// ⚠️ You must defer HGLOBAL.GlobalUnlock(). After that, the slice must not be
// used.
//
// # Example
//
//	hMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 50)
//	defer hMem.GlobalFree()
//
//	sliceMem := hMem.GlobalLock(hMem.GlobalSize())
//	defer hMem.GlobalUnlock()
//
// [GlobalLock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
func (hGlobal HGLOBAL) GlobalLock(numBytes int) []byte {
	ret, _, err := syscall.SyscallN(proc.GlobalLock.Addr(),
		uintptr(hGlobal))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), numBytes)
}

// [GlobalReAlloc] function.
//
// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// [GlobalReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalrealloc
func (hGlobal HGLOBAL) GlobalReAlloc(numBytes int, uFlags co.GMEM) HGLOBAL {
	ret, _, err := syscall.SyscallN(proc.GlobalReAlloc.Addr(),
		uintptr(hGlobal), uintptr(numBytes), uintptr(uFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// [GlobalSize] function.
//
// [GlobalSize]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalsize
func (hGlobal HGLOBAL) GlobalSize() int {
	ret, _, err := syscall.SyscallN(proc.GlobalSize.Addr(),
		uintptr(hGlobal))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int(ret)
}

// [GlobalUnlock] function.
//
// [GlobalUnlock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
func (hGlobal HGLOBAL) GlobalUnlock() error {
	ret, _, err := syscall.SyscallN(proc.GlobalUnlock.Addr(),
		uintptr(hGlobal))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}
