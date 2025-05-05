//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// [DwmEnableMMCSS] function.
//
// [DwmEnableMMCSS]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmenablemmcss
func DwmEnableMMCSS(enable bool) error {
	ret, _, _ := syscall.SyscallN(_DwmEnableMMCSS.Addr(),
		util.BoolToUintptr(enable))
	return util.ErrorAsHResult(ret)
}

var _DwmEnableMMCSS = dll.Dwmapi.NewProc("DwmEnableMMCSS")

// [DwmFlush] function.
//
// [DwmFlush]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmflush
func DwmFlush() error {
	ret, _, _ := syscall.SyscallN(_DwmFlush.Addr())
	return util.ErrorAsHResult(ret)
}

var _DwmFlush = dll.Dwmapi.NewProc("DwmFlush")

// [DwmGetColorizationColor] function.
//
// [DwmGetColorizationColor]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetcolorizationcolor
func DwmGetColorizationColor() (color COLORREF, isOpaqueBlend bool, hr error) {
	var clr COLORREF
	var bOpaque int32 // BOOL

	ret, _, _ := syscall.SyscallN(_DwmGetColorizationColor.Addr(),
		uintptr(unsafe.Pointer(&clr)), uintptr(unsafe.Pointer(&bOpaque)))
	if hr = co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return COLORREF(0), false, hr
	}
	return clr, bOpaque != 0, nil
}

var _DwmGetColorizationColor = dll.Dwmapi.NewProc("DwmGetColorizationColor")

// [DwmIsCompositionEnabled] function.
//
// [DwmIsCompositionEnabled]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmiscompositionenabled
func DwmIsCompositionEnabled() (bool, error) {
	var pfEnabled int32 // BOOL
	ret, _, _ := syscall.SyscallN(_DwmIsCompositionEnabled.Addr(),
		uintptr(unsafe.Pointer(&pfEnabled)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		panic(hr)
	}
	return pfEnabled != 0, nil
}

var _DwmIsCompositionEnabled = dll.Dwmapi.NewProc("DwmIsCompositionEnabled")
