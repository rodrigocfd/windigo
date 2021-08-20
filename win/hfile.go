package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a file.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HFILE HANDLE

// ⚠️ You must defer HFILE.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
func CreateFile(fileName string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition co.DISPOSITION, attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG, security co.SECURITY,
	hTemplateFile HFILE) (HFILE, error) {

	ret, _, err := syscall.Syscall9(proc.CreateFile.Addr(), 7,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(fileName))),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile), 0, 0)

	if int(ret) == _INVALID_HANDLE_VALUE {
		return HFILE(0), errco.ERROR(err)
	}
	return HFILE(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() error {
	ret, _, err := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() (uint64, error) {
	retSz := int64(0)
	ret, _, err := syscall.Syscall(proc.GetFileSizeEx.Addr(), 2,
		uintptr(hFile), uintptr(unsafe.Pointer(&retSz)), 0)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return 0, wErr
	}
	return uint64(retSz), nil
}

// ⚠️ You must defer HFILEMAP.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE, protectSec co.SEC, maxSize uint32,
	objectName string) (HFILEMAP, error) {

	ret, _, err := syscall.Syscall6(proc.CreateFileMapping.Addr(), 6,
		uintptr(hFile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		0, uintptr(maxSize),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(objectName))))

	if ret == 0 {
		return HFILEMAP(0), errco.ERROR(err)
	}
	return HFILEMAP(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(buf []byte, numBytesToRead uint32) error {
	numRead := uint32(0) // not used for anything, but must be passed to the call
	ret, _, err := syscall.Syscall6(proc.ReadFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(numBytesToRead), uintptr(unsafe.Pointer(&numRead)), 0, 0) // OVERLAPPED not even considered

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return err
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() error {
	ret, _, err := syscall.Syscall(proc.SetEndOfFile.Addr(), 1,
		uintptr(hFile), 0, 0)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return err
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
//
// Returns the new pointer offset.
func (hFile HFILE) SetFilePointerEx(
	distanceToMove int64, moveMethod co.FILE_FROM) (uint64, error) {

	newOff := int64(0)
	ret, _, err := syscall.Syscall6(proc.SetFilePointerEx.Addr(), 4,
		uintptr(hFile), uintptr(distanceToMove),
		uintptr(unsafe.Pointer(&newOff)), uintptr(moveMethod), 0, 0)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return 0, err
	}
	return uint64(newOff), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(buf []byte) error {
	written := uint32(0)
	ret, _, err := syscall.Syscall6(proc.WriteFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0) // OVERLAPPED not even considered

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return err
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
	ret, _, err := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hMap), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// ⚠️ You must defer HFILEMAPVIEW.UnmapViewOfFile().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
func (hMap HFILEMAP) MapViewOfFile(desiredAccess co.FILE_MAP,
	offset uint32, numBytesToMap uintptr) (HFILEMAPVIEW, error) {

	ret, _, err := syscall.Syscall6(proc.MapViewOfFile.Addr(), 5,
		uintptr(hMap), uintptr(desiredAccess), 0, uintptr(offset),
		numBytesToMap, 0)
	if ret == 0 {
		return HFILEMAPVIEW(0), errco.ERROR(err)
	}
	return HFILEMAPVIEW(ret), nil
}

//------------------------------------------------------------------------------

// A handle to the memory block of a memory-mapped file. Actually, this is the
// starting address of the mapped view.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
type HFILEMAPVIEW HANDLE

// Returns a pointer to the beginning of the mapped memory block.
func (hMem HFILEMAPVIEW) Ptr() *byte {
	return (*byte)(unsafe.Pointer(hMem))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-unmapviewoffile
func (hMem HFILEMAPVIEW) UnmapViewOfFile() error {
	ret, _, err := syscall.Syscall(proc.UnmapViewOfFile.Addr(), 1,
		uintptr(hMem), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
