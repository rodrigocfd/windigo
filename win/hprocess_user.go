//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [SetUserObjectInformation] function.
//
// [SetUserObjectInformation]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setuserobjectinformationw
func (hProcess HPROCESS) SetUserObjectInformation(
	index co.UOI,
	info unsafe.Pointer,
	infoLen uintptr) error {

	ret, _, err := syscall.SyscallN(proc.SetUserObjectInformation.Addr(),
		uintptr(hProcess), uintptr(index), uintptr(info), uintptr(infoLen))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
