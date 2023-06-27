//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle returned by [FindFirstFile] function.
//
// [FindFirstFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
type HFIND HANDLE

// [FindFirstFile] function.
//
// Returns true if a file was found.
//
// ⚠️ You must defer HFIND.FindClose().
//
// [FindFirstFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func FindFirstFile(fileName string,
	findFileData *WIN32_FIND_DATA) (HFIND, bool, error) {

	ret, _, err := syscall.SyscallN(proc.FindFirstFile.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(findFileData)))

	if int(ret) == _INVALID_HANDLE_VALUE {
		if wErr := errco.ERROR(err); wErr == errco.FILE_NOT_FOUND || wErr == errco.PATH_NOT_FOUND {
			return HFIND(0), false, nil // no matching files, not an error
		} else {
			return HFIND(0), false, wErr
		}
	}
	return HFIND(ret), true, nil // a file was found
}

// [FindClose] function.
//
// [FindClose]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() error {
	ret, _, err := syscall.SyscallN(proc.FindClose.Addr(),
		uintptr(hFind))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [FindNextFile] function.
//
// Returns true if a file was found.
//
// [FindNextFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(findFileData *WIN32_FIND_DATA) (bool, error) {
	ret, _, err := syscall.SyscallN(proc.FindNextFile.Addr(),
		uintptr(hFind), uintptr(unsafe.Pointer(findFileData)))

	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a file was found
}
