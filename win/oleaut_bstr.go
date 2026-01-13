//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/wstr"
)

// [BSTR] is the string type used in COM Automation.
//
// [BSTR]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/bstr
type BSTR uintptr

// [SysAllocString] function.
//
// ⚠️ You must defer [BSTR.SysFreeString].
//
// Example:
//
//	bstr, _ := win.SysAllocString("hello")
//	defer bstr.SysFreeString()
//
// [SysAllocString]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysallocstring
func SysAllocString(s string) (BSTR, error) {
	var wS wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Oleaut.Load(&_oleaut_SysAllocString, "SysAllocString"),
		uintptr(wS.AllowEmpty(s)))
	if ret == 0 {
		return BSTR(0), co.HRESULT_E_OUTOFMEMORY
	}
	return BSTR(ret), nil
}

var _oleaut_SysAllocString *syscall.Proc

// [SysFreeString] function.
//
// [SysFreeString]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysfreestring
func (bstr BSTR) SysFreeString() {
	if bstr != 0 {
		syscall.SyscallN(
			dll.Oleaut.Load(&_oleaut_SysFreeString, "SysFreeString"),
			uintptr(bstr))
	}
}

var _oleaut_SysFreeString *syscall.Proc

// [SysReAllocString] function.
//
// Be careful when using this function. It returns a new [BSTR] handle, which
// invalidates the previous one – that is, you should not call
// [BSTR.SysFreeString] on the previous one. This can become tricky if you
// used defer.
//
// ⚠️ You must defer [BSTR.SysFreeString].
//
// [SysReAllocString]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysreallocstring
func (bstr BSTR) SysReAllocString(s string) (BSTR, error) {
	var wS wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		dll.Oleaut.Load(&_oleaut_SysReAllocString, "SysReAllocString"),
		uintptr(bstr),
		uintptr(wS.AllowEmpty(s)))
	if ret == 0 {
		return BSTR(0), co.HRESULT_E_OUTOFMEMORY
	}
	return BSTR(ret), nil
}

var _oleaut_SysReAllocString *syscall.Proc

// Converts the BSTR pointer to a string.
func (bstr BSTR) String() string {
	return wstr.DecodePtr((*uint16)(unsafe.Pointer(bstr)))
}
