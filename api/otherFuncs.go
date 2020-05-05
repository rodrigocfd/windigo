package api

import (
	"fmt"
	"syscall"
	"unsafe"
	"winffi/api/proc"
	c "winffi/consts"
)

// Wrapper to syscall.UTF16PtrFromString(), panics in error.
func StrToUtf16Ptr(s string) *uint16 {
	// We won't return an uintptr right away because it has no pointer semantics,
	// it's just a number, so pointed memory can be garbage-collected.
	// https://stackoverflow.com/a/51188315
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(fmt.Sprintf("ToUtf16Ptr failed \"%s\": %s\n",
			s, err))
	}
	return pstr
}

// Wrapper to syscall.UTF16PtrFromString(), panics in error. A blank string will
// return a null pointer.
func StrToUtf16PtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return StrToUtf16Ptr(s)
	}
	return nil
}

func boolToUintptr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

//------------------------------------------------------------------------------

func HiWord(value uint32) uint16 {
	return uint16(value >> 16 & 0xffff)
}

func LoWord(value uint32) uint16 {
	return uint16(value)
}

func HiByte(value uint16) uint8 {
	return uint8(value >> 8 & 0xff)
}

func LoByte(value uint16) uint8 {
	return uint8(value)
}

//------------------------------------------------------------------------------

func InitCommonControls() {
	syscall.Syscall(proc.InitCommonControls.Addr(), 0,
		0, 0, 0)
}

func GetSystemMetrics(index c.SM) int32 {
	ret, _, _ := syscall.Syscall(proc.GetSystemMetrics.Addr(), 1,
		uintptr(index), 0, 0)
	return int32(ret)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(proc.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}

func SystemParametersInfo(action c.SPI, param uint32,
	pvParam unsafe.Pointer, winIni uint32) {

	ret, _, errno := syscall.Syscall6(proc.SystemParametersInfo.Addr(), 4,
		uintptr(action), uintptr(param), uintptr(pvParam), uintptr(winIni),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SystemParametersInfo failed: %d %s\n",
			errno, errno.Error()))
	}
}
