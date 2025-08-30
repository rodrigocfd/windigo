//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [DwmEnableMMCSS] function.
//
// [DwmEnableMMCSS]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmenablemmcss
func DwmEnableMMCSS(enable bool) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_DwmEnableMMCSS, "DwmEnableMMCSS"),
		utl.BoolToUintptr(enable))
	return utl.ErrorAsHResult(ret)
}

var _DwmEnableMMCSS *syscall.Proc

// [DwmFlush] function.
//
// [DwmFlush]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmflush
func DwmFlush() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_DwmFlush, "DwmFlush"))
	return utl.ErrorAsHResult(ret)
}

var _DwmFlush *syscall.Proc

// [DwmGetColorizationColor] function.
//
// [DwmGetColorizationColor]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetcolorizationcolor
func DwmGetColorizationColor() (color COLORREF, isOpaqueBlend bool, hr error) {
	var clr COLORREF
	var bOpaque int32 // BOOL

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_DwmGetColorizationColor, "DwmGetColorizationColor"),
		uintptr(unsafe.Pointer(&clr)),
		uintptr(unsafe.Pointer(&bOpaque)))
	if hr = co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return COLORREF(0), false, hr
	}
	return clr, bOpaque != 0, nil
}

var _DwmGetColorizationColor *syscall.Proc

// [DwmIsCompositionEnabled] function.
//
// [DwmIsCompositionEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmiscompositionenabled
func DwmIsCompositionEnabled() (bool, error) {
	var pfEnabled int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_DwmIsCompositionEnabled, "DwmIsCompositionEnabled"),
		uintptr(unsafe.Pointer(&pfEnabled)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		panic(hr)
	}
	return pfEnabled != 0, nil
}

var _DwmIsCompositionEnabled *syscall.Proc

// [DwmShowContact] function.
//
// [DwmShowContact]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmshowcontact
func DwmShowContact(pointerId uint, showContact co.DWMSC) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.DWMAPI, &_DwmShowContact, "DwmShowContact"),
		uintptr(uint32(pointerId)),
		uintptr(showContact))
	return utl.ErrorAsHResult(ret)
}

var _DwmShowContact *syscall.Proc
