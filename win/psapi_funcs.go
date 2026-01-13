//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
)

// [GetPerformanceInfo] function.
//
// [GetPerformanceInfo]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getperformanceinfo
func GetPerformanceInfo() (PERFORMANCE_INFORMATION, error) {
	var pi PERFORMANCE_INFORMATION
	pi.SetCb()

	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_K32GetPerformanceInfo, "K32GetPerformanceInfo"),
		uintptr(unsafe.Pointer(&pi)),
		uintptr(unsafe.Sizeof(pi)))
	if ret == 0 {
		return PERFORMANCE_INFORMATION{}, co.ERROR(err)
	}
	return pi, nil
}

var _psapi_K32GetPerformanceInfo *syscall.Proc
