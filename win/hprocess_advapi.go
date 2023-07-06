//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [OpenProcessToken] function.
//
// ⚠️ You must defer HACCESSTOKEN.CloseHandle().
//
// # Example
//
//	hProcess := GetCurrentProcess()
//
//	hToken, _ := hProcess.OpenProcessToken(co.TOKEN_EXECUTE)
//	defer hToken.CloseHandle()
//
// [OpenProcessToken]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocesstoken
func (hProcess HPROCESS) OpenProcessToken(
	desiredAccess co.TOKEN) (HACCESSTOKEN, error) {

	var hToken HACCESSTOKEN
	ret, _, err := syscall.SyscallN(proc.OpenProcessToken.Addr(),
		uintptr(hProcess), uintptr(desiredAccess),
		uintptr(unsafe.Pointer(&hToken)))
	if ret == 0 {
		return HACCESSTOKEN(0), errco.ERROR(err)
	}
	return hToken, nil
}
