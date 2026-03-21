//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [RegDisablePredefinedCache] function.
//
// [RegDisablePredefinedCache]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablepredefinedcache
func RegDisablePredefinedCache() error {
	ret, _, _ := syscall.SyscallN(
		dll.Advapi.Load(&_advapi_RegDisablePredefinedCache, "RegDisablePredefinedCache"))
	return utl.ZeroAsSysError(ret)
}

var _advapi_RegDisablePredefinedCache *syscall.Proc

// [RegDisablePredefinedCacheEx] function.
//
// [RegDisablePredefinedCacheEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablepredefinedcacheex
func RegDisablePredefinedCacheEx() error {
	ret, _, _ := syscall.SyscallN(
		dll.Advapi.Load(&_advapi_RegDisablePredefinedCacheEx, "RegDisablePredefinedCacheEx"))
	return utl.ZeroAsSysError(ret)
}

var _advapi_RegDisablePredefinedCacheEx *syscall.Proc

// [LookupPrivilegeValue] function.
//
// [LookupPrivilegeValue]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-lookupprivilegevaluew
func LookupPrivilegeValue(systemname string, name string) (co.LUID, error) {
	var wsystemname wstr.BufEncoder
	var wname wstr.BufEncoder
	var result co.LUID

	ret, _, err := syscall.SyscallN(
		dll.Advapi.Load(&_advapi_LookupPrivilegeValueW, "LookupPrivilegeValueW"),
		uintptr(wsystemname.EmptyIsNil(systemname)),
		uintptr(wname.AllowEmpty(name)),
		uintptr(unsafe.Pointer(&result)))
	return result, utl.ZeroAsGetLastError(ret, err)
}

var _advapi_LookupPrivilegeValueW *syscall.Proc
