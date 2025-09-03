//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a [heap] object.
//
// [heap]: https://learn.microsoft.com/en-us/windows/win32/Memory/heap-functions
type HHEAP HANDLE

// [GetProcessHeap] function.
//
// [GetProcessHeap]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-getprocessheap
func GetProcessHeap() (HHEAP, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessHeap, "GetProcessHeap"))
	if ret == 0 {
		return HHEAP(0), co.ERROR(err)
	}
	return HHEAP(ret), nil
}

var _GetProcessHeap *syscall.Proc

// [HeapCreate] function.
//
// Panics if initialSize or maximumSize is negative.
//
// ⚠️ You must defer [HHEAP.HeapDestroy].
//
// [HeapCreate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcreate
func HeapCreate(options co.HEAP_CREATE, initialSize, maximumSize int) (HHEAP, error) {
	utl.PanicNeg(initialSize, maximumSize)
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapCreate, "HeapCreate"),
		uintptr(options),
		uintptr(uint64(initialSize)),
		uintptr(uint64(maximumSize)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return HHEAP(ret), nil
}

var _HeapCreate *syscall.Proc

// [HeapAlloc] function.
//
// Panics if numBytes is negative.
//
// ⚠️ You must defer [HHEAP.HeapFree].
//
// Example:
//
//	hHeap, _ := win.GetProcessHeap()
//	ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, 10)
//	defer hHeap.HeapFree(co.HEAP_NS_NONE, ptr)
//
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
func (hHeap HHEAP) HeapAlloc(flags co.HEAP_ALLOC, numBytes int) (unsafe.Pointer, error) {
	utl.PanicNeg(numBytes)
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapAlloc, "HeapAlloc"),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(uint64(numBytes)))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

var _HeapAlloc *syscall.Proc

// [HeapCompact] function.
//
// [HeapCompact]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcompact
func (hHeap HHEAP) HeapCompact(flags co.HEAP_NS) (int, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapCompact, "HeapCompact"),
		uintptr(hHeap),
		uintptr(flags))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(uint64(ret)), nil
}

var _HeapCompact *syscall.Proc

// [HeapDestroy] function.
//
// Paired with [HeapCreate].
//
// [HeapDestroy]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapdestroy
func (hHeap HHEAP) HeapDestroy() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapDestroy, "HeapDestroy"),
		uintptr(hHeap))
	return utl.ZeroAsGetLastError(ret, err)
}

var _HeapDestroy *syscall.Proc

// [HeapFree] function.
//
// Paired with [HHEAP.HeapAlloc] and [HHEAP.HeapReAlloc].
//
// This method is safe to be called if block is nil.
//
// Example:
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

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapFree, "HeapFree"),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr))
	return utl.ZeroAsGetLastError(ret, err)
}

var _HeapFree *syscall.Proc

// [HeapReAlloc] function.
//
// If block is nil, simple calls [HHEAP.HeapAlloc].
//
// Panics if numBytes is negative.
//
// ⚠️ You must defer [HHEAP.HeapFree].
//
// Example:
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
	numBytes int,
) (unsafe.Pointer, error) {
	if ptr == nil {
		return hHeap.HeapAlloc(co.HEAP_ALLOC(flags), numBytes) // simply forward
	}

	utl.PanicNeg(numBytes)
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapReAlloc, "HeapReAlloc"),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr),
		uintptr(uint64(numBytes)))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

var _HeapReAlloc *syscall.Proc

// [HeapSize] function.
//
// [HeapSize]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapsize
func (hHeap HHEAP) HeapSize(flags co.HEAP_NS, ptr unsafe.Pointer) (int, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapSize, "HeapSize"),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr))
	if int64(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return int(uint64(ret)), nil
}

var _HeapSize *syscall.Proc

// [HeapValidate] function.
//
// [HeapValidate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapvalidate
func (hHeap HHEAP) HeapValidate(flags co.HEAP_NS, ptr unsafe.Pointer) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_HeapValidate, "HeapValidate"),
		uintptr(hHeap),
		uintptr(flags),
		uintptr(ptr))
	return ret != 0
}

var _HeapValidate *syscall.Proc
