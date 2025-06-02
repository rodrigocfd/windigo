//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// [GetPerformanceInfo] function.
//
// [GetPerformanceInfo]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getperformanceinfo
func GetPerformanceInfo() (PERFORMANCE_INFORMATION, error) {
	var pi PERFORMANCE_INFORMATION
	pi.SetCb()

	ret, _, err := syscall.SyscallN(_K32GetPerformanceInfo.Addr(),
		uintptr(unsafe.Pointer(&pi)),
		uintptr(unsafe.Sizeof(pi)))
	if ret == 0 {
		return PERFORMANCE_INFORMATION{}, co.ERROR(err)
	}
	return pi, nil
}

var _K32GetPerformanceInfo = dll.Psapi.NewProc("K32GetPerformanceInfo")
