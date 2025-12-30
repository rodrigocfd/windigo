//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to a [find object], used to find files.
//
// [find object]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
type HFIND HANDLE

// [FindFirstFile] function.
//
// Returns true if a file was found.
//
// This is a low-level function, prefer using [PathEnum] or [PathEnumDeep].
//
// ⚠️ You must defer [HFIND.FindClose].
//
// [FindFirstFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func FindFirstFile(fileName string, findFileData *WIN32_FIND_DATA) (HFIND, bool, error) {
	var wFileName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_FindFirstFileW, "FindFirstFileW"),
		uintptr(wFileName.EmptyIsNil(fileName)),
		uintptr(unsafe.Pointer(findFileData)))

	if int(ret) == utl.INVALID_HANDLE_VALUE {
		if wErr := co.ERROR(err); wErr == co.ERROR_FILE_NOT_FOUND || wErr == co.ERROR_PATH_NOT_FOUND {
			return HFIND(0), false, nil // no matching files, not an error
		} else {
			return HFIND(0), false, wErr
		}
	}
	return HFIND(ret), true, nil // a file was found
}

var _kernel_FindFirstFileW *syscall.Proc

// [FindClose] function.
//
// Paired with [FindFirstFile].
//
// [FindClose]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_FindClose, "FindClose"),
		uintptr(hFind))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_FindClose *syscall.Proc

// [FindNextFile] function.
//
// Returns true if a file was found.
//
// This is a low-level function, prefer using [PathEnum] or [PathEnumDeep].
//
// [FindNextFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(findFileData *WIN32_FIND_DATA) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_FindNextFileW, "FindNextFileW"),
		uintptr(hFind),
		uintptr(unsafe.Pointer(findFileData)))

	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a file was found
}

var _kernel_FindNextFileW *syscall.Proc
