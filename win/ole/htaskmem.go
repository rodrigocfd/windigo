//go:build windows

package ole

import (
	"syscall"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to an OLE [block of memory].
//
// [block of memory]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
type HTASKMEM win.HANDLE

// [CoTaskMemAlloc] function.
//
// ⚠️ You must defer [HTASKMEM.CoTaskMemFree].
//
// # Example
//
//	hMem, _ := ole.CoTaskMemAlloc(uint(unsafe.Sizeof(win.MSG{})))
//	defer hMem.CoTaskMemFree()
//
//	pMsg := (*win.MSG)(unsafe.Pointer(hMem))
//
// [CoTaskMemAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(numBytes uint) (HTASKMEM, error) {
	ret, _, _ := syscall.SyscallN(dllOle(_PROC_CoTaskMemAlloc),
		uintptr(numBytes))
	if ret == 0 {
		return HTASKMEM(0), co.HRESULT_E_OUTOFMEMORY
	}
	return HTASKMEM(ret), nil
}

// [CoTaskMemFree] function.
//
// This method is safe to be called if hMem is zero.
//
// Paired with [CoTaskMemAlloc] and [HTASKMEM.CoTaskMemRealloc].
//
// [CoTaskMemFree]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func (hMem HTASKMEM) CoTaskMemFree() {
	if hMem != 0 {
		syscall.SyscallN(dllOle(_PROC_CoTaskMemFree),
			uintptr(hMem))
	}
}

// [CoTaskMemRealloc] function.
//
// ⚠️ You must defer [HTASKMEM.CoTaskMemFree].
//
// [CoTaskMemRealloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func (hMem HTASKMEM) CoTaskMemRealloc(numBytes uint) (HTASKMEM, error) {
	ret, _, _ := syscall.SyscallN(dllOle(_PROC_CoTaskMemRealloc),
		uintptr(hMem),
		uintptr(numBytes))
	if ret == 0 {
		return HTASKMEM(0), co.HRESULT_E_OUTOFMEMORY
	}
	return HTASKMEM(ret), nil
}
