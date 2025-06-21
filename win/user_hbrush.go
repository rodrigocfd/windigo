//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// [GetSysColorBrush] function.
//
// [GetSysColorBrush]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(index co.COLOR) (HBRUSH, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.USER32, &_GetSysColorBrush, "GetSysColorBrush"),
		uintptr(index))
	if ret == 0 {
		return HBRUSH(0), co.ERROR_INVALID_PARAMETER
	}
	return HBRUSH(ret), nil
}

var _GetSysColorBrush *syscall.Proc
