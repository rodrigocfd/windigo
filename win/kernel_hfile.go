//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Handle to a [file].
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HFILE HANDLE

// [CreateFile] function.
//
// ⚠️ You must defer HFILE.CloseHandle().
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
	fileName16 := wstr.NewBufWith[wstr.Stack20](fileName, wstr.EMPTY_IS_NIL)

	ret, _, err := syscall.SyscallN(_CreateFileW.Addr(),
		uintptr(fileName16.UnsafePtr()),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile))

	if int(ret) == util.INVALID_HANDLE_VALUE {
		return HFILE(0), co.ERROR(err)
	}
	return HFILE(ret), nil
}

var _CreateFileW = dll.Kernel32.NewProc("CreateFileW")

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() error {
	return HANDLE(hFile).CloseHandle()
}

// [GetFileSizeEx] function.
//
// [GetFileSizeEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() (uint, error) {
	var retSz int64
	ret, _, err := syscall.SyscallN(_GetFileSizeEx.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&retSz)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(retSz), nil
}

var _GetFileSizeEx = dll.Kernel32.NewProc("GetFileSizeEx")

// [CreateFileMapping] function.
//
// ⚠️ You must defer HFILEMAP.CloseHandle().
//
// [CreateFileMapping]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(
	securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE,
	protectSec co.SEC,
	maxSize uint,
	objectName string,
) (HFILEMAP, error) {
	objectName16 := wstr.NewBufWith[wstr.Stack20](objectName, wstr.EMPTY_IS_NIL)

	ret, _, err := syscall.SyscallN(_CreateFileMappingFromApp.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)), uintptr(maxSize),
		uintptr(objectName16.UnsafePtr()))
	if ret == 0 {
		return HFILEMAP(0), co.ERROR(err)
	}
	return HFILEMAP(ret), nil
}

var _CreateFileMappingFromApp = dll.Kernel32.NewProc("CreateFileMappingFromApp")

// [GetFileTime] function.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfiletime
func (hFile HFILE) GetFileTime() (creation, lastAccess, lastWrite time.Time, wErr error) {
	var ftCreation FILETIME
	var ftLastAccess FILETIME
	var ftLastWrite FILETIME

	ret, _, err := syscall.SyscallN(_GetFileTime.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftLastAccess)),
		uintptr(unsafe.Pointer(&ftLastWrite)))
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return ftCreation.ToTime(), ftLastAccess.ToTime(), ftLastWrite.ToTime(), nil
}

var _GetFileTime = dll.Kernel32.NewProc("GetFileTime")

// [LockFile] function.
//
// ⚠️ You must defer HFILE.UnlockFile().
//
// [LockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfile
func (hFile HFILE) LockFile(offset, numBytes uint) error {
	offsetLo, offsetHi := util.Break64(uint64(offset))
	numBytesLo, numBytesHi := util.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(_LockFile.Addr(),
		uintptr(hFile), uintptr(offsetLo), uintptr(offsetHi),
		uintptr(numBytesLo), uintptr(numBytesHi))
	return util.ZeroAsGetLastError(ret, err)
}

var _LockFile = dll.Kernel32.NewProc("LockFile")

// [LockFileEx] function.
//
// ⚠️ You must defer HFILE.UnlockFileEx().
//
// [LockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfileex
func (hFile HFILE) LockFileEx(
	flags co.LOCKFILE,
	numBytes uint,
	overlapped *OVERLAPPED,
) error {
	numBytesLo, numBytesHi := util.Break64(uint64(numBytes))
	ret, _, err := syscall.SyscallN(_LockFileEx.Addr(),
		uintptr(hFile), uintptr(flags), 0,
		uintptr(numBytesLo), uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return util.ZeroAsGetLastError(ret, err)
}

var _LockFileEx = dll.Kernel32.NewProc("LockFileEx")

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(
	buffer []byte,
	overlapped *OVERLAPPED,
) (numBytesRead uint, wErr error) {
	var numBytesRead32 uint32
	ret, _, err := syscall.SyscallN(_ReadFile.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(uint32(len(buffer))), uintptr(unsafe.Pointer(&numBytesRead32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		numBytesRead = 0
	} else {
		numBytesRead, wErr = uint(numBytesRead32), nil
	}
	return
}

var _ReadFile = dll.Kernel32.NewProc("ReadFile")

// [SetEndOfFile] function.
//
// [SetEndOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, err := syscall.SyscallN(_SetEndOfFile.Addr(),
		uintptr(hFile))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return err
	}
	return nil
}

var _SetEndOfFile = dll.Kernel32.NewProc("SetEndOfFile")

// [UnlockFile] function.
//
// Paired with [LockFile].
//
// [UnlockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfile
// [LockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfile
func (hFile HFILE) UnlockFile(offset, numBytes uint) error {
	offsetLo, offsetHi := util.Break64(uint64(offset))
	numBytesLo, numBytesHi := util.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(_UnlockFile.Addr(),
		uintptr(hFile), uintptr(offsetLo), uintptr(offsetHi),
		uintptr(numBytesLo), uintptr(numBytesHi))
	return util.ZeroAsGetLastError(ret, err)
}

var _UnlockFile = dll.Kernel32.NewProc("UnlockFile")

// [UnlockFileEx] function.
//
// Paired with [LockFileEx].
//
// [UnlockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfileex
// [LockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfileex
func (hFile HFILE) UnlockFileEx(numBytes uint, overlapped *OVERLAPPED) error {
	numBytesLo, numBytesHi := util.Break64(uint64(numBytes))
	ret, _, err := syscall.SyscallN(_UnlockFileEx.Addr(),
		uintptr(hFile), 0, uintptr(numBytesLo), uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return util.ZeroAsGetLastError(ret, err)
}

var _UnlockFileEx = dll.Kernel32.NewProc("UnlockFileEx")

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(
	data []byte,
	overlapped *OVERLAPPED,
) (numBytesWritten uint, wErr error) {
	var numBytesWritten32 uint32
	ret, _, err := syscall.SyscallN(_WriteFile.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&data[0])),
		uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&numBytesWritten32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		numBytesWritten = 0
	} else {
		numBytesWritten, wErr = uint(numBytesWritten32), nil
	}
	return
}

var _WriteFile = dll.Kernel32.NewProc("WriteFile")
