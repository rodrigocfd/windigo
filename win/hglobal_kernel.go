//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a global memory block.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hglobal
type HGLOBAL HANDLE

// With co.GMEM_FIXED, the handle itself is the pointer to the memory block, and
// it can optionally be passed to unsafe.Slice() to create a slice over the
// memory block.
//
// With co.GMEM_MOVEABLE, you must call HGLOBAL.GlobalLock() to retrieve the
// pointer.
//
// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// Example:
//
//		hMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 50)
//		defer hMem.GlobalFree()
//
//		sliceMem := hMem.GlobalLock(50)
//		defer hMem.GlobalUnlock()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(uFlags co.GMEM, numBytes int) HGLOBAL {
	ret, _, err := syscall.Syscall(proc.GlobalAlloc.Addr(), 2,
		uintptr(uFlags), uintptr(numBytes), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// Allocs a null-terminated *uint16.
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
// Example:
//
//		hMem := win.GlobalAllocStr(co.GMEM_FIXED, "my text")
//		defer hMem.GlobalFree()
//
//		charSlice := hMem.GlobalLock(hMem.GlobalSize())
//		defer hMem.GlobalUnlock()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAllocStr(uFlags co.GMEM, s string) HGLOBAL {
	sliceStr16 := Str.ToNativeSlice(s) // null-terminated
	sliceStr8 := unsafe.Slice((*byte)(unsafe.Pointer(&sliceStr16[0])), len(sliceStr16)*2)

	hMem := GlobalAlloc(uFlags, len(sliceStr8))
	if (uFlags & co.GMEM_MOVEABLE) != 0 {
		dest := hMem.GlobalLock(len(sliceStr8))
		copy(dest, sliceStr8)
		hMem.GlobalUnlock()
	} else {
		dest := unsafe.Slice((*byte)(unsafe.Pointer(hMem)), len(sliceStr8))
		copy(dest, sliceStr8)
	}
	return hMem
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalflags
func (hGlobal HGLOBAL) GlobalFlags() co.GMEM {
	ret, _, err := syscall.Syscall(proc.GlobalFlags.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == _GMEM_INVALID_HANDLE {
		panic(errco.ERROR(err))
	}
	return co.GMEM(ret)
}

// This method is safe to be called if hGlobal is zero.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
func (hGlobal HGLOBAL) GlobalFree() {
	ret, _, err := syscall.Syscall(proc.GlobalFree.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret != 0 {
		panic(errco.ERROR(err))
	}
}

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
// Example:
//
//		hMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 50)
//		defer hMem.GlobalFree()
//
//		sliceMem := hMem.GlobalLock(hMem.GlobalSize())
//		defer hMem.GlobalUnlock()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
func (hGlobal HGLOBAL) GlobalLock(numBytes int) []byte {
	ret, _, err := syscall.Syscall(proc.GlobalLock.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), numBytes)
}

// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalrealloc
func (hGlobal HGLOBAL) GlobalReAlloc(numBytes int, uFlags co.GMEM) HGLOBAL {
	ret, _, err := syscall.Syscall(proc.GlobalReAlloc.Addr(), 3,
		uintptr(hGlobal), uintptr(numBytes), uintptr(uFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalsize
func (hGlobal HGLOBAL) GlobalSize() int {
	ret, _, err := syscall.Syscall(proc.GlobalSize.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
func (hGlobal HGLOBAL) GlobalUnlock() {
	ret, _, err := syscall.Syscall(proc.GlobalUnlock.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
}
