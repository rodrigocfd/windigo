//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// A handle returned by [FindFirstFile] function.
//
// [FindFirstFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
type HFIND HANDLE

// [FindFirstFile] function.
//
// Returns true if a file was found.
//
// This is a low-level function, prefer using [EnumFiles] or [EnumFilesDeep].
//
// ⚠️ You must defer [HFIND.FindClose].
//
// [FindFirstFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func FindFirstFile(fileName string, findFileData *WIN32_FIND_DATA) (HFIND, bool, error) {
	fileName16 := wstr.NewBufWith[wstr.Stack20](fileName, wstr.EMPTY_IS_NIL)

	ret, _, err := syscall.SyscallN(_FindFirstFileW.Addr(),
		uintptr(fileName16.UnsafePtr()),
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

var _FindFirstFileW = dll.Kernel32.NewProc("FindFirstFileW")

// [FindClose] function.
//
// Paired with [FindFirstFile].
//
// [FindClose]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() error {
	ret, _, err := syscall.SyscallN(_FindClose.Addr(),
		uintptr(hFind))
	return utl.ZeroAsGetLastError(ret, err)
}

var _FindClose = dll.Kernel32.NewProc("FindClose")

// [FindNextFile] function.
//
// Returns true if a file was found.
//
// This is a low-level function, prefer using [EnumFiles] or [EnumFilesDeep].
//
// [FindNextFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(findFileData *WIN32_FIND_DATA) (bool, error) {
	ret, _, err := syscall.SyscallN(_FindNextFileW.Addr(),
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

var _FindNextFileW = dll.Kernel32.NewProc("FindNextFileW")
