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

// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(uFlags co.GMEM, dwBytes uint64) HGLOBAL {
	ret, _, err := syscall.Syscall(proc.GlobalAlloc.Addr(), 2,
		uintptr(uFlags), uintptr(dwBytes), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
func (hGlobal HGLOBAL) GlobalFree() {
	ret, _, err := syscall.Syscall(proc.GlobalFree.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret != 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HGLOBAL.GlobalUnlock(). After that, the slice must not be used.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
func (hGlobal HGLOBAL) GlobalLock() []byte {
	ret, _, err := syscall.Syscall(proc.GlobalLock.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), hGlobal.GlobalSize())
}

// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalrealloc
func (hGlobal HGLOBAL) GlobalReAlloc(dwBytes uint64, uFlags co.GMEM) HGLOBAL {
	ret, _, err := syscall.Syscall(proc.GlobalReAlloc.Addr(), 3,
		uintptr(hGlobal), uintptr(dwBytes), uintptr(uFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalsize
func (hGlobal HGLOBAL) GlobalSize() uint64 {
	ret, _, err := syscall.Syscall(proc.GlobalSize.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint64(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
func (hGlobal HGLOBAL) GlobalUnlock() {
	ret, _, err := syscall.Syscall(proc.GlobalUnlock.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
}
