//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a heap object.
type HHEAP HANDLE

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-getprocessheap
func GetProcessHeap() (HHEAP, error) {
	ret, _, err := syscall.SyscallN(proc.GetProcessHeap.Addr())
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return HHEAP(ret), nil
}

// ⚠️ You must defer HHEAP.HeapDestroy().
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcreate
func HeapCreate(
	options co.HEAP_CREATE, initialSize, maximumSize uint) (HHEAP, error) {

	ret, _, err := syscall.SyscallN(proc.HeapCreate.Addr(),
		uintptr(options), uintptr(initialSize), uintptr(maximumSize))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return HHEAP(ret), nil
}

// ⚠️ You must defer HHEAP.HeapFree().
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
func (hHeap HHEAP) HeapAlloc(
	flags co.HEAP_ALLOC, num_bytes uint) ([]byte, error) {

	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, err := syscall.SyscallN(proc.HeapAlloc.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(num_bytes))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), num_bytes), nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcompact
func (hHeap HHEAP) HeapCompact(flags co.HEAP_NS) (uint, error) {
	ret, _, err := syscall.SyscallN(proc.HeapCompact.Addr(),
		uintptr(hHeap), uintptr(flags))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint(ret), nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapdestroy
func (hHeap HHEAP) HeapDestroy() error {
	ret, _, err := syscall.SyscallN(proc.HeapDestroy.Addr(),
		uintptr(hHeap))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapfree
func (hHeap HHEAP) HeapFree(flags co.HEAP_NS, block []byte) error {
	ret, _, err := syscall.SyscallN(proc.HeapFree.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(unsafe.Pointer(&block[0])))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// ⚠️ You must defer HHEAP.HeapFree().
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heaprealloc
func (hHeap HHEAP) HeapReAlloc(
	flags co.HEAP_REALLOC, block []byte, num_bytes uint) ([]byte, error) {

	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, err := syscall.SyscallN(proc.HeapReAlloc.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(unsafe.Pointer(&block[0])),
		uintptr(num_bytes))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), num_bytes), nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapsize
func (hHeap HHEAP) HeapSize(flags co.HEAP_NS, block []byte) (uint, error) {
	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, err := syscall.SyscallN(proc.HeapSize.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(unsafe.Pointer(&block[0])))
	if int64(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return uint(ret), nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapvalidate
func (hHeap HHEAP) HeapValidate(flags co.HEAP_NS, block []byte) bool {
	syscall.SyscallN(proc.SetLastError.Addr(), 0)

	ret, _, _ := syscall.SyscallN(proc.HeapValidate.Addr(),
		uintptr(hHeap), uintptr(flags), uintptr(unsafe.Pointer(&block[0])))
	return ret != 0
}
