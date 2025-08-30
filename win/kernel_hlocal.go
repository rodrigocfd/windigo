//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a local [memory block].
//
// [memory block]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hlocal
type HLOCAL HANDLE

// [LocalAlloc] function.
//
// With [co.LMEM_FIXED], the handle itself is the pointer to the memory block,
// and it can optionally be passed to unsafe.Slice() to create a slice over the
// memory block.
//
// With [co.LMEM_MOVEABLE], you must call [HLOCAL.LocalLock] to retrieve the
// pointer.
//
// ⚠️ You must defer [HLOCAL.LocalFree].
//
// Example:
//
//	hMem, _ := win.LocalAlloc(co.LMEM_FIXED|co.LMEM_ZEROINIT, 10)
//	defer hMem.LocalFree()
//
//	sliceMem, _ := hMem.LocalLockSlice()
//	defer hMem.LocalUnlock()
//
//	println(len(sliceMem))
//
// [LocalAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localalloc
func LocalAlloc(flags co.LMEM, numBytes uint) (HLOCAL, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalAlloc, "LocalAlloc"),
		uintptr(flags),
		uintptr(numBytes))
	if ret == 0 {
		return HLOCAL(0), co.ERROR(err)
	}
	return HLOCAL(ret), nil
}

var _LocalAlloc *syscall.Proc

// [LocalFlags] function.
//
// [LocalFlags]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localflags
func (hLocal HLOCAL) LocalFlags() (co.LMEM, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalFlags, "LocalFlags"),
		uintptr(hLocal))
	if ret == utl.LMEM_INVALID_HANDLE {
		return co.LMEM(0), co.ERROR(err)
	}
	return co.LMEM(ret), nil
}

var _LocalFlags *syscall.Proc

// [LocalFree] function.
//
// Paired with [LocalAlloc] and [HLOCAL.LocalReAlloc].
//
// This method is safe to be called if hLocal is zero.
//
// [LocalFree]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localfree
func (hLocal HLOCAL) LocalFree() error {
	if hLocal == 0 {
		return nil // nothing to do
	}

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalFree, "LocalFree"),
		uintptr(hLocal))
	if ret != 0 {
		return co.ERROR(err)
	}
	return nil
}

var _LocalFree *syscall.Proc

// [LocalLock] function.
//
// If you called [LocalAlloc] with [co.LMEM_FIXED], technically you don't need
// to call this method, because the handle itself is the pointer to the memory
// block; however, this method is safer.
//
// If you need to access the memory block as a slice, see
// [HLOCAL.LocalLockSlice].
//
// ⚠️ You must defer [HLOCAL.LocalUnlock].
//
// Example:
//
//	hMem, _ := win.LocalAlloc(co.LMEM_FIXED|co.LMEM_ZEROINIT, 10)
//	defer hMem.LocalFree()
//
//	szMem, _ := hMem.LocalSize()
//
//	ptrMem, _ := hMem.LocalLock()
//	defer hMem.LocalUnlock()
//
// [LocalLock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-locallock
func (hLocal HLOCAL) LocalLock() (unsafe.Pointer, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalLock, "LocalLock"),
		uintptr(hLocal))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

var _LocalLock *syscall.Proc

// Calls [HLOCAL.LocalSize] and [HLOCAL.LocalLock] to retrieve the size and
// pointer to the allocated memory, then converts it to a slice over it.
//
// ⚠️ You must defer [HLOCAL.LocalUnlock].
//
// Example:
//
//	hMem, _ := win.LocalAlloc(co.LMEM_FIXED|co.LMEM_ZEROINIT, 10)
//	defer hMem.LocalFree()
//
//	sliceMem, _ := hMem.LocalLockSlice()
//	defer hMem.LocalUnlock()
//
//	println(len(sliceMem))
func (hLocal HLOCAL) LocalLockSlice() ([]byte, error) {
	sz, wErr := hLocal.LocalSize()
	if wErr != nil {
		return nil, wErr
	}

	ptr, wErr := hLocal.LocalLock()
	if wErr != nil {
		return nil, wErr
	}

	return unsafe.Slice((*byte)(ptr), sz), nil
}

// [LocalReAlloc] function.
//
// Be careful when using this function. It returns a new [HLOCAL] handle, which
// invalidates the previous one – that is, you should not call
// [HLOCAL.LocalFree] on the previous one. This can become tricky if you
// used defer.
//
// ⚠️ You must defer [HLOCAL.LocalFree].
//
// [LocalReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localrealloc
func (hLocal HLOCAL) LocalReAlloc(numBytes uint, flags co.LMEM) (HLOCAL, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalReAlloc, "LocalReAlloc"),
		uintptr(hLocal),
		uintptr(numBytes),
		uintptr(flags))
	if ret == 0 {
		return HLOCAL(0), co.ERROR(err)
	}
	return HLOCAL(ret), nil
}

var _LocalReAlloc *syscall.Proc

// [LocalSize] function.
//
// [LocalSize]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localsize
func (hLocal HLOCAL) LocalSize() (uint, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalSize, "LocalSize"),
		uintptr(hLocal))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _LocalSize *syscall.Proc

// [LocalUnlock] function.
//
// [LocalUnlock]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localunlock
func (hLocal HLOCAL) LocalUnlock() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LocalUnlock, "LocalUnlock"),
		uintptr(hLocal))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return wErr
	}
	return nil
}

var _LocalUnlock *syscall.Proc
