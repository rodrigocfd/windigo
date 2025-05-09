//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
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
	ret, _, err := syscall.SyscallN(_GetProcessHeap.Addr())
	if ret == 0 {
		return HHEAP(0), co.ERROR(err)
	}
	return HHEAP(ret), nil
}

var _GetProcessHeap = dll.Kernel32.NewProc("GetProcessHeap")

// [HeapCreate] function.
//
// ⚠️ You must defer HHEAP.HeapDestroy().
//
// [HeapCreate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcreate
func HeapCreate(options co.HEAP_CREATE, initialSize, maximumSize uint) (HHEAP, error) {
	ret, _, err := syscall.SyscallN(_HeapCreate.Addr(),
		uintptr(options), uintptr(initialSize), uintptr(maximumSize))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return HHEAP(ret), nil
}

var _HeapCreate = dll.Kernel32.NewProc("HeapCreate")

// [HeapAlloc] function.
//
// ⚠️ You must defer HHEAP.HeapFree().
//
// # Example
//
//	hHeap, _ := win.GetProcessHeap()
//	ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, 10)
//	defer hHeap.HeapFree(co.HEAP_NS_NONE, ptr)
//
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
func (hHeap HHEAP) HeapAlloc(flags co.HEAP_ALLOC, numBytes uint) (unsafe.Pointer, error) {
	ret, _, err := syscall.SyscallN(_HeapAlloc.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(numBytes))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

var _HeapAlloc = dll.Kernel32.NewProc("HeapAlloc")

// [HeapCompact] function.
//
// [HeapCompact]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcompact
func (hHeap HHEAP) HeapCompact(flags co.HEAP_NS) (uint, error) {
	ret, _, err := syscall.SyscallN(_HeapCompact.Addr(),
		uintptr(hHeap), uintptr(flags))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _HeapCompact = dll.Kernel32.NewProc("HeapCompact")

// [HeapDestroy] function.
//
// Paired with [HeapCreate].
//
// [HeapDestroy]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapdestroy
// [HeapCreate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcreate
func (hHeap HHEAP) HeapDestroy() error {
	ret, _, err := syscall.SyscallN(_HeapDestroy.Addr(),
		uintptr(hHeap))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _HeapDestroy = dll.Kernel32.NewProc("HeapDestroy")

// [HeapFree] function.
//
// Paired with [HeapAlloc] and [HeapReAlloc].
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
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
// [HeapReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heaprealloc
func (hHeap HHEAP) HeapFree(flags co.HEAP_NS, ptr unsafe.Pointer) error {
	if ptr == nil {
		return nil // nothing to do
	}

	ret, _, err := syscall.SyscallN(_HeapFree.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(ptr))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _HeapFree = dll.Kernel32.NewProc("HeapFree")

// [HeapReAlloc] function.
//
// If block is nil, simple calls [HeapAlloc].
//
// ⚠️ You must defer HHEAP.HeapFree().
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
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
func (hHeap HHEAP) HeapReAlloc(
	flags co.HEAP_REALLOC,
	ptr unsafe.Pointer,
	numBytes uint,
) (unsafe.Pointer, error) {
	if ptr == nil {
		return hHeap.HeapAlloc(co.HEAP_ALLOC(flags), numBytes) // simply forward
	}

	ret, _, err := syscall.SyscallN(_HeapReAlloc.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(ptr),
		uintptr(numBytes))
	if ret == 0 {
		return nil, co.ERROR(err)
	}
	return unsafe.Pointer(ret), nil
}

var _HeapReAlloc = dll.Kernel32.NewProc("HeapReAlloc")

// [HeapSize] function.
//
// [HeapSize]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapsize
func (hHeap HHEAP) HeapSize(flags co.HEAP_NS, ptr unsafe.Pointer) (uint, error) {
	ret, _, err := syscall.SyscallN(_HeapSize.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(ptr))
	if int64(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _HeapSize = dll.Kernel32.NewProc("HeapSize")

// [HeapValidate] function.
//
// [HeapValidate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapvalidate
func (hHeap HHEAP) HeapValidate(flags co.HEAP_NS, ptr unsafe.Pointer) bool {
	ret, _, _ := syscall.SyscallN(_HeapValidate.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(ptr))
	return ret != 0
}

var _HeapValidate = dll.Kernel32.NewProc("HeapValidate")
