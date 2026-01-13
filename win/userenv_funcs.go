//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/wstr"
)

// [GetAllUsersProfileDirectory] function.
//
// [GetAllUsersProfileDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/userenv/nf-userenv-getallusersprofiledirectoryw
func GetAllUsersProfileDirectory() (string, error) {
	var szBuf uint32
	_, _, err := syscall.SyscallN(
		dll.Userenv.Load(&_userenv_GetAllUsersProfileDirectoryW, "GetAllUsersProfileDirectoryW"),
		0,
		uintptr(unsafe.Pointer(&szBuf)))
	if wErr := co.ERROR(err); wErr != co.ERROR_INSUFFICIENT_BUFFER {
		return "", wErr
	}

	buf := make([]uint16, szBuf)
	ret, _, err := syscall.SyscallN(
		dll.Userenv.Load(&_userenv_GetAllUsersProfileDirectoryW, "GetAllUsersProfileDirectoryW"),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szBuf)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.DecodeSlice(buf[:]), nil
}

var _userenv_GetAllUsersProfileDirectoryW *syscall.Proc

// [GetDefaultUserProfileDirectory] function.
//
// [GetDefaultUserProfileDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/userenv/nf-userenv-getdefaultuserprofiledirectoryw
func GetDefaultUserProfileDirectory() (string, error) {
	var szBuf uint32
	_, _, err := syscall.SyscallN(
		dll.Userenv.Load(&_userenv_GetDefaultUserProfileDirectoryW, "GetDefaultUserProfileDirectoryW"),
		0,
		uintptr(unsafe.Pointer(&szBuf)))
	if wErr := co.ERROR(err); wErr != co.ERROR_INSUFFICIENT_BUFFER {
		return "", wErr
	}

	buf := make([]uint16, szBuf)
	ret, _, err := syscall.SyscallN(
		dll.Userenv.Load(&_userenv_GetDefaultUserProfileDirectoryW, "GetDefaultUserProfileDirectoryW"),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szBuf)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.DecodeSlice(buf[:]), nil
}

var _userenv_GetDefaultUserProfileDirectoryW *syscall.Proc

// [GetProfilesDirectory] function.
//
// [GetProfilesDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/userenv/nf-userenv-getprofilesdirectoryw
func GetProfilesDirectory() (string, error) {
	var szBuf uint32
	_, _, err := syscall.SyscallN(
		dll.Userenv.Load(&_userenv_GetProfilesDirectoryW, "GetProfilesDirectoryW"),
		0,
		uintptr(unsafe.Pointer(&szBuf)))
	if wErr := co.ERROR(err); wErr != co.ERROR_INSUFFICIENT_BUFFER {
		return "", wErr
	}

	buf := make([]uint16, szBuf)
	ret, _, err := syscall.SyscallN(
		dll.Userenv.Load(&_userenv_GetProfilesDirectoryW, "GetProfilesDirectoryW"),
		uintptr(unsafe.Pointer(unsafe.SliceData(buf))),
		uintptr(unsafe.Pointer(&szBuf)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.DecodeSlice(buf[:]), nil
}

var _userenv_GetProfilesDirectoryW *syscall.Proc
