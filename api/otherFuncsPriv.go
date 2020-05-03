package api

import (
	"syscall"
	"unsafe"
)

func toUtf16ToUintptr(s string) uintptr {
	pstr, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(pstr))
}

func toUtf16BlankIsNilToUintptr(s string) uintptr {
	pstr := uintptr(0)
	if s != "" {
		pstr = toUtf16ToUintptr(s)
	}
	return pstr
}

func boolToUintptr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

func hiWord(value uint32) uint16 {
	return uint16(value >> 16 & 0xffff)
}

func loWord(value uint32) uint16 {
	return uint16(value)
}

func hiByte(value uint16) uint8 {
	return uint8(value >> 8 & 0xff)
}

func loByte(value uint16) uint8 {
	return uint8(value)
}
