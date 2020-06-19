/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HFILE HANDLE

func (hfile HFILE) CloseHandle() {
	ret, _, lerr := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hfile), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CloseHandle failed: %d %s",
			lerr, lerr.Error()))
	}
}

func CreateDirectory(pathName string, securityAttributes *SECURITY_ATTRIBUTES) {
	ret, _, lerr := syscall.Syscall(proc.CreateDirectory.Addr(), 2,
		uintptr(unsafe.Pointer(StrToUtf16Ptr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)), 0)
	if ret == 0 {
		panic(fmt.Sprintf("CreateDirectory failed: %d %s",
			lerr, lerr.Error()))
	}
}

func CreateFile(fileName string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition co.FILE_DISPO, attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG, security co.SECURITY, hTemplateFile HFILE) HFILE {

	ret, _, lerr := syscall.Syscall9(proc.CreateFile.Addr(), 7,
		uintptr(unsafe.Pointer(StrToUtf16Ptr(fileName))),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile), 0, 0)
	if int32(ret) == -1 {
		panic(fmt.Sprintf("CreateFile failed: %d %s",
			lerr, lerr.Error()))
	}
	return HFILE(ret)
}

func (hfile HFILE) CreateFileMapping(securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE, protectSec co.SEC, maxSize uint32,
	objectName string) HFILEMAP {

	ret, _, lerr := syscall.Syscall6(proc.CreateFileMapping.Addr(), 6,
		uintptr(hfile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		0, uintptr(maxSize),
		uintptr(unsafe.Pointer(StrToUtf16PtrBlankIsNil(objectName))))
	if lerr != 0 {
		panic(fmt.Sprintf("CreateFileMapping failed: %d %s",
			lerr, lerr.Error()))
	}
	return HFILEMAP(ret)
}

func (hfile HFILE) DeleteFile(fileName string) {
	ret, _, lerr := syscall.Syscall(proc.DeleteFile.Addr(), 1,
		uintptr(unsafe.Pointer(StrToUtf16Ptr(fileName))), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("DeleteFile failed: %d %s",
			lerr, lerr.Error()))
	}
}

func GetFileAttributes(fileName string) co.FILE_ATTRIBUTE {
	ret, _, lerr := syscall.Syscall(proc.GetFileAttributes.Addr(), 1,
		uintptr(unsafe.Pointer(StrToUtf16Ptr(fileName))), 0, 0)
	if int32(ret) == -1 {
		panic(fmt.Sprintf("GetFileAttributes failed: %d %s",
			lerr, lerr.Error()))
	}
	return co.FILE_ATTRIBUTE(ret)
}

func (hfile HFILE) GetFileSize() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetFileSize.Addr(), 1,
		uintptr(hfile), 0, 0)
	if ret == 0xFFFFFFFF && lerr != 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("GetFileSize failed: %d %s",
			lerr, lerr.Error()))
	}
	return uint32(ret)
}

func (hfile HFILE) GetFileSizeEx() int64 {
	buf := int64(0)
	ret, _, lerr := syscall.Syscall(proc.GetFileSizeEx.Addr(), 2,
		uintptr(hfile), uintptr(unsafe.Pointer(&buf)), 0)
	if ret == 0 && lerr != 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("GetFileSizeEx failed: %d %s",
			lerr, lerr.Error()))
	}
	return buf
}

func (hfile HFILE) ReadFile(buf []uint8, bytesToRead uint32) {
	read := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.ReadFile.Addr(), 5,
		uintptr(hfile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(bytesToRead), uintptr(unsafe.Pointer(&read)), 0, 0) // OVERLAPPED not even considered
	if ret == 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("ReadFile failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hfile HFILE) SetEndOfFile() {
	ret, _, lerr := syscall.Syscall(proc.SetEndOfFile.Addr(), 1,
		uintptr(hfile), 0, 0)
	if ret == 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("SetEndOfFile failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hfile HFILE) SetFilePointer(distanceToMove int32,
	moveMethod co.FILE_SETPTR) {

	ret, _, lerr := syscall.Syscall6(proc.SetFilePointer.Addr(), 4,
		uintptr(hfile), uintptr(distanceToMove), 0, uintptr(moveMethod),
		0, 0)
	if int32(ret) == -1 && lerr != 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("SetFilePointer failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hfile HFILE) SetFilePointerEx(distanceToMove int64,
	moveMethod co.FILE_SETPTR) {

	ret, _, lerr := syscall.Syscall6(proc.SetFilePointer.Addr(), 4,
		uintptr(hfile), uintptr(distanceToMove), 0, uintptr(moveMethod),
		0, 0)
	if ret == 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("SetFilePointerEx failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hfile HFILE) WriteFile(buf []uint8) {
	written := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.WriteFile.Addr(), 5,
		uintptr(hfile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0) // OVERLAPPED not even considered
	if ret == 0 {
		hfile.CloseHandle()
		panic(fmt.Sprintf("WriteFile failed: %d %s",
			lerr, lerr.Error()))
	}
}
