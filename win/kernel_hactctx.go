//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
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
		dll.Kernel.Load(&_kernel_CreateActCtx, "CreateActCtx"),
		uintptr(unsafe.Pointer(actctx)))

	if int(ret) == utl.INVALID_HANDLE_VALUE {
		return HACTCTX(0), co.ERROR(err)
	}
	return HACTCTX(ret), nil
}

var _kernel_CreateActCtx *syscall.Proc

// [GetCurrentActCtx] function.
//
// ⚠️ You must defer [HACTCTX.ReleaseActCtx].
//
// [GetCurrentActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-getcurrentactctx
func GetCurrentActCtx() (HACTCTX, error) {
	var hActCtx HACTCTX
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetCurrentActCtx, "GetCurrentActCtx"),
		uintptr(unsafe.Pointer(hActCtx)))
	if ret == 0 {
		return HACTCTX(0), co.ERROR(err)
	}
	return HACTCTX(ret), nil
}

var _kernel_GetCurrentActCtx *syscall.Proc

// [ActivateActCtx] function.
//
// Deactivation is made by [DeactivateActCtx].
//
// [ActivateActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-activateactctx
func (hActCtx HACTCTX) ActivateActCtx() (int, error) {
	var cookie uintptr
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ActivateActCtx, "ActivateActCtx"),
		uintptr(unsafe.Pointer(hActCtx)),
		uintptr(unsafe.Pointer(&cookie)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(cookie), nil
}

var _kernel_ActivateActCtx *syscall.Proc

// [AddRefActCtx] function.
//
// ⚠️ You must defer [HACTCTX.ReleaseActCtx].
//
// [AddRefActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-addrefactctx
func (hActCtx HACTCTX) AddRefActCtx() HACTCTX {
	ret, _, _ := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_AddRefActCtx, "AddRefActCtx"),
		uintptr(hActCtx))
	return HACTCTX(ret)
}

var _kernel_AddRefActCtx *syscall.Proc

// [ReleaseActCtx] function.
//
// [ReleaseActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-releaseactctx
func (hActCtx HACTCTX) ReleaseActCtx() {
	syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ReleaseActCtx, "ReleaseActCtx"),
		uintptr(hActCtx))
}

var _kernel_ReleaseActCtx *syscall.Proc

// [ZombifyActCtx] function.
//
// [ZombifyActCtx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-zombifyactctx
func (hActCtx HACTCTX) ZombifyActCtx() error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ZombifyActCtx, "ZombifyActCtx"),
		uintptr(hActCtx))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_ZombifyActCtx *syscall.Proc
