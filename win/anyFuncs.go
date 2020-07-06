/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// Returns *uint16.StrToPtr
// Wrapper to syscall.UTF16PtrFromString().
// Panics on error.
func StrToPtr(s string) *uint16 {
	// We won't return an uintptr right away because it has no pointer semantics,
	// it's just a number, so pointed memory can be garbage-collected.
	// https://stackoverflow.com/a/51188315
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("ToUtf16Ptr failed \"%s\": %s",
			s, err))
	}
	return pstr
}

// Returns *uint16, or nil of empty string.
// Wrapper to syscall.UTF16PtrFromString().
// Panics on error.
func StrToPtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return StrToPtr(s)
	}
	return nil
}

// Simple conversion for syscalls.
func boolToUintptr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

func hiWord(value uint32) uint16 { return uint16(value >> 16 & 0xffff) }
func loWord(value uint32) uint16 { return uint16(value) }
func hiByte(value uint16) uint8  { return uint8(value >> 8 & 0xff) }
func loByte(value uint16) uint8  { return uint8(value) }

//------------------------------------------------------------------------------

// Don't forget to call CoUninitialize().
func CoInitializeEx(dwCoInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(dwCoInit), 0)
	if co.ERROR(ret) != co.ERROR_S_OK && co.ERROR(ret) != co.ERROR_S_FALSE {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("CoInitializeEx failed: %d %s",
			lerr, lerr.Error()))
	}
}

func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}

func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0, 0, 0, 0)
}

func GetCursorPos() POINT {
	ptBuf := POINT{}
	ret, _, lerr := syscall.Syscall(proc.GetCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(&ptBuf)), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetCursorPos failed: %d %s",
			lerr, lerr.Error()))
	}
	return ptBuf
}

// Available in Windows 10, version 1607.
func GetDpiForSystem() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetDpiForSystem.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

// Multiplies two 32-bit values and then divides the 64-bit result by a third
// 32-bit value. The final result is rounded to the nearest integer.
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := syscall.Syscall(proc.MulDiv.Addr(), 3,
		uintptr(number), uintptr(numerator), uintptr(denominator))
	return int32(ret)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}

// Available in Windows 10, version 1703.
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) {
	ret, _, lerr := syscall.Syscall(proc.SetProcessDpiAwarenessContext.Addr(), 1,
		uintptr(value), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetProcessDpiAwarenessContext failed: %d %s",
			lerr, lerr.Error()))
	}
}

// Available in Windows Vista.
func SetProcessDPIAware() {
	ret, _, _ := syscall.Syscall(proc.SetProcessDPIAware.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic("SetProcessDPIAware failed.")
	}
}

func SystemParametersInfo(uiAction co.SPI, uiParam uint32,
	pvParam unsafe.Pointer, fWinIni uint32) {

	ret, _, lerr := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SystemParametersInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}