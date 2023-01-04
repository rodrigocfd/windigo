//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
)

// Handle to an OLE block of memory.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
type HTASKMEM HANDLE

// ⚠️ You must defer HTASKMEM.CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(numBytes int) HTASKMEM {
	ret, _, _ := syscall.Syscall(proc.CoTaskMemAlloc.Addr(), 1,
		uintptr(numBytes), 0, 0)
	if ret == 0 {
		panic("CoTaskMemAlloc() failed.")
	}
	return HTASKMEM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func (hMem HTASKMEM) CoTaskMemFree() {
	syscall.Syscall(proc.CoTaskMemFree.Addr(), 1,
		uintptr(hMem), 0, 0)
}

// ⚠️ You must defer CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func (hMem HTASKMEM) CoTaskMemRealloc(numBytes int) HTASKMEM {
	ret, _, _ := syscall.Syscall(proc.CoTaskMemRealloc.Addr(), 2,
		uintptr(hMem), uintptr(numBytes), 0)
	if ret == 0 {
		panic("CoTaskMemRealloc() failed.")
	}
	return HTASKMEM(ret)
}
