//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
)

// [RegDisablePredefinedCache] function.
//
// [RegDisablePredefinedCache]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablepredefinedcache
func RegDisablePredefinedCache() error {
	ret, _, _ := syscall.SyscallN(_RegDisablePredefinedCache.Addr())
	return util.ZeroAsSysError(ret)
}

var _RegDisablePredefinedCache = dll.Advapi32.NewProc("RegDisablePredefinedCache")

// [RegDisablePredefinedCacheEx] function.
//
// [RegDisablePredefinedCacheEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablepredefinedcacheex
func RegDisablePredefinedCacheEx() error {
	ret, _, _ := syscall.SyscallN(_RegDisablePredefinedCacheEx.Addr())
	return util.ZeroAsSysError(ret)
}

var _RegDisablePredefinedCacheEx = dll.Advapi32.NewProc("RegDisablePredefinedCacheEx")
