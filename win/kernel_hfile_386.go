//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// [SetFilePointer] function.
//
// [SetFilePointer]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointer
func (hFile HFILE) SetFilePointerEx(
	distanceToMove int,
	moveMethod co.FILE_FROM,
) (newPointerOffset int, wErr error) {
	var newOff32 int32
	ret, _, err := syscall.SyscallN(procKernel(dll.SetFilePointer),
		uintptr(hFile),
		uintptr(int32(distanceToMove)),
		uintptr(unsafe.Pointer(&newOff32)),
		uintptr(moveMethod))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	} else {
		return int(newOff32), nil
	}
	return
}
