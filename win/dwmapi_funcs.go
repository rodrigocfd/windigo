//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [DwmEnableMMCSS] function.
//
// [DwmEnableMMCSS]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmenablemmcss
func DwmEnableMMCSS(enable bool) error {
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmEnableMMCSS),
		utl.BoolToUintptr(enable))
	return utl.ErrorAsHResult(ret)
}

// [DwmFlush] function.
//
// [DwmFlush]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmflush
func DwmFlush() error {
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmFlush))
	return utl.ErrorAsHResult(ret)
}

// [DwmGetColorizationColor] function.
//
// [DwmGetColorizationColor]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetcolorizationcolor
func DwmGetColorizationColor() (color COLORREF, isOpaqueBlend bool, hr error) {
	var clr COLORREF
	var bOpaque int32 // BOOL

	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmGetColorizationColor),
		uintptr(unsafe.Pointer(&clr)),
		uintptr(unsafe.Pointer(&bOpaque)))
	if hr = co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return COLORREF(0), false, hr
	}
	return clr, bOpaque != 0, nil
}

// [DwmIsCompositionEnabled] function.
//
// [DwmIsCompositionEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmiscompositionenabled
func DwmIsCompositionEnabled() (bool, error) {
	var pfEnabled int32 // BOOL
	ret, _, _ := syscall.SyscallN(dll.Dwmapi(dll.PROC_DwmIsCompositionEnabled),
		uintptr(unsafe.Pointer(&pfEnabled)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		panic(hr)
	}
	return pfEnabled != 0, nil
}
