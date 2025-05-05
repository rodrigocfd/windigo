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
	ret, _, err := syscall.SyscallN(_SetFilePointer.Addr(),
		uintptr(hFile), uintptr(distanceToMove),
		uintptr(unsafe.Pointer(&newPointerOffset)), uintptr(moveMethod))

	if wErr = co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		newPointerOffset = 0
	} else {
		wErr = nil
	}
	return
}

var _SetFilePointer = dll.Kernel32.NewProc("SetFilePointer")
