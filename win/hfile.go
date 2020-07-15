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

func (hFile HFILE) CloseHandle() {
	lerr := hFile.closeHandleNoPanic()
	if lerr != co.ERROR_SUCCESS {
		panic(lerr.Format("CloseHandle failed."))
	}
}

func (hFile HFILE) closeHandleNoPanic() co.ERROR {
	return freeNoPanic(HANDLE(hFile), proc.CloseHandle)
}

func CreateDirectory(pathName string, securityAttributes *SECURITY_ATTRIBUTES) {
	ret, _, lerr := syscall.Syscall(proc.CreateDirectory.Addr(), 2,
		uintptr(unsafe.Pointer(StrToPtr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("CreateDirectory failed."))
	}
}

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

func DeleteFile(fileName string) {
	ret, _, lerr := syscall.Syscall(proc.DeleteFile.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtr(fileName))), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DeleteFile failed."))
	}
}

func GetFileAttributes(lpFileName string) (co.FILE_ATTRIBUTE, co.ERROR) {
	ret, _, lerr := syscall.Syscall(proc.GetFileAttributes.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtr(lpFileName))), 0, 0)
	return co.FILE_ATTRIBUTE(ret), co.ERROR(lerr)
}

func (hFile HFILE) GetFileSize() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetFileSize.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0xFFFFFFFF && lerr != 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("GetFileSize failed."))
	}
	return uint32(ret)
}

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

// Returns the number of bytes actually read.
// Buffer must be previously allocated.
func (hFile HFILE) ReadFile(buf []uint8, bytesToRead uint32) uint32 {
	numRead := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.ReadFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(bytesToRead), uintptr(unsafe.Pointer(&numRead)), 0, 0) // OVERLAPPED not even considered

	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("ReadFile failed."))
	}
	return numRead
}

func (hFile HFILE) SetEndOfFile() {
	ret, _, lerr := syscall.Syscall(proc.SetEndOfFile.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("SetEndOfFile failed."))
	}
}

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

func (hFile HFILE) WriteFile(buf []uint8) {
	written := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.WriteFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0) // OVERLAPPED not even considered
	if ret == 0 {
		hFile.closeHandleNoPanic() // free resource
		panic(co.ERROR(lerr).Format("WriteFile failed."))
	}
}
