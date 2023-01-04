//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmenablemmcss
func DwmEnableMMCSS(enable bool) error {
	ret, _, _ := syscall.Syscall(proc.DwmEnableMMCSS.Addr(), 1,
		util.BoolToUintptr(enable), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmflush
func DwmFlush() error {
	ret, _, _ := syscall.Syscall(proc.DwmFlush.Addr(), 0,
		0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetcolorizationcolor
func DwmGetColorizationColor() (color COLORREF, isOpaqueBlend bool) {
	bOpaqueBlend := int32(util.BoolToUintptr(isOpaqueBlend)) // BOOL
	ret, _, _ := syscall.Syscall(proc.DwmGetColorizationColor.Addr(), 2,
		uintptr(unsafe.Pointer(&color)), uintptr(unsafe.Pointer(&bOpaqueBlend)),
		0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	isOpaqueBlend = bOpaqueBlend != 0
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmiscompositionenabled
func DwmIsCompositionEnabled() bool {
	var pfEnabled int32 // BOOL
	ret, _, _ := syscall.Syscall(proc.DwmIsCompositionEnabled.Addr(), 1,
		uintptr(unsafe.Pointer(&pfEnabled)), 0, 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return pfEnabled != 0
}
