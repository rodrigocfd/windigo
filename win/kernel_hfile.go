//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to a [file].
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HFILE HANDLE

// [CreateFile] function.
//
// ⚠️ You must defer [HFILE.CloseHandle].
//
// [CreateFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
func CreateFile(
	fileName string,
	desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE,
	securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition co.DISPOSITION,
	attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG,
	security co.SECURITY,
	hTemplateFile HFILE,
) (HFILE, error) {
	var wFileName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CreateFileW, "CreateFileW"),
		uintptr(wFileName.EmptyIsNil(fileName)),
		uintptr(desiredAccess),
		uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile))

	if int(ret) == utl.INVALID_HANDLE_VALUE {
		return HFILE(0), co.ERROR(err)
	}
	return HFILE(ret), nil
}

var _CreateFileW *syscall.Proc

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() error {
	return HANDLE(hFile).CloseHandle()
}

// [GetFileSizeEx] function.
//
// [GetFileSizeEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() (int, error) {
	var retSz int64
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetFileSizeEx, "GetFileSizeEx"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&retSz)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(retSz), nil
}

var _GetFileSizeEx *syscall.Proc

// [CreateFileMapping] function.
//
// Panics if maxSize is negative.
//
// ⚠️ You must defer [HFILEMAP.CloseHandle].
//
// [CreateFileMapping]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(
	securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE,
	protectSec co.SEC,
	maxSize int,
	objectName string,
) (HFILEMAP, error) {
	utl.PanicNeg(maxSize)

	maxLo, maxHi := utl.Break64(uint64(maxSize))
	var wObjectName wstr.BufEncoder

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CreateFileMappingFromApp, "CreateFileMappingFromApp"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		uintptr(maxHi),
		uintptr(maxLo),
		uintptr(wObjectName.EmptyIsNil(objectName)))
	if ret == 0 {
		return HFILEMAP(0), co.ERROR(err)
	}
	return HFILEMAP(ret), nil
}

var _CreateFileMappingFromApp *syscall.Proc

// [GetFileTime] function.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfiletime
func (hFile HFILE) GetFileTime() (creation, lastAccess, lastWrite time.Time, wErr error) {
	var ftCreation, ftLastAccess, ftLastWrite FILETIME
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetFileTime, "GetFileTime"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftLastAccess)),
		uintptr(unsafe.Pointer(&ftLastWrite)))
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return ftCreation.ToTime(), ftLastAccess.ToTime(), ftLastWrite.ToTime(), nil
}

var _GetFileTime *syscall.Proc

// [LockFile] function.
//
// Panics if offset or numBytes is negative.
//
// ⚠️ You must defer [HFILE.UnlockFile].
//
// [LockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfile
func (hFile HFILE) LockFile(offset, numBytes int) error {
	utl.PanicNeg(offset, numBytes)
	offsetLo, offsetHi := utl.Break64(uint64(offset))
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LockFile, "LockFile"),
		uintptr(hFile),
		uintptr(offsetLo),
		uintptr(offsetHi),
		uintptr(numBytesLo),
		uintptr(numBytesHi))
	return utl.ZeroAsGetLastError(ret, err)
}

var _LockFile *syscall.Proc

// [LockFileEx] function.
//
// Panics if numBytes is negative.
//
// ⚠️ You must defer [HFILE.UnlockFileEx].
//
// [LockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfileex
func (hFile HFILE) LockFileEx(
	flags co.LOCKFILE,
	numBytes int,
	overlapped *OVERLAPPED,
) error {
	utl.PanicNeg(numBytes)
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LockFileEx, "LockFileEx"),
		uintptr(hFile),
		uintptr(flags),
		0,
		uintptr(numBytesLo),
		uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _LockFileEx *syscall.Proc

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(
	buffer []byte,
	overlapped *OVERLAPPED,
) (numBytesRead int, wErr error) {
	var numBytesRead32 uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ReadFile, "ReadFile"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(uint32(len(buffer))),
		uintptr(unsafe.Pointer(&numBytesRead32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(numBytesRead32), nil
}

var _ReadFile *syscall.Proc

// [SetEndOfFile] function.
//
// [SetEndOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SetEndOfFile, "SetEndOfFile"),
		uintptr(hFile))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return err
	}
	return nil
}

var _SetEndOfFile *syscall.Proc

// [UnlockFile] function.
//
// Paired with [HFILE.LockFile].
//
// Panics if offset or numBytes is negative.
//
// [UnlockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfile
func (hFile HFILE) UnlockFile(offset, numBytes int) error {
	utl.PanicNeg(offset, numBytes)
	offsetLo, offsetHi := utl.Break64(uint64(offset))
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_UnlockFile, "UnlockFile"),
		uintptr(hFile),
		uintptr(offsetLo),
		uintptr(offsetHi),
		uintptr(numBytesLo),
		uintptr(numBytesHi))
	return utl.ZeroAsGetLastError(ret, err)
}

var _UnlockFile *syscall.Proc

// [UnlockFileEx] function.
//
// Paired with [HFILE.LockFileEx].
//
// Panics if numBytes is negative.
//
// [UnlockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfileex
func (hFile HFILE) UnlockFileEx(numBytes int, overlapped *OVERLAPPED) error {
	utl.PanicNeg(numBytes)
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_UnlockFileEx, "UnlockFileEx"),
		uintptr(hFile),
		0,
		uintptr(numBytesLo),
		uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _UnlockFileEx *syscall.Proc

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(
	data []byte,
	overlapped *OVERLAPPED,
) (numBytesWritten int, wErr error) {
	var written32 uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_WriteFile, "WriteFile"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(uint32(len(data))),
		uintptr(unsafe.Pointer(&written32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(written32), nil
}

var _WriteFile *syscall.Proc
