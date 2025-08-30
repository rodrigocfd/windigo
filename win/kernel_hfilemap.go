//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a memory-mapped [file].
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type HFILEMAP HANDLE

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hMap HFILEMAP) CloseHandle() error {
	return HANDLE(hMap).CloseHandle()
}

// [MapViewOfFile] function.
//
// The offset will be rounded down to a multiple of the allocation granularity,
// which is taken with [GetSystemInfo].
//
// Note that this function may present issues in x86 architectures.
//
// ⚠️ You must defer [HFILEMAPVIEW.UnmapViewOfFile].
//
// [MapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(
	desiredAccess co.FILE_MAP,
	offset uint64,
	numBytesToMap uint,
) (HFILEMAPVIEW, error) {
	si := GetSystemInfo()
	if (offset % uint64(si.DwAllocationGranularity)) != 0 {
		offset -= offset % uint64(si.DwAllocationGranularity)
	}

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_MapViewOfFileFromApp, "MapViewOfFileFromApp"),
		uintptr(hMap),
		uintptr(desiredAccess),
		uintptr(offset),
		uintptr(numBytesToMap))
	if ret == 0 {
		return HFILEMAPVIEW(0), co.ERROR(err)
	}
	return HFILEMAPVIEW(ret), nil
}

var _MapViewOfFileFromApp *syscall.Proc

// Handle to the memory block of a memory-mapped [file]. Actually, this is the
// starting address of the mapped view.
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
type HFILEMAPVIEW HANDLE

// Returns a pointer to the beginning of the mapped memory block.
func (hMem HFILEMAPVIEW) Ptr() *byte {
	return (*byte)(unsafe.Pointer(hMem))
}

// [FlushViewOfFile] function.
//
// [FlushViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-flushviewoffile
func (hMem HFILEMAPVIEW) FlushViewOfFile(numBytes uint) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_FlushViewOfFile, "FlushViewOfFile"),
		uintptr(hMem),
		uintptr(numBytes))
	return utl.ZeroAsGetLastError(ret, err)
}

var _FlushViewOfFile *syscall.Proc

// [UnmapViewOfFile] function.
//
// Paired with [HFILEMAP.MapViewOfFile].
//
// [UnmapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (hMem HFILEMAPVIEW) UnmapViewOfFile() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_UnmapViewOfFile, "UnmapViewOfFile"),
		uintptr(hMem))
	return utl.ZeroAsGetLastError(ret, err)
}

var _UnmapViewOfFile *syscall.Proc
