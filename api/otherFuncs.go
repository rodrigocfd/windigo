package api

import (
	"syscall"
	p "winffi/procs"
)

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

func InitCommonControls() {
	syscall.Syscall(p.InitCommonControls.Addr(), 0,
		0, 0, 0)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(p.PostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
}
