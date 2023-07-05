//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to an access token.
type HACCESSTOKEN HANDLE

// [GetCurrentProcessToken] function.
//
// [GetCurrentProcessToken]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocesstoken
func GetCurrentProcessToken() HACCESSTOKEN {
	ret, _, _ := syscall.SyscallN(proc.GetCurrentProcessToken.Addr())
	return HACCESSTOKEN(ret)
}

// [GetCurrentThreadEffectiveToken] function.
//
// [GetCurrentThreadEffectiveToken]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadeffectivetoken
func GetCurrentThreadEffectiveToken() HACCESSTOKEN {
	ret, _, _ := syscall.SyscallN(proc.GetCurrentThreadEffectiveToken.Addr())
	return HACCESSTOKEN(ret)
}

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hToken HACCESSTOKEN) CloseHandle() error {
	ret, _, err := syscall.SyscallN(proc.CloseHandle.Addr(),
		uintptr(hToken))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
