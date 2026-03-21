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

// BOOL AdjustTokenPrivileges(
//[in]            HANDLE            TokenHandle,
//[in]            BOOL              DisableAllPrivileges,
//[in, optional]  PTOKEN_PRIVILEGES NewState,
//[in]            DWORD             BufferLength,
//[out, optional] PTOKEN_PRIVILEGES PreviousState,
//[out, optional] PDWORD            ReturnLength

type TOKEN_PRIVILEGES struct {
	PrivilegeCount uint32
	Privileges     [1]co.LUID_AND_ATTRIBUTES
}

// AllPrivileges returns a slice that can be used to iterate over the privileges in p.
func (p *TOKEN_PRIVILEGES) AllPrivileges() []co.LUID_AND_ATTRIBUTES {
	return (*[(1 << 27) - 1]co.LUID_AND_ATTRIBUTES)(unsafe.Pointer(&p.Privileges[0]))[:p.PrivilegeCount:p.PrivilegeCount]
}

func (hToken HTOKEN) AdjustTokenPrivilege(luid co.LUID, attribute uint32) (TOKEN_PRIVILEGES, error) {
	privilege := co.LUID_AND_ATTRIBUTES{luid, attribute}
	in := TOKEN_PRIVILEGES{PrivilegeCount: 1, Privileges: [1]co.LUID_AND_ATTRIBUTES{privilege}}
	out := TOKEN_PRIVILEGES{}
	var outsize uint32
	ret, _, err := syscall.SyscallN(
		dll.Advapi.Load(&_advapi_AdjustTokenPrivileges, "AdjustTokenPrivileges"),
		uintptr(hToken),
		uintptr(0), // do NOT disable all privileges
		uintptr(unsafe.Pointer(&in)),
		uintptr(unsafe.Sizeof(out)),
		uintptr(unsafe.Pointer(&out)),
		uintptr(unsafe.Pointer(&outsize)))
	return out, utl.ZeroAsGetLastError(ret, err) // XXX check error handling
}

var _advapi_AdjustTokenPrivileges *syscall.Proc
