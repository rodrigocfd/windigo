/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HFILE HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hFile HFILE) CloseHandle() {
	if hFile != 0 {
		syscall.Syscall(proc.CloseHandle.Addr(), 1,
			uintptr(hFile), 0, 0)
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createdirectoryw
func CreateDirectory(pathName string,
	securityAttributes *SECURITY_ATTRIBUTES) *WinError {

	ret, _, lerr := syscall.Syscall(proc.CreateDirectory.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(pathName))),
		uintptr(unsafe.Pointer(securityAttributes)), 0)
	if ret == 0 {
		return NewWinError(co.ERROR(lerr), "CreateDirectory")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
func CreateFile(fileName string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition co.FILE_DISPO, attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG, security co.SECURITY,
	hTemplateFile HFILE) (HFILE, *WinError) {

	ret, _, lerr := syscall.Syscall9(proc.CreateFile.Addr(), 7,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(fileName))),
		uintptr(desiredAccess), uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(uint32(attributes)|uint32(flags)|uint32(security)),
		uintptr(hTemplateFile), 0, 0)
	if int(ret) == -1 { // INVALID_HANDLE_VALUE
		return HFILE(0), NewWinError(co.ERROR(lerr), "CreateFile")
	}
	return HFILE(ret), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
func (hFile HFILE) CreateFileMapping(securityAttributes *SECURITY_ATTRIBUTES,
	protectPage co.PAGE, protectSec co.SEC, maxSize uint32,
	objectName string) (HFILEMAP, *WinError) {

	ret, _, lerr := syscall.Syscall6(proc.CreateFileMapping.Addr(), 6,
		uintptr(hFile), uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(protectPage)|uint32(protectSec)),
		0, uintptr(maxSize),
		uintptr(unsafe.Pointer(Str.ToUint16PtrBlankIsNil(objectName))))
	if ret == 0 {
		return HFILEMAP(0), NewWinError(co.ERROR(lerr), "CreateFileMapping")
	}
	return HFILEMAP(ret), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-deletefilew
func DeleteFile(fileName string) *WinError {
	ret, _, lerr := syscall.Syscall(proc.DeleteFile.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(fileName))), 0, 0)
	if ret == 0 {
		return NewWinError(co.ERROR(lerr), "DeleteFile")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfileattributesw
func GetFileAttributes(lpFileName string) (co.FILE_ATTRIBUTE, *WinError) {
	ret, _, lerr := syscall.Syscall(proc.GetFileAttributes.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpFileName))), 0, 0)
	retAttr := co.FILE_ATTRIBUTE(ret)
	if retAttr == co.FILE_ATTRIBUTE_INVALID {
		return retAttr, NewWinError(co.ERROR(lerr), "GetFileAttributes")
	}
	return retAttr, nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getfilesizeex
func (hFile HFILE) GetFileSizeEx() (uint64, *WinError) {
	retSz := int64(0)
	ret, _, lerr := syscall.Syscall(proc.GetFileSizeEx.Addr(), 2,
		uintptr(hFile), uintptr(unsafe.Pointer(&retSz)), 0)
	if ret == 0 && co.ERROR(lerr) != co.ERROR_SUCCESS {
		return 0, NewWinError(co.ERROR(lerr), "GetFileSizeEx")
	}
	return uint64(retSz), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hFile HFILE) ReadFile(buf []byte, numBytesToRead uint32) *WinError {
	if len(buf) < int(numBytesToRead) {
		panic("ReadFile failed: not enough room in buffer.")
	}

	numRead := uint32(0) // not used for anything, but must be passed to the call
	ret, _, lerr := syscall.Syscall6(proc.ReadFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(numBytesToRead), uintptr(unsafe.Pointer(&numRead)), 0, 0) // OVERLAPPED not even considered
	if ret == 0 {
		return NewWinError(co.ERROR(lerr), "ReadFile")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setendoffile
func (hFile HFILE) SetEndOfFile() *WinError {
	ret, _, lerr := syscall.Syscall(proc.SetEndOfFile.Addr(), 1,
		uintptr(hFile), 0, 0)
	if ret == 0 {
		return NewWinError(co.ERROR(lerr), "SetEndOfFile")
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
//
// Returns the new pointer offset.
func (hFile HFILE) SetFilePointerEx(
	distanceToMove int64, moveMethod co.FILE_SETPTR) (uint64, *WinError) {

	newOff := int64(0)
	ret, _, lerr := syscall.Syscall6(proc.SetFilePointerEx.Addr(), 4,
		uintptr(hFile), uintptr(distanceToMove),
		uintptr(unsafe.Pointer(&newOff)), uintptr(moveMethod), 0, 0)
	if ret == 0 {
		return 0, NewWinError(co.ERROR(lerr), "SetFilePointerEx")
	}
	return uint64(newOff), nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hFile HFILE) WriteFile(buf []byte) *WinError {
	written := uint32(0)
	ret, _, lerr := syscall.Syscall6(proc.WriteFile.Addr(), 5,
		uintptr(hFile), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0) // OVERLAPPED not even considered
	if ret == 0 {
		return NewWinError(co.ERROR(lerr), "WriteFile")
	}
	return nil
}
