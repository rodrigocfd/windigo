//go:build windows

package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
)

// [BSTR] is the string type used in COM Automation.
//
// [BSTR]: https://docs.microsoft.com/en-us/previous-versions/windows/desktop/automat/bstr
type BSTR uintptr

// [SysAllocString] function.
//
// ⚠️ You must defer BSTR.SysFreeString(), unless you call
// BSTR.SysReAllocString().
//
// [SysAllocString]: https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysallocstring
func SysAllocString(s string) BSTR {
	ret, _, _ := syscall.SyscallN(proc.SysAllocString.Addr(),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(s))))
	if ret == 0 {
		panic("SysAllocString() failed.")
	}
	return BSTR(ret)
}

// [SysFreeString] function.
//
// [SysFreeString]: https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysfreestring
func (bstr BSTR) SysFreeString() {
	syscall.SyscallN(proc.SysFreeString.Addr(),
		uintptr(bstr))
}

// [SysReAllocString] function.
//
// ⚠️ You must defer BSTR.SysFreeString(), unless you call
// BSTR.SysReAllocString().
//
// [SysReAllocString]: https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysreallocstring
func (bstr BSTR) SysReAllocString(s string) BSTR {
	ret, _, _ := syscall.SyscallN(proc.SysReAllocString.Addr(),
		uintptr(bstr), uintptr(unsafe.Pointer(win.Str.ToNativePtr(s))))
	if ret == 0 {
		panic("SysReAllocString() failed.")
	}
	return BSTR(ret)
}

// Converts the BSTR pointer to a string.
func (bstr BSTR) String() string {
	return win.Str.FromNativePtr((*uint16)(unsafe.Pointer(bstr)))
}
