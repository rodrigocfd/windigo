/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// This type doesn't exist in Win32, it's just a HANDLE.
type HFIND HANDLE

func (hFind HFIND) FindClose() {
	ret, lerr := hFind.findCloseNoPanic()
	if ret == 0 {
		panic(fmt.Sprintf("FindClose failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hFind HFIND) findCloseNoPanic() (uintptr, syscall.Errno) {
	ret, _, lerr := syscall.Syscall(proc.FindClose.Addr(), 1,
		uintptr(hFind), 0, 0)
	return ret, lerr
}

// Returns true/false if a file was found or not.
func FindFirstFile(lpFileName string,
	lpFindFileData *WIN32_FIND_DATA) (bool, HFIND) {

	ret, _, lerr := syscall.Syscall(proc.FindFirstFile.Addr(), 2,
		uintptr(unsafe.Pointer(StrToPtr(lpFileName))),
		uintptr(unsafe.Pointer(lpFindFileData)), 0)

	if int32(ret) == -1 { // INVALID_HANDLE_VALUE
		if co.ERROR(lerr) == co.ERROR_FILE_NOT_FOUND { // no matching files, not an error
			return false, 0
		} else {
			panic(fmt.Sprintf("FindFirstFile failed: %d %s",
				lerr, lerr.Error()))
		}
	}
	return true, HFIND(ret)
}

func (hFind HFIND) FindNextFile(lpFindFileData *WIN32_FIND_DATA) bool {
	ret, _, lerr := syscall.Syscall(proc.FindNextFile.Addr(), 2,
		uintptr(hFind), uintptr(unsafe.Pointer(lpFindFileData)), 0)

	if ret == 0 {
		if co.ERROR(lerr) == co.ERROR_NO_MORE_FILES { // not an error, search ended
			return false
		} else {
			hFind.findCloseNoPanic() // cleanup
			panic(fmt.Sprintf("FindNextFile failed: %d %s",
				lerr, lerr.Error()))
		}
	}
	return true
}
