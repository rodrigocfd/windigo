//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to an OLE [block of memory].
//
// [block of memory]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
type HTASKMEM HANDLE

// [CoTaskMemAlloc] function.
//
// Panics if numBytes is negative.
//
// ⚠️ You must defer [HTASKMEM.CoTaskMemFree].
//
// Example:
//
//	hMem, _ := win.CoTaskMemAlloc(int(unsafe.Sizeof(win.MSG{})))
//	defer hMem.CoTaskMemFree()
//
//	pMsg := (*win.MSG)(unsafe.Pointer(hMem))
//
// [CoTaskMemAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(numBytes int) (HTASKMEM, error) {
	utl.PanicNeg(numBytes)
	ret, _, _ := syscall.SyscallN(
		dll.Ole.Load(&_ole_CoTaskMemAlloc, "CoTaskMemAlloc"),
		uintptr(uint64(numBytes)))
	if ret == 0 {
		return HTASKMEM(0), co.HRESULT_E_OUTOFMEMORY
	}
	return HTASKMEM(ret), nil
}

var _ole_CoTaskMemAlloc *syscall.Proc

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
			dll.Ole.Load(&_ole_CoTaskMemFree, "CoTaskMemFree"),
			uintptr(hMem))
	}
}

var _ole_CoTaskMemFree *syscall.Proc

// [CoTaskMemRealloc] function.
//
// Be careful when using this function. It returns a new [HTASKMEM] handle,
// which invalidates the previous one – that is, you should not call
// [HTASKMEM.CoTaskMemFree] on the previous one. This can become tricky if you
// used defer.
//
// Panics if numBytes is negative.
//
// ⚠️ You must defer [HTASKMEM.CoTaskMemFree].
//
// [CoTaskMemRealloc]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func (hMem HTASKMEM) CoTaskMemRealloc(numBytes int) (HTASKMEM, error) {
	utl.PanicNeg(numBytes)
	ret, _, _ := syscall.SyscallN(
		dll.Ole.Load(&_ole_CoTaskMemRealloc, "CoTaskMemRealloc"),
		uintptr(hMem),
		uintptr(uint64(numBytes)))
	if ret == 0 {
		return HTASKMEM(0), co.HRESULT_E_OUTOFMEMORY
	}
	return HTASKMEM(ret), nil
}

var _ole_CoTaskMemRealloc *syscall.Proc
