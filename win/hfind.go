package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle returned by FindFirstFile() function.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
type HFIND HANDLE

// Returns true if a file was found.
//
// ⚠️ You must defer FindClose().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func FindFirstFile(lpFileName string,
	lpFindFileData *WIN32_FIND_DATA) (HFIND, bool, error) {

	ret, _, err := syscall.Syscall(proc.FindFirstFile.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpFileName))),
		uintptr(unsafe.Pointer(lpFindFileData)), 0)

	errCode := errco.ERROR(err)
	if int(ret) == _INVALID_HANDLE_VALUE {
		if errCode == errco.FILE_NOT_FOUND || errCode == errco.PATH_NOT_FOUND { // no matching files, not an error
			return HFIND(0), false, nil
		} else {
			return HFIND(0), false, errCode
		}
	}
	return HFIND(ret), true, nil // a file was found
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() {
	ret, _, err := syscall.Syscall(proc.FindClose.Addr(), 1,
		uintptr(hFind), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Returns true if a file was found.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(lpFindFileData *WIN32_FIND_DATA) (bool, error) {
	ret, _, err := syscall.Syscall(proc.FindNextFile.Addr(), 2,
		uintptr(hFind), uintptr(unsafe.Pointer(lpFindFileData)), 0)

	errCode := errco.ERROR(err)
	if ret == 0 {
		if errCode == errco.NO_MORE_FILES { // not an error, search ended
			return false, nil
		} else {
			return false, errCode
		}
	}
	return true, nil // a file was found
}
