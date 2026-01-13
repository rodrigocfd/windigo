//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [SetUserObjectInformation] function.
//
// [SetUserObjectInformation]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setuserobjectinformationw
func (hProcess HPROCESS) SetUserObjectInformation(
	index co.UOI,
	info unsafe.Pointer,
	infoLen uintptr,
) error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_SetUserObjectInformationW, "SetUserObjectInformationW"),
		uintptr(hProcess),
		uintptr(index),
		uintptr(info),
		uintptr(infoLen))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_SetUserObjectInformationW *syscall.Proc
