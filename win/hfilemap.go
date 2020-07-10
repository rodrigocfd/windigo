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

func (hMap HFILEMAP) CloseHandle() {
	lerr := hMap.closeHandleNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("CloseHandle failed."))
	}
}

func (hMap HFILEMAP) closeHandleNoPanic() co.ERROR {
	return freeNoPanic(HANDLE(hMap), proc.CloseHandle)
}

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

//------------------------------------------------------------------------------

// This type doesn't exist in Win32, just a BYTE pointer to memory address.
type HFILEMAP_PTR uintptr

func (mappedPtr HFILEMAP_PTR) UnmapViewOfFile() {
	ret, _, lerr := syscall.Syscall(proc.UnmapViewOfFile.Addr(), 1,
		uintptr(mappedPtr), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("UnmapViewOfFile failed."))
	}
}
