//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to an [activation context].
//
// [activation context]: https://learn.microsoft.com/en-us/windows/win32/sbscs/activation-contexts
type HACTCTX HANDLE

// [CreateActCtx] function.
//
// ⚠️ You must defer [HACTCTX.ReleaseActCtx].
//
// [CreateActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-createactctxw
func CreateActCtx(actctx *ACTCTX) (HACTCTX, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CreateActCtx, "CreateActCtx"),
		uintptr(unsafe.Pointer(actctx)))

	if int(ret) == utl.INVALID_HANDLE_VALUE {
		return HACTCTX(0), co.ERROR(err)
	}
	return HACTCTX(ret), nil
}

var _CreateActCtx *syscall.Proc

// [GetCurrentActCtx] function.
//
// ⚠️ You must defer [HACTCTX.ReleaseActCtx].
//
// [GetCurrentActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-getcurrentactctx
func GetCurrentActCtx() (HACTCTX, error) {
	var hActCtx HACTCTX
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetCurrentActCtx, "GetCurrentActCtx"),
		uintptr(unsafe.Pointer(hActCtx)))
	if ret == 0 {
		return HACTCTX(0), co.ERROR(err)
	}
	return HACTCTX(ret), nil
}

var _GetCurrentActCtx *syscall.Proc

// [ActivateActCtx] function.
//
// Deactivation is made by [DeactivateActCtx].
//
// [ActivateActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-activateactctx
func (hActCtx HACTCTX) ActivateActCtx() (uint, error) {
	var cookie uint
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ActivateActCtx, "ActivateActCtx"),
		uintptr(unsafe.Pointer(hActCtx)),
		uintptr(unsafe.Pointer(&cookie)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return cookie, nil
}

var _ActivateActCtx *syscall.Proc

// [AddRefActCtx] function.
//
// ⚠️ You must defer [HACTCTX.ReleaseActCtx].
//
// [AddRefActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-addrefactctx
func (hActCtx HACTCTX) AddRefActCtx() HACTCTX {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_AddRefActCtx, "AddRefActCtx"),
		uintptr(hActCtx))
	return HACTCTX(ret)
}

var _AddRefActCtx *syscall.Proc

// [ReleaseActCtx] function.
//
// [ReleaseActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-releaseactctx
func (hActCtx HACTCTX) ReleaseActCtx() {
	syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ReleaseActCtx, "ReleaseActCtx"),
		uintptr(hActCtx))
}

var _ReleaseActCtx *syscall.Proc

// [ZombifyActCtx] function.
//
// [ZombifyActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-zombifyactctx
func (hActCtx HACTCTX) ZombifyActCtx() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ZombifyActCtx, "ZombifyActCtx"),
		uintptr(hActCtx))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ZombifyActCtx *syscall.Proc
