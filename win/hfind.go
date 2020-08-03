/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// This type doesn't exist in Win32, it's just a HANDLE.
type HFIND HANDLE

func (hFind HFIND) FindClose() {
	lerr := hFind.findCloseNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("FindClose failed."))
	}
}

// Returns true/false if a file was found or not.
func FindFirstFile(lpFileName string,
	lpFindFileData *WIN32_FIND_DATA) (HFIND, bool) {

	ret, _, lerr := syscall.Syscall(proc.FindFirstFile.Addr(), 2,
		uintptr(unsafe.Pointer(StrToPtr(lpFileName))),
		uintptr(unsafe.Pointer(lpFindFileData)), 0)

	lerr2 := co.ERROR(lerr)
	if int(ret) == -1 { // INVALID_HANDLE_VALUE
		if lerr2 == co.ERROR_FILE_NOT_FOUND ||
			lerr2 == co.ERROR_PATH_NOT_FOUND {
			// No matching files, not an error.
			return 0, false
		} else {
			panic(lerr2.Format("FindFirstFile failed."))
		}
	}
	return HFIND(ret), true
}

func (hFind HFIND) FindNextFile(lpFindFileData *WIN32_FIND_DATA) bool {
	ret, _, lerr := syscall.Syscall(proc.FindNextFile.Addr(), 2,
		uintptr(hFind), uintptr(unsafe.Pointer(lpFindFileData)), 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 {
		if lerr2 == co.ERROR_NO_MORE_FILES { // not an error, search ended
			return false
		} else {
			hFind.findCloseNoPanic() // free resource
			panic(lerr2.Format("FindNextFile failed."))
		}
	}
	return true
}

func (hFind HFIND) findCloseNoPanic() co.ERROR {
	if hFind == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(proc.FindClose.Addr(), 1,
		uintptr(hFind), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}
