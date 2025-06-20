//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// [SetFilePointerEx] function.
//
// [SetFilePointerEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
func (hFile HFILE) SetFilePointerEx(
	distanceToMove int,
	moveMethod co.FILE_FROM,
) (newPointerOffset int, wErr error) {
	var newOff64 int64
	ret, _, err := syscall.SyscallN(
		dll.Kernel(&_SetFilePointerEx, "SetFilePointerEx"),
		uintptr(hFile),
		uintptr(int64(distanceToMove)),
		uintptr(unsafe.Pointer(&newOff64)),
		uintptr(moveMethod))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	} else {
		return int(newOff64), nil
	}
}

var _SetFilePointerEx *syscall.Proc
