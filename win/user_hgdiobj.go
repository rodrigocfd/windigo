package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(index co.COLOR) HBRUSH {
	ret, _, err := syscall.Syscall(proc.GetSysColorBrush.Addr(), 1,
		uintptr(index), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}
