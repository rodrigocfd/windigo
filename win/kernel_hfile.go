//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pFileName := wbuf.PtrEmptyIsNil(fileName)

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_CreateFileW),
		uintptr(pFileName),
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
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetFileSizeEx),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&retSz)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(retSz), nil
}

// [CreateFileMapping] function.
//
// ⚠️ You must defer [HFILEMAP.CloseHandle].
//
// [CreateFileMapping]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(
	securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE,
	protectSec co.SEC,
	maxSize uint,
	objectName string,
) (HFILEMAP, error) {
	maxLo, maxHi := utl.Break64(uint64(maxSize))

	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pObjectName := wbuf.PtrEmptyIsNil(objectName)

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_CreateFileMappingFromApp),
		uintptr(hFile),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		uintptr(maxHi),
		uintptr(maxLo),
		uintptr(pObjectName))
	if ret == 0 {
		return HFILEMAP(0), co.ERROR(err)
	}
	return HFILEMAP(ret), nil
}

// [GetFileTime] function.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfiletime
func (hFile HFILE) GetFileTime() (creation, lastAccess, lastWrite time.Time, wErr error) {
	var ftCreation FILETIME
	var ftLastAccess FILETIME
	var ftLastWrite FILETIME

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetFileTime),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftLastAccess)),
		uintptr(unsafe.Pointer(&ftLastWrite)))
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return ftCreation.ToTime(), ftLastAccess.ToTime(), ftLastWrite.ToTime(), nil
}

// [LockFile] function.
//
// ⚠️ You must defer [HFILE.UnlockFile].
//
// [LockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfile
func (hFile HFILE) LockFile(offset, numBytes uint) error {
	offsetLo, offsetHi := utl.Break64(uint64(offset))
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_LockFile),
		uintptr(hFile),
		uintptr(offsetLo),
		uintptr(offsetHi),
		uintptr(numBytesLo),
		uintptr(numBytesHi))
	return utl.ZeroAsGetLastError(ret, err)
}

// [LockFileEx] function.
//
// ⚠️ You must defer [HFILE.UnlockFileEx].
//
// [LockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfileex
func (hFile HFILE) LockFileEx(
	flags co.LOCKFILE,
	numBytes uint,
	overlapped *OVERLAPPED,
) error {
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_LockFileEx),
		uintptr(hFile),
		uintptr(flags),
		0,
		uintptr(numBytesLo),
		uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(
	buffer []byte,
	overlapped *OVERLAPPED,
) (numBytesRead uint, wErr error) {
	var numBytesRead32 uint32
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_ReadFile),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(uint32(len(buffer))),
		uintptr(unsafe.Pointer(&numBytesRead32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		numBytesRead = 0
	} else {
		numBytesRead, wErr = uint(numBytesRead32), nil
	}
	return
}

// [SetEndOfFile] function.
//
// [SetEndOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetEndOfFile),
		uintptr(hFile))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return err
	}
	return nil
}

// [UnlockFile] function.
//
// Paired with [HFILE.LockFile].
//
// [UnlockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfile
func (hFile HFILE) UnlockFile(offset, numBytes uint) error {
	offsetLo, offsetHi := utl.Break64(uint64(offset))
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_UnlockFile),
		uintptr(hFile),
		uintptr(offsetLo),
		uintptr(offsetHi),
		uintptr(numBytesLo),
		uintptr(numBytesHi))
	return utl.ZeroAsGetLastError(ret, err)
}

// [UnlockFileEx] function.
//
// Paired with [HFILE.LockFileEx].
//
// [UnlockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfileex
func (hFile HFILE) UnlockFileEx(numBytes uint, overlapped *OVERLAPPED) error {
	numBytesLo, numBytesHi := utl.Break64(uint64(numBytes))
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_UnlockFileEx),
		uintptr(hFile),
		0,
		uintptr(numBytesLo),
		uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return utl.ZeroAsGetLastError(ret, err)
}

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(
	data []byte,
	overlapped *OVERLAPPED,
) (numBytesWritten uint, wErr error) {
	var numBytesWritten32 uint32
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_WriteFile),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(uint32(len(data))),
		uintptr(unsafe.Pointer(&numBytesWritten32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		numBytesWritten = 0
	} else {
		numBytesWritten, wErr = uint(numBytesWritten32), nil
	}
	return
}
