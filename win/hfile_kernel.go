//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [file].
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
	hTemplateFile HFILE) (HFILE, error) {

	ret, _, err := syscall.SyscallN(proc.CreateFile.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile))

	if int(ret) == _INVALID_HANDLE_VALUE {
		return HFILE(0), errco.ERROR(err)
	}
	return HFILE(ret), nil
}

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() error {
	ret, _, err := syscall.SyscallN(proc.CloseHandle.Addr(),
		uintptr(hFile))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [GetFileSizeEx] function.
//
// [GetFileSizeEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() (uint64, error) {
	var retSz int64
	ret, _, err := syscall.SyscallN(proc.GetFileSizeEx.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&retSz)))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return 0, wErr
	}
	return uint64(retSz), nil
}

// [CreateFileMapping] function.
//
// ⚠️ You must defer HFILEMAP.CloseHandle().
//
// [CreateFileMapping]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(
	securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE,
	protectSec co.SEC,
	maxSize uint64,
	objectName StrOpt) (HFILEMAP, error) {

	ret, _, err := syscall.SyscallN(proc.CreateFileMappingFromApp.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		uintptr(maxSize), uintptr(objectName.Raw()))
	if ret == 0 {
		return HFILEMAP(0), errco.ERROR(err)
	}
	return HFILEMAP(ret), nil
}

// [LockFile] function.
//
// ⚠️ You must defer HFILE.UnlockFile().
//
// [LockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfile
func (hFile HFILE) LockFile(offset, numBytes uint64) error {
	offsetLo, offsetHi := util.Break64(offset)
	numBytesLo, numBytesHi := util.Break64(numBytes)

	ret, _, err := syscall.SyscallN(proc.LockFile.Addr(),
		uintptr(hFile), uintptr(offsetLo), uintptr(offsetHi),
		uintptr(numBytesLo), uintptr(numBytesHi))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [LockFileEx] function.
//
// ⚠️ You must defer HFILE.UnlockFileEx().
//
// [LockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfileex
func (hFile HFILE) LockFileEx(
	flags co.LOCKFILE,
	numBytes uint64,
	overlapped *OVERLAPPED) error {

	numBytesLo, numBytesHi := util.Break64(numBytes)
	ret, _, err := syscall.SyscallN(proc.LockFileEx.Addr(),
		uintptr(hFile), uintptr(flags), 0,
		uintptr(numBytesLo), uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(
	buffer []byte, overlapped *OVERLAPPED) (numBytesRead uint32, e error) {

	ret, _, err := syscall.SyscallN(proc.ReadFile.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(uint32(len(buffer))), uintptr(unsafe.Pointer(&numBytesRead)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		numBytesRead, e = 0, wErr
	}
	return
}

// [SetEndOfFile] function.
//
// [SetEndOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, err := syscall.SyscallN(proc.SetEndOfFile.Addr(),
		uintptr(hFile))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return err
	}
	return nil
}

// [SetFilePointerEx] function.
//
// In x86 architecture, [SetFilePointer] will be called instead.
//
// [SetFilePointerEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
// [SetFilePointer]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointer
func (hFile HFILE) SetFilePointerEx(
	distanceToMove int64,
	moveMethod co.FILE_FROM) (newPointerOffset int64, e error) {

	pfunc := proc.SetFilePointerEx
	if runtime.GOARCH == "386" {
		pfunc = proc.SetFilePointer
	}

	ret, _, err := syscall.SyscallN(pfunc.Addr(),
		uintptr(hFile), uintptr(distanceToMove),
		uintptr(unsafe.Pointer(&newPointerOffset)), uintptr(moveMethod))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		newPointerOffset, e = 0, wErr
	}
	return
}

// [UnlockFile] function.
//
// [UnlockFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfile
func (hFile HFILE) UnlockFile(offset, numBytes uint64) error {
	offsetLo, offsetHi := util.Break64(offset)
	numBytesLo, numBytesHi := util.Break64(numBytes)

	ret, _, err := syscall.SyscallN(proc.UnlockFile.Addr(),
		uintptr(hFile), uintptr(offsetLo), uintptr(offsetHi),
		uintptr(numBytesLo), uintptr(numBytesHi))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [UnlockFileEx] function.
//
// [UnlockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-unlockfileex
func (hFile HFILE) UnlockFileEx(numBytes uint64, overlapped *OVERLAPPED) error {
	numBytesLo, numBytesHi := util.Break64(numBytes)
	ret, _, err := syscall.SyscallN(proc.UnlockFileEx.Addr(),
		uintptr(hFile), 0, uintptr(numBytesLo), uintptr(numBytesHi),
		uintptr(unsafe.Pointer(overlapped)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(
	data []byte, overlapped *OVERLAPPED) (numBytesWritten uint32, e error) {

	ret, _, err := syscall.SyscallN(proc.WriteFile.Addr(),
		uintptr(hFile), uintptr(unsafe.Pointer(&data[0])),
		uintptr(uint32(len(data))), uintptr(unsafe.Pointer(&numBytesWritten)),
		uintptr(unsafe.Pointer(overlapped)))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		numBytesWritten, e = 0, wErr
	}
	return
}

//------------------------------------------------------------------------------

// A handle to a memory-mapped [file].
//
// [file]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type HFILEMAP HANDLE

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hMap HFILEMAP) CloseHandle() error {
	ret, _, err := syscall.SyscallN(proc.CloseHandle.Addr(),
		uintptr(hMap))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [MapViewOfFile] function.
//
// The offset will be rounded down to a multiple of the allocation granularity,
// which is taken with GetSystemInfo().
//
// Note that this function may present issues in x86 architectures.
//
// ⚠️ You must defer HFILEMAPVIEW.UnmapViewOfFile().
//
// [MapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(
	desiredAccess co.FILE_MAP,
	offset uint64,
	numBytesToMap uint) (HFILEMAPVIEW, error) {

	si := SYSTEM_INFO{}
	GetSystemInfo(&si)
	if (offset % uint64(si.DwAllocationGranularity)) != 0 {
		offset -= offset % uint64(si.DwAllocationGranularity)
	}

	ret, _, err := syscall.SyscallN(proc.MapViewOfFileFromApp.Addr(),
		uintptr(hMap), uintptr(desiredAccess), uintptr(offset),
		uintptr(numBytesToMap))
	if ret == 0 {
		return HFILEMAPVIEW(0), errco.ERROR(err)
	}
	return HFILEMAPVIEW(ret), nil
}

//------------------------------------------------------------------------------

// A handle to the memory block of a memory-mapped [file]. Actually, this is the
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
// [FlushViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-flushviewoffile
func (hMem HFILEMAPVIEW) FlushViewOfFile(numBytes uint64) error {
	ret, _, err := syscall.SyscallN(proc.FlushViewOfFile.Addr(),
		uintptr(hMem), uintptr(numBytes))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [UnmapViewOfFile] function.
//
// [UnmapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (hMem HFILEMAPVIEW) UnmapViewOfFile() error {
	ret, _, err := syscall.SyscallN(proc.UnmapViewOfFile.Addr(),
		uintptr(hMem))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
