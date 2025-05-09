//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
)

// A [handle] to an object. This generic handle is used throughout the whole
// API, with different meanings.
//
// [handle]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HANDLE syscall.Handle

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (h HANDLE) CloseHandle() error {
	ret, _, err := syscall.SyscallN(_CloseHandle.Addr(),
		uintptr(h))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _CloseHandle = dll.Kernel32.NewProc("CloseHandle")
