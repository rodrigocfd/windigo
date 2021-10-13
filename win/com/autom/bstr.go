package autom

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
)

// String type used in COM Automation.
//
// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/automat/bstr
type BSTR uintptr

// ⚠️ You must defer BSTR.SysFreeString(), unless you call
// BSTR.SysReAllocString() or pass it to NewVariantBstr().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysallocstring
func SysAllocString(s string) BSTR {
	ret, _, _ := syscall.Syscall(proc.SysAllocString.Addr(), 1,
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(s))), 0, 0)
	if ret == 0 {
		panic("SysAllocString() failed.")
	}
	return BSTR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysfreestring
func (bstr BSTR) SysFreeString() {
	syscall.Syscall(proc.SysFreeString.Addr(), 1,
		uintptr(bstr), 0, 0)
}

// ⚠️ You must defer BSTR.SysFreeString(), unless you call
// BSTR.SysReAllocString() or pass it to NewVariantBstr().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysreallocstring
func (bstr BSTR) SysReAllocString(s string) BSTR {
	ret, _, _ := syscall.Syscall(proc.SysReAllocString.Addr(), 2,
		uintptr(bstr), uintptr(unsafe.Pointer(win.Str.ToNativePtr(s))), 0)
	if ret == 0 {
		panic("SysReAllocString() failed.")
	}
	return BSTR(ret)
}

// Creates a string from the BSTR pointer.
func (bstr BSTR) ToString() string {
	return win.Str.FromNativePtr((*uint16)(unsafe.Pointer(bstr)))
}
