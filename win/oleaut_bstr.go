//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [BSTR] is the string type used in COM Automation.
//
// [BSTR]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/bstr
type BSTR uintptr

// [SysAllocString] function.
//
// ⚠️ You must defer [BSTR.SysFreeString].
//
// # Example
//
//	bstr := win.SysAllocString("hello")
//	defer bstr.SysFreeString()
//
// [SysAllocString]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysallocstring
func SysAllocString(s string) (BSTR, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pS := wbuf.PtrAllowEmpty(s)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLEAUT32, &_SysAllocString, "SysAllocString"),
		uintptr(pS))
	if ret == 0 {
		return BSTR(0), co.HRESULT_E_OUTOFMEMORY
	}
	return BSTR(ret), nil
}

var _SysAllocString *syscall.Proc

// [SysFreeString] function.
//
// [SysFreeString]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysfreestring
func (bstr BSTR) SysFreeString() {
	if bstr != 0 {
		syscall.SyscallN(
			dll.Load(dll.OLEAUT32, &_SysFreeString, "SysFreeString"),
			uintptr(bstr))
	}
}

var _SysFreeString *syscall.Proc

// [SysReAllocString] function.
//
// # Example
//
//	bstr := win.SysAllocString("hello")
//	defer bstr.SysFreeString()
//
//	bstr = bstr.SysReAllocString("another")
//
// [SysReAllocString]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysreallocstring
func (bstr BSTR) SysReAllocString(s string) (BSTR, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pS := wbuf.PtrAllowEmpty(s)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLEAUT32, &_SysReAllocString, "SysReAllocString"),
		uintptr(bstr),
		uintptr(pS))
	if ret == 0 {
		return BSTR(0), co.HRESULT_E_OUTOFMEMORY
	}
	return BSTR(ret), nil
}

var _SysReAllocString *syscall.Proc

// Converts the BSTR pointer to a string.
func (bstr BSTR) String() string {
	return wstr.WinPtrToGo((*uint16)(unsafe.Pointer(bstr)))
}
