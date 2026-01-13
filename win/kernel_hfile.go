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
// Unless you're doing something specific, consider using the high-level [File]
// abstraction.
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
		dll.Kernel.Load(&_kernel_CreateFileW, "CreateFileW"),
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

var _kernel_CreateFileW *syscall.Proc

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
		dll.Kernel.Load(&_kernel_GetFileSizeEx, "GetFileSizeEx"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&retSz)))

	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(retSz), nil
}

var _kernel_GetFileSizeEx *syscall.Proc

// [CreateFileMapping] function.
//
// Unless you're doing something specific, consider using the high-level
// [FileMap] abstraction.
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
		dll.Kernel.Load(&_kernel_CreateFileMappingFromApp, "CreateFileMappingFromApp"),
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

var _kernel_CreateFileMappingFromApp *syscall.Proc

// [GetFileTime] function.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfiletime
func (hFile HFILE) GetFileTime() (creation, lastAccess, lastWrite time.Time, wErr error) {
	var ftCreation, ftLastAccess, ftLastWrite FILETIME
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetFileTime, "GetFileTime"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftLastAccess)),
		uintptr(unsafe.Pointer(&ftLastWrite)))
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return ftCreation.ToTime(), ftLastAccess.ToTime(), ftLastWrite.ToTime(), nil
}

var _kernel_GetFileTime *syscall.Proc

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
		dll.Kernel.Load(&_kernel_LockFile, "LockFile"),
		uintptr(hFile),
		uintptr(offsetLo),
		uintptr(offsetHi),
		uintptr(numBytesLo),
		uintptr(numBytesHi))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_LockFile *syscall.Proc

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
		dll.Kernel.Load(&_kernel_LockFileEx, "LockFileEx"),
		uintptr(hFile),
		uintptr(flags),
		0,
		uintptr(numBytesLo),
		uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_LockFileEx *syscall.Proc

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(
	destBuf []byte,
	overlapped *OVERLAPPED,
) (numBytesRead int, wErr error) {
	var numBytesRead32 uint32
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ReadFile, "ReadFile"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(unsafe.SliceData(destBuf))),
		uintptr(uint32(len(destBuf))),
		uintptr(unsafe.Pointer(&numBytesRead32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(numBytesRead32), nil
}

var _kernel_ReadFile *syscall.Proc

// [SetEndOfFile] function.
//
// [SetEndOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_SetEndOfFile, "SetEndOfFile"),
		uintptr(hFile))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return err
	}
	return nil
}

var _kernel_SetEndOfFile *syscall.Proc

// [SetHandleInformation] function.
//
// [SetHandleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-sethandleinformation
func (hFile HFILE) SetHandleInformation(mask, flags co.HANDLE_FLAG) error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_SetHandleInformation, "SetHandleInformation"),
		uintptr(hFile),
		uintptr(mask),
		uintptr(flags))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_SetHandleInformation *syscall.Proc

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
		dll.Kernel.Load(&_kernel_UnlockFile, "UnlockFile"),
		uintptr(hFile),
		uintptr(offsetLo),
		uintptr(offsetHi),
		uintptr(numBytesLo),
		uintptr(numBytesHi))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_UnlockFile *syscall.Proc

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
		dll.Kernel.Load(&_kernel_UnlockFileEx, "UnlockFileEx"),
		uintptr(hFile),
		0,
		uintptr(numBytesLo),
		uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_UnlockFileEx *syscall.Proc

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(
	data []byte,
	overlapped *OVERLAPPED,
) (numBytesWritten int, wErr error) {
	var written32 uint32
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_WriteFile, "WriteFile"),
		uintptr(hFile),
		uintptr(unsafe.Pointer(unsafe.SliceData(data))),
		uintptr(uint32(len(data))),
		uintptr(unsafe.Pointer(&written32)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(written32), nil
}

var _kernel_WriteFile *syscall.Proc

// Handle to a memory-mapped [file].
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type HFILEMAP HANDLE

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hMap HFILEMAP) CloseHandle() error {
	return HANDLE(hMap).CloseHandle()
}

// [MapViewOfFile] function.
//
// The offset will be rounded down to a multiple of the allocation granularity,
// which is taken with [GetSystemInfo].
//
// Note that this function may present issues in x86 architectures.
//
// Panics if offset or numBytesToMap is negative.
//
// ⚠️ You must defer [HFILEMAPVIEW.UnmapViewOfFile].
//
// [MapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(
	desiredAccess co.FILE_MAP,
	offset int,
	numBytesToMap int,
) (HFILEMAPVIEW, error) {
	utl.PanicNeg(offset, numBytesToMap)

	si := GetSystemInfo()
	offset64 := uint64(offset)
	if (offset64 % uint64(si.DwAllocationGranularity)) != 0 {
		offset64 -= offset64 % uint64(si.DwAllocationGranularity)
	}

	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_MapViewOfFileFromApp, "MapViewOfFileFromApp"),
		uintptr(hMap),
		uintptr(desiredAccess),
		uintptr(offset64),
		uintptr(uint64(numBytesToMap)))
	if ret == 0 {
		return HFILEMAPVIEW(0), co.ERROR(err)
	}
	return HFILEMAPVIEW(ret), nil
}

var _kernel_MapViewOfFileFromApp *syscall.Proc

// [SetHandleInformation] function.
//
// [SetHandleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-sethandleinformation
func (hMap HFILEMAP) SetHandleInformation(mask, flags co.HANDLE_FLAG) error {
	return HFILE(hMap).SetHandleInformation(mask, flags)
}

// Handle to the memory block of a memory-mapped [file]. Actually, this is the
// starting address of the mapped view.
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
type HFILEMAPVIEW HANDLE

// Returns a pointer to the beginning of the mapped memory block.
func (hMem HFILEMAPVIEW) Ptr() *byte {
	return (*byte)(unsafe.Pointer(hMem))
}

// [FlushViewOfFile] function.
//
// Panics if numBytes is negative.
//
// [FlushViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-flushviewoffile
func (hMem HFILEMAPVIEW) FlushViewOfFile(numBytes int) error {
	utl.PanicNeg(numBytes)
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_FlushViewOfFile, "FlushViewOfFile"),
		uintptr(hMem),
		uintptr(uint64(numBytes)))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_FlushViewOfFile *syscall.Proc

// [UnmapViewOfFile] function.
//
// Paired with [HFILEMAP.MapViewOfFile].
//
// [UnmapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (hMem HFILEMAPVIEW) UnmapViewOfFile() error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_UnmapViewOfFile, "UnmapViewOfFile"),
		uintptr(hMem))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_UnmapViewOfFile *syscall.Proc
