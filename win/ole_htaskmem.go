//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
)

// Handle to an OLE [block of memory].
//
// [block of memory]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
type HTASKMEM HANDLE

// [CoTaskMemAlloc] function.
//
// ⚠️ You must defer [HTASKMEM.CoTaskMemFree].
//
// Example:
//
//	hMem, _ := win.CoTaskMemAlloc(uint(unsafe.Sizeof(win.MSG{})))
//	defer hMem.CoTaskMemFree()
//
//	pMsg := (*win.MSG)(unsafe.Pointer(hMem))
//
// [CoTaskMemAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(numBytes uint) (HTASKMEM, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_CoTaskMemAlloc, "CoTaskMemAlloc"),
		uintptr(numBytes))
	if ret == 0 {
		return HTASKMEM(0), co.HRESULT_E_OUTOFMEMORY
	}
	return HTASKMEM(ret), nil
}

var _CoTaskMemAlloc *syscall.Proc

// [CoTaskMemFree] function.
//
// This method is safe to be called if hMem is zero.
//
// Paired with [CoTaskMemAlloc] and [HTASKMEM.CoTaskMemRealloc].
//
// [CoTaskMemFree]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func (hMem HTASKMEM) CoTaskMemFree() {
	if hMem != 0 {
		syscall.SyscallN(
			dll.Load(dll.OLE32, &_CoTaskMemFree, "CoTaskMemFree"),
			uintptr(hMem))
	}
}

var _CoTaskMemFree *syscall.Proc

// [CoTaskMemRealloc] function.
//
// Be careful when using this function. It returns a new [HTASKMEM] handle,
// which invalidates the previous one – that is, you should not call
// [HTASKMEM.CoTaskMemFree] on the previous one. This can become tricky if you
// used defer.
//
// ⚠️ You must defer [HTASKMEM.CoTaskMemFree].
//
// [CoTaskMemRealloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func (hMem HTASKMEM) CoTaskMemRealloc(numBytes uint) (HTASKMEM, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_CoTaskMemRealloc, "CoTaskMemRealloc"),
		uintptr(hMem),
		uintptr(numBytes))
	if ret == 0 {
		return HTASKMEM(0), co.HRESULT_E_OUTOFMEMORY
	}
	return HTASKMEM(ret), nil
}

var _CoTaskMemRealloc *syscall.Proc
