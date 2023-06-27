//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
)

// Handle to an OLE [block of memory].
//
// [block of memory]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
type HTASKMEM HANDLE

// [CoTaskMemAlloc] function.
//
// ⚠️ You must defer HTASKMEM.CoTaskMemFree().
//
// [CoTaskMemAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(numBytes int) HTASKMEM {
	ret, _, _ := syscall.SyscallN(proc.CoTaskMemAlloc.Addr(),
		uintptr(numBytes))
	if ret == 0 {
		panic("CoTaskMemAlloc() failed.")
	}
	return HTASKMEM(ret)
}

// [CoTaskMemFree] function.
//
// [CoTaskMemFree]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func (hMem HTASKMEM) CoTaskMemFree() {
	syscall.SyscallN(proc.CoTaskMemFree.Addr(),
		uintptr(hMem))
}

// [CoTaskMemRealloc] function.
//
// ⚠️ You must defer CoTaskMemFree().
//
// [CoTaskMemRealloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func (hMem HTASKMEM) CoTaskMemRealloc(numBytes int) HTASKMEM {
	ret, _, _ := syscall.SyscallN(proc.CoTaskMemRealloc.Addr(),
		uintptr(hMem), uintptr(numBytes))
	if ret == 0 {
		panic("CoTaskMemRealloc() failed.")
	}
	return HTASKMEM(ret)
}
