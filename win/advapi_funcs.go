//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [RegDisablePredefinedCache] function.
//
// [RegDisablePredefinedCache]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablepredefinedcache
func RegDisablePredefinedCache() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_advapi_RegDisablePredefinedCache, "RegDisablePredefinedCache"))
	return utl.ZeroAsSysError(ret)
}

var _advapi_RegDisablePredefinedCache *syscall.Proc

// [RegDisablePredefinedCacheEx] function.
//
// [RegDisablePredefinedCacheEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablepredefinedcacheex
func RegDisablePredefinedCacheEx() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_advapi_RegDisablePredefinedCacheEx, "RegDisablePredefinedCacheEx"))
	return utl.ZeroAsSysError(ret)
}

var _advapi_RegDisablePredefinedCacheEx *syscall.Proc
