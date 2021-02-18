/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	proc "github.com/rodrigocfd/windigo/win/internal"
)

// This type doesn't exist in Win32, it's just a HANDLE. It's defined here so we
// can restrict its methods.
type HFILEMAP HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hMap HFILEMAP) CloseHandle() {
	if hMap != 0 {
		syscall.Syscall(proc.CloseHandle.Addr(), 1,
			uintptr(hMap), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(desiredAccess co.FILE_MAP,
	offset uint32, numBytesToMap uintptr) (HFILEMAP_PTR, error) {

	ret, _, lerr := syscall.Syscall6(proc.MapViewOfFile.Addr(), 5,
		uintptr(hMap), uintptr(desiredAccess), 0, uintptr(offset),
		numBytesToMap, 0)
	if ret == 0 {
		return HFILEMAP_PTR(0), NewWinError(co.ERROR(lerr), "MapViewOfFile")
	}
	return HFILEMAP_PTR(ret), nil
}

//------------------------------------------------------------------------------

// This type doesn't exist in Win32, it's just a BYTE pointer to memory address.
// It's defined here so we can restrict its methods.
type HFILEMAP_PTR uintptr

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (mappedPtr HFILEMAP_PTR) UnmapViewOfFile() {
	if mappedPtr != 0 {
		syscall.Syscall(proc.UnmapViewOfFile.Addr(), 1,
			uintptr(mappedPtr), 0, 0)
	}
}
