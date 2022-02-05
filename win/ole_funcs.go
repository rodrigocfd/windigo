package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Loads the COM module. This needs to be done only once in your application.
// Typically uses COINIT_APARTMENTTHREADED.
//
// ⚠️ You must defer CoUninitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(coInit), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK && hr != errco.S_FALSE {
		panic(hr)
	}
}

// ⚠️ You must defer CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(size int) uintptr {
	ret, _, _ := syscall.Syscall(proc.CoTaskMemAlloc.Addr(), 1,
		uintptr(size), 0, 0)
	if ret == 0 {
		panic("CoTaskMemAlloc() failed.")
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func CoTaskMemFree(pv uintptr) {
	syscall.Syscall(proc.CoTaskMemFree.Addr(), 1,
		pv, 0, 0)
}

// ⚠️ You must defer CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func CoTaskMemRealloc(pv uintptr, size int) uintptr {
	ret, _, _ := syscall.Syscall(proc.CoTaskMemRealloc.Addr(), 2,
		pv, uintptr(size), 0)
	if ret == 0 {
		panic("CoTaskMemRealloc() failed.")
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
