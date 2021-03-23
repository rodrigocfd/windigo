package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// A handle to a file.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HFILE HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() error {
	ret, _, lerr := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0 {
		return err.ERROR(lerr)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() (uint64, error) {
	retSz := int64(0)
	ret, _, lerr := syscall.Syscall(proc.GetFileSizeEx.Addr(), 2,
		uintptr(hFile), uintptr(unsafe.Pointer(&retSz)), 0)

	if errCode := err.ERROR(lerr); ret == 0 && errCode != err.SUCCESS {
		return 0, errCode
	}
	return uint64(retSz), nil
}

// ⚠️ You must defer CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE, protectSec co.SEC, maxSize uint32,
	objectName string) (HFILEMAP, error) {

	ret, _, lerr := syscall.Syscall6(proc.CreateFileMapping.Addr(), 6,
		uintptr(hFile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		0, uintptr(maxSize),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(objectName))))

	if ret == 0 {
		return HFILEMAP(0), err.ERROR(lerr)
	}
	return HFILEMAP(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(buf []byte, numBytesToRead uint32) error {
	numRead := uint32(0) // not used for anything, but must be passed to the call
	ret, _, lerr := syscall.Syscall6(proc.ReadFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(numBytesToRead), uintptr(unsafe.Pointer(&numRead)), 0, 0) // OVERLAPPED not even considered

	if errCode := err.ERROR(lerr); ret == 0 && errCode != err.SUCCESS {
		return lerr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, lerr := syscall.Syscall(proc.SetEndOfFile.Addr(), 1,
		uintptr(hFile), 0, 0)

	if errCode := err.ERROR(lerr); ret == 0 && errCode != err.SUCCESS {
		return lerr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
//
// Returns the new pointer offset.
func (hFile HFILE) SetFilePointerEx(
	distanceToMove int64, moveMethod co.FILE_FROM) (uint64, error) {

	newOff := int64(0)
	ret, _, lerr := syscall.Syscall6(proc.SetFilePointerEx.Addr(), 4,
		uintptr(hFile), uintptr(distanceToMove),
		uintptr(unsafe.Pointer(&newOff)), uintptr(moveMethod), 0, 0)

	if errCode := err.ERROR(lerr); ret == 0 && errCode != err.SUCCESS {
		return 0, lerr
	}
	return uint64(newOff), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(buf []byte) error {
	written := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.WriteFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0) // OVERLAPPED not even considered

	if errCode := err.ERROR(lerr); ret == 0 && errCode != err.SUCCESS {
		return lerr
	}
	return nil
}

//------------------------------------------------------------------------------

// A handle to a memory-mapped file.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type HFILEMAP HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hMap HFILEMAP) CloseHandle() error {
	ret, _, lerr := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hMap), 0, 0)
	if ret == 0 {
		return err.ERROR(lerr)
	}
	return nil
}

// ⚠️ You must defer UnmapViewOfFile().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(desiredAccess co.FILE_MAP,
	offset uint32, numBytesToMap uintptr) (HFILEMAPMEM, error) {

	ret, _, lerr := syscall.Syscall6(proc.MapViewOfFile.Addr(), 5,
		uintptr(hMap), uintptr(desiredAccess), 0, uintptr(offset),
		numBytesToMap, 0)
	if ret == 0 {
		return HFILEMAPMEM(0), err.ERROR(lerr)
	}
	return HFILEMAPMEM(ret), nil
}

//------------------------------------------------------------------------------

// A handle to the memory block of a memory-mapped file.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
type HFILEMAPMEM HANDLE

// Returns a pointer to the beginning of the mapped memory block.
func (hMem HFILEMAPMEM) Ptr() *byte {
	return (*byte)(unsafe.Pointer(hMem))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (hMem HFILEMAPMEM) UnmapViewOfFile() error {
	ret, _, lerr := syscall.Syscall(proc.UnmapViewOfFile.Addr(), 1,
		uintptr(hMem), 0, 0)
	if ret == 0 {
		return err.ERROR(lerr)
	}
	return nil
}
