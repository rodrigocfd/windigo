//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a [heap] object.
//
// [heap]: https://learn.microsoft.com/en-us/windows/win32/Memory/heap-functions
type HHEAP HANDLE

// [GetProcessHeap] function.
//
// [GetProcessHeap]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-getprocessheap
func GetProcessHeap() (HHEAP, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessHeap))
	if ret == 0 {
		return HHEAP(0), co.ERROR(err)
	}
	return HHEAP(ret), nil
}

// [HeapCreate] function.
//
// ⚠️ You must defer [HHEAP.HeapDestroy].
//
// [HeapCreate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcreate
func HeapCreate(options co.HEAP_CREATE, initialSize, maximumSize uint) (HHEAP, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapCreate),
		uintptr(options),
		uintptr(initialSize),
		uintptr(maximumSize))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return HHEAP(ret), nil
}

// [HeapAlloc] function.
//
// ⚠️ You must defer [HHEAP.HeapFree].
//
// # Example
//
//	hHeap, _ := win.GetProcessHeap()
//	ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, 10)
//	defer hHeap.HeapFree(co.HEAP_NS_NONE, ptr)
//
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
func (hHeap HHEAP) HeapAlloc(flags co.HEAP_ALLOC, numBytes uint) (unsafe.Pointer, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapAlloc),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(numBytes))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

// [HeapCompact] function.
//
// [HeapCompact]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcompact
func (hHeap HHEAP) HeapCompact(flags co.HEAP_NS) (uint, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapCompact),
		uintptr(hHeap),
		uintptr(flags))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

// [HeapDestroy] function.
//
// Paired with [HeapCreate].
//
// [HeapDestroy]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapdestroy
func (hHeap HHEAP) HeapDestroy() error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapDestroy),
		uintptr(hHeap))
	return utl.ZeroAsGetLastError(ret, err)
}

// [HeapFree] function.
//
// Paired with [HHEAP.HeapAlloc] and [HHEAP.HeapReAlloc].
//
// This method is safe to be called if block is nil.
//
// # Example
//
//	hHeap, _ := win.GetProcessHeap()
//	ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, 10)
//	defer hHeap.HeapFree(co.HEAP_NS_NONE, ptr)
//
// [HeapFree]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapfree
func (hHeap HHEAP) HeapFree(flags co.HEAP_NS, ptr unsafe.Pointer) error {
	if ptr == nil {
		return nil // nothing to do
	}

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapFree),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr))
	return utl.ZeroAsGetLastError(ret, err)
}

// [HeapReAlloc] function.
//
// If block is nil, simple calls [HHEAP.HeapAlloc].
//
// ⚠️ You must defer [HHEAP.HeapFree].
//
// # Example
//
//	hHeap, _ := win.GetProcessHeap()
//	ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, 10)
//	defer hHeap.HeapFree(co.HEAP_NS_NONE, ptr)
//
//	ptr, _ = hHeap.HeapReAlloc(
//		co.HEAP_REALLOC_ZERO_MEMORY, ptr, 20)
//
// [HeapReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heaprealloc
func (hHeap HHEAP) HeapReAlloc(
	flags co.HEAP_REALLOC,
	ptr unsafe.Pointer,
	numBytes uint,
) (unsafe.Pointer, error) {
	if ptr == nil {
		return hHeap.HeapAlloc(co.HEAP_ALLOC(flags), numBytes) // simply forward
	}

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapReAlloc),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr),
		uintptr(numBytes))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

// [HeapSize] function.
//
// [HeapSize]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapsize
func (hHeap HHEAP) HeapSize(flags co.HEAP_NS, ptr unsafe.Pointer) (uint, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_HeapSize),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr))
	if int64(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

// [HeapValidate] function.
//
// [HeapValidate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapvalidate
func (hHeap HHEAP) HeapValidate(flags co.HEAP_NS, ptr unsafe.Pointer) bool {
	ret, _, _ := syscall.SyscallN(dll.Kernel(dll.PROC_HeapValidate),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr))
	return ret != 0
}
