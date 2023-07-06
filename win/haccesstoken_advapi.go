//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
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

// [GetTokenInformation] function.
//
// # Example
//
// Checking of the current process has elevated privileges:
//
//	hToken, _ := win.GetCurrentProcess().
//		OpenProcessToken(co.TOKEN_QUERY)
//	defer hToken.CloseHandle()
//
//	var elevation win.TOKEN_ELEVATION
//	hToken.GetTokenInformation(
//		co.TOKEN_INFO_Elevation,
//		unsafe.Pointer(&elevation),
//		uint32(unsafe.Sizeof(elevation)),
//	)
//
// [GetTokenInformation]: https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-gettokeninformation
func (hToken HACCESSTOKEN) GetTokenInformation(
	infoClass co.TOKEN_INFO, pInfo unsafe.Pointer, szInfo uint32) error {

	var retLen uint32
	ret, _, err := syscall.SyscallN(proc.GetTokenInformation.Addr(),
		uintptr(hToken), uintptr(infoClass),
		uintptr(pInfo), uintptr(szInfo), uintptr(unsafe.Pointer(&retLen)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
