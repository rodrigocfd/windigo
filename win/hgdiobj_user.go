//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [GetSysColorBrush] function.
//
// [GetSysColorBrush]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(index co.COLOR) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.GetSysColorBrush.Addr(),
		uintptr(index))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}
