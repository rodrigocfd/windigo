/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
	proc "windigo/win/internal"
)

// This type doesn't exist in Win32, it's just a HANDLE.
type HFIND HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() {
	if hFind != 0 {
		syscall.Syscall(proc.FindClose.Addr(), 1,
			uintptr(hFind), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
//
// Returns true if any file was found.
func FindFirstFile(lpFileName string,
	lpFindFileData *WIN32_FIND_DATA) (HFIND, bool, *WinError) {

	ret, _, lerr := syscall.Syscall(proc.FindFirstFile.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpFileName))),
		uintptr(unsafe.Pointer(lpFindFileData)), 0)

	lerr2 := co.ERROR(lerr)
	if int(ret) == _INVALID_HANDLE_VALUE {
		if lerr2 == co.ERROR_FILE_NOT_FOUND ||
			lerr2 == co.ERROR_PATH_NOT_FOUND {
			// No matching files, not an error.
			return HFIND(0), false, nil
		} else {
			return HFIND(0), false, NewWinError(lerr2, "FindFirstFile")
		}
	}
	return HFIND(ret), true, nil // a file was found
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(
	lpFindFileData *WIN32_FIND_DATA) (bool, *WinError) {

	ret, _, lerr := syscall.Syscall(proc.FindNextFile.Addr(), 2,
		uintptr(hFind), uintptr(unsafe.Pointer(lpFindFileData)), 0)

	lerr2 := co.ERROR(lerr)
	if ret == 0 {
		if lerr2 == co.ERROR_NO_MORE_FILES { // not an error, search ended
			return false, nil
		} else {
			return false, NewWinError(lerr2, "FindNextFile")
		}
	}
	return true, nil // a file was found
}
