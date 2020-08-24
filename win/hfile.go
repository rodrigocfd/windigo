/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HFILE HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() {
	lerr := hFile.closeHandleNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("CloseHandle failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(pathName string, securityAttributes *SECURITY_ATTRIBUTES) {
	ret, _, lerr := syscall.Syscall(proc.CreateDirectory.Addr(), 2,
		uintptr(unsafe.Pointer(StrToPtr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("CreateDirectory failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
func CreateFile(fileName string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition co.FILE_DISPO, attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG, security co.SECURITY, hTemplateFile HFILE) HFILE {

	ret, _, lerr := syscall.Syscall9(proc.CreateFile.Addr(), 7,
		uintptr(unsafe.Pointer(StrToPtr(fileName))),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile), 0, 0)
	if int(ret) == -1 {
		panic(co.ERROR(lerr).Format("CreateFile failed."))
	}
	return HFILE(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE, protectSec co.SEC, maxSize uint32,
	objectName string) HFILEMAP {

	ret, _, lerr := syscall.Syscall6(proc.CreateFileMapping.Addr(), 6,
		uintptr(hFile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		0, uintptr(maxSize),
		uintptr(unsafe.Pointer(StrToPtrBlankIsNil(objectName))))

	if lerr != 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("CreateFileMapping failed."))
	}
	return HFILEMAP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) {
	ret, _, lerr := syscall.Syscall(proc.DeleteFile.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtr(fileName))), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DeleteFile failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
func GetFileAttributes(lpFileName string) (co.FILE_ATTRIBUTE, co.ERROR) {
	ret, _, lerr := syscall.Syscall(proc.GetFileAttributes.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtr(lpFileName))), 0, 0)
	return co.FILE_ATTRIBUTE(ret), co.ERROR(lerr)
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesize
func (hFile HFILE) GetFileSize() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetFileSize.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0xFFFFFFFF && lerr != 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("GetFileSize failed."))
	}
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() int64 {
	buf := int64(0)
	ret, _, lerr := syscall.Syscall(proc.GetFileSizeEx.Addr(), 2,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf)), 0)
	if ret == 0 && lerr != 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("GetFileSizeEx failed."))
	}
	return buf
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(buf []byte, numBytesToRead uint32) uint32 {
	if len(buf) < int(numBytesToRead) {
		panic("ReadFile failed: not enough room in buffer.")
	}

	numRead := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.ReadFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(numBytesToRead), uintptr(unsafe.Pointer(&numRead)), 0, 0) // OVERLAPPED not even considered

	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("ReadFile failed."))
	}
	return numRead
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() {
	ret, _, lerr := syscall.Syscall(proc.SetEndOfFile.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("SetEndOfFile failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointer
func (hFile HFILE) SetFilePointer(distanceToMove int32,
	moveMethod co.FILE_SETPTR) {

	ret, _, lerr := syscall.Syscall6(proc.SetFilePointer.Addr(), 4,
		uintptr(hFile), uintptr(distanceToMove), 0, uintptr(moveMethod),
		0, 0)
	if int(ret) == -1 && lerr != 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("SetFilePointer failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
func (hFile HFILE) SetFilePointerEx(distanceToMove int64,
	moveMethod co.FILE_SETPTR) {

	ret, _, lerr := syscall.Syscall6(proc.SetFilePointer.Addr(), 4,
		uintptr(hFile), uintptr(distanceToMove), 0, uintptr(moveMethod),
		0, 0)
	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("SetFilePointerEx failed."))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(buf []byte) {
	written := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.WriteFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0) // OVERLAPPED not even considered
	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("WriteFile failed."))
	}
}

func (hFile HFILE) closeHandleNoPanic() co.ERROR {
	if hFile == 0 { // handle is null, do nothing
		return co.ERROR_SUCCESS
	}
	ret, _, lerr := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0 { // an error occurred
		return co.ERROR(lerr)
	}
	return co.ERROR_SUCCESS
}
