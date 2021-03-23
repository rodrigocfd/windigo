package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle returned by FindFirstFile() function.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
type HFIND HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() {
	ret, _, lerr := syscall.Syscall(proc.FindClose.Addr(), 1,
		uintptr(hFind), 0, 0)
	if ret == 0 {
		panic(err.ERROR(lerr))
	}
}

// Returns true if a file was found.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(lpFindFileData *WIN32_FIND_DATA) (bool, error) {
	ret, _, lerr := syscall.Syscall(proc.FindNextFile.Addr(), 2,
		uintptr(hFind), uintptr(unsafe.Pointer(lpFindFileData)), 0)

	errCode := err.ERROR(lerr)
	if ret == 0 {
		if errCode == err.NO_MORE_FILES { // not an error, search ended
			return false, nil
		} else {
			return false, errCode
		}
	}
	return true, nil // a file was found
}
