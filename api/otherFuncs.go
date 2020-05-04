package api

import (
	"log"
	"syscall"
	p "winffi/api/proc"
)

func ToUtf16Ptr(s string) *uint16 {
	// We won't return an uintptr right away because it has no pointer semantics,
	// it's just a number, so pointed memory can be garbage-collected.
	// https://stackoverflow.com/a/51188315
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		log.Panicf("toUtf16Ptr failed \"%s\": %s\n", s, err)
	}
	return pstr
}

func ToUtf16PtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return ToUtf16Ptr(s)
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
	syscall.Syscall(p.InitCommonControls.Addr(), 0,
		0, 0, 0)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(p.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}
