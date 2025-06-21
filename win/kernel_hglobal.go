//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a global [memory block].
//
// [memory block]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hglobal
type HGLOBAL HANDLE

// [GlobalAlloc] function.
//
// With [co.GMEM_FIXED], the handle itself is the pointer to the memory block,
// and it can optionally be passed to unsafe.Slice() to create a slice over the
// memory block.
//
// With [co.GMEM_MOVEABLE], you must call [HGLOBAL.GlobalLock] to retrieve the
// pointer.
//
// ⚠️ You must defer [HGLOBAL.GlobalFree].
//
// # Example
//
//	hMem, _ := GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 10)
//	defer hMem.GlobalFree()
//
//	sliceMem, _ := hMem.GlobalLockSlice()
//	defer hMem.GlobalUnlock()
//
//	println(len(sliceMem))
//
// [GlobalAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(uFlags co.GMEM, numBytes uint) (HGLOBAL, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalAlloc, "GlobalAlloc"),
		uintptr(uFlags),
		uintptr(numBytes))
	if ret == 0 {
		return HGLOBAL(0), co.ERROR(err)
	}
	return HGLOBAL(ret), nil
}

var _GlobalAlloc *syscall.Proc

// [GlobalFlags] function.
//
// [GlobalFlags]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalflags
func (hGlobal HGLOBAL) GlobalFlags() (co.GMEM, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalFlags, "GlobalFlags"),
		uintptr(hGlobal))
	if ret == utl.GMEM_INVALID_HANDLE {
		return co.GMEM(0), co.ERROR(err)
	}
	return co.GMEM(ret), nil
}

var _GlobalFlags *syscall.Proc

// [GlobalFree] function.
//
// Paired with [GlobalAlloc] and [HGLOBAL.GlobalReAlloc].
//
// This method is safe to be called if hGlobal is zero.
//
// [GlobalFree]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
func (hGlobal HGLOBAL) GlobalFree() error {
	if hGlobal == 0 {
		return nil // nothing to do
	}

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalFree, "GlobalFree"),
		uintptr(hGlobal))
	if ret != 0 {
		return co.ERROR(err)
	}
	return nil
}

var _GlobalFree *syscall.Proc

// [GlobalLock] function.
//
// If you called [GlobalAlloc] with [co.GMEM_FIXED], technically you don't need
// to call this method, because the handle itself is the pointer to the memory
// block; however, this method is safer.
//
// If you need to access the memory block as a slice, see
// [HGLOBAL.GlobalLockSlice].
//
// ⚠️ You must defer [HGLOBAL.GlobalUnlock].
//
// # Example
//
//	hMem, _ := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 10)
//	defer hMem.GlobalFree()
//
//	szMem, _ := hMem.GlobalSize()
//
//	ptrMem, _ := hMem.GlobalLock()
//	defer hMem.GlobalUnlock()
//
// [GlobalLock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
func (hGlobal HGLOBAL) GlobalLock() (unsafe.Pointer, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalLock, "GlobalLock"),
		uintptr(hGlobal))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

var _GlobalLock *syscall.Proc

// Calls [HGLOBAL.GlobalSize] and [HGLOBAL.GlobalLock] to retrieve the size and
// pointer to the allocated memory, then converts it to a slice over it.
//
// ⚠️ You must defer [HGLOBAL.GlobalUnlock].
//
// # Example
//
//	hMem, _ := GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 10)
//	defer hMem.GlobalFree()
//
//	sliceMem, _ := hMem.GlobalLockSlice()
//	defer hMem.GlobalUnlock()
//
//	println(len(sliceMem))
func (hGlobal HGLOBAL) GlobalLockSlice() ([]byte, error) {
	sz, wErr := hGlobal.GlobalSize()
	if wErr != nil {
		return nil, wErr
	}

	ptr, wErr := hGlobal.GlobalLock()
	if wErr != nil {
		return nil, wErr
	}

	return unsafe.Slice((*byte)(ptr), sz), nil
}

// [GlobalReAlloc] function.
//
// ⚠️ You must defer [HGLOBAL.GlobalFree].
//
// [GlobalReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalrealloc
func (hGlobal HGLOBAL) GlobalReAlloc(numBytes uint, uFlags co.GMEM) (HGLOBAL, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalReAlloc, "GlobalReAlloc"),
		uintptr(hGlobal),
		uintptr(numBytes),
		uintptr(uFlags))
	if ret == 0 {
		return HGLOBAL(0), co.ERROR(err)
	}
	return HGLOBAL(ret), nil
}

var _GlobalReAlloc *syscall.Proc

// [GlobalSize] function.
//
// [GlobalSize]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalsize
func (hGlobal HGLOBAL) GlobalSize() (uint, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalSize, "GlobalSize"),
		uintptr(hGlobal))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _GlobalSize *syscall.Proc

// [GlobalUnlock] function.
//
// [GlobalUnlock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
func (hGlobal HGLOBAL) GlobalUnlock() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GlobalUnlock, "GlobalUnlock"),
		uintptr(hGlobal))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return wErr
	}
	return nil
}

var _GlobalUnlock *syscall.Proc
