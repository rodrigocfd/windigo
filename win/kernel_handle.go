//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to an [internal object]. This generic handle is used throughout the
// whole API, with different meanings.
//
// [internal object]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HANDLE syscall.Handle

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (h HANDLE) CloseHandle() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_CloseHandle, "CloseHandle"),
		uintptr(h))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_CloseHandle *syscall.Proc
