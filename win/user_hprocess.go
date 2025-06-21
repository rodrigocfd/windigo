//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
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
		dll.Load(dll.USER32, &_SetUserObjectInformationW, "SetUserObjectInformationW"),
		uintptr(hProcess),
		uintptr(index),
		uintptr(info),
		uintptr(infoLen))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetUserObjectInformationW *syscall.Proc
