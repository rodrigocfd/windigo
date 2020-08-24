/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"wingows/co"
	"wingows/win/proc"
)

// This type doesn't exist in Win32, just a HANDLE.
type HFILEMAP HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hMap HFILEMAP) CloseHandle() {
	lerr := hMap.closeHandleNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("CloseHandle failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(desiredAccess co.FILE_MAP,
	offset uint32, numBytesToMap uintptr) HFILEMAP_PTR {

	ret, _, lerr := syscall.Syscall6(proc.MapViewOfFile.Addr(), 5,
		uintptr(hMap), uintptr(desiredAccess), 0, uintptr(offset),
		numBytesToMap, 0)
	if ret == 0 {
		hMap.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("MapViewOfFile failed."))
	}
	return HFILEMAP_PTR(ret)
}

func (hMap HFILEMAP) closeHandleNoPanic() co.ERROR {
	if hMap == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hMap), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}

//------------------------------------------------------------------------------

// This type doesn't exist in Win32, just a BYTE pointer to memory address.
type HFILEMAP_PTR uintptr

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (mappedPtr HFILEMAP_PTR) UnmapViewOfFile() {
	ret, _, lerr := syscall.Syscall(proc.UnmapViewOfFile.Addr(), 1,
		uintptr(mappedPtr), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("UnmapViewOfFile failed."))
	}
}
